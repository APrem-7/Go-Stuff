package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_PAVBHAJI_PRICE float32 = 10
var MAX_PURI_PRICE float32 = 10

func main() {
	var pavBhaji = make(chan string)
	var puri = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "walgreens.com"}
	for i := range websites {
		go getChickenPrices(websites[i], pavBhaji)
		go getPuriPrice(websites[i],puri)
	}

	sendMessage(pavBhaji)
	sendMessage(puri)

}

func getChickenPrices(websites string, pavBhaji chan string) {
	for {
		time.Sleep(time.Second * 1)
		pavBhajiPrice := rand.Float32() * 20
		if pavBhajiPrice <= MAX_PAVBHAJI_PRICE {
			pavBhaji <- websites
			break
		}

	}

}

func getPuriPrice(websites string,puri chan string){
	for{
		time.Sleep(time.Second * 1)
		puriPrice := rand.Float32()*20
		if puriPrice <= MAX_PURI_PRICE{
			puri <- websites
		}
	}
}

func sendMessage(pavBhaji chan string) {
	fmt.Printf("Found a deal on chicken at %s", <-pavBhaji)
}
