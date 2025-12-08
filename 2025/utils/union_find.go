package utils

import "sort"

type UnionFind struct {
	Parent []int
	Size   []int
	Count  int
}

func NewUnionFind(numOfElements int) *UnionFind {
	parent := make([]int, numOfElements)
	size := make([]int, numOfElements)
	for i := range numOfElements {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		Parent: parent,
		Size:   size,
		Count:  numOfElements,
	}
}

func (uf *UnionFind) Find(node int) int {
	for node != uf.Parent[node] {
		uf.Parent[node] = uf.Parent[uf.Parent[node]]
		node = uf.Parent[node]
	}
	return node
}

func (uf *UnionFind) Union(node1, node2 int) {
	root1 := uf.Find(node1)
	root2 := uf.Find(node2)

	if root1 == root2 {
		return
	}

	if uf.Size[root1] > uf.Size[root2] {
		uf.Parent[root2] = root1
		uf.Size[root1] += 1
	} else {
		uf.Parent[root1] = root2
		uf.Size[root2] += 1
	}

	uf.Count -= 1
}

func (uf *UnionFind) GetSetsNumberOfElements() []int {
	sets := make(map[int]int)

	for _, node := range uf.Parent {
		root := uf.Find(node)
		sets[root] += 1
	}

	var sizes []int
	for _, size := range sets {
		sizes = append(sizes, size)
	}

	sort.Ints(sizes)

	return sizes
}
