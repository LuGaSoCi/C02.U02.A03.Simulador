package main

import (
    "parking_simulator/models"
    "parking_simulator/service"
    "parking_simulator/ui"
    "fyne.io/fyne/v2/app"
)

func main() {
    simulatorApp := app.New()
    parkingStructure := models.NewParkingStructure(20)
    parkingServiceHandler := service.NewParkingServiceHandler(parkingStructure)
    
    ui.GenerateWindow(simulatorApp, parkingServiceHandler, 250, 100)
    simulatorApp.Run()
}
