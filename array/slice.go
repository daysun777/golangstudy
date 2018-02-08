package array

import "fmt"

func init2() {

	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//s[3] = "d" 		// error 得用s = append(s, "d")
	fmt.Println(s)      //[a b c]
	fmt.Println(s[2])   //c
	fmt.Println(len(s)) //3
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s) //[a b c d e]

	c := make([]string, len(s))

	copy(c, s)
	fmt.Println(c) //[a b c d e]

	l := s[2:5]
	fmt.Println(l) //[c d e]

	l = s[1:]
	fmt.Println(l) //[b c d e]

	l = s[:5]
	fmt.Println(l) //[a b c d e]

	t := []string{"g","h","i"}
	fmt.Println(t) //[g h i]

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
