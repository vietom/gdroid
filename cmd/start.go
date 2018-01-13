package cmd

import (
	"os/exec"
	"github.com/spf13/cobra"
	"github.com/nhooyr/color"
)

func isAvailable(available_avds []string, requested_avd string) bool {
	for _, avd := range available_avds {
		if avd == requested_avd {
			return true
		}
	}
	return false
}

func isRunning(running_emulated_devices []string, requested_avd string) bool {
	for _, device := range running_emulated_devices {
		name, err := GetAVDNameForDevice(device)
		if err != nil {
			color.Printf("%h[fgRed]%s%r\n", err)
			continue
		}

		if name == requested_avd {
			return true
		}
	}
	return false
}

var startCmd = &cobra.Command {
    Use: "start",
	Short: "Start specified AVDs",
    Run: func(cmd *cobra.Command, args []string) {
		auto, _ := cmd.Flags().GetInt("auto")
		avds := make(map[string]bool, 5)

		available_avds := GetAVDs()
		running_emulated_devices := GetRunningEmulatedDevices()

		// add the explicitly named avds
		for _, avd := range args {
			// check if avd is available and not already running
			if isAvailable(available_avds, avd) &&
				!isRunning(running_emulated_devices, avd) {
					avds[avd] = true
			}
		}

		// add the specified automatic avd selection
		if auto > 0 {
			color.Printf("%h[fgCyan]Automatically selecting %h[fgYellow]%d%h[fgCyan] AVDs%r\n", auto)
			// iterate over available avds, if avd not yet in list (and free!) add
			for _, avd := range available_avds {
				if auto == 0 { // found enough already
					break
				}
				// if already specified continue
				_, defined := avds[avd]
				if defined {
					continue
				}
				// if not yet running add it
				if !isRunning(running_emulated_devices, avd) {
					avds[avd] = true
					auto--
				}
			}
		}

		// start specified avds
		for avd := range avds {
			color.Printf("Starting device '%h[fgYellow]%s%r'\n", avd)
			exec.Command(
				emulator,
				"-writable-system",
				"-netdelay", "none",
				"-netspeed", "full",
				"-dns-server", "192.168.98.14,8.8.8.8",
				"-avd", avd,
			).Start()
		}
    },
}

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.Flags().IntP("auto", "a", 0, "Automatically determine which devices to start up")
}
