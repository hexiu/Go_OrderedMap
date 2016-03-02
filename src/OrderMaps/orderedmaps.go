package OrderMaps

import (
	"fmt"
)

type OrderedMap struct {
	keys []interface{}
	m    map[interface{}]interface{}
}

func (omap *OrderedMap) Len() int {
	return len(omap.keys)
}

func (omap *OrderedMap) Less(i, j int) bool {

}

func (omap *OrderedMap) Swap(i, j int) {

}
