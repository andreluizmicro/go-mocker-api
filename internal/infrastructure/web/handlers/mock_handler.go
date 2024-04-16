package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andreluizmicro/go-mocker-api/internal/application"
)

type ResponseData struct {
	Message string `json:"id"`
}

type MockHandler struct {
	mockService *application.MockService
}

func NewMockHandler(mockService *application.MockService) *MockHandler {
	return &MockHandler{
		mockService: mockService,
	}
}

func (m *MockHandler) Create(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro interno do servidor"))
		return
	}

	sessionId, err := m.mockService.Create(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao salvar mock"))
		fmt.Println(err)
		return
	}

	responseData := ResponseData{Message: *sessionId}
	jsonResponse, err := json.Marshal(responseData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (m *MockHandler) Find(w http.ResponseWriter, r *http.Request) {
	sessionId := r.PathValue("session_id")

	file, err := m.mockService.Find(sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data interface{}

	if err := json.Unmarshal(file, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
