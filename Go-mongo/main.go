package main

import (
	"fmt"
	"net/http"
	"github.com/bk167465/Go_Tutorial/Go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main(){
    r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

//MGO is not maintained
func getSession() *mgo.Session{
	s, err := mgo.Dial("127.0.0.1")
	if err != nil{
		panic(err)
	}
	fmt.Println("Connected")
	return s
}