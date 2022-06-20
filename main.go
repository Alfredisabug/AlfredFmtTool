package main

import (
	"bufio"
	"fmt"
	"os"

	formatter "AlfredFmtTool/formatter"

	"github.com/sirupsen/logrus"
)

var optionString = `
Enter the code number to execute...
1.Cpp Header comment
2.Cpp function comment
3.CppDescrioptionBlock
4.exit
`

var endString = "Press enter key to return."

func init() {
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	var optionCode int
	var headerString string
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Print(optionString)
		fmt.Scanln(&optionCode)

		switch optionCode {
		case 1:
			fmt.Print("\033[H\033[2J")
			logrus.Debug("Enter Cpp Header comment format")

			fmt.Println("Enter header block name:")
			inputReader := bufio.NewReader(os.Stdin)
			headerString, _ = inputReader.ReadString('\n')
			result := formatter.CppHeaderBlockFmtFormatter(headerString)
			fmt.Print("\033[H\033[2J")
			fmt.Println(result)
			fmt.Println(endString)
			fmt.Scanln()
		case 2:
			fmt.Print("\033[H\033[2J")
			fmt.Print("Comming soon!")
			fmt.Scanln()
		case 3:
			fmt.Print("\033[H\033[2J")
			logrus.Debug("Enter Cpp Descrioption comment")
			fmt.Println("Enter Descrioption:")
			inputReader := bufio.NewReader(os.Stdin)
			headerString, _ = inputReader.ReadString('\n')
			result := formatter.CppDescrioptionBlockFormatter(headerString)

			// fmt.Print("\033[H\033[2J")
			fmt.Println(result)
			fmt.Println(endString)
			fmt.Scanln()
		case 4:
			fmt.Print("\033[H\033[2J")
			os.Exit(0)
		default:
			fmt.Println("Wrong input!")
		}
	}
}
