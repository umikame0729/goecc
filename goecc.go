package oc

import (
	"embed"

	"github.com/liuzl/gocc"
	"github.com/umikame0729/goefs"
)

//go:embed config
var config embed.FS

//go:embed dictionary
var dictionary embed.FS

var OpenCC *gocc.OpenCC = nil

func New(conv string, option ...gocc.Option) (*gocc.OpenCC, error) {
	if err := goefs.CreateDirFromEmbedFS(&config); err != nil {
		return nil, err
	}

	if err := goefs.CreateDirFromEmbedFS(&dictionary); err != nil {
		return nil, err
	}

	return gocc.New(conv, option...)
}
