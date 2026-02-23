package handlers

import (
	"ecommerce/config"
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

	cnf := config.GetConfig()
	accessToken, err := utils.CreateJWT(cnf.JwtSecret, utils.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.SendJSONData(w, accessToken, http.StatusOK)
}
