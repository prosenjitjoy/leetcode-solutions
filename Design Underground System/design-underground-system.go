type stationInfo struct{
    name string
    time int 
}
type travelInfo struct{
    start string
    end string
}

type UndergroundSystem struct {
    checkInMap map[int]stationInfo
    totalTimeMap map[travelInfo][2]int
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{
        checkInMap: map[int]stationInfo{},
        totalTimeMap: map[travelInfo][2]int{},
    }
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.checkInMap[id] = stationInfo{
        name: stationName,
        time: t,
    }
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    station := this.checkInMap[id]

    route := travelInfo{
        start: station.name,
        end: stationName,
    }
    
    total := this.totalTimeMap[route]
    total[0] += t - station.time
    total[1] += 1
    this.totalTimeMap[route] = total
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    total := this.totalTimeMap[travelInfo{
        start: startStation,
        end: endStation,
    }]

    return float64(total[0]) / float64(total[1])
} 


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */