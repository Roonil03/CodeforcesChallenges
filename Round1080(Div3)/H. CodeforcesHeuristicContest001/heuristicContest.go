package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(line)
}

type Point struct {
	x, y int
}

type Triangle struct {
	p1, p2, p3 int
}

var base5x3 []int
var points5x3 []Point

func crossProduct(a, b, c Point) int {
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
}

func intersect(a, b, c, d Point) bool {
	cp1 := crossProduct(c, d, a)
	cp2 := crossProduct(c, d, b)
	cp3 := crossProduct(a, b, c)
	cp4 := crossProduct(a, b, d)

	if ((cp1 > 0 && cp2 < 0) || (cp1 < 0 && cp2 > 0)) &&
		((cp3 > 0 && cp4 < 0) || (cp3 < 0 && cp4 > 0)) {
		return true
	}
	return false
}

func trianglesIntersect(t1, t2 Triangle) bool {
	pts1 := []Point{points5x3[t1.p1], points5x3[t1.p2], points5x3[t1.p3]}
	pts2 := []Point{points5x3[t2.p1], points5x3[t2.p2], points5x3[t2.p3]}

	for i := 0; i < 3; i++ {
		u, v := pts1[i], pts1[(i+1)%3]
		for j := 0; j < 3; j++ {
			w, z := pts2[j], pts2[(j+1)%3]
			if intersect(u, v, w, z) {
				return true
			}
		}
	}
	return false
}

func init() {
	points5x3 = make([]Point, 15)
	for y := 0; y < 3; y++ {
		for x := 0; x < 5; x++ {
			points5x3[y*5+x] = Point{x, y}
		}
	}
	used := 0
	currentSol := make([]Triangle, 0, 5)
	var solve func(idx int) bool
	solve = func(idx int) bool {
		firstUnused := -1
		for i := 0; i < 15; i++ {
			if (used>>i)&1 == 0 {
				firstUnused = i
				break
			}
		}
		if firstUnused == -1 {
			return len(currentSol) == 5
		}
		p1 := firstUnused
		used |= (1 << p1)
		for p2 := p1 + 1; p2 < 15; p2++ {
			if (used>>p2)&1 != 0 {
				continue
			}
			used |= (1 << p2)
			for p3 := p2 + 1; p3 < 15; p3++ {
				if (used>>p3)&1 != 0 {
					continue
				}
				cp := crossProduct(points5x3[p1], points5x3[p2], points5x3[p3])
				if cp != 1 && cp != -1 {
					continue
				}
				newTri := Triangle{p1, p2, p3}
				ok := true
				for _, existing := range currentSol {
					if trianglesIntersect(newTri, existing) {
						ok = false
						break
					}
				}
				if ok {
					used |= (1 << p3)
					currentSol = append(currentSol, newTri)
					if solve(idx + 1) {
						return true
					}
					currentSol = currentSol[:len(currentSol)-1]
					used &^= (1 << p3)
				}
			}
			used &^= (1 << p2)
		}
		used &^= (1 << p1)
		return false
	}
	if !solve(0) {
		base5x3 = []int{
			0, 0, 1, 0, 0, 1,
			0, 2, 1, 1, 1, 2,
			2, 0, 2, 1, 3, 1,
			3, 0, 4, 0, 4, 1,
			2, 2, 3, 2, 4, 2,
		}
	} else {
		base5x3 = make([]int, 0, 30)
		for _, t := range currentSol {
			p1, p2, p3 := points5x3[t.p1], points5x3[t.p2], points5x3[t.p3]
			base5x3 = append(base5x3, p1.x, p1.y, p2.x, p2.y, p3.x, p3.y)
		}
	}
}

func printTri(x1, y1, x2, y2, x3, y3 int) {
	fmt.Fprintf(out, "%d %d %d %d %d %d\n", x1, y1, x2, y2, x3, y3)
}

func fillRect(x, y, w, h int) {
	if w%2 == 0 && h%3 == 0 {
		for i := 0; i < w; i += 2 {
			for j := 0; j < h; j += 3 {
				printTri(x+i+1, y+j+1, x+i+1, y+j+2, x+i+2, y+j+1)
				printTri(x+i+2, y+j+2, x+i+2, y+j+3, x+i+1, y+j+3)
			}
		}
	} else if w%3 == 0 && h%2 == 0 {
		for i := 0; i < w; i += 3 {
			for j := 0; j < h; j += 2 {
				printTri(x+i+1, y+j+1, x+i+2, y+j+1, x+i+1, y+j+2)
				printTri(x+i+3, y+j+1, x+i+3, y+j+2, x+i+2, y+j+2)
			}
		}
	}
}

func solveOdd(n, offX, offY int) {
	if n == 3 {
		for strip := 0; strip < 3; strip++ {
			sy := offY + strip*3
			fillRect(offX, sy, 4, 3)
			baseX := offX + 4
			baseY := sy
			for k := 0; k < len(base5x3); k += 6 {
				printTri(
					baseX+base5x3[k]+1, baseY+base5x3[k+1]+1,
					baseX+base5x3[k+2]+1, baseY+base5x3[k+3]+1,
					baseX+base5x3[k+4]+1, baseY+base5x3[k+5]+1,
				)
			}
		}
		return
	}
	fillRect(offX, offY+3*n-6, 3*n, 6)
	fillRect(offX, offY, 6, 3*(n-2))
	solveOdd(n-2, offX+6, offY)
}

func solve() {
	line := readLine()
	if line == "" {
		return
	}
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	if n == 1 {
		fmt.Fprintln(out, 2)
		fmt.Fprintln(out, "1 1 1 2 2 1")
		fmt.Fprintln(out, "2 3 3 2 3 3")
		return
	}
	m := 3 * n * n
	fmt.Fprintln(out, m)
	if n%2 == 0 {
		fillRect(0, 0, 3*n, 3*n)
	} else {
		solveOdd(n, 0, 0)
	}
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	line := readLine()
	if line == "" {
		return
	}
	t, _ := strconv.Atoi(line)
	for t > 0 {
		t--
		solve()
	}
}
