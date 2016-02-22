package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"os"
)

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	handler(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestHandleIndexReturnsWithHostname(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	h, _ := os.Hostname()
	expectedBody := fmt.Sprintf("Hi there, I'm served from %s!", h)
	handler(response, request)
	assert.Equal(t, expectedBody, response.Body.String())
}
