package main

import "fmt"

type RegestRequest struct {
	UserName string `validate:"gt=6"`
	Password string `validate:"min=6,max=12"`
	PassRepeat string `validate:"eqfield=Passwd"`
	Email string `validate:"email"`
}

func main(){
	req :=RegestRequest{
		UserName:"zcy",
		Password: "121312312",
		PassRepeat: "231232",
		Email: "dwd",
	}
	fmt.Println(req)

}