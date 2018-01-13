package cmd

import (
	"fmt"
	"strings"
	"strconv"
	"net"
	"bufio"
	"io/ioutil"
)

func discardUntilOK(rw *bufio.ReadWriter) error {
	for {
		line, err := rw.ReadString('\n')
		if err != nil {
			return err
		}
		if line == "OK\r\n" {
			return nil
		}
	}
}

func read(rw *bufio.ReadWriter) (string, error) {
	line, err := rw.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.TrimSpace(line)
	return line, nil
}

func write(rw *bufio.ReadWriter, cmd string) error {
		s := fmt.Sprintf("%s\r\n", cmd)
		_, err := rw.WriteString(s)
		if err != nil {
			return err
		}
		rw.Flush()
		return nil
}


func GetAVDNameForDevice(device string) (string, error) {
		// fetch key needed for the emulator telnet connection
		b, err := ioutil.ReadFile(emulator_auth_token_file) // just pass the file name
		if err != nil {
			return "", err
		}
		key := strings.TrimSpace(string(b))

		// open the telnet connection
		port, err := strconv.Atoi(strings.Split(device, "-")[1])
		if err != nil {
			return "", err
		}
		address := fmt.Sprintf("localhost:%d", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			return "", err
		}
		rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		defer conn.Close()

		// authenticate against the device
		err = discardUntilOK(rw)
		if err != nil {
			return "", err
		}
		auth_cmd := fmt.Sprintf("auth %s", key)
		err = write(rw, auth_cmd)
		if err != nil {
			return "", err
		}

		// fetch avd name
		err = discardUntilOK(rw)
		if err != nil {
			return "", err
		}
		err = write(rw, "avd name")
		if err != nil {
			return "", err
		}
		name, err := read(rw)
		if err != nil {
			return "", err
		}

		return name, nil
}
