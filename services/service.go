package services

import (
	"context"
	"fmt"
	"github.com/avishkar-activex/chms-auth/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type WebService struct {
	Router *mux.Router
	Server *http.Server
}

func NewWebService() *WebService {

	ws := WebService{}
	ws.Router = mux.NewRouter()

	routes.Register(ws.Router)

	ws.Server = &http.Server{Addr: ":" + "8001", Handler: enableCORS()(ws.Router)}

	return &ws
}

func enableCORS() func(http.Handler) http.Handler {

	allowCred := handlers.AllowCredentials()
	headersOk := handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Length", "Content-Type", "Content-Disposition", "Accept", "Origin", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	return handlers.CORS(allowCred, headersOk, originsOk, methodsOk)
}

func (ws *WebService) Run() error {
	//fmt.Printf("auth service starting on port: %s", viper.GetString("services.scheduling_svc.port"))
	return ws.Server.ListenAndServe()
}

func (ws *WebService) Start() {
	//go func() {
	err := ws.Run()
	if err != nil && err.Error() != "http: server closed" {
		log.Fatal("could not start webservice", err)
	}
	//}()
}

func (ws *WebService) Shutdown() error {
	fmt.Println("shutting down scheduling service")
	return ws.Server.Shutdown(context.Background())
}
