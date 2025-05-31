package main

import (
	"errors"
	"fmt"
)

type Node interface {
	SetValue(v int) error
	GetValue() int
}

type SSLNode struct {
	next         *SSLNode
	value        int
	sNodeMessage string
}

func (sNode *SSLNode) SetValue(v int) error {
	if sNode == nil {
		return errors.New("Node is not valid")
	}

	sNode.value = v
	return nil
}

func (sNode *SSLNode) GetValue() int {
	return sNode.value
}

func NewSSLNode() *SSLNode {
	return &SSLNode{sNodeMessage: "Message from the normal Node type."}
}

type PowerNode struct {
	next         *PowerNode
	value        int
	pNodeMessage string
}

func (sNode *PowerNode) SetValue(v int) error {
	if sNode == nil {
		return errors.New("Node is not valid")
	}

	sNode.value = 10 * v
	return nil
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{pNodeMessage: "Message from the Power Node type."}
}

func createNode(v int) Node {
	node := NewPowerNode()
	node.SetValue(v)
	return node
}

func main() {
	node := createNode(5)

	switch concrete := node.(type) {
	case *SSLNode:
		fmt.Println("Type is SSLNode, message:", concrete.sNodeMessage)
	case *PowerNode:
		fmt.Println("Type is PowerNode, message:", concrete.pNodeMessage)
	}

	var n2 Node
	n2 = &PowerNode{value: 17}
	fmt.Printf("node: %v\n", n2)

	iStore := NewMagicStore("Integer Store")
	iStore.SetValue(4)
	if v, ok := iStore.GetValue().(int); ok {
		v *= 4
		fmt.Printf("v: %v\n", v)
	}

	sStore := NewMagicStore("String Store")
	sStore.SetValue("Hello")
	if v, ok := sStore.GetValue().(string); ok {
		v += " World"
		fmt.Printf("v: %v\n", v)
	}
}

type magicStore struct {
	value any // use in generic interfaces
	name  string
}

func (ms *magicStore) SetValue(v any) {
	ms.value = v
}

func (ms *magicStore) GetValue() any {
	return ms.value
}

func NewMagicStore(name string) *magicStore {
	return &magicStore{name: name}
}
