package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	g     SimpleGraph
	data0 = `
0:1 3 4
1:0 2
2:1
3:0
4:0
`
	data1 = `
0 1
1 2
2 0
3 4
`
	data2 = `
0:1 5
1:0 5 3 2
2:1 3
3:1 2 4 5
4:3 5
5:0 1 3 4
`
	data3 = `
0:1
1:2 3
2:1 4
3:1
4:2
`
	data4 = `
0:1 2
1:2 3
2:1 4
3:1
4:2
`
)

func init() {
	g = NewGraphByData(data1)
}

func TestNewGraphByData(t *testing.T) {
	g = NewGraphByData(data1)
	t.Log(g)
	t.Log(g.Adj(0))
}

func TestNewGraphByAdjacencyList(t *testing.T) {
	g = NewGraphByAdjacencyList(data2)
	t.Log(g)
	t.Log(g.Adj(1))
}

func TestSearch1(t *testing.T) {
	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> DFSearch >>>>")
	g = NewGraphByAdjacencyList(data0)
	search := new(DFSearch).GenSearch(g, 0)
	t.Log("path_to [0-2]:", search.PathTo(2))
}
func TestSearch2(t *testing.T) {
	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> DFSearch >>>>")
	search := new(DFSearch).GenSearch(g, 0)
	t.Log("graph", g)
	t.Log("total number of 0 connected", search.Count())
	t.Log("is 2 connected with 0 :", search.Marked(2))
	t.Log("is 3 connected with 0 :", search.Marked(3))
	t.Log("is 4 connected with 0 :", search.Marked(4))
	t.Log("path_to [0-2]:", search.PathTo(2))

	t.Log("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< DFSearch <<<<")
	t.Log(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> BFSearch >>>>")
	search = new(BFSearch).GenSearch(g, 0)
	t.Log("graph", g)
	t.Log("total number of 0 connected", search.Count())
	t.Log("is 2 connected with 0 :", search.Marked(2))
	t.Log("is 3 connected with 0 :", search.Marked(3))
	t.Log("is 4 connected with 0 :", search.Marked(4))
	t.Log("path_to [0-2]:", search.PathTo(2))
	t.Log("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< BFSearch <<<<")
}
func TestCycle(t *testing.T) {
	c := NewCycle(g)
	t.Log(c.HasCycle())
}

func TestTowColor(t *testing.T) {
	yes := NewGraphByAdjacencyList(data3)
	no := NewGraphByAdjacencyList(data4)
	assert.Equal(t, NewTowColor(yes).IsBipartite(), true)
	assert.Equal(t, NewTowColor(no).IsBipartite(), false)
}

func TestNewDigraph(t *testing.T) {
	var dig SimpleDigraph = NewDigraph(10)
	dig.AddEdge(0, 3)
	dig.AddEdge(1, 5)
	dig.AddEdge(0, 9)
	t.Log(dig)
	t.Log("-----------------")
	t.Log(dig.Reverse())
}

func TestDirectedSearchDFS_PathTo(t *testing.T) {
	var dig SimpleDigraph = NewDigraph(5)
	dig.AddEdge(0, 1)
	dig.AddEdge(0, 2)
	dig.AddEdge(0, 3)
	dig.AddEdge(1, 2)
	dig.AddEdge(2, 0)
	dig.AddEdge(3, 4)
	t.Log(dig)
	s := new(DirectedSearchDFS).GenSearch(dig, 0,1 )
	r := s.PathTo(4)
	t.Log(r)

}
