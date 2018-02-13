package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/spf13/cobra"
	"github.com/nhooyr/color"
)

func clearAllContacts(device string) error {
	//color.Printf("Clearing all contacts...'\n")
	err := exec.Command(
		adb,
		"-s",
		device,
		"shell",
		"pm",
		"clear",
		"com.android.providers.contacts",
	).Start()
	return err
}

var contactsCmd = &cobra.Command {
	Use: "contacts",
	Short: "Clear/Import contacts from phonebook",
	Run: func(cmd *cobra.Command, args []string) {
		auto, _ := cmd.Flags().GetBool("auto")
		var devices_to_be_considered []string

		if auto {
			devices_to_be_considered = GetRunningEmulatedDevices()
			if len(devices_to_be_considered) == 0 {
				color.Printf("%h[fgYellow]No running devices detected%r\n")
			}
		}else { // clear on the specified devices
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
						devices_to_be_considered = append(devices_to_be_considered, device)
						is_running = true
					}
				}
				if !is_running {
					color.Printf("%h[fgYellow]Skip%r not running device %h[fgCyan]%s%r\n", arg)
				}
			}

			// actually do stuff to them
			for _, device := range devices_to_be_considered {
				color.Printf("Attempting to %h[fgYellow]clear contacts%r on the device %h[fgCyan]%s%r\n", device)
				//_, err := exec.Command(adb, "-s", device, "emu", "kill").Output()
				err := clearAllContacts(device)
				if err != nil {
					color.Printf("%h[fgRed]%v%r\n", err)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(contactsCmd)
	contactsCmd.Flags().BoolP("auto", "a", true, "Clear all contacts from the device")
}
