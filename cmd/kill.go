package cmd

import (
	"fmt"
	"strings"
	"os/exec"
	"github.com/spf13/cobra"
	"github.com/nhooyr/color"
)

var killCmd = &cobra.Command {
	Use: "kill",
	Short: "Kill specified running emulated devices",
	Run: func(cmd *cobra.Command, args []string) {
		auto, _ := cmd.Flags().GetBool("auto")
		var devices_to_be_killed []string

		// craft a list of devices to be killed
		if auto { // add all running avds
			devices_to_be_killed = GetRunningEmulatedDevices()
			if len(devices_to_be_killed) == 0 {
				color.Printf("%h[fgYellow]No running devices detected%r\n")
			}
		}else { // add the explicitly named avds
			// prefix the specified names with 'emulator-" if ommited
			for index, arg := range args {
				if !strings.HasPrefix(arg, "emulator-") {
					args[index] = fmt.Sprintf("emulator-%s", arg)
				}
			}

			// figure out which of the specified devices are running and valid to be killed
			running_emulated_devices := GetRunningEmulatedDevices()
			for _, arg := range args {
				is_running := false
				for _, device := range running_emulated_devices {
					if device == arg {
						devices_to_be_killed = append(devices_to_be_killed, device)
						is_running = true
					}
				}
				if !is_running {
					color.Printf("%h[fgYellow]Skip%r not running device %h[fgCyan]%s%r\n", arg)
				}
			}
		}

		// actually kill them
		for _, device := range devices_to_be_killed {
			color.Printf("Attempting to %h[fgYellow]kill%r device %h[fgCyan]%s%r\n", device)
			_, err := exec.Command(adb, "-s", device, "emu", "kill").Output()
			if err != nil {
				color.Printf("%h[fgRed]%v%r\n", err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(killCmd)
	killCmd.Flags().BoolP("auto", "a", false, "Kill all emulated devices")
}
