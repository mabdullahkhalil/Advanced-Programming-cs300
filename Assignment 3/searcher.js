// Muhammad Abdullah Khalil
// 19100142






const fs=require ("fs")
const path= require ("path")

const dirname= process.argv[4]
const jsonfile= process.argv[3]
const task = process.argv[2]
const dirname2= "/Users/mabdullahk/Desktop/"+jsonfile

myJson = {}
var files=[]
var onefile=[]
var ans=[]


const readFile = file => new Promise ((resolve,reject) => 
	fs.readFile(file,'utf8',(err,data) => err ? reject (err) : resolve(file+"\n"+data)))

const readFile2 = file => new Promise ((resolve,reject) => 
	fs.readFile(file,'utf8',(err,data) => err ? reject (err) : resolve(data)))



const readDir = file => new Promise ((resolve,reject) => 
	fs.readdir(file,(err,list)  => err ? reject (err) : resolve(list)))


const isDir = file => new Promise ((resolve,reject) => 
	fs.lstat(file,(err,stats)  => err ? reject (err) : resolve(stats.isDirectory())))


const writeFile = (file,data) => new Promise ((resolve,reject) => 
	fs.writeFile(file,data,err  => err ? reject (err) : resolve()))


const readFiles= (dirname) =>
new Promise (
	(resolve) =>   {
	readDir(dirname).then(filenames=>{
		for (var i = filenames.length-1 ; i >= 0; i--) {
			var name=dirname+'/'+filenames[i];
			if(name.includes(".")){
  			if (name.includes(".txt")){
  				//files.push(name)
  			//	console.log(`hellllodaofjadndfn.      ${files[(files.length)-1]}`)
  				readFile(name).then((data) =>{
  									var lines = data.split('\n')
  				  					name=lines[0]
  				  					lines.shift()
  				  					onefile =[]
  				  				checkWordsInFile(name,lines).then(onefile=>{
  				  					console.log(`Done Reading File ${name}`)
  				  				}).catch(err => console.log ("error"))
  				  })
  			}}

  			else{
  				readFiles(name)
  			}
		}
	}).catch(err =>
	console.log(`Reading Directory error`))
})

const findword = (word,name,linee) =>
new Promise (
	(resolve)=> {
		for(i in myJson){
			if(i == word){
			//console.log(`${word} ---- ${linee}`)
				inserting (name,word,linee).then(resolve(word))
				return;
			}
		}
	myJson[word]= ([{file: name , lines:[linee]}])
	resolve(word)

})

const inserting = (name,word,linenu) => 
new Promise (
	(resolve) => {
		for ( i in myJson[word]){

			if (myJson[word][i].file == name) {

				myJson[word][i].lines.push(linenu)
				// console.log(myJson[word][i])
				resolve(myJson)
				return;
			}
		}

		myJson[word].push({file: name , lines:[linenu]})
		resolve(myJson)
		return;


	})





const makingJSon = (name,lines) => 
new Promise (
	(resolve) => {
		for (var i = 0; i < lines.length ; i++) {
			var one_line = lines[i].split(' ') //get words
			for (var j = 0; j < one_line.length ; j++) {
				// console.log(one_line[j])
				word = one_line[j] //get individual word
				if (word.length > 3)
					{findword(word,name,i+1).then(ans => {
									
				
			})}
			}
		}
		resolve(myJson)


})





const checkWordsInFile = (name,lines) => 
new Promise(
	(resolve,reject) => {
		if (true)
		{
		makingJSon(name,lines).then(myJsonn => {
			writeFile(jsonfile,JSON.stringify(myJsonn)).catch(err=>console.log("can't write"))
			resolve(name)
		})

		} 
		else {
			const msg= new Error ('empty everything')
			if (onefile.length == 0) {
				reject(msg)
			}
			reject(msg)
		}



})

let obj={}


const printfile = (word, path,list) => new Promise (
resolve => {
	readFile2(path).then(data => {
		const alllines=data.split('\n')
		for (var i = 0; i < list.length; i++) {
			let m=list[i]
			console.log(`${word} : ${path} : ${alllines[m-1]}`)
		}
	})

	
})


const objectfile = (word, obj) =>
new Promise (
	resolve => {
		let count=obj.length
		for (var i = obj.length - 1; i >= 0; i--) {
			 printfile(word, obj[i].file,obj[i].lines).then(count--)
			}
			count || resolve()
})




const search = (list_of_words,JSONfilename) => new Promise(
	(resolve) => {
	readFile2(JSONfilename).then(data=>{
		obj = JSON.parse(data)
		let list = list_of_words.slice(1,(list_of_words.length-1))
		let finlist = list.split(",")
		console.log(finlist)
		for (var i = 0; i < finlist.length; i++) {
			for (j in obj){
				if (j == finlist[i]) {
					
					objectfile(finlist[i], obj[j])

				}
			}
		}

	}).then(resolve()).catch(err => console.log(`${err} occured`))

	})


if (task == "index")
{readFiles(dirname).then(()=>console.log("ALL DOEN!"))
}
if (task == "search") {
search(dirname,dirname2).then(()=>console.log("all printed"))
}














