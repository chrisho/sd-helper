package uniqueset

type uniInterface interface {
	getMap() map[interface{}]bool
}

func UniqueAdd(data uniInterface, v interface{}) {
	m := data.getMap()
	if !m[v] {
		m[v] = true
	}
}

type StrSet struct {
	Unimap map[interface{}]bool
}

func (s StrSet) getMap() map[interface{}]bool {
	return s.Unimap
}
func (s StrSet) List() []string {
	strList := make([]string, 0, len(s.Unimap))
	for k, _ := range s.Unimap {
		if value, ok := k.(string); ok {
			strList = append(strList, value)
		}
	}
	return strList
}

type Int64Set struct {
	Unimap map[interface{}]bool
}

func (s Int64Set) getMap() map[interface{}]bool {
	return s.Unimap
}
func (s Int64Set) List() []int64 {
	intList := make([]int64, 0, len(s.Unimap))
	for k, _ := range s.Unimap {
		if value, ok := k.(int64); ok {
			intList = append(intList, value)
		}
	}
	return intList
}

type IntSet struct {
	Unimap map[interface{}]bool
}

func (s IntSet) getMap() map[interface{}]bool {
	return s.Unimap
}
func (s IntSet) List() []int {
	intList := make([]int, 0, len(s.Unimap))
	for k, _ := range s.Unimap {
		if value, ok := k.(int); ok {
			intList = append(intList, value)
		}
	}
	return intList
}
func NewStrSet() *StrSet {
	return &StrSet{Unimap: make(map[interface{}]bool)}
}
func NewInt64Set() *Int64Set {
	return &Int64Set{Unimap: make(map[interface{}]bool)}
}
func NewIntSet() *IntSet {
	return &IntSet{Unimap: make(map[interface{}]bool)}
}
