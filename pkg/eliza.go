package pkg

import (
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var (
	EmptyAnswer = ""
)

type ChatBot struct {
	Replaces    map[string]string
	Reflections map[string]string
	Patterns    [][]string
	Answers     [][]string
}

// func New(Reflections map[string]string,
// 	Patterns []string,
// 	Answers [][]string,
// 	Replaces map[string]string) *ChatBot {
// 	return &ChatBot{
// 		Reflections: Reflections,
// 		Patterns:    Patterns,
// 		Answers:     Answers,
// 		Replaces:    Replaces,
// 	}
// }

func (c *ChatBot) cleanQuestion(question string) string {
	clean := question

	for replaced, toReplace := range c.Replaces {
		clean = strings.ReplaceAll(clean, replaced, toReplace)
	}

	return clean
}

func (c *ChatBot) reflectAnswer(answer string, matchGroup []string) string {
	reflected := answer

	index := strings.Index(reflected, "%")
	for index != -1 {
		patternGroup, err := strconv.Atoi(reflected[index+1 : index+2])
		if err != nil {
			log.Fatalln(err)
		}

		matchValue := matchGroup[patternGroup]
		substitution, mustSubstitute := c.Reflections[matchValue]
		if mustSubstitute {
			reflected = reflected[:index] + substitution + reflected[index+2:]
		} else {
			reflected = reflected[:index] + matchValue + reflected[index+2:]
		}

		index = strings.Index(reflected, "%")
	}

	return reflected
}

func (c *ChatBot) choiceAnswer(AnswersIndex int) string {
	Answers := c.Answers[AnswersIndex]
	randIndex := rand.Intn(len(Answers))
	return Answers[randIndex]
}

func (c *ChatBot) Response(question string) string {
	c.cleanQuestion(question)
	// c.matchQuestion(question)
	lower := strings.ToLower(question)

	for i, patterns := range c.Patterns {
		for _, pattern := range patterns {
			regexpPattern := regexp.MustCompile(pattern)
			match := regexpPattern.MatchString(lower)
			if match {
				groups := regexpPattern.FindAllStringSubmatch(lower, -1)
				answer := c.choiceAnswer(i)
				return c.reflectAnswer(answer, groups[0])
			}
		}
	}

	return EmptyAnswer
}
