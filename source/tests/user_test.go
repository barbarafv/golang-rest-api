package tests

import (
	_ "app/source/_testinit/fixture"
	"app/source/controllers/requests"
	"net/http"
	"testing"
	"time"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	runTest(func() {
		rest := &gofight.RequestConfig{Debug: true}

		rest.POST("/users").SetJSONInterface(requests.UserRequest{
			Login:     "login",
			Email:     "login@email.com",
			Password:  "login1234",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}).Run(router, func(response gofight.HTTPResponse, request gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, response.Code)
		})
	})
}
