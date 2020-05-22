package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adyatlov/xp/gql"

	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
)

type Server struct {
	box    *packr.Box
	schema *gql.Schema
}

func New() *Server {
	s := &Server{}
	s.schema = gql.NewSchema()
	s.box = packr.New("client", "../client/build")
	return s
}

func (s *Server) Serve() error {
	schema := graphql.MustParseSchema(gql.SchemaString, s.schema)
	r := mux.NewRouter()
	r.PathPrefix("/client").Handler(http.StripPrefix("/client", http.FileServer(s.box)))
	r.Handle("/graphql", &DebugHandler{&relay.Handler{Schema: schema}})
	r.HandleFunc("/", s.redirectToClient)
	corsHandler := cors.Default().Handler(r)
	return http.ListenAndServe(fmt.Sprintf("%v:%v", "localhost", "7777"), corsHandler)
}

func (s *Server) redirectToClient(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/client", http.StatusTemporaryRedirect)
}

type DebugHandler struct {
	handler http.Handler
}

func (d *DebugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request :)")
	d.handler.ServeHTTP(w, r)
}
