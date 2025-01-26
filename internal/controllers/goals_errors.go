package controllers

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, httpStatusCode int, errCode, errMessage string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Add("Content-Type", "application/json")
	resp := errorResp{
		ErrCode:    "ERR00001",
		ErrMessage: "Error decoding the request",
	}
	json.NewEncoder(w).Encode(resp)
}
