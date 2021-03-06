package util

import (
	"sort"
)

func NewPAscTimeRanges(timeRanges []*TimeRange) (result AscTimeRanges) {
	result = make(AscTimeRanges, 0)

	sort.Slice(timeRanges, func(i, j int) bool {
		it := timeRanges[i]
		jt := timeRanges[j]
		if compare := it.Compare(jt); compare == -1 {
			return true
		}
		return false
	})
	for _, v := range timeRanges {
		result = append(result, *v)
	}

	return
}

func NewAscTimeRanges(timeRanges []TimeRange) (result AscTimeRanges) {
	result = make(AscTimeRanges, 0)

	sort.Slice(timeRanges, func(i, j int) bool {
		it := timeRanges[i]
		jt := timeRanges[j]
		if compare := it.Compare(&jt); compare == -1 {
			return true
		}
		return false
	})
	for _, v := range timeRanges {
		result = append(result, v)
	}

	return
}

type TimeRanges []TimeRange

type AscTimeRanges TimeRanges

func (trs AscTimeRanges) Append(newInsertTimeRanges ...TimeRange) (newTimeRanges AscTimeRanges) {
	for i, newInsertTimeRange := range newInsertTimeRanges {
		if i == 0 {
			newTimeRanges = trs.insert(newInsertTimeRange)
		} else {
			newTimeRanges = newTimeRanges.insert(newInsertTimeRange)
		}
	}
	return
}

func (trs AscTimeRanges) insert(newInsertTimeRange TimeRange) (newTimeRanges AscTimeRanges) {
	newTimeRanges = make(AscTimeRanges, 0)

	if len(trs) > 0 {
		previousIndex, nextIndex := SearchUpDown(
			0, len(trs)-1,
			func(index int) int {
				t := trs[index]
				return newInsertTimeRange.Compare(&t)
			},
			false,
		)
		if isFirst := previousIndex == -1; isFirst {
			newTimeRanges = append(newTimeRanges, newInsertTimeRange)
			newTimeRanges = append(newTimeRanges, trs...)
		} else if isLast := nextIndex == -1; isLast {
			newTimeRanges = append(newTimeRanges, trs...)
			newTimeRanges = append(newTimeRanges, newInsertTimeRange)
		} else {
			insertIndex := nextIndex
			newTimeRanges = append(newTimeRanges, trs[:insertIndex]...)
			newTimeRanges = append(newTimeRanges, newInsertTimeRange)
			newTimeRanges = append(newTimeRanges, trs[insertIndex:]...)
		}
	} else {
		newTimeRanges = append(newTimeRanges, newInsertTimeRange)
	}

	return
}

func (trs AscTimeRanges) Contain(t TimeRange) *TimeRange {
	// TODO: use binary search
	for _, v := range trs {
		fromCompare := v.CompareTime(&t.From)
		if fromCompare == -1 {
			return nil
		} else if fromCompare == 1 {
			continue
		}
		toCompare := v.CompareTime(&t.To)
		if toCompare == 0 {
			return &v
		}
	}
	return nil
}

// newTimeRanges:?????????????????????
// intersectionTimeRanges:???????????????
// leftTimeRanges:??????????????????
func (trs AscTimeRanges) Sub(sub TimeRange) (
	newTimeRanges,
	intersectionTimeRanges,
	leftTimeRanges AscTimeRanges,
) {
	newTimeRanges = make(AscTimeRanges, 0)

	for i, timeRange := range trs {
		if isSubFromBefore := sub.From.Before(timeRange.From); isSubFromBefore {
			// current  | |
			// minus   |
			if isSubToBeforeFrom := sub.To.Before(timeRange.From); isSubToBeforeFrom {
				// current    | |
				// minus   | |
				// ????????? ??????
				leftTimeRanges = append(leftTimeRanges, sub)
				newTimeRanges = append(newTimeRanges, trs[i:]...)
				return
			} else {
				// current   | |
				// minus   |  |
				// or
				// minus   |    |
				// ????????? ???????????????
				leftTimeRanges = append(leftTimeRanges, TimeRange{
					From: sub.From,
					To:   timeRange.From,
				})
				sub.From = timeRange.From
			}
		} else if isSubFromAfterOrEqualTo := !sub.From.Before(timeRange.To); isSubFromAfterOrEqualTo {
			// current | |
			// minus     |
			// or
			// minus      |
			// ????????? ??????????????????
			newTimeRanges = append(newTimeRanges, timeRange)
			continue
		} else if isSubFromAfter := sub.From.After(timeRange.From); isSubFromAfter {
			// current   | |
			// minus      |
			// ????????????
			newTimeRanges = append(newTimeRanges, TimeRange{
				From: timeRange.From,
				To:   sub.From,
			})
			timeRange.From = sub.From
		}

		// current | |
		// minus   |

		if isSubToAfter := sub.To.After(timeRange.To); isSubToAfter {
			// current   ||
			// minus     | |
			// cut all current
			intersectionTimeRanges = append(intersectionTimeRanges, timeRange)

			sub.From = timeRange.To
			continue
		} else if sub.To.Equal(timeRange.To) {
			// current   ||
			// minus     ||
			// cut all current ??????
			intersectionTimeRanges = append(intersectionTimeRanges, timeRange)

			newTimeRanges = append(newTimeRanges, trs[i+1:]...)
			return
		}

		// current | |
		// minus   ||
		intersectionTimeRanges = append(intersectionTimeRanges, sub)

		// ????????????????????????
		currentLeftTimeRanges := trs[i+1:]
		newTimeRanges = append(newTimeRanges,
			currentLeftTimeRanges.Append(TimeRange{
				From: sub.To,
				To:   timeRange.To,
			})...,
		)
		return
	}

	if sub.To.After(sub.From) {
		leftTimeRanges = append(leftTimeRanges, sub)
	}
	return
}

func (trs AscTimeRanges) CombineByCount() (countAscTimeRangesMap map[int]AscTimeRanges) {
	countAscTimeRangesMap = make(map[int]AscTimeRanges)
	if len(trs) == 0 {
		return
	}

	currentTimeRange := trs[0]
	currentTargetTo := currentTimeRange.To
	currentTargetFrom := &currentTimeRange.From

	// ????????????????????????????????????
	nextAscTrs := NewAscTimeRanges(make([]TimeRange, 0))
	nextFromIndex := 1
	for ; nextFromIndex < len(trs); nextFromIndex++ {
		nextIndex := nextFromIndex
		next := trs[nextIndex]
		if !next.From.After(*currentTargetFrom) {
			// ??????????????????
			if isNotSame := next.To.After(currentTargetTo); isNotSame {
				// ??????????????????
				nextAscTrs = append(nextAscTrs, TimeRange{
					From: currentTargetTo,
					To:   next.To,
				})
			}
			continue
		}

		if next.From.Before(currentTargetTo) {
			// ????????????????????????????????????????????????
			// ????????????????????????????????????????????????
			currentTargetTo = next.From

			othersTimeRanges := make([]TimeRange, 0)
			// ???????????????????????????????????????
			nextFromsTimeRanges := make([]TimeRange, 0)
			for _, v := range trs[0:nextFromIndex] {
				v.From = next.From
				nextFromsTimeRanges = append(nextFromsTimeRanges, v)
			}
			for _, v := range trs[nextFromIndex:] {
				if v.From.After(next.From) {
					// ???????????????
					othersTimeRanges = append(othersTimeRanges, v)
					continue
				}
				nextFromsTimeRanges = append(nextFromsTimeRanges, v)
			}
			nextAscTrs = NewAscTimeRanges(nextFromsTimeRanges)

			// ?????????????????????
			nextAscTrs = append(nextAscTrs, othersTimeRanges...)
		}

		break
	}

	count := nextFromIndex
	_, exist := countAscTimeRangesMap[count]
	if !exist {
		countAscTimeRangesMap[count] = make(AscTimeRanges, 0)
	}
	countAscTimeRangesMap[count] = append(countAscTimeRangesMap[count], TimeRange{
		From: *currentTargetFrom,
		To:   currentTargetTo,
	})

	// ????????????????????????????????????
	m := nextAscTrs.CombineByCount()
	for count, v := range m {
		_, exist := countAscTimeRangesMap[count]
		if !exist {
			countAscTimeRangesMap[count] = make(AscTimeRanges, 0)
		}
		countAscTimeRangesMap[count] = append(countAscTimeRangesMap[count], v...)
	}

	return
}
