package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"satya.com/http_json_rest/messages"
	"testing"
)

/*
func TestQueryHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/query", nil)
	w := httptest.NewRecorder()
	queryHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}
	log.Printf("TestUpperCaseHandler : Done")
} */

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	homeHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	decoder := json.NewDecoder(res.Body)
	var str messages.JsonResponse
	decoder.Decode(&t)
	log.Println("Got response string : ", str.JsonResponseString)

	// if string(data) != "ABC" {
	// 	t.Errorf("expected ABC got %v", string(data))
	// }
	log.Printf("TestHomeHandler : Done")
}
