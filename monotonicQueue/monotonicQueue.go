type MonotonicQueue struct {
	maxq []interface{}
	minq []interface{}
	q    []interface{}
}

func (mq *MonotonicQueue) push(elem interface{}) {
	if elem == nil || elem == "" {
		return
	}

	mq.q = append(mq.q, elem)

	for len(maxq) > 0 {
		lastIndex = len(maxq) - 1
		if mq.maxq[lastIndex].(Comparable).CompareTo(elem.(Comparable)) < 0 {
			mq.maxq = mq.maxq[:lastIndex]
		} else {
			break
		}
	}

	mq.maxq = append(mq.maxq, elem)

	for len(minq) > 0 {
		lastIndex = len(minq) - 1
		if mq.minq[lastIndex].(Comparable).CompareTo(elem.(Comparable)) > 0 {
			mq.minq = mq.minq[:lastIndex]
		} else {
			break
		}
	}

	mq.minq = append(mq.minq, elem)

}

func (mq *MonotonicQueue) max() interface{} {
	if len(maxq) > 0 {
		return mq.maxq[0]
	}
	return nil
}

func (mq *MonotonicQueue) min() interface{} {
	if len(minq) > 0 {
		return mq.minq[0]
	}
	return nil
}

func (mq *MonotonicQueue) pop() interface{} {
	if len(mq.q) == 0 {
		return
	}

	deleteVal := mq.q[0]
	mq.q = mq.q[1:]

	if len(mq.maxq) > 0 && deleteVal == mq.maxq[0] {
		mq.maxq = mq.maxq[1:]
	}
	if len(mq.minq) > 0 && deleteVal == mq.minq[0] {
		mq.minq = mq.minq[1:]
	}

	return deleteVal
}

func (mq *MonotonicQueue) size() int {
	return len(mq.q)
}

func (mq *MonotonicQueue) isEmpty() bool {
	return len(mq.q) == 0
}

type Comparable[T any] interface {
	CompareTo(other T) int
}

//单调队列可以和 滑动窗口 算法结合，在窗口滑动的过程中快速计算窗口内部的最值