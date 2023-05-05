package main

type Graph struct {
	// 节点个数
	vertices int
	// 边 权重
	edges map[int]map[int]int
	// true:有向;false无向
	Directed bool
}

func New(v int) *Graph {
	return &Graph{
		vertices: v,
	}
}

func (g *Graph) AddVertex(v int)  {
	if g.edges == nil {
		g.edges = make(map[int]map[int]int)
	}

	_, ok := g.edges[v]
	if !ok {
		g.edges[v] = make(map[int]int)
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.AddWeightedEdge(v1, v2, 0)
}

func (g *Graph) AddWeightedEdge(v1, v2, weight int) {
	g.AddVertex(v1)
	g.AddVertex(v2)

	g.edges[v1][v2] = weight
	// 无向图两个节点互相添加权重
	if !g.Directed {
		g.edges[v2][v1] = weight
	}

}