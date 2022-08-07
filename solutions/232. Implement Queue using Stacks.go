package solutions

type MyQueue struct {
	stack []int
	qp    int //low bound
	sp    int //high bound
}

func Constructor() MyQueue {
	return MyQueue{
		qp: 0,
		sp: 0,
	}
}

func (this *MyQueue) Push(x int) {
	if this.sp < len(this.stack) {
		this.stack[this.sp] = x
	} else {
		this.stack = append(this.stack, x)
	}
	this.sp++
}

func (this *MyQueue) Pop() int {
	if this.qp < this.sp {
		ret := this.stack[this.qp]
		this.qp++
		if this.qp == this.sp {
			this.qp = 0
			this.sp = 0
		}
		return ret
	}
	return 0
}

func (this *MyQueue) Peek() int {
	if this.qp < len(this.stack) {
		return this.stack[this.qp]
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	if this.qp < len(this.stack) {
		return false
	}
	return true
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

/*
type MyQueue struct {
	stack []int
	qp    int
}

func Constructor() MyQueue {
	return MyQueue{
		qp: 0,
	}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	if this.qp < len(this.stack) {
		this.qp++
		return this.stack[this.qp-1]
	}
	return 0
}

func (this *MyQueue) Peek() int {
	if this.qp < len(this.stack) {
		return this.stack[this.qp]
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	if this.qp < len(this.stack) {
		return false
	}
	return true
}
*/
