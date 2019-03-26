package main

func main() {
	var a interface{} = "55"
	//var a  = "55"
	switch a.(type) {
	case string:
		println("string")
	}

	var marks = 90
	switch marks {
	case 90:
		println("A")
	case 80:
		println("B")
	case 50, 60, 70:
		println("C")
	default:
		println("D")
	}
}
