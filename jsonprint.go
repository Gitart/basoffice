
// 
// Title               : Просмотр JSON 
// Date & Time         : 12.02.2016 12:00
// 
func jsonPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
