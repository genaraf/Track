package move

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

const (
	Motor_A_EN = 4
	Motor_B_EN = 17

	Motor_A_Pin1 = 26
	Motor_A_Pin2 = 21
	Motor_B_Pin1 = 27
	Motor_B_Pin2 = 18
)

type Direction int

const (
	STOP     Direction = 0
	FORWARD  Direction = 1
	BACKWARD Direction = 2
)

type Turn int

const (
	LEFT   Turn = -1
	DIRECT Turn = 0
	RIGHT  Turn = 1
)

var (
	aEN   gpio.PinIO
	aPin1 gpio.PinIO
	/*
		aEN   = rpio.Pin(Motor_A_EN)
		bEN   = rpio.Pin(Motor_B_EN)
		aPin1 = rpio.Pin(Motor_A_Pin1)
		aPin2 = rpio.Pin(Motor_A_Pin2)
		bPin1 = rpio.Pin(Motor_B_Pin1)
		bPin2 = rpio.Pin(Motor_B_Pin2)
	*/
	motorPwm = 2000
)

func left(dir Direction, speed float32) {
	if dir == FORWARD {

	} else if dir == BACKWARD {

	} else {

	}
}

func right(dir Direction, speed float32) {
	if dir == FORWARD {

	} else if dir == BACKWARD {

	} else {

	}
}

func Init() {
	aEN = gpioreg.ByName("GPIO4")
	if aEN == nil {
		log.Fatal("Failed to find GPIO4")
	}
	aEN.Out(gpio.Low)
	fmt.Printf("%s: %s\n", aEN, aEN.Function())

	/*
		aEN.Pwm()
		aEN.Freq(motorPwm)
		aPin1.Output()
		aPin2.Output()

		bEN.Pwm()
		bEN.Freq(motorPwm)
		bPin1.Output()
		bPin2.Output()
	*/
	Stop()
}

func Stop() {
	/*
		aEN.Low()
		aPin1.Low()
		aPin2.Low()

		bEN.Low()
		bPin1.Low()
		bPin2.Low()
	*/
}

// 0 < radius <= 1
func Move(speed float32, dir Direction, turn Turn, radius float32) {
	if dir == FORWARD {
		if turn == LEFT {
			left(BACKWARD, speed*radius)
			right(FORWARD, speed)
		} else if turn == RIGHT {
			left(BACKWARD, speed*radius)
			right(FORWARD, speed)
		} else {
			left(FORWARD, speed)
			right(FORWARD, speed)

		}
	} else if dir == BACKWARD {
		if turn == LEFT {
			left(BACKWARD, speed*radius)
			right(FORWARD, speed)
		} else if turn == RIGHT {
			left(BACKWARD, speed*radius)
			right(FORWARD, speed)
		} else {
			left(BACKWARD, speed)
			right(BACKWARD, speed)
		}
	} else {
		Stop()
	}
}
