package c4f
import (
	"fmt"
)

func Println(str string) {
	if len(str) == 0 {
		fmt.Println("Bad string!")
	}
	fmt.Println("Awesome! %s", str);
}

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

}

func test() {
	
}
