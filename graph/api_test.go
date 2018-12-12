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
4 5
5 3
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
	t.Log(c.Cycle())
}

func TestDirectedCycle(t *testing.T) {
	dig := NewDigraph(5)
	dig.AddEdge(0, 1)
	dig.AddEdge(1, 2)
	dig.AddEdge(3, 4)
	dig.AddEdge(4, 2)

	c1 := NewDirectedCycle(dig)
	assert.Equal(t, c1.HasCycle(), false)
	t.Log(c1.Cycle())

	dig.AddEdge(2, 0)
	c2 := NewDirectedCycle(dig)
	assert.Equal(t, c2.HasCycle(), true)
	t.Log(c2.Cycle())
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
	s := new(DirectedSearchDFS).GenSearch(dig, 0)
	r := s.PathTo(2)
	t.Log(r)
}

func TestDirectedSearchBFS_PathTo(t *testing.T) {
	var dig SimpleDigraph = NewDigraph(5)
	dig.AddEdge(0, 1)
	dig.AddEdge(0, 2)
	dig.AddEdge(0, 3)
	dig.AddEdge(1, 2)
	dig.AddEdge(2, 0)
	dig.AddEdge(3, 4)
	s := new(DirectedSearchBFS).GenSearch(dig, 0)
	r := s.PathTo(2)
	t.Log(r)
}

func TestNewDFDigOrder(t *testing.T) {
	var dig SimpleDigraph = NewDigraph(13)
	dig.AddEdge(0, 1)
	dig.AddEdge(0, 5)
	dig.AddEdge(0, 6)

	dig.AddEdge(2, 0)
	dig.AddEdge(2, 3)

	dig.AddEdge(3, 5)

	dig.AddEdge(5, 4)

	dig.AddEdge(6, 4)
	dig.AddEdge(6, 9)

	dig.AddEdge(7, 6)

	dig.AddEdge(8, 7)

	dig.AddEdge(9, 10)
	dig.AddEdge(9, 11)
	dig.AddEdge(9, 12)

	dig.AddEdge(11, 12)

	t.Log(dig)

	o := NewDFDigOrder(dig)
	t.Log(o.Per())
	t.Log(o.Post())
	t.Log(o.ReversePost())
}