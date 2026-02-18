package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var LoginData LoginPayload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&LoginData)
	if err != nil {
		http.Error(w, "Invalid Requested Data", http.StatusBadRequest)
		return
	}
	user := database.MatchUserCredentials(LoginData.Email, LoginData.Password)
	fmt.Println(user)
	if user == nil {
		http.Error(w, "Invalid user credentials", http.StatusBadRequest)
		return
	}
	utils.SendJSONData(w, user, http.StatusOK)
}
