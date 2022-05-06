package main

import "fmt"

type Vertex struct {
	value             string
	adjacent_vertices []*Vertex
}

func initializeVertex(value string) Vertex {
	newV := Vertex{}
	newV.value = value
	newV.adjacent_vertices = []*Vertex{}
	return newV
}

func (v *Vertex) add_adjacent_vertex(toAdd *Vertex) {
	var alreadyThere bool = false
	for _, elem := range v.adjacent_vertices {
		if elem == toAdd {
			alreadyThere = true
		}
	}
	if !alreadyThere {
		v.adjacent_vertices = append(v.adjacent_vertices, toAdd)
		toAdd.add_adjacent_vertex(v)
	}

}

func dfsTraversal(v *Vertex, visited map[*Vertex]bool) {
	visited[v] = true

	fmt.Println(v.value)

	for _, elem := range v.adjacent_vertices {
		if !(visited[elem] == true) {
			dfsTraversal(elem, visited)
		}
	}
}

func dfsSearch(v *Vertex, visited map[*Vertex]bool, search *Vertex) bool {
	if search == v {
		return true
	}

	visited[v] = true

	for _, elem := range v.adjacent_vertices {
		if search == elem {
			return true
		}
		if !(visited[elem] == true) {
			dfsSearch(elem, visited, search)
		}
	}

	return false
}

func bfsTraversal(v *Vertex) {
	queue := []*Vertex{}
	queue = append(queue, v)

	visited := map[*Vertex]bool{}
	visited[v] = true

	for len(queue) > 0 {
		currentV := queue[0]
		queue = queue[1:]
		fmt.Println(currentV.value)

		for _, elem := range currentV.adjacent_vertices {
			_, exists := visited[elem]
			if !exists {
				visited[elem] = true
				queue = append(queue, elem)
			}
		}

	}
}

func main() {

	alice := initializeVertex("alice")
	bob := initializeVertex("bob")
	cynthia := initializeVertex("cynthia")
	//pete := initializeVertex("pete")
	alice.add_adjacent_vertex(&bob)
	alice.add_adjacent_vertex(&cynthia)
	bob.add_adjacent_vertex(&cynthia)
	cynthia.add_adjacent_vertex(&bob)
	//dfsTraversal(&alice, map[*Vertex]bool{})
	// fmt.Println(dfsSearch(&alice, map[*Vertex]bool{}, &pete))
	// fmt.Println(dfsSearch(&alice, map[*Vertex]bool{}, &bob))
	bfsTraversal(&alice)

}
