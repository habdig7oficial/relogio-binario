package main

import (
    "machine"
    "time"
)

func to_binary(value uint){
	for ; value > 0;
	{
		print (value & 1)
		value >>= 1
	} 

	//for ;bits >= 0; bits-- {
	//	aux := value >> bits
	//	println("aux ", aux)
		//if ((aux) & 1){
		//	println(1)
		//} else {
		//	println(0)
		//}
	//}

}

func main() {
	horas_led := []machine.Pin{ machine.D53, machine.D52, machine.D51, machine.D50, machine.D49 }
	for i:=0; i < len(horas_led); i++ {
		horas_led[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}


    for {
		for i := 0; i< len(horas_led); i++{
			horas_led[i].Low()
			time.Sleep(time.Millisecond * 500)
	
			horas_led[i].High()
			time.Sleep(time.Millisecond * 500)
		} 

		to_binary(25)

		for i := 0; i< len(horas_led); i++{
			horas_led[i].Low()
		} 
    }
}
//tinygo flash --port /dev/ttyUSB0 --target arduino-mega2560 -monitor  -baudrate 9600 trigger.go 