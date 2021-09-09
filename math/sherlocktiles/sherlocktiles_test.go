package sherlocktiles

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMovingTiles(t *testing.T) {
	tt := []struct {
		name   string
		width  int       // width of both squares in m
		v1, v2 int       // velocity of each along y=x in m/s
		in     []int64   // slice of areas
		want   []float64 // slice of times
	}{
		{
			name:  "example1 from hacker rank",
			width: 10,
			v1:    1,
			v2:    2,
			in:    []int64{50, 100},
			want:  []float64{4.1421356237, 0},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := movingTiles(tc.width, tc.v1, tc.v2, tc.in)

			if len(got) != len(tc.want) {
				t.Fatalf("sent query of length %d, but got %d results",
					len(tc.want), len(got))
			}
			for i, time := range tc.want {
				assertClose(t, time, got[i])
			}
		})
	}
}

func TestMovingTiles_FromInput(t *testing.T) {
	tt := []struct {
		infile, outfile string
	}{
		{"input03.txt", "output03.txt"},
		{"input05.txt", "output05.txt"},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.infile, func(t *testing.T) {
			t.Parallel()
			width, v1, v2, queries := parseInput(t, tc.infile)
			got := movingTiles(width, v1, v2, queries)
			want := parseOutput(t, tc.outfile)
			for i, seconds := range want {
				assertClose(t, seconds, got[i])
			}
		})
	}
}

func parseInput(t *testing.T, filename string) (width, v1, v2 int, queries []int64) {
	t.Helper()
	in, err := os.Open(filename)
	if err != nil {
		t.Fatalf("unexpected error opening test input: %s", err)
	}
	defer in.Close()

	s := bufio.NewScanner(in)
	ok := s.Scan()
	if !ok {
		t.Fatalf("missing first line from input file")
	}

	header := strings.Split(strings.TrimSpace(s.Text()), " ")

	if width, err = strconv.Atoi(header[0]); err != nil {
		t.Fatalf("error parsing width of tiles: %s", err)
	}

	if v1, err = strconv.Atoi(header[1]); err != nil {
		t.Fatalf("error parsing speed of square 1: %s", err)
	}

	if v2, err = strconv.Atoi(header[2]); err != nil {
		t.Fatalf("error parsing speed of square 2: %s", err)
	}

	ok = s.Scan()
	if !ok {
		t.Fatalf("missing second line from input file")
	}

	queries = make([]int64, 0, 1024*8)
	for s.Scan() {
		if err := s.Err(); err != nil {
			t.Fatalf("unexpected error reading input: %s", err)
		}
		q, err := strconv.Atoi(s.Text())
		if err != nil {
			t.Fatalf("error parsing input: %s", err)
		}
		queries = append(queries, int64(q))
	}

	return width, v1, v2, queries
}

func parseOutput(t *testing.T, filename string) []float64 {
	t.Helper()
	out, err := os.Open(filename)
	if err != nil {
		t.Fatalf("unexpected error opening desired test results: %s", err)
	}
	defer out.Close()

	s := bufio.NewScanner(out)

	want := make([]float64, 0, 1024*8)
	for s.Scan() {
		if err := s.Err(); err != nil {
			t.Fatalf("unexpected error reading desired test results: %s", err)
		}
		seconds, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			t.Fatalf("error parsing output: %s", err)
		}
		want = append(want, seconds)
	}
	return want
}

func assertClose(t *testing.T, want, got float64) {
	t.Helper()
	delta := want - got
	if math.Abs(delta) > 10e-5 {
		t.Logf("got %8.5f   want %8.5f", got, want)
		t.Fail()
	}
}
