package vec

import (
    "math"
    "math/rand"
    "github.com/golang/geo/r2"
)


func RandomUnitVector() *r2.Point {
    ang := rand.Float64() * 2 * math.Pi
    return &r2.Point{X: math.Cos(ang), Y: math.Sin(ang)}
}

func PointAngle(p r2.Point) float64 {
    return math.Atan2(p.Y, p.X)
}

func RotateVector2d(v r2.Point, ang float64) r2.Point {
    sinAng, cosAng := math.Sincos(ang)
    return r2.Point{
        X: v.X * cosAng - v.Y * sinAng,
        Y: v.X * sinAng + v.Y * cosAng,
    }
}

func WrapPoint(p *r2.Point, w r2.Point) {
    // Wrap.
    if p.X > w.X {
        p.X -= w.X
    } else if p.X < 0 {
        p.X += w.X
    }
    if p.Y > w.Y {
        p.Y -= w.Y
    } else if p.Y < 0 {
        p.Y += w.Y
    }    
}


func closeWrapDist(x, w float64) float64 {
    if 2 * x > w {
        return x - w
    } else if 2 * x < -w {
        return x + w
    }
    return x
}

func ClosestWrappedDistance(d, w r2.Point) r2.Point {
    d.X = closeWrapDist(d.X, w.X)
    d.Y = closeWrapDist(d.Y, w.Y)
    return d
}

func SubWrap(pTo, pFrom, w r2.Point) r2.Point {
    return ClosestWrappedDistance(pTo.Sub(pFrom), w)
}

func MoveWrappedPoint(p *r2.Point, d, w r2.Point) {
    *p = p.Add(d)
    WrapPoint(p, w)
}

func RandSymmetricFloat64(scale float64) float64 {
    return (rand.Float64() - 0.5) * 2 * scale
}

func UnitPointFromAngle(a float64) r2.Point {
    return r2.Point{X: math.Cos(a), Y: math.Sin(a)}
}

func Round(f float64) float64 {
    return math.Floor(f + .5)
}
