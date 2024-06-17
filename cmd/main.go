package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/NalbertLeal/chatbot-regexp/pkg"
)

var (
	filename string
)

func main() {
	flag.StringVar(&filename, "file", "", "The file location with all the patterns need to the talk")
	flag.Parse()

	bot := pkg.ReadBotPatterns(filename)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		answer := bot.Response(input[:len(input)-2])
		fmt.Println(answer)
		if input[:len(input)-2] == "quit" {
			return
		}
	}

}
