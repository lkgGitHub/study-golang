package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	mySolution()
	//number, letter := make(chan bool), make(chan int)
	//
	//go func() {
	//	defer println("number over")
	//	i := 1
	//	println("number")
	//	for true {
	//		select {
	//		case <-number:
	//			print(i)
	//			i++
	//			letter <- i
	//		}
	//	}
	//}()
	//
	//wg := sync.WaitGroup{}
	//go func(w *sync.WaitGroup) {
	//	wg.Add(1)
	//	defer println("letter over")
	//	const str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//	strArr := strings.Split(str, "")
	//
	//	for true {
	//		select {
	//		case i := <-letter:
	//			if i-2 < len(strArr) {
	//				print(strArr[i-2])
	//				number <- true
	//			} else {
	//				close(number)
	//				wg.Done()
	//			}
	//		}
	//	}
	//}(&wg)
	//
	//number <- true
	//wg.Wait()
}

func mySolution() {
	number, letter := make(chan bool), make(chan int)

	go func() {
		defer println("over number")
		i := 1
	loop:
		for true {
			select {
			case b, isOpen := <-number:
				if isOpen {
					if b {
						print(i)
						i++
						print(i)
						i++
						letter <- i
					} else {
						println()
						close(letter)
					}
				} else {
					break loop
				}
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(w *sync.WaitGroup) {
		defer println("over letter")
		const str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		strArr := strings.Split(str, "")
	loop:
		for true {
			select {
			case i, isOpen := <-letter:
				if isOpen {
					if i-3 < len(strArr) {
						for j := i - 3; j < i-1; j++ {
							print(strArr[j])
						}
						number <- true
					} else {
						w.Done()
						number <- false
						close(number)
					}
				} else {
					break loop
				}
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}

func solution() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}

				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}

		}
	}(&wait)
	number <- true
	wait.Wait()
}
