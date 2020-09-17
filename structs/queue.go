package structs

// 最基本的队列
type Queue struct {
  items []int
}

func (q *Queue) New() *Queue {
  q.items = []int{}
  return q
}

func (q *Queue) Enqueue(item int) {
  q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
  dequeueItem := q.items[0]
  q.items = q.items[1:]
  return dequeueItem
}
