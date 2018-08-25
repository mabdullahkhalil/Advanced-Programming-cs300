/*

Muhammad Abdullah Khalil
19100142

part 5 and 6:

there may be an error that says, index out of range, run the code again. 
i asked sir about this, he said some slice might get lost. 
*/



package main

import (
"fmt"
"os"
"strconv"
"math"
"encoding/csv"
"sync"
)

type CensusGroup struct {
	population int
	latitude, longitude float64
}

type Vertex struct {
	Lat, Long float64
}

func task(censusData []CensusGroup,min_lat1 chan float64,min_long1 chan float64,max_lat1 chan float64, max_long1 chan float64, donee chan int){
	var min_lat float64 = 1
	var min_long float64 = 0
	var max_lat float64 = 0
	var max_long float64 = -200

	for i := 0; i < len(censusData); i++ {
				// fmt.Println("hello1")
			if censusData[i].latitude < min_lat{
				min_lat = censusData[i].latitude

			}
			if censusData[i].latitude > max_lat {
				max_lat = censusData[i].latitude
			}
			if censusData[i].longitude < min_long{
				min_long = censusData[i].longitude
			}
			if censusData[i].longitude > max_long {
				max_long = censusData[i].longitude
			}
	}

	go func(){
		select{
			case checkminlat := <-min_lat1:
				if min_lat < checkminlat {
					go func(){min_lat1 <- min_lat}()
				} else{
					go func(){min_lat1 <- checkminlat}()
				}
			default:
				go func(){min_lat1 <- min_lat}()
		}
	}()

	go func(){
		select{
			case checkmaxlat := <-max_lat1:
				if max_lat > checkmaxlat {
					go func(){max_lat1 <- max_lat}()
				} else{
					go func(){max_lat1 <- checkmaxlat}()
				}
			default:
				go func(){max_lat1 <- max_lat}()
		}
	}()


	go func(){
		select{
			case checkminlong := <-min_long1:
				if min_long < checkminlong {
					go func(){min_long1 <- min_long}()
				} else{
					go func(){min_long1 <- checkminlong}()
				}
			default:
				go func(){min_long1 <- min_long}()
		}
	}()


	go func(){
		select{
			case checkmaxlong := <-max_long1:
				if max_long > checkmaxlong {
					go func(){max_long1 <- max_long}()
				} else{
					go func(){max_long1 <- checkmaxlong}()
				}
			default:
				go func(){max_long1 <- max_long}()
		}
	}()


	donee<-0

}

func v4Setup(censusData []CensusGroup, min_lat float64 , min_long float64, latIncrement float64, longIncrement float64,xdim int, ydim int, done chan int) [][] int{
	  mainArray := make([][]int, xdim)


		for i := 0; i < xdim; i++ {
		    mainArray[i] = make([]int, ydim)
		}

	for i := 0; i < len(censusData); i++ {
		thisdata := censusData[i]
		lat1 := int(math.Floor((thisdata.latitude - min_lat)/ latIncrement))
		long1 := int(math.Floor((thisdata.longitude - min_long )/ longIncrement ))
		if lat1 == -1 {
			lat1=0
		}
		if xdim-1-long1 == -1 {
			long1=0
		}
		if lat1 == ydim {
			lat1= lat1-1
		}

		// fmt.Println(xdim-1-long1 , lat1)
		mainArray[xdim-1-long1][lat1]=mainArray[xdim-1-long1][lat1] + thisdata.population
	}
	done <- 1
	return mainArray
}

func v2Query (censusData []CensusGroup, population chan int,done chan int, min_lat float64, min_long float64, west int, south int, east int, north int, latIncrement float64 , longIncrement float64 ){
	pop := 0
	for i := 0; i < len(censusData); i++ {
		thisdata := censusData[i]
		lat1 := int(math.Ceil((thisdata.latitude - min_lat)/ latIncrement))
		long1 := int(math.Ceil((thisdata.longitude - min_long )/ longIncrement ))
		if lat1 >= south && lat1 <= north && long1 >= west && long1 <= east {
			pop = pop + thisdata.population
		}
	}

	go func(){
		select{
			case checkpop := <-population:
				go func(){population <- pop+checkpop}()
			default:
				go func(){population <- pop}()
		}
	}()

	go func(){
		done <- 0
	}()
}


func taskv2(censusData []CensusGroup,min_lat1 chan float64,min_long1 chan float64,max_lat1 chan float64, max_long1 chan float64){

	if len(censusData) > 10000{
			mid := len(censusData)/2
			max_latitude1  := make(chan float64)
			max_latitude2  := make(chan float64)
			min_latitude1  := make(chan float64)
			min_latitude2  := make(chan float64)
			max_longitude1 := make(chan float64)
			max_longitude2 := make(chan float64)
			min_longitude1 := make(chan float64)
			min_longitude2 := make(chan float64)
			go taskv2(censusData[:mid], min_latitude1 ,min_longitude1 , max_latitude1 , max_longitude1)
			go taskv2(censusData[mid:], min_latitude2 ,min_longitude2 , max_latitude2 , max_longitude2)
			backmaxlat1  := <- max_latitude1
			backmaxlat2  := <- max_latitude2
			backmaxlong1 := <- max_longitude1
			backmaxlong2 := <- max_longitude2	
			backminlat1  := <- min_latitude1
			backminlat2	 := <- min_latitude2			
			backminlong1 := <- min_longitude1
			backminlong2 := <- min_longitude2
		fmt.Println("#################################")

			if backmaxlong1 < backmaxlong2 {
				go func(){max_long1 <- backmaxlong2}()
			}else {
				go func(){max_long1 <- backmaxlong1}()
			}

			
			if backmaxlat1 < backmaxlat2 {
				go func(){max_lat1 <- backmaxlat2}()
			}else {
				go func(){max_lat1 <- backmaxlat1}()
			}

			if backminlat1 < backminlat2 {
				go func(){min_lat1 <- backminlat1}()
			}else {
				go func(){min_lat1 <- backminlat2}()
			}

			if backminlong1 < backminlong2 {
				go func(){min_long1 <- backminlong1}()
			}else {
				go func(){min_long1 <- backminlong2}()
			}

	}else{
		var min_lat float64 = 1
		var min_long float64 = 0
		var max_lat float64 = 0
		var max_long float64 = -200
		fmt.Println("length of data", len(censusData))
		for i := 0; i < len(censusData); i++ {
					// fmt.Println("hello1")
				if censusData[i].latitude < min_lat{
					min_lat = censusData[i].latitude

				}
				if censusData[i].latitude > max_lat {
					max_lat = censusData[i].latitude
				}
				if censusData[i].longitude < min_long{
					min_long = censusData[i].longitude
				}
				if censusData[i].longitude > max_long {
					max_long = censusData[i].longitude
				}
		}

		go func(){
			select{
				case checkminlat := <-min_lat1:
					if min_lat < checkminlat {
						go func(){min_lat1 <- min_lat}()
					} else{
						go func(){min_lat1 <- checkminlat}()
					}
				default:
					go func(){min_lat1 <- min_lat}()
			}
		}()

		go func(){
			select{
				case checkmaxlat := <-max_lat1:
					if max_lat > checkmaxlat {
						go func(){max_lat1 <- max_lat}()
					} else{
						go func(){max_lat1 <- checkmaxlat}()
					}
				default:
					go func(){max_lat1 <- max_lat}()
			}
		}()


		go func(){
			select{
				case checkminlong := <-min_long1:
					if min_long < checkminlong {
						go func(){min_long1 <- min_long}()
					} else{
						go func(){min_long1 <- checkminlong}()
					}
				default:
					go func(){min_long1 <- min_long}()
			}
		}()


		go func(){
			select{
				case checkmaxlong := <-max_long1:
					if max_long > checkmaxlong {
						go func(){max_long1 <- max_long}()
					} else{
						go func(){max_long1 <- checkmaxlong}()
					}
				default:
					go func(){max_long1 <- max_long}()
			}
		}()
			fmt.Println(max_long , min_long , max_lat, min_lat, "all the length")

	}

	// donee<-0

}

func v2(censusData []CensusGroup, max_lat chan float64){
	if len(censusData) >1 {
		


	} else{
		maxlatitude := censusData[0].latitude
		select{
			case checkmaxlat := <-max_lat:
				if maxlatitude > checkmaxlat {
					go func(){max_lat <- maxlatitude}()
				} else{
					go func(){max_lat <- checkmaxlat}()
				}
			default:
				go func(){max_lat <- maxlatitude}()
		}
	}
}






func v5Setup(censusData []CensusGroup, min_lat float64 , min_long float64, latIncrement float64, longIncrement float64,xdim int, ydim int , mainArray *[][]int , mainLockArray *[][]sync.Mutex, done chan int){

	for i := 0; i < len(censusData); i++ {
		thisdata := censusData[i]
		lat1 := int(math.Floor((thisdata.latitude - min_lat)/ latIncrement))
		long1 := int(math.Floor((thisdata.longitude - min_long )/ longIncrement ))
		if lat1 == -1 {
			lat1=0
		}
		if xdim-1-long1 == -1 {
			long1=0
		}
		if lat1 == ydim {
			lat1= lat1-1
		}

		// fmt.Println("check ",xdim-1-long1 , lat1)

		// fmt.Println("check ....",(*mainLockArray)[xdim-1-long1][lat1] , xdim-1-long1 , lat1)
		(*mainLockArray)[xdim-1-long1][lat1].Lock()
		(*mainArray)[xdim-1-long1][lat1]=(*mainArray)[xdim-1-long1][lat1] + thisdata.population
		(*mainLockArray)[xdim-1-long1][lat1].Unlock()
	}

	done<-1
}


func ParseCensusData(fname string) ([]CensusGroup, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	censusData := make([]CensusGroup, 0, len(records))

	for _, rec := range records {
		if len(rec) == 7 {
			population, err1 := strconv.Atoi(rec[4])
			latitude, err2 := strconv.ParseFloat(rec[5], 64)
			longitude, err3 := strconv.ParseFloat(rec[6], 64)
			if err1 == nil && err2 == nil && err3 == nil {
				latpi := latitude * math.Pi / 180
				latitude = math.Log(math.Tan(latpi) + 1 / math.Cos(latpi))
				censusData = append(censusData, CensusGroup{population, latitude, longitude})
			}
		}
	}

	return censusData, nil
}


func main () {
	if len(os.Args) < 4 {
		fmt.Printf("Usage:\nArg 1: file name for input data\nArg 2: number of x-dim buckets\nArg 3: number of y-dim buckets\nArg 4: -v1, -v2, -v3, -v4, -v5, or -v6\n")
		return
	}
	fname, ver := os.Args[1], os.Args[4]
	xdim, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	ydim, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	censusData, err := ParseCensusData(fname)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalpop:=float64(0)
	for _,data := range censusData {
		totalpop += float64(data.population)
	}

	latIncrement  := float64(0)
	longIncrement := float64(0)
	min_lat  := float64(1)
	min_long := float64(0)
	max_lat  := float64(0)
	max_long := float64(-200)
 // 	minlat  := float64(1)
	// minlong := float64(0)
	// maxlat  := float64(0)
	// maxlong := float64(-200)
 	min_lat1  := make(chan float64)
 	min_long1 := make(chan float64)
 	max_lat1  := make(chan float64)
 	max_long1 := make(chan float64)
	mainLockArray := make([][] sync.Mutex, xdim)

	done := make(chan int,1)
    mainArray := make([][]int, xdim)
    // grid:= make([][]int, xdim)

    // Some parts may need no setup code
	switch ver {
	case "-v1":

// go run PopulationQuery.go CenPop2010.txt 23 12 -v1


				// fmt.Println(<-max_lat1)

		for i := 0; i < len(censusData); i++ {
			if censusData[i].latitude < min_lat{
				min_lat = censusData[i].latitude
			}
			if censusData[i].latitude > max_lat {
				max_lat = censusData[i].latitude
			}
			if censusData[i].longitude < min_long{
				min_long = censusData[i].longitude
			}
			if censusData[i].longitude > max_long {
				max_long = censusData[i].longitude
			}
		}

		latIncrement  = (max_lat  - min_lat)  / float64(ydim)
		longIncrement = (max_long - min_long) / float64(xdim)

		fmt.Println("latdiff:", latIncrement,"longdiff:" , longIncrement)

	case "-v2":

		go task(censusData[:50001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[50001:100001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[100001:150001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[150001:200001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[200001:len(censusData)],min_lat1, min_long1, max_lat1, max_long1,done)
		// go taskv2(censusData,min_lat1,min_long1,max_lat1, max_long1)


		// go v2(censusData , max_lat1)

		// fmt.Println(<-max_lat1, <-max_long1 , <-min_lat1 , <-min_long1)
		<- done
		<- done
		<- done
		<- done
		<- done

		min_lat = <-min_lat1
		min_long = <-min_long1
		latIncrement  = (<-max_lat1  - min_lat)  / float64(ydim)
		longIncrement = (<-max_long1 - min_long) / float64(xdim)

		fmt.Println("latdiff:", latIncrement,"longdiff:" , longIncrement)

	case "-v3":


  


		for i := 0; i < xdim; i++ {
		    mainArray[i] = make([]int, ydim)
		}


		for i := 0; i < len(censusData); i++ {
			if censusData[i].latitude < min_lat{
				min_lat = censusData[i].latitude
			}
			if censusData[i].latitude > max_lat {
				max_lat = censusData[i].latitude
			}
			if censusData[i].longitude < min_long{
				min_long = censusData[i].longitude
			}
			if censusData[i].longitude > max_long {
				max_long = censusData[i].longitude
			}
		}


		latIncrement  = (max_lat  - min_lat)  / float64(ydim)
		longIncrement = (max_long - min_long) / float64(xdim)


		for i := 0; i < len(censusData); i++ {
			thisdata := censusData[i]
			lat1 := int(math.Floor((thisdata.latitude - min_lat)/ latIncrement))
			long1 := int(math.Floor((thisdata.longitude - min_long )/ longIncrement ))
			if lat1 == -1 {
				lat1=0
			}
			if xdim-1-long1 == -1 {
				long1=0
			}
			if lat1 == ydim {
				lat1= lat1-1
			}

			// fmt.Println(xdim-1-long1 , lat1)
			mainArray[xdim-1-long1][lat1]=mainArray[xdim-1-long1][lat1] + thisdata.population
		}


		// fmt.Println(mainArray)

		for i := 1; i < ydim; i++ {
			mainArray[xdim-1][i]=mainArray[xdim-1][i]+mainArray[xdim-1][i-1]
		}
		for i := xdim-2; i >=0; i-- {
			mainArray[i][0]=mainArray[i][0]+mainArray[i+1][0]
		}

		for i := xdim-2; i >=0; i-- {
			for j := 1; j < ydim; j++ {
				 mainArray[i][j] = mainArray[i][j]+mainArray[i+1][j]+mainArray[i][j-1]-mainArray[i+1][j-1] 
			}
		}

		fmt.Println("latdiff:", latIncrement,"longdiff:" , longIncrement)
		fmt.Println("SETUP READY ....")





	case "-v4":
        
		for i := 0; i < xdim; i++ {
		    mainArray[i] = make([]int, ydim)
		}

		go task(censusData[:50001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[50001:100001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[100001:150001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[150001:200001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[200001:len(censusData)],min_lat1, min_long1, max_lat1, max_long1,done)
		
		<- done
		<- done
		<- done
		<- done
		<- done

		min_lat = <-min_lat1
		min_long = <-min_long1
		latIncrement  = (<-max_lat1  - min_lat)  / float64(ydim)
		longIncrement = (<-max_long1 - min_long) / float64(xdim)

		fmt.Println("latdiff:", latIncrement,"longdiff:" , longIncrement)

		array1 := mainArray
		array2 := mainArray
	 	array3 := mainArray
	 	array4 := mainArray
	 	array5 := mainArray

	 	fmt.Println(array1)

	 	go func () {
	 		array1 = v4Setup(censusData[:50001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, done)}()

	 	go func () {
	 		array2 = v4Setup(censusData[50001:100001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, done)}()

	 	go func () {
	 		array3 = v4Setup(censusData[100001:150001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, done)}()

	 	go func () {
	 		array4 = v4Setup(censusData[150001:200001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, done)}()

	 	go func () {
	 		array5 = v4Setup(censusData[200001:len(censusData)] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, done)}()

	 	<-done
	 	<-done
	 	<-done
	 	<-done
	 	<-done


		for i := 0; i < xdim; i++ {
			for j := 0; j < ydim; j++ {
				mainArray[i][j]=array1[i][j]+array2[i][j]+array3[i][j]+array4[i][j]+array5[i][j]
				
			}	
		} 	


		for i := 1; i < ydim; i++ {
			mainArray[xdim-1][i]=mainArray[xdim-1][i]+mainArray[xdim-1][i-1]
		}
		for i := xdim-2; i >=0; i-- {
			mainArray[i][0]=mainArray[i][0]+mainArray[i+1][0]
		}

		for i := xdim-2; i >=0; i-- {
			for j := 1; j < ydim; j++ {
				 mainArray[i][j] = mainArray[i][j]+mainArray[i+1][j]+mainArray[i][j-1]-mainArray[i+1][j-1] 
			}
		}


		fmt.Println(mainArray)


	case "-v5":
		// var thislock sync.mutex


		for i := 0; i < xdim; i++ {
		    mainLockArray[i] = make([]sync.Mutex, ydim)
		    mainArray[i] = make([]int, ydim)
		}

		fmt.Println(mainLockArray)


	    go task(censusData[:50001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[50001:100001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[100001:150001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[150001:200001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[200001:len(censusData)],min_lat1, min_long1, max_lat1, max_long1,done)
		
		<- done
		<- done
		<- done
		<- done
		<- done

		min_lat = <-min_lat1
		min_long = <-min_long1
		latIncrement  = (<-max_lat1  - min_lat)  / float64(ydim)
		longIncrement = (<-max_long1 - min_long) / float64(xdim)

		fmt.Println("latdiff:", latIncrement,"longdiff:" , longIncrement)

	 	go v5Setup(censusData[:50001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, &mainArray , &mainLockArray, done)

	 	go v5Setup(censusData[50001:100001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, &mainArray , &mainLockArray, done)


	 	go v5Setup(censusData[100001:150001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim,&mainArray , &mainLockArray, done)

	 	go v5Setup(censusData[150001:200001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim,&mainArray , &mainLockArray, done)

	 	go  v5Setup(censusData[200001:len(censusData)] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, &mainArray , &mainLockArray, done)


	 	<-done
	 	<-done	 	
	 	<-done
	 	<-done	 
	 	<-done
	 	
		for i := 1; i < ydim; i++ {
			mainArray[xdim-1][i]=mainArray[xdim-1][i]+mainArray[xdim-1][i-1]
			// fmt.Print(mainArray[xdim-1][i],".")

		}
		// fmt.Println()
		for i := xdim-2; i >=0; i-- {
			mainArray[i][0]=mainArray[i][0]+mainArray[i+1][0]
			// fmt.Print(mainArray[i][0],".")

		}
		// fmt.Println()
		for i := xdim-2; i >=0; i-- {
			for j := 1; j < ydim; j++ {
				 mainArray[i][j] = mainArray[i][j]+mainArray[i+1][j]+mainArray[i][j-1]-mainArray[i+1][j-1] 
				 // fmt.Print(mainArray[i][j],".")
			}
			// fmt.Println()
		}

		// fmt.Println()

	 	for i := 0; i < len(mainArray); i++ {
	 			 	fmt.Println(mainArray[i])

	 	}

        // YOUR SETUP CODE FOR PART 5
	case "-v6":
       	// mainLockArray := make([][] sync.Mutex, xdim)

		for i := 0; i < xdim; i++ {
		    mainLockArray[i] = make([]sync.Mutex, ydim)
		    mainArray[i] = make([]int, ydim)
		}

		fmt.Println(mainLockArray)


	    go task(censusData[:50001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[50001:100001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[100001:150001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[150001:200001],min_lat1, min_long1, max_lat1, max_long1,done)
		go task(censusData[200001:len(censusData)],min_lat1, min_long1, max_lat1, max_long1,done)
		
		<- done
		<- done
		<- done
		<- done
		<- done

		min_lat = <-min_lat1
		min_long = <-min_long1
		latIncrement  = (<-max_lat1  - min_lat)  / float64(ydim)
		longIncrement = (<-max_long1 - min_long) / float64(xdim)

		fmt.Println("latdiff:", latIncrement,"longdiff:" , longIncrement)

	 	go v5Setup(censusData[:50001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, &mainArray , &mainLockArray, done)

	 	go v5Setup(censusData[50001:100001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, &mainArray , &mainLockArray, done)


	 	go v5Setup(censusData[100001:150001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim,&mainArray , &mainLockArray, done)

	 	go v5Setup(censusData[150001:200001] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim,&mainArray , &mainLockArray, done)

	 	go v5Setup(censusData[200001:len(censusData)] , min_lat, min_long, latIncrement, longIncrement,xdim, ydim, &mainArray , &mainLockArray, done)


	 	<-done
	 	<-done	 	
	 	<-done
	 	<-done	 
	 	<-done

	 	var mu sync.Mutex

	 	if xdim<=ydim {
	 		for i := xdim-1; i >=0; i-- {
	 			go func(){
	 				if i==xdim-1 {
	 					for j := 1; j < ydim; j++ {
	 						// fmt.Println("row",i , j)
	 						mainArray[i][j]=mainArray[i][j-1]+mainArray[i][j]
	 										 // fmt.Print(mainArray[i][j],".")


	 					}
	 						 			// fmt.Println("--row----")

	 				}else{
	 					for j := xdim-i; j < ydim; j++ {

	 						if j==xdim-i {
	 						fmt.Println("lock here at upper", j , i)

	 							mu.Lock()
								mainArray[i][j] = mainArray[i][j]+mainArray[i+1][j]+mainArray[i][j-1]-mainArray[i+1][j-1] 
								mu.Unlock()
	 							fmt.Println("un lock here at upper", j , i)

	 						}else{
	 						// fmt.Println("row",i , j)
							 mainArray[i][j] = mainArray[i][j]+mainArray[i+1][j]+mainArray[i][j-1]-mainArray[i+1][j-1] 
	 									 // fmt.Print(mainArray[i][j],".")
							}

	 					}
	 						 			// fmt.Println("---row---")

	 				}
	 				done<-1
	 			}()
	 			// <-done
	 			if i==xdim-1 {
	 				for j := i-1; j >=0; j-- {
						mainArray[j][0]=mainArray[j][0]+mainArray[j+1][0]
						// fmt.Print(mainArray[j][0],".")
					}
						 			// fmt.Println("---col---")

	 			}else {
	 				for j := i; j >=0; j-- {
	 					// fmt.Println(mainArray[j+1][xdim-1-i] , mainArray[j][xdim-1-i-1] , mainArray[j+1][xdim-1-i-1])

	 					if j==i {
	 						fmt.Println("lock here at", j , xdim-1-i)
	 						mu.Lock()
							mainArray[j][xdim-1-i]=mainArray[j][xdim-1-i]+mainArray[j+1][xdim-1-i]+mainArray[j][xdim-1-i-1]-mainArray[j+1][xdim-1-i-1]
							mu.Unlock()
							fmt.Println("un lock here at", j , xdim-1-i)

	 					} else{

						mainArray[j][xdim-1-i]=mainArray[j][xdim-1-i]+mainArray[j+1][xdim-1-i]+mainArray[j][xdim-1-i-1]-mainArray[j+1][xdim-1-i-1]
						// fmt.Print(mainArray[j][xdim-1-i],".")
					}
					}	
						 			// fmt.Println("---col---")

	 			}
	 			// fmt.Println("--------")
	 			<-done
	 		}
	 	}else if ydim<xdim{

	 		for i := 0; i < ydim; i++ {
	 			p :=i
	 			go func() {
	 				fmt.Println("in go",p)

	 				if i==0 {

	 					for j := xdim-2; j >= 0; j-- {
	 							fmt.Println("........", j,i)

	 						mainArray[j][i]=mainArray[j][i]+mainArray[j+1][i]
	 					}
	 				} else{
	 					for j := xdim-i-2 ; j >= 0 ; j-- {
	 						if j==xdim-i-2{
	 							mu.Lock()
	 							fmt.Println("........", j,i)
	 							mainArray[j][i] = mainArray[j][i] + mainArray[j+1][i] + mainArray[j][i-1] - mainArray[j+1][i-1]    
	 							mu.Unlock()
	 						} else {
	 							mainArray[j][i] = mainArray[j][i] + mainArray[j+1][i] + mainArray[j][i-1] - mainArray[j+1][i-1]    
	 						}

	 					}

	 				}
	 				done<-1
	 			}()
	 			if i==0 {
	 				for j := 1; j < ydim; j++ {
	 					mainArray[xdim-1][j]=mainArray[xdim-1][j]+mainArray[xdim-1][j-1]
	 				}
	 			} else {
	 				for j := i; j < ydim; j++ {
	 					if j==i {
	 						mu.Lock()
	 						mainArray[xdim-i-1][j] = mainArray[xdim-i-1][j] + mainArray[xdim-i-1+1][j]+mainArray[xdim-i-1][j-1]-mainArray[xdim-i-1+1][j-1] 
	 						mu.Unlock()
	 						
	 					} else{
	 						mainArray[xdim-i-1][j] = mainArray[xdim-i-1][j] + mainArray[xdim-i-1+1][j]+mainArray[xdim-i-1][j-1]-mainArray[xdim-i-1+1][j-1] 
	 					}
	 				}
	 			}

	 			<-done

	 		}
	 	}

	 	for i := 0; i < len(mainArray); i++ {
	 			 	fmt.Println("no",mainArray[i])
	 	}


	default:
		fmt.Println("Invalid version argument")
		return
	}

	for {
		var west, south, east, north int
		n, err := fmt.Scanln(&west, &south, &east, &north)
		if n != 4 || err != nil || west<1 || west>xdim || south<1 || south>ydim || east<west || east>xdim || north<south || north>ydim {
			break
		}

		var population int

		var percentage float64
		switch ver {
		case "-v1":


			var totalPop float64

			for i := 0; i < len(censusData); i++ {
				thisdata := censusData[i]
				lat1 := int(math.Ceil((thisdata.latitude - min_lat)/ latIncrement))
				long1 := int(math.Ceil((thisdata.longitude - min_long )/ longIncrement ))


				if lat1 >= south && lat1 <= north && long1 >= west && long1 <= east {
					population = population+thisdata.population
				}

				totalPop = totalPop + float64(thisdata.population)
			}

			percentage = float64(float64(population) / totalPop)* float64(100)

		case "-v2":

			pop := make(chan int)
			
			go v2Query(censusData[:50001] , pop , done, min_lat, min_long,west, south, east, north,latIncrement,longIncrement)
			go v2Query(censusData[50001:100001] , pop , done, min_lat, min_long,west, south, east, north,latIncrement,longIncrement)
			go v2Query(censusData[100001:150001] , pop , done, min_lat, min_long,west, south, east, north,latIncrement,longIncrement)
			go v2Query(censusData[150001:200001] , pop , done, min_lat, min_long,west, south, east, north,latIncrement,longIncrement)
			go v2Query(censusData[200001:len(censusData)] , pop , done, min_lat, min_long,west, south, east, north,latIncrement,longIncrement)
			<-done
			<-done
			<-done
			<-done
			<-done
			// fmt.Println(<-pop,totalpop)
			population = <- pop
			percentage = float64(float64(population) / totalpop)* float64(100)



		case "-v3":
            west = xdim-west
            south  = south-1
            east = xdim-east
            north  = north-1

         
            northeast := mainArray[east][north]
            southeast := mainArray[east][south]
            northwest := mainArray[west][north]
            southwest := mainArray[west][south]

            if west==xdim-1 && south==0 {
            	population=mainArray[east][north]


            } else if west!=xdim-1 && south==0{
            	west=west+1
          	  northwest = mainArray[west][north]
           	 population = northeast-northwest

            }else if south!=0 && west==xdim-1{
            	south=south-1

            	southeast = mainArray[east][south]

            	population = northeast - southeast
            } else {
                south=south-1
                west=west+1
            southeast = mainArray[east][south]
            northwest = mainArray[west][north]
            southwest = mainArray[west][south]

            population= northeast - southeast - northwest + southwest
        	}

			percentage = float64(float64(population) / totalpop)* float64(100)

		case "-v4":
                        west = xdim-west
            south  = south-1
            east = xdim-east
            north  = north-1

         
            northeast := mainArray[east][north]
            southeast := mainArray[east][south]
            northwest := mainArray[west][north]
            southwest := mainArray[west][south]

            if west==xdim-1 && south==0 {
            	population=mainArray[east][north]


            } else if west!=xdim-1 && south==0{
            	west=west+1
          	  northwest = mainArray[west][north]
           	 population = northeast-northwest

            }else if south!=0 && west==xdim-1{
            	south=south-1

            	southeast = mainArray[east][south]

            	population = northeast - southeast
            } else {
                south=south-1
                west=west+1
            southeast = mainArray[east][south]
            northwest = mainArray[west][north]
            southwest = mainArray[west][south]

            population= northeast - southeast - northwest + southwest
        	}

			percentage = float64(float64(population) / totalpop)* float64(100)
		case "-v5":
                        west = xdim-west
            south  = south-1
            east = xdim-east
            north  = north-1

         
            northeast := mainArray[east][north]
            southeast := mainArray[east][south]
            northwest := mainArray[west][north]
            southwest := mainArray[west][south]

            if west==xdim-1 && south==0 {
            	population=mainArray[east][north]


            } else if west!=xdim-1 && south==0{
            	west=west+1
          	  northwest = mainArray[west][north]
           	 population = northeast-northwest

            }else if south!=0 && west==xdim-1{
            	south=south-1

            	southeast = mainArray[east][south]

            	population = northeast - southeast
            } else {
                south=south-1
                west=west+1
            southeast = mainArray[east][south]
            northwest = mainArray[west][north]
            southwest = mainArray[west][south]

            population= northeast - southeast - northwest + southwest
        	}

			percentage = float64(float64(population) / totalpop)* float64(100)

            // YOUR QUERY CODE FOR PART 5
		case "-v6":
                  west = xdim-west
            south  = south-1
            east = xdim-east
            north  = north-1

         
            northeast := mainArray[east][north]
            southeast := mainArray[east][south]
            northwest := mainArray[west][north]
            southwest := mainArray[west][south]

            if west==xdim-1 && south==0 {
            	population=mainArray[east][north]


            } else if west!=xdim-1 && south==0{
            	west=west+1
          	  northwest = mainArray[west][north]
           	 population = northeast-northwest

            }else if south!=0 && west==xdim-1{
            	south=south-1

            	southeast = mainArray[east][south]

            	population = northeast - southeast
            } else {
                south=south-1
                west=west+1
            southeast = mainArray[east][south]
            northwest = mainArray[west][north]
            southwest = mainArray[west][south]

            population= northeast - southeast - northwest + southwest
        	}

			percentage = float64(float64(population) / totalpop)* float64(100)
		}

		fmt.Printf("%v %.2f%%\n", population, percentage)
	}
}
