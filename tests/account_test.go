package tests

import (
	"api/bookstoreApi/consts"
	usermodels "api/bookstoreApi/models/userModels"
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



func TestFindAllRouteAccount(t *testing.T) {
	router := server.SetupServer()
	assert := assert.New(t)

	w := httptest.NewRecorder()
	req, errRequest := http.NewRequest(http.MethodGet, os.Getenv("ROOT_API")+consts.AccountModelName+consts.RouteFindAll, nil)

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

func TestCreateAccount(t *testing.T) {
	router := server.SetupServer()
	assert := assert.New(t)

	w := httptest.NewRecorder()

	accountMock := usermodels.Account{
		Username:  faker.Name(),
		Password:  faker.Password(),
	}

	accountMockJson, errMarshal := json.Marshal(accountMock)

	assert.NoError(errMarshal)

	req, errRequest := http.NewRequest(http.MethodPost, os.Getenv("ROOT_API")+consts.AccountModelName+consts.RouteCreate, bytes.NewReader(accountMockJson))

	req.Header.Add("authorization", os.Getenv("JWT_KEY_TESTS"))
	req.Header.Set("Content-Type", "application/json")

	assert.NoError(errRequest, "no error when requesting")

	router.ServeHTTP(w, req)

	response := w.Result()
	var body ApiCreateResponse

	errBody := json.Unmarshal(w.Body.Bytes(), &body)

	assert.NoError(errBody, "should unmarshal response body")

	assert.Equal(http.StatusOK, response.StatusCode, "status code should be 200")
	assert.Containsf(w.Body.String(), "created", "response should contain created account")
}

func TestUpdateAccount(t *testing.T) {
	router := server.SetupServer()
	assert := assert.New(t)

	w := httptest.NewRecorder()

	accountMock := usermodels.Account{
		Username:  faker.Name(),
		Password:  faker.Password(),
	}


	accountMockJson, errMarshal := json.Marshal(accountMock)

	assert.NoError(errMarshal)

	req, errRequest := http.NewRequest(http.MethodPut, os.Getenv("ROOT_API")+consts.AccountModelName+"/update/1", bytes.NewReader(accountMockJson))

	req.Header.Add("authorization", os.Getenv("JWT_KEY_TESTS"))
	req.Header.Set("Content-Type", "application/json")

	assert.NoError(errRequest, "no error when requesting")

	router.ServeHTTP(w, req)

	response := w.Result()
	var body ApiUpdateResponse

	errBody := json.Unmarshal(w.Body.Bytes(), &body)

	assert.NoError(errBody, "should unmarshal response body")

	assert.Equal(http.StatusOK, response.StatusCode, "status code should be 200")
	assert.Containsf(w.Body.String(), "updated", "response should contain updated role")
}
