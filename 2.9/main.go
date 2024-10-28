package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("my-shell> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error: ", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		if input == "\\quit" {
			break
		}
		err = executeCommand(input)
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}

func executeCommand(input string) error {
	commands := parseInput(input)
	if len(commands) == 0 {
		return nil
	}
	if len(commands) == 1 {
		handled, err := handleBuildInCommand(commands[0])
		if handled {
			return err
		}
		return executeExturnalCommand(commands[0])
	}

	return fmt.Errorf("pipeline is not supported yet")
}

func executeExturnalCommand(cmd []string) error {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func parseInput(input string) [][]string {
	pipeCommands := strings.Split(input, "|")
	commands := [][]string{}
	for _, cmd := range pipeCommands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}
		parts := strings.Fields(cmd)
		commands = append(commands, parts)
	}
	return commands
}

func handleBuildInCommand(cmd []string) (bool, error) {
	switch cmd[0] {
	case "cd":
		if len(cmd) < 2 {
			return true, fmt.Errorf("cd: missing argument")
		}
		err := os.Chdir(cmd[1])
		return true, err
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return true, err
		}
		fmt.Println(dir)
		return true, nil
	case "echo":
		fmt.Println(strings.Join(cmd[1:], " "))
		return true, nil
	case "kill":
		if len(cmd) < 2 {
			return true, fmt.Errorf("kill: missing PID")
		}
		pid, err := strconv.Atoi(cmd[1])
		if err != nil {
			return true, fmt.Errorf("kill: incorrect PID")
		}
		prosess, err := os.FindProcess(pid)
		if err != nil {
			return true, err
		}
		err = prosess.Kill()
		return true, err
	case "ps":
		err := handlePsCommand()
		return true, err
	default:
		return false, nil
	}
}

func handlePsCommand() error {
	processes, err := process.Processes()
	if err != nil {
		return fmt.Errorf("не удалось получить список процессов: %v", err)
	}

	fmt.Println("PID\tTTY\t\tTIME\tCMD")

	for _, p := range processes {
		pid := p.Pid

		cmd, err := p.Name()
		if err != nil {
			cmd = "Unknown"
		}

		tty, err := p.Terminal()
		if err != nil || tty == "" {
			tty = "?"
		}

		cpuTimes, err := p.Times()
		if err != nil {
			continue
		}
		totalTime := cpuTimes.User + cpuTimes.System
		timeStr := formatCPUTime(totalTime)

		fmt.Printf("%d\t%s\t%s\t%s\n", pid, tty, timeStr, cmd)
	}

	return nil
}

func formatCPUTime(totalSeconds float64) string {
	total := int(totalSeconds)
	hours := total / 3600
	minutes := (total % 3600) / 60
	seconds := total % 60

	return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
}
