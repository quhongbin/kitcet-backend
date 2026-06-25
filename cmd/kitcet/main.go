package main

import "fmt"

type User struct{
	Name string
	Age int
	Height float32
}

func (user User)Greet(){
	fmt.Printf("hello,I am %s, %d years old", user.Name, user.Age)
}

func sendData(ch chan string){
	ch <- "data from goroutine"
}

func main(){
	ch := make(chan string)

	go sendData(ch)

	msg := <-ch
	fmt.Println(msg)
}
