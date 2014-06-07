package graph

import (
	"fmt"
	"io"
)

type Edge interface{}

type Vertex interface {
	GetLabel() string
}

type Graph interface {
	Edges() chan (Edge)
	GetVertices(e Edge) (Vertex, Vertex)
}

func WriteDot(g Graph, w io.Writer) error {
	_, err := io.WriteString(w, "graph G {\n")
	if err != nil {
		return err
	}

	i := 0
	for e := range g.Edges() {
		u, v := g.GetVertices(e)
		line := fmt.Sprintf("\t\"%s\" -- \"%s\" // e=%d\n", u.GetLabel(), v.GetLabel(), i)
		_, err := io.WriteString(w, line)
		if err != nil {
			return err
		}
		i++
	}

	_, err = io.WriteString(w, "}\n")
	if err != nil {
		return err
	}

	return nil
}

func WriteDiDot(g Graph, w io.Writer) error {
	_, err := io.WriteString(w, "digraph G {\n")
	if err != nil {
		return err
	}

	i := 0
	for e := range g.Edges() {
		u, v := g.GetVertices(e)
		line := fmt.Sprintf("\t\"%s\" -> \"%s\" // e=%d\n", u.GetLabel(), v.GetLabel(), i)
		_, err := io.WriteString(w, line)
		if err != nil {
			return err
		}
		i++
	}

	_, err = io.WriteString(w, "}\n")
	if err != nil {
		return err
	}

	return nil
}
