package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPserver struct {
	httpHandlers *HTTPhandlers
}

func NewHTTPServer(httpHandler *HTTPhandlers) *HTTPserver {
	return &HTTPserver{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPserver) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetAllTasks)
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandleGetAlluncompletedTasks)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleCompleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		} else {
			return err
		}
	}

	return nil

}
