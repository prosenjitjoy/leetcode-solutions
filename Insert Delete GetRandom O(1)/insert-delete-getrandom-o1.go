type RandomizedSet struct {
    Map map[int]int
    List []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        Map: map[int]int{},
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