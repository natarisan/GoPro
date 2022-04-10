package app

import (
	"GOP/service"
	"GOP/dto"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type AccountHandler struct{
	service service.AccountService
}

func(h AccountHandler) newAccount(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		}else{
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

//まずはgetリクエストのURLから値を抽出
func(h AccountHandler)MakeTransaction(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.TransRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil{
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId
	}

	account, appError := h.service.MakeTransaction(request)

	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, account)
	}
}