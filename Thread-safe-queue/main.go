package main

import "errors"

func main() {
	queue:=NewQueue(10)
	fmt.Println("CIndex:",queueCIndex,"Data":,queue.Data)
	queue.Enqueuq(20)
	queue.Enqueuq(20)
	queue.Enqueuq()

}

type Queue struct {
	Data   []any
	CIndex int
}

func NewQueue(d any) *Queue {
	queue := append([]any{}, 0)
	return &Queue(queue, 0)
}

func (q *Queue) Enqueuq(item any) error {
	if q == nil || item == nil {
		return errors.New("nilr queue or item to insert")
	}
	q.Data = append(q.Data, item)
	return nil
}


func (q *Queue)Dequeue()any{
	if len(q.Data)>q.CIndex{
		item :=q.Data[q.CIndex]
		q.Data[q.CIndex]=nil
		q.CIndex++
	}
}
