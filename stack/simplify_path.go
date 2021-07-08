//以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。
//
//在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径
//
//请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 / 结尾。此外，规范路径必须是表示绝对路径的最短字符串。
//
//
//
//示例 1：
//
//输入："/home/"
//输出："/home"
//解释：注意，最后一个目录名后面没有斜杠。
//示例 2：
//
//输入："/../"
//输出："/"
//解释：从根目录向上一级是不可行的，因为根是你可以到达的最高级。
//示例 3：
//
//输入："/home//foo/"
//输出："/home/foo"
//解释：在规范路径中，多个连续斜杠需要用一个斜杠替换。
//示例 4：
//
//输入："/a/./b/../../c/"
//输出："/c"
//示例 5：
//
//输入："/a/../../b/../c//.//"
//输出："/c"
//示例 6：
//
//输入："/a//b////c/d//././/.."
//输出："/a/b/c"

package stack

import (
	"strings"
)

func simplifyPath(path string) string {
	s := make([]string, 0)
	paths := strings.Split(path, "/")
	l := len(paths)

	for i := 0; i < l; i++ {
		switch paths[i] {
		case "":
			continue
		case ".":
			continue
		case "..":
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		default:
			s = append(s, paths[i])
		}
	}

	ret := strings.Join(s, "/")
	return "/" + ret
}

// func simplifyPath(path string) string {
// 	slicePath := strings.Split(path, "/")
// 	lenSlicePath := len(slicePath)
// 	slicePathRet := make([]string, 0)
// 	for i := 0; i < lenSlicePath; i++ {
// 		switch slicePath[i] {
// 		case "..":
// 			if len(slicePathRet) > 0 {
// 				slicePathRet = slicePathRet[:len(slicePathRet)-1]
// 			}
// 			break
// 		case "/":
// 			break
// 		case ".":
// 			break
// 		case " ":
// 			break
// 		case "":
// 			break
// 		default:
// 			slicePathRet = append(slicePathRet, slicePath[i])
// 			break
// 		}
// 	}
// 	return "/" + strings.Join(slicePathRet, "/")
// }
