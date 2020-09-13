package behavioral

import "testing"

func TestSliceIterator(t *testing.T) {
	iterator := sliceIterator{
		S: []interface{}{1, 2, 3},
	}
	for index, expected := range []int{1, 2, 3} {
		actual, hasNext, _ := iterator.next()
		if actual.(int) != expected {
			t.Errorf("[%d] expected:'%d', actual: '%d'", index, expected, actual.(int))
		}

		if index == 2 && hasNext == true {
			t.Errorf("hasNext should be false on the last interation")
		}
	}
}
