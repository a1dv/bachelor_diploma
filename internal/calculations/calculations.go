package calculations

import (
  "fmt"
  "math"
  "gonum.org/v1/gonum/diff/fd"
  "gonum.org/v1/plot/plotter"
  "gonum.org/v1/plot"
  "image/color"
)

func Calculations(m, L, w, a float64) {

    fmt.Println(m, L, w, a)
    T1 := 0.05 * 0.707 / 3

    Hg := math.Sqrt(2) / T1

    Hg1 := 1 / (T1 * T1)

    f := func(x float64) float64{
        return 0.5 * math.Sin(x * 10) + 1.5
    }
    df := fd.Derivative(f, 0, &fd.Settings{
        Formula: fd.Forward,
        Step:    1e-10,
    })

    //F1 := 1
    //z := 1
  //  F12 := 0
    //F2 := 1
  //  F22 := 0

    secf := func(t float64) float64{
        return 0.15 * w * math.Cos(t * w)
    }

    secdf := fd.Derivative(secf, 1. / 10., &fd.Settings{
        Formula: fd.Forward,
        Step:    1e-10,
    })

    J := m * a * a / 4
    fmt.Println(J, L)

    fmt.Println("f'(0) ≈", df)
    fmt.Println("2f'(0) ≈", secdf)

    t0 := 0.
    t1 := 1.5
    Y0 := []float64{0, 0, 0, 0.5, 0, 0}
    N := 1000.

    Rkadapt(Y0, t0, t1, N, w, Hg, Hg1, m, L, T1,a, J)

    fmt.Println(T1)
    fmt.Println(Hg)
    fmt.Println(Hg1)
}


func Rkadapt(Y0 []float64, a, b, n, w, Hg, Hg1, m float64, L, T1, A, J float64) {
    pts := plotter.XYs{}
    pts1 := plotter.XYs{}
    pts2 := plotter.XYs{}
    pts3 := plotter.XYs{}
    pts4 := plotter.XYs{}

    p := plot.New()
    p1 := plot.New()
    p2 := plot.New()

    Y := make([]float64, len(Y0))
    h := math.Abs(b - a)/n

    for i := 0.; i < n; i++ {

        if !math.IsNaN(Y[0]) {

            pts = append(pts, plotter.XY{a + h * float64(i), Y[1]})
            pts1 = append(pts1, plotter.XY{a + h * float64(i), Y[3]})
            pts2 = append(pts2, plotter.XY{a + h * float64(i), Y[5]})
            pts3 = append(pts3, plotter.XY{a + h * float64(i), Y[0]})
            pts4 = append(pts4, plotter.XY{a + h * float64(i), Y[4]})
        }

        Y = D(a + h*float64(i), w, A, Hg, Hg1, m, J, Y0)
    }

    line, err := plotter.NewLine(pts)
    if err != nil {
        panic(err)
    }

    line.LineStyle.Color = color.RGBA{uint8(255), uint8(0), uint8(0), 0}
    line1, err := plotter.NewLine(pts1)
    if err != nil {
        panic(err)
    }

    line2, err := plotter.NewLine(pts2)
    if err != nil {
        panic(err)
    }

    line3, err := plotter.NewLine(pts3)
    if err != nil {
        panic(err)
    }

    line4, err := plotter.NewLine(pts4)
    if err != nil {
        panic(err)
    }

    scatter, err := plotter.NewScatter(pts)
    if err != nil {
        panic(err)
    }

    scatter.GlyphStyle.Color = color.RGBA{uint8(255), uint8(0), uint8(0), 0}
    scatter2, err := plotter.NewScatter(pts2)
    if err != nil {
        panic(err)
    }

    scatter2.GlyphStyle.Color = color.RGBA{uint8(255), uint8(0), uint8(255), 0}
    scatter1, err := plotter.NewScatter(pts1)
    if err != nil {
        panic(err)
    }

    scatter3, err := plotter.NewScatter(pts3)
    scatter1.GlyphStyle.Color = color.RGBA{uint8(0), uint8(255), uint8(0), 0}
    if err != nil {
        panic(err)
    }

    scatter4, err := plotter.NewScatter(pts4)
    if err != nil {
        panic(err)
    }

    p.Add(line, scatter)
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"
    p.Add(line2, scatter2)
    p.Add(line1, scatter1)
    p1.Add(line3, scatter3)
    p2.Add(line4, scatter4)

    if err := p.Save(1600, 1600, "assets/function.png"); err != nil {
        panic(err)
    }

    if err := p1.Save(1600, 1600, "assets/x.png"); err != nil {
        panic(err)
    }

    if err := p2.Save(1600, 1600, "assets/angle.png"); err != nil {
        panic(err)
    }
    fmt.Println(Y)
}


func D(t, w, a, Hg, Hg1, m, J float64, Y []float64) []float64{
    Y01 := 0.5 * math.Sin(t * w) + 1.5
    X01 := 0.
    phi01 := 0.

    partsin := a * math.Sin(Y[5]) / 2
    partcos := a * math.Cos(Y[5]) / 2
    Y[1] = 0.5 * w * math.Cos(t * w)
    Y[2] = 0.5 * w * math.Cos(t * w)
    Y[3] = 0.5 * math.Sin(t * w)
    Y[5] = 0

    q1 := math.Sqrt(math.Pow(Y[3] - partsin, 2.) + math.Pow(2 * a / 3 + Y[1] - partcos, 2.))
    q2 := math.Pow(Y[3] + partsin, 2.) + math.Pow(Y[1] - 2 * a / 3 + partcos, 2.)
    q3 := math.Sqrt(math.Pow(Y[3] + partsin, 2.) + math.Pow(Y[1] - a + partcos, 2.))

    A1 := (4 * a / 3 + 2 * Y[1] - a * math.Cos(Y[5])) / (2 * q1)
    B1 := (2 * Y[1] - 4 * a / 3 + a * math.Cos(Y[5])) / (2 * q1)
    C1 := (2 * Y[1] - 2 * a + a * math.Cos(Y[5])) / (2 * q1)
    A2 := (2 * Y[3] - a * math.Sin(Y[5])) / (2 * math.Sqrt(q2))
    B2 := (2 * Y[3] + a * math.Sin(Y[5])) / (2 * math.Sqrt(q2))
    C2 := (2 * Y[3] + a * math.Sin(Y[5])) / (2 * math.Sqrt(q2))
    A3 := (a * math.Sin(Y[5]) * (2 * a / 3 + Y[1] - partcos) - a * math.Cos(Y[5]) * (Y[3] - partsin)) / (2 * q3)
    B3 := (a * math.Sin(Y[5]) * (Y[1] - 2 * a / 3 + partcos) - a * math.Cos(Y[5]) * (Y[3] + partsin)) / (2 * q3)
    C3 := (a * math.Sin(Y[5]) * (Y[1] - a + partcos) - a * math.Cos(Y[5]) * (Y[3] + partsin)) / (2 * q3)
    X1 := 0.
    phi1 := 0.
    X11 := 0.
    Y11 := -0.5 * w * w * math.Sin(t * w)
    phi11 := 0.
    fmt.Println(A1, B1, C1, A2, B2, C2, A3, B3, C3)
    firstbot := A1*B2*C3 - A1*B3*C2 - A2*B1*C3 + A2*B3*C1 + A3*B1*C2 - A3*B2*C1
    A11 := (B2 * C3 - B3 * C2) / firstbot
    B11 := - (B1 * C3 - B3 * C1) / firstbot
    C11 := (B1 * C2 - B2 * C1) / firstbot
    A22 := - (A2 * C3 - A3 * C2) / firstbot
    B22 := (A1 * C3 - A3 * C1) / firstbot
    C22 := - (A1 * C2 - A2 * C1) / firstbot
    A33 := (A2 * B3 - A3 * B2) / firstbot
    B33 := - (A1 * B3 - A3 * B1) / firstbot

    C33 := (A1 * B2 - A2 * B1) / firstbot

// закон управления:
    x20 := X11 - Hg * (Y[0] - X1) - Hg1 * (Y[1] - X01)
    y20 := Y11 - Hg * (Y[2] - Y[1]) - Hg1 * (Y[3] - Y01)
    phi20 := phi11 - Hg * (Y[4] - phi1) - Hg1 * (Y[5] * phi01)

    P1 := m * x20 * A11 + m * y20 * A22 + J * phi20 * A33
    P2 := m * x20 * B11 + m * y20 * B22 + J * phi20 * B33
    P3 := m * x20 * C11 + m * y20 * C22 + J * phi20 * C33
    fmt.Printf("\nP1=%f A1 = %f P2 = %f B1 = %f P3 = %f C1 = %f\n", P1, A1, P2, B1, P3, C1)
    fmt.Println("M ", partsin)
    ax := (P1 * A1 + P2 * B1 + P3 * C1) / m

    return []float64{ax, 0, (P1 * A2 + P2 * B2 + P3 * C2) / m, Y[1], P1 * A3 + P2 * B3 + P3 * C3, 0}
}
