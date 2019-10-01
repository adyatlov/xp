package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adyatlov/bunxp/objects"
)

type Server struct {
	explorer *objects.Explorer
}

func New(e *objects.Explorer) *Server {
	return &Server{
		explorer: e,
	}
}

func (e *Server) Serve() error {
	r := mux.NewRouter()
	r.HandleFunc("/objects/cluster", e.cluster)
	r.HandleFunc("/objects/{type}/{id}", e.object)
	r.HandleFunc("/objectTypes", e.objectTypes)
	r.HandleFunc("/metricTypes", e.metricTypes)
	return http.ListenAndServe(fmt.Sprintf("%v:%v", "localhost", "7777"), r)
}

func (e *Server) cluster(w http.ResponseWriter, r *http.Request) {
	t := objects.ObjectTypeName("cluster")
	e.objectTypeId(t, "", w, r)
}

func (e *Server) object(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := objects.ObjectTypeName(vars["type"])
	id := objects.ObjectId(vars["id"])
	e.objectTypeId(t, id, w, r)
}

func (e *Server) objectTypes(w http.ResponseWriter, r *http.Request) {
	types := objects.ObjectTypes()
	write(types, w)
}

func (e *Server) metricTypes(w http.ResponseWriter, r *http.Request) {
	types := objects.MetricTypes()
	write(types, w)
}

func (e *Server) objectTypeId(t objects.ObjectTypeName, id objects.ObjectId, w http.ResponseWriter, r *http.Request) {
	object, err := e.explorer.Object(t, id)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("Error: cannot get object: %v\n", err.Error()),
			http.StatusNotFound)
		return
	}
	write(object, w)
}

func write(i interface{}, w http.ResponseWriter) {
	objectBytes, err := json.Marshal(i)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("Error: cannot encode object: %v\n", err.Error()),
			http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprint(w, string(objectBytes)); err != nil {
		log.Printf("Error occurred when sending response: %v", err)
	}
}
