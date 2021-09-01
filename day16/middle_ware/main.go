package main

import (
	"log"
	"net/http"
	"time"
)

var limitChan =make(chan struct{},100)

func getBoy(w http.ResponseWriter, r *http.Request) {
	time.Sleep(150 * time.Millisecond)
	w.Write([]byte("hi boy"))
}

func timeMiddleWare(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		inner.ServeHTTP(rw, r)
		_= time.Since(begin)
		//userTime := time.Since(begin)
		//log.Printf("use time %d ms\n", userTime.Milliseconds())
	})

}
//对并发数进行限制，规定并发次数。
func limitMiddleWare(inner http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		limitChan<- struct{}{}
		log.Printf("coccurrence %d\n",len(limitChan))
		inner.ServeHTTP(rw,r)
		<-limitChan
	})
}

func main() {
	http.Handle("/", limitMiddleWare(timeMiddleWare(http.HandlerFunc(getBoy))))
	http.ListenAndServe(":5656", nil)
}
