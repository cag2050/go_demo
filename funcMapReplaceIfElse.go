package main

func main() {
	// 使用map[string]func优化if-else语句，参考：https://segmentfault.com/q/1010000017173434

	var ColShowTypeFuncMap = map[string]func(jsonStr string) (sqlStr string, err error){}

	ColShowTypeFuncMap["string"] = func(s string) (string, error) {
		println("sql:" + s)
		return "sql" + s, nil
	}

	ColShowTypeFuncMap["string"]("ssss")
}
