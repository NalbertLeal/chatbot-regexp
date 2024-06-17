package pkg

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func ReadBotPatterns(filename string) *ChatBot {
	content := readFile(filename)

	return readBotPatterns([]byte(content))
}

func readFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	content := ""
	for scanner.Scan() {
		content += scanner.Text()
	}

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func readBotPatterns(fileContent []byte) *ChatBot {
	bot := &ChatBot{}

	err := json.Unmarshal(fileContent, &bot)
	if err != nil {
		log.Fatal(err)
	}

	return bot
}
