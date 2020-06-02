package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestStatusRoutes(t *testing.T) {
	testRouter := mux.NewRouter()
	serviceName := "%CUSTOM_PLUGIN_SERVICE_NAME%"
	serviceVersion := "0.0.0"
	StatusRoutes(testRouter, serviceName, serviceVersion)

	t.Run("/-/healthz - ok", func(assert *testing.T) {
		expectedResponse := fmt.Sprintf("{\"status\":\"OK\",\"name\":\"%s\",\"version\":\"%s\"}", serviceName, serviceVersion)
		responseRecorder := httptest.NewRecorder()
		request, error := http.NewRequest(http.MethodGet, "/-/healthz", nil)
		if error != nil {
			assert.Fatal("Error creating the /-/healthz request")
		}
		testRouter.ServeHTTP(responseRecorder, request)
		statusCode := responseRecorder.Result().StatusCode
		require.Equal(assert, http.StatusOK, statusCode, "The response statusCode should be 200")

		rawBody := responseRecorder.Result().Body
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(rawBody)
		require.Equal(assert, expectedResponse, buffer.String(), "The response body should be the expected one")
	})

	t.Run("/-/ready - ok", func(assert *testing.T) {
		expectedResponse := fmt.Sprintf("{\"status\":\"OK\",\"name\":\"%s\",\"version\":\"%s\"}", serviceName, serviceVersion)
		responseRecorder := httptest.NewRecorder()
		request, error := http.NewRequest(http.MethodGet, "/-/ready", nil)
		if error != nil {
			assert.Fatal("Error creating the /-/ready request")
		}
		testRouter.ServeHTTP(responseRecorder, request)
		statusCode := responseRecorder.Result().StatusCode
		require.Equal(assert, http.StatusOK, statusCode, "The response statusCode should be 200")

		rawBody := responseRecorder.Result().Body
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(rawBody)
		require.Equal(assert, expectedResponse, buffer.String(), "The response body should be the expected one")
	})

	t.Run("/-/check-up - ok", func(assert *testing.T) {
		expectedResponse := fmt.Sprintf("{\"status\":\"OK\",\"name\":\"%s\",\"version\":\"%s\"}", serviceName, serviceVersion)
		responseRecorder := httptest.NewRecorder()
		request, error := http.NewRequest(http.MethodGet, "/-/check-up", nil)
		if error != nil {
			assert.Fatal("Error creating the /-/check-up request")
		}
		testRouter.ServeHTTP(responseRecorder, request)
		statusCode := responseRecorder.Result().StatusCode
		require.Equal(assert, http.StatusOK, statusCode, "The response statusCode should be 200")

		rawBody := responseRecorder.Result().Body
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(rawBody)
		require.Equal(assert, expectedResponse, buffer.String(), "The response body should be the expected one")
	})
}
