package routes

import (
	"encoding/json"
	"github.com/avishkar-activex/chms-auth/services/auth"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func Register(r *mux.Router) {
	r.HandleFunc("/healthcheck", healthCheck).Methods(http.MethodGet)
	r.HandleFunc("/v1/auth", authenticate).Methods(http.MethodPost)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

type AuthRequestBody struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserId    int64  `json:"user_id"`
	Name      string `json:"name"`
	AccountId int64  `json:"account_id"`
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	var body AuthRequestBody

	readBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("failed to read request")
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	err = json.Unmarshal(readBody, &body)
	if err != nil {
		log.Errorf("failed to unmarshal request")
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	defer r.Body.Close()

	usr, err := auth.AuthenticateUser(body.UserName, body.Password)
	if err != nil {
		log.Errorf("failed to autheticate user")
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
	response := AuthResponse{UserId: usr.Id, AccountId: usr.AccountId, Name: usr.Name}
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "response marshal error", http.StatusInternalServerError)
	}

	w.Write(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
