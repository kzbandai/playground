package main

import (
	"fmt"
	"log"
	"time"
)

func threeSecond() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("3 seconds passed")
}

func oneSecond() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("1 second passed")
}

func twoSecond() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("2 second passed")
}

func section1() {
	// https://qiita.com/TakaakiFuruse/items/241578174fd2f00aaa8a
	fmt.Println(time.Now())
	go threeSecond()
	oneSecond()
	twoSecond()
	fmt.Println(time.Now())
}

func wait1(c chan string) {
	time.Sleep(1 * time.Second)
	log.Print("waited 1 sec")
	c <- "wait1 finished"

}

func wait2(c chan string) {
	time.Sleep(2 * time.Second)
	log.Print("waited 2 sec")
	c <- "wait2 finished"
}

func wait5(c chan string) {
	time.Sleep(5 * time.Second)
	log.Print("waited 5 sec")
	c <- "wait5 finished"
}

func section2() {
	// https://qiita.com/TakaakiFuruse/items/241578174fd2f00aaa8a
	// channelを作ると他の処理が全部終わるまで待ってくれる
	c := make(chan string)

	log.Print("started")

	go wait1(c)
	go wait2(c)
	go wait5(c)

	w1, w2, w3 := <-c, <-c, <-c
	log.Print("finished")

	fmt.Println(w1)
	fmt.Println(w2)
	fmt.Println(w3) // この三行は一気に表示される
}

func section3() {
	// https://qiita.com/kkohtaka/items/c42bfc75bede7cd8dc50
	res := make(chan int)
	go func(res chan int) {
		fmt.Println("go func print")
		time.Sleep(2 * time.Second)
		res <- 100
	}(res)

	fmt.Println("normal print", <-res)
}

func section4() {
	// https://qiita.com/kkohtaka/items/c42bfc75bede7cd8dc50
	res := make(chan int)

	// buffered channel
	//res := make(chan int, 5)
	//res := make(chan int, 10)

	go func(res chan int) {
		for i := 0; i < 42; i++ {
			log.Println("another goroutine is sending:", i)
			res <- i
		}
		close(res)
	}(res)

	for {
		i, ok := <-res
		if !ok {
			break
		}

		log.Println("the main goroutine receives:", i)
	}
}

func section5() {
	// https://qiita.com/kkohtaka/items/c42bfc75bede7cd8dc50
	// buffered channel では, 未受信の値でバッファが埋められている間はバッファに空きが生まれるまでブロックする
	// 複数の並行処理の並行度を制限することができる
	sem := make(chan struct{}, 3) // concurrency: 3

	for i := 0; i < 42; i++ {
		sem <- struct{}{}

		go func(i int, sem chan struct{}) {
			log.Println("goroutine: ", i)
			<-sem // 値を受信＝解放する
		}(i, sem)
	}

	log.Println("the main goroutine")
}

func section6() {
	// https://qiita.com/kkohtaka/items/c42bfc75bede7cd8dc50
	// BとCは実質同じ
	// Aはchannelを渡すためのchannel、という前提を忘れない

	A := make(chan chan int)

	go func(A chan chan int) {
		for {
			select {
			case B := <-A: // 値を受信して代入、Bはchannel
				log.Println("another goroutine receives a request")
				B <- 100 // 値を送信
			}
		}
	}(A)

	request := func() int {
		C := make(chan int)
		A <- C     // 値を送信
		return <-C // 値を受信
	}

	time.Sleep(3 * time.Second)
	log.Println("the main goroutine receives a response:", request())
}

func main() {
	//section1()
	//section2()
	//section3()
	//section4()
	//section5()
	section6()
}
