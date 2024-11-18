package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sinhanamanav030/chat-bot/server"
)

func main() {
	handler := server.NewHandler()

	go func() {
		router := gin.Default()
		router.POST("/messages", handler.PostMessageHandler)
		router.GET("/messages/:id", handler.GetMessageHandler)

		router.Run(":8080")
	}()

	reader := bufio.NewReader(os.Stdin)
	baseURL := "http://localhost:8080"

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Post a message")
		fmt.Println("2. Get a message by ID")
		fmt.Println("3. Exit")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			message = strings.TrimSpace(message)
			postMessage(baseURL+"/messages", message)

		case "2":
			fmt.Print("Enter message ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			getMessage(baseURL + "/messages/" + id)

		case "3":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func postMessage(url, message string) {
	payload := server.PostMessageRequest{
		Message: message,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding message:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error posting message:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}

func getMessage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting message:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}
