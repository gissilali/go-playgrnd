package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// cmdFreq returns the frequency of "go" subcommand usage in ZSH history
func cmdFreq(fileName string) (map[string]int, error) {
	commandMap := make(map[string]int)
	var commands []string
	// read file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	// defer file closure
	defer file.Close()
	// define the regex
	goCommandRegex, err := regexp.Compile("go .*")
	goSubCommandRegex := regexp.MustCompile("(go) (.*)")
	// scan each line for matches
	s := bufio.NewScanner(file)
	for s.Scan() {
		if goCommandRegex.MatchString(s.Text()) {
			//fmt.Println(s.Text())
			matches := goSubCommandRegex.FindStringSubmatch(s.Text())
			fmt.Println(matches)
			if strings.Split(matches[2], " ")[0] != "" {
				commands = append(commands, strings.Split(matches[2], " ")[0])
			}
		}
	}

	for _, command := range commands {
		commandMap[command]++
	}

	return commandMap, nil
}

func main() {
	freqs, err := cmdFreq("/home/silali/Work/playground/linkedin-go/Ch03/challenge/zsh_history")
	if err != nil {
		log.Fatal(err)
	}

	for cmd, count := range freqs {
		fmt.Printf("%s -> %d\n", cmd, count)
	}
}
