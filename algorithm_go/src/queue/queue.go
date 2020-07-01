package queue

type Queue struct {
	items []int
	head int
	tail int
}

func (q *Queue) Push(val int) {
	q.items[q.tail] = val
	q.tail++
}

func (q *Queue) Pop() int {
	if q.head == q.tail{
		return -1
	}

	val := q.items[q.head]
	q.head++

	return val
}

func (q *Queue) Front() int {
	if q.Empty() == 1{
		return -1
	}

	return q.items[q.head]
}

func (q *Queue) Back() int {
	if q.Empty() == 1{
		return -1
	}

	return q.items[q.tail - 1]
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Empty() int {
	if q.head == q.tail{
		return 1
	}

	return 0
}

