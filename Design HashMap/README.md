# [Problem statement](https://leetcode.com/problems/design-hashmap)

Design a HashMap without using any built-in hash table libraries.

Implement the `MyHashMap` class:

* `MyHashMap()` initializes the object with an empty map.
* `void put(int key, int value)` inserts a `(key, value)` pair into the HashMap. If the `key` already exists in the map, update the corresponding `value`.
* `int get(int key)` returns the `value` to which the specified `key` is mapped, or `-1` if this map contains no mapping for the `key`.
* `void remove(key)` removes the `key` and its corresponding `value` if the map contains the mapping for the `key`.

**Example 1:**


**Input**
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
**Output**
[null, null, null, 1, -1, null, 1, null, -1]

**Explanation**
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // The map is now [[1,1]]
myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]

**Constraints:**

* `0 <= key, value <= 106`
* At most `104` calls will be made to `put`, `get`, and `remove`.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/956077880/)

```go
type node struct {
    key int
    val int
    next *node
}

func NewNode(key, value int) *node {
    return &node {
        key: key,
        val: value,
        next: nil,
    }
}

type MyHashMap struct {
    arr []*node
}

func Constructor() MyHashMap {
    hsize := 1250
    array := make([]*node, hsize)

    for i:=0; i<hsize; i++ {
        array[i] = NewNode(-1, -1)
    }

    return MyHashMap{
        arr: array,
    }
}

func (this *MyHashMap) hash(key int) int {
    return key % len(this.arr)
}

func (this *MyHashMap) Put(key int, value int)  {
    cur := this.arr[this.hash(key)]

    for cur.next != nil {
        if cur.next.key == key {
            cur.next.val = value
            return
        }
        cur = cur.next
    }

    cur.next = NewNode(key, value)
}

func (this *MyHashMap) Get(key int) int {
    cur := this.arr[this.hash(key)]

    for cur.next != nil {
        if cur.next.key == key {
            return cur.next.val
        }
        cur = cur.next
    }

    return -1
}

func (this *MyHashMap) Remove(key int)  {
    cur := this.arr[this.hash(key)]

    for cur != nil && cur.next != nil {
        if cur.next.key == key {
            cur.next = cur.next.next
            return
        }
        cur = cur.next
    }
}
```