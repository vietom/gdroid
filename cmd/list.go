package cmd

import (
	"os/exec"
	"strings"
	"unicode"
	"github.com/spf13/cobra"
	"github.com/nhooyr/color"
)

func GetRunningEmulatedDevices() []string{
	out, err := exec.Command(adb, "devices").Output()
	if err != nil {
		color.Printf("%h[fgRed]%v%r\n", err)
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	running_devices := make([]string, 0, 5)
	for _, line := range lines {
		for index, char := range line {
			if unicode.IsSpace(char) {
				running_devices = append(running_devices, line[:index])
				break
			}
		}
	}
	emulated_devices := make([]string, 0, 5)
	for _, device := range running_devices {
		if strings.HasPrefix(device, "emulator-") {
			emulated_devices = append(emulated_devices, device)
		}
	}
	return emulated_devices
}

func GetAVDs() []string{
	out, err := exec.Command(emulator, "-list-avds").Output()
	if err != nil {
		color.Printf("%h[fgRed]%v%r\n", err)
	}
	avds := strings.Split(strings.TrimSpace(string(out)), "\n")
	return avds
}

var listCmd = &cobra.Command {
	Use: "list",
	Short: "List available AVDs and running emulated devices",
	Run: func(cmd *cobra.Command, args []string) {
		color.Printf("%h[fgCyan]Available AVDs:%r\n")
		available_avds := GetAVDs()
		if len(available_avds) == 0 {
			color.Printf("%h[fgYellow]No avds available%r\n")
		}
		for _, avd := range available_avds {
			color.Println(avd)
		}

		color.Printf("%h[fgCyan]Running emulated devices:%r\n")
		running_emulated_devices := GetRunningEmulatedDevices()
		if len(running_emulated_devices) == 0 {
			color.Printf("%h[fgYellow]No running devices detected%r\n")
		}
		for _, device := range running_emulated_devices {
			color.Println(device)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
