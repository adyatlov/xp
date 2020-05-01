package server

import (
	"fmt"
	"net/http"

	"github.com/adyatlov/xp/xp"
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
)

type Server struct {
	box      *packr.Box
	explorer *xp.Explorer
}

func New(e *xp.Explorer) *Server {
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
	r.Handle("/graphql", &relay.Handler{Schema: schema})
	r.HandleFunc("/", s.redirectToClient)
	corsHandler := cors.Default().Handler(r)
	return http.ListenAndServe(fmt.Sprintf("%v:%v", "localhost", "7777"), corsHandler)
}

func (s *Server) redirectToClient(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/client", http.StatusTemporaryRedirect)
}
