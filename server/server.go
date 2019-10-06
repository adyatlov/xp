package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"

	"github.com/adyatlov/bunxp/objects"
)

type Server struct {
	explorer *objects.Explorer
	box      *packr.Box
}

func New(e *objects.Explorer) *Server {
	s := &Server{}
	s.explorer = e
	s.box = packr.New("client", "../client/build")
	return s
}

func (s *Server) Serve() error {
	r := mux.NewRouter()
	r.PathPrefix("/client").Handler(http.StripPrefix("/client", http.FileServer(s.box)))
	r.HandleFunc("/api/objects/cluster", s.cluster)
	r.HandleFunc("/api/objects/{type}/{id}", s.object)
	r.HandleFunc("/api/objectTypes", s.objectTypes)
	r.HandleFunc("/api/metricTypes", s.metricTypes)
	r.HandleFunc("/", s.redirectToClient)
	return http.ListenAndServe(fmt.Sprintf("%v:%v", "localhost", "7777"), r)
}

func (s *Server) redirectToClient(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/client", http.StatusTemporaryRedirect)
}

func (s *Server) cluster(w http.ResponseWriter, r *http.Request) {
	t := objects.ObjectTypeName("cluster")
	s.objectTypeId(t, "", w, r)
}

func (s *Server) object(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := objects.ObjectTypeName(vars["type"])
	id := objects.ObjectId(vars["id"])
	s.objectTypeId(t, id, w, r)
}

func (s *Server) objectTypes(w http.ResponseWriter, r *http.Request) {
	types := objects.ObjectTypes()
	write(types, w)
}

func (s *Server) metricTypes(w http.ResponseWriter, r *http.Request) {
	types := objects.MetricTypes()
	write(types, w)
}

func (s *Server) objectTypeId(t objects.ObjectTypeName, id objects.ObjectId, w http.ResponseWriter, r *http.Request) {
	object, err := s.explorer.Object(t, id)
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
	if _, err := fmt.Fprint(w, string(objectBytes)); err != nil {
		log.Printf("Error occurred when sending response: %v", err)
	}
}
