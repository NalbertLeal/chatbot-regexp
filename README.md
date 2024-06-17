# ChatBot

![GitHub](https://img.shields.io/github/license/NalbertLeal/chatbot-regexp)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/NalbertLeal/chatbot-regexp)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/NalbertLeal/chatbot-regexp)

## Description

_ChatBot_ is a chat bot developed to receive text and match it with patterns previous writen, with that the answers (also previous writen) can be selected.

## How to Use

### Prerequisites

Ensure you have Go installed on your machine. For instructions on Go installation, visit [https://golang.org/doc/install](https://golang.org/doc/install).

### Installation

```bash
git clone https://github.com/NalbertLeal/chatbot-regexp
```

### Basic Usage

To start a conversation you need a json file with all the chatbot patterns of question and answers. The folder "chat-bot-patterns" already have some configurations prepared. For example, if you want to talk with the chatbot doctor:

```bash
go run .\cmd\main.go -file .\chat-bot-patterns\aquarium-helper.json
```

The parameter **-file** is the path to the patterns file.

## Contribution

If you wish to contribute improvements, bug fixes, or new features, please feel free to open an issue or submit a pull request. Your contributions are highly appreciated!

## License

This project is licensed under the [GNU LESSER GENERAL PUBLIC LICENSE](https://choosealicense.com/licenses/lgpl-2.1/).