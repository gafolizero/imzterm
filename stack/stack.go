package stack

import "os"

type Stack struct {
	History []int
}

func (s Stack) Peek() int {
	return s.History[len(s.History)-1]
}

func (s *Stack) Push(val int) {
	s.History = append(s.History, val)
}

func (s *Stack) Pop() {
	s.History = s.History[:len(s.History)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.History) == 0
}

func NewStack() *Stack {
	return &Stack{
		History: []int{},
	}
}

type ImgStack struct {
	ImgHistory []*os.File
}

func (is ImgStack) Peek() *os.File {
	return is.ImgHistory[len(is.ImgHistory)-1]
}

func (is *ImgStack) Push(newImgFile *os.File) {
	is.ImgHistory = append(is.ImgHistory, newImgFile)
}

func (is *ImgStack) Pop() {
	is.ImgHistory = is.ImgHistory[:len(is.ImgHistory)-1]
}

func (is *ImgStack) IsEmpty() bool {
	return len(is.ImgHistory) == 0
}

func NewImgStack() *ImgStack {
	return &ImgStack{
		ImgHistory: []*os.File{},
	}
}
