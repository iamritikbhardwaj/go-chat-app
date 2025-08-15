# 💬 Go WebSocket Chat App

A real-time chat application built using **Go (Golang)**, **WebSockets**, and **Vanilla JavaScript**, with a clean and responsive UI.

---

## 🚀 Features

- 🔥 Real-time messaging with WebSockets
- 🌐 Modern, responsive UI
- 👤 Sender identification (you vs others)
- ⚡ Lightweight and dependency-free frontend
- 🔒 Simple and secure WebSocket setup

---

## 🧰 Tech Stack

- [Go](https://golang.org/)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- HTML5, CSS3, JavaScript

---

## 🗂️ Project Structure

```

go-chat-app/
├── main.go                 # Go WebSocket server
├── templates/
│   └── index.html          # Chat UI
├── static/
│   └── index.js            # Frontend WebSocket logic
├── go.mod
└── go.sum

````

---

## 📦 Installation

### ✅ Prerequisites

- Go 1.18+
- Git

### 🔧 Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-chat-app.git
   cd go-chat-app
````

2. Install dependencies:

   ```bash
   go get github.com/gorilla/websocket
   ```

3. Run the app:

   ```bash
   go run .
   ```

4. Open in browser:

   ```
   http://localhost:8008
   ```

---

## 🔍 How It Works

* Clients connect to `ws://localhost:8008/ws`
* Each message sent includes:

  ```json
  {
    "user": "YourName",
    "text": "Hello world"
  }
  ```
* Server broadcasts the message to all connected clients
* Each client distinguishes between **own messages** and **others**

---

## 🖼️ UI Preview

* Messages you send: right-aligned blue bubbles
* Messages from others: left-aligned gray bubbles
* Clean chat layout with auto-scrolling

---

## 📌 Example Screenshot

![Preveiw](/go-chat-app/image.png)

---

## ✨ Future Improvements

* ✅ Add timestamps
* ✅ Store chat history in a database
* ✅ User authentication
* ✅ Private chats or chat rooms
* ✅ Emojis and attachments

---

## 📄 License

MIT License — free to use and modify.

---

## 🙋 Author

Made with ❤️ by **Ritik Singh**

[GitHub](https://github.com/yourusername) | [LinkedIn](https://linkedin.com/in/yourprofile)

