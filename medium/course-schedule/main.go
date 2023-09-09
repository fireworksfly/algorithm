package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	afterMp := map[int][]int{}           // 构建邻接表（v是k的后修课程）, k: 课程1，v：[课程0，课程2]
	inDegrees := make([]int, numCourses) // index: 某课程，v：该课程的入度
	for _, p := range prerequisites {
		after, first := p[0], p[1]                     // 后修课程, 先修课程
		afterMp[first] = append(afterMp[first], after) // 通过邻接表: 构建课程依赖关系
		inDegrees[after]++                             // 因学after之前必须先修first，所以after的入度+1
	}
	q := []int{}                                 // BFS的queue
	for courseIdx, inDegree := range inDegrees { // 初始化q: 入度为0的课程
		if inDegree == 0 {
			q = append(q, courseIdx)
		}
	}
	learnedCoursesCount := 0 // 学过的课程数量
	for len(q) > 0 {
		learnedCoursesCount++ // 学过的课程数量++
		learnedCourse := q[0] // q内均是入度为0的课, q.pop()则意为学完了此learnedCourse课
		q = q[1:]             // 同上, 共同完成q.pop()动作
		for _, after := range afterMp[learnedCourse] {
			inDegrees[after]-- // 因已学完after的先修课程, 故after的入度--
			if inDegrees[after] == 0 {
				q = append(q, after)
			} // 因after已无先修课程, 故可入队
		}
	}
	return learnedCoursesCount == numCourses // 是否课程全部学完
}

func main() {

}
