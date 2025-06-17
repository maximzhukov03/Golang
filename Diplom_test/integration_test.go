// integration_test.go
// Интеграционные тесты End-to-End для демонстрационного приложения
package main

import (
    "bytes"
    "encoding/json"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"
    "testing"
    "github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8080"

func TestE2E_RegisterLoginUploadList(t *testing.T) {
    // Регистрация
    regBody := map[string]string{"email": "test@exa.com", "password": "Secret123"}
    regBts, _ := json.Marshal(regBody)
    resp, err := http.Post(baseURL+"/api/register", "application/json", bytes.NewReader(regBts))
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
    resp.Body.Close()
    // Логин
    loginBody := map[string]string{"email": "test@exa.com", "password": "Secret123"}
    loginBts, _ := json.Marshal(loginBody)
    resp, err = http.Post(baseURL+"/api/login", "application/json", bytes.NewReader(loginBts))
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    var loginResp struct { Token string `json:"token"` }
    json.NewDecoder(resp.Body).Decode(&loginResp)
    resp.Body.Close()
    assert.NotEmpty(t, loginResp.Token)

    // 3. Загрузка файла
    var buf bytes.Buffer
    writer := multipart.NewWriter(&buf)
    part, _ := writer.CreateFormFile("file", filepath.Base("icons8-golang-480.png"))
    f, _ := os.Open("icons8-golang-480.png")
    io.Copy(part, f)
    f.Close()
    writer.Close()

    req, _ := http.NewRequest("POST", baseURL+"/api/files", &buf)
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("Authorization", "Bearer "+loginResp.Token)
    resp, err = http.DefaultClient.Do(req)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
    resp.Body.Close()

    // 4. Получение списка файлов
    req, _ = http.NewRequest("GET", baseURL+"/api/files", nil)
    req.Header.Set("Authorization", "Bearer "+loginResp.Token)
    resp, err = http.DefaultClient.Do(req)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    var files []struct { URL string `json:"url"` }
    json.NewDecoder(resp.Body).Decode(&files)
    resp.Body.Close()
    assert.GreaterOrEqual(t, len(files), 1)
    assert.Contains(t, files[0].URL, "http")
}

func TestAdminOperations(t *testing.T) {
    // Предполагается, что существует админ с токеном
    adminToken := os.Getenv("ADMIN_TOKEN")
    if adminToken == "" {
        t.Skip("ADMIN_TOKEN not set")
    }

    // Promote пользователя с ID 1
    req, _ := http.NewRequest("POST", baseURL+"/api/admin/users/1/promote", nil)
    req.Header.Set("Authorization", "Bearer "+adminToken)
    resp, err := http.DefaultClient.Do(req)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    resp.Body.Close()

    // Получение списка всех пользователей
    req, _ = http.NewRequest("GET", baseURL+"/api/admin/users", nil)
    req.Header.Set("Authorization", "Bearer "+adminToken)
    resp, err = http.DefaultClient.Do(req)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    var users []map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&users)
    resp.Body.Close()
    assert.Greater(t, len(users), 0)
}
