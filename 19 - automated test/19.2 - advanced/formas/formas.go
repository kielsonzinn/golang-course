package formas

import (
	"math"
)

// https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/learn/lecture/22117190?start=179#notes
type Forma interface {
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
