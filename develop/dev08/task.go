package main

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с
поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи
запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись
ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный
сеанс поддерживается до тех пор, пока не будет введена
команда выхода (например \quit).
*/

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var homeDir string

// Буфер для реализации конвейера
var buf *bytes.Buffer

// Вывод текущего положения относительно домашней директории
func printDir() {
	curDir, _ := os.Getwd()
	fmt.Print(strings.TrimPrefix(curDir, homeDir), ":~$	")
}

func main() {
	// Чтение из стандартного потока
	file := os.Stdin
	scanner := bufio.NewScanner(file)

	// Вывод текущего положения относительно домашней директории
	homeDir, _ = os.UserHomeDir()
	printDir()

	buf = new(bytes.Buffer)

	// Считывание команд
	for scanner.Scan() {

		// Разделение команд в пайплайне
		pipelines := strings.Split(scanner.Text(), "|")

		// Выполнение каждой команды
		for cmdnum, line := range pipelines {
			first := (cmdnum == 0)
			last := (cmdnum == len(pipelines)-1)
			args := parseArgs(line)
			if args[0] == `\quit` {
				return
			}
			execCmd(args, first, last)
		}
		printDir()
	}
}

func execCmd(args []string, first, last bool) {
	var input io.Reader
	var output io.Writer
	var err error
	command := args[0]
	args = args[1:]
	if first {
		input = os.Stdin
	} else {
		input = buf
	}
	if last {
		output = os.Stdout
	} else {
		output = buf
	}
	// Выполнение команды
	switch command {
	case "cd":
		err = os.Chdir(args[0])
		if err != nil {
			log.Println(err)
		}
	case "echo":
		execEcho(args)
	case "kill":
		execKill(args)
	case "pwd":
		execPwd(args)
	default:
		standartCmd(command, input, output, args)
	}
}

// Выполнение обычных команд
func standartCmd(command string, input io.Reader, output io.Writer, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdin = input
	cmd.Stdout = output
	err := cmd.Run()
	if err != nil {
		log.Printf("error while executing command:%s", err)
	}
}

// Выполнение Echo
func execEcho(args []string) {
	for _, arg := range args {
		variable := strings.TrimLeft(arg, "$")
		if variable == arg {
			fmt.Print(arg, " ")
		} else {
			fmt.Print(os.Getenv(variable), " ")
		}
	}
	fmt.Println()
}

func execKill(args []string) {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println("invalid process id")
	}
	process := os.Process{
		Pid: pid,
	}
	err = process.Kill()
	if err != nil {
		log.Println("cannot kill process: ", err)
	}
}

func execPwd(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(dir)
}

func parseArgs(cmdLine string) []string {
	args := strings.FieldsFunc(cmdLine, func(c rune) bool {
		return c == ' ' || c == '\t'
	})
	return args
}
