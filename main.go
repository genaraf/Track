package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

var scon socketio.Conn = nil

func reportLoop() {
	var temp int = 0
	v, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err == nil {
		_, err := fmt.Sscanf(string(v), "%d", &temp)
		if err == nil {
			temp /= 1000
		}
	}

	if scon != nil {
		scon.Emit("temp", strconv.Itoa(temp))
		scon.Emit("volt", 12.60)
	}
}

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		scon = s
		return nil
	})

	server.OnEvent("/", "pos", func(s socketio.Conn, msx int, msy int) {
		fmt.Printf("msx:%d msy:%d\n", msx, msy)
	})

	server.OnEvent("/", "light", func(s socketio.Conn, togle int) {
		fmt.Printf("light:%d\n", togle)
	})

	server.OnEvent("/", "power", func(s socketio.Conn, val int) {
		fmt.Printf("power:%d\n", val)
		exec.Command("sudo", "shoutdown", "now")
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		scon = nil
	})

	go server.Serve()
	defer server.Close()

	// 5 second ticker
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	go func() {
		for {
			<-ticker.C
			reportLoop()
		}
	}()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./www")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
