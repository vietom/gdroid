package cmd

import (
	"os"
	"fmt"
	"runtime"
	"github.com/spf13/cobra"
	"github.com/nhooyr/color"
)

var emulator = "emulator"
var adb = "adb"
var emulator_auth_token_file = fmt.Sprintf("%s/.emulator_console_auth_token", os.Getenv("HOME"))

func existsAndIsExecutable(file string) bool {
	// file exists?
	file_info, err := os.Stat(file)
	if err != nil {
		return false
	}

	// file is executable?
	mode := file_info.Mode()
	if mode & 0111 == 0 {
		return false
	}

	// it's an actual file?
	if !mode.IsRegular() {
		return false
	}

	return true
}

func checkAndroidBinaries(path string) (string, string) {
	emulator := fmt.Sprintf("%s/emulator/emulator", path)
	adb := fmt.Sprintf("%s/platform-tools/adb", path)
	if existsAndIsExecutable(emulator) && existsAndIsExecutable(adb) {
		return emulator, adb
	}
	return "", ""
}

func searchAndSetExecutables() {
	var e, a string
	if len(os.Getenv("ANDROID_HOME")) > 0 {
		e, a = checkAndroidBinaries(os.Getenv("ANDROID_HOME"))
	}else if len(os.Getenv("ANDROID_SDK_HOME")) > 0 {
		e, a = checkAndroidBinaries(os.Getenv("ANDROID_SDK_HOME"))
	}else if len(os.Getenv("ANDROID_SDK_ROOT")) > 0 {
		e, a = checkAndroidBinaries(os.Getenv("ANDROID_SDK_ROOT"))
	}else {
		switch(runtime.GOOS) {
		case "darwin": 
			e, a = checkAndroidBinaries(fmt.Sprintf("%s/Library/Android/sdk", os.Getenv("HOME")))
		default:
			e, a = checkAndroidBinaries(fmt.Sprintf("%s/Android/Sdk", os.Getenv("HOME")))
		}
	}

	if len(e) == 0 {
		color.Printf("%h[fgYellow]Couldn't find 'emulator' binary\nPlease point the $ANDROID_SDK_ROOT environment variable to the correct SDK location.%r\n")
	}else {
		emulator = e
	}
	if len(a) == 0 {
		color.Printf("%h[fgYellow]Couldn't find 'adb' binary\nPlease point the $ANDROID_SDK_ROOT environment variable to the correct SDK location.%r\n")
	}else {
		adb = a
	}
}

func init() {
	searchAndSetExecutables()
}

var RootCmd = &cobra.Command {
	Use:   "gdroid",
	Short: "gdroid is a command line manager for the android emulator",
}
