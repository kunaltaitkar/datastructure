#  Data structure Implementation in Go Lang 


- [Linked List](https://github.com/kunaltaitkar/datastructure/tree/master/linkedList)

```

    package main

    import (
    	"datastructure/linkedList"
    	"fmt"
    )

    func main() {

    	list := linkedList.NewList()

    	input := []int{5, 4, 6, 7, 9, 10, 54, 99, 60} // random input for linked list

    	for index := 0; index < len(input); index++ {

    		newNode := list.CreateNode(input[index])

    		list.Insert(newNode) // to insert single node at last of the linked list

    	}

    	list.Display() //to- display value of nodes

    	list.Remove(&linkedList.Node{Value: 99})

    	fmt.Println("After remove")

    	list.Display() //to- display value of nodes

    }



```
