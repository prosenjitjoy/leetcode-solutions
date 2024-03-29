func pushDominoes(dominoes string) string {
    line := []byte(dominoes)
    queue := Queue{}

    for i, v := range line {
        if v != '.' {
            queue.Enqueue(pair{index: i, value: v})
        }
    }

    for !queue.Empty() {
        d := queue.Dequeue()

        if d.value == 'L' && d.index > 0 && line[d.index-1] == '.' {
            queue.Enqueue(pair{index: d.index-1, value: 'L'})
            line[d.index-1] = 'L'
        } else if d.value == 'R' && d.index+1 < len(line) && line[d.index+1] == '.' {
            if d.index+2 < len(line) && line[d.index+2] == 'L' {
                queue.Dequeue()
            } else {
                queue.Enqueue(pair{index: d.index+1, value: 'R'})
                line[d.index+1] = 'R'
            }
        }
    }

    return string(line)
}

type pair struct {
    index int
    value byte
}
type Queue struct {
    items []pair
}
func (q *Queue) Enqueue(val pair) {
    q.items = append(q.items, val)
}
func (q *Queue) Dequeue() pair {
    toRemove := q.items[0]
    q.items = q.items[1:]
    return toRemove
}
func (q *Queue) Empty() bool {
    return len(q.items) == 0
}