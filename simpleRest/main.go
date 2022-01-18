package main

import (
	"net/http"
	"simpleRest/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	//creating new instance of router and assigning it to r var
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	//this creates Go server on the port and we pass the router here
	http.ListenAndServe("localhost:8080", r)
}

//returns a pointer to mongodb session
//*mgo.Session is return type of the  method
func getSession() *mgo.Session {
	//gets session from running db in s and if error then in variabe in err
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		//we can do whatever we want this, e.g: stop everything, send some message to front-end
		panic(err)
	}

	return s
}
