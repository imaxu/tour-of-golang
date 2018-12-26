package main

import "fmt"

func main() {
	
	s := make([]string,3)
	fmt.Println(s)

	s[0] = "i"
	s[1] = "m"
	s[2] = "a"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println(len(s))

	s = append(s,"x")
	
	s = append(s,"u","!")
	fmt.Println(s)


	c := make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy:", c)

	l := s[:5]
	fmt.Println("sl1:", l)

	l = s[2:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"x","w","h"}
	fmt.Println("dcl:", t)

	twoD := make([][]int,3)
	for i := 0;i<3;i++{
		innerLen := i + 1
		twoD[i] = make([]int,innerLen)
		for j := 0;j<innerLen;j++{
			twoD[i][j] = i * j
		}
	}
	fmt.Println(twoD)
}