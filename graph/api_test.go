package graph

import (
	"testing"
)

var (
	g     SimpleGraph
	data1 = `
0 1
1 2
2 0
3 4
`
	data2 = `
0:1,5
1:0,5,3,2
2:1,3
3:1,2,4,5
4:3,5
5:0,1,3,4
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

func TestSearch(t *testing.T) {
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
