package main

import(
	"fmt"
	"net/http"
	"hutuguaner/message"
	"log"
)

func main(){
	fmt.Println("main start ...")
	
	http.HandleFunc("/putmsg/",message.Putmsg)
	http.HandleFunc("/getmsg/",message.Getmsg)

	if err := http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal("ListenAndServe: ",err)
	}
}