package calc

import (
	"calc-api/types"
	"calc-api/utils"

	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /add", h.Add)
	router.HandleFunc("POST /subtract", h.Subtract)
	router.HandleFunc("POST /multiply", h.Multiply)
	router.HandleFunc("POST /divide", h.Divide)
	router.HandleFunc("POST /sum", h.Sum)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	utils.LogCalcReceived(r)

	// Check if content type is application/json
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		utils.WriteError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
		return
	}

	values := types.CalculatorRequest{}
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		fmt.Println(err)
		switch {
		case errors.Is(err, io.EOF):
			utils.WriteError(w, http.StatusBadRequest, "Request body is empty")
		case errors.Is(err, io.ErrUnexpectedEOF):
			utils.WriteError(w, http.StatusBadRequest, "Request body contains invalid JSON")
		default:
			utils.WriteError(w, http.StatusBadRequest, "Body not in correct format")
		}
		return
	}

	// validate values
	if err := validateValues(values); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Process the request
	result := types.CalculatorResponse{Result: values.First + values.Second}
	utils.WriteJSON(w, http.StatusOK, result)
	utils.LogCalcCompleted(values, result)
}

func (h *Handler) Subtract(w http.ResponseWriter, r *http.Request) {
	utils.LogCalcReceived(r)

	// Check if content type is application/json
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		utils.WriteError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
		return
	}

	values := types.CalculatorRequest{}
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		fmt.Println(err)
		switch {
		case errors.Is(err, io.EOF):
			utils.WriteError(w, http.StatusBadRequest, "Request body is empty")
		case errors.Is(err, io.ErrUnexpectedEOF):
			utils.WriteError(w, http.StatusBadRequest, "Request body contains invalid JSON")
		default:
			utils.WriteError(w, http.StatusBadRequest, "Body not in correct format")
		}
		return
	}

	// validate values
	if err := validateValues(values); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Process the request
	result := types.CalculatorResponse{Result: values.First - values.Second}
	utils.WriteJSON(w, http.StatusOK, result)
	utils.LogCalcCompleted(values, result)
}
func (h *Handler) Multiply(w http.ResponseWriter, r *http.Request) {
	utils.LogCalcReceived(r)

	// Check if content type is application/json
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		utils.WriteError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
		return
	}

	values := types.CalculatorRequest{}
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		fmt.Println(err)
		switch {
		case errors.Is(err, io.EOF):
			utils.WriteError(w, http.StatusBadRequest, "Request body is empty")
		case errors.Is(err, io.ErrUnexpectedEOF):
			utils.WriteError(w, http.StatusBadRequest, "Request body contains invalid JSON")
		default:
			utils.WriteError(w, http.StatusBadRequest, "Body not in correct format")
		}
		return
	}

	// validate values
	if err := validateValues(values); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Process the request
	result := types.CalculatorResponse{Result: values.First * values.Second}
	utils.WriteJSON(w, http.StatusOK, result)
	utils.LogCalcCompleted(values, result)
}
func (h *Handler) Divide(w http.ResponseWriter, r *http.Request) {

	// Check if content type is application/json
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		utils.WriteError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
		return
	}

	values := types.CalculatorRequest{}
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		fmt.Println(err)
		switch {
		case errors.Is(err, io.EOF):
			utils.WriteError(w, http.StatusBadRequest, "Request body is empty")
		case errors.Is(err, io.ErrUnexpectedEOF):
			utils.WriteError(w, http.StatusBadRequest, "Request body contains invalid JSON")
		default:
			utils.WriteError(w, http.StatusBadRequest, "Body not in correct format")
		}
		return
	}

	// validate values
	if err := validateValues(values); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Process the request
	result := types.CalculatorResponse{Result: values.First / values.Second}
	utils.WriteJSON(w, http.StatusOK, result)

	utils.LogCalcCompleted(values, result)
}
func (h *Handler) Sum(w http.ResponseWriter, r *http.Request) {
	utils.LogCalcReceived(r)

	// Check if content type is application/json
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		utils.WriteError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
		return
	}

	body, _ := io.ReadAll(r.Body)
	var sumData []int
	json.Unmarshal([]byte(body), &sumData)

	// validate values
	if err := validateSumValues(sumData); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Process the request
	result := types.CalculatorResponse{Result: sum(sumData)}
	utils.WriteJSON(w, http.StatusOK, result)
	slog.Info("calculation completed",
		slog.Int("result", result.Result),
		slog.Int("status", http.StatusOK),
	)
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func validateValues(values types.CalculatorRequest) error {
	var first, second int = values.First, values.Second

	// values cannot be float
	if (first%1 != 0) || (second%1 != 0) {
		return errors.New("values cannot be float")
	}

	// values cannot be negative
	if (first < 0) || (second < 0) {
		return errors.New("values cannot be negative")
	}

	return nil
}

func validateSumValues(values []int) error {
	for _, v := range values {
		if v%1 != 0 {
			return errors.New("values cannot be float")
		}
		if v < 0 {
			return errors.New("values cannot be negative")
		}
	}
	return nil
}
