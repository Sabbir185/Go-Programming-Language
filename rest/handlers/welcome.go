package handlers

import (
	"ecommerce/utils"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONData(w, "Welcome to the Go Land...", 200)
}
