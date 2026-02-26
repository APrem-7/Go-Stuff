package main

import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"
)
	
var MAX_PAVBHAJI_PRICEfloat32=10

func main() {
	var pavBhaji= make(chan string)
	var websites=[]string{"walmart.com","costco.com","walgreens.com"}
	for i:= range websites{
		getChickenPrices(i,pavBhaji)
	}

	
}

func getChickenPrices(websties string,pavBhaji chan string){
	for {
		time.Sleep(time.Seconds*1)
		pavBhajiPrice=rand.Float32()*20
		if pavBhajiPrice<=MAX_PAVBHAJI_PRICEfloat32{
			chickenChannel <- websites
			break
		}

	}

}

func sendMessage(chickenChannel chan string){
	fmt.Printf("Found a deal on chicken at %s",<- chickenChannel)
}
