package game

import (
	"fmt"
	"io"
	"os"
)

const N = 10
const Black = '1'
const White = '2'

var table [N][N]byte

type Man struct {
	r io.Reader
	w io.Writer
	c byte
}

func NewMan(r io.Reader, w io.Writer, c byte) *Man {
	return &Man{r, w, c}
}
func (m Man) Prepare() {
	fmt.Fprintf(m.w, "%c 游戏开始!\r\n", m.c)
}
func Init() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			table[i][j] = '_'
		}
	}
}

func Prepare(A, B *Man) {
	Init()
	A.Prepare()
	B.Prepare()
}

func Do(m *Man, x, y int) (winned bool) {
	if x >= N || y >= N || table[x][y] != '_' {
		fmt.Fprint(m.w, "那里不能下!\r\n")
		panic("下错位置!")
	}
	table[x][y] = m.c
	return Judge(m.c, x, y)
}

func Round(d, w *Man) (winned bool) {
	Show(d.w)
	var x, y int
	fmt.Fscanf(os.Stdin, "%d,%d\r\n", &x, &y)
	// fmt.Fprintln(d.w, x, y)
	return Do(d, x, y)
}

func Win(d, w *Man) {
	fmt.Fprintf(d.w, "%c Win\r\n", d.c)
	fmt.Fprintf(w.w, "%c Win\r\n", d.c)
}

func Show(w io.Writer) {
	for i := 0; i < N; i++ {
		fmt.Fprintf(w, " %d", i)
	}

	for lno, ln := range table {
		fmt.Fprintf(w, "\n%d", lno)
		for _, s := range ln {
			fmt.Fprintf(w, " %c", s)
		}
	}
	fmt.Fprintln(w)
}
func Judge(c byte, x, y int) (winned bool) {
	directions := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, d := range directions {
		t, x1, y1 := 0, x, y
		for i := 0; i < 4; i++ {
			x1, y1 = x1+d[0], y1+d[1]
			if x1 < 0 || y1 < 0 || x1 >= N || y1 >= N || table[x1][y1] != c {
				break
			}
			t++
		}
		if t == 4 {
			return true
		}
	}
	return false
}
