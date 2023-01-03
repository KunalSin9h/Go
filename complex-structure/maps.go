package main

/*
Map
	map[x]y
	x -> y
*/

import "fmt"

func main() {

	userEmails := make(map[int]string)

	userEmails[1] = "kunal@gmail.com"
	userEmails[2] = "kunal@gmail.com"

	fmt.Println(userEmails)

	var userPassword map[int]string = map[int]string{
		1: "39f9sek",
		2: "d9hs89jife",
	}

	userPassword[3] = "d9j39fd"
	userPassword[1] = "k9"

	fmt.Println(userPassword)

	cnt := make(map[int]int)

	cnt[0]++
	cnt[1]++
	cnt[2]++
	cnt[1]++

	delete(cnt, 0)

	fmt.Println(cnt)

	/* Map lookup */
	if val, ok := cnt[9]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("NOT FOUND")
	}

}
