package main

/*
	Copyright © 2024 Mateus Felipe da Silveira Vieira
    This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

    You should have received a copy of the GNU General Public License along with this program. If not, see <https://www.gnu.org/licenses/>. 
*/

import (
    "machine"
    "time"

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
		b_arr := to_binary(7, 0x10)

		for i := 0; i < len(b_arr); i++ {

			if b_arr[i] == 1 {
				horas_led[i].High()
			}  
			//println( b_arr[i] )
		}
		println("")

		b_arr = to_binary(31, 0x20)
		for i := 0; i < len(b_arr); i++ {

			if b_arr[i] == 1 {
				minutos_led[i].High()
			}  
			//println( b_arr[i] )
		}
		println(time.Now().UTC().String())
    }
}
//tinygo flash --port /dev/ttyUSB0 --target arduino-mega2560 -monitor  -baudrate 9600 trigger.go 