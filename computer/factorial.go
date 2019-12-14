package computer

func calculate(num uint) uint {
	if num == 0 {
		return 1
	}
	return num * calculate(num-1)
}

func CalculateUsingGoRoutines(num uint, channel chan uint) {
	channel <- calculate(num)
}
