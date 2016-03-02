package Key

import (
	"reflect"
	"sort"
)

type myKeys struct {
	container   []interface{}
	compareFunc func(interface{}, interface{}) int8
	elemType    reflect.Type
}

type keys interface {
	sort.Interface
	Add(k interface{}) bool
	Remove(k interface{}) bool
	Clear()
	Get(index int) interface{}
	GetAll() []interface{}
	Search(k interface{}) (index int, contains bool)
	ElemType() reflect.Type
	CompareFunc() func(interface{}, interface{}) int8
}

func main() {
	// This is init myKeys!
	/*
	   int64Keys:=&myKeys{
	           container : make([]interface{},0)
	           compareFunc func(e1 interface{},e2 interface{}) int8 {
	                   k1:=e1.(int64)
	                   k2:=e2.(int64)
	                   if k1<k2 {
	                           return -1
	                   }else if k1>k2 {
	                           return 1
	                   } else {
	                           return 0
	                   }
	           },
	           elemType : reflect.TypeOf(int64(1))
	   }
	*/

}

/*
func NewMyKeys() {
        return &myKeys{
                   container : make([]interface{},0),
                   compareFunc func(e1 interface{},e2 interface{}) int8 {
                           k1:=e1.(int64)
                           k2:=e2.(int64)
                           if k1<k2 {
                                   return -1
                           }else if k1>k2 {
                                   return 1
                           } else {
                                   return 0
                           }
                   },
                   elemType : reflect.TypeOf(int64(1))}
}
*/

func (keys *myKeys) Len() int {
	return len(keys.container)
}

// if is true , k1<k2,    else k1>=k2
func (keys *myKeys) Less(i, j int) bool {
	return keys.compareFunc(keys.container[i], keys.container[j]) == -1
}

func (keys *myKeys) Swap(i, j int) {
	keys.container[i], keys.container[j] = keys.container[j], keys.container[i]
}

func (keys *myKeys) isAcceptableElem(k interface{}) bool {
	if k == nil {
		return false
	}
	if reflect.TypeOf(k) != keys.elemType {
		return false
	}
	return true
}

func (keys *myKeys) Add(k interface{}) bool {
	ok := keys.isAcceptableElem(k)
	if !ok {
		return false
	}
	keys.container = append(keys.container, k)
	sort.Sort(keys)
	return true
}

func Search(n int, f func(int) bool) int {

}

func (keys *myKeys) Clear() {
	keys.container = make([]interface{}, 0)
}

func (keys *myKeys) Get(index int) interface{} {
	if index >= key.Len() {
		return nil
	}
	return keys.container[index]
}

func (keys *myKeys) GetAll() []interface{} {

}
