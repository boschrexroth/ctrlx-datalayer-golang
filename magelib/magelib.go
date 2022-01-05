package magelib

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

// BuildExecutables (cross-) builds executables for an operating system and an architecture.
// Supported goos_arch combinations are linux_amd64, linux_arm64 and windows_amd64.
// sdkPath is the path where the ctrlX SDK is installed.
// The executables are build to the subdir public/goos_arch of outputPath.
// flags and packages are same a for the go build command.
func BuildExecutables(ctx context.Context, goos, arch, sdkPath, outputPath string, flags, packages []string) error {
	return build(ctx, goos, arch, sdkPath, outputPath, false, flags, packages)
}

// BuildPackages (cross-) builds packages for an operating system (linux, windows) and an architecture (amd64, arm64).
// Supported goos_arch combinations are linux_amd64, linux_arm64 and windows_amd64.
// sdkPath is the path where the ctrlX SDK is installed.
// flags and packages are same a for the go build command.
func BuildPackages(ctx context.Context, goos, arch, sdkPath string, flags, packages []string) error {
	return build(ctx, goos, arch, sdkPath, "", true, flags, packages)
}

// RunGoCommand runs a go command (vet, test, ...) with the right CGO_LDFLAGS and LD_LIBRARY_PATH
// when using the pkg/datalayer package as dependency.
// sdkPath is the path where the ctrlX SDK is installed.
func RunGoCommand(ctx context.Context, sdkPath string, goArgs ...string) error {
	ldFlags, err := GetLdFlags(ctx, runtime.GOOS, "amd64", sdkPath)
	if err != nil {
		return err
	}
	ldLibraryPath, err := GetLdLibraryPath(sdkPath)
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(ctx, "go", goArgs...)
	cmd.Env = append(os.Environ(), ldFlags, ldLibraryPath, "CGO_ENABLED=1")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%v\n%v\n\n%w", cmd, string(out), err)
	}
	return nil
}

// GetLdFlags returns the CGO_LDFLAGS needed to build when using the pkg/datalayer package as dependency.
// Supported goos_arch combinations are linux_amd64, linux_arm64 and windows_amd64.
// sdkPath is the path where the ctrlX SDK is installed.
//
// Note that getting the flags for windows_amd64 is not side effect free because of https://github.com/golang/go/issues/43244.
// Because of the bug, a directory containing only comm_datalayer.dll needs to be created.
// GetLdFlags creates ./temp/windows_amd64 for this purpose.
func GetLdFlags(ctx context.Context, goos, arch, sdkPath string) (string, error) {
	var err error
	if sdkPath, err = filepath.Abs(sdkPath); err != nil {
		return "", fmt.Errorf("make %v absolute: %w", sdkPath, err)
	}
	switch goos {
	case "linux":
		var gccArch string
		switch arch {
		case "amd64":
			gccArch = "linux-gcc-x64"
		case "arm64":
			gccArch = "linux-gcc-aarch64"
		default:
			return "", fmt.Errorf("invalid arch %v", arch)
		}
		return "CGO_LDFLAGS=" +
			fmt.Sprintf("-L%v/public/bin/comm.datalayer/%v/release ", sdkPath, gccArch), nil
	case "windows":
		// Bug https://github.com/golang/go/issues/43244
		// dll and lib together in library path does not work -> Use only the dll
		tmpDepDir, err := filepath.Abs("temp/windows_amd64")
		if err != nil {
			return "", err
		}
		if err := mkDir(tmpDepDir); err != nil {
			return "", err
		}
		// copy comm_datalayer.dll to temp dir
		winDllPath := path.Join(sdkPath, "public/bin/comm.datalayer/win-msvc-x64/release")
		if winDllPath, err = filepath.Abs(winDllPath); err != nil {
			return "", fmt.Errorf("make %v absolute: %w", winDllPath, err)
		}
		if err := copy(path.Join(winDllPath, "comm_datalayer.dll"), path.Join(tmpDepDir, "comm_datalayer.dll")); err != nil {
			return "", err
		}
		return "CGO_LDFLAGS=-L" + tmpDepDir, nil
	default:
		return "", fmt.Errorf("invalid OS %v", runtime.GOOS)
	}

}

// GetLdLibraryPath returns the LD_LIBRARY_PATH needed for running executables
// when using the pkg/datalayer package as dependency.
func GetLdLibraryPath(sdkPath string) (string, error) {
	var err error
	if sdkPath, err = filepath.Abs(sdkPath); err != nil {
		return "", err
	}
	var gccArch string
	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			gccArch = "linux-gcc-x64"
		case "arm64":
			gccArch = "linux-gcc-aarch64"
		}
	case "windows":
		gccArch = "win-msvc-x64"
	default:
		return "", fmt.Errorf("invalid OS %v", runtime.GOOS)
	}
	return "LD_LIBRARY_PATH=" + os.Getenv("LD_LIBRARY_PATH") + ":" +
		fmt.Sprintf("%v/public/bin/comm.datalayer/%v/release", sdkPath, gccArch), nil
}

// PrepareWin64Runtime copies all necessary dlls for running an excecutable using datalayer to the outputPath.
// To get the OpenSSL dlls, add
// <DependsOn Org="com.boschrexroth" Name="oss.openssl" Version="1.1.1-OpenSSL_1_1_1-stab.328" Destination="./deps/openssl" UsedFor="interface"/>
// to your cdf.xml
func PrepareWin64Runtime(ctx context.Context, sdkPath, opensslPath, outputPath string) error {
	fmt.Printf("Preparing Win64 Runtime %v\n", outputPath)
	if err := copy(path.Join(sdkPath, "public/bin/comm.datalayer/win-msvc-x64/release", "comm_datalayer.dll"), path.Join(outputPath, "comm_datalayer.dll")); err != nil {
		return err
	}
	if err := copy(path.Join(sdkPath, "public/bin/zmq/win-msvc-x64/release", "libzmq.dll"), path.Join(outputPath, "libzmq.dll")); err != nil {
		return err
	}
	if err := copy(path.Join(opensslPath, "public/bin/win-msvc-x64/release", "libcrypto-1_1-x64.dll"), path.Join(outputPath, "libcrypto-1_1-x64.dll")); err != nil {
		return err
	}
	if err := copy(path.Join(opensslPath, "public/bin/win-msvc-x64/release", "libssl-1_1-x64.dll"), path.Join(outputPath, "libssl-1_1-x64.dll")); err != nil {
		return err
	}
	return nil
}

func build(ctx context.Context, goos, arch, sdkPath, outputPath string, isLib bool, flags, packages []string) error {
	fmt.Printf("Building %v for %v %v...\n", packages, goos, arch)
	ldFlags, err := GetLdFlags(ctx, goos, arch, sdkPath)
	if err != nil {
		return err
	}
	// set c compiler
	var cc string
	if goos == "linux" && arch == "arm64" {
		cc = "CC=aarch64-linux-gnu-gcc"
	} else if goos == "windows" {
		cc = "CC=x86_64-w64-mingw32-gcc"
	}
	args := []string{"build"}
	if !isLib {
		// make out dir
		specificOutputPath := filepath.Clean(path.Join(outputPath, "public/"+goos+"_"+arch))
		if err := mkDir(specificOutputPath); err != nil {
			return err
		}
		args = append(args, "-o", specificOutputPath)
	}
	// build it
	args = append(args, flags...)
	args = append(args, packages...)
	buildCmd := exec.CommandContext(ctx, "go", args...)
	buildCmd.Env = append(os.Environ(), ldFlags, cc, "GOARCH="+arch, "GOOS="+goos, "CGO_ENABLED=1")
	if out, err := buildCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%v\n%v\n%w", buildCmd, string(out), err)
	}
	return nil
}

func mkDir(newPath string) error {
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		if err := os.MkdirAll(newPath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func copy(sourcePath, destPath string) error {
	bytesRead, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}
	if err := os.Remove(destPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	fileInfo, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(destPath, bytesRead, fileInfo.Mode()); err != nil {
		return err
	}
	return nil
}
