package main

import mediator "p16/mediator"

func main() {
	stationManager := mediator.NewStationManger()

    passengerTrain := &mediator.PassengerTrain{
        Mediator: stationManager,
    }
    freightTrain := &mediator.FreightTrain{
        Mediator: stationManager,
    }

    passengerTrain.Arrive()
    freightTrain.Arrive()
    passengerTrain.Depart()
}