package main

import (
    "machine"
    //"time"

)

func to_binary(value uint, mask uint) ([6]int){

	arr := [6]int{}

	for i := 0 ; mask > 0 ; i++{
		if ((value & mask) > 0){
			//print (1)
			arr[i] = 1
		} else {
			//print (0)
			arr[i] = 0
		}
		mask >>= 1
	}
	//println("")

	return arr
}

func main() {
	horas_led := []machine.Pin{ machine.D53, machine.D52, machine.D51, machine.D50, machine.D49 }
	minutos_led := []machine.Pin{ machine.D43, machine.D42, machine.D41, machine.D40, machine.D39, machine.D38,}

	for i:=0; i < len(horas_led); i++ {
		horas_led[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	for i:=0; i < len(minutos_led); i++ {
		minutos_led[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}


    for {
		//		b_arr := to_binary(10, 0x20)
		b_arr := to_binary(10, 0x10)

		for i := 0; i < len(b_arr); i++ {

			if b_arr[i] == 1 {
				horas_led[i].High()
			}  
			println( b_arr[i] )
		}
		println("")

		minutos_led[5].High()
    }
}
//tinygo flash --port /dev/ttyUSB0 --target arduino-mega2560 -monitor  -baudrate 9600 trigger.go 