package main

/* лидирующие нули
lpad("3","0",2) result: "03"
lpad("12","0",6) result: "000012"
*/
func lpad(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}

// обработка err
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
