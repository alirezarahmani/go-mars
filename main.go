package main

import (
	"./Domain"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	argsWithProg := os.Args
	var lens = len(argsWithProg)
	lens--
	if lens >= 6 && ((lens-2)%4) == 0 {
		var plateauCoordinatesX, _ = strconv.Atoi(argsWithProg[1:][0])
		var plateauCoordinatesY, _ = strconv.Atoi(argsWithProg[1:][1])

		var plateau, _ = Domain.NewPlateau(plateauCoordinatesX, plateauCoordinatesY)

		// iterate over all rovers positions and headings
		for i := 0; i < lens; i =i+6 {
			var x, _ = strconv.Atoi(argsWithProg[1:][i+2])
			var y, _ = strconv.Atoi(argsWithProg[1:][i+3])
			// get arguments with order
			// i= X  i+1= Y i+2= Heading i+3= movements
			var position, _ = Domain.NewPosition(x, y)
			marsRover, err := Domain.NewMarsRover(plateau, position, Domain.DirectionType(argsWithProg[1:][i+4]))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			moves := strings.Split(argsWithProg[1:][i+5], "")
			/** @var RoverMove $move */
			for _, element := range moves {
				// index is the index where we are
				// element is the element from someSlice for where we are
				err := marsRover.Move(Domain.RoverMoveType(element))
				if err != nil {
					fmt.Println(err.Error())
					return 
				}
			}
			xout, yout := marsRover.ReportPosition()
			fmt.Println(xout, yout, marsRover.ReportHeading())
		}
	} else {
		fmt.Println("make sure you add right argument")
	}
}
