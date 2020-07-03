package main

import (
	"errors"
	"fmt"
)


type Node struct {
	Val  interface{}
	Key interface{}
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}


type LRU struct {
	capacity int
	currSize int
	items map[interface{}]*Node
	dll DoublyLinkedList
}


func NewLRU(capacity int)(*LRU, error){
	if capacity <= 0 {
		return nil, errors.New("Capacity must be greater than zero")
	}
	lru := &LRU{
		capacity: capacity,
		items: make(map[interface{}]*Node),
	}
	return lru, nil
}


func (l *LRU)get(key interface{})(value interface{}, ok bool){
		if node,ok:=l.items[key];ok{
				l.dll.remove(node)
				l.dll.setHead(node)
				return node.Val,true
		}
		return
}

func ( l *LRU)set(key interface{}, val interface{}){
	if l.currSize>=l.capacity{
			deleted:=l.dll.removeLast()
			delete(l.items,deleted)
			value:=l.dll.setNewHead(key,val)
			l.items[key]=value
			return
	}
	
		value:=l.dll.setNewHead(key,val)
		l.items[key]=value
		l.currSize++
}

func (d *DoublyLinkedList) setNewHead(key interface{},val interface{})*Node {
	if d.Head==nil{
		value:=&Node{Val: val,Key:key}
		d.Head = value
		return value
	}
	curHead := d.Head
	value:=&Node{Key:key,Val: val, Next: curHead}
	d.Head = value
	curHead.Prev = value
	return value
} 


func (d *DoublyLinkedList) setHead(node *Node) {
	curHead := d.Head
	d.Head = node
	curHead.Prev = node
} 


func (d *DoublyLinkedList) removeLast() interface{} {
	if d.Head == nil {
		return -1
	}
	curr := d.Head
	for curr != nil {
		if curr.Next ==nil {
			prevNode :=curr.Prev
			prevNode.Next=nil
			return curr.Key
		}
		curr = curr.Next
	}
	return -1
}

func (d *DoublyLinkedList) remove(node *Node) {
	if d.Head == nil {
		return
	}
	curr := d.Head
	if (curr==node){
			curr.Next=nil
			return
	}

	for curr != nil {
		if curr ==node {
			if(curr.Next==nil){
				curr.Prev=nil
				return
			}
			if(curr.Prev==nil){
				curr.Next=nil
				return
			}

			nextNode := curr.Next
			prevNode := curr.Prev
			prevNode.Next = nextNode
			nextNode.Prev = prevNode
		}
		curr = curr.Next
	}
}

func(l *DoublyLinkedList)traverse(){
	curr:=l.Head
	for curr!=nil{
		fmt.Println("CURRR",curr.Key)
		curr=curr.Next
		}

}
func(l *LRU)traverse(){
for k, v := range l.items {
			fmt.Println("KEY",k,"VALUE",v)
	}
}

func main() {
	lru,_:=NewLRU(3)

	lru.set("A","Anakin")
	lru.set("B","Bobba Fett")
	lru.set("C","Captain Solo")
	lru.set("D","Darth Vader")
	lru.set("E","Emperor Palpatine ")

	fmt.Println("SIZE",lru.currSize)
	lru.dll.traverse()
	lru.traverse()
	val1,_:=lru.get("B")
	val2,_:=lru.get("C")
	val3,_:=lru.get("D")

	fmt.Println("VALUE 1",val1)
	fmt.Println("VALUE 2 ",val2)
	fmt.Println("VALUE 3",val3)

}

