package handler

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	gin.DefaultWriter = io.Discard
	r := NewRouter()
	file, err := os.Open("chairs.ifc")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// フォームデータを作成
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// ファイルを追加
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(part, file)

	// フォームの終わりを書き込む
	writer.Close()

	// リクエストを作成
	req := httptest.NewRequest("POST", "/exec", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
}
