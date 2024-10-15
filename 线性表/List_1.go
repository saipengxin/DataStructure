package main_1

// 线性表顺序存储结构（数组实现）

const (
	MaxSize = 10 // 线性表最大容量
	ERROR   = -1 // 获取失败的返回结果
	OK      = 1  // 操作成功返回结果
)

type List struct {
	data   [MaxSize]int // 存储元素的数组
	length int          // 线性表当前长度
}

// GetElem 用e返回L线性表中第i个位置的数据（i是位置而不是下标）
func (L *List) GetElem(i int, e *int) int {
	if L.length <= 0 || i < 1 || i > L.length {
		return ERROR
	}
	*e = L.data[i-1]
	return OK
}

// ListInsert 在线性表L的第i个位置新增元素e
func (L *List) ListInsert(i int, e int) int {
	// i 小于 0 需要报错
	// i 大于当前线性表长度+1需要报错，否则数据插入后就不连续了
	// 线性表满了需要报错
	if i < 0 || i > L.length+1 || L.length == MaxSize {
		return ERROR
	}
	// 如果插入的数据不在表尾，那么后面的数据都需要移动
	if i <= L.length {
		// 将第i个位置的数据都向后移动一位，需要从后向前遍历，从后开始移动
		for index := L.length - 1; index > i-1; index-- {
			L.data[index+1] = L.data[index]
		}
	}
	// 将元素放在第i个位置
	L.data[i-1] = e
	L.length++
	return OK
}

// ListDelete 将顺序表L第i个位置的元素删除，并通过e返回其值
func (L *List) ListDelete(i int, e *int) int {
	// 删除位置小于1，报错
	// 删除位置比顺序表长度长，报错
	// 顺序表为空报错
	if i < 1 || i > L.length || L.length == 0 {
		return ERROR
	}
	*e = L.data[i-1]
	if i < L.length {
		for index := i; index <= L.length-1; index++ {
			L.data[index-1] = L.data[index]
		}
	}
	return OK
}
