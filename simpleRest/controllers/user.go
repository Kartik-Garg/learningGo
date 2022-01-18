package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simpleRest/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	// & represents the address so here we are passing the session and returning address of that session variable
	return &UserController{s}
}

//Read about struct methods in GoLang

//we defined struct so now we can create methods for controller with the session like getUser and so on.
//basically we created a custom class (struct) and instantiated it which stores the value of the pointer which contains the session with mongoDb

//writing controller methods with the retrieved session id

//These are struct methods
//GetUser
//CreateUser

//this is a struct method, this method is extending struct UserController (uc UserController) which can be used to access members of the given struct
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	//returns where the parameter is valid hex representation of the object
	if !bson.IsObjectIdHex(id) {
		//sending message that it does not exist
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	//.C makes collection of users
	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return 
	}

	//found data in u and then marshal it and put it in uj and pass it to front-end
	//NOTE: Read about marshal and unmarshal
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//fprintf - formats according to the formatter and writes to the 'w'
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(u)

	//converting it to Json after storing it to DB.
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s/n", uj)
}
