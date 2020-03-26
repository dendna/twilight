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
		// {40.9, -74.3, twilight.DuskTypeSimple, 1990, 6, 25},

		// {55.75, 37.62, twilight.DuskTypeAstronomical, 2019, 03, 25},
		// {55.75, 37.62, twilight.DuskTypeBorders, 2019, 03, 25},
		// {55.75, 37.62, twilight.DuskTypeAstronomical, 2020, 03, 25},

		// {55.7857, 49.1199, twilight.DuskTypeAstronomical, 2020, 03, 25},
		// {55.7857, 49.1199, twilight.DuskTypeNautical, 2020, 03, 25},
		// {55.7857, 49.1199, twilight.DuskTypeCivil, 2020, 03, 25},
		// {55.7857, 49.1199, twilight.DuskTypeMid, 2020, 03, 25},
		// {55.7857, 49.1199, twilight.DuskTypeSimple, 2020, 03, 25},
		// {55.7857, 49.1199, twilight.DuskTypeBorders, 2020, 03, 25},

		// kirov
		// {58.6018, 49.6706, twilight.DuskTypeAstronomical, 2019, 01, 02},
		// {58.6018, 49.6706, twilight.DuskTypeNautical, 2019, 01, 02},
		// {58.6018, 49.6706, twilight.DuskTypeCivil, 2019, 01, 02},
		// {58.6018, 49.6706, twilight.DuskTypeMid, 2019, 01, 02},
		// {58.6018, 49.6706, twilight.DuskTypeSimple, 2019, 01, 02},
		// {58.6018, 49.6706, twilight.DuskTypeBorders, 2019, 01, 02},

		// //{58.603591, 49.668014, twilight.DuskTypeAstronomical, 2019, 01, 02},
		// {58.603591, 49.668014, twilight.DuskTypeNautical, 2019, 01, 02},
		// {58.603591, 49.668014, twilight.DuskTypeCivil, 2019, 01, 02},
		// //{58.603591, 49.668014, twilight.DuskTypeMid, 2019, 01, 02},
		// {58.603591, 49.668014, twilight.DuskTypeSimple, 2019, 01, 02},
		// {58.603591, 49.668014, twilight.DuskTypeBorders, 2019, 01, 02},

		// {58.6018, 49.6706, twilight.DuskTypeAstronomical, 2019, 06, 19},
		// {58.6018, 49.6706, twilight.DuskTypeNautical, 2019, 06, 19},
		// {58.6018, 49.6706, twilight.DuskTypeCivil, 2019, 06, 19},

		{58.6018, 49.6706, twilight.DuskTypeBorders, 2019, 01, 02},
		{58.6018, 49.6706, twilight.DuskTypeSimple, 2019, 01, 02},
		{58.6018, 49.6706, twilight.DuskTypeCivil, 2019, 01, 02},
		{58.6018, 49.6706, twilight.DuskTypeNautical, 2019, 01, 02},

		// {58.6018, 49.6706, twilight.DuskTypeBorders, 2019, 03, 25},
		// {58.6018, 49.6706, twilight.DuskTypeSimple, 2019, 03, 25},
		// {58.6018, 49.6706, twilight.DuskTypeCivil, 2019, 03, 25},
		// {58.6018, 49.6706, twilight.DuskTypeNautical, 2019, 03, 25},

		// {52.5, -1.9167, twilight.DuskTypeCivil, 1998, 10, 25},

	}

	for i := range input {
		fmt.Println()
		fmt.Println(twilight.Calc(input[i].latitude, input[i].longitude, input[i].dusk, input[i].year, input[i].month, input[i].day))
		fmt.Println(twilight.CalcRise(input[i].latitude, input[i].longitude, input[i].dusk, input[i].year, input[i].month, input[i].day))
		fmt.Println(twilight.CalcSet(input[i].latitude, input[i].longitude, input[i].dusk, input[i].year, input[i].month, input[i].day))
	}

}
