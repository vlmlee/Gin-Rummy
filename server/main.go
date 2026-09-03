package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/creack/pty"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins for dev
}

func handleXterm(w http.ResponseWriter, r *http.Request) {
	// 1. Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	// 2. Target the Go program you want to run (e.g., your compiled CLI app)
	// You can also target "bash" or "sh" if you want a full terminal wrapper.
	cmd := exec.Command("./my_go_cli_app") 

	// 3. Start the process tied to a Pseudo-Terminal (PTY)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Printf("Failed to start PTY: %v", err)
		return
	}
	defer ptmx.Close()

	// Channel to signal routine completion
	done := make(chan struct{})

	// 4. Goroutine: Pipe browser input -> Go App stdin (via PTY)
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				close(done)
				return
			}
			ptmx.Write(message)
		}
	}()

	// 5. Goroutine: Pipe Go App stdout/stderr -> Browser output (via WebSocket)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptmx.Read(buf)
			if err != nil {
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, buf[:n])
			if err != nil {
				return
			}
		}
	}()

	<-done
	cmd.Wait()
}

func main() {
	http.HandleFunc("/ws", handleXterm)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}