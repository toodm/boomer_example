package main

import (
	"orm_test_platform/action"

	"github.com/toodm/boomer"
)

func main() {

	task1 := &boomer.Task{
		Name:   "foo",
		Weight: 10,
		Fn:     action.Action_10001,
	}

	//	task2 := &boomer.Task{
	//		Name:   "bar",
	//		Weight: 20,
	//		Fn:     bar,
	//	}

	boomer.Run(task1)
	//	for i := 1; i <= 10; i++ {
	//		action.Action_10001()
	//	}

}
