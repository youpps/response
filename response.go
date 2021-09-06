package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type J = map[string]interface{}

type responseWithStatus struct {
	rw         http.ResponseWriter
	statusCode int
}

type ResponseWithStatus interface {
	JSON(J)
	String(string)
	Bytes([]byte)
	HTML(string)
}

func Status(rw http.ResponseWriter, statusCode int) ResponseWithStatus {
	return &responseWithStatus{rw, statusCode}
}

func (r *responseWithStatus) JSON(jsonMap J) {
	r.rw.Header().Add("Content-Type", "application/json; charset=UTF-8")
	bytes, err := jsonParse(jsonMap)
	if err != nil {
		r.rw.Write([]byte("No valid json."))
		return
	}
	r.rw.WriteHeader(r.statusCode)
	r.rw.Write(bytes)
}

func (r *responseWithStatus) Bytes(bytes []byte) {
	r.rw.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	r.rw.WriteHeader(r.statusCode)
	r.rw.Write(bytes)
}

func (r *responseWithStatus) String(str string) {
	r.rw.Header().Add("Content-Type", "text/plain; charset=UTF-8")
	r.rw.WriteHeader(r.statusCode)
	r.rw.Write([]byte(str))
}

func (r *responseWithStatus) HTML(html string) {
	r.rw.Header().Add("Content-Type", "text/html; charset=UTF-8")
	r.rw.WriteHeader(r.statusCode)
	r.rw.Write([]byte(html))
}

func Cookie(rw http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(rw, cookie)
}

func jsonParse(jsonMap J) ([]byte, error) {
	bytes, err := json.Marshal(jsonMap)
	if err != nil {
		return nil, errors.New("No valid json: " + err.Error())
	}
	return bytes, nil
}
