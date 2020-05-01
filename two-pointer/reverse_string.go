package main

func main() {
	s := []byte {'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	println(string(s))
}

func reverseString(s []byte)  {
	lens := len(s)
	j := lens - 1
	for i := 0; i < lens && i <= j; i++ {
		c := s[i]
		s[i] = s[j]
		s[j] = c
		j--
	}
}