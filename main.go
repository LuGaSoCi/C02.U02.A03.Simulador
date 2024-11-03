package main

import (
    "parking_simulator/models"
    "parking_simulator/service"
    "parking_simulator/ui"
    "fyne.io/fyne/v2/app"
)

func main() {
    myApp := app.New()
    parkingLot := models.NewParkingLot(20)
    parkingService := service.NewParkingService(parkingLot)
    
    ui.CreateWindow(myApp, parkingService, 250, 100)
    myApp.Run()
}
