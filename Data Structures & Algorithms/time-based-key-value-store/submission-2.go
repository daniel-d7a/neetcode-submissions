type Item struct {
	Ts int
	Value string
}

type TimeMap struct {
	Data map[string][]Item
}

func Constructor() TimeMap {
	return TimeMap{Data: map[string][]Item{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.Data[key]; ok {
		this.Data[key] = append(this.Data[key], Item{timestamp, value})
	} else {
		this.Data[key] = []Item{{timestamp, value}}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if data, ok := this.Data[key]; ok {
		return binarySearch(data, timestamp)
	} else {
		return ""
	}

}

func binarySearch(items []Item, target int) string {

	mid := len(items) / 2
	l := 0
	r := len(items) - 1

	for l <= r {

		if items[mid].Ts < target {
			l = mid + 1
		} else if items[mid].Ts > target {
			r = mid - 1
		} else {
			return items[mid].Value
		}

		mid = (r + l) / 2
	} 

	if items[mid].Ts <= target {
		return items[mid].Value
	} else {
		return ""
	}
}
