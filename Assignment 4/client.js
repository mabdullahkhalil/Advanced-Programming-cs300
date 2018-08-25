const socket = io()
let gamenumber=0
let playerNumber=0

const clickdiv = () => {
	socket.emit('myalert', gamenumber+"@"+"abdul")
	console.log('clicked on play')
}

const handleMouse = event => {
	event.preventDefault()
	const btn_id = event.target.id
	document.getElementById(btn_id).style.boxShadow=  '20px 10px 20px red';
	document.getElementById(btn_id).style.outlineOffset =  '15px';
	document.getElementById(btn_id).style.border= '10px solid';
	document.getElementById(btn_id).style.borderColor='green';
	socket.emit('hoverCheck',gamenumber+"@"+event.target.id+"@"+playerNumber)
}


const mouseOut = event => {
	const btn_id = event.target.id
	document.getElementById(btn_id).style.boxShadow=  '0px 0px 0px ';
	document.getElementById(btn_id).style.border= '';

	const col = btn_id.split(' ')[1]


	// for (var i = 0; i < 6; i++) {
	// 	const id= document.getElementById(`${i} ${col}`)
	// 	//console.log(id.id)

	// 	if (id.innerHTML != '.') {
	// 		id.style.background = "white"
	// 	}
		
	// }
}


const handleClick = event => {
	console.log(`this is: ${event.target.id}`);
	socket.emit('click',gamenumber+"@"+event.target.id+"@"+playerNumber)
}


ReactDOM.render(
	React.createElement('button',
	{	
		id : 'bttn',
		onClick: () => clickdiv(),
		style: {
			fontSize : '100px',
			background : 'red',
			color  :'yellow',
			position : 'fixed',
			margin: '-100px 0 0 -100px',
			top : '70%',
			left : '50%'

		}
	}, 'PLAY' ),
	document.getElementById('root'))

socket.on('colorBox',data=> {
	const work = data.split('---')
	document.getElementById(work[0]).style.background = work[1]

})

socket.on('gameNumber',data=>{
	gamenumber=data
	console.log(`my game number is  ${data}`)
})


socket.on('noColor',data => {
	console.log(data)
})




socket.on('winning',data =>{
	ReactDOM.render(
		React.createElement('center',{},
	React.createElement('img', {src : 'http://fairtranslation.co.za/wp-content/uploads/2016/05/youwon.gif',
style :{
			height: '500px',
			width : '1000px'
		}}))

	, document.getElementById('root'))

})


socket.on('losing',data =>{
	ReactDOM.render(
		React.createElement('center',{},
	React.createElement('img', {src : 'https://3.bp.blogspot.com/-b3Hw2NVCCgc/V8lpcgiVnRI/AAAAAAAAM-M/q82O0nB-iAIyHIFrRTQ0mNBehIuDFlyFwCLcB/s1600/you-lose-banner-sm-%2540x2.png',
style :{
			height: '500px',
			width : '1000px'
		}}))

	, document.getElementById('root'))
		//window.location.reload()

})


socket.on('hoverAnswer',data=>{
	const data1 = data.split('@')
	const col = data1[0].split(' ')[1]
		console.log(col)

	// let color1 = ""	
	// if (data1[1]=="red") {
	// 	color1 = "#ff8080"
	// }
	// else{
	// 	color1 = "#66ff99"
	// }

	// for (var i = 0; i < 6; i++) {
	// 	const id= document.getElementById(`${i} ${col}`)
	// 	console.log(id.id)
	// 	id.style.background = color1
	// }


})

socket.on('makeboard',data => {
	if (data=="you have to wait for second player") {
		console.log("waiting in progress....")
	}
	else { 
		playerNumber = data
		console.log(`your are connected as player ${data}`)

		const divs = []
		for (var j = 0; j < 6; j++) {
			const but = []
			for (var i = 0; i < 7; i++) {
				const buttons= React.createElement('button' ,
					{id : `${j} ${i}`,
					onClick : () => handleClick(event), 
					onMouseOver : () => handleMouse(event),
					onMouseOut : () => mouseOut (event),
					style : {
						height : '100px',
						width  :'100px',
						borderRadius : '50%',
						background : 'white'
					}} ,
					" ")

				but.push(buttons)
			}
			divs.push(React.createElement('div',{id : `row ${j}`},but))
		}	
		ReactDOM.render(
			React.createElement('center',{id : "game Board", style : {
				background : 'blue'
			}},divs),
			document.getElementById('root')
			

			)}	
	})


// redraw()

