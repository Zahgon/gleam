package ui

const (
	m = 10
)
const (
	linesize    = 4
	labelfs     = 16
	notchsize   = 8
	groupmargin = 10

	lcolor      = "rgb(190,190,190)"
	boxradius   = 10
	lopacity    = "1.0"
	defcolor    = "black"
	linefmt     = "stroke:%s;fill:none"
	globalstyle = "font-family:Calibri,sans-serif;font-size:%dpx;fill:black;text-anchor:middle;stroke-linecap:round;stroke-width:%dpx;stroke-opacity:%s"
	ltstyle     = "text-anchor:%s;fill:black"
	legendstyle = "text-anchor:start;fill:black;font-size:%dpx"
	gridstyle   = "fill:none; stroke:gray; stroke-opacity:0.3"
	notefmt     = "font-size:%dpx"
	ntfmt       = "text-anchor:%s;fill:%s"
)

var ()

type point struct {
	x int
	y int
}

func connect(canvas *svg.SVG, a, b point, label string) { _ = "STUB: not implemented"; return }

// message object
func message(canvas *svg.SVG, x, y, w, h, l int, bcolor, scolor string) {
	_ = "STUB: not implemented"
	return
}

// linelabel determines the connection and arrow geometry
func linelabel(canvas *svg.SVG, x1, y1, x2, y2 int, label string, mark string, d1 string, d2 string, dir string, color string) {
	_ = "STUB: not implemented"
	return
}

// linestyle returns the style for lines
func linestyle(color string) string { _ = "STUB: not implemented"; return "" }

// arrow constructs line-ending arrows according to connecting points
func arrow(canvas *svg.SVG, x, y, w, h int, dir string, color string) (xl, yl int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// doline draws a line between to coordinates
func doline(canvas *svg.SVG, x1, y1, x2, y2 int, style, direction, label string) {
	_ = "STUB: not implemented"
	return
}

// upwards line

// horizontal line

// downwards line

// initial control points
// fmt.Fprintf(os.Stderr, "%s slope = %.3f\n", label, m)

// midpoint

// sloper computes the slope and r of a line
func sloper(x1, y1, x2, y2 int) (m, r float64) { _ = "STUB: not implemented"; return 0, 0 }
