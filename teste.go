package main

import (
    "time"
)

func main() {
    count := 0
    for {
        println(count, ": Hello, World")
        time.Sleep(time.Millisecond * 1000)
        count++
    }
}
/*		for i := 0; i< len(horas_led); i++{
			horas_led[i].Low()
			time.Sleep(time.Millisecond * 500)
	
			horas_led[i].High()
			time.Sleep(time.Millisecond * 500)
		} 
		
				for i := 0; i< len(horas_led); i++{
			horas_led[i].Low()
		} */