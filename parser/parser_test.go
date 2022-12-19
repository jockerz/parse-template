package parser_test

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jockerz/parse-template/parser"
)

type Template struct {
	templates *template.Template
}

func prepareTestEnvirontment(ext string) (string, error) {
	// Prepare file and directory for test
	baseDir := filepath.Join("..", "test_template")
	// if _, err := os.Stat(baseDir); !os.IsNotExist(err) {
	// 	return "", errors.New("Test directory is existed")
	// }
	os.Mkdir(baseDir, 0755)

	// test_template/root.html
	filename0 := filepath.Join(baseDir, fmt.Sprintf("root.%s", ext))
	file0, err := os.Create(filename0)
	if err != nil {
		return "", err
	}
	defer file0.Close()
	file0.WriteString(fmt.Sprintf("Content of root.html\n{{ template \"parts/part.%s\" }}", ext))

	level1 := filepath.Join(baseDir, "level1")
	os.Mkdir(level1, 0755)

	// test_template/level1/level1.html
	filename1 := filepath.Join(level1, "level1.html")
	file1, err := os.Create(filename1)
	if err != nil {
		return "", err
	}
	defer file1.Close()
	file1.WriteString(fmt.Sprintf("Content of level1.html\n{{ template \"parts/%s.html\" }}", ext))

	level2 := filepath.Join(level1, "level2")
	os.Mkdir(level2, 0755)

	// test_template/level1/level2/level2.html
	filename2 := filepath.Join(level2, fmt.Sprintf("level2.%s", ext))
	file2, err := os.Create(filename2)
	if err != nil {
		return "", err
	}
	defer file2.Close()
	file2.WriteString(fmt.Sprintf("Content of level2.html\n{{ template \"parts/part.%s\" }}", ext))

	parts := filepath.Join(baseDir, "parts")
	os.Mkdir(parts, 0755)

	// test_template/parts/part.html
	filename3 := filepath.Join(parts, fmt.Sprintf("part.%s", ext))
	file3, err := os.Create(filename3)
	if err != nil {
		return "", err
	}
	defer file3.Close()
	file3.WriteString(fmt.Sprintf("Content of part.%s", ext))

	return baseDir, nil
}

func TestParseTemplate(t *testing.T) {
	// test that template file is valid and process as expected

	baseDir, err := prepareTestEnvirontment("html")
	// Remove template test directories
	defer os.RemoveAll(baseDir)
	if err != nil {
		t.Error(err)
	}

	// Test
	tmpl, err := parser.ParseTemplate("../test_template", "html", nil)
	if err != nil {
		panic(err)
	}
	tt := &Template{
		templates: tmpl,
	}

	var out bytes.Buffer

	tt.templates.ExecuteTemplate(&out, "root.html", nil)
	if !strings.Contains(out.String(), "part.html") {
		t.Errorf("part.html not included in root.html: %s", out.String())
	}

	tt.templates.ExecuteTemplate(&out, "level1/level1.html", nil)
	if !strings.Contains(out.String(), "part.html") {
		t.Errorf("part.html not included in level1/level1.html: %s", out.String())
	}

	tt.templates.ExecuteTemplate(&out, "level1/level2/level2.html", nil)
	if !strings.Contains(out.String(), "part.html") {
		t.Errorf("part.html not included in level1/level2/level2.html: %s", out.String())
	}

	tt.templates.ExecuteTemplate(&out, "parts/part.html", nil)
	if !strings.Contains(out.String(), "part.html") {
		t.Errorf("part.html not included in parts/part.html: %s", out.String())
	}
}
