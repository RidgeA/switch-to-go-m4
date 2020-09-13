package behavioral

import "fmt"

type (
	iterator interface {
		next() (value interface{}, hasNext bool, err error)
	}

	sliceIterator struct {
		S         []interface{}
		currIndex int
	}
)

func (si *sliceIterator) next() (interface{}, bool, error) {
	if si.currIndex == len(si.S) {
		return nil, false, fmt.Errorf("iterator out of range")
	}
	v := si.S[si.currIndex]
	hasNext := false
	if si.currIndex == len(si.S) {
		hasNext = true
	}
	si.currIndex++
	return v, hasNext, nil
}
