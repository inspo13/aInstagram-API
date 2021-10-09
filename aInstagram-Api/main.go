package main

import (
	"fmt"
	"net/http"
	"time"

	//"log"
	"github.com/inspo13/aInstagram-Api/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	//check and print the DB connectivity
	if connectToMongo() {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connected")
	}
	//creating new instance of http router and assign it to r
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())
	//Requests using GET should only be used to
	//request data (they shouldn't include data).

	//The POST method is used to request that the origin server
	//accept the entity enclosed  in the request as a new subordinate
	// of the resource identified by the Request-URI in the Request-Line.

	//create the user details
	r.POST("/users", uc.CreateUser)
	//get the user details
	r.GET("/users/:id", uc.GetUser)

	//create the post
	r.POST("/posts", pc.CreatePost)
	//read or get the posts created by th user
	r.GET("/posts/:id", pc.GetPost)

	//get all the post created by user with given id
	r.GET("/posts/users/:id", pc.GetAllPost)
	//helps in deleting the user info
	r.DELETE("/users/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)

}

//checking DB connectivity
func connectToMongo() bool {
	ret := false
	fmt.Println("enter main - connecting to mongo")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Detected panic")
			var ok bool
			err, ok := r.(error)
			if !ok {
				fmt.Printf("pkg:  %v,  error: %s", r, err)
			}
		}
	}()

	maxWait := time.Duration(5 * time.Second)
	session, sessionErr := mgo.DialWithTimeout("localhost:27017", maxWait)
	if sessionErr == nil {
		session.SetMode(mgo.Monotonic, true)
		coll := session.DB("MyDB").C("MyCollection")
		if coll != nil {
			fmt.Println("Got a collection object")
			ret = true
		}
	} else { // never gets here
		fmt.Println("Unable to connect to local mongo instance!")
	}
	return ret
}

// whatever is in getSession will be passed to NewUserController and NewPostController
func getSession() *mgo.Session {
	//server connectivity
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}
	return s

}
