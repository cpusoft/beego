package utils

import (
	"github.com/beego/beego/core/utils"
)

// GetGOPATHs returns all paths in GOPATH variable.
func GetGOPATHs() []string {
	return utils.GetGOPATHs()
}
