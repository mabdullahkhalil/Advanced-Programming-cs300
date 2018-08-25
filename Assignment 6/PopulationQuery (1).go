/*

Muhammad Abdullah Khalil
19100142

part1: 
 the code is sequential.
 setup: the code for O(n) and finds out the max/min latitude and longitude.
 		then find the difference of each box horizontally and vertically.
 query: in the query part, it finds the box number and compares the limits.

part 2:
	setup:  made 5 go functions, it is kind of hardcoded. made everything channels. 
			every go routine before writing to channel, checks using slects and cases
			that in case anyhting is present in the channel before, if not, it adds the value,
			in case something is present it takes it out and adds new value on the basis of that. 

	query: query codes is divided in 5 go routines too. and after every routine it checks the channels and adds the value 
			to it.

part 3: 
	setup : exactly done as in handout.
	query: exactly done as in handout.

*/



package main

import (
"fmt"
"os"
"strconv"
"math"
"encoding/csv"
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
 	
 	min_lat1  := make(chan float64)
 	min_long1 := make(chan float64)
 	max_lat1  := make(chan float64)
 	max_long1 := make(chan float64)

	done := make(chan int,1)
    mainArray := make([][]int, xdim)

    // Some parts may need no setup code
	switch ver {
	case "-v1":


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
				if long1 == -1 {
					long1=0
				}
				if lat1 == ydim {
					lat1= lat1-1
				}
				mainArray[xdim-1-long1][lat1]=mainArray[xdim-1-long1][lat1] + thisdata.population
			}

			pop:=0
			for i := 0; i < xdim; i++ {
				for j := 0; j < ydim; j++ {
					pop = pop+mainArray[i][j]
					
				}	
			}
			fmt.Println(pop)

			for i := 1; i < ydim; i++ {
				mainArray[22][i]=mainArray[22][i]+mainArray[22][i-1]
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
        // YOUR SETUP CODE FOR PART 4
	case "-v5":
        // YOUR SETUP CODE FOR PART 5
	case "-v6":
        // YOUR SETUP CODE FOR PART 6
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
            // YOUR QUERY CODE FOR PART 4
		case "-v5":
            // YOUR QUERY CODE FOR PART 5
		case "-v6":
            // YOUR QUERY CODE FOR PART 6
		}

		fmt.Printf("%v %.2f%%\n", population, percentage)
	}
}
