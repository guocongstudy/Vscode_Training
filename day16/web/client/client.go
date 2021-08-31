package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func simpleGet(){
	if resp,err :=http.Get("http://127.0.0.1:5656");err !=nil{
		fmt.Println(err)
	}else{
		defer resp.Body.Close()
		fmt.Printf("proto:%s\n",resp.Proto)
		fmt.Printf("status:%v\n",resp.Status)
		fmt.Println("header")
		for key,values :=range resp.Header{
			fmt.Printf("%s,%v",key,values)
		}
		fmt.Println()
		fmt.Printf("body:")
		io.Copy(os.Stdout,resp.Body)
	}
}

func simplePost(){
	reader :=strings.NewReader("Hi Server")
	if resp,err :=http.Post("http://127.0.0.1:5656","text/plain",reader) ;err !=nil{
		fmt.Println(err)
	}else{
		defer resp.Body.Close()
		fmt.Printf("proro:%s\n",resp.Proto)
		fmt.Printf("status:%v\n",resp.Status)
		fmt.Println("header")
		for key,value :=range resp.Header{
			fmt.Printf("%s=%v\n",key,value)
		}
		fmt.Println()
		fmt.Printf("body: ")
		io.Copy(os.Stdout,resp.Body)
		fmt.Println()
	}
}

func main(){
	simpleGet()
	simplePost()

}