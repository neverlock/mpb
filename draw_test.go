package mpb

import (
	"bytes"
	"testing"
)

func TestFillBar(t *testing.T) {
	// key is termWidth
	testSuite := map[int]map[string]struct {
		total, current int64
		barWidth       int
		barRefill      *refill
		want           string
	}{
		100: {
			"t,c,bw{100,100,0}": {
				total:    100,
				current:  0,
				barWidth: 100,
				want:     "[--------------------------------------------------------------------------------------------------]",
			},
			"t,c,bw{100,1,100}": {
				total:    100,
				current:  1,
				barWidth: 100,
				want:     "[>-------------------------------------------------------------------------------------------------]",
			},
			"t,c,bw{100,40,100}": {
				total:    100,
				current:  40,
				barWidth: 100,
				want:     "[======================================>-----------------------------------------------------------]",
			},
			"t,c,bw{100,40,100}refill{'+', 32}": {
				total:     100,
				current:   40,
				barWidth:  100,
				barRefill: &refill{'+', 32},
				want:      "[+++++++++++++++++++++++++++++++=======>-----------------------------------------------------------]",
			},
			"t,c,bw{100,99,100}": {
				total:    100,
				current:  99,
				barWidth: 100,
				want:     "[================================================================================================>-]",
			},
			"t,c,bw{100,100,100}": {
				total:    100,
				current:  100,
				barWidth: 100,
				want:     "[==================================================================================================]",
			},
		},
		2: {
			"t,c,bw{0,0,100}": {
				barWidth: 100,
				want:     "[]",
			},
			"t,c,bw{60,20,80}": {
				total:    60,
				current:  20,
				barWidth: 80,
				want:     "[]",
			},
		},
		3: {
			"t,c,bw{100,20,100}": {
				total:    100,
				current:  20,
				barWidth: 100,
				want:     "[-]",
			},
			"t,c,bw{100,98,100}": {
				total:    100,
				current:  98,
				barWidth: 100,
				want:     "[=]",
			},
			"t,c,bw{100,100,100}": {
				total:    100,
				current:  100,
				barWidth: 100,
				want:     "[=]",
			},
		},
		5: {
			"t,c,bw{100,20,100}": {
				total:    100,
				current:  20,
				barWidth: 100,
				want:     "[>--]",
			},
			"t,c,bw{100,98,100}": {
				total:    100,
				current:  98,
				barWidth: 100,
				want:     "[===]",
			},
			"t,c,bw{100,100,100}": {
				total:    100,
				current:  100,
				barWidth: 100,
				want:     "[===]",
			},
		},
		6: {
			"t,c,bw{100,20,100}": {
				total:    100,
				current:  20,
				barWidth: 100,
				want:     "[>---]",
			},
			"t,c,bw{100,98,100}": {
				total:    100,
				current:  98,
				barWidth: 100,
				want:     "[====]",
			},
			"t,c,bw{100,100,100}": {
				total:    100,
				current:  100,
				barWidth: 100,
				want:     "[====]",
			},
		},
		20: {
			"t,c,bw{100,20,100}": {
				total:    100,
				current:  20,
				barWidth: 100,
				want:     "[===>--------------]",
			},
			"t,c,bw{100,60,100}": {
				total:    100,
				current:  60,
				barWidth: 100,
				want:     "[==========>-------]",
			},
			"t,c,bw{100,98,100}": {
				total:    100,
				current:  98,
				barWidth: 100,
				want:     "[==================]",
			},
			"t,c,bw{100,100,100}": {
				total:    100,
				current:  100,
				barWidth: 100,
				want:     "[==================]",
			},
		},
		50: {
			"t,c,bw{100,20,100}": {
				total:    100,
				current:  20,
				barWidth: 100,
				want:     "[=========>--------------------------------------]",
			},
			"t,c,bw{100,60,100}": {
				total:    100,
				current:  60,
				barWidth: 100,
				want:     "[============================>-------------------]",
			},
			"t,c,bw{100,98,100}": {
				total:    100,
				current:  98,
				barWidth: 100,
				want:     "[==============================================>-]",
			},
			"t,c,bw{100,100,100}": {
				total:    100,
				current:  100,
				barWidth: 100,
				want:     "[================================================]",
			},
		},
	}

	prependWs := newWidthSync(nil, 1, 0)
	appendWs := newWidthSync(nil, 1, 0)
	for termWidth, cases := range testSuite {
		for name, tc := range cases {
			s := newTestState()
			s.width = tc.barWidth
			s.total = tc.total
			s.current = tc.current
			if tc.barRefill != nil {
				s.refill = tc.barRefill
			}
			s.draw(termWidth, prependWs, appendWs)
			got := s.bufB.String()
			if got != tc.want {
				t.Errorf("termWidth %d; %s: want: %s %d, got: %s %d\n", termWidth, name, tc.want, len(tc.want), got, len(got))
			}
		}
	}
}

func newTestState() *bState {
	s := &bState{
		trimLeftSpace:  true,
		trimRightSpace: true,
		bufP:           new(bytes.Buffer),
		bufB:           new(bytes.Buffer),
		bufA:           new(bytes.Buffer),
	}
	s.updateFormat("[=>-]")
	return s
}
