package main

import (
	"fmt"
	"time"
)

var ch=make(chan int);
var num int;
func fnct1() (int,int){
time.Sleep(time.Second * 4);
return 2,3;
}

func fnct2() {
	<-ch;
	fmt.Print("hello brothher" );
}

func main() {
	go fnct2(); 
	val1,_:=fnct1();
	ch<-val1;
 fmt.Scan(&num);
	fmt.Print("helll broter")
}