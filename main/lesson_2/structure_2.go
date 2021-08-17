package main

type testStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (input *testStruct) Shoot() bool {
	if input.On == false {
		return false
	} else if input.Ammo != 0 {
		input.Ammo -= 1
		return true
	} else {
		return false
	}
}

func (input *testStruct) RideBike() bool {
	if input.On == false {
		return false
	} else if input.Power != 0 {
		input.Power -= 1
		return true
	} else {
		return false
	}
}

func main() {
	var str = *new(testStruct)
}
