package structs

type Item interface {}
// 最基本的队列
type Queue struct {
  items []Item
}

func NewQueue() *Queue {
  return &Queue{items: []Item{}}
}

func (q *Queue) Enqueue(item Item) {
  q.items = append(q.items, item)
}

func (q *Queue) Dequeue() Item {
  dequeueItem := q.items[0]
  q.items = q.items[1:]
  return dequeueItem
}

func (q *Queue) Empty() bool {
  return len(q.items) == 0
}
