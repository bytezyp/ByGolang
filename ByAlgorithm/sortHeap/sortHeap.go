package main

func sortHeap(nums []int) []int  {
	n := len(nums)
	for i := (n-1-1)/2; i >=0 ; i-- {
		heapify(nums,n,i)
	}
	for i := n-1; i >0 ; i-- {
		nums[i],nums[0] = nums[0], nums[i]
		heapify(nums,i,0)
	}
	return nums
}

func heapify(nums []int, n,i int)  {
	/**
	n ： 数组的长度
	i： 待维护结点的下标，既不满足大根堆的定义的节点的下标
	 */
	largest := i
	lson := i*2 +1
	rson := i*2 +2
	if lson < n && nums[lson] <nums[largest] {
		largest = lson
	}
	if rson <n && nums[rson] < nums[largest] {
		largest = rson
	}
	if largest != i {
		nums[largest],nums[i] = nums[i],nums[largest]
		heapify(nums,n, largest)
	}
}

//---------------------------------------------

type MinHeap struct {
	items []int
}

func (h *MinHeap) Insert(item  int)  {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items)-1)
}

func (h *MinHeap)heapifyUp(index int)  {
	for index >0 {
		parentIndex := (index-1)/2
		if h.items[index] > h.items[parentIndex] {
			break
		}
		// 交换值
		h.items[index], h.items[parentIndex] = h.items[parentIndex], h.items[index]
		index = parentIndex
	}
}

func main()  {
	//arr := []int{1,3,77,44,4,33}
	//sortHeap(arr)
	//fmt.Print(arr)
	//minHeap := &MinHeap{}
	//minHeap.Insert(3)
	//minHeap.Insert(7)
	//minHeap.Insert(22)
	//minHeap.Insert(11)
	//minHeap.Insert(1)
	//minHeap.Insert(4)
	//fmt.Println(minHeap.items)
	//arr1 := make([]int, 10,20)
	//arr2 := arr1[8:]
	//arr2[0] = 11
	//fmt.Println(arr1,arr2)
}



