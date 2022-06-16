package webSockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	db "github.com/vendenta/database"
	"github.com/vendenta/models"
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

	pingPeriod = (pongWait * 9) / 10

	//writeWait = 10 * time.Second

	maxMessageSize = 512
)

type chanelsQuiz struct {
	conn            *websocket.Conn
	timer           time.Timer
	winChan         chan int
	UpdateScoreChan chan int
}

func NewchanelsQuiz(connection *websocket.Conn, tm time.Timer, win chan int, update chan int) *chanelsQuiz {
	return &chanelsQuiz{
		conn:            connection,
		timer:           tm,
		winChan:         win,
		UpdateScoreChan: update,
	}
}

func (w *chanelsQuiz) timerQuiz() {

	puntaje := 0
	_ = puntaje

	pingTicker := time.NewTicker(pingPeriod)

	for {

		select {
		case <-w.timer.C:

			data, _ := json.Marshal("se acabo el tiempo... puntaje: " + strconv.Itoa(puntaje))

			err := w.conn.WriteMessage(websocket.TextMessage, data)

			if err != nil {
				log.Println("Error on write message: ", err.Error())
			}

			w.conn.Close()

			pingTicker.Stop()

			return

		case coins := <-w.winChan:

			data, _ := json.Marshal("ganaste... puntaje: " + strconv.FormatInt(int64(coins), 10))

			w.conn.WriteMessage(websocket.TextMessage, data)
			
			w.conn.Close()

			w.timer.Stop()

			pingTicker.Stop()

			return

		case sf := <-w.UpdateScoreChan:
			puntaje = sf

		case <-pingTicker.C:
			//w.conn.SetWriteDeadline(time.Now().Add(writeWait))  //esta vaina establece el tiempo limite para que escriba, si no lo hace se jode el ws
			//fmt.Println("ping")	
			if err := w.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				fmt.Println("ping: ", err)		
			}
		
		}

	}

}

func getANS(order int, quest []*models.Questions) (map[string]bool, string) {

	ans := make(map[string]bool)

	options := ""

	for _, n := range quest[order].Answer {

		ans[n.Answer] = n.Correct

		options = options + " " + n.Answer

	}

	return ans, options
}

func HandlerWebSocket(rw http.ResponseWriter, r *http.Request) {

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

	win := make(chan int)

	scoreChan := make(chan int)

	timer := time.NewTimer(15 * time.Second)

	var score int = 0

	quizChanels := NewchanelsQuiz(connection, *timer, win, scoreChan)

	go quizChanels.timerQuiz()

	ID := r.URL.Query().Get("id")

	quiz := db.SearshQuizDB(ID)

	question := quiz.Questions

	first := question[0]

	orden := 0

	ans, options := getANS(orden, question)

	data, _ := json.Marshal(first.Question + "   opciones: (" + options + " )")

	connection.WriteMessage(websocket.TextMessage, data)

	for {

		if _, message, err := connection.ReadMessage(); err != nil {
			log.Println("Error on read message: ", err.Error())
			break
		} else {

			if ans_final, ok := ans[string(message)]; ok && ans_final {

				score = score + question[orden].Point

				scoreChan <- score

				orden++

				if orden >= len(question) {

					win <- score

					break

				} else {

					ans, options = getANS(orden, question)

					data, _ := json.Marshal(question[orden].Question + "   opciones: (" + options + " )")

					connection.WriteMessage(websocket.TextMessage, data)

				}

			}

		}

	}
}
