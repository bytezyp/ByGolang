package main

import "fmt"

type Price struct {
	Value int `json:"value"`
}

func processSlice(price []Price) {
	fmt.Printf("222————1 %p %p %v %v %v\n", &price, &price[0], price, len(price), cap(price))
	for i := range price {
		if i == 1 {
			price = append(price[:i], price[i+1:]...)
		}
	}

	fmt.Printf("222————2 %p %p %v %v %v \n", &price, &price[0], price, len(price), cap(price))
	//price = append(price, Price{Value: 10})
}

func main() {
	var priceList []Price

	// len 3 cap 4
	priceList = append(priceList, Price{Value: 10}, Price{Value: 12}, Price{Value: 15})
	fmt.Printf("111 %p %p %v %v %v \n", &priceList, &priceList[0], priceList, len(priceList), cap(priceList))
	processSlice(priceList)
	fmt.Printf("333 %p %p %v %v %v \n", &priceList, &priceList[0], priceList, len(priceList), cap(priceList))
	priceList = append(priceList, Price{Value: 10})
	fmt.Printf("444 %p %p %v %v %v \n", &priceList, &priceList[0], priceList, len(priceList), cap(priceList))

}
