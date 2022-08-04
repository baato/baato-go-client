package main

import (
	"fmt"

	baato "github.com/baato/baato-go-client/lib"
)

func main() {

	baatoMap, _ := baato.Baato("ssd")

	reverseGeocode := baatoMap.ReverseGeocode.GetGeocode()
	// fmt.Println(reverseGeocode)
	fmt.Println(reverseGeocode.Data[0].Name)
}
