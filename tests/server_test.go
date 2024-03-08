package tests

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/server"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)



type ApiResponse struct {
	Models []interface{} `json:"models"`
}

func TestFindAllRouteRole(t *testing.T) {

	router := server.SetupServer()
  assert := assert.New(t)

	w := httptest.NewRecorder()
	req, errRequest := http.NewRequest("GET", os.Getenv("ROOT_API") + consts.RoleModelName + consts.RouteFindAll, nil)

	req.Header.Add("authorization", os.Getenv("JWT_KEY_TESTS"))

	assert.NoError(errRequest, "no error when requesting")

	router.ServeHTTP(w, req)

	response := w.Result()
	var body ApiResponse

	errBody := json.Unmarshal(w.Body.Bytes(), &body)

	assert.NoError(errBody , "should unmarshal response body")

	assert.Equal(http.StatusOK, response.StatusCode, "status code should be 200")
	assert.Containsf( w.Body.String(), "models", "response should contain at least models")
	assert.True(reflect.TypeOf(body.Models).Kind() == reflect.Slice, "response models should be slice")
}
