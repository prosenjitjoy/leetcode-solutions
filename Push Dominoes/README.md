# [Problem statement](https://leetcode.com/problems/push-dominoes)

There are `n` dominoes in a line, and we place each domino vertically upright. In the beginning, we simultaneously push some of the dominoes either to the left or to the right.

After each second, each domino that is falling to the left pushes the adjacent domino on the left. Similarly, the dominoes falling to the right push their adjacent dominoes standing on the right.

When a vertical domino has dominoes falling on it from both sides, it stays still due to the balance of the forces.

For the purposes of this question, we will consider that a falling domino expends no additional force to a falling or already fallen domino.

You are given a string `dominoes` representing the initial state where:

* `dominoes[i] = 'L'`, if the `ith` domino has been pushed to the left,
* `dominoes[i] = 'R'`, if the `ith` domino has been pushed to the right, and
* `dominoes[i] = '.'`, if the `ith` domino has not been pushed.

Return _a string representing the final state_.

**Example 1:**


**Input:** dominoes = "RR.L"
**Output:** "RR.L"
**Explanation:** The first domino expends no additional force on the second domino.

**Example 2:**

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/05/18/domino.png) 


**Input:** dominoes = ".L.R...LR..L.."
**Output:** "LL.RR.LLRRLL.."

**Constraints:**

* `n == dominoes.length`
* `1 <= n <= 105`
* `dominoes[i]` is either `'L'`, `'R'`, or `'.'`.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1147939339/)

```go
func pushDominoes(dominoes string) string {
    line := []byte(dominoes)
    queue := Queue{}
    N := len(line)

    for i, ch := range line {
        if ch != '.' {
            queue.Enqueue(pair{index: i, value: ch})
        }
    }

    for !queue.Empty() {
        d := queue.Dequeue()

        if d.value == 'L' {
            if d.index > 0 && line[d.index-1] == '.' {
                queue.Enqueue(pair{index: d.index-1, value: 'L'})
                line[d.index-1] = 'L'
            }
        } else {
            if d.index+1 < N && line[d.index+1] == '.' {
                if d.index+2 < N && line[d.index+2] == 'L' {
                    queue.Dequeue()
                } else {
                    queue.Enqueue(pair{index: d.index+1, value: 'R'})
                    line[d.index+1] = 'R'
                }
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
```