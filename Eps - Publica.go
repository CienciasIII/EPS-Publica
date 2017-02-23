package main

import (
	"fmt"
)

var lista[] Estructura

type Node struct{
	name string
	id string
	Eps string
}

type Estructura struct{
	nom string
	conteo int
}

func (n *Node) String() string {
	return fmt.Sprint(n.name, "->" ,n.id)
}

func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	buscar(node.Eps)
	return node
}

func buscar(eps string){
	var token bool = false
	for i := 0; i < len(lista); i++ {
		if lista[i].nom == eps{
			token = true
			lista[i].conteo = lista[i].conteo+1
		}
	}
	if token==false{
		s := Estructura{nom:eps,conteo:1}
		lista = append(lista,s)
	}
} 

func main() {

  fmt.Println("Banco de Datos EPS")
	q := NewQueue(1)
	q.Push(&Node{"Roberto","1022455","PremiSalud"})
	q.Push(&Node{"Viviana","3516333","Sura"})
	q.Push(&Node{"Carlos","3552123","Sisben"})
	q.Push(&Node{"Daniel","876465","Famisanar"})
	q.Push(&Node{"Santiago","76598","SaludCop"})
	q.Push(&Node{"Santiago","76598","SaludCop"})
	q.Push(&Node{"Santiago","76598","SaludCop"})
	fmt.Println(q.Pop(), q.Pop(), q.Pop(),q.Pop(),q.Pop(),q.Pop(),q.Pop())
	fmt.Println(lista)
}