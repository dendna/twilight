# twilight

Go package used to calculate sunrise and sunset times based on twilight type, latitude, longitude and date.

------



### General

Based on http://www.edwilliams.org/sunrise_sunset_algorithm.htm

##### Twilight valid types:

- Borders			  = 90&deg;20' (90.3)
- Simple				= 90&deg;50' (90.83)
- Mid					 = 93&deg;
- Civil					= 96&deg;
- Nautical			 = 102&deg;
- Astronomical	= 108&deg;



### Installation

```bash
go get github.com/dendna/twilight
cd ${GOPATH}/src/github.com/dendna/twilight/cmd/twilight
go build
```



### Usage

```go
package main

import (
	"fmt"

	"github.com/dendna/twilight"
)

func main() {
	var input = []struct {
		latitude  float64
		longitude float64
		dusk      twilight.DuskType
		year      int
		month     int
		day       int
	}{		
		{58.6018, 49.6706, twilight.DuskTypeBorders, 2019, 01, 02},
		{58.6018, 49.6706, twilight.DuskTypeSimple, 2019, 01, 02},
    	{58.6018, 49.6706, twilight.DuskTypeMid, 2019, 01, 02},
		{58.6018, 49.6706, twilight.DuskTypeCivil, 2019, 01, 02},
		{58.6018, 49.6706, twilight.DuskTypeNautical, 2019, 01, 02},
		{58.6018, 49.6706, twilight.DuskTypeAstronomical, 2019, 01, 02},
	}

	for i := range input {
        // calculate the sunrise and sunset times
        sunrise, sunset, err := twilight.Calc(
            input[i].latitude, 
            input[i].longitude, 
            input[i].dusk, 
            input[i].year, 
            input[i].month, 
            input[i].day
        )
        
        // print results
        if err == nil {
            fmt.Println("Sunrise:", sunrise.Format("15:04:05"))
	        fmt.Println("Sunset:", sunset.Format("15:04:05"))
        } else {
            fmt.Println(err)
		}	
    
    	// separated calculations
        
    	// sunrise, err := twilight.CalcRise(
        //    input[i].latitude, 
        //    input[i].longitude, 
        //    input[i].dusk, 
        //    input[i].year, 
        //    input[i].month, 
        //    input[i].day
        // )        
        // if err == nil {
        //    fmt.Println("Sunrise:", sunrise.Format("15:04:05"))
        // } else {
        //    fmt.Println(err)
		// }
        
        // sunset, err := twilight.CalcSet(
        //    input[i].latitude, 
        //    input[i].longitude, 
        //    input[i].dusk, 
        //    input[i].year, 
        //    input[i].month, 
        //    input[i].day
        // )        
        // if err == nil {
        //    fmt.Println("Sunset:", sunset.Format("15:04:05"))
        // } else {
        //    fmt.Println(err)
		// }
	}

}

```



### License