package util

import (
	"github.com/truanter/yizhigo/config"
	"strconv"
	"strings"
)

func IsBlock(keyword, isBlockStr string) bool {
	isBlock, _ := strconv.Atoi(isBlockStr)
	if isBlock == 1 {
		for _, v := range config.GetVirtualProductKeyWords() {
			if strings.Contains(keyword, v) {
				return true
			}
		}
	}
	return false
}
