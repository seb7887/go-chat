package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/challenge/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidator(t *testing.T) {
	setupTesting()

	req := httptest.NewRequest("GET", "/", nil)
	login, err := helpers.GenerateJwt(1, "john")
	require.NoError(t, err)

	token := fmt.Sprintf("Bearer %s", login.Token)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	h := ValidateUser(testEndpoint)

	h.ServeHTTP(rec, req)

	res := rec.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestValidatorFail(t *testing.T) {
	setupTesting()

	req := httptest.NewRequest("GET", "/", nil)

	token := "Bearer huiohasxcbui"
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()

	h := ValidateUser(testEndpoint)

	h.ServeHTTP(rec, req)

	res := rec.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

// Change working directory to project root
func setupTesting() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func testEndpoint(w http.ResponseWriter, r *http.Request) {}
