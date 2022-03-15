package main

import "log"

type rover struct {
	xCoordinate int
	yCoordinate int
	direction   string // better to be enum S,N,E & W
	finalX      int
	finalY      int
}

// scan input co-ordinate
// scan scan no. of rovers
// get starting co-ordinate,
// instruction text with L,M, or R
// split down instruction to individual characters
// check direction movement or not(L,R): face down
// find the next co-ordinate by passing current state and one instruction character
// display after the all iterations
func main() {
	processInput(1, 2, 5, 5, "N", "LMLMLMLMM")
}

func processInput(xCoordinate, yCoordinate, finalX, finalY int, direction string, commands string) {
	// initialize rover with co-ordinates
	roverOne := &rover{
		xCoordinate: xCoordinate,
		yCoordinate: yCoordinate,
		direction:   direction,
	}

	chars := []rune(commands)

	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		roverOne.getNextState(finalX, finalY, char)
	}

	log.Println("final co-ordinate", roverOne.xCoordinate, roverOne.yCoordinate, roverOne.direction)

}

// utility function to process each command character
func (r *rover) getNextState(finalX, finalY int, command string) {
	log.Println("before command", command, r.xCoordinate, r.yCoordinate, r.direction)
	switch command {
	case "L":
		// output direction dependes on existing face down direction
		switch r.direction {
		case "N":
			r.direction = "W"
		case "S":
			r.direction = "E"
		case "E":
			r.direction = "N"
		case "W":
			r.direction = "S"
		}
	case "R":
		switch r.direction {
		case "N":
			r.direction = "E"
		case "S":
			r.direction = "W"	
		case "E":
			r.direction = "S"
		case "W":
			r.direction = "N"

		}
	case "M":
		// decide to increase or decrease x or y co-ordinate
		switch r.direction {
		case "N":
			// check its in final Y point or not
			if r.yCoordinate != finalY {
				r.yCoordinate++
			}
		case "S":
			// lowest point
			if r.yCoordinate != 0 {
				r.yCoordinate--
			}
		case "E":
			if r.xCoordinate != finalX {
				r.xCoordinate++
			}
		case "W":
			if r.xCoordinate != 0 {
				r.xCoordinate--
			}
		}
	default:
		log.Println("Unknown command")

	}
	log.Println("final after", command, r.xCoordinate, r.yCoordinate, r.direction)
	return

}
