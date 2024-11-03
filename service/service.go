package service

import "parking_simulator/models"

type ParkingService struct {
    parkingLot *models.ParkingLot 
}

func NewParkingService(parkingLot *models.ParkingLot) *ParkingService {
    return &ParkingService{parkingLot: parkingLot}
}

func (ps *ParkingService) Arrive(vehicle *models.Vehicle) {
    ps.parkingLot.Arrive(vehicle)
}

func (ps *ParkingService) Depart(vehicle *models.Vehicle) {
    ps.parkingLot.Depart(vehicle)
}

func (ps *ParkingService) GetOccupiedSpaces() ([]bool, []int) {
    return ps.parkingLot.GetOccupiedSpaces()
}

func (ps *ParkingService) Capacity() int {
    return ps.parkingLot.Capacity()
}
