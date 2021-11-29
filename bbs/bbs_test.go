package bbs

import (
	"reflect"
	"testing"
)

func countSeries(bits []bool, valToCount bool) (seriesNum map[int]int) {
	seriesNum = make(map[int]int)

	var currLength int
	for _, val := range bits {
		if val == valToCount {
			currLength += 1
		} else if currLength != 0 {
			seriesNum[currLength] += 1
			currLength = 0
		}
	}

	if currLength != 0 {
		seriesNum[currLength] += 1
	}

	return
}

func countBits(bits []bool, valToCount bool) (number int) {
	for _, val := range bits {
		if val == valToCount {
			number += 1
		}
	}

	return
}

func TestFIPS140_2Level1(t *testing.T) {
	bits, _ := BBS(90187, 31511, 3, 20000)

	onesNum := countBits(bits, true)
	if onesNum <= 9725 || onesNum >= 10275 {
		t.Errorf("number of binary ones (%v) expected to be between 9725 and 19275", onesNum)
	}
}

func TestFIPS140_2Level2(t *testing.T) {
	bits, _ := BBS(90187, 31511, 3, 20000)

	check := func(series map[int]int, seriesLength int, min int, max int, tag string) {
		if series[seriesLength] < min || series[seriesLength] > max {
			t.Errorf("number of binary %v with length %v expected to be between %v and %v, actual value is %v", tag, seriesLength, min, max, series[seriesLength])
		}
	}

	onesSeries := countSeries(bits, true)
	zerosSeries := countSeries(bits, false)

	check(onesSeries, 1, 2315, 2685, "ones")
	check(zerosSeries, 1, 2315, 2685, "zeros")

	check(onesSeries, 2, 1114, 1386, "ones")
	check(zerosSeries, 2, 1114, 1386, "zeros")

	check(onesSeries, 3, 527, 723, "ones")
	check(zerosSeries, 3, 527, 723, "zeros")

	check(onesSeries, 4, 240, 384, "ones")
	check(zerosSeries, 4, 240, 384, "zeros")

	check(onesSeries, 5, 103, 209, "ones")
	check(zerosSeries, 5, 103, 209, "zeros")

	sumSeriesOnes := 0
	for seriesLength := range onesSeries {
		if seriesLength > 5 {
			sumSeriesOnes += onesSeries[seriesLength]
		}
	}

	if 103 > sumSeriesOnes || sumSeriesOnes > 209 {
		t.Errorf("number of binary ones with length 6+ expected to be between 103 and 209, actual value is %v", sumSeriesOnes)
	}

	sumSeriesZeros := 0
	for seriesLength := range zerosSeries {
		if seriesLength >= 6 {
			sumSeriesZeros += zerosSeries[seriesLength]
		}
	}

	if 103 > sumSeriesZeros || sumSeriesZeros > 209 {
		t.Errorf("number of binary ones with length 6+ expected to be between 103 and 209, actual value is %v", sumSeriesZeros)
	}
}

func TestFIPS140_2Level3(t *testing.T) {
	bits, _ := BBS(90187, 31511, 3, 20000)

	onesSeries := countSeries(bits, true)
	zerosSeries := countSeries(bits, false)

	for k := range onesSeries {
		if k >= 26 {
			t.Errorf("exists series with the length of %v, but expected no series with length >= 26", k)
		}
	}

	for k := range zerosSeries {
		if k >= 26 {
			t.Errorf("exists series with the length of %v, but expected no series with length >= 26", k)
		}
	}
}

func TestFIPS140_2Level4For20000Bits(t *testing.T) {
	bits, _ := BBS(90187, 31511, 3, 20000)

	counts := make(map[int]int)
	for i := 0; i < 20000; i += 4 {
		var decimal int
		if bits[i] {
			decimal |= 0b0001
		}
		if bits[i+1] {
			decimal |= 0b0010
		}
		if bits[i+2] {
			decimal |= 0b0100
		}
		if bits[i+3] {
			decimal |= 0b1000
		}
		counts[decimal] += 1
	}

	var sum int
	for i := 0; i < 16; i++ {
		sum += counts[i] * counts[i]
	}

	x := 16.0/5000.0*float64(sum) - 5000

	if x <= 2.16 || 46.17 <= x {
		t.Errorf("expected x (%v) to be between 2.16 and 46.17", x)
	}
}

func TestBBSValue1(t *testing.T) {
	actual, err := BBS(11, 23, 3, 6)
	if err != nil {
		t.Fatalf("error was not expected: %q", err)
	} else if expected := [...]bool{true, true, false, false, true, false}; !reflect.DeepEqual(expected[:], actual) {
		t.Fatalf("expected %v, but got %v", expected, actual)
	}
}

func TestBBSValue2(t *testing.T) {
	actual, err := BBS(11, 19, 3, 6)
	if err != nil {
		t.Fatalf("error was not expected: %q", err)
	} else if expected := [...]bool{true, true, false, false, false, false}; !reflect.DeepEqual(expected[:], actual) {
		t.Fatalf("expected %v, but got %v", expected, actual)
	}
}

func TestTestingTools(t *testing.T) {
	bits, _ := BBS(11, 23, 3, 6)

	actualCountZeroSeries := countSeries(bits, false)
	expectedCountZeroSeries := map[int]int{
		1: 1,
		2: 1,
	}
	if !reflect.DeepEqual(expectedCountZeroSeries, actualCountZeroSeries) {
		t.Errorf("expected counts to be %v, but got %v", expectedCountZeroSeries, actualCountZeroSeries)
	}

	actualCountOnesSeries := countSeries(bits, true)
	expectedCountOnesSeries := map[int]int{
		1: 1,
		2: 1,
	}
	if !reflect.DeepEqual(expectedCountOnesSeries, actualCountOnesSeries) {
		t.Errorf("expected counts to be %v, but got %v", expectedCountOnesSeries, expectedCountOnesSeries)
	}
}

func TestBadInput(t *testing.T) {
	if bits, err := BBS(1, 2, 0, 20000); bits != nil || err == nil {
		t.Fatal("error was expected")
	}
}
