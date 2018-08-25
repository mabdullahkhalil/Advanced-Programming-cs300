const socket = io()

const delay = secs => new Promise(resolve => setTimeout(resolve, 1000*secs))

const shipsize = {
	'aircraft_carrier': 5,
	'battleship': 4,
	'cruiser': 3,
	'destroyer': 2,
	'submarine': 1
}
const state = {}

// const shipnumber = 

let shipsizeturn = 5


const checkHorizontal=  (row,column) => new Promise ((resolve,reject) =>{

	let col=column
	let count1=0
	let count2=0

	//cehcking any duplicates
	for (var i = 0; i < state.grid.length; i++) {
		for (var j = 0; j < 10; j++) {
			if (state.grid[i][j] == shipsizeturn) {
				state.grid[i][j]=0
			}
		}
	}
	// cehcking cols are zero
	while(count2 < shipsizeturn){
		if(state.grid[row][col] == 0){
			count1++
		}
		col++
		count2++
	}

	let count=0
	let c =column
	let last_col=c+shipsizeturn-1
 	while(column <= last_col && count1 == shipsizeturn){
		state.grid[row][column] = shipsizeturn
		column ++;
		count++;
 	}

	if (count == shipsizeturn+10) {
		reject()
	}
	resolve()

})

const checkVertical=  (row,column) => new Promise ((resolve,reject) =>{

	let col=row
	let count1=0
	let count2=0

	//cehcking any duplicates
	for (var i = 0; i < state.grid.length; i++) {
		for (var j = 0; j < 10; j++) {
			if (state.grid[i][j] == shipsizeturn) {
				state.grid[i][j]=0
			}
		}
	}
	// cehcking cols are zero
	while(count2 < shipsizeturn){
		if(state.grid[col][column] == 0){
			count1++
		}
		col++
		count2++
	}

	let count=0
	let r =row
	let last_row=r+shipsizeturn-1
 	while(row <= last_row && count1 == shipsizeturn){
		state.grid[row][column] = shipsizeturn
		row++;
		count++;
 	}

	if (count == shipsizeturn+10) {
		reject()
	}
	resolve()

})


const handleClick = event => {
	let row = Number(event.target.id.split(' ')[1])

	let column = Number(event.target.id.split(' ')[2])

	checkHorizontal(row,column).then(()=>	setState(state.grid)).catch(()=>setState(state.grid))
	

	
}

const handleRightClick = event => {
	event.preventDefault()
	let row = Number(event.target.id.split(' ')[1])

	let column = Number(event.target.id.split(' ')[2])

	checkVertical(row,column).then(()=>	setState(state.grid)).catch(()=>setState(state.grid))
	

}

const changedEle = event => {
	shipsizeturn = shipsize[event.target.value]
}

const sendData = ()=>{
	console.log("data sent")
	socket.emit('mydata', state.grid)
}


const setState = updates => {
	Object.assign(state, updates)
	const rows= []
	for (var j = 0; j < state.grid.length ; j++){
		const divs = []
		for (var i = 0; i < 10; i++) {
			if (state.grid[j][i]==0){
				divs.push(React.createElement('div',{className : 'box',id : `M ${j} ${i}`, onClick : () => handleClick(event), onContextMenu : ()=> handleRightClick(event)}))
			}
			if (state.grid[j][i]==5){
				divs.push(React.createElement('div',{className : 'box ship-aircraft_carrier',id : `M ${j} ${i}`, onClick : () => handleClick(event), onContextMenu : ()=> handleRightClick(event)}))
			}
			if (state.grid[j][i]==4){
				divs.push(React.createElement('div',{className : 'box ship-battleship',id : `M ${j} ${i}`, onClick : () => handleClick(event), onContextMenu : ()=> handleRightClick(event)}))
			}
			if (state.grid[j][i]==3){
				divs.push(React.createElement('div',{className : 'box ship-cruiser',id : `M ${j} ${i}`, onClick : () => handleClick(event), onContextMenu : ()=> handleRightClick(event)}))
			}
			if (state.grid[j][i]==2){
				divs.push(React.createElement('div',{className : 'box ship-destroyer',id : `M ${j} ${i}`, onClick : () => handleClick(event),  onContextMenu : ()=> handleRightClick(event)}))
			}
			if (state.grid[j][i]==1){
				divs.push(React.createElement('div',{className : 'box ship-submarine',id : `M ${j} ${i}`, onClick : () => handleClick(event), onContextMenu : ()=> handleRightClick(event)}))
			}

		}
		rows.push(React.createElement('div',{id :`${j}`},divs))

	}
		const rows1= []
	for (var j = 0; j < state.grid.length ; j++){
		const divs1 = []
		for (var i = 0; i < 10; i++) {
			if (state.guessboard[j][i]==0){
				divs1.push(React.createElement('div',{className : 'box',id : `G ${j} ${i}`}))
			}
		}
		rows1.push(React.createElement('div',{id :`${j}`},divs1))

	}
const options=[
			React.createElement('option',{},'aircraft_carrier'),
			React.createElement('option',{},'battleship'),
			React.createElement('option',{},'cruiser'),
			React.createElement('option',{},'destroyer'),
			React.createElement('option',{},'submarine')
			]
	const dropdown=	React.createElement('select',{
		onChange: ev => changedEle(ev)
	},options)

	ReactDOM.render(React.createElement('div', null, rows,dropdown, rows1, 
		React.createElement('button',{onClick : ()=>sendData()}, "START GAME"),
		state.msg.map(m=> React.createElement('div',{},m))),
		document.getElementById('root'))
}



const array = [0,0,0,0,0,0,0,0,0,0]
const make =[]
	for (var i = 0; i < 10; i++) {
		make[i] = [0,0,0,0,0,0,0,0,0,0]
	}


setState({grid :make, guessboard: make, msg:[]})

socket.on('validdata',(data)=>{
	console.log(data)
	state.msg.push(data)
	setState(state)
})

socket.on('invalid',(data)=>{
	console.log(data)
	state.msg.push(data)
	setState(state)
})