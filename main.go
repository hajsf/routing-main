package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://developers.google.com/maps/documentation/distance-matrix/distance-matrix

func main() {

	api_key := "AIzaSyBC6tf0jSZTjZ7BbJWWy8ZOzrKFraReqkk" //"AIzaSyCA4CWdFzFq1ceXh5f0rkoPxXfkCtH5ic0"
	url := "https://maps.googleapis.com/maps/api/distancematrix/json" +
		"?origins=40.6655101,-73.89188969999998" +
		"&destinations=" +
		"40.659569,-73.933783" +
		"%7C40.729029,-73.851524" +
		"%7C40.6860072,-73.6334271" +
		"%7C40.598566,-73.7527626" +
		//	"origins=Washington,%20DC" +
		//	"&destinations=New%20York%20City,%20NY" +
		//"&units=imperial"+
		"&key=" + api_key +
		"&arrival_time=2022-08-31T23:50:21.817Z" //+
		//"&language=ar"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
