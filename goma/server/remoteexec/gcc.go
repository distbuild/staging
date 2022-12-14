// Copyright 2018 The Goma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package remoteexec

import (
	"errors"
	"fmt"
	"strings"
)

// longest first
var pathFlags = []string{
	"-fcoverage-compilation-dir=",
	"-fcrash-diagnostics-dir=",
	"-fdebug-compilation-dir=",
	"-ffile-compilation-dir=",
	"-fprofile-sample-use=",
	"-fsanitize-blacklist=",
	"-fprofile-instr-use=",
	"-fprofile-list=",
	"-resource-dir=",
	"--include=",
	"--sysroot=",
	"--include",
	"-include=",
	"-include",
	"-isystem",
	"-o",
	"-B",
	"-F",
	"-I",
}

// TODO: share exec/gcc.go ?

// gccRelocatableReq checks if the request (args, envs) uses relative
// paths only and doesn't use flags that generates output including cwd,
// so will generate cwd-agnostic outputs
// (files/stdout/stderr will not include cwd dependent paths).
//
// The request will be relocatable if path used in arg is cwd relative.
//
// The request will NOT be relocatable, that is, generate
// outputs that would contain absolute path names (DW_AT_comp_dir etc),
// if
//  debug build (-g* except -g0) -> DW_AT_comp_dir or other filepaths.
//    this will be canceled by -fdebug-compilation-dir
//  --pnacl-allow-translate  crbug.com/685461
//
// The following flags would NOT be relocatable
//  absolute input filename (debug build)
//      *.d file output will not be cwd-agnostic.
//      DW_AT_name (debug build)
//  -I<path>
//      *.d file output will not be cwd-agnostic.
//      directory table in debug info (debug build)
//  -B<path>
//  -isystem<path> --sysroot=<path>
//  ...
//  TODO: these could be normalized to cwd relative?
//
// ref:
// https://docs.google.com/spreadsheets/d/1_-ZJhqy7WhSFYuZU2QkmQ4Ed9182bWfKg09EfBAkVf8/edit#gid=759603323
//
// TODO: http://b/150662978 relocatableReq should check input and output file path too.
func gccRelocatableReq(filepath clientFilePath, args, envs []string) error {
	var debugFlags []string
	debugCompilationDir := false
	subArgs := map[string][]string{}
	pathFlag := false
	var subCmd string
Loop:
	for _, arg := range args {
		if pathFlag {
			if filepath.IsAbs(arg) {
				return fmt.Errorf("abs path: %s", arg)
			}
			// TODO: When clang supports relative paths in hmap,
			// instead check that hmap does not have abs paths.
			if strings.HasSuffix(arg, ".hmap") {
				return fmt.Errorf("hmap file: %s", arg)
			}
			pathFlag = false
			continue
		}
		// TODO: When clang supports relative paths in hmap,
		// instead check that hmap does not have abs paths.
		if strings.HasPrefix(arg, "-I") && strings.HasSuffix(arg, ".hmap") {
			return fmt.Errorf("hmap file: %s", arg)
		}
		for _, fp := range pathFlags {
			if arg != fp && strings.HasPrefix(arg, fp) {
				if filepath.IsAbs(arg[len(fp):]) {
					return fmt.Errorf("abs path: %s", arg)
				}
				if fp == "-fdebug-compilation-dir=" || fp == "-ffile-compilation-dir=" {
					switch subCmd {
					case "clang":
						subArgs[subCmd] = append(subArgs[subCmd], arg)
						fallthrough
					case "":
						debugCompilationDir = true
						continue Loop
					}
					return errors.New("-fdebug-compilation-dir= and -ffile-compilation-dir= not supported for " + subCmd)
				}
				continue Loop
			}
		}
		switch {
		case arg == "-fdebug-compilation-dir":
			// We can stop checking the rest of the flags.
			// When seeing "-fdebug-compilation-dir",
			// we could cancel non-cwd agnosticsness due to
			// debug flags.
			//
			// Note that this check applies to both GCC and Clang
			// -xx -fdebug-compilation-dir . -yy ...                 <- GCC flag
			// -xx -Xclang -fdebug-compilation-dir -Xclang . -yy ... <- Clang flag
			//
			// As a result, clangArgRelocatableReq() doesn't need to check this again.
			// The value of -fdebug-compilation-dir is used
			// just for DW_AT_comp_dir, so no need to check it.
			switch subCmd {
			case "clang":
				subArgs[subCmd] = append(subArgs[subCmd], arg)
				fallthrough
			case "":
				debugCompilationDir = true

				continue Loop
			}
			return errors.New("fdebug-compilation-dir not supported for " + subCmd)

		case subCmd != "":
			subArgs[subCmd] = append(subArgs[subCmd], arg)
			subCmd = ""

		case strings.HasPrefix(arg, "-g"):
			if arg == "-g0" {
				debugFlags = nil
				continue
			}
			debugFlags = append(debugFlags, arg)
		case arg == "--pnacl-allow-translate": // crbug.com/685461
			return errors.New("pnacl-allow-translate")

		case strings.HasPrefix(arg, "-Wa,"): // assembler arg
			subArgs["as"] = append(subArgs["as"], strings.Split(arg[len("-Wa,"):], ",")...)
		case strings.HasPrefix(arg, "-Wl,"): // linker arg
			subArgs["ld"] = append(subArgs["ld"], strings.Split(arg[len("-Wl,"):], ",")...)
		case strings.HasPrefix(arg, "-Wp,"): // preproc arg
			subArgs["cpp"] = append(subArgs["cpp"], strings.Split(arg[len("-Wp,"):], ",")...)
		case arg == "-Xclang":
			subCmd = "clang"

		case arg == "-mllvm":
			// -mllvm <value>  Additional arguments to forward to LLVM's option processing
			subCmd = "llvm"

		case strings.HasPrefix(arg, "-w"): // inhibit all warnings
		case strings.HasPrefix(arg, "-W"): // warning
		case strings.HasPrefix(arg, "-D"): // define
		case strings.HasPrefix(arg, "-U"): // undefine
		case strings.HasPrefix(arg, "-O"): // optimize
		case strings.HasPrefix(arg, "-f"): // feature
		case strings.HasPrefix(arg, "-m"):
			// -m64, -march=x86-64
		case arg == "-arch":
		case arg == "-target":
		case strings.HasPrefix(arg, "--target="):

		case strings.HasPrefix(arg, "-no"):
			// -no-canonical-prefixes, -nostdinc++
		case arg == "-integrated-as":
		case arg == "-pedantic":
		case arg == "-pipe":
		case arg == "-pie":
		case arg == "-pthread":
		case arg == "-c":
		case strings.HasPrefix(arg, "-std"):
		case strings.HasPrefix(arg, "--param="):
		case arg == "-MMD" || arg == "-MD" || arg == "-M":
		case arg == "-Qunused-arguments":
		case arg == "-static-libgcc":
		case strings.HasPrefix(arg, "--rtlib="):
			continue

		case arg == "-o":
			pathFlag = true
		case arg == "-I" || arg == "-B" || arg == "-F" || arg == "-isystem" || arg == "-include" || arg == "-iframework":
			pathFlag = true
		case arg == "-MF":
			pathFlag = true
		case arg == "-isysroot":
			pathFlag = true
		case arg == "--sysroot":
			pathFlag = true
		case arg == "-idirafter":
			pathFlag = true

		case strings.HasPrefix(arg, "-"): // unknown flag?
			return unknownFlagError{arg: arg}

		default: // input file?
			if filepath.IsAbs(arg) {
				return fmt.Errorf("abs path: %s", arg)
			}
		}
	}

	if len(debugFlags) > 0 && !debugCompilationDir {
		return fmt.Errorf("debug build: %q", debugFlags)
	}
	if len(subArgs) > 0 {
		for cmd, args := range subArgs {
			switch cmd {
			case "clang":
				err := clangArgRelocatable(filepath, args)
				if err != nil {
					return err
				}
			case "llvm":
				err := llvmArgRelocatable(filepath, args)
				if err != nil {
					return err
				}
			case "as":
				err := asArgRelocatable(filepath, args)
				if err != nil {
					return err
				}
			default:
				return unknownFlagError{arg: fmt.Sprintf("unsupported subcommand %s: %s", cmd, args)}
			}
		}
	}

	for _, env := range envs {
		e := strings.SplitN(env, "=", 2)
		if len(e) != 2 {
			return fmt.Errorf("bad environment variable: %s", env)
		}
		if e[0] == "PWD" {
			continue
		}
		if filepath.IsAbs(e[1]) {
			return fmt.Errorf("abs path in env %s=%s", e[0], e[1])
		}
	}
	return nil
}

func clangArgRelocatable(filepath clientFilePath, args []string) error {
	pathFlag := false
	skipFlag := false
	for _, arg := range args {
		switch {
		case pathFlag:
			if filepath.IsAbs(arg) {
				return fmt.Errorf("clang abs path: %s", arg)
			}
			pathFlag = false
		case skipFlag:
			skipFlag = false

		case arg == "-mllvm", arg == "-add-plugin", arg == "-fdebug-compilation-dir", arg == "-target-feature":
			// TODO: pass llvmArgRelocatable for -mllvm?
			skipFlag = true
		case strings.HasPrefix(arg, "-plugin-arg-"):
			skipFlag = true
		case arg == "-load":
			pathFlag = true
		case strings.HasPrefix(arg, "-f"): // feature
		case strings.HasPrefix(arg, "-debug-info-kind="):
		case arg == "-no-opaque-pointers":
		default:
			return unknownFlagError{arg: fmt.Sprintf("clang: %s", arg)}
		}
	}
	return nil
}

func llvmArgRelocatable(filepath clientFilePath, args []string) error {
	for _, arg := range args {
		switch {
		case strings.HasPrefix(arg, "-asan-"):
			// https://b/issues/141210713#comment3
			// -mllvm -asan-globals=0
			// https://github.com/llvm-mirror/llvm/blob/ef512ca8e66e2d6abee71b9729b2887cb094cb6e/lib/Transforms/Instrumentation/AddressSanitizer.cpp
			// -asan-* has no path related options

		case strings.HasPrefix(arg, "-regalloc="):
			// https://b/issues/141210713#comment4
			// -mllvm -regalloc=pbqp
			// https://github.com/llvm-mirror/llvm/blob/be9f44f943df228dbca68139efef55f2c7666563/lib/CodeGen/TargetPassConfig.cpp
			// -regalloc= doesn't take path related value,
			// "basic", "fast", "greedy", "pbqp", etc.

		case strings.HasPrefix(arg, "-pbqp-"):
			// https://b/issues/141210713#comment4
			// -mllvm -pbqp-coalescing
			// https://github.com/llvm-mirror/llvm/blob/114087caa6f95b526861c3af94b3093d9444c57b/lib/CodeGen/RegAllocPBQP.cpp

		case strings.HasPrefix(arg, "-instcombine-"):
			// https://b/issues/161304121
			// -mllvm -instcombine-lower-dbg-declare=0
			// -instcombine-* defined in
			// https://github.com/llvm/llvm-project/blob/8e9a505139fbef7d2e6e9d0adfe1efc87326f9ef/llvm/lib/Transforms/InstCombine/InstructionCombining.cpp

		case strings.HasPrefix(arg, "-basic-aa"):
			// -mllvm -basic-aa-recphi=0
			// https://github.com/llvm/llvm-project/blob/694ded37b9d70e385addfc482d298b054073ebe1/llvm/lib/Analysis/BasicAliasAnalysis.cpp

		case strings.HasPrefix(arg, "-sanitizer-coverage-"):
			// -mllvm -sanitizer-coverage-prune-blocks=1
			// https://github.com/llvm/llvm-project/blob/93ec6cd684265161623b4ea67836f022cd18c224/llvm/lib/Transforms/Instrumentation/SanitizerCoverage.cpp

		case strings.HasPrefix(arg, "-enable-dse-memoryssa="):
			// https://crbug.com/1127713
			// -mllvm -enable-dse-memoryssa={true,false}
			// doesn't take a path related value.

		case strings.HasPrefix(arg, "-enable-ml-inliner="):
			// --mllvm -enable-ml-inliner={development,release}
			// does't take a path related value.

		case strings.HasPrefix(arg, "-training-log="):
			if filepath.IsAbs(arg[len("-training-log="):]) {
				return fmt.Errorf("abs path: %s", arg)
			}

		case strings.HasPrefix(arg, "-ml-inliner-model-under-training="):
			if filepath.IsAbs(arg[len("-ml-inliner-model-under-training="):]) {
				return fmt.Errorf("abs path: %s", arg)
			}

		case strings.HasPrefix(arg, "-limited-coverage-experimental="):
			// b/229939600 -limited-coverage-experimental=true
			// doesn't take a path related value.

		default:
			return unknownFlagError{arg: fmt.Sprintf("llvm: %s", arg)}
		}
	}
	return nil
}

func asArgRelocatable(filepath clientFilePath, args []string) error {
	for _, arg := range args {
		switch {
		case arg == "--fatal-warnings":
			// b/173641495
			// https://github.com/llvm/llvm-project/blob/ffc5d98d2c0df5f72ce67e5dcb724b64f03f639b/llvm/lib/MC/MCTargetOptionsCommandFlags.cpp
		default:
			return unknownFlagError{arg: fmt.Sprintf("as: %s", arg)}
		}
	}
	return nil
}

// gccOutputs returns output files from gcc command line.
// TODO: implicit obj output (without -o, but -c).
// TODO: -MD / -MMD without -MF case.
func gccOutputs(args []string) []string {
	var outputs []string
	var objout string
	outputArg := false
	splitDwarf := false
	mfArg := false

	for _, arg := range args {
		switch {
		case arg == "-o":
			outputArg = true
		case outputArg:
			objout = arg
			outputArg = false
		case strings.HasPrefix(arg, "-o"):
			objout = arg[2:]

		case arg == "-gsplit-dwarf":
			splitDwarf = true

		case arg == "-MF":
			mfArg = true
		case mfArg:
			outputs = append(outputs, arg)
			mfArg = false
		case strings.HasPrefix(arg, "-MF"):
			outputs = append(outputs, arg[3:])

		}
	}
	if objout != "" {
		outputs = append(outputs, objout)
		if splitDwarf {
			outputs = append(outputs, strings.TrimSuffix(objout, ".o")+".dwo")
		}
	}
	return outputs
}
