package app

import (
	"encoding/json"
    "net/http"
    "GOP/dto"
    "GOP/service"
    "github.com/gorilla/mux"
)

type CustomerHandlers struct{
    service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request){
    status := r.URL.Query().Get("status")

    customers, err := ch.service.GetAllCustomer(status)

    if err != nil {
        writeResponse(w, err.Code, err.AsMessage())
    } else {
        writeResponse(w, http.StatusOK, customers)
    }
    
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["customer_id"]

    customer, err := ch.service.GetCustomer(id)
    if err != nil {
        writeResponse(w, err.Code, err.AsMessage())
    }else{
        writeResponse(w, http.StatusOK, customer)
    }
}

func (ch *CustomerHandlers) getImages(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
	customerId := vars["customer_id"]
    images, appError := ch.service.GetImages(customerId)
    if appError != nil {
        writeResponse(w, appError.Code, appError.AsMessage())
    } else {
        writeResponse(w, http.StatusOK, images)
    }
}

func (ch *CustomerHandlers) postImage(w http.ResponseWriter, r *http.Request) {

    var request dto.PostImageRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        writeResponse(w, http.StatusBadRequest, err.Error())
    }

    appError := ch.service.PostImage(request)

    if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, "")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}){
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(data); err != nil {
        panic(err)
    }
}


