package time

var globalTime = time{day: 0}

type time struct {
	day int
}

func Day() int {
	return globalTime.day
}

func Set(day int) {
	globalTime.day = day
}
