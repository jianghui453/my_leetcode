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

import "testing"

func TestSimplifyPath(t *testing.T) {
	var path, hope, ret string

	path = "/home/"
	hope = "/home"
	ret = simplifyPath(path)
	t.Logf("%t path=%s hope=%s ret=%s", ret==hope, path, hope, ret)

	path = "/../"
	hope = "/"
	ret = simplifyPath(path)
	t.Logf("%t path=%s hope=%s ret=%s", ret==hope, path, hope, ret)

	path = "/./"
	hope = "/"
	ret = simplifyPath(path)
	t.Logf("%t path=%s hope=%s ret=%s", ret==hope, path, hope, ret)

	path = "/a/../"
	hope = "/"
	ret = simplifyPath(path)
	t.Logf("%t path=%s hope=%s ret=%s", ret==hope, path, hope, ret)

	path = "/home//foo/"
	hope = "/home/foo"
	ret = simplifyPath(path)
	t.Logf("%t path=%s hope=%s ret=%s", ret==hope, path, hope, ret)

	path = "/a/./b/../../c/"
	hope = "/c"
	ret = simplifyPath(path)
	t.Logf("%t path=%s hope=%s ret=%s", ret==hope, path, hope, ret)

	path = "/a/../../b/../c//.//"
	hope = "/c"
	ret = simplifyPath(path)
	t.Logf("%t path=%s hope=%s ret=%s", ret==hope, path, hope, ret)

	path = "/a//b////c/d//././/.."
	hope = "/a/b/c"
	ret = simplifyPath(path)
	t.Logf("%t path=%s hope=%s ret=%s", ret==hope, path, hope, ret)
}
