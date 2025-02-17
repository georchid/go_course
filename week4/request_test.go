package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type TestCase struct {
	ID         string
	Response   string
	StatusCode int
}
