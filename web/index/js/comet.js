// 'user strict'

var ws;

function login() {
	server = document.getElementById('server').value || document.getElementById('server').getAttribute('placeholder')
	userId = document.getElementById('user_id').value
	if (!server || !userId) {
		printAlert('ws address, user_id can not be null');
		return
	}
	document.getElementById('msg').innerHTML = '</br>';
	document.getElementById('chat').innerHTML = '';
	connect(server.trim(), userId.trim())
}

function logout() {
	userId = document.getElementById('user_id').value
	if (!userId) {
		printAlert('ws address, user_id can not be null');
		return
	}
	if (null != ws) {
		var logout = '{"action":"logout"}'
		ws.send(logout);
	}
}

function connect(server, userId) {
	try {
		ws = new WebSocket('ws://' + server + '?user_id=' + userId)
		counter = 0
		ws.onopen = function(e) { 
		    printAlert('connected to ws server ' + server + ' user_id: ' + userId);
		}; 
		ws.onclose = function(e) { 
			ws.close()
			ws = null
		    printAlert('connection closed')
		}; 
		ws.onmessage = function(e) { 
		    printMessage(counter++, e.data)
		}; 
		ws.onerror = function(e) { 
		    printAlert('connect error');
		};
	} catch (e) {
		printAlert('connect error, please check your ws server address');
	}
	
}

function printAlert(msg) {
	document.getElementById('msg').innerHTML = msg;
}

function printMessage(counter, msg) {
	document.getElementById('chat').innerHTML += 'received message:' + counter + ' ' + msg + '</br>';
}