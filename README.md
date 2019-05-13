# A light-weight Golang WebSocket Comet Server

You can use Comet for realtime peer-to-peer message push or broadcasting.

The Comet API provide two different protocols for receiver and sender.
Message receiver using websocket protocal, as well as Messsage sender using HTTP protocal. 
<div align=center><img width="664" height="261" src="https://github.com/olzhy/comet/blob/master/comet-api.png"/></div><br/>
The design is very simple and light-weighted.
When a web browser client connected, the Comet server launch a goroutine simultaneously and monitor it.
Monitor corresponding to the client send a heartbeat message per 5 second in order to check the connection. Once the connection break, the Comet will remove this client.
So, in this way, the frontend will never pay attention to the Comet server and never should send heartbeat to server.
Control is in server side completely.
<div align=center><img width="408" height="220" src="https://github.com/olzhy/comet/blob/master/comet-heartbeat.png"/></div>

## 1) Run 
```Bash
$ go get github.com/olzhy/comet
$ go run main.go
```

## 2) API for web Browser
You can use JavaScript WebSocket API to establish a connection to Comet server.
The address is:
```JavaScript
ws://localhost:8080/comet?user_id=:user_id
```
Please give a arbitrary valid :user_id param, such as x.

## 3) Send message
You can send a message to the browser user client in step 2.
Method is post, body is 'aplication/json' format.
user_id and message param required.
such as:
```Bash
curl -d '{"user_id": "x", "message": "test"}' http://localhost:8080/messages
```
Then the user x you specified will receive the test message.

For more details, please visit my blog.
https://leileiluoluo.com/posts/golang-websocket-combine-consistent-hashing.html
