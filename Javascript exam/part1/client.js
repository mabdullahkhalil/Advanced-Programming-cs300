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

const setState = updates => {
	Object.assign(state, updates)
	const rows= []


	for (var j = 0; j < state.grid.length ; j++){
		const divs = []
		for (var i = 0; i < 10; i++) {
			if (state.grid[j][i]==0){
				divs.push(React.createElement('div',{className : 'box',id : `${j} ${i}`, onClick : () => handleClick(event)}))
			}
		}
		rows.push(React.createElement('div',{id : `${i}`},divs))

	}
	ReactDOM.render(React.createElement('div', null, rows), 
		document.getElementById('root'))
}



array = [[0,0,0,0,0,0,0,0,0,0]]
setState({grid :array})