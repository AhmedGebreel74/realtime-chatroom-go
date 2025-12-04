# 💬 Go Chatroom 🚀

A simple real-time TCP chatroom built with **Go** and **goroutines**!

---

## ✨ Features

🆔 **Auto-Generated IDs** – Each user gets a unique ID automatically  
👤 **Custom Names** – Enter your display name when connecting  
🎯 **Real-Time Chat** – Send and receive messages instantly  
👥 **Multiple Users** – Connect many clients at once  
📢 **Broadcasts** – Everyone sees new messages  
👋 **Join Notifications** – Know when users arrive and leave  
🔄 **Concurrent Handling** – Smooth multi-client support  
🛡️ **Thread-Safe** – Mutex protection for shared data  

---

## 📁 Files

| File | Purpose |
|------|---------|
| 🖥️ `server.go` | Listens for connections & broadcasts messages |
| 💻 `client.go` | Connects to server & sends/receives messages |
| 📖 `README.md` | This file! |

---

## 🚀 Quick Start

### Terminal 1 - Start Server
```bash
go run server.go
```
✅ Server runs on **localhost:8888**

### Terminal 2+ - Run Clients
```bash
go run client.go
```
Each client:
- 🆔 Gets auto-generated ID (User-XXXX)
- 👤 Prompts for your name
- 💬 Starts chatting!

---

## 💡 How It Works

### Server (`server.go`)
- 🎧 Listens on port **8888**
- 👥 Stores **Client struct** (ID, Name, Connection)
- 📤 Broadcasts messages to all except sender
- 🔐 Uses `sync.Mutex` for thread safety
- 🔔 Sends join/leave notifications with emojis
- 📋 Shows current users when someone connects

### Client (`client.go`)
- 🔌 Connects to server via TCP
- 🆔 Auto-generates unique random ID
- 👤 Prompts for name entry
- 📨 Sends messages with name + ID
- 🎵 Receives messages in real-time
- 🎨 Color-coded output:
  - 🔵 Blue = System messages
  - 🟢 Green = Your messages
  - 🟡 Yellow = Others' messages

---

## 📊 Architecture

```
┌─────────────────────────────────────┐
│      SERVER (Port 8888)             │
│                                     │
│  ┌─────────────────────────────┐   │
│  │   Client Map (Mutex)        │   │
│  │  - User-1234 (Alice)        │   │
│  │  - User-5678 (Bob)          │   │
│  │  - User-9012 (Charlie)      │   │
│  └─────────────────────────────┘   │
│                                     │
│  ┌─────────────────────────────┐   │
│  │   Broadcast Channel         │   │
│  │  (Relays to all except      │   │
│  │   sender)                   │   │
│  └─────────────────────────────┘   │
└─────────────────────────────────────┘
        ▲          ▲          ▲
        │          │          │
    Client 1   Client 2   Client 3
   (Auto ID)  (Auto ID)  (Auto ID)
```

---

## 🎮 Usage Example

```
🖥️ Terminal 1:
$ go run server.go
🚀 Server running on port 8888...

💻 Terminal 2:
$ go run client.go
🆔 Your ID: User-1234
👤 Enter your name: Alice
🟢 Current users in chat: none

💻 Terminal 3:
$ go run client.go
🆔 Your ID: User-5678
👤 Enter your name: Bob
🟢 Current users in chat: Alice (User-1234)

Alice receives:
🔔 ** Bob (User-5678) joined the chat **

Alice types: Hello Bob!
💬 Alice >

Bob receives:
[Alice (User-1234)]: Hello Bob!

Bob types: Hi Alice! 👋
💬 Bob >

Alice receives:
[Bob (User-5678)]: Hi Alice! 👋
```

---

## 🛠️ Technology Stack

| Component | Tech |
|-----------|------|
| Language | 🐹 Go |
| Networking | 🔗 TCP Sockets |
| Concurrency | ⚡ Goroutines |
| Synchronization | 🔐 Mutex |
| Message Format | 📝 Text (ID\|Name format) |
| Port | 🔌 8888 |

---

## 🎓 Key Concepts

- **Goroutines** – Lightweight threads for handling each client independently
- **Mutex** – Prevents race conditions when accessing the shared clients map
- **Broadcasting** – One goroutine relays messages to all connected clients
- **Client Struct** – Encapsulates ID, Name, and Connection for each user
- **TCP Protocol** – Direct socket communication between server and clients

---

## 🚦 Status

✅ **Working & Tested!**  
Ready for local chatting with multiple users! 🎉

---

## 🎯 Why This Project is Cool

✅ **Auto IDs** – No manual ID entry needed, auto-generated  
✅ **Custom Names** – Personalize your chat presence  
✅ **Simple Code** – Just a few lines of Go handle complex networking  
✅ **Educational** – Learn goroutines, mutexes, and TCP sockets  
✅ **Scalable** – Handles many concurrent users efficiently  
✅ **Real-World** – Same concepts used in production systems  

---

## 🔮 Future Enhancements

🔧 **Possible Additions:**
- Message history persistence
- Private messaging between users
- User authentication
- Web UI with WebSockets
- Command support (/help, /users, /leave)
- Typing indicators

💡 **Learning Opportunities:**
- Explore context.Context for graceful shutdown
- Add TLS/SSL encryption
- Implement advanced message queueing
- Load testing and performance tuning

---

## 📞 Support

Found a bug? Have ideas? Feel free to contribute! 🙌

**Happy chatting!** 💬✨
