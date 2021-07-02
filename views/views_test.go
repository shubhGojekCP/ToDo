package views

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask(t *testing.T) {

	var mockedTask = []byte(`{"Id":1,"Task":"Running","Status":true}`)
	req, err := http.NewRequest("POST", "/api/task", bytes.NewBuffer(mockedTask))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := `{"Status":"Creates Successfully","StatusCode":201,"Proto":"","ProtoMajor":0,"ProtoMinor":0,"Header":null,"Body":null,"ContentLength":0,"TransferEncoding":null,"Close":false,"Uncompressed":false,"Trailer":null,"Request":null,"TLS":null}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body, expected)
	}

}
