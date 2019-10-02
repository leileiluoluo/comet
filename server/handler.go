package server

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

const UserId = "user_id"

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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Error{"params_invalid", "params invalid"})
		return
	}
	h.httpServer.SendMessage(m.UserId, m.Message)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func (h *Handler) CometHandler(conn *websocket.Conn) {
	userId := conn.Request().URL.Query().Get(UserId)
	defer conn.Close()
	if "" != strings.TrimSpace(userId) {
		millis := time.Now().UnixNano() / 1e6
		c := &Client{userId, millis, conn, h.wsServer}
		h.wsServer.AddCli <- c
		c.Listen()
	}
}
