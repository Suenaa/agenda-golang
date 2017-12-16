package service

import (
	"encoding/json"
	"net/http"

	"github.com/unrolled/render"
	"github.com/Suenaa/agenda-golang/service/entities"
)

func userLoginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		err := UserLogin(req.FormValue("username"),req.FormValue("password"))
		check(err)
		if err == nil {
			formatter.JSON(w, http.StatusOK, struct {
			Name      string `json:"name"`
			Status string `json:"status"`
			}{Name: req.FormValue("username"), Status: "Is login"})
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func createNewUserHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		decoder := json.NewDecoder(req.Body)
		var user entities.User
		err := decoder.Decode(&user)
		check(err)
		err = UserRegister(user.Username, user.Password, user.Email, user.Phone)
		if err == nil {
			formatter.JSON(w, http.StatusCreated, user)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func listAllUsersHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		userList := QueryAllUsers()
		for i := range userList {
			userList[i].Password = "******"
		}
		formatter.JSON(w, http.StatusOK, userList)
	}

}

func deleteCurrentUserHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		err := DeleteUser(req.FormValue("password"))
		check(err)
		if err == nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func listAllMeetingsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		meetingList, err := QueryMeeting(req.FormValue("startdate"), req.FormValue("enddate"))
		check(err)
		if err == nil {
			formatter.JSON(w, http.StatusOK, meetingList)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func createNewMeetingHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		decoder := json.NewDecoder(req.Body)
		var meeting entities.Meeting
		err := decoder.Decode(&meeting)
		check(err)
		err = CreateMeeting(meeting.Title, meeting.Start, meeting.End, meeting.Participators)
		check(err)
		if err == nil {
			formatter.JSON(w, http.StatusOK, meeting)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func clearMeetingsHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		err := DeleteAllMeeting()
		check(err)
		if err == nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
