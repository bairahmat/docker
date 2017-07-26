package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/kanciogo/kancio-chat/controllers"
	"github.com/gorilla/websocket"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
 //
 //func serveHome(w http.ResponseWriter, r *http.Request) {
 //	c := controllers.NewControlers(tpl)
 //	log.Println(r.URL)
 //	if r.URL.Path != "/" {
 //		http.HandleFunc("/", c.Login)
 //		return
 //	}
 //	if r.Method != "GET" {
 //		http.Error(w, "Method not allowed", 405)
 //		return
 //	}
	//
 //}

//websocket chat
var clients = make(map[*websocket.Conn]bool)
var broadcast =  make(chan Message)

var upgrader = websocket.Upgrader{}
type Message struct {
        Email    string `json:"email"`
        Username string `json:"username"`
        Message  string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r , nil)
	if err != nil {
		log.Fatal("Error", err)
	}
	defer ws.Close()
	//register a new global clients
	clients[ws]=true
	for {
		var msg Message
		 // Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error, cannot read JSON file: %v", err)
			delete (clients, ws)
			break
		}
		fmt.Println("Data JSON:", err)
		fmt.Println("Data Msg:", msg)
		// Send the newly received message to the broadcast channel
		broadcast <- msg
		fmt.Println("Data broadcast:", broadcast)
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		// send it out to  every clients that is currently connected
		for client := range clients{
			err := client.WriteJSON(&msg)
			if err != nil {
				log.Printf("Error, Cannot write JSON file: %v", err)
				client.Close()
				delete(clients, client)
			}
			fmt.Println("Baca data Json", err)
			fmt.Println("Baca data Msg", msg)
		}

	}
}

func main() {
	c := controllers.NewControlers(tpl)
	//chatting websocket
	http.HandleFunc("/ws", handleConnections)
	//http.HandleFunc("/chat", c.Chat)
	go handleMessages()
	 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", nil)
	 })
	http.HandleFunc("/home", c.Home)
	http.HandleFunc("/daftar", c.Daftar)
	http.HandleFunc("/login", c.Login)

	//http.HandleFunc("/", serveHome)
	http.HandleFunc("/logout", c.Logout)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
