type MinStack struct {
    items []int
}

func Constructor() MinStack {
    return MinStack{
        items: []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
    if (len(this.items) == 0) {
        return
    }
    topIndex := len(this.items) - 1
    this.items = this.items[:topIndex]
}

func (this *MinStack) Top() int {
    if (len(this.items) == 0) {
        return 0
    }
    topIndex := len(this.items) - 1
    return this.items[topIndex]
}

func (this *MinStack) GetMin() int {
    min := math.MaxInt32
    for _, v := range this.items {
        if v < min {
            min = v
        }
    }
    return min
}
