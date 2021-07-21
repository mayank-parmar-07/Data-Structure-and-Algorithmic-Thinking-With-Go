package stacks

import (
	"errors"
	"fmt"
	"strconv"
)

type StacksStringNode struct {
	Val  string
	Next *StacksStringNode
}

type StringStack struct {
	Top  *StacksStringNode
	Size int
}

type IntNode struct {
	Val  int
	Next *IntNode
}

type IntStack struct {
	Top  *IntNode
	Size int
}

var mapOperatorPrecedence = map[string]int{"-": 1, "+": 2, "*": 3, "/": 4}

func NewStackString(type_of string) *StringStack {
	if type_of == "string" {
		stack := new(StringStack)
		stack.Size = 0
		stack.Top = nil
		return stack
	} else {
		return nil
	}
}

func NewStackInt() *IntStack {
	stack := new(IntStack)
	stack.Size = 0
	stack.Top = nil
	return stack

}

func (stack *StringStack) Push(val string) {
	newNode := new(StacksStringNode)
	newNode.Val = val
	newNode.Next = nil
	if stack.Top == nil {
		stack.Top = newNode
	} else {
		newNode.Next = stack.Top
		stack.Top = newNode
	}
	stack.Size += 1
}

func (stack *StringStack) Pop() (string, error) {
	if stack.Top == nil {
		return "", errors.New("Empty Stack")
	} else {
		val := stack.Top.Val
		stack.Top = stack.Top.Next
		stack.Size -= 1
		return val, nil
	}
}

func (stack *IntStack) Push(val int) {
	newNode := new(IntNode)
	newNode.Val = val
	newNode.Next = nil
	if stack.Top == nil {
		stack.Top = newNode
	} else {
		newNode.Next = stack.Top
		stack.Top = newNode
	}
	stack.Size += 1
}

func (stack *IntStack) Pop() (int, error) {
	if stack.Top == nil {
		return 0, errors.New("Empty Stack")
	} else {
		val := stack.Top.Val
		stack.Top = stack.Top.Next
		stack.Size -= 1
		return val, nil
	}
}

func (stack *IntStack) SeekTop() (int, error) {
	if stack.Top == nil {
		return 0, errors.New("Empty Stack")
	} else {
		return stack.Top.Val, nil
	}
}

func (stack *StringStack) GetSize() int {
	return stack.Size
}

func (stack *StringStack) GetTopValue() string {
	return stack.Top.Val
}

func (stack *StringStack) Display() {
	fmt.Println("Stacks Display")
	traverseNode := stack.Top
	for traverseNode != nil {
		fmt.Println(traverseNode.Val)
		traverseNode = traverseNode.Next
	}
}

func (stack *IntStack) Display() {
	fmt.Println("Stacks Display")
	traverseNode := stack.Top
	for traverseNode != nil {
		fmt.Println(traverseNode.Val)
		traverseNode = traverseNode.Next
	}
}

func (stack *StringStack) IsEmpty() bool {
	if stack.Top == nil {
		return true
	} else {
		return false
	}
}

func ConvertInfixToPostfix(expr string) string {
	newStack := NewStackString("string")
	ans := ""
	for _, char := range expr {
		if val_opr, ok := mapOperatorPrecedence[string(char)]; ok {
			if newStack.IsEmpty() {
				newStack.Push(string(char))
			} else {
				for true {
					if val, ok_second := mapOperatorPrecedence[newStack.GetTopValue()]; ok_second {
						if val >= val_opr {
							top_val, err := newStack.Pop()
							if err != nil {
								break
							}
							ans += string(top_val)
						} else {
							break
						}
					}
				}
				newStack.Push(string(char))
			}
		} else {
			ans += string(char)
		}
	}
	for !newStack.IsEmpty() {
		ans += newStack.GetTopValue()
		newStack.Pop()
	}
	return ans
}

func EvaluatePostFix(expr string) (int, error) {
	newStack := NewStackString("string")
	for _, char := range expr {
		if _, ok := mapOperatorPrecedence[string(char)]; ok {
			first_elem, err := newStack.Pop()
			if err != nil {
				return 0, errors.New("Incorrect Expression")
			}
			second_elem, err := newStack.Pop()
			if err != nil {
				return 0, errors.New("Incorrect Expression")
			}
			first_int, err := strconv.Atoi(first_elem)
			if err != nil {
				return 0, errors.New("Incorrect Expression")
			}
			second_int, err := strconv.Atoi(second_elem)
			if err != nil {
				return 0, errors.New("Incorrect Expression")
			}
			if string(char) == "+" {
				ans := second_int + first_int
				newStack.Push(fmt.Sprint(ans))
			} else if string(char) == "-" {
				ans := second_int - first_int
				newStack.Push(fmt.Sprint(ans))
			} else if string(char) == "*" {
				ans := second_int * first_int
				newStack.Push(fmt.Sprint(ans))
			} else if string(char) == "/" {
				ans := second_int / first_int
				newStack.Push(fmt.Sprint(ans))
			} else {
				return 0, errors.New("Incorrect Expression")
			}
		} else {
			newStack.Push(string(char))
		}
	}
	if newStack.Size == 1 {
		ans, err := newStack.Pop()
		if err != nil {
			return 0, errors.New("Incorrect Expression")
		}
		int_val, err := strconv.Atoi(ans)
		if err != nil {
			return 0, errors.New("Incorrect Expression")
		}
		return int_val, nil
	} else {
		return 0, errors.New("Incorrect Expression")
	}
}

func FindLargestSpan(arr []int) []int {
	span := make([]int, len(arr))
	span[0] = 1
	stack := NewStackInt()
	stack.Push(0)
	for index, val := range arr[1:] {
		index = index + 1
		stack_top := 0
		for true {
			if stack.Top == nil {
				break
			}
			stack_top, _ = stack.SeekTop()
			if arr[stack_top] <= val {
				_, err := stack.Pop()
				if err != nil {
					break
				}
			} else {
				break
			}
		}
		if stack.Top == nil {
			span[index] = index + 1
		} else {
			span[index] = index - stack_top
		}
		stack.Push(index)
	}
	return span
}

func FindMaxAreaHistogram(arr []int) int {
	stack := NewStackInt()
	ans := 0
	for index, val := range arr {
		max := 0
		stack_top, err := stack.SeekTop()
		if err != nil {
			stack.Push(index)
			continue
		}
		if val > arr[stack_top] {
			stack.Push(index)
		} else {
			stack.Pop()
			for stack.Top != nil {
				top_val, _ := stack.SeekTop()
				if arr[top_val] <= arr[stack_top] {
					stack.Pop()
				} else {
					break
				}
			}
			if stack.Top == nil {
				max = (stack_top + 1) * arr[stack_top]
				if max > ans {
					ans = max
				}
			} else {
				new_top, _ := stack.SeekTop()
				max = (index - new_top - 1) * arr[stack_top]
				if max > ans {
					ans = max
				}
			}
		}
		stack.Push(index)
	}
	return ans
}
func Panel() {
	stringStack := NewStackString("string")
	for i := 0; i < 10; i++ {
		stringStack.Push(fmt.Sprint(i))
	}
	stringStack.Display()
	for i := 0; i < 11; i++ {
		_, err := stringStack.Pop()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	postfix := ConvertInfixToPostfix("1+1")
	fmt.Println("Postfix is", postfix)
	fmt.Println("Postfix is", ConvertInfixToPostfix("2-1+2"))
	expr_result, err := EvaluatePostFix("123*+5-")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result is", expr_result)
	}

	fmt.Println("max span")
	arr := []int{6, 3, 4, 5, 2}
	result := FindLargestSpan(arr)
	fmt.Println("Span array is", result)
	histogram := []int{6, 2, 5, 4, 5, 1, 6}
	fmt.Println("Max histogram area is", FindMaxAreaHistogram(histogram))

}
