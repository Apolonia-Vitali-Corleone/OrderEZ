package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"user-service/internal/model/po"
	"user-service/internal/service"
	"user-service/util"
)

type testServer struct {
	router      *gin.Engine
	db          *gorm.DB
	redisClient *redis.Client
	cleanup     func()
}

func setupTestServer(t *testing.T) *testServer {
	t.Helper()

	gin.SetMode(gin.TestMode)

	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite database: %v", err)
	}

	if err := db.AutoMigrate(&po.User{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("failed to start miniredis: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{Addr: mr.Addr()})

	userService := service.NewUserService(db, redisClient, nil)

	snowflake, err := util.NewSnowflake(1, 1)
	if err != nil {
		t.Fatalf("failed to create snowflake generator: %v", err)
	}

	userHandler := NewUserHandler(userService, snowflake)

	router := gin.New()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
		userGroup.POST("/logout", userHandler.Logout)
		userGroup.GET("/", userHandler.GetAllUsers)
	}

	cleanup := func() {
		redisClient.Close()
		mr.Close()
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}

	return &testServer{
		router:      router,
		db:          db,
		redisClient: redisClient,
		cleanup:     cleanup,
	}
}

func performJSONRequest(t *testing.T, router http.Handler, method, target string, body interface{}) *httptest.ResponseRecorder {
	t.Helper()

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("failed to encode request body: %v", err)
		}
	}

	req := httptest.NewRequest(method, target, &buf)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}

func uniqueCredentials(prefix string) (string, string) {
	return fmt.Sprintf("%s_%d", prefix, time.Now().UnixNano()), "Password!234"
}

func decodeResponse(t *testing.T, resp *httptest.ResponseRecorder, v interface{}) {
	t.Helper()
	if err := json.Unmarshal(resp.Body.Bytes(), v); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
}

func TestRegisterEndpoint(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.cleanup()

	username, password := uniqueCredentials("register_user")
	resp := performJSONRequest(t, ts.router, http.MethodPost, "/user/register", map[string]string{
		"username": username,
		"password": password,
	})

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", resp.Code, resp.Body.String())
	}

	var body struct {
		Token string `json:"token"`
	}
	decodeResponse(t, resp, &body)

	if body.Token == "" {
		t.Fatalf("expected token in register response")
	}

	var storedUser po.User
	if err := ts.db.Where("username = ?", username).First(&storedUser).Error; err != nil {
		t.Fatalf("expected user to be stored, got error: %v", err)
	}
}

func TestLoginEndpoint(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.cleanup()

	username, password := uniqueCredentials("login_user")

	registerResp := performJSONRequest(t, ts.router, http.MethodPost, "/user/register", map[string]string{
		"username": username,
		"password": password,
	})
	if registerResp.Code != http.StatusOK {
		t.Fatalf("register failed with status %d: %s", registerResp.Code, registerResp.Body.String())
	}

	loginResp := performJSONRequest(t, ts.router, http.MethodPost, "/user/login", map[string]string{
		"username": username,
		"password": password,
	})

	if loginResp.Code != http.StatusOK {
		t.Fatalf("expected login status 200, got %d: %s", loginResp.Code, loginResp.Body.String())
	}

	var body struct {
		Token string `json:"token"`
	}
	decodeResponse(t, loginResp, &body)

	if body.Token == "" {
		t.Fatalf("expected token in login response")
	}
}

func TestLogoutEndpoint(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.cleanup()

	username, password := uniqueCredentials("logout_user")

	registerResp := performJSONRequest(t, ts.router, http.MethodPost, "/user/register", map[string]string{
		"username": username,
		"password": password,
	})
	if registerResp.Code != http.StatusOK {
		t.Fatalf("register failed with status %d: %s", registerResp.Code, registerResp.Body.String())
	}

	loginResp := performJSONRequest(t, ts.router, http.MethodPost, "/user/login", map[string]string{
		"username": username,
		"password": password,
	})
	if loginResp.Code != http.StatusOK {
		t.Fatalf("login failed with status %d: %s", loginResp.Code, loginResp.Body.String())
	}

	var loginBody struct {
		Token string `json:"token"`
	}
	decodeResponse(t, loginResp, &loginBody)
	if loginBody.Token == "" {
		t.Fatalf("expected login token")
	}

	req := httptest.NewRequest(http.MethodPost, "/user/logout", nil)
	req.Header.Set("Authorization", loginBody.Token)
	resp := httptest.NewRecorder()
	ts.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected logout status 200, got %d: %s", resp.Code, resp.Body.String())
	}

	var body struct {
		Message string `json:"message"`
	}
	decodeResponse(t, resp, &body)
	if body.Message != "登出成功" {
		t.Fatalf("unexpected logout message: %q", body.Message)
	}
}

func TestGetAllUsersEndpoint(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.cleanup()

	username, password := uniqueCredentials("list_user")

	registerResp := performJSONRequest(t, ts.router, http.MethodPost, "/user/register", map[string]string{
		"username": username,
		"password": password,
	})
	if registerResp.Code != http.StatusOK {
		t.Fatalf("register failed with status %d: %s", registerResp.Code, registerResp.Body.String())
	}

	loginResp := performJSONRequest(t, ts.router, http.MethodPost, "/user/login", map[string]string{
		"username": username,
		"password": password,
	})
	if loginResp.Code != http.StatusOK {
		t.Fatalf("login failed with status %d: %s", loginResp.Code, loginResp.Body.String())
	}

	req := httptest.NewRequest(http.MethodGet, "/user/?page=1&pageSize=5", nil)
	resp := httptest.NewRecorder()
	ts.router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", resp.Code, resp.Body.String())
	}

	var body struct {
		Users []po.User `json:"users"`
	}
	decodeResponse(t, resp, &body)

	found := false
	for _, u := range body.Users {
		if u.Username == username {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected to find user %q in list", username)
	}
}
