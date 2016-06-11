package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/kazukgw/coa"
	"github.com/spf13/hugo/parser"
)

type buildIndexes struct {
	GetCurrentDir
	GetFmPath
	coa.DoSelf
	DefaultErrorHandler
}

type metadataSet struct {
	Path  string
	Value interface{}
}

func (ag *buildIndexes) Do(c coa.Context) error {
	metas := map[string][]metadataSet{}
	err := filepath.Walk(ag.GetCurrentDir.Dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if strings.HasSuffix(path, ".git") {
				return filepath.SkipDir
			}
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return nil
		}
		defer f.Close()
		p, err := parser.ReadFrom(f)
		if err != nil {
			HasNoHugoFrontMatter := "unable to read frontmatter at filepos"
			if strings.HasPrefix(err.Error(), HasNoHugoFrontMatter) {
				return nil
			}
			return err
		}

		m, err := p.Metadata()
		if err != nil {
			return nil
		}
		if metamap, ok := m.(map[string]interface{}); ok {
			for k, v := range metamap {
				ms, ok := metas[k]
				if !ok {
					ms = []metadataSet{}
				}
				ms = append(ms, metadataSet{path, v})
				metas[k] = ms
			}
		} else {
			reflect.TypeOf(m)
		}
		return nil
	})
	if err != nil {
		panic(err.Error())
	}

	for k, meta := range metas {
		dat := map[string]interface{}{}
		for _, v := range meta {
			dat[v.Path] = v.Value
		}
		fpath := filepath.Join(ag.FmPath, k+".json")
		f, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()
		if err := json.NewEncoder(f).Encode(dat); err != nil {
			panic(err.Error())
		}
	}

	return nil
}
