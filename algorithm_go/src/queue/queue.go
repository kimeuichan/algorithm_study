package queue

const MAX = 1000005

type queue struct {
	items []interface{}
	head int
	tail int
}

func Queue() *queue{
	return &queue{
		items: make([]interface{}, MAX),
		head:  0,
		tail:  0,
	}
}

func (q *queue) Push(val interface{}) {
	q.items[q.tail] = val
	q.tail++
}

func (q *queue) Pop() interface{} {
	if q.head == q.tail{
		return -1
	}

	val := q.items[q.head]
	q.head++

	return val
}

func (q *queue) Front() interface{} {
	if q.Empty() == 1{
		return -1
	}

	return q.items[q.head]
}

func (q *queue) Back() interface{} {
	if q.Empty() == 1{
		return -1
	}

	return q.items[q.tail - 1]
}

func (q *queue) Size() int {
	return q.tail - q.head
}

func (q *queue) Empty() int {
	if q.head == q.tail{
		return 1
	}

	return 0
}

