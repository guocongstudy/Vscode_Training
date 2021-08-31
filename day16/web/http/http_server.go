package main

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter,r *http.Request){
	fmt.Printf("host %s\n",r.Host)
	fmt.Printf("method %s\n",r.Method)
	fmt.Printf("header")
	for key,values :=range r.Header{
		fmt.Printf("%s = %v\n",key,values)
	}
	fmt.Println("cookie")
	for _,cookie:=range r.Cookies(){
		fmt.Printf("name:%s,values:%s,path:%s,domain:%s\n",cookie.Name,cookie.Value,cookie.Path,cookie.Domain)
	}
	fmt.Fprintf(w,"Hi boy")
}


func main(){
	http.HandleFunc("/",RootHandler)
	http.ListenAndServe(":5656",nil)
}


