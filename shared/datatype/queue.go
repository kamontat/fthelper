package datatype

type Queue struct {
	limit    int
	internal []interface{}
}

// If queue size is exceed limit, remove oldest data
// return true if data got removed
func (q *Queue) Dequeue() bool {
	if q.limit > 0 && len(q.internal) > q.limit {
		q.internal[0] = nil // free memory
		q.internal = q.internal[1:]
		return true
	}

	return false
}

// Add data to end of queue and maintain queue size
func (q *Queue) Enqueue(d interface{}) *Queue {
	q.internal = append(q.internal, d)
	q.Dequeue()
	return q
}

// Get first queue and remove from the queue
func (q *Queue) Get() interface{} {
	var data = q.Head()

	if data != nil {
		// Remove data from first queue
		q.internal[0] = nil // free memory
		q.internal = q.internal[1:]
	}
	return data
}

// Head will get first queue but not change queue size
func (q *Queue) Head() interface{} {
	if q.Size() == 0 {
		return nil
	}
	return q.internal[0]
}

// Tail will get last queue but not change queue size
func (q *Queue) Tail() interface{} {
	return q.internal[q.Size()-1]
}

// Size will get the size of queue
func (q *Queue) Size() int {
	return len(q.internal)
}

// Empty will check whether queue empty or not
func (q *Queue) Empty() bool {
	return q.Size() == 0
}

// NewQueue will create new infinity queue
func NewQueue() *Queue {
	return NewLimitQueue(-1)
}

// NewLimitQueue will create new limit queue
func NewLimitQueue(size int) *Queue {
	return ToLimitQueue(size, make([]interface{}, 0))
}

// ToQueue will try to convert array to queue
func ToQueue(v []interface{}) *Queue {
	return ToLimitQueue(-1, v)
}

// ToLimitQueue will try to convert array to limit queue
func ToLimitQueue(size int, v []interface{}) *Queue {
	return &Queue{
		limit:    size,
		internal: v,
	}
}
