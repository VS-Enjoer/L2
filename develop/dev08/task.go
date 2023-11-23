package task

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func executeCommand(command string) {
	// Разделяем команду на аргументы
	args := strings.Fields(command)

	// Проверка наличия команды
	if len(args) == 0 {
		return
	}

	// Обработка встроенных команд
	switch args[0] {
	case "cd":
		if len(args) > 1 {
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("cd:", err)
			}
		} else {
			fmt.Println("cd: missing argument")
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("pwd:", err)
		} else {
			fmt.Println(dir)
		}
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) > 1 {
			cmd := exec.Command("kill", args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("kill:", err)
			}
		} else {
			fmt.Println("kill: missing argument")
		}
	case "ps":
		cmd := exec.Command("ps", "aux")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("ps:", err)
		}
	default:
		// Обработка внешних команд с использованием fork и exec
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Выводим приглашение для пользователя
		fmt.Print("> ")

		// Считываем ввод пользователя
		scanner.Scan()
		userInput := scanner.Text()

		// Проверка на команду выхода
		if userInput == "\\quit" {
			break
		}

		// Поддержка конвейеров
		if strings.Contains(userInput, "|") {
			commands := strings.Split(userInput, "|")

			var cmd *exec.Cmd
			var err error

			for _, command := range commands {
				// Обработка каждой команды в конвейере
				args := strings.Fields(command)
				if cmd == nil {
					cmd = exec.Command(args[0], args[1:]...)
				} else {
					cmd = exec.Command(args[0], args[1:]...)
					cmd.Stdin, err = cmd.StdoutPipe()
					if err != nil {
						fmt.Println("Error:", err)
						break
					}
				}
			}

			if cmd != nil {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		} else {
			// Выполнение обычной команды
			executeCommand(userInput)
		}
	}
}
