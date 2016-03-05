package OrderMaps

import (
	"fmt"
	"key"
	"reflect"
)

type OrderedMap struct {
	keys     key.Keys
	elemType reflect.Type
	m        map[interface{}]interface{}
}

type OrderedMap interface {
        Get(key interface{}) interface{}
        Put(key interface{},elem interface{}) (interface{},bool)
        Remove(key interface{}) interface{}
        Clear()
        Len()
        Contains(key interface{}) bool 
        FirstKey() interface{}
        LastKey() interface{}
        HeadMap(toKey interface{}) OrderedMap
        SubMap(fromKey interface{},toKey interface{}) OrderedMap
        TailMap(fromKey interface{}) OrderedMap
        Keys() []interface{}
        Elems() []interface{}
        ToMap() map[interface{}]interface{}
        KeyType() reflect.Type
        ElemType() reflect.Type
}


func (omap *OrderedMap) Get(key interface{}) interface{} {
        
}

func (omap *OrderedMap) Put(key interface{},elem interface{}) (interface{},bool){

}

func (omap *OrderedMap) Remove(key interface{}) interface{]{
        
}

func (omap *OrderedMap) Clear() {
        
}

func (omap *OrderedMap) Len() int{
        
}

func (omap *OrderedMap) Contains(key interface{}) bool{
        
}

func (omap *OrderedMap) FirstKey() interface{} {
        
}

func (omap *OrderedMap) LastKey() interface{} {
        
}

func (omap *OrderedMap) HeadMap(toKey interface{}) OrderedMap {
        
}

func (omap *OrderedMap) SubMAp(fromKey interface{},toKey interface{}) OrderedMap {
        
}

func (omap *OrderedMap) TailMap(fromKey interface{}) OrderedMap{
        
}

func (omap *OrderedMap) Keys() []interface{} {
        
}

func (omap *OrderedMap) Elems() []interface{}{
        
}

func (omap *OrderedMap) ToMap() map[interface{}]interface{} {
        
}

func (omap *OrderedMap) KeyType() reflect.Type {
        
}

func (omap *OrderedMap) ElemType() reflect.Type{

}

func NewOrderedMap(keys key.Keys,elem reflect.Type) *OrderedMap{
        return &OrderedMap{
                keys : key,
                elemType:elem,
                m:make(map[interface{}]interface{}),
        }
}
