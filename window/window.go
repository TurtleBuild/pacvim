package window

import (
	"strconv"

	"pacvim/buffer"

	termbox "github.com/nsf/termbox-go"
)

type Window struct {
	lines []*buffer.Line
}

func New(b *buffer.Buffer) *Window {
	w := new(Window)
	w.copyBufToWindow(b)
	return w
}

func (w *Window) copyBufToWindow(b *buffer.Buffer) {
	w.lines = []*buffer.Line{}
	winWidth, winHeight := termbox.Size()
	for i := 0; i < len(b.Lines); i++ {
		if i > winHeight-1 {
			break
		}
		w.lines = append(w.lines, &buffer.Line{})
		for j := 0; j < len(b.Lines[i].Text); j++ {
			if j+getDigit(b.NumOfLines())+1 > winWidth-1 {
				break
			}
			w.lines[i].Text = append(w.lines[i].Text, b.Lines[i].Text[j])
		}
	}
}

// 行番号なしで表示
func (w *Window) Show(b *buffer.Buffer) error {
	err := termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	for y, l := range w.lines {
		for x, r := range l.Text {
			termbox.SetCell(x, y, r, termbox.ColorYellow, termbox.ColorBlack)
		}
	}
	return err
}

// 行番号ありで表示
func (w *Window) ShowWithLineNum(b *buffer.Buffer) error {
	err := termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	offset := getDigit(b.NumOfLines())
	b.Offset = offset + 1
	for y, l := range w.lines {
		linenums := makeLineNum(y+1, offset)
		t := append(linenums, l.Text...)
		for x, r := range t {
			termbox.SetCell(x, y, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
	return err
}
func makeLineNum(num int, digit int) []rune {
	numstr := strconv.Itoa(num)
	lineNum := make([]rune, digit+1)
	for i := 0; i < len(lineNum); i++ {
		lineNum[i] = ' '
	}
	cdigit := getDigit(num)
	for i, c := range numstr {
		lineNum[i+(digit-cdigit)] = c
	}
	return lineNum
}
func getDigit(linenum int) int {
	d := 0
	for linenum != 0 {
		linenum = linenum / 10
		d++
	}
	return d
}
