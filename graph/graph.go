package graph

import (
	"strings"
	"strconv"
	"bytes"
	"github.com/cc14514/go-cookiekit/collections/bag"
)

// 无向图
type Graph struct {
	v, e int
	adj  []*bag.Bag //邻接表
}

func (self *Graph) GetAdj() []*bag.Bag {
	return self.adj
}

func (self *Graph) V() int {
	return self.v
}

func (self *Graph) E() int {
	return self.e
}

func (self *Graph) AddEdge(v, w int) {
	if v > self.V() || w > self.V() {
		panic("error number")
	}
	if self.adj[v] == nil {
		self.adj[v] = bag.New()
	}
	if self.adj[v].Count(w) < 1 {
		self.adj[v].Insert(w)
		if v != w { // 自环
			if self.adj[w] == nil {
				self.adj[w] = bag.New()
			}
			self.adj[w].Insert(v)
		}
		self.e ++
	}
}

func (self *Graph) Adj(v int) []int {
	if v >= len(self.adj) {
		return nil
	}
	bag := self.adj[v]
	r := make([]int, 0)
	bag.Items(func(i interface{}) {
		r = append(r, i.(int))
	})
	return r
}

func (self *Graph) String() string {
	var buf bytes.Buffer
	buf.WriteString("\n")
	buf.WriteString(strconv.Itoa(int(self.V())))
	buf.WriteString("\n")
	buf.WriteString(strconv.Itoa(int(self.E())))
	buf.WriteString("\n")
	filter := make(map[struct{ a, b interface{} }]bool)
	for i, n := range self.adj {
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

func NewGraph(v int) (g *Graph) {
	g = new(Graph)
	g.v = v
	g.adj = make([]*bag.Bag, v, v)
	return
}

/*
-------------
 data format
-------------
p1 p2
p2 p3
p3 p4
......
*/
func NewGraphByData(data string) (g *Graph) {
	if data == "" {
		return
	}
	g = new(Graph)
	da := strings.Split(data, "\n")
	el := make([]int, 0)
	for _, d := range da {
		if d != "" {
			d = strings.Trim(d, " ")
			if dda := strings.Split(d, " "); len(dda) == 2 {
				_v, _ := strconv.ParseInt(dda[0], 10, 32)
				_w, _ := strconv.ParseInt(dda[1], 10, 32)
				el = append(el, int(_v), int(_w))
			}
		}
	}
	b := bag.New()
	for _, i := range el {
		b.Insert(i)
	}
	g.v = b.ItemSize()
	g.adj = make([]*bag.Bag, g.v)
	for i := 0; i < len(el); i += 2 {
		g.AddEdge(el[i], el[i+1])
	}
	return
}

/*
-------------
 data format
-------------
p1: a,b,c,d
p2: e,f,g
p3: h,i,j,k,l
......
*/
func NewGraphByAdjacencyList(data string) (g *Graph) {
	if data == "" {
		return
	}
	g = new(Graph)
	da := strings.Split(data, "\n")
	sl := make([]string, 0)
	for _, d := range da {
		if d != "" {
			sl = append(sl, d)
		}
	}
	if len(sl) <= 0 {
		return nil
	}
	g.v = len(sl)
	g.adj = make([]*bag.Bag, g.v)
	for _, s := range sl {
		ss := strings.Split(s, ":")
		if len(ss) < 2 {
			continue
		}
		ee := strings.Split(ss[1], " ")
		if len(ee) < 1 {
			continue
		}
		v, _ := strconv.ParseInt(strings.Trim(ss[0], " "), 10, 32)
		for _, e := range ee {
			w, _ := strconv.ParseInt(strings.Trim(e, " "), 10, 32)
			g.AddEdge(int(v), int(w))
		}
	}
	return
}
