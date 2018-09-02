# comet
golang comet server

go run main.go

1) websocket 
ws://localhost:8080/comet?user_id=x

2) send message
curl -d '{"user_id": "x", "message": "test"}' http://localhost:8080/messages
