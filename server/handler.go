package server

import (
	"net/http"
	"encoding/json"
	"golang.org/x/net/websocket"
	"time"
)

type Error struct {
	Code   string `json:"code"`
	Reason string `json:"reason"`
}

type Handler struct {
	wsServer   *WsServer
	httpServer *HttpServer
}

func NewHandler(wsServer *WsServer, httpServer *HttpServer) *Handler {
	return &Handler{wsServer, httpServer}
}

func (h *Handler) MessageHandler(w http.ResponseWriter, r *http.Request) {
	m := &Message{}
	err := json.NewDecoder(r.Body).Decode(m)
	defer r.Body.Close()
	if nil != err {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&Error{"params_invalid", "params invalid"})
		return
	}
	h.httpServer.SendMessage(m.UserId, m.Message)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(m)
}

func (h *Handler) CometHandler(conn *websocket.Conn) {
	userId := conn.Request().URL.Query().Get("user_id")
	defer conn.Close()
	if len(userId) > 0 {
		millis := time.Now().UnixNano() / 1000000
		c := &Client{userId, millis, conn, h.wsServer}
		h.wsServer.AddCli <- c
		c.Listen()
	}
}
