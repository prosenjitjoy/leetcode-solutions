# [Problem statement](https://leetcode.com/problems/insert-delete-getrandom-o1)

Implement the `RandomizedSet` class:

* `RandomizedSet()` Initializes the `RandomizedSet` object.
* `bool insert(int val)` Inserts an item `val` into the set if not present. Returns `true` if the item was not present, `false` otherwise.
* `bool remove(int val)` Removes an item `val` from the set if present. Returns `true` if the item was present, `false` otherwise.
* `int getRandom()` Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the **same probability** of being returned.

You must implement the functions of the class such that each function works in **average** `O(1)` time complexity.

**Example 1:**


**Input**
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
**Output**
[null, true, false, true, 2, true, false, 2]

**Explanation**
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.

**Constraints:**

* `-231 <= val <= 231 - 1`
* At most `2 * ` `105` calls will be made to `insert`, `remove`, and `getRandom`.
* There will be **at least one** element in the data structure when `getRandom` is called.

<br />

# [Solution in go](https://leetcode.com/submissions/detail/1149960875/)

```go
type RandomizedSet struct {
    Map map[int]int
    List []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        Map: make(map[int]int),
        List: []int{},
    }
}

func (rs *RandomizedSet) Insert(val int) bool {
    if _, ok := rs.Map[val]; !ok {
        rs.Map[val] = len(rs.List)
        rs.List = append(rs.List, val)
        return true
    }
    return false
}

func (rs *RandomizedSet) Remove(val int) bool {
    if v, ok := rs.Map[val]; ok {
        lastIdx := len(rs.List)-1
        lastVal := rs.List[lastIdx]
        rs.List[v] = lastVal
        rs.List = rs.List[:lastIdx]
        rs.Map[lastVal] = v
        delete(rs.Map, val)
        return true
    }
    return false
}


func (rs *RandomizedSet) GetRandom() int {
    return rs.List[rand.Intn(len(rs.List))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
```