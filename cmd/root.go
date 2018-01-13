package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

var emulator = "emulator"
var adb = "adb"
var emulator_auth_token_file = fmt.Sprintf("%s/.emulator_console_auth_token", os.Getenv("HOME"))

func init() {
	fmt.Println(os.Getenv("ANDROID_HOME"))
	if len(os.Getenv("ANDROID_HOME")) > 0 {
		emulator = fmt.Sprintf("%s/emulator/emulator", os.Getenv("ANDROID_HOME"))
		adb = fmt.Sprintf("%s/platform-tools/adb", os.Getenv("ANDROID_HOME"))
	}else if len(os.Getenv("ANDROID_SDK_HOME")) > 0 {
		emulator = fmt.Sprintf("%s/emulator/emulator", os.Getenv("ANDROID_SDK_HOME"))
		adb = fmt.Sprintf("%s/platform-tools/adb", os.Getenv("ANDROID_SDK_HOME"))
	}else if len(os.Getenv("ANDROID_SDK_ROOT")) > 0 {
		emulator = fmt.Sprintf("%s/emulator/emulator", os.Getenv("ANDROID_SDK_ROOT"))
		adb = fmt.Sprintf("%s/platform-tools/adb", os.Getenv("ANDROID_SDK_ROOT"))
	}
}

var RootCmd = &cobra.Command {
	Use:   "gdroid",
	Short: "gdroid is a command line manager for the android emulator",
}
