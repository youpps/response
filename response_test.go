package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	r := httptest.NewRecorder()

	Status(r, 200).String("Success!")

	if r.Result().StatusCode != 200 {
		t.Fatal("No equal status codes.")
	}
	if r.Body.String() != "Success!" {
		t.Fatal("No valid success response.")
	}

	r = httptest.NewRecorder()

	Status(r, 404).String("Error!")

	if r.Result().StatusCode != 404 {
		t.Fatal("No equal status codes.")
	}
	if r.Body.String() != "Error!" {
		t.Fatal("No valid error response.")
	}
}
func TestJSON(t *testing.T) {
	r := httptest.NewRecorder()

	Status(r, 200).JSON(J{
		"status": "success",
		"user": J{
			"name":     "youpps",
			"password": "qwerty",
		},
	})

	if r.Body.String() != `{"status":"success","user":{"name":"youpps","password":"qwerty"}}` {
		t.Fatal("No valid response body.")
	}
	if r.Header().Get("Content-type") != "application/json; charset=UTF-8" {
		t.Fatal("No valid response header.")
	}
}
func TestJsonParse(t *testing.T) {
	bytes, err := jsonParse(J{
		"name": "youpps",
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != `{"name":"youpps"}` {
		t.Fatal("No valid json result.")
	}
}

func TestString(t *testing.T) {
	r := httptest.NewRecorder()

	Status(r, 200).String("youpps")

	if r.Body.String() != "youpps" {
		t.Fatal("No valid string result.")
	}
}
func TestBytes(t *testing.T) {
	r := httptest.NewRecorder()

	Status(r, 200).Bytes([]byte("youpps"))

	if r.Body.String() != "youpps" {
		t.Fatal("No valid bytes-string result.")
	}
}
func TestCookie(t *testing.T) {
	r := httptest.NewRecorder()

	Cookie(r, &http.Cookie{
		Name:  "access-token",
		Value: "j3yhu1837rh21f9u1hfp912",
	})

	if r.Header().Get("Set-Cookie") != "access-token=j3yhu1837rh21f9u1hfp912" {
		t.Fatal("No valid set-cookie header.")
	}
}
