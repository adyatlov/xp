package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/adyatlov/bunxp/explorer"
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
)

type Server struct {
	box      *packr.Box
	explorer *explorer.Explorer
}

func New(e *explorer.Explorer) *Server {
	s := &Server{}
	s.explorer = e
	s.box = packr.New("client", "../client/build")
	return s
}

func (s *Server) Serve() error {
	schema := graphql.MustParseSchema(schemaString,
		&resolver{explorer: s.explorer})
	r := mux.NewRouter()
	r.PathPrefix("/client").Handler(http.StripPrefix("/client", http.FileServer(s.box)))
	r.Handle("/query", &relay.Handler{Schema: schema})
	r.HandleFunc("/", s.redirectToClient)
	corsHandler := cors.Default().Handler(r)
	return http.ListenAndServe(fmt.Sprintf("%v:%v", "localhost", "7777"), corsHandler)
}

func (s *Server) redirectToClient(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/client", http.StatusTemporaryRedirect)
}

func (s *Server) cluster(w http.ResponseWriter, r *http.Request) {
	t := explorer.ObjectTypeName("cluster")
	s.objectTypeId(t, "", w, r)
}

func (s *Server) object(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := explorer.ObjectTypeName(vars["type"])
	id := explorer.ObjectId(vars["id"])
	s.objectTypeId(t, id, w, r)
}

func (s *Server) objectTypes(w http.ResponseWriter, r *http.Request) {
	types := explorer.GetObjectTypes()
	write(types, w)
}

func (s *Server) metricTypes(w http.ResponseWriter, r *http.Request) {
	types := explorer.GetMetricTypes()
	write(types, w)
}

func (s *Server) objectTypeId(t explorer.ObjectTypeName, id explorer.ObjectId, w http.ResponseWriter, r *http.Request) {
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
