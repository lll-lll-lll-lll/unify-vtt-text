package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestTextSegment(t *testing.T) {
	token := "you need to know it, ♪ I know ♪"
	filename := "testvtt.en-ehkg1hFWq8A.vtt"
	t.Run("get 「.」and 「?」", func(t *testing.T) {
		got := CheckTerminalFlag(token)
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
		t.Log(got)
	})
	t.Run("", func(t *testing.T) {
		wantpr := token[:19] // you need to know it
		wantba := token[20:] // ♪ I know ♪
		wantterminal := "?"
		locs := SearchTerminalTokenRegexp(token)
		fmt.Print(locs)
		_, pr, ba, ter, flag := SplitByCommaAndQuestion(token)
		if flag {
			t.Log("no word in text. `?` or `.`")
		} else {
			if pr != wantpr {
				t.Errorf("got %s want %s", pr, wantpr)
			}
			if ba != wantba {
				t.Errorf("got %s want %s", ba, wantba)
			}

			if ter != wantterminal {
				t.Errorf("got %s want %s", "?", wantterminal)
			}
		}
	})

	t.Run("Unify Text by comma or question", func(t *testing.T) {
		f, err := ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		webVtt := NewWebVtt(f)
		webVtt.ScanLines(ScanTimeLineSplitFunc)
		w := UnifyTextByTerminalPoint(webVtt)
		got := w.VttElements[1].Text
		want := "- Yo what is going on guys,welcome back to the channel."
		fmt.Println(got)
		fmt.Println(want)
		PrintlnJson(w.VttElements)
	})
}

func PrintlnJson(elements []*VTTElement) {
	for _, e := range elements {
		var out bytes.Buffer
		b, _ := json.Marshal(e)
		err := json.Indent(&out, b, "", "  ")
		if err != nil {
			panic(err)

		}
		fmt.Println(out.String())
	}
}
