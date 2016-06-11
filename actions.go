package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kazukgw/coa"
)

type DefaultErrorHandler struct {
}

func (a *DefaultErrorHandler) HandleError(c coa.Context, err error) error {
	fmt.Println("[error] ", err.Error())
	return err
}

type GetCurrentDir struct {
	Dir string
}

func (a *GetCurrentDir) Do(c coa.Context) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	a.Dir = dir
	return nil
}

type GetFmPath struct {
	FmPath string
}

func (a *GetFmPath) Do(c coa.Context) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	a.FmPath = filepath.Join(dir, ".fm")

	_, err = os.Stat(a.FmPath)
	if err == os.ErrNotExist {
		if err := os.Mkdir(a.FmPath, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

type GetListKey struct {
	Key string
}

func (a *GetListKey) Do(c coa.Context) error {
	clictx := c.(*Ctx).Context
	args := clictx.Args()
	if len(args) < 1 {
		return nil
	}
	a.Key = args[0]
	return nil
}
