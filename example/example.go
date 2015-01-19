package main

import (
	"fmt"
	"os"

	"../uaparser" // You could change this to a github repo as well
)

func main() {
	testStr := "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-80) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16 Silk-Accelerated=true"
	parser, err := uaparser.New()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	client := parser.Parse(testStr)
	fmt.Println(testStr)
	fmt.Println("UserAgent: " + client.UserAgent.ToString())
	fmt.Println("OS: " + client.Os.ToString())
	fmt.Println("Device: " + client.Device.ToString())
}
