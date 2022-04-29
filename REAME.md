# golang_car_fleet

There are `n` cars going to the same destination along a one-lane road. The destination is `target` miles away.

You are given two integer array `position` and `speed`, both of length `n`, where `position[i]` is the position of the `ith` car and `speed[i]` is the speed of the `ith` car (in miles per hour).

A car can never pass another car ahead of it, but it can catch up to it and drive bumper to bumper **at the same speed**. The faster car will **slow down** to match the slower car's speed. The distance between these two cars is ignored (i.e., they are assumed to have the same position).

A **car fleet** is some non-empty set of cars driving at the same position and same speed. Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.

Return *the **number of car fleets** that will arrive at the destination*.

## Examples

**Example 1:**

```
Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
Output: 3
Explanation:
The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12.
The car starting at 0 does not catch up to any other car, so it is a fleet by itself.
The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.
Note that no other cars meet these fleets before the destination, so the answer is 3.

```

**Example 2:**

```
Input: target = 10, position = [3], speed = [3]
Output: 1
Explanation: There is only one car, hence there is only one fleet.

```

**Example 3:**

```
Input: target = 100, position = [0,2,4], speed = [4,2,1]
Output: 1
Explanation:
The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4. The fleet moves at speed 2.
Then, the fleet (speed 2) and the car starting at 4 (speed 1) become one fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.

```

**Constraints:**

- `n == position.length == speed.length`
- `1 <= n <= 10^5`
- `0 < target <= $10^6$`
- `0 <= position[i] < target`
- All the values of `position` are **unique**.
- `0 < speed[i] <= $10^6$`

## 解析

所謂 Car Fleet 是指同時以同速度到達的車群

而因為只有一條路所以如果後車追上前車則以前車速度為主

因此只要把每個車子的位置做排序並且算出每台到達終點所需時間

然後從最接近終點的車子開始往後來看， 如果有時間大於目前所需時間的車， 代表後車跟前車不會同時到達屬於新的車群

## 解法

step 1: 初始化一個 物件陣列 cars: []Car，物件內裏面紀錄 所有車到達終點的時間與當下位置

step 2: 排序 cars

step 3: 初始化 index = 0,  last_time = 0, car_fleet = 0

step 4: 當 index < car.length 檢查 last_time < cars[index].time 

step 5:  如果 last_time < cars[index].time， 更新 last_time = cars[idx].time , car_fleet += 1

step 6:  index = index +1 回到 step 4

step 7: 回傳 car_fleet

## 程式碼

```go
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

```

## 困難點

1. 需要知道題目所說的Car Fleet
2. 此題用 stack 並非直覺解法
3. 需要知道語言的 sort 作法 

# Solve point

- [x]  Understand car fleet definition
- [x]  Figure what data structure for return type
- [x]  Analysis Complexity