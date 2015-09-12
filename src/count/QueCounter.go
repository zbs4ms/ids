package count

type QueCounter struct{
    OrderNow uint64
    BeginOrder uint64
    MaxOrder uint64
    Step uint64
}

func NewQueCounter(begin uint64, max uint64, step uint64) *QueCounter {
    return &QueCounter{begin, begin, max, step}
}

func (q *QueCounter) NextOrder() uint64 {
    if q.OrderNow >= q.MaxOrder || q.OrderNow + q.Step > q.MaxOrder { //重新开始计数
        q.OrderNow = q.BeginOrder
    }
    q.OrderNow = q.OrderNow + q.Step
    return q.OrderNow
}
