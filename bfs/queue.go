package bfs

import "image"

type Queue struct {
	elements []image.Point
}

func (q *Queue) pop() image.Point {
	f := len(q.elements)
	rv := q.elements[f-1]
	q.elements = append((q.elements)[:f-1])
	return rv
}

func (q *Queue) push(point image.Point) {
	q.elements = append(q.elements, point)
}

func (q Queue) len() int {
	return len(q.elements)
}
