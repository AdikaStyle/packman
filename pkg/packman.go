package packman

import (
	"github.com/AdikaStyle/packman/internal"
)

func Unpack(remote, path string, flagsMap map[string]string) error {
	return internal.M.TemplatingService.Unpack(remote, path, flagsMap)
}
