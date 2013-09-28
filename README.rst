Priority Queue in Go
====================

This package provides a priority queue implementation and scaffold
interfaces.

Installation
------------
Use the ``goinstall`` tool::

	$ goinstall github.com/nu7hatch/gopqueue

... or install it manually::

	$ git clone git://github.com/nu7hatch/gopqueue.git
	$ cd gopqueue
	$ make install

Usage
-----
Here's trivial example of the fast queue usage::

	package main

	import pqueue "github.com/nu7hatch/gopqueue"
	
	type Task struct {
	    Name     string
	    priority int
	}

	func (t *Task) Less(other interface{}) bool {
	    return t.priority < other.(*Task).priority
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

	// Produces:
	//
	//     two
	//     three
	//     four
	//     one

For more information and examples check the package documentation.
	
Copyright
---------
Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>

See COPYING file for details.
