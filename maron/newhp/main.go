package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"math/rand"

	"github.com/google/uuid" 
)

type ApiErr struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}

func (a ApiErr) Error() string {
	return a.Err
}

type AddService interface {
	Add(uuid.UUID) (uuid.UUID, float64, error)
}

type addService struct {
}

func NewAddService() AddService {
	return &addService{}
}

func (as *addService) Add(id uuid.UUID) (uuid.UUID, float64, error) {
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	return uuid.New(), 69.69 + n, nil
}

type AddRequest struct {
	AdPlacementID uuid.UUID `json:"addPlacementID"`
}

type AddResponse struct {
	AddID    uuid.UUID `json:"addID"`
	BidPrice float64   `json:"bidPrice"`
}

func main() {
	router := http.NewServeMux()
	svc := NewAddService()
	h := NewAddRequestHandler(svc)
	router.HandleFunc("/add", TransHttpHandlefunc(h.handleAddRequest))
	http.ListenAndServe(":40001", router)
}

type addRequestHandler struct {
	svc AddService
}

func NewAddRequestHandler(svc AddService) *addRequestHandler {
	return &addRequestHandler{
		svc: svc,
	}
}

func (h addRequestHandler) handleAddRequest(w http.ResponseWriter, r *http.Request) error {
	addID := uuid.New()
	id, bidprice, err := h.svc.Add(addID)
	if err != nil {
		slog.Error("add service returned non 200 response", "err", err)
		return WriteJSON(w, http.StatusNoContent, map[string]any{"error": err.Error()})
	}
	resp := AddResponse{
		AddID:    id,
		BidPrice: bidprice,
	}
	return WriteJSON(w, http.StatusOK, resp)
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func TransHttpHandlefunc(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("error", "err", err)
			WriteJSON(w, http.StatusInternalServerError, ApiErr{
				Status: http.StatusInternalServerError,
				Err:    err.Error(),
			})
		}
	}
}

//	func HandleAddRequest(w http.ResponseWriter, r *http.Request) error {
//		resp := AddResponse{
//			BidPrice: 69.69,
//			AddID:    uuid.New(),
//		}
//		return WriteJSON(w, http.StatusOK, resp)
//	}
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&v)
}
