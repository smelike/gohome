package main

// Bit 数组

/*
Go 语言中的集合一般会用 map[T]bool 形式来表示，T 代表元素类型。

集合用 map 类型来表示虽然非常灵活，但我们可以以一种更好的形式来表示它。
例如在数据流分析领域，集合元素通常是一个非负整数，集合会包含很多元素，并且集合
会经常进行并集、交集操作，这种情况下，bit 数组会比 map 表现更加理想。

比如执行一个 http 下载任务，把文件按照 16 kb 一块划分为很多块，需要有一个全局
变量来标识哪些块下载完成了，这种时候也需要用到 bit 数组。

一个 bit 数组通常会用一个无符号数或称之为“字”的 slice 表示，每个元素的每一位都
表示集合里的一个值。当集合的第 i 位被设置时，我们才说这个集合包含元素 i。

因为每一个字都有 64 个二进制位，所以为了定位 x 的 bit 位，我们用了 x/64 的商
作为字的下标，并且用 x%64 得到的值作为这个字内的 bit 的所在位置。

UinonWith 这个方法里用到了 bit 位的“或”逻辑操作符号 | 来一次完成 64 个元素的或计算。

*/

/*
An IntSet is a set of small non-negative integers.
Its zero value represents the empty set.
*/
type IntSet struct {
	words []uint64
}

//
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	// s.words[word]&(1<<bit) != 0
	// word < len(s.words)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// |= 位或
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
