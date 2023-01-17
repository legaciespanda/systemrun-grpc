/*
Copyright © 2023 Ernest Obot <ernestobot.dev@gmail.com>

*/

package helper

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

//run cmd commands
func run(cmd string) (string, error) {
	c := exec.Command("cmd", "/C", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
		return "Unable to shutdown system", fmt.Errorf("%v", err)
	}
	return "system has been shutdown", nil
}

//execute remote command on the PC
func RemoteCommand(cmd string) (string, error) {
	CommandWork := exec.Command("cmd", "/C", cmd)
	CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, err := CommandWork.Output()
	if err != nil {
		return "Unable to execute command", fmt.Errorf("%v", err)
	}
	return string(History), nil
}

//Kill or terminate ("Tool.exe")
func Kill(name string) (string, error) {
	c := exec.Command("cmd", "/C", "taskkill /F /IM "+name)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := c.Run(); err != nil {
		return "Unable to execute command", fmt.Errorf("%v", err) //RPCUpdate("Kill: " + err.Error())
	}
	return "File " + name + "is Killed", nil
}

//Start an exe; example calc
func StartEXE(name string) (string, error) {
	if strings.Contains(name, ".exe") {
		binary, err := exec.LookPath(name)
		if err != nil {
			return "Unable to start executable file", fmt.Errorf("%v", err)
		}
		exec.Command(binary).Run()
		return "Executable file launched", nil
	}
	return "Unrecognised file format", nil
}

//Opens a URL in a default browser
func OpenURL(URL string) (string, error) {
	err := exec.Command("cmd", "/c", "start", URL).Start()
	if err != nil {
		return "Unable to launch url in the browser", fmt.Errorf("%v", err)
	}
	return "Url launched on the server browser", nil
}

//remotely shutdown, restart, logoff, hibernate
func PowerOptions(mode string) (string, error) {
	if mode == "shutdown" {
		result, err := run("shutdown -s -t 00")
		if err != nil {
			return "", fmt.Errorf("%v", err)
		}
		return result, nil
	} else if mode == "restart" {
		result, err := run("shutdown -r -t 00")
		if err != nil {
			return "", fmt.Errorf("%v", err)
		}
		return result, nil
	} else if mode == "logoff" {
		result, err := run("shutdown -l -t 00")
		if err != nil {
			return "", fmt.Errorf("%v", err)
		}
		return result, nil
	} else if mode == "hibernate" {
		result, err := run("shutdown -h -t 00")
		if err != nil {
			return "", fmt.Errorf("%v", err)
		}
		return result, nil
	}
	return "Error", fmt.Errorf("%v", "Invalid shutdown command")
}
