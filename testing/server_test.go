package testing

import (
	"github.com/tandonraghav/go-testing/server"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogle(t *testing.T){
	req,_:=http.NewRequest("GET", "/", nil)

	rr := httptest.NewRecorder()

	server.Routers().ServeHTTP(rr,req)

	assert.Equal(t, 200, rr.Code, "Expecting `200`")

}