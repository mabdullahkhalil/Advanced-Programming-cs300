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

const handleClick = event => {
	let row = Number(event.target.id.split(' ')[0])

	let column = Number(event.target.id.split(' ')[1])
	console.log(row,column)

	let col=column
	let count1=0
	let count2=0
	while(count2 < 5){
		if(state.grid[row][col]==5){
			count1++
		}
		col++
		count2++
	}
	let count=0
	console.log("count", count1)
	let c =column
	while(column < 10 && (10-c) > 4 && count1==0 && count < 5){
		console.log(row,column)
		state.grid[row][column] = 5
		// row++
		column ++;
		count++;
	}
	console.log(state.grid)
	setState(state.grid)
	
}




const setState = updates => {
	Object.assign(state, updates)
	const rows= []

	for (var j = 0; j < state.grid.length ; j++){
		const divs = []
		for (var i = 0; i < 10; i++) {
			if (state.grid[j][i]==0){
				divs.push(React.createElement('div',{className : 'box',id : `${j} ${i}`, onClick : () => handleClick(event)}))
			}
			if (state.grid[j][i]==5){
				divs.push(React.createElement('div',{className : 'box ship-aircraft_carrier',id : `${j} ${i}`, onClick : () => handleClick(event)}))
			}
		}
		rows.push(React.createElement('div',{id :`${j}`},divs))

	}
	ReactDOM.render(React.createElement('div', null, rows), 
		document.getElementById('root'))
}



const array = [0,0,0,0,0,0,0,0,0,0]
const make =[]
	for (var i = 0; i < 10; i++) {
		make[i] = [0,0,0,0,0,0,0,0,0,0]
	}


setState({grid :make})