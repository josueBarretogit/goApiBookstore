package tests

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/server"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestFindAllRoute(t *testing.T) {

	router := server.SetupServer()


	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", os.Getenv("ROOT_API") + consts.RoleModelName + consts.RouteFindAll, nil)
	req.Header.Add("authorization", os.Getenv("JWT_KEY_TESTS"))

	fmt.Println(req.Header.Get("authorization"))

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
