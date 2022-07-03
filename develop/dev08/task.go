package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/**
8. Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд
*/

/**
Пакет exec запускает внешние команды. Он обертывает os.StartProcess,
чтобы сделать его проще переназначить stdin и stdout,
соединить ввод /вывод с помощью каналов и сделать другие корректировки.
*/

const (
	CommandEcho = "echo"
	CommandCd   = "cd"
	CommandKill = "kill"
	CommandPwd  = "pwd"
	CommandExit = "quit"
	CommandPs   = "ps"
	ExitText    = "Exit"
)

type Command interface {
	Exec(args ...string) ([]byte, error)
}

// echo command interface
type echoCmd struct {
}

func (e *echoCmd) Exec(args ...string) ([]byte, error) {
	return exec.Command("echo", args...).Output()
}

// cd command interface
type cdCmd struct {
}

func (c *cdCmd) Exec(args ...string) ([]byte, error) {
	dir := args[0]
	// change directory
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	//get current path
	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	return []byte(dir), nil
}

// pwd command interface
type pwdCmd struct {
}

func (p *pwdCmd) Exec(args ...string) ([]byte, error) {
	// get rooted path name of current directory
	dir, err := os.Getwd()
	return []byte(dir), err
}

// kill command interface
type killCmd struct {
}

func (k *killCmd) Exec(args ...string) ([]byte, error) {
	// convet pid to int
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// try to find process using pid
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// kill process
	err = process.Kill()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return []byte("killed"), nil
}

// ps interface
type psCmd struct {
}

func (p *psCmd) Exec(args ...string) ([]byte, error) {
	return exec.Command("ps", args...).Output()
}

// Shell - UNIX-шелл-утилита с поддержкой ряда простейших команд
type Shell struct {
	command Command
	output  io.Writer
}

func (s *Shell) SetCommand(cmd Command) {
	s.command = cmd
}

// Run - выполнение конкретной команды
func (s *Shell) run(args ...string) {
	b, err := s.command.Exec(args...)
	_, err = fmt.Fprintln(s.output, string(b))
	if err != nil {
		fmt.Println("[err]", err.Error())
		return
	}
}

// ExecuteCommands Исполняет команды, которые ввел пользователь
func (s *Shell) ExecuteCommands(cmds []string) {
	for _, command := range cmds {
		args := strings.Split(command, " ")

		com := args[0]
		if len(args) > 1 {
			args = args[1:]
		}

		switch com {
		case CommandEcho:
			cmd := &echoCmd{}
			s.SetCommand(cmd)

		case CommandCd:
			cmd := &cdCmd{}
			s.SetCommand(cmd)

		case CommandKill:
			cmd := &killCmd{}
			s.SetCommand(cmd)

		case CommandPwd:
			cmd := &pwdCmd{}
			s.SetCommand(cmd)

		case CommandPs:
			cmd := &psCmd{}
			s.SetCommand(cmd)

		case CommandExit:
			_, err := fmt.Fprintln(s.output, ExitText)
			if err != nil {
				fmt.Println("[err]", err.Error())
				return
			}
			os.Exit(1)
		default:
			fmt.Println("wrong command")
			continue
		}
		s.run(args...)
	}
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	var output = os.Stdout

	shell := &Shell{output: output}
	for {
		fmt.Print(">: ")

		if scan.Scan() {
			line := scan.Text()
			cmds := strings.Split(line, " | ")

			shell.ExecuteCommands(cmds)
		}
	}
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print("> ")
// 		// Read the keyboad input.
// 		input, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 		}

// 		// Handle the execution of the input.
// 		if err = execInput(input); err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 		}
// 	}
// }

// // ErrNoPath is returned when 'cd' was called without a second argument.
// var ErrNoPath = errors.New("path required")

// func execInput(input string) error {
// 	// Remove the newline character.
// 	input = strings.TrimSuffix(input, "\n")

// 	// Split the input separate the command and the arguments.
// 	args := strings.Split(input, " ")

// 	// Check for built-in commands.
// 	switch args[0] {
// 	case "cd":
// 		// 'cd' to home with empty path not yet supported.
// 		if len(args) < 2 {
// 			return ErrNoPath
// 		}
// 		// Change the directory and return the error.
// 		return os.Chdir(args[1])
// 	case "exit":
// 		os.Exit(0)
// 	}

// 	// Prepare the command to execute.
// 	cmd := exec.Command(args[0], args[1:]...)

// 	// Set the correct output device.
// 	cmd.Stderr = os.Stderr
// 	cmd.Stdout = os.Stdout

// 	// Execute the command and return the error.
// 	return cmd.Run()
// }
