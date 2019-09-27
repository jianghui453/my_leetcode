package stack

import "strconv"

//根据逆波兰表示法，求表达式的值。
//
//有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
//说明：
//
//整数除法只保留整数部分。
//给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
//示例 1：
//
//输入: ["2", "1", "+", "3", "*"]
//输出: 9
//解释: ((2 + 1) * 3) = 9
//示例 2：
//
//输入: ["4", "13", "5", "/", "+"]
//输出: 6
//解释: (4 + (13 / 5)) = 6
//示例 3：
//
//输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
//输出: 22
//解释:
//  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22

func evalRPN(tokens []string) int {
    t := make([]string, 0)
    v := make([]int, 0)

    tmap := map[string]bool{
        "+": true,
        "-": true,
        "*": true,
        "/": true,
    }
    tl := len(tokens)
    for i := 0; i < tl; i ++ {
        token := tokens[i]
        if _, ok := tmap[token]; !ok {
            v2, err := strconv.Atoi(token)
            if err != nil {
                return 0
            }
            v = append(v, v2)
            continue
        }
        if len(v) > 1 {
            v1 := v[len(v)-2]
            v2 := v[len(v)-1]
            v = v[: len(v)-1]
            var v0 int
            switch token {
            case "+":
                v0 = v1+v2
                break
            case "-":
                v0 = v1-v2
                break
            case "*":
                v0 = v1*v2
                break
            case "/":
                v0 = v1/v2
            default:
                return 0
            }
            v[len(v)-1] = v0
            continue
        }
        t = append(t, token)
    }
    if len(v) == 0 {
        return 0
    }
    return v[0]
}
