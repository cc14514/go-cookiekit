package graph

import (
	"github.com/cc14514/go-cookiekit/collections/bag"
	"bytes"
	"strconv"
)

// 有向图
type Digraph struct {
	v, e int
	adj  []*bag.Bag //邻接表
}

func (self *Digraph) V() int {
	return self.v
}

func (self *Digraph) E() int {
	return self.e
}

func (self *Digraph) AddEdge(v, w int) {
	if v > self.V() || w > self.V() {
		panic("error number")
	}
	if self.adj[v] == nil {
		self.adj[v] = bag.New()
	}
	if self.adj[v].Count(w) < 1 {
		self.adj[v].Insert(w)
		self.e ++
	}
}

func (self *Digraph) GetAdj() []*bag.Bag {
	return self.adj
}

func (self *Digraph) Adj(v int) []int {
	if v >= len(self.adj) {
		return nil
	}
	bag := self.adj[v]
	if bag == nil {
		return nil
	}
	r := make([]int, 0)
	bag.Items(func(i interface{}) {
		r = append(r, i.(int))
	})
	return r
}

func (self *Digraph) String() string {
	var buf bytes.Buffer
	buf.WriteString("\n")
	buf.WriteString(strconv.Itoa(int(self.V())))
	buf.WriteString("\n")
	buf.WriteString(strconv.Itoa(int(self.E())))
	buf.WriteString("\n")
	filter := make(map[struct{ a, b interface{} }]bool)
	for i, n := range self.adj {
		if n == nil {
			continue
		}
		n.Items(func(val interface{}) {
			if !filter[struct{ a, b interface{} }{i, val}] {
				filter[struct{ a, b interface{} }{val, i}] = true
				buf.WriteString(strconv.Itoa(i))
				buf.WriteString(" ")
				buf.WriteString(strconv.Itoa(val.(int)))
				buf.WriteString("\n")
			}
		})
	}
	return buf.String()
}

func (self *Digraph) Reverse() SimpleDigraph {
	dig := NewDigraph(self.V())
	for v := 0; v < self.V(); v++ {
		if self.Adj(v) == nil {
			continue
		}
		for _, w := range self.Adj(v) {
			dig.AddEdge(w, v)
		}
	}
	return dig
}

func NewDigraph(v int) (g *Digraph) {
	g = new(Digraph)
	g.v = v
	g.adj = make([]*bag.Bag, v, v)
	return
}
