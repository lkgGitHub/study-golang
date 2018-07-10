package main

type PersonInfo struct {
	Id string
	Name string
	Address string
}

func main() {
	//var personDB map[string] PersonInfo
	personDB := make(map[string] PersonInfo)

	personDB["1"] = PersonInfo{"1","Tome","Room"}
	personDB["2"] = PersonInfo{"2","2Tome","2Room"}

	person, ok := personDB["1"]

	if ok {
		println("find persion ", person.Name)
	}else {
		println("no find")
	}

	println("return",example(3))
}

func example(x int) int {
	if x == 0 {
		return 0
	}else {
		return 1
	}
}