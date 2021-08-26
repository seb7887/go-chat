package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/challenge/pkg/controller"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/services"
	"github.com/challenge/pkg/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var start uint

func TestHealcheck(t *testing.T) {
	setupTesting()

	h := buildHandler()

	req := httptest.NewRequest("POST", CheckEndpoint, nil)
	rec := httptest.NewRecorder()

	h.Check(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var respMsg models.Health
	err := json.NewDecoder(res.Body).Decode(&respMsg)
	require.NoError(t, err)
	assert.Equal(t, respMsg.Health, "ok")

	cleanup()
	seedDB()
}

func TestCreateUser(t *testing.T) {
	setupTesting()

	h := buildHandler()

	bodyJSON := []byte(`{
		"username": "test",
		"password": "test123"
		}`)

	req := httptest.NewRequest("POST", UsersEndpoint, bytes.NewBuffer(bodyJSON))
	rec := httptest.NewRecorder()

	h.CreateUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	var respMsg models.NewUserResp
	err := json.NewDecoder(res.Body).Decode(&respMsg)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestLogin(t *testing.T) {
	setupTesting()

	h := buildHandler()

	bodyJSON := []byte(`{
		"username": "john",
		"password": "john123"
		}`)

	req := httptest.NewRequest("POST", LoginEndpoint, bytes.NewBuffer(bodyJSON))
	rec := httptest.NewRecorder()

	h.Login(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var respMsg models.Login
	err := json.NewDecoder(res.Body).Decode(&respMsg)
	require.NoError(t, err)
	assert.Equal(t, uint(1), respMsg.Id)
}

func TestInvalidLogin(t *testing.T) {
	setupTesting()

	h := buildHandler()

	bodyJSON := []byte(`{
		"username": "test",
		"password": "wrongPassword"
		}`)

	req := httptest.NewRequest("POST", LoginEndpoint, bytes.NewBuffer(bodyJSON))
	rec := httptest.NewRecorder()

	h.Login(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestSendTextMessage(t *testing.T) {
	setupTesting()

	h := buildHandler()

	bodyJSON := []byte(`{
		"sender": 1,
		"recipient": 1,
		"content": {"type": "text", "text": "Hello World!"}
		}`)

	req := httptest.NewRequest("POST", MessagesEndpoint, bytes.NewBuffer(bodyJSON))
	ctx := context.WithValue(req.Context(), "user_id", float64(1))
	req = req.WithContext(ctx)
	rec := httptest.NewRecorder()

	h.SendMessage(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	var respMsg models.NewMsgResp
	err := json.NewDecoder(res.Body).Decode(&respMsg)
	require.NoError(t, err)
	start = respMsg.Id
}

func TestSendImageMessage(t *testing.T) {
	setupTesting()

	h := buildHandler()

	bodyJSON := []byte(`{
		"sender": 1,
		"recipient": 1,
		"content": {"type": "image", "url": "http://example.com", "height": 200, "width": 200}
		}`)

	req := httptest.NewRequest("POST", MessagesEndpoint, bytes.NewBuffer(bodyJSON))
	ctx := context.WithValue(req.Context(), "user_id", float64(1))
	req = req.WithContext(ctx)
	rec := httptest.NewRecorder()

	h.SendMessage(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestSendVideoMessage(t *testing.T) {
	setupTesting()

	h := buildHandler()

	bodyJSON := []byte(`{
		"sender": 1,
		"recipient": 1,
		"content": {"type": "video", "url": "http://example.com", "source": "youtube"}
		}`)

	req := httptest.NewRequest("POST", MessagesEndpoint, bytes.NewBuffer(bodyJSON))
	ctx := context.WithValue(req.Context(), "user_id", float64(1))
	req = req.WithContext(ctx)
	rec := httptest.NewRecorder()

	h.SendMessage(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetMessages(t *testing.T) {
	setupTesting()

	h := buildHandler()

	url := fmt.Sprintf("%s?recipient=1&start=%d&limit=10", MessagesEndpoint, start)

	req := httptest.NewRequest("GET", url, nil)
	ctx := context.WithValue(req.Context(), "user_id", float64(1))
	req = req.WithContext(ctx)
	rec := httptest.NewRecorder()

	h.GetMessages(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var respMsg []models.MessageResp
	err := json.NewDecoder(res.Body).Decode(&respMsg)
	require.NoError(t, err)
	assert.Equal(t, 3, len(respMsg))

	cleanup()
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

func buildHandler() controller.Handler {
	userRepository := storage.NewUserRepository()
	messageRepository := storage.NewMessageRepository()
	userService := services.NewUserService(userRepository)
	messageService := services.NewMessageService(messageRepository)
	return controller.NewHandler(userService, messageService)
}

func cleanup() {
	db := storage.GetInstance()

	db.Exec("DELETE FROM videos")
	db.Exec("DELETE FROM images")
	db.Exec("DELETE FROM texts")
	db.Exec("DELETE FROM messages")
	db.Exec("DELETE FROM users")
}

func seedDB() {
	db := storage.GetInstance()

	sampleUser := models.User{
		ID:       1,
		Username: "john",
		Password: "john123",
	}

	db.Create(&sampleUser)
}
