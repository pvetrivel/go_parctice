package main

import "bitbucket.org/tekion/rollOut_Backend/test"

func main() {

	test.Start()
}

//import (
//	"fmt"
//	"time"
//)
//
//func producer(iters int) <-chan int {
//	c := make(chan int)
//	go func() {
//		for i := 0; i < iters; i++ {
//			c <- i
//			time.Sleep(1 * time.Second)
//		}
//		close(c)
//	}()
//	return c
//}
//
//func consumer(cin <-chan int) {
//	for i := range cin {
//		fmt.Println(i)
//	}
//}
//
//func fanOutUnbuffered(ch <-chan int, size int) []chan int {
//	cs := make([]chan int, size)
//	for i, _ := range cs {
//		cs[i] = make(chan int)
//	}
//	go func() {
//		for i := range ch {
//			for _, c := range cs {
//				c <- i
//			}
//		}
//		for _, c := range cs {
//			close(c)
//		}
//	}()
//	return cs
//}
//
//func main() {
//	c := producer(10)
//	chans := fanOutUnbuffered(c, 3)
//	go consumer(chans[0])
//	go consumer(chans[1])
//	consumer(chans[2])
//}
