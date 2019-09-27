package hash_table

import "sort"

//给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。
//
//示例 1:
//
//输入: [[1,1],[2,2],[3,3]]
//输出: 3
//解释:
//^
//|
//|        o
//|     o
//|  o  
//+------------->
//0  1  2  3  4
//示例 2:
//
//输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
//输出: 4
//解释:
//^
//|
//|  o
//|     o        o
//|        o
//|  o        o
//+------------------->
//0  1  2  3  4  5  6

func maxPoints(points [][]int) int {
    if len(points) < 2 {
        return len(points)
    }
    answer := 2
    for len(points) > answer {
        duplicate := 0
        for idx, point := range points {
            if point[0] < points[0][0] || (point[0] == points[0][0] && point[1] < points[0][1]) {
                p := points[0]
                points[0] = points[idx]
                points[idx] = p
            }
        }
        basePoint := points[0]
        points = points[1:]
        for i := 0; i < len(points); i++ {
            if points[i][0] == basePoint[0] && points[i][1] == basePoint[1] {
                duplicate++
                points[i] = points[len(points)-1]
                points = points[:len(points)-1]
            }
        }

        sort.Slice(points, func(i int, j int) bool {
            return (points[i][1] - basePoint[1]) * (points[j][0] - basePoint[0]) -
                (points[j][1] - basePoint[1]) * (points[i][0] - basePoint[0]) < 0
        })

        cc := 1
        for i := 0; i < len(points); i++ {
            if i == 0 || (points[i][1] - basePoint[1]) * (points[i-1][0] - basePoint[0]) -
                (points[i-1][1] - basePoint[1]) * (points[i][0] - basePoint[0]) == 0 {
                cc++
            } else {
                cc = 2
            }
            if cc + duplicate > answer {
                answer = cc + duplicate
            }
        }


    }
    return answer
}

//func maxPoints(points [][]int) int {
//    if len(points) < 3 {
//        return len(points)
//    }
//    rcd := make(map[int]map[int]map[string]int)
//    max := 0
//    newPoints := make([][]int, 0)
//    his := make(map[int]map[int]bool)
//    for i := 0; i < len(points); i ++ {
//        if _, ok := his[points[i][0]][points[i][1]]; !ok {
//            if _, ok := his[points[i][0]]; !ok {
//                his[points[i][0]] = make(map[int]bool)
//            }
//            his[points[i][0]][points[i][1]] = true
//            newPoints = append(newPoints, points[i])
//        }
//        if _, ok := rcd[points[i][0]]; !ok {
//            rcd[points[i][0]] = make(map[int]map[string]int)
//        }
//        if _, ok := rcd[points[i][0]][points[i][1]]; !ok {
//            rcd[points[i][0]][points[i][1]] = make(map[string]int)
//        }
//        if _, ok := rcd[points[i][0]][points[i][1]]["cnt"]; !ok {
//            rcd[points[i][0]][points[i][1]]["cnt"] = 1
//        } else {
//            rcd[points[i][0]][points[i][1]]["cnt"] ++
//        }
//        max = getMax(max, rcd[points[i][0]][points[i][1]]["cnt"])
//    }
//
//    //fmt.Printf("newPoints=%v\n", newPoints)
//    for i := 0; i < len(newPoints); i ++ {
//        for j := i+1; j < len(newPoints); j ++ {
//            x1, x2, y1, y2 := newPoints[i][0], newPoints[j][0], newPoints[i][1], newPoints[j][1]
//            gcd := getGCD(y2-y1, x2-x1)
//            y := (y2-y1)/gcd
//            x := (x2-x1)/gcd
//
//            var k string
//            if x*y<0 {
//                k = "-" + strconv.Itoa(x) + "-" + strconv.Itoa(y)
//            } else {
//                k = strconv.Itoa(x) + "-" + strconv.Itoa(y)
//            }
//
//            if _, ok := rcd[x1][y1][k]; !ok {
//                if _, ok := rcd[x1]; !ok {
//                    rcd[x1] = make(map[int]map[string]int)
//                    rcd[x1][y1] = make(map[string]int)
//                } else if _, ok := rcd[x1][y1]; !ok {
//                    rcd[x1][y1] = make(map[string]int)
//                }
//                rcd[x1][y1][k] = rcd[x1][y1]["cnt"] + rcd[x2][y2]["cnt"]
//            } else {
//                rcd[x1][y1][k] += rcd[x2][y2]["cnt"]
//            }
//            max = getMax(max, rcd[x1][y1][k])
//
//            if _, ok := rcd[x2][y2]; !ok {
//                if _, ok := rcd[x2]; !ok {
//                    rcd[x2] = make(map[int]map[string]int)
//                    rcd[x2][y2] = make(map[string]int)
//                } else if _, ok := rcd[x2][y2]; !ok {
//                    rcd[x2][y2] = make(map[string]int)
//                }
//                rcd[x2][y2][k] = rcd[x2][y2]["cnt"] + rcd[x1][y1]["cnt"]
//            } else {
//                rcd[x2][y2][k] += rcd[x1][y1]["cnt"]
//                max = getMax(max, rcd[x2][y2][k])
//            }
//        }
//    }
//    return max
//}
//
//func getGCD(x, y int) int {
//    if y == 0 {
//        return x
//    }
//    return getGCD(y, x%y)
//}
//
//func getMax(x, y int) int {
//    if x > y {
//        return x
//    }
//    return y
//}
