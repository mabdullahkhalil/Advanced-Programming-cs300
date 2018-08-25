const fs = require('fs')
const http = require('http')
const socketio= require('socket.io')

const readFile = f => new Promise((resolve, reject) =>
	fs.readFile(f, (e, d) => e?reject(e):resolve(d)))

const server = http.createServer(async (req,resp) =>
	resp.end(await readFile(req.url.substr(1))))


const gamestat = {}
const playerslist = {}
const gameDecision ={}
const playerturn ={}


const io = socketio(server)
const socklist =[]

const clickDecision = (col , game) => new Promise (
	(resolve,reject) => {
		for (var i = game.length - 1; i >= 0; i--) {
			if(game[i][col]==0){
				resolve(`${i} ${col}`)
			}
		}
		reject()
	})



const checkHorizontal = (check,player) => new Promise ((resolve,reject) => {
	for (var i = 0; i < 6; i++) {
		let count = 0
		for (var j = 0; j < 6; j++) {
			if (check[i][j] == check[i][j+1] && check[i][j]==player) {
				count++
			}
			else{
				count=0
			}
			if (count==3) {
				// console.log('horizontal',count)
				resolve()

			}
		}


	}
	reject([check,player])
	
})


const checkVertical = (check,player) => new Promise ((resolve,reject) => {
	for (var i = 0; i < 7; i++) {
		let count = 0
		for (var j = 0; j < 5; j++) {
			if (check[j][i] == check[j+1][i] && check[j][i]==player) {
				count++
			}
			else{
				count=0
			}
			if (count==3) {
				// console.log("found",i)
				resolve()

			}
		}


	}
	reject([check,player])
})


const checkDiag = (check,player) => new Promise((resolve,reject)=> {
	const colLength = check[0].length
	const rowLength = check.length

	for (var i = 0; i < colLength; i++) {
		for (var j = 0; j < rowLength; j++) {
			let jj = i+1
			let ii = j+1
			let count=0
			while(ii < rowLength && jj < colLength){
				if (check[ii-1][jj-1] == check[ii][jj] && check[ii][jj]==player) {
					count++
				}
				else{
					count=0
				}
				ii++
				jj++
			}
			if (count==3) {
				resolve()
			}
		}	
	}


	for (var i = 0; i < rowLength; i++) {
		for (var j = 1; j < colLength; j++) {
			let jj = i+1
			let ii = j+1
			let count=0
			while(ii < rowLength && jj < colLength){
				if (check[ii-1][jj-1] == check[ii][jj] && check[ii][jj]==player) {
					count++
				}
				else{
					count=0
				}
				ii++
				jj++
			}
			// console.log("count", count)
			if (count==3) {
				resolve()
			}
		}
		
	}


	for (var i = colLength; i >= 0; i--) {
		for (var j = 0; j < rowLength; j++) {
			let ii=i-1//col
			let jj=j+1 // row
			let count=0
			while(jj < rowLength && ii >= 0) {
				if (check[jj-1][ii+1] == check[jj][ii] && check[jj][ii]==player) {
					count++
				}
				else{
					count=0
				}
				ii--
				jj++
			}
			if (count==3) {
				resolve()
			}
		}	
	}

	

	reject([check,player])

})


const checkarr= (original_game,arr,count,player) => 
Promise.all(arr.map(i => new Promise (resolve => {
	// let count1= ;
	// let count=0;
	clickDecision(i,original_game).then(data1=>{
		// console.log("clickDecision",data1)
		const row_col1 = data1.split(' ')
		original_game[row_col1[0]][row_col1[1]]=player
	// console.log(original_game)

		checkHorizontal(original_game,player)
		.then(()=>{	
			// console.log("hori")
			count--;
			resolve(count)
		})
		.catch(r => {
			checkVertical(r[0],r[1])
			.then(()=>{		
				// console.log("vert")
		
				count--;
				resolve(count)

			})
			.catch(p => {
				checkDiag(p[0],p[1])
				.then(() => {
					// console.log("diag")

					count--;
					resolve(count)
				})
				.catch (()=> {

					// console.log("adding")
					count++;
					resolve(count)


				})
			})
		})

		original_game[row_col1[0]][row_col1[1]]=0

	}).catch(()=>(resolve(count)))

})))




io.sockets.on('connection', socket => {
	socklist.push(socket)
	console.log(`${socklist.indexOf(socket)} -- ${socket.id} connected.`)
	// playerslist[socket]=Math.ceil((socklist.length)/2)

	if (!gamestat[Math.ceil((socklist.length)/2)]) {
		gamestat[Math.ceil((socklist.length)/2)] = [socket]
		socket.emit('gameNumber',Math.ceil((socklist.length)/2))
	}
	else{
		gamestat[Math.ceil((socklist.length)/2)].push(socket)
		socket.emit('gameNumber',Math.ceil((socklist.length)/2))


	}

	socket.on('myalert', data => {
		const gamenum = Number(data.split('@')[0])

		if(gamestat[gamenum].length == 1){
			gamestat[gamenum][0].emit('makeboard',"you have to wait for second player")
		}
		if (gamestat[gamenum].length ==2){
			playerturn[gamenum]=1
			gamestat[gamenum][0].emit('makeboard',1)
			gamestat[gamenum][1].emit('makeboard',2)
			const row = []
			for(var i=0 ; i < 6;i++){
				const column =[]
				for(var j=0 ; j < 7;j++){
					column.push(0)
				}
				row.push(column)
			}
			gameDecision[gamenum]=row


		}

	})


	socket.on('click', data => {
		// console.log("clicked")
		const gamenum = Number(data.split('@')[0])
		const rowcol = data.split('@')[1]
		const player = Number(data.split('@')[2])

		// console.log(rowcol)
		const col_clicked = Number(rowcol.split(" ")[1])
		clickDecision(col_clicked,gameDecision[gamenum])
		.then(d =>{
			const row_col = d.split(' ')
			if (player == 1 && playerturn[gamenum]==1){
				gameDecision[gamenum][row_col[0]][row_col[1]]=1 

				checkHorizontal(gameDecision[gamenum],1)
				.then(()=>{				
					gamestat[gamenum][0].emit('winning',"YOU WON !")
					gamestat[gamenum][1].emit('losing',"YOU LOST !")})
				.catch((r)=>{
					checkVertical(r[0],r[1])
					.then(()=>{				
						gamestat[gamenum][0].emit('winning',"YOU WON !")
						gamestat[gamenum][1].emit('losing',"YOU LOST !")})
					.catch((p)=>{checkDiag(p[0],p[1])
						.then(()=>{
							gamestat[gamenum][0].emit('winning',"YOU WON !")
							gamestat[gamenum][1].emit('losing',"YOU LOST !")

						})
						.catch(()=>{
							gamestat[gamenum][0].emit('colorBox',d+"---"+"yellow")
							gamestat[gamenum][1].emit('colorBox',d+"---"+"yellow")
							playerturn[gamenum]=2

						})})
				})

			}
			else if (player == 2 && playerturn[gamenum]==2){
				gameDecision[gamenum][row_col[0]][row_col[1]]=2

				checkHorizontal(gameDecision[gamenum],2)
				.then(()=>{				
					gamestat[gamenum][1].emit('winning',"YOU WON !")
					gamestat[gamenum][0].emit('losing',"YOU LOST !")})
				.catch((r)=>{
					// console.log("checked horizontal")
					checkVertical(r[0],r[1])
					.then(()=>{				
						gamestat[gamenum][1].emit('winning',"YOU WON !")
						gamestat[gamenum][0].emit('losing',"YOU LOST !")})
					.catch((p)=>{checkDiag(p[0],p[1])
						.then(()=>{
							gamestat[gamenum][1].emit('winning',"YOU WON !")
							gamestat[gamenum][0].emit('losing',"YOU LOST !")

						})
						.catch(()=>{
							// console.log(gamestat)
							gamestat[gamenum][1].emit('colorBox',d+"---"+"red")
							gamestat[gamenum][0].emit('colorBox',d+"---"+"red")
							playerturn[gamenum]=1

						})})
				})



			}	
			else {
				
				gamestat[gamenum][player-1].emit('noColor',"It is not your turn yet. waiting for the other player to take the turn")
				
			}



		})
		.catch(()=>{})

	})



	// socket.on('hover', data => {
	// 	const gamenum = Number(data.split('@')[0])
	// 	const rowcol = data.split('@')[1]
	// 	const player = Number(data.split('@')[2])

	// 	console.log(rowcol)
	// 	const col_clicked = Number(rowcol.split(" ")[1])
	// 	const original_game=gameDecision[gamenum]

	// 	clickDecision(col_clicked,original_game)
	// 	.then(d =>{
	// 		const row_col = d.split(' ')
	// 		console.log("row col", row_col)
	// 		if (player == 1 && playerturn[gamenum]==1){
	// 			original_game[row_col[0]][row_col[1]]=1 

	// 			for (var i = 0; i < 6; i++) {
	// 				clickDecision(i,original_game).then(data1=>{
	// 					const row_col1 = data1.split(' ')
	// 					original_game[row_col1[0]][row_col1[1]]=2


	// 					checkHorizontal(original_game,2)
	// 					.then(()=>{				
	// 						gamestat[gamenum][0].emit('hoverAnswer',"red")	

	// 					})
	// 					.catch((r)=>{
	// 						checkVertical(r[0],r[1])
	// 						.then(()=>{				
	// 							gamestat[gamenum][0].emit('hoverAnswer',"red")	

	// 						})
	// 						.catch((p)=>{checkDiag(p[0],p[1])
	// 							.then(()=>{
	// 								gamestat[gamenum][0].emit('hoverAnswer',"red")	


	// 							})
	// 							.catch(()=>{
	// 									gamestat[gamenum][0].emit('hoverAnswer',"green")	


	// 							})})
	// 					})
	// 				}).catch(()=>gamestat[gamenum][0].emit('hoverAnswer',"green"))

	// 			}
	// 		}
	// 		// else if (player == 2 && playerturn[gamenum]==2){
	// 		// 	original_game[row_col[0]][row_col[1]]=2

	// 		// 	for (var i = 0; i < 6; i++) {
	// 		// 		clickDecision(i,original_game).then(data1=>{
	// 		// 			const row_col1 = data1.split(' ')
	// 		// 			original_game[row_co1[0]][row_col1[1]]=2

	// 		// 			checkHorizontal(original_game,1)
	// 		// 			.then(()=>{				
	// 		// 				gamestat[gamenum][1].emit('hoverAnswer',"red")
	// 		// 				.catch((r)=>{
	// 		// 					checkVertical(r[0],r[1])
	// 		// 					.then(()=>{				
	// 		// 						gamestat[gamenum][1].emit('hoverAnswer',"red")
	// 		// 					})
	// 		// 					.catch((p)=>{checkDiag(p[0],p[1])
	// 		// 						.then(()=>{
	// 		// 							gamestat[gamenum][1].emit('hoverAnswer',"red")

	// 		// 						})
	// 		// 						.catch(()=>{
	// 		// 							gamestat[gamenum][1].emit('hoverAnswer',"green")

	// 		// 						})})
	// 		// 				})

	// 		// 			})
	// 		// 		})

	// 		// 	}
	// 		// }	

	// 	}).catch(()=>console.log("nothing found"))


	// })



	socket.on('hoverCheck', data => {
		const gamenum = Number(data.split('@')[0])
		const rowcol = data.split('@')[1]
		const player = Number(data.split('@')[2])

		const col_clicked = Number(rowcol.split(" ")[1])
		// const original_game=Array.from(gameDecision[gamenum])
		// console.log(original_game)


		var originalgame = JSON.parse(JSON.stringify(gameDecision))
		const original_game = originalgame[gamenum]

		clickDecision(col_clicked,original_game)
		.then(d =>{
			const row_col = d.split(' ')
			if (player == 1 && playerturn[gamenum]==1){
				original_game[row_col[0]][row_col[1]]=1
				const count=0
				const arr = [0,1,2,3,4,5,6]
				checkarr(original_game,arr,count,2).then(data => {
					if (data[data.length-1] == 7) {
						gamestat[gamenum][0].emit('hoverAnswer',rowcol+"@"+"green")
					}
					else{
						gamestat[gamenum][0].emit('hoverAnswer',rowcol+"@"+"red")

					}
				})
				// console.log("count",count)
			}	
			else if (player == 2 && playerturn[gamenum]==2){
				original_game[row_col[0]][row_col[1]]=2
				const count=0
				const arr = [0,1,2,3,4,5,6]
				checkarr(original_game,arr,count,1).then(data => {
					if (data[data.length-1] == 7) {
						gamestat[gamenum][1].emit('hoverAnswer',rowcol+"@"+"green")
					}
					else{
						gamestat[gamenum][1].emit('hoverAnswer',rowcol+"@"+"red")

					}
				})
				// console.log("count",count)
			}	


		}).catch(()=>{})

	})

})





server.listen(8032, () => console.log('Started...'))

