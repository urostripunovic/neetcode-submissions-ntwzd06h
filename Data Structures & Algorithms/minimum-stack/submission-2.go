type MinStack struct {
    items []int
    min []int
}

func Constructor() MinStack {
    return MinStack{
        items: []int{},
        min: []int{},
    }
}

func (this *MinStack) Push(val int) {
    if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
        this.min = append(this.min, val)
    }
    this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
    topIndex := len(this.items) - 1
    val := this.items[topIndex]
    if val == this.min[len(this.min) - 1] {
        topMinIndex := len(this.min) - 1
        this.min = this.min[:topMinIndex]
    }
    this.items = this.items[:topIndex]
}

func (this *MinStack) Top() int {
    topIndex := len(this.items) - 1
    val := this.items[topIndex]
    return val  
}

func (this *MinStack) GetMin() int {
    return this.min[len(this.min) - 1]
}

func min(a, b int) bool {
    if a < b {
        return true
    }
    return false
}
