# A light-weight Golang WebSocket Comet Server

You can use Comet for realtime peer-to-peer message push or broadcasting.

<div align=center><img width="664" height="261" src="https://leileiluoluo.com/wp-content/uploads/2018/09/comet-api.png"/></div><br/>
<div align=center><img width="408" height="220" src="https://leileiluoluo.com/wp-content/uploads/2018/09/comet-heartbeat.png"/></div>

## 1) Run 
```Bash
$ go get github.com/olzhy/comet
$ go run main.go
```

## 2) API for Browser
You can use JavaScript WebSocket API to establish a connection to Comet server.
The address is:
```JavaScript
ws://localhost:8080/comet?user_id=:user_id
```
See the sample page:
```JavaScript
http://localhost:8080
```
Then input a user ID such as 1.

## 3) Send message
You can send a message to the browser user client in step 2.
Method is post, body is 'aplication/json' format.
user_id and message param required.
such as:
```Bash
curl -d '{"user_id": "x", "message": "test"}' http://localhost:8080/messages
```
Then the browser user client will receive this message.
