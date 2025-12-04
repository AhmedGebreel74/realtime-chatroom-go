package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	// Auto-generate ID
	rand.Seed(time.Now().UnixNano())
	clientID := fmt.Sprintf("User-%d", 1000+rand.Intn(9000))
	fmt.Printf("🆔 Your ID: %s\n", clientID)

	// Ask for name
	fmt.Print("👤 Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Anonymous"
	}

	// Send ID and name to server
	fmt.Fprintf(conn, "%s|%s\n", clientID, name)

	serverReader := bufio.NewReader(conn)

	msg, err := serverReader.ReadString('\n')
	if err == nil {
		msg = strings.TrimSpace(msg)
		if msg != "" && strings.HasPrefix(msg, "**") {
			fmt.Printf("\033[1;34m%s\033[0m\n\n", msg)
		}
	}

	go func() {
		for {
			msg, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println("\n❌ Disconnected from server")
				os.Exit(0)
			}

			msg = strings.TrimSpace(msg)
			displayMsg := msg

			if strings.HasPrefix(msg, "**") {
				fmt.Printf("\033[1;34m%s\033[0m\n\n", displayMsg)
			} else if strings.HasPrefix(msg, "["+clientID+"]") {
				displayMsg = strings.Replace(msg, "["+clientID+"]", "[you]", 1)
				fmt.Printf("\033[1;32m%s\033[0m\n\n", displayMsg)
			} else {
				fmt.Printf("\033[1;33m%s\033[0m\n\n", displayMsg)
			}

			fmt.Printf("💬 %s > ", name)
		}
	}()

	for {
		fmt.Printf("💬 %s > ", name)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Fprintln(conn, "exit")
			conn.Close()
			fmt.Println("👋 You left the chat.")
			os.Exit(0)
		}

		if text != "" {
			fmt.Fprintln(conn, text)
		}
	}
}
