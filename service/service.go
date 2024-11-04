package service

import "parking_simulator/models"

type ParkingServiceHandler struct {
    parkingStructure *models.ParkingStructure
}

func NewParkingServiceHandler(parkingStructure *models.ParkingStructure) *ParkingServiceHandler {
    return &ParkingServiceHandler{parkingStructure: parkingStructure}
}

func (handler *ParkingServiceHandler) RegisterArrival(vehicle *models.Vehicle) {
    handler.parkingStructure.RegisterArrival(vehicle)
}

func (handler *ParkingServiceHandler) RegisterDeparture(vehicle *models.Vehicle) {
    handler.parkingStructure.RegisterDeparture(vehicle)
}

func (handler *ParkingServiceHandler) GetOccupiedSlots() ([]bool, []int) {
    return handler.parkingStructure.GetOccupiedSlots()
}

func (handler *ParkingServiceHandler) TotalCapacity() int {
    return handler.parkingStructure.TotalCapacity()
}
