package main

import pqueue "github.com/nu7hatch/gopqueue"

type Task struct {
	Name     string
	priority int
}

func (t *Task) Priority() int {
	return t.priority 
}

func main() {
	q := pqueue.New(0)
	q.Enqueue(&Task{"one", 10})
	q.Enqueue(&Task{"two", 2})
	q.Enqueue(&Task{"three", 5})
	q.Enqueue(&Task{"four", 7})

	for i := 0; i < 4; i += 1 {
		task := q.Dequeue()
	    println(task.(*Task).Name)
	}
}