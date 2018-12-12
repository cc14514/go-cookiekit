package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	g     SimpleGraph
	dig   SimpleDigraph
	dag   SimpleDigraph
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

	dag = NewDigraph(13)
	dag.AddEdge(0, 1)
	dag.AddEdge(0, 5)
	dag.AddEdge(0, 6)
	dag.AddEdge(2, 0)
	dag.AddEdge(2, 3)
	dag.AddEdge(3, 5)
	dag.AddEdge(5, 4)
	dag.AddEdge(6, 4)
	dag.AddEdge(6, 9)
	dag.AddEdge(7, 6)
	dag.AddEdge(8, 7)
	dag.AddEdge(9, 10)
	dag.AddEdge(9, 11)
	dag.AddEdge(9, 12)
	dag.AddEdge(11, 12)

	dig = NewDigraph(5)
	dig.AddEdge(0, 1)
	dig.AddEdge(0, 2)
	dig.AddEdge(0, 3)
	dig.AddEdge(1, 2)
	dig.AddEdge(2, 0)
	dig.AddEdge(3, 4)

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

func TestCCImpl(t *testing.T) {
	g = NewGraph(7)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	cc := NewCC(g)
	t.Log(cc.Count())
	t.Log(cc.ID(0), cc.ID(1), cc.ID(2))
	t.Log(cc.ID(3), cc.ID(4))
	t.Log(cc.ID(5), cc.ID(6))
	t.Log(cc.Connected(3,6))
	t.Log(cc.Connected(1,2))
}

func TestCycle(t *testing.T) {
	c := NewCycle(g)
	t.Log(c.HasCycle())
	t.Log(c.Cycle())
}

func TestDirectedCycle(t *testing.T) {
	dig = NewDigraph(5)
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
	dig = NewDigraph(10)
	dig.AddEdge(0, 3)
	dig.AddEdge(1, 5)
	dig.AddEdge(0, 9)
	t.Log(dig)
	t.Log("-----------------")
	t.Log(dig.Reverse())
}

func TestDirectedSearchDFS_PathTo(t *testing.T) {

	s := new(DirectedSearchDFS).GenSearch(dig, 0)
	r := s.PathTo(2)
	t.Log(r)
}

func TestDirectedSearchBFS_PathTo(t *testing.T) {
	s := new(DirectedSearchBFS).GenSearch(dig, 0)
	r := s.PathTo(2)
	t.Log(r)
}

func TestNewDFDigOrder(t *testing.T) {
	t.Log(dag)
	o := NewDFOrder(dag)
	t.Log(o.Per())
	t.Log(o.Post())
	t.Log(o.ReversePost())
}

func TestNewDigTopological(t *testing.T) {
	t.Log(dag)
	tl := NewDigTopological(dag)
	t.Log(tl.IsDAG())
	t.Log(tl.Order())
}
