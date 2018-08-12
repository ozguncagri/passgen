package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/AlecAivazis/survey.v1"
)

func askForString(question string, errorString string) string {
	var out string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(question)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			out = scanner.Text()
			break
		} else {
			fmt.Print(errorString)
		}
	}

	return strings.TrimSpace(out)
}

func askForPassword(question string, errorString string) string {
	var (
		pwd []byte
		err error
	)

	fmt.Print(question)
	for len(pwd) == 0 {
		pwd, err = terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatalf(errorString)
		}
	}
	fmt.Print("\n")

	return strings.TrimSpace(string(pwd))
}

func askForInt(question string, errorString string) (out int) {
	scanner := bufio.NewScanner(os.Stdin)

	for out <= 0 {
		var err error
		fmt.Print(question)
		scanner.Scan()
		out, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(errorString)
		}
	}

	return
}

func askForCharPool() string {
	charPools := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select Character Pool Items:",
		Options: []string{"Upper", "Lower", "Number", "Symbols"},
	}
	err := survey.AskOne(prompt, &charPools, nil)
	if err != nil {
		os.Exit(1)
	}

	for i, v := range charPools {
		charPools[i] = string(v[0])
	}

	return strings.ToUpper(strings.Join(charPools, ""))
}
