package templatex

import (
	"bytes"
	"github.com/civet148/log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	defaultTemplateName = "default_template_name"
)

func Generate(htmlTemplate string, data any, toFiles ...string) (htmlContent string, err error) {
	var t *template.Template

	if !isBuiltinTemplate(htmlTemplate) {
		t = template.Must(template.New(filepath.Base(htmlTemplate)).ParseFiles(htmlTemplate))
	} else {
		t = template.Must(template.New(defaultTemplateName).Parse(htmlTemplate))
	}

	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, data)
	if err != nil {
		return htmlContent, err
	}
	htmlContent = buffer.String()
	for _, to := range toFiles {
		var fp *os.File
		fp, err = createFile(to)
		if err != nil {
			return htmlContent, err
		}
		_, err = fp.WriteString(htmlContent)
		if err != nil {
			return htmlContent, err
		}
	}
	return htmlContent, nil
}

func isBuiltinTemplate(str string) bool {
	//<!DOCTYPE html>
	if strings.Contains(str, "<!") && strings.Contains(str, "DOCTYPE") && strings.Contains(str, "html") && strings.Contains(str, ">") {
		return true
	}
	if strings.Contains(str, "{{") && strings.Contains(str, "}}") {
		return true
	}
	return false
}

func createFile(strFilePath string) (f *os.File, err error) {
	dir := extractDir(strFilePath)
	if dir != "" {
		err = createDirIfNotExist(dir)
		if err != nil {
			log.Errorf(err)
			return nil, err
		}
	}
	return os.OpenFile(strFilePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
}

// extractDir从给定的完整文件路径中提取出目录部分
func extractDir(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			return path[:i]
		}
	}
	return ""
}

// createDirIfNotExist检查目录是否存在，如果不存在则创建
func createDirIfNotExist(dir string) error {
	var ignoreDirs = []string{
		"",
		".",
		"..",
		"./",
		"../",
	}
	dir = strings.TrimSpace(dir)
	for _, d := range ignoreDirs {
		if d == dir {
			return nil
		}
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 如果目录不存在，则创建目录，权限设置为0755（可根据需求调整）
		return os.MkdirAll(dir, os.ModePerm)
	}
	return err
}

