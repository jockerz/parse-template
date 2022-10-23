package parser

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ParseTemplate(root string, ext string, t *template.Template) (*template.Template, error) {
	base_path := fmt.Sprintf("%s%c", root, os.PathSeparator)

	e := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, fmt.Sprintf(".%s", ext)) {
			// Template name
			name := strings.Replace(path, base_path, "", 1)

			// Template content
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			s := string(b)

			var tmpl *template.Template
			if t == nil {
				t = template.New(name)
				tmpl = t
			} else {
				tmpl = t.New(name)
			}

			_, err = tmpl.Parse(s)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if e != nil {
		return nil, e
	} else {
		return t, nil
	}
}
