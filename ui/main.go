package ui

import (
	"github.com/GeertJohan/go.rice"
)

func UI() *rice.Box {
	return rice.MustFindBox("_build")
}
