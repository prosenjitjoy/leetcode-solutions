# [Problem statement](https://leetcode.com/problems/design-hashset)

Design a HashSet without using any built-in hash table libraries.

Implement `MyHashSet` class:

* `void add(key)` Inserts the value `key` into the HashSet.
* `bool contains(key)` Returns whether the value `key` exists in the HashSet or not.
* `void remove(key)` Removes the value `key` in the HashSet. If `key` does not exist in the HashSet, do nothing.

**Example 1:**


**Input**
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
**Output**
[null, null, null, true, false, null, true, null, false]

**Explanation**
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // return True
myHashSet.contains(3); // return False, (not found)
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // return True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // return False, (already removed)

**Constraints:**

* `0 <= key <= 106`
* At most `104` calls will be made to `add`, `remove`, and `contains`.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1133234351/)

```go
type node struct {
    key int
    next *node
}

type MyHashSet struct {
    set []*node
}


func Constructor() MyHashSet {
    hsize := 1000
    arr := make([]*node, hsize)

    for i:=0; i<hsize; i++ {
        arr[i] = &node{}
    }

    return MyHashSet{set: arr}
}

func (this *MyHashSet) hash(key int) int {
    return key % len(this.set)
}

func (this *MyHashSet) Add(key int)  {
    cur := this.set[this.hash(key)]

    for cur.next != nil {
        if cur.next.key == key {
            return
        }
        cur = cur.next
    }

    cur.next = &node{key: key}
}


func (this *MyHashSet) Remove(key int)  {
    cur := this.set[this.hash(key)]

    for cur.next != nil {
        if cur.next.key == key {
            cur.next = cur.next.next
            return
        }
        cur = cur.next
    }
}


func (this *MyHashSet) Contains(key int) bool {
    cur := this.set[this.hash(key)]

    for cur.next != nil {
        if cur.next.key == key {
            return true
        }
        cur = cur.next
    }

    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
```