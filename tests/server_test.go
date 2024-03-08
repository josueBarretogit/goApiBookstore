package tests

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/server"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

type ApiFindAllResponse struct {
	Models []interface{} `json:"models"`
}

type ApiCreateResponse struct {
	Created interface{} `json:"created"`
}

type ApiUpdateResponse struct {
	Updated interface{} `json:"updated"`
}

func TestFindAllRouteRole(t *testing.T) {
	router := server.SetupServer()
	assert := assert.New(t)

	w := httptest.NewRecorder()
	req, errRequest := http.NewRequest("GET", os.Getenv("ROOT_API")+consts.RoleModelName+consts.RouteFindAll, nil)

	req.Header.Add("authorization", os.Getenv("JWT_KEY_TESTS"))

	assert.NoError(errRequest, "no error when requesting")

	router.ServeHTTP(w, req)

	response := w.Result()
	var body ApiFindAllResponse

	errBody := json.Unmarshal(w.Body.Bytes(), &body)

	assert.NoError(errBody, "should unmarshal response body")

	assert.Equal(http.StatusOK, response.StatusCode, "status code should be 200")
	assert.Containsf(w.Body.String(), "models", "response should contain at least models")
	assert.True(reflect.TypeOf(body.Models).Kind() == reflect.Slice, "response models should be slice")
}

func TestCreateRole(t *testing.T) {
	router := server.SetupServer()
	assert := assert.New(t)

	w := httptest.NewRecorder()

	roleMock := map[string]interface{}{"rolename": faker.Name()}

	roleMockJson, errMarshal := json.Marshal(roleMock)

	assert.NoError(errMarshal)

	req, errRequest := http.NewRequest(http.MethodPost, os.Getenv("ROOT_API")+consts.RoleModelName+consts.RouteCreate, bytes.NewReader(roleMockJson))

	req.Header.Add("authorization", os.Getenv("JWT_KEY_TESTS"))
	req.Header.Set("Content-Type", "application/json")

	assert.NoError(errRequest, "no error when requesting")

	router.ServeHTTP(w, req)

	response := w.Result()
	var body ApiCreateResponse

	errBody := json.Unmarshal(w.Body.Bytes(), &body)

	assert.NoError(errBody, "should unmarshal response body")

	assert.Equal(http.StatusOK, response.StatusCode, "status code should be 200")
	assert.Containsf(w.Body.String(), "created", "response should contain created role")
}

func TestUpdateRole(t *testing.T) {
	router := server.SetupServer()
	assert := assert.New(t)

	w := httptest.NewRecorder()

	roleMock := map[string]interface{}{"rolename": faker.Name()}

	roleMockJson, errMarshal := json.Marshal(roleMock)

	assert.NoError(errMarshal)

	req, errRequest := http.NewRequest(http.MethodPut, os.Getenv("ROOT_API")+consts.RoleModelName+"/update/3", bytes.NewReader(roleMockJson))

	req.Header.Add("authorization", os.Getenv("JWT_KEY_TESTS"))
	req.Header.Set("Content-Type", "application/json")

	assert.NoError(errRequest, "no error when requesting")

	router.ServeHTTP(w, req)

	response := w.Result()
	var body ApiCreateResponse

	errBody := json.Unmarshal(w.Body.Bytes(), &body)

	assert.NoError(errBody, "should unmarshal response body")

	assert.Equal(http.StatusOK, response.StatusCode, "status code should be 200")
	assert.Containsf(w.Body.String(), "updated", "response should contain updated role")
}
