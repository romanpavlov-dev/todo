package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"todo/logics"

	"github.com/gorilla/mux"
)

type HTTPhandlers struct {
	todoList *logics.List
}

func NewHTTPHandlers(todoList *logics.List) *HTTPhandlers {
	return &HTTPhandlers{
		todoList: todoList,
	}
}

/*
pattern - /tasks
Method - Post
info - JSON in http request body
*/
func (h *HTTPhandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		ErrDto := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, ErrDto.ToString(), http.StatusBadRequest)
		return
	}

	todoTask := logics.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {

		ErrDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, logics.ErrTaskNotFound) {
			http.Error(w, ErrDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, ErrDTO.ToString(), http.StatusInternalServerError)
		}

		return

	}

	b, err := json.MarshalIndent(todoTask, "", "    ")

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response")
		return
	}

}

/*
pattern - tasks/{title}
Method - GET
asnwer - JSON of found task in http request body
*/
func (h *HTTPhandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)

	if err != nil {
		ErrDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, logics.ErrTaskNotFound) {
			http.Error(w, ErrDTO.ToString(), http.StatusBadRequest)
		} else {
			http.Error(w, ErrDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response")
		return
	}
}

/*

pattern - tasks/
Method - GET
answer - JSON of all tasks in http request body

*/

func (h *HTTPhandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()

	b, err := json.MarshalIndent(tasks, "", "    ")

	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write response", err)
		return
	}
}

/*

pattern - tasks/
Method - GET
answer - JSON in http request body

*/

func (h *HTTPhandlers) HandleGetAlluncompletedTasks(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.ListUnompletedTasks()
	b, err := json.MarshalIndent(uncompletedTasks, "", "    ")

	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write response", err)
		return
	}
}

/*

pattern - tasks/{title}
Method - PATCH
asnwer - JSON of completed task in http request body

*/

func (h *HTTPhandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeDTO completeTaskDTO
	Err := json.NewDecoder(r.Body).Decode(&completeDTO)
	if Err != nil {
		errDTO := ErrorDto{
			Message: Err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	title := mux.Vars(r)["title"]

	var (
		changedTask logics.Task
		err         error
	)

	if completeDTO.Complete {
		changedTask, err = h.todoList.CompleteTask(title)
	} else {
		changedTask, err = h.todoList.UncompleteTask(title)
	}

	if err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, logics.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(changedTask, "", "    ")

	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to sent response")

		return
	}

}

/*

pattern - tasks/{title}
Method - DELETE
answer - nothing?


*/

func (h *HTTPhandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		errDTO := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, logics.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return

	}

}
