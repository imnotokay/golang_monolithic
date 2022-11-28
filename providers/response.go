package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status         int         `json:"status"`
	Data           interface{} `json:"data"`
	Message        string      `json:"message"`
	contentType    string
	responseWriter http.ResponseWriter
}

func DefaultResponse(w http.ResponseWriter) Response {
	return Response{
		Status:         http.StatusOK,
		contentType:    "application/json",
		responseWriter: w,
	}
}

func createInternalServerError(r *Response, message string) {

	if message == "" {
		r.Message = "Internal server error"
	} else {
		r.Message = message
	}
	r.Data = nil
	r.Status = http.StatusInternalServerError
}

func createNotFound(r *Response, message string) {
	if message == "" {
		r.Message = "Resource not found"
	} else {
		r.Message = message
	}
	r.Data = nil
	r.Status = http.StatusNotFound
}

func createBadRequest(r *Response, message string) {
	if message == "" {
		r.Message = "Bad Request"
	} else {
		r.Message = message
	}
	r.Data = nil
	r.Status = http.StatusNotFound
}

func (r *Response) Send() {
	r.responseWriter.Header().Set("Content-Type", r.contentType)
	r.responseWriter.WriteHeader(r.Status)

	output, _ := json.Marshal(&r)
	fmt.Fprintln(r.responseWriter, string(output))
}

func (r *Response) InternalServerError(message string) {
	createInternalServerError(r, message)
	r.Send()
}

func (r *Response) NotFound(message string) {
	createNotFound(r, message)
	r.Send()
}

func (r *Response) BadRequest(message string) {
	createBadRequest(r, message)
	r.Send()
}

func CatchExceptions(resp Response) {
	if err := recover(); err != nil {
		resp.InternalServerError(fmt.Sprint(err))
	}
}
