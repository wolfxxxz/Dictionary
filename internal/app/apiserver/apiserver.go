package apiserver

import (
	"net/http"

	"github.com/Wolfxxxz/Dictionary/internal/app/middleware"
	"github.com/Wolfxxxz/Dictionary/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

// type for APIServer object for instancing server
type APIServer struct {
	//Unexported field
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// APIServer constructor
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server and connection to db and logger confs
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting api server at port :", s.config.BindAddr)
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// func for configureate logger, should be unexported
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return nil
	}
	s.logger.SetLevel(level)

	return nil
}

// func for configure Router
func (s *APIServer) configureRouter() {
	s.router.HandleFunc(prefix+"/words", s.GetAllWords).Methods("GET")         // http://localhost:8081/api/v1/words
	s.router.HandleFunc(prefix+"/words"+"/{id}", s.GetWordById).Methods("GET") // http://localhost:8081/api/v1/words/2
	//s.router.HandleFunc(prefix+"/words"+"/{id}", s.DeleteWordById).Methods("DELETE") // http://localhost:8081/api/v1/words/1
	s.router.Handle(prefix+"/words"+"/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.DeleteWordById))).Methods("DELETE") // http://localhost:8081/api/v1/words/1
	s.router.HandleFunc(prefix+"/words/register", s.PostWord).Methods("POST")             // http://localhost:8081/api/v1/words/register
	s.router.HandleFunc(prefix+"/words/registermany", s.PostAllWords).Methods("POST")     // http://localhost:8081/api/v1/words/registermany
	s.router.HandleFunc(prefix+"/words/registerAllWords", s.PostAllWords).Methods("POST") // http://localhost:8081/api/v1/words/registerAllWords
	s.router.HandleFunc(prefix+"/user/register", s.PostUserRegister).Methods("POST")      // http://localhost:8081/api/v1/user/register
	//new pair for auth / middleware /
	s.router.HandleFunc(prefix+"/user/auth", s.PostToAuth).Methods("POST") // http://localhost:8081/api/v1/user/auth

	/*
		s.router.HandleFunc(prefix+"/articles", s.GetAllArticles).Methods("GET")
		//s.router.Handle(prefix+"/articles", middleware.JwtMiddleware.Handler(
		//	http.HandlerFunc(s.GetAllArticles))).Methods("GET")
		//Было до jwt token
		s.router.HandleFunc(prefix+"/articles"+"/{id}", s.GetArticleById).Methods("GET")
		//Теперь требует токена
		// s.router.Handle(prefix+"/articles"+"/{id}", middleware.JwtMiddleware.Handler(
		// 	http.HandlerFunc(s.GetArticleById))).Methods("GET")
		s.router.HandleFunc(prefix+"/articles"+"/{id}", s.DeleteArticleById).Methods("DELETE")
		// s.router.Handle(prefix+"/articles"+"/{id}", middleware.JwtMiddleware.Handler(
		// 	http.HandlerFunc(s.DeleteArticleById))).Methods("DELETE")
		s.router.HandleFunc(prefix+"/articles", s.PostArticle).Methods("POST")
		// s.router.Handle(prefix+"/articles", middleware.JwtMiddleware.Handler(
		// 	http.HandlerFunc(s.PostArticle))).Methods("POST")
		s.router.HandleFunc(prefix+"/articles"+"/{id}", s.PutArticleById).Methods("PUT")
		// s.router.Handle(prefix+"/articles"+"/{id}", middleware.JwtMiddleware.Handler(
		// 	http.HandlerFunc(s.PutArticleById))).Methods("PUT")
		s.router.HandleFunc(prefix+"/user/register", s.PostUserRegister).Methods("POST")
		// s.router.Handle(prefix+"/user/register", middleware.JwtMiddleware.Handler(
		// 	http.HandlerFunc(s.PostUserRegister))).Methods("POST")
		//new pair for auth / middleware /
		s.router.HandleFunc(prefix+"/user/auth", s.PostToAuth).Methods("POST")
	*/
}

// configureStore method
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
