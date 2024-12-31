//No longer working on theese test cases.
//There must be a better way of defining the test cases.
//Investigation needed and to be done in subsequent riddles

package day2

import (
	"slices"
	"testing"
)

func TestPart2Addition(t *testing.T) {
	nbTestsRun := 0
	nbFailedTests := 0
	//create input
	//increasing array
	inc_arr := [6]int{1, 2, 3, 4, 5, 6}
	//decreasing array
	dec_arr := [6]int{6, 5, 4, 3, 2, 1}
	//increasing array 1 repetition
	inc_arr_1rep := [6]int{1, 2, 3, 3, 5, 6}
	//decreasing array 1 repetition
	dec_arr_1rep := [6]int{6, 5, 3, 3, 2, 1}
	//increasing array 1 disruption
	inc_arr_1dis := [6]int{1, 2, 3, 8, 5, 6}
	//decreasing array 1 disruption
	dec_arr_1dis := [6]int{6, 5, 3, 8, 2, 1}
	//increasing array 1 disruption 1 repetition
	inc_arr_1dis_1rep := [6]int{1, 2, 2, 8, 5, 6}
	//decreasing array 1 disruption 1 repetition
	dec_arr_1dis_1rep := [6]int{6, 5, 2, 8, 2, 1}
	//increasing array 1 disruption 1 repetition 1 disruption at beginning
	inc_arr_1dis_1rep_beg := [6]int{8, 8, 2, 3, 5, 6}
	//decreasing array 1 disruption 1 repetition 1 disruption at beginning
	dec_arr_1dis_1rep_beg := [6]int{4, 4, 5, 3, 2, 1}
	//increasing array 1 disruption 1 repetition 1 disruption at end
	inc_arr_1dis_1rep_end := [6]int{2, 3, 4, 5, 1, 1}
	//decreasing array 1 disruption 1 repetition 1 disruption at end
	dec_arr_1dis_1rep_end := [6]int{6, 5, 4, 3, 7, 7}
	//increasing array 1 disruption 1 repetition 1 disruption at end 1 disruption at beginning
	inc_arr_1dis_1rep_beg_end := [6]int{2, 2, 4, 5, 1, 1}
	//decreasing array 1 disruption 1 repetition 1 disruption at end 1 disruption at beginning
	dec_arr_1dis_1rep_beg_end := [6]int{6, 6, 5, 4, 7, 7}
	//disrupted array
	dis_arr := [6]int{1, 8, 5, 7, 6, 3}
	//disrupted array 1 repetition
	dis_arr_1rep := [6]int{1, 8, 8, 2, 6, 3}
	//disrupted array 1 disruption 1 repetition
	dis_arr_1dis_1rep := [6]int{1, 2, 8, 8, 5, 6}

	//describe expected output
	want_inc_sice := []level{{inc_arr[:], 0}}
	want_dec_slice := []level{{dec_arr[:], 0}}
	want_inc_slice_1rep := []level{{inc_arr_1rep[:], 0}}
	want_dec_slice_1rep := []level{{dec_arr_1rep[:], 0}}
	want_inc_slice_1dis := []level{{[]int{1, 2, 3, 5, 6}, 1}}
	want_dec_slice_1dis := []level{{[]int{6, 5, 3, 2, 1}, 1}}
	want_inc_slice_1dis_1rep := []level{{[]int{1, 2, 2, 5, 6}, 1}}
	want_dec_slice_1dis_1rep := []level{{[]int{6, 5, 2, 2, 1}, 1}}
	want_inc_slice_1dis_1rep_beg := []level{}
	want_dec_slice_1dis_1rep_beg := []level{{[]int{4, 4, 3, 2, 1}, 1}}
	want_inc_slice_1dis_1rep_end := []level{}
	want_dec_slice_1dis_1rep_end := []level{}
	want_inc_slice_1dis_1rep_beg_end := []level{}
	want_dec_slice_1dis_1rep_beg_end := []level{}
	want_dis_slice := []level{}
	want_dis_slice_1rep := []level{}
	want_dis_slice_1dis_1rep := []level{}

	//run function
	output_inc_slice := part2Addition(inc_arr[:])
	nbTestsRun++
	output_dec_slice := part2Addition(dec_arr[:])
	nbTestsRun++
	output_inc_slice_1rep := part2Addition(inc_arr_1rep[:])
	nbTestsRun++
	output_dec_slice_1rep := part2Addition(dec_arr_1rep[:])
	nbTestsRun++
	output_inc_slice_1dis := part2Addition(inc_arr_1dis[:])
	nbTestsRun++
	output_dec_slice_1dis := part2Addition(dec_arr_1dis[:])
	nbTestsRun++
	output_inc_slice_1dis_1rep := part2Addition(inc_arr_1dis_1rep[:])
	nbTestsRun++
	output_dec_slice_1dis_1rep := part2Addition(dec_arr_1dis_1rep[:])
	nbTestsRun++
	output_inc_slice_1dis_1rep_beg := part2Addition(inc_arr_1dis_1rep_beg[:])
	nbTestsRun++
	output_dec_slice_1dis_1rep_beg := part2Addition(dec_arr_1dis_1rep_beg[:])
	nbTestsRun++
	output_inc_slice_1dis_1rep_end := part2Addition(inc_arr_1dis_1rep_end[:])
	nbTestsRun++
	output_dec_slice_1dis_1rep_end := part2Addition(dec_arr_1dis_1rep_end[:])
	nbTestsRun++
	output_inc_slice_1dis_1rep_beg_end := part2Addition(inc_arr_1dis_1rep_beg_end[:])
	nbTestsRun++
	output_dec_slice_1dis_1rep_beg_end := part2Addition(dec_arr_1dis_1rep_beg_end[:])
	nbTestsRun++
	output_dis_slice := part2Addition(dis_arr[:])
	nbTestsRun++
	output_dis_slice_1rep := part2Addition(dis_arr_1rep[:])
	nbTestsRun++
	output_dis_slice_1dis_1rep := part2Addition(dis_arr_1dis_1rep[:])
	nbTestsRun++

	//check output
	if !slices.EqualFunc(want_inc_sice, output_inc_slice, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)

	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_inc_slice, want_inc_sice)
	}

	if !slices.EqualFunc(want_dec_slice, output_dec_slice, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)

	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dec_slice, want_dec_slice)
	}

	if !slices.EqualFunc(want_inc_slice_1rep, output_inc_slice_1rep, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)

	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_inc_slice_1rep, want_inc_slice_1rep)
	}

	if !slices.EqualFunc(want_dec_slice_1rep, output_dec_slice_1rep, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)

	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dec_slice_1rep, want_dec_slice_1rep)
	}

	if !slices.EqualFunc(want_inc_slice_1dis, output_inc_slice_1dis, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)

	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_inc_slice_1dis, want_inc_slice_1dis)
	}

	if !slices.EqualFunc(want_dec_slice_1dis, output_dec_slice_1dis, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)

	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dec_slice_1dis, want_dec_slice_1dis)
	}

	if !slices.EqualFunc(want_inc_slice_1dis_1rep, output_inc_slice_1dis_1rep, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)

	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_inc_slice_1dis_1rep, want_inc_slice_1dis_1rep)
	}

	if !slices.EqualFunc(want_dec_slice_1dis_1rep, output_dec_slice_1dis_1rep, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)

	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dec_slice_1dis_1rep, want_dec_slice_1dis_1rep)
	}

	if !slices.EqualFunc(want_inc_slice_1dis_1rep_beg, output_inc_slice_1dis_1rep_beg, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_inc_slice_1dis_1rep_beg, want_inc_slice_1dis_1rep_beg)
	}

	if !slices.EqualFunc(want_dec_slice_1dis_1rep_beg, output_dec_slice_1dis_1rep_beg, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dec_slice_1dis_1rep_beg, want_dec_slice_1dis_1rep_beg)
	}
	if !slices.EqualFunc(want_inc_slice_1dis_1rep_end, output_inc_slice_1dis_1rep_end, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_inc_slice_1dis_1rep_end, want_inc_slice_1dis_1rep_end)
	}

	if !slices.EqualFunc(want_dec_slice_1dis_1rep_end, output_dec_slice_1dis_1rep_end, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dec_slice_1dis_1rep_end, want_dec_slice_1dis_1rep_end)
	}

	if !slices.EqualFunc(want_inc_slice_1dis_1rep_beg_end, output_inc_slice_1dis_1rep_beg_end, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_inc_slice_1dis_1rep_beg_end, want_inc_slice_1dis_1rep_beg_end)
	}

	if !slices.EqualFunc(want_dec_slice_1dis_1rep_beg_end, output_dec_slice_1dis_1rep_beg_end, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dec_slice_1dis_1rep_beg_end, want_dec_slice_1dis_1rep_beg_end)
	}

	if !slices.EqualFunc(want_dis_slice, output_dis_slice, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dis_slice, want_dis_slice)
	}

	if !slices.EqualFunc(want_dis_slice_1rep, output_dis_slice_1rep, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dis_slice_1rep, want_dis_slice_1rep)
	}

	if !slices.EqualFunc(want_dis_slice_1dis_1rep, output_dis_slice_1dis_1rep, func(l level, r level) bool {
		return l.numberOfTransgressions == r.numberOfTransgressions && slices.Equal(l.arr, r.arr)
	}) {
		nbFailedTests++
		t.Errorf("part2Addition() = %v, want %v", output_dis_slice_1dis_1rep, want_dis_slice_1dis_1rep)
	}

	t.Log(nbFailedTests, "out of", nbTestsRun, "tests failed.")
}

func TestAmIaTrueLevel(t *testing.T) {
	//create input
	//describe expected output
	//run function
	//check output
}

func TestMain(t *testing.T) {
	//create input
	//describe expected output
	//run function
	//check output
}
