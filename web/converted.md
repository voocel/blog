>Golang并没有内置set这种数据结构，我们知道，map 的 key 肯定是唯一的，而这恰好与 set 的特性一致，天然保证 set 中成员的唯一性。

```
package main

import (
	"fmt"
	"sync"
)

type set struct {
	m  map[string]struct{}
	mu sync.Mutex
}

func NewSet() *set {
	return &set{
		make(map[string]struct{}),
		sync.Mutex{},
	}
}

// 添加一个成员到集合中
func (s *set) Add(item ...string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range item {
		if v == "" {
			continue
		}
		s.m[v] = struct{}{}
	}
}

// 是否是集合成员
func (s *set) IsMember(item string) bool {
	_, ok := s.m[item]
	return ok
}

// 移除一个成员
func (s *set) Rem(item ...string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range item {
		delete(s.m, v)
	}
}

// 返回集合所有成员
func (s *set) Members() []string {
	var all []string
	for k:= range s.m {
		all = append(all, k)
	}
	return all
}

// 返回集合成员数量
func (s *set) Card() int {
	return len(s.m)
}

func main() {
	s := NewSet()
	//s.Add("a")
	//s.Add("b")
	//s.Add("c")
	//s.Add("")
	s.Add("a","","b","c")

	//fmt.Println(s.IsMember("a"))
	////s.Rem("a")
	//fmt.Println(s.IsMember("a"))
	fmt.Println(s.Card())
	fmt.Println(s.Members())
}

```