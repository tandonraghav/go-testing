package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tandonraghav/go-testing/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGoogle(t *testing.T){
	req,_:=http.NewRequest("GET", "/", nil)

	rr := httptest.NewRecorder()

	server.Routers().ServeHTTP(rr,req)
	fmt.Print(rr.Code)

	assert.Equal(t, 200, rr.Code, "Expecting `200`")

}