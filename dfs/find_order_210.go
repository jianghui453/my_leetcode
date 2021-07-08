package dfs

func findOrder210(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		courses []int
		state   = make([]int, numCourses)
		dfs     func(int)
		invalid = false
	)

	dfs = func(i int) {
		state[i] = 1
		for _, course := range edges[i] {
			if state[course] == 2 {
				continue
			}
			if state[course] == 1 {
				invalid = true
				return
			}
			dfs(course)
			if invalid {
				return
			}
		}
		state[i] = 2
		courses = append(courses, i)
	}

	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
	}

	for i := 0; i < numCourses; i++ {
		if state[i] == 2 {
			continue
		}
		dfs(i)
		if invalid {
			return []int{}
		}
	}

	for i := 0; i < numCourses/2; i++ {
		courses[i], courses[numCourses - 1 - i] = courses[numCourses - 1 - i], courses[i]
	}

	return courses
}
