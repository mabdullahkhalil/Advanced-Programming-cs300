const fs = require('fs')
const http = require('http')
const socketio = require('socket.io')

const readFile = file => new Promise((resolve, reject) =>
    fs.readFile(file, (err, data) => err ? reject(err) : resolve(data)))

const delay = msecs => new Promise(resolve => setTimeout(resolve, msecs))

const server = http.createServer(async (request, response) =>
    response.end(await readFile(request.url.substr(1))))

const checkdata = data=> new Promise ((resolve,reject) =>{

	let checkarr =[0,0,0,0,0]
	for (var i = 0; i < 10; i++) {
		for (var j = 0; j < 10; j++) {
		 	if (data[i][j]==1) {
		 		checkarr[0]=1
		 	}
		}
	}
	for (var i = 0; i < 10; i++) {
	  for (var j = 0; j < 10; j++) {
		 	if (data[i][j]==2 && data[i][j+1]==2&& j+1<10) {
		 		checkarr[1]=1
		 		i=10
		 		j=10

		 	}
		 	else if(data[i][j]==2 && data[i+1][j]==2 && i+1<10){
		 		checkarr[1]=1
		 		i=10
		 		j=10
		 	}
		}
	}

	for (var i = 0; i < 10; i++) {
	  for (var j = 0; j < 10; j++) {
		 	if (data[i][j]==3 && data[i][j+1]==3 && data[i][j+2]==3 && j+2<10) {
		 		checkarr[2]=1
		 		i=10
		 		j=10

		 	}
		 	else if(data[i][j]==3 && data[i+1][j]==3 && data[i+2][j]==3 && i+2<10){
		 		checkarr[2]=1
		 		i=10
		 		j=10
		 	}
		}
	}

	for (var i = 0; i < 10; i++) {
	  for (var j = 0; j < 10; j++) {
		 	if (data[i][j]==4 && data[i][j+1]==4 && data[i][j+2]==4 && data[i][j+3]==4 && j+3<10) {
		 		checkarr[3]=1
		 		i=10
		 		j=10

		 	}
		 	else if(data[i][j]==4 && data[i+1][j]==4 && data[i+2][j]==4 && data[i+3][j]==4 && i+3<10){
		 		checkarr[3]=1
		 		i=10
		 		j=10
		 	}
		}
	}
	for (var i = 0; i < 10; i++) {
	  for (var j = 0; j < 10; j++) {
		 	if (data[i][j]==5 && data[i][j+1]==5 && data[i][j+2]==5 && data[i][j+3]==5 && data[i][j+4]==5 && j+4<10) {
		 		checkarr[4]=1
		 		console.log("yayy")
		 		i=10
		 		j=10

		 	}
		 	else if(data[i][j]==5 && data[i+1][j]==5 && data[i+2][j]==5 && data[i+3][j]==5 && data[i+4][j]==5 && i+4<10){
		 		checkarr[4]=1
		 		i=10
		 		j=10
		 	}
		}
	}

	let count =0;
	while(count<5){
		if (checkarr[count]==0) {
			reject()
		}
		count++
	}
	resolve()


})

const gamestat={}
const games= {}
const socklist =[]
const io = socketio(server)

io.sockets.on('connection', socket => {
    console.log('a client connected')
    socklist.push(socket)

    if (!gamestat[Math.ceil((socklist.length)/2)]) {
    	gamestat[Math.ceil((socklist.length)/2)] = [socket]
    	socket.emit('gamenumber', Math.ceil((socklist.length)/2))
    }
    else{
    	gamestat[Math.ceil((socklist.length)/2)].push(socket)
    	socket.emit('gamenumber', Math.ceil((socklist.length)/2))

    }


    socket.on('mydata', data => {
    	const gamenum = data[data.length -1]

    	if (gamestat[gamenum].length == 1) {


    	const arr=[]
    	for (var i = 0; i < data.length-1; i++) {
    		arr.push(data[i])
    	}
    	console.log(arr)

    	checkdata(arr).then(() => {
    		gamestat[gamenum][0].emit('validdata',"your data is valid")
    	    gamestat[gamenum][0].emit('connect',"you have to wait for the second palyer")}).catch(() => gamestat[gamenum][0].emit('invalid',"your data is invalid"))
    	}
    	else{
    		const arr=[]
    	for (var i = 0; i < data.length-1; i++) {
    		arr.push(data[i])
    	}
    	console.log(arr)

    	checkdata(arr).then(() => {
    					gamestat[gamenum][1].emit('validdata',"your data is valid")
    					gamestat[gamenum][0].emit('connect',"you are now connected to play")
    					gamestat[gamenum][1].emit('connect',"you are now connected to play")}).catch(() => gamestat[gamenum][1].emit('invalid',"your data is invalid"))
    	}


    })




    socket.on('disconnect', () => console.log('a client disconnected'))
})

server.listen(8000)
