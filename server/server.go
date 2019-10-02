package server

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

const (
	CliAddChanSize  = 20
	CliDelChanSize  = 20
	MessageChanSize = 1000
)

type HttpServer struct {
	wsServer *WsServer
}

type WsServer struct {
	Clients map[string][]*Client
	mutex   sync.RWMutex
	AddCli  chan *Client
	DelCli  chan *Client
	Message chan *Message
}

type Client struct {
	UserId    string
	Timestamp int64
	conn      *websocket.Conn
	wsServer  *WsServer
}

type Message struct {
	UserId  string `json:"user_id"`
	Message string `json:"message"`
}

func NewWsServer() *WsServer {
	return &WsServer{
		Clients: make(map[string][]*Client),
		AddCli:  make(chan *Client, CliAddChanSize),
		DelCli:  make(chan *Client, CliDelChanSize),
		Message: make(chan *Message, MessageChanSize),
	}
}

func NewHttpServer(wsServer *WsServer) *HttpServer {
	return &HttpServer{wsServer}
}

func (httpServer *HttpServer) SendMessage(userId, message string) {
	log.Printf("message reveived, user_id: %s, message: %s", userId, message)
	httpServer.wsServer.Message <- &Message{userId, message}
}

func (wsServer *WsServer) SendMessage(userId, message string) {
	// LOCK
	wsServer.mutex.RLock()
	defer wsServer.mutex.RUnlock()

	clients := wsServer.Clients[userId]
	if len(clients) <= 0 {
		log.Printf("client not found, user_id: %s", userId)
		return
	}
	for _, c := range clients {
		if nil != c {
			// async message send
			go c.sendMessage(userId, message)
		}
	}
	log.Printf("message success sent to client, user_id: %s", userId)
}

func (wsServer *WsServer) addClient(c *Client) {
	// LOCK
	wsServer.mutex.Lock()
	defer wsServer.mutex.Unlock()

	clients := wsServer.Clients[c.UserId]
	wsServer.Clients[c.UserId] = append(clients, c)
	log.Printf("a client added, userId: %s, timestamp: %d", c.UserId, c.Timestamp)
}

func (wsServer *WsServer) delClient(c *Client) {
	// LOCK
	wsServer.mutex.Lock()
	defer wsServer.mutex.Unlock()

	clients := wsServer.Clients[c.UserId]
	if 0 == len(clients) {
		delete(wsServer.Clients, c.UserId)
	} else if len(clients) > 0 {
		for i, client := range clients {
			if client.Timestamp == c.Timestamp {
				wsServer.Clients[c.UserId] = append(clients[:i], clients[i+1:]...)
				break
			}
		}
	}
	log.Printf("a client deleted, user_id: %s, timestamp: %d", c.UserId, c.Timestamp)
}

func (c *Client) sendMessage(userId, message string) {
	if nil != c.conn {
		c.conn.Write([]byte(message))
	}
}

func (wsServer *WsServer) Start() {
	for {
		select {
		case msg := <-wsServer.Message:
			wsServer.SendMessage(msg.UserId, msg.Message)
		case c := <-wsServer.AddCli:
			wsServer.addClient(c)
		case c := <-wsServer.DelCli:
			wsServer.delClient(c)
		}
	}
}

func (c *Client) heartbeat() error {
	millis := time.Now().UnixNano() / 1e6
	heartbeat := struct {
		Heartbeat int64 `json:"heartbeat"`
	}{millis}
	bytes, _ := json.Marshal(heartbeat)
	_, err := c.conn.Write(bytes)
	return err
}

func (c *Client) Listen() {
	for range time.Tick(5 * time.Second) {
		err := c.heartbeat()
		if nil != err {
			log.Printf("client heartbeat error, user_id: %v, timestamp: %d, err: %s", c.UserId, c.Timestamp, err)
			c.wsServer.DelCli <- c
			return
		}
	}
}
