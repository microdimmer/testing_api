package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/microdimmer/testing_api/internal/app/models"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitControllers(l *logrus.Logger) {
	logger = l
}

var GetCourse = func(w http.ResponseWriter, r *http.Request) {
	logger.Info("get-course")

	//get the UUID of course
	course_attemt := &models.Course_expiring{}
	course_attemt.Link = mux.Vars(r)["link"]

	valid := course_attemt.Validate()

	if !valid {
		respondWithMessage(w, 401, "There is no such link or it is expired")
		logger.Warn("There is no such link or it is expired")
		return
	}

	course := course_attemt.FindCourse()

	if course == nil {
		respondWithMessage(w, 400, "There is not such course")
		return
	}

	respondWithJSON(w, 200, course)
}

var StartCourse = func(w http.ResponseWriter, r *http.Request) {
	logger.Info("start-course")
	course_attemt := &models.Course_expiring{}
	err := json.NewDecoder(r.Body).Decode(course_attemt)
	if err != nil {
		logger.Error(err)
		respondWithError(w, 400, "Wrong payload")
		return
	}

	//get the UUID of course
	course_attemt.Link = mux.Vars(r)["link"]

	if !course_attemt.Validate() {
		respondWithMessage(w, 403, "There is no such link or it is expired")
		return
	}

	if !course_attemt.Start() {
		respondWithMessage(w, 403, "Course is already started")
		return
	}
	respondWithMessage(w, 200, fmt.Sprintf("Course attemt from %s %s started", course_attemt.Empl_name, course_attemt.Empl_dep))
}

var ProcessCourse = func(w http.ResponseWriter, r *http.Request) {
	logger.Info("process-course")
	course_attemt := &models.Course_expiring{}
	err := json.NewDecoder(r.Body).Decode(course_attemt)
	if err != nil {
		logger.Error(err)
		respondWithError(w, 400, "Wrong payload")
		return
	}

	//get the UUID of course
	course_attemt.Link = mux.Vars(r)["link"]

	if !course_attemt.Validate() {
		respondWithMessage(w, 403, "There is no such link or it is expired")
		return
	}

	if !course_attemt.Process() {
		respondWithMessage(w, 403, "Cannot pass course")
		return
	}
	if course_attemt.Passing == 10 {
		respondWithMessage(w, 200, "Course finished")
	} else {
		respondWithMessage(w, 200, "Course passed")
	}
}

func respondWithMessage(w http.ResponseWriter, code int, message string) {
	logger.Info(message)
	respondWithJSON(w, code, map[string]string{"message": message})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	logger.Warn(message)
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
