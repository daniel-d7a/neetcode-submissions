import "slices"

func carFleet(target int, position []int, speed []int) int {

	cars := []Car{}
	for i, p := range position {
		cars = append(cars, Car{p, speed[i]})
	}

	slices.SortFunc(cars, func(a, b Car) int {
		return b.Pos - a.Pos
	})


	fleets := []Car{}
	for _, c := range cars {
		if len(fleets) == 0 {
			fleets = append(fleets, c)
		} else {
			top := fleets[len(fleets) - 1]
			
			topTime := float64(target - top.Pos) / float64(top.Speed)
			cTime := float64(target - c.Pos) / float64(c.Speed)

			if cTime > topTime {
				fleets = append(fleets, c)
			} 
		}
	}
	return len(fleets)

}

type Car struct {
	Pos int
	Speed int
}