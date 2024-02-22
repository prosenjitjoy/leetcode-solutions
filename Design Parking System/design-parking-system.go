type ParkingSystem struct {
    space [3]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{
        space: [3]int{big, medium, small},
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    if this.space[carType-1] > 0 {
        this.space[carType-1]--
        return true
    }

    return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */