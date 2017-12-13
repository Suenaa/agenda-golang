package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)


func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{IndentJSON: true})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {

	//user login
	mx.HandleFunc("/user/login", userLoginHandler(formatter)).Methods("GET")

	//Create a new user
	mx.HandleFunc("/user/register", createNewUserHandler(formatter)).Methods("POST")

	//List all Users
	mx.HandleFunc("/users/allusers", listAllUsersHandler(formatter)).Methods("GET")

	//Delete current user
	mx.HandleFunc("/users/deleteuser", deleteCurrentUserHandler(formatter)).Methods("DELETE")

	//List all meetings
	mx.HandleFunc("/meetings/allmeetings", listAllMeetingsHandler(formatter)).Methods("GET")

	//Create a new meeting
	mx.HandleFunc("/meetings/newmeeting", createNewMeetingHandler(formatter)).Methods("POST")

	//Clear all meetings
	mx.HandleFunc("/meetings/deletemeetings", clearMeetingsHandler(formatter)).Methods("DELETE")
}
