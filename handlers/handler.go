package handler

import (
	"encoding/json"
	"net/http"
	"sort"

	model "github.com/mjehanno/todo-back/models"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	switch r.Method {
	case "GET":
		GetHandler(w, r)
	case "POST":
		PostHandler(w, r, decoder)
	case "PUT":
		PutHandler(w, r, decoder)
	case "DELETE":
		DeleteHandler(w, r, decoder)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(nil)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request, d *json.Decoder) {
	var t model.Task
	d.Decode(&t)
	maxId := 0
	for _, task := range model.Db {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	t.Id = maxId + 1
	model.Db = append(model.Db, t)
	w.WriteHeader(http.StatusCreated)
	w.Write(nil)

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var dbCopy []model.Task
	dbCopy = append(dbCopy, model.Db...)

	sort.Slice(dbCopy, func(i, j int) bool {
		return dbCopy[i].Priority < dbCopy[j].Priority
	})
	body, err := json.Marshal(dbCopy)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(nil)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func PutHandler(w http.ResponseWriter, r *http.Request, d *json.Decoder) {
	var t model.Task
	d.Decode(&t)

	for i, task := range model.Db {
		if task.Id == t.Id {
			model.Db[i] = t
		}
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request, d *json.Decoder) {
	var t model.Task
	d.Decode(&t)

	for i, task := range model.Db {
		if task.Id == t.Id {
			model.Db = append(model.Db[:i], model.Db[i+1:]...)
		}
	}
}
