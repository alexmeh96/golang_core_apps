package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleGetFoo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FOO"))
}

// тестирование с использованиме сервера предоставляемого httptest
func TestHandleByServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expected := "FOO"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}

// тестирование с использованиме ResponseRecorder предоставляемого httptest,
// который передаётся в тестируемый хэндлер, для записи в него результата выполнения этого хэндлера
func TestHandleByRecorder(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}

	handleGetFoo(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", rr.Result().StatusCode)
	}
	defer rr.Result().Body.Close()

	expected := "FOO"
	b, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}
