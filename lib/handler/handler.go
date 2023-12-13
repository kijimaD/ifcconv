package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

const reqFileForm = "file"
const ifcCmd = "IfcConvert"

func ExecHandler(c *gin.Context) {
	// リクエストのときのフォームに対応する名前
	file, err := c.FormFile(reqFileForm)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}

	// 一時ディレクトリにファイルを保存する
	tempFile, _ := os.CreateTemp(os.TempDir(), "tempfile-*.ifc")
	c.SaveUploadedFile(file, tempFile.Name())
	defer os.Remove(tempFile.Name())

	var out bytes.Buffer
	var stderr bytes.Buffer
	distPath := fmt.Sprintf("%s.obj", tempFile.Name()) // 変換先
	// ここの拡張子でどの形式に変換するか判断しているようだ
	task := fmt.Sprintf("%s %s %s -y", ifcCmd, tempFile.Name(), distPath)
	cmd := exec.Command("bash", "-c", task)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	fmt.Printf("実行された: %s\n", task)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("%s", err))
		return
	}

	fmt.Printf("結果: %s\n", out.String())

	c.File(distPath)
}
