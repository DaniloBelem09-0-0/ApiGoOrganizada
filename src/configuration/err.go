package rest_err

import (
	"net/http"
)

// Estrutura padrão para erro REST
type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int    `json:"code"`
	Causes  []Causes `json:"causes,omitempty"` // Omitir se vazio
}

// Estrutura para detalhar os erros
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Função genérica para criar um erro REST
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// 400 - Bad Request
func NewBadRequestError(message string) *RestErr {
	return NewRestErr(message, "bad_request", http.StatusBadRequest, nil)
}

// 400 - Bad Request com validação
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, "bad_request", http.StatusBadRequest, causes)
}

// 401 - Unauthorized
func NewUnauthorizedError(message string) *RestErr {
	return NewRestErr(message, "unauthorized", http.StatusUnauthorized, nil)
}

// 403 - Forbidden
func NewForbiddenError(message string) *RestErr {
	return NewRestErr(message, "forbidden", http.StatusForbidden, nil)
}

// 404 - Not Found
func NewNotFoundError(message string) *RestErr {
	return NewRestErr(message, "not_found", http.StatusNotFound, nil)
}

// 409 - Conflict
func NewConflictError(message string) *RestErr {
	return NewRestErr(message, "conflict", http.StatusConflict, nil)
}

// 422 - Unprocessable Entity
func NewUnprocessableEntityError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, "unprocessable_entity", http.StatusUnprocessableEntity, causes)
}

// 500 - Internal Server Error
func NewInternalServerError(message string) *RestErr {
	return NewRestErr(message, "internal_server_error", http.StatusInternalServerError, nil)
}

// Implementamos a interface erro para devolver o RestErr
func (r *RestErr) Error() string {
	return r.Message
}