package HashTable

import "fmt"

const HASHSIZE = 5

type HashKey int32

type MyHashNode struct {
	Key   HashKey
	Value string
	Next  *MyHashNode
}
type MyHashTable struct {
	nodes []*MyHashNode
}

//hash for hashing algorithm
func (h *MyHashTable) hash(src string) HashKey {
	var hashValue HashKey
	for _, v := range src {
		hashValue = hashValue*33 + HashKey(v)
	}
	return hashValue % HASHSIZE
}

func NewHashTable() *MyHashTable {
	res := make([]*MyHashNode, HASHSIZE)
	ht := new(MyHashTable)
	ht.nodes = res
	return ht
}

func (h *MyHashTable) LookUp(src string) *MyHashNode {
	index := h.hash(src)
	for v := h.nodes[index]; v != nil; v = v.Next {
		if v.Value == src {
			return v
		}
	}
	return nil
}

func (h *MyHashTable) Install(src string) {
	if np := h.LookUp(src); np == nil {
		index := h.hash(src)
		newNode := new(MyHashNode)
		newNode.Value = src
		newNode.Key = index
		newNode.Next = h.nodes[index]
		h.nodes[index] = newNode
		fmt.Println("add new HashNode")
	} else {
		np.Value = src
	}

}
