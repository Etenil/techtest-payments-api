package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/etenil/techtest-payments-api/models"
)

type PaymentsApi struct {
	paymentModel *models.PaymentModel
}

type ApiError struct {
	Error string `json:"error"`
}

func main() {
	log.Print("API started")

	paymentModel := models.NewPaymentModel()

	api := &PaymentsApi{
		paymentModel: paymentModel,
	}

	api.start()
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s ", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func apiError(w http.ResponseWriter, s int, err string) {
	log.Print(err)
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(&ApiError{Error: err})
}

func (api *PaymentsApi) start() {
	router := mux.NewRouter()
	// Access Logs
	router.Use(accessLogMiddleware)

	router.HandleFunc("/payments", api.ListPayments).Methods("GET")
	router.HandleFunc("/payments", api.CreatePayment).Methods("POST")

	router.HandleFunc("/payments/{id}", api.GetPayment).Methods("GET")
	router.HandleFunc("/payments/{id}", api.UpdatePayment).Methods("PUT")
	router.HandleFunc("/payments/{id}", api.DeletePayment).Methods("DELETE")

	// CORS boilerplate, allow everything everywhere
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "X-Owner-ID"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err := http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods)(router))
	if err != nil {
		log.Fatalf("Failed to start HTTP server on at localhost:8080")
	}
}

func (api *PaymentsApi) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var p *models.Payment
	json_err := json.NewDecoder(r.Body).Decode(&p)

	if json_err != nil {
		apiError(w, http.StatusBadRequest, "error parsing payment data")
		return
	}

	p.Id = -1
	p.Currency = "GBP"

	err := api.paymentModel.CreatePayment(p)

	if err != nil {
		apiError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (api *PaymentsApi) ListPayments(w http.ResponseWriter, r *http.Request) {
	p := api.paymentModel.GetPayments()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (api *PaymentsApi) GetPayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, converr := strconv.Atoi(vars["id"])

	if converr != nil {
		apiError(w, http.StatusBadRequest, "invalid payment id provided")
		return
	}

	p, err := api.paymentModel.GetPaymentById(id)

	if err != nil {
		apiError(w, http.StatusNotFound, "payment resource not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (api *PaymentsApi) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, converr := strconv.Atoi(vars["id"])

	if converr != nil {
		log.Print("Invalid payment ID provided")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	p, p_err := api.paymentModel.GetPaymentById(id)

	if p_err != nil {
		apiError(w, http.StatusNotFound, "payment resource not found")
		return
	}

	json_err := json.NewDecoder(r.Body).Decode(&p)

	if json_err != nil {
		apiError(w, http.StatusBadRequest, "invalid payment data")
		return
	}

	p.Id = id
	p.Currency = "GBP"

	err := api.paymentModel.UpdatePayment(p)

	if err != nil {
		apiError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (api *PaymentsApi) DeletePayment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, converr := strconv.Atoi(vars["id"])

	if converr != nil {
		log.Print("Invalid payment ID provided")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	p, err := api.paymentModel.GetPaymentById(id)

	if err != nil {
		apiError(w, http.StatusNotFound, "payment resource not found")
		return
	}

	api.paymentModel.DeletePayment(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
