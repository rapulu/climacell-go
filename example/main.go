package main

import (
	"fmt"

	"github.com/andyhaskell/climacell"
)

const(
	apiKey = "2ZZDGH0XHABKxzuQBTW1YoqO65fDfxwR"
)

func main()  {
	c := climacell.New(apiKey)

	fArgs := climacell.ForecastArgs{}
	fArgs.Location = climacell.LatLon{Lat:6.5244, Lon:3.3792}
	fArgs.Fields = []string{"humidity", "temp",}

	w, err := c.Nowcast(fArgs)
		if err != nil{
			fmt.Println(err)
			return
		}

	for _, weather := range w{
		temp, ok := weather.Humidity.GetValue()
		if !ok {
			/* handle a temp value being absent */
		}
		fmt.Printf("%v \n", temp)
	}
}