package main

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
