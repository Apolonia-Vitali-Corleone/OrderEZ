package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"user-service/internal/model/po"
	"user-service/internal/service"
	"user-service/util"
)

type testContext struct {
	router      *gin.Engine
	redisServer *miniredis.Miniredis
	redisClient *redis.Client
	db          *gorm.DB
}

func setupTestContext(t *testing.T) *testContext {
	t.Helper()

	gin.SetMode(gin.TestMode)

	mr := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	if err := db.AutoMigrate(&po.User{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	userService := service.NewUserService(db, redisClient, nil)
	idGen, err := util.NewSnowflake(1, 1)
	if err != nil {
		t.Fatalf("failed to create snowflake: %v", err)
	}
	handler := NewUserHandler(userService, idGen)

	router := gin.New()
	router.Use(gin.Recovery())

	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.POST("/logout", handler.Logout)
		userGroup.GET("/", handler.GetAllUsers)
	}

	return &testContext{
		router:      router,
		redisServer: mr,
		redisClient: redisClient,
		db:          db,
	}
}

func (tc *testContext) cleanup() {
	if tc.redisClient != nil {
		_ = tc.redisClient.Close()
	}
	if tc.redisServer != nil {
		tc.redisServer.Close()
	}
	if tc.db != nil {
		if sqlDB, err := tc.db.DB(); err == nil {
			_ = sqlDB.Close()
		}
	}
}

func uniqueUsername(prefix string) string {
	return fmt.Sprintf("%s_%d", prefix, time.Now().UnixNano())
}

func TestRegisterEndpoint(t *testing.T) {
	tc := setupTestContext(t)
	defer tc.cleanup()

	payload := map[string]string{
		"username": uniqueUsername("register"),
		"password": "secret-password",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	tc.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	var responseBody map[string]string
	if err := json.Unmarshal(resp.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if responseBody["token"] == "" {
		t.Fatalf("expected token in response, got: %v", responseBody)
	}
}

func TestLoginEndpoint(t *testing.T) {
	tc := setupTestContext(t)
	defer tc.cleanup()

	username := uniqueUsername("login")
	registerPayload := map[string]string{
		"username": username,
		"password": "login-password",
	}
	registerBody, err := json.Marshal(registerPayload)
	if err != nil {
		t.Fatalf("failed to marshal register payload: %v", err)
	}
	registerReq := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(registerBody))
	registerReq.Header.Set("Content-Type", "application/json")
	registerResp := httptest.NewRecorder()
	tc.router.ServeHTTP(registerResp, registerReq)
	if registerResp.Code != http.StatusOK {
		t.Fatalf("register failed with status %d", registerResp.Code)
	}

	loginPayload := map[string]string{
		"username": username,
		"password": "login-password",
	}
	loginBody, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("failed to marshal login payload: %v", err)
	}

	loginReq := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginResp := httptest.NewRecorder()

	tc.router.ServeHTTP(loginResp, loginReq)

	if loginResp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, loginResp.Code)
	}

	var responseBody map[string]string
	if err := json.Unmarshal(loginResp.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal login response: %v", err)
	}
	if responseBody["token"] == "" {
		t.Fatalf("expected token in login response, got: %v", responseBody)
	}
}

func TestGetAllUsersEndpoint(t *testing.T) {
	tc := setupTestContext(t)
	defer tc.cleanup()

	username := uniqueUsername("list")
	registerPayload := map[string]string{
		"username": username,
		"password": "list-password",
	}
	registerBody, err := json.Marshal(registerPayload)
	if err != nil {
		t.Fatalf("failed to marshal register payload: %v", err)
	}
	registerReq := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(registerBody))
	registerReq.Header.Set("Content-Type", "application/json")
	registerResp := httptest.NewRecorder()
	tc.router.ServeHTTP(registerResp, registerReq)
	if registerResp.Code != http.StatusOK {
		t.Fatalf("register failed with status %d", registerResp.Code)
	}

	loginPayload := map[string]string{
		"username": username,
		"password": "list-password",
	}
	loginBody, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("failed to marshal login payload: %v", err)
	}
	loginReq := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginResp := httptest.NewRecorder()
	tc.router.ServeHTTP(loginResp, loginReq)
	if loginResp.Code != http.StatusOK {
		t.Fatalf("login failed with status %d", loginResp.Code)
	}
	var loginBodyResp map[string]string
	if err := json.Unmarshal(loginResp.Body.Bytes(), &loginBodyResp); err != nil {
		t.Fatalf("failed to unmarshal login response: %v", err)
	}
	token := loginBodyResp["token"]
	if token == "" {
		t.Fatalf("expected token after login, got: %v", loginBodyResp)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/user/?page=1&pageSize=10", nil)
	listReq.Header.Set("Authorization", "Bearer "+token)
	listResp := httptest.NewRecorder()
	tc.router.ServeHTTP(listResp, listReq)

	if listResp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, listResp.Code)
	}

	var responseBody struct {
		Users []po.User `json:"users"`
	}
	if err := json.Unmarshal(listResp.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal list response: %v", err)
	}
	if len(responseBody.Users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(responseBody.Users))
	}
	if responseBody.Users[0].Username != username {
		t.Fatalf("expected username %s, got %s", username, responseBody.Users[0].Username)
	}
}

func TestLogoutEndpoint(t *testing.T) {
	tc := setupTestContext(t)
	defer tc.cleanup()

	username := uniqueUsername("logout")
	registerPayload := map[string]string{
		"username": username,
		"password": "logout-password",
	}
	registerBody, err := json.Marshal(registerPayload)
	if err != nil {
		t.Fatalf("failed to marshal register payload: %v", err)
	}
	registerReq := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewReader(registerBody))
	registerReq.Header.Set("Content-Type", "application/json")
	registerResp := httptest.NewRecorder()
	tc.router.ServeHTTP(registerResp, registerReq)
	if registerResp.Code != http.StatusOK {
		t.Fatalf("register failed with status %d", registerResp.Code)
	}

	loginPayload := map[string]string{
		"username": username,
		"password": "logout-password",
	}
	loginBody, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("failed to marshal login payload: %v", err)
	}
	loginReq := httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader(loginBody))
	loginReq.Header.Set("Content-Type", "application/json")
	loginResp := httptest.NewRecorder()
	tc.router.ServeHTTP(loginResp, loginReq)
	if loginResp.Code != http.StatusOK {
		t.Fatalf("login failed with status %d", loginResp.Code)
	}
	var loginBodyResp map[string]string
	if err := json.Unmarshal(loginResp.Body.Bytes(), &loginBodyResp); err != nil {
		t.Fatalf("failed to unmarshal login response: %v", err)
	}
	token := loginBodyResp["token"]
	if token == "" {
		t.Fatalf("expected token after login, got: %v", loginBodyResp)
	}

	logoutReq := httptest.NewRequest(http.MethodPost, "/user/logout", nil)
	logoutReq.Header.Set("Authorization", "Bearer "+token)
	logoutResp := httptest.NewRecorder()
	tc.router.ServeHTTP(logoutResp, logoutReq)

	if logoutResp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, logoutResp.Code)
	}

	var responseBody map[string]string
	if err := json.Unmarshal(logoutResp.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal logout response: %v", err)
	}
	if responseBody["message"] != "登出成功" {
		t.Fatalf("unexpected logout response: %v", responseBody)
	}
}
