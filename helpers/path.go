package helpers

import (
	"path"
	"runtime"
)

func Path(dirName string) string {
	_, root, _, _ := runtime.Caller(0)
	switch dirName {
	case "root":
		return path.Join(path.Dir(root), "..")
	case "storage":
		return path.Join(path.Dir(root), "../storage")
	case "temp":
		return path.Join(path.Dir(root), "../storage/temp")
	case "media":
		return path.Join(path.Dir(root), "../storage/media")
	default:
		return path.Join(path.Dir(root), "..")
	}
}
