package OrderedMaps

import (
	"key"
	"reflect"
)

type myOrderedMap struct {
	keys     key.Keys
	elemType reflect.Type
	m        map[interface{}]interface{}
}

type OrderedMap interface {
	Get(key interface{}) interface{}
	Put(key interface{}, elem interface{}) (interface{}, bool)
	Remove(key interface{}) interface{}
	Clear()
	Len() int
	Contains(key interface{}) bool
	FirstKey() interface{}
	LastKey() interface{}
	HeadMap(toKey interface{}) OrderedMap
	SubMap(fromKey interface{}, toKey interface{}) OrderedMap
	TailMap(fromKey interface{}) OrderedMap
	Keys() []interface{}
	Elems() []interface{}
	ToMap() map[interface{}]interface{}
	KeyType() reflect.Type
	ElemType() reflect.Type
}

func (omap *myOrderedMap) Get(key interface{}) interface{} {
	if !omap.Contains(key) {
		return nil
	} else {
		return omap.m[key]
	}
}

func (omap *myOrderedMap) Put(key interface{}, elem interface{}) (interface{}, bool) {
	if !omap.Contains(key) {
		omap.m[key] = elem
		omap.keys.Add(key)
		return nil, true
	} else {
		oldMap := omap.m[key]
		omap.m[key] = elem
		return oldMap, false
	}
}

func (omap *myOrderedMap) Remove(key interface{}) interface{} {
	if !omap.Contains(key) {
		return nil
	} else {
		oldKey := omap.m[key]
		omap.m[key] = nil
		omap.keys.Remove(key)
		return oldKey
	}
}

func (omap *myOrderedMap) Clear() {
	omap.m = make(map[interface{}]interface{}, 0)
	omap.keys.Clear()
}

func (omap *myOrderedMap) Len() int {
	return omap.keys.Len()
}

func (omap *myOrderedMap) Contains(key interface{}) bool {
	if omap.m[key] == nil {
		return false
	} else {
		return true
	}
}

func (omap *myOrderedMap) FirstKey() interface{} {
	err := omap.Contains(omap.keys.Get(0))
	if err {
		return omap.keys.Get(0)
	} else {
		return nil
	}
}

func (omap *myOrderedMap) LastKey() interface{} {
	err := omap.Contains(omap.keys.Get(omap.keys.Len() - 1))
	if err {
		return omap.keys.Get(omap.keys.Len() - 1)
	} else {
		return nil
	}
}

func (omap *myOrderedMap) HeadMap(toKey interface{}) OrderedMap {
	if !omap.Contains(toKey) {
		return nil
	} else {
		newMap := NewOrderedMap(omap.keys, omap.elemType)
		for _, value := range omap.keys.GetAll() {
			if value == toKey {
				break
			}
			newMap.keys.Add(value)
			newMap.m[value] = omap.m[value]
		}
		var newMapX OrderedMap
		newMapX = newMap
		return newMapX
	}
}

func (omap *myOrderedMap) SubMap(fromKey interface{}, toKey interface{}) OrderedMap {
	if !omap.Contains(fromKey) || !omap.Contains(toKey) {
		return nil
	} else {
		newMap := NewOrderedMap(omap.keys, omap.elemType)
		for _, value := range omap.keys.GetAll() {
			if value != fromKey {
				continue
			}
			newMap.keys.Add(value)
			newMap.m[value] = omap.m[value]
			if value == toKey {
				break
			}
		}
		var newMapX OrderedMap
		newMapX = newMap
		return newMapX
	}
}

func (omap *myOrderedMap) TailMap(fromKey interface{}) OrderedMap {
	if !omap.Contains(fromKey) {
		return nil
	} else {
		newMap := NewOrderedMap(omap.keys, omap.elemType)
		for _, value := range omap.keys.GetAll() {
			if value != fromKey {
				continue
			}
			newMap.keys.Add(value)
			newMap.m[value] = omap.m[value]
		}
		var newMapX OrderedMap
		newMapX = newMap
		return newMapX
	}
}

func (omap *myOrderedMap) Keys() []interface{} {
	var index int = 0
	mapKeySlice := make([]interface{}, omap.keys.Len())
	for index < omap.keys.Len() {
		t := omap.keys.Get(index)
		mapKeySlice[index] = t
		index++
	}
	if index == omap.Len() {
		return mapKeySlice
	} else {
		return nil
	}
}

func (omap *myOrderedMap) Elems() []interface{} {
	index := 0
	mapValueSlice := make([]interface{}, omap.Len())
	for index < omap.Len() {
		mapValueSlice[index] = omap.m[omap.keys.Get(index)]
		index++
	}
	if index == omap.Len() {
		return mapValueSlice
	} else {
		return nil
	}
}

func (omap *myOrderedMap) ToMap() map[interface{}]interface{} {
	newMap := make(map[interface{}]interface{}, 0)
	valueSlice := omap.Elems()
	for index, key := range omap.Keys() {
		newMap[key] = valueSlice[index]
	}
	return newMap
}

func (omap *myOrderedMap) KeyType() reflect.Type {
	return omap.keys.ElemType()
}

func (omap *myOrderedMap) ElemType() reflect.Type {
	return omap.elemType
}

func NewOrderedMap(keys key.Keys, elem reflect.Type) *myOrderedMap {
	return &myOrderedMap{
		keys:     keys,
		elemType: elem,
		m:        make(map[interface{}]interface{}),
	}
}
