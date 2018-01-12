package cmd

import (
	"os"
	"fmt"
	//"time"
	"github.com/spf13/cobra"
)

var emulator = fmt.Sprintf("%s/emulator/emulator", os.Getenv("ANDROID_HOME"))
var adb = fmt.Sprintf("%s/platform-tools/adb", os.Getenv("ANDROID_HOME"))
var emulator_auth_token_file = fmt.Sprintf("%s/.emulator_console_auth_token", os.Getenv("HOME"))
//const timeout = 45 * time.Second

var RootCmd = &cobra.Command {
  Use:   "gdroid",
  Short: "gdroid is a command line manager for the android emulator",
}
