package car_fleet

import "sort"

type Car struct {
	time float64
	pos  int
}
type ByPosition []Car

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].pos > a[j].pos }
func carFleet(target int, position []int, speed []int) int {
	pLen := len(position)
	if pLen == 1 {
		return 1
	}
	cars := make([]Car, pLen)
	for idx := 0; idx < pLen; idx++ {
		cars[idx] = Car{time: float64(target-position[idx]) / float64(speed[idx]), pos: position[idx]}
	}
	sort.Sort(ByPosition(cars))
	carFleet, lastTime := 0, 0.0
	for idx := 0; idx < pLen; idx++ {
		if cars[idx].time > lastTime {
			carFleet++
			lastTime = cars[idx].time
		}
	}
	return carFleet
}
