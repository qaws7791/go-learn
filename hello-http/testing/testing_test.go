package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil) // 요청 URL은 /로 한다.

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)       // 상태 코드가 200인지 확인한다.
	data, _ := io.ReadAll(res.Body)             // 응답 바디를 읽어온다.
	assert.Equal("Hello, world!", string(data)) // 응답 바디가 "Hello, world!"인지 확인한다.
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/bar", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello, bar!", string(data))
}
