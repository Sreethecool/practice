package main

import (
	"fmt"
	"sync"
	"github.com/practice/sorting"
)

var odd,even chan int32

var servChan = make(chan int32)


func Serve(){
	val := <- servChan
	if val%2 == 0 {
		even <- val
	}else{
		odd <- val
	}

}

func main(){

	strArr := []string{"a","ab","bc","abc"}
    result := sort.RemainderSorting(strArr)
	fmt.Println(result)

}

func main1(){
	fmt.Println("Hello World")
	var wg sync.WaitGroup
	arr := []int32{1,2,3,4,5,6,7,8,9,10}

	odd := make(chan int32)
	even := make(chan int32)
 


	for _,v := range arr{
		servChan <- v
	}

	go func(){
		defer wg.Done()
		a := <-odd

		for elem := range queue {
			fmt.Println(elem)
		}

	}

	go func(){
		defer wg.Done()
		a := <-odd

		for elem := range queue {
			fmt.Println(elem)
		}
		
	}



}