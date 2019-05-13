# A light-weight Golang WebSocket Comet Server

You can use Comet for realtime peer-to-peer message push or broadcasting.

1) Run
$ go get github.com/olzhy/comet
$ go run main.go

2) API for Browser
You can use JavaScript WebSocket API to establish a connection to Comet server.
The address is ws://localhost:8080/comet?user_id=:user_id
See the sample page http://localhost:8080
Then input a user ID such as 1.

3) Send message
You can send a message to the browser user client in step 2.
Method is post, body is 'aplication/json' format.
user_id and message param required.
such as:
curl -d '{"user_id": "x", "message": "test"}' http://localhost:8080/messages
Then the browser user client will receive this message.
