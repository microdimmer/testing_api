package rest_server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/microdimmer/testing_api/internal/app/controllers"
	"github.com/microdimmer/testing_api/internal/app/models"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type RESTServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *RESTServer {
	return &RESTServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *RESTServer) Run() error {
	//init logger
	if err := s.configureLogger(); err != nil {
		return err
	}

	//init controllers
	controllers.InitControllers(s.logger)

	//init models
	if err := s.InitModels(); err != nil {
		return err
	}
	s.logger.Info("models initialized..")

	//init routes
	s.configureRouter()

	s.logger.Info("server listening at port ", s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *RESTServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	s.logger.Info("logs are writing..")

	return nil
}

func (s *RESTServer) configureRouter() {
	s.router.HandleFunc("/get-course/{link}/", controllers.GetCourse).Methods("GET")
	s.router.HandleFunc("/get-course/{link}/", controllers.StartCourse).Methods("POST")
	s.router.HandleFunc("/process-course/{link}/", controllers.ProcessCourse).Methods("POST")
	// s.router.HandleFunc("/generate-courses", s.handleGenerateCourses()).Methods("POST")
	s.logger.Info("routes initialized..")
}

func (s *RESTServer) InitModels() error {
	s.logger.Info("connecting to ", s.config.DBHost)
	dsnUrl := "host=%v port=%v user=%v password=%v dbname=%v sslmode=disable"
	dsnUrl = fmt.Sprintf(dsnUrl, s.config.DBHost, s.config.DBPort, s.config.DBUser, s.config.DBPassword, s.config.DBName)
	return models.InitConnection(dsnUrl, s.logger)
}
