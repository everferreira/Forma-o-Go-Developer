package main

import "fmt"

const ebulicaok = 370.0

func main()  {
	
	tempK := ebulicaok
	tempC := tempK - 273.0

	fmt.printf("A temperatura de ebulição da água em k = %g , temperatura de ebulição da água em °C =%g.", tempK, tempC)
}