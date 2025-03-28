package rest_err

import (
	"net/http"
)

// Estrutura para erros REST
type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int64    `json:"code"`
	Causes  []Causes `json:"causes,omitempty"` // `omitempty` evita JSON com array vazio
}

// Estrutura para detalhar os erros
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Função para criar um erro REST genérico
func NewRestErr(message, err string, code int64, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// 400 - Bad Request
func NewBadRequestError(message string) *RestErr {
	return NewRestErr(message, "bad request", http.StatusBadRequest, nil)
}

// 400 - Bad Request com validação
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, "bad request", http.StatusBadRequest, causes)
}

// 401 - Unauthorized
func NewUnauthorizedError(message string) *RestErr {
	return NewRestErr(message, "unauthorized", http.StatusUnauthorized, nil)
}

// 403 - Forbidden com validação
func NewForbiddenValidationError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, "forbidden", http.StatusForbidden, causes)
}

// 403 - Forbidden
func NewForbiddenError(message string) *RestErr {
	return NewRestErr(message, "forbidden", http.StatusForbidden, nil)
}

// 500 - Internal Server Error
func NewInternalServerError(message string) *RestErr {
	return NewRestErr(message, "internal server error", http.StatusInternalServerError, nil)
}
