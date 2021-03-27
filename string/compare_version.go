package string

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	versionStrs1 := strings.Split(version1, ".")
	versionStrs2 := strings.Split(version2, ".")
	
	var trimAndConv func(string) int
	trimAndConv = func(versionStr string) int {
		if strings.HasPrefix(versionStr, "0") {
			return trimAndConv(strings.TrimPrefix(versionStr, "0"))
		}
		num, _ := strconv.Atoi(versionStr)
		return num
	}

	for i := 0; i < len(versionStrs1) && i < len(versionStrs2); i++ {
		num1, num2 := trimAndConv(versionStrs1[i]), trimAndConv(versionStrs2[i])
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}

	if len(versionStrs1) > len(versionStrs2) {
		for i := len(versionStrs2); i < len(versionStrs1); i++ {
			if trimAndConv(versionStrs1[i]) > 0 {
				return 1
			}
		}
	} else {
		for i := len(versionStrs1); i < len(versionStrs2); i++ {
			if trimAndConv(versionStrs2[i]) > 0 {
				return -1
			}
		}
	}
	return 0
}
