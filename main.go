package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type IntFlag struct {
	Value int
	IsSet bool
}

func (f *IntFlag) Set(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	f.Value = v
	f.IsSet = true
	return nil
}

func (f *IntFlag) String() string {
	return fmt.Sprint(f.Value)
}

func readConfig(filePath string) (botToken, chatID string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "botToken=") {
			botToken = strings.TrimPrefix(line, "botToken=")
		} else if strings.HasPrefix(line, "chatID=") {
			chatID = strings.TrimPrefix(line, "chatID=")
		}
	}

	return botToken, chatID, scanner.Err()
}

func sendMessage(botToken, chatID, message string) {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	data := url.Values{
		"chat_id":    {chatID},
		"text":       {message},
		"parse_mode": {"markdown"},
	}

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		fmt.Println("[-] Failed to send:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if body, err := io.ReadAll(resp.Body); err == nil {
			fmt.Println("[-] Failed to send. Response:", string(body))
		}
	} else {
		fmt.Println("[+] Success.")
	}
}

func promptForInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func main() {
	fmt.Println(`
    ╔╦╗┌─┐┬  ┌─┐╔╗ ┌─┐┌┬┐╔═╗╔═╗╔═╗╔╦╗
     ║ ├┤ │  ├┤ ╠╩╗│ │ │ ╚═╗╠═╝╠═╣║║║
     ╩ └─┘┴─┘└─┘╚═╝└─┘ ┴ ╚═╝╩  ╩ ╩╩ ╩
    By @elliottophellia        v1.0.0
    `)

	botToken, chatID, err := readConfig("config")
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	var requestCount IntFlag
	flag.Var(&requestCount, "loop", "How many times to send the message")
	messagePtr := flag.String("msg", "", "The message to send")
	flag.Parse()

	if *messagePtr == "" {
		*messagePtr = promptForInput("The message to send? ")
		if *messagePtr == "" {
			fmt.Println("Please enter a message to send.")
			return
		}
	}

	if !requestCount.IsSet {
		input := promptForInput("How many times to send the message? ")
		if input == "" {
			fmt.Println("No request count specified, using default value of 1.")
			requestCount.Value = 1
		} else {
			count, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid request count, using default value of 1.")
				requestCount.Value = 1
			} else {
				requestCount.Value = count
			}
		}
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nReceived an interrupt, stopping...")
		os.Exit(0)
	}()

	for i := 0; i < requestCount.Value || requestCount.Value == 0; i++ {
		sendMessage(botToken, chatID, *messagePtr)
		if requestCount.Value == 0 {
			i = 0
		}
	}
}
