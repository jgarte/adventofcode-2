package day19

import (
	"testing"

	"github.com/Saser/adventofcode/internal/testcase"
)

const (
	inputFile        = "../testdata/19"
	part1ExampleFile = "testdata/p1example"
	part2ExampleFile = "testdata/p2example"
)

var (
	tcPart1 = testcase.NewFile("input", inputFile, "111")
	tcPart2 = testcase.NewFile("input", inputFile, "343")
)

func TestPart1(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", part1ExampleFile, "2"),
		tcPart1,
	} {
		tc.Test(t, Part1)
	}
}

func BenchmarkPart1(b *testing.B) {
	tcPart1.Benchmark(b, Part1)
}

func TestPart2(t *testing.T) {
	for _, tc := range []testcase.TestCase{
		testcase.NewFile("example", part2ExampleFile, "12"),
		tcPart2,
	} {
		tc.Test(t, Part2)
	}
}

func BenchmarkPart2(b *testing.B) {
	tcPart2.Benchmark(b, Part2)
}
