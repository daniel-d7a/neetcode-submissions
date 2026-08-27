type MinStack struct {
	data []int
	min []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.min) > 0{
		this.min = append(this.min, min(this.min[len(this.min) - 1], val))
	}else{
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[: len(this.data) - 1]
	this.min = this.min[: len(this.min) - 1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data) - 1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min) - 1]
}
