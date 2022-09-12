package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"transporte/src/controller"
	"transporte/src/library/sqlquery"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type jwtclaim struct {
	User   string `json:"user"`
	Nombre string `json:"nombre"`
	Cargo  int64  `json:"cargo"`
	jwt.StandardClaims
}

func RutasAuth(r *mux.Router) {

	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/", auth).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")

}

func auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()

	// Get the request body
	req_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//objeto map
	body := make(map[string]interface{})
	//leer the body y convertir en objeto
	json.Unmarshal(req_body, &body)

	dataUser := sqlquery.NewQuerys("Seguridad").Select().Where("users", "=", body["users"]).Exec().One()
	if len(dataUser) <= 0 {
		response.Msg = "Usuario y contrasenia Incorrecto"
		response.StatusCode = 300
		response.Status = "Usuario y contrasenia Incorrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//int8, int64, int32

	err = bcrypt.CompareHashAndPassword([]byte(dataUser["l_pass"].(string)), []byte(body["password"].(string)))
	if err != nil {
		response.Msg = "Contraseña Incorrecta o Usuario Incorrecto"
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	//Tokens
	var key_token interface{}
	key_token = []byte("supervisor")
	claims := jwtclaim{
		dataUser["users"].(string),
		dataUser["l_nomb"].(string),
		dataUser["k_carg"].(int64),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * (60 * 24)).Unix(),
			Issuer:    "pdt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err_token := token.SignedString(key_token)
	if err_token != nil {
		response.Msg = "Error signing" + err_token.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["token"] = token_string
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
