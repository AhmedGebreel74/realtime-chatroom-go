package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Client struct {
	ID   string
	Name string
	Conn net.Conn
}

var (
	clients   = make(map[string]*Client)
	clientsMu sync.Mutex
	broadcast = make(chan struct {
		SenderID string
		Content  string
	})
)

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)

	// Read client ID and name
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	parts := strings.Split(data, "|")

	clientID := parts[0]
	clientName := parts[0]
	if len(parts) > 1 {
		clientName = parts[1]
	}

	client := &Client{
		ID:   clientID,
		Name: clientName,
		Conn: conn,
	}

	clientsMu.Lock()

	// Send current users
	if len(clients) > 0 {
		var currentUsers []string
		for _, c := range clients {
			currentUsers = append(currentUsers, fmt.Sprintf("%s (%s)", c.Name, c.ID))
		}
		fmt.Fprintf(conn, "🟢 Current users in chat: %s\n", strings.Join(currentUsers, ", "))
	} else {
		fmt.Fprintf(conn, "🟢 Current users in chat: none\n")
	}

	// Add new user
	clients[clientID] = client
	clientsMu.Unlock()

	fmt.Printf("✅ [SERVER] User '%s' (%s) joined the chat\n", clientName, clientID)

	// Notify others
	clientsMu.Lock()
	for uid, c := range clients {
		if uid != clientID {
			fmt.Fprintf(c.Conn, "🔔 ** %s (%s) joined the chat **\n", clientName, clientID)
		}
	}
	clientsMu.Unlock()

	// Message loop
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		msg = strings.TrimSpace(msg)

		// Exit command
		if msg == "exit" {
			break
		}

		if msg != "" {
			broadcast <- struct {
				SenderID string
				Content  string
			}{
				SenderID: clientID,
				Content:  fmt.Sprintf("[%s (%s)]: %s", clientName, clientID, msg),
			}
		}
	}

	// Remove user
	clientsMu.Lock()
	delete(clients, clientID)
	clientsMu.Unlock()

	fmt.Printf("❌ [SERVER] User '%s' (%s) left the chat\n", clientName, clientID)

	// Notify others
	broadcast <- struct {
		SenderID string
		Content  string
	}{
		SenderID: clientID,
		Content:  fmt.Sprintf("👋 ** %s (%s) left the chat **", clientName, clientID),
	}

	conn.Close()
}

func broadcaster() {
	for {
		msg := <-broadcast

		clientsMu.Lock()
		for uid, c := range clients {
			if uid == msg.SenderID {
				continue
			}
			fmt.Fprintln(c.Conn, msg.Content)
		}
		clientsMu.Unlock()
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	fmt.Println("🚀 Server running on port 8888...")
	go broadcaster()

	for {
		conn, err := ln.Accept()
		if err == nil {
			go handleClient(conn)
		}
	}
}
