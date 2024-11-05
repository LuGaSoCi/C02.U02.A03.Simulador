package ui

import (
    "math/rand"
    "parking_simulator/models"
    "parking_simulator/service"
    "time"
)

func simulateVehicleFlow(carCount int, duration float64, parkingServiceHandler *service.ParkingServiceHandler) {
    for i := 1; i <= carCount; i++ {
        time.Sleep(time.Duration(rand.ExpFloat64() * duration) * time.Millisecond)
        vehicle := &models.Vehicle{ID: i}
        go parkingServiceHandler.RegisterArrival(vehicle)
    }
}
