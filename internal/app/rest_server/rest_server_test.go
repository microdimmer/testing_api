package rest_server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/microdimmer/testing_api/internal/app/controllers"
	"github.com/microdimmer/testing_api/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestRESTServer_GetCourse(t *testing.T) {
	initTest()

	course_exp := &models.Course_expiring{}
	course_exp.ID = 1
	course_exp.CourseID = 1
	course_exp.Link = "e8712100-600b-4a85-b4c3-87d33d2171ef"
	course_exp.Expiring = time.Now().Add(1 * time.Hour)
	course_exp.Create()

	payload := `{"ID":1,"CreatedAt":"2021-05-25T20:00:00Z","UpdatedAt":"2021-05-25T20:00:00Z","DeletedAt":null,"name":"Промышленная безопасность"}`

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/get-course/{link}/", nil)

	vars := map[string]string{
		"link": "e8712100-600b-4a85-b4c3-87d33d2171ef",
	}
	controllers.GetCourse(rec, mux.SetURLVars(req, vars))
	assert.Equal(t, payload, rec.Body.String())
}

func TestRESTServer_StartCourse(t *testing.T) {
	initTest()

	course_exp := &models.Course_expiring{}
	course_exp.ID = 1
	course_exp.CourseID = 1
	course_exp.Link = "e8712100-600b-4a85-b4c3-87d33d2171ef"
	course_exp.Expiring = time.Now().Add(1 * time.Hour)
	course_exp.Create()

	rec := httptest.NewRecorder()
	payload := `{"empl_name": "Vasya","empl_dep": "Kykyshechka"}`
	req, _ := http.NewRequest(http.MethodPost, "/get-course/{link}/", bytes.NewBuffer([]byte(payload)))

	vars := map[string]string{
		"link": "e8712100-600b-4a85-b4c3-87d33d2171ef",
	}
	controllers.StartCourse(rec, mux.SetURLVars(req, vars))
	answer := `{"message":"Course attemt from Vasya Kykyshechka started"}`
	assert.Equal(t, answer, rec.Body.String())

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/get-course/{link}/", bytes.NewBuffer([]byte(payload)))
	vars = map[string]string{
		"link": "e8712100-600b-4a85-b4c3-87d33d2171ef",
	}
	controllers.StartCourse(rec, mux.SetURLVars(req, vars))
	answer = `{"message":"Course is already started"}`
	assert.Equal(t, answer, rec.Body.String())
}

func TestRESTServer_ProcessCourse(t *testing.T) {
	initTest()

	course_exp := &models.Course_expiring{}
	course_exp.ID = 1
	course_exp.CourseID = 1
	course_exp.Link = "e8712100-600b-4a85-b4c3-87d33d2171ef"
	course_exp.Expiring = time.Now().Add(1 * time.Hour)
	course_exp.Passing = 2
	course_exp.Create()

	rec := httptest.NewRecorder()
	payload := `{"passing": }`
	req, _ := http.NewRequest(http.MethodPost, "/process-course/{link}/", bytes.NewBuffer([]byte(payload)))

	vars := map[string]string{
		"link": "e8712100-600b-4a85-b4c3-87d33d2171ef",
	}
	controllers.ProcessCourse(rec, mux.SetURLVars(req, vars))
	answer := `{"error":"Wrong payload"}`
	assert.Equal(t, answer, rec.Body.String())

	rec = httptest.NewRecorder()
	payload = `{"passing": 0}`
	req, _ = http.NewRequest(http.MethodPost, "/process-course/{link}/", bytes.NewBuffer([]byte(payload)))

	controllers.ProcessCourse(rec, mux.SetURLVars(req, vars))
	answer = `{"message":"Cannot pass course"}`
	assert.Equal(t, answer, rec.Body.String())

	rec = httptest.NewRecorder()
	payload = `{"passing": 5}`
	req, _ = http.NewRequest(http.MethodPost, "/process-course/{link}/", bytes.NewBuffer([]byte(payload)))

	controllers.ProcessCourse(rec, mux.SetURLVars(req, vars))
	answer = `{"message":"Course passed"}`
	assert.Equal(t, answer, rec.Body.String())

	rec = httptest.NewRecorder()
	payload = `{"passing": 10}`
	req, _ = http.NewRequest(http.MethodPost, "/process-course/{link}/", bytes.NewBuffer([]byte(payload)))

	controllers.ProcessCourse(rec, mux.SetURLVars(req, vars))
	answer = `{"message":"Course finished"}`
	assert.Equal(t, answer, rec.Body.String())
}

func initTest() {
	s := New(NewConfig())
	controllers.InitControllers(s.logger)
	s.InitModels()

	models.ClearTables()

	course := &models.Course{}
	course.ID = 1
	course.Name = "Промышленная безопасность"
	course.CreatedAt = time.Date(2021, time.May, 25, 20, 0, 0, 0, &time.Location{})
	course.UpdatedAt = time.Date(2021, time.May, 25, 20, 0, 0, 0, &time.Location{})
	course.Create()
}
