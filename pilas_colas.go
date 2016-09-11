package main

import (
	"fmt"
	"bufio"
	"os"
)

type Node struct {
	Ano string
  Genero string
	Titulo string
	Actor string
}


func (n *Node) String() string {
	return fmt.Sprint(n.Ano, "->" ,n.Genero,"->",n.Titulo,"->",n.Actor)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}
// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
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

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}
func buscar(s *Stack) *Stack{
	var b *Node
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Desea buscar la película por:\n 1. Año \n 2. Género \n 3. Título \n 4. Actor \n Digite el número de la opción elegida")
	campo, _ :=reader.ReadString('\n')
	campo = campo[:(len(campo)-1)]
	reader2 := bufio.NewReader(os.Stdin)
	var nombreCampo string= ""
	switch campo {
	case "1":
			nombreCampo="Año"
	case "2":
			nombreCampo="Género"
	case "3":
			nombreCampo="Título"
	default:
			nombreCampo="Actor"
	}
	fmt.Println("Digite el "+nombreCampo+" de la película")
	valor, _ :=reader2.ReadString('\n')
	valor = valor[:(len(valor)-1)]
	q := NewQueue(1)
	switch campo {
	case "1":
		b =s.Pop()
		for b!=nil{
			q.Push(b)
			if(b.Ano==valor){
				fmt.Println(b)
			}
			b =s.Pop()
		}
	case "2":
		b =s.Pop()
		for b!=nil{
			q.Push(b)
			if(b.Genero==valor){
				fmt.Println(b)
			}
			b =s.Pop()
		}
	case "3":
		b =s.Pop()
		for b!=nil{
			q.Push(b)
			if(b.Titulo==valor){
				fmt.Println(b)
			}
			b =s.Pop()
		}
	case "4":
		b =s.Pop()
		for b!=nil{
			q.Push(b)
			if(b.Actor==valor){
				fmt.Println(b)
			}
			b =s.Pop()
		}
	}
	var a int =1
	s = NewStack()
	for a==1{
		var b *Node =q.Pop()
		if(b != nil){
			s.Push(b)
		}else{
			a++
		}
	}
	return s
}

func agregar(s *Stack)  {
	fmt.Println("Digite el año en que se realizó la película")
	reader := bufio.NewReader(os.Stdin)
	ano, _ :=reader.ReadString('\n')
	ano = ano[:(len(ano)-1)]
	fmt.Println("Digite género de la película")
	genero, _ :=reader.ReadString('\n')
	genero = genero[:(len(genero)-1)]
	fmt.Println("Digite títitulo de la película")
	titulo, _ :=reader.ReadString('\n')
	titulo = titulo[:(len(titulo)-1)]
	fmt.Println("Digite el actor principal de la película")
	actor, _ :=reader.ReadString('\n')
	actor = actor[:(len(actor)-1)]
	s.Push(&Node{ano,genero,titulo,actor})

}
func imprimir (s *Stack) *Stack{
	var a int = 1
	q := NewQueue(1)
	for a==1{
		var b *Node =s.Pop()
		if(b != nil){
			q.Push(b)
			fmt.Println(b)
		}else{
			a++
		}
	}
	a=1
	s = NewStack()
	for a==1{
		var b *Node =q.Pop()
		if(b != nil){
			s.Push(b)
		}else{
			a++
		}
	}
	return s
}

func main() {
	s := NewStack()
	s.Push(&Node{"2012","Ciencia ficción","Los juegos del hambre", "Jennifer Lawrence"})
	s.Push(&Node{"2010","Acción","Karate Kid","Jennifer Lawrence"})
	s.Push(&Node{"2007","Terror","Rec","Pepito Perez"})
	s.Push(&Node{"2012","Terror","Rec","Pepito Perez"})
	s.Push(&Node{"2008","Terror","El payaso2","Jack Harrison"})
	//s=buscar(s)
	var a int=1
	for a==1{
		fmt.Println("\n --------- \n ")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Que le gustaría hacer"+
			""+":\n 1. Añadir una película"+
			""+"\n 2. Ver el catalogo de las películas existentes"+
			""+" \n 3. Buscar una película"+
			""+" \n 4. Salir")
		opcion, _ :=reader.ReadString('\n')
		opcion = opcion[:(len(opcion)-1)]
		switch opcion {
		case "1":
			agregar(s)
		case "2":
			s=imprimir(s)
		case "3":
			s=buscar(s)
		default:
			a=2
		}
	}


}
