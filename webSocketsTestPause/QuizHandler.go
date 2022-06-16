package websocketstestpause

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func check(r *http.Request) bool {
	log.Printf("%s %s%s %v", r.Method, r.Host, r.RequestURI, r.Proto)
	return r.Method == http.MethodGet
}

var upGradeWebSocket = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     check,
}

const (
	pongWait = 60 * time.Second

	maxMessageSize = 512
)

type chanelscontrollers struct {
	Conn  *websocket.Conn
	Pause chan int
	Play  chan int
	Quit  chan int
}

func CreateChannelsDisplay(connection *websocket.Conn, pause chan int, play chan int, quit chan int) *chanelscontrollers {
	return &chanelscontrollers{
		Conn:  connection,
		Pause: pause,
		Play:  play,
		Quit:  quit,
	}
}

func (w *chanelscontrollers) displaying() {

	var i = 0

	for {

		select {

		case <-w.Play:
			continue

		case <-w.Pause:

			select {

			case <-w.Pause:
				continue
				
			case <-w.Play:
				continue

			case <-w.Quit:
				w.Conn.Close()
				close(w.Play)
				close(w.Pause)
				close(w.Quit)
				return
			}

		case <-w.Quit:
			w.Conn.Close()
			return

		default:
			time.Sleep(1000 * time.Millisecond)
			i++
			data, _ := json.Marshal(i)
			w.Conn.WriteMessage(websocket.TextMessage, data)
		}

	}

}

func HandlerWebSocketPause(rw http.ResponseWriter, r *http.Request) {

	connection, err := upGradeWebSocket.Upgrade(rw, r, nil)

	if err != nil {
		log.Println("No se abrió la connextion")
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Connection failed.")
		return
	}

	connection.SetReadLimit(maxMessageSize)
	connection.SetReadDeadline(time.Now().Add(pongWait))
	connection.SetPongHandler(func(string) error { connection.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	play := make(chan int)
	pause := make(chan int)
	quit := make(chan int)

	Chanels := CreateChannelsDisplay(connection, play, pause, quit)

	go Chanels.displaying()

	data, _ := json.Marshal("bienvenido!")

	connection.WriteMessage(websocket.TextMessage, data)

	for {

		if _, message, err := connection.ReadMessage(); err != nil {
			log.Println("Error on read message: ", err.Error())
			break
		} else {

			switch string(message) {

			case "play":

				Chanels.Play <- 0

			case "pause":

				Chanels.Pause <- 0

			case "quit":

				Chanels.Quit <- 0

			}

		}

	}
}
