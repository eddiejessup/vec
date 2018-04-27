package vec

import (
    "math"
    "github.com/golang/geo/r2"
)

type Circle struct {
    Centre r2.Point
    R      float64
}

func (c *Circle) Offset(p r2.Point) *Circle {
    return &Circle{Centre: c.Centre.Add(p), R: c.R}
}

func (c *Circle) Area() float64 {
    return math.Pi * c.R * c.R
}

func SquaresIntersect(sep r2.Point, dA, dB float64) bool {
    return (math.Abs(sep.X) < dA + dB) && (math.Abs(sep.Y) < dA + dB)
}

func WrappedSquaresIntersect(rA, rB r2.Point, dA, dB float64, w r2.Point) bool {
    sep := ClosestWrappedDistance(rA.Sub(rB), w)
    return (math.Abs(sep.X) < dA + dB) && (math.Abs(sep.Y) < dA + dB)
}

func CirclesIntersect(sep r2.Point, r1, r2 float64) bool {
    return sep.Dot(sep) < (r1 + r2) * (r1 + r2)
}

func (c *Circle) Intersects(cO *Circle) bool {
    return CirclesIntersect(cO.Centre.Sub(c.Centre), c.R, cO.R)
}

func (c *Circle) WrappedIntersects(cO *Circle, w r2.Point) bool {
    sep := ClosestWrappedDistance(cO.Centre.Sub(c.Centre), w)
    return CirclesIntersect(sep, c.R, cO.R)
}
