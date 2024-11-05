package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ParkingStructure struct {
    maxCapacity      int
    currentVehicles  int
    accessMutex      sync.RWMutex
    vehicleQueue     chan *Vehicle
    openSlots        chan struct{}
    occupiedSlots    []bool
    vehicleSlotIDs   []int
    nextAvailableIdx int
}

func NewParkingStructure(maxCapacity int) *ParkingStructure {
    structure := &ParkingStructure{
        maxCapacity:     maxCapacity,
        vehicleQueue:    make(chan *Vehicle, maxCapacity),
        openSlots:       make(chan struct{}, maxCapacity),
        occupiedSlots:   make([]bool, maxCapacity),
        vehicleSlotIDs:  make([]int, maxCapacity),
        nextAvailableIdx: 0,
    }
    for i := 0; i < maxCapacity; i++ {
        structure.openSlots <- struct{}{}
    }
    return structure
}

func (structure *ParkingStructure) findNextSlot() int {
    for i := 0; i < structure.maxCapacity; i++ {
        idx := (structure.nextAvailableIdx + i) % structure.maxCapacity
        if !structure.occupiedSlots[idx] {
            structure.nextAvailableIdx = (idx + 1) % structure.maxCapacity
            return idx
        }
    }
    return -1
}

func (structure *ParkingStructure) RegisterArrival(vehicle *Vehicle) {
	select {
    case <-structure.openSlots:
        fmt.Printf("Vehículo %d ingresando al estacionamiento.\n", vehicle.ID)
        structure.accessMutex.Lock()
        slotIndex := structure.findNextSlot()
        if slotIndex != -1 {
            structure.occupiedSlots[slotIndex] = true
            structure.vehicleSlotIDs[slotIndex] = vehicle.ID
            fmt.Printf("Vehículo %d asignado al espacio %d.\n", vehicle.ID, slotIndex+1)
        }
        structure.accessMutex.Unlock()

        time.Sleep(time.Duration(3+rand.Intn(3)) * time.Second)
        structure.RegisterDeparture(vehicle)

    default:
        fmt.Printf("Vehículo %d esperando por un espacio.\n", vehicle.ID)
        structure.vehicleQueue <- vehicle
    }
}

func (structure *ParkingStructure) RegisterDeparture(vehicle *Vehicle) {
	structure.accessMutex.Lock()
    slotFound := false
    slotIndex := -1

    for i := 0; i < structure.maxCapacity; i++ {
        if structure.vehicleSlotIDs[i] == vehicle.ID {
            slotFound = true
            slotIndex = i
            structure.occupiedSlots[i] = false
            structure.vehicleSlotIDs[i] = 0
            break
        }
    }
    structure.accessMutex.Unlock()

    if slotFound {
        fmt.Printf("Vehículo %d saliendo del espacio %d.\n", vehicle.ID, slotIndex+1)
        structure.openSlots <- struct{}{}

        select {
        case nextVehicle := <-structure.vehicleQueue:
            go structure.RegisterArrival(nextVehicle)
        default:
        }
    }
}

func (structure *ParkingStructure) GetOccupiedSlots() ([]bool, []int) {
    structure.accessMutex.RLock()
    defer structure.accessMutex.RUnlock()
    slotsCopy := make([]bool, len(structure.occupiedSlots))
    idsCopy := make([]int, len(structure.vehicleSlotIDs))
    copy(slotsCopy, structure.occupiedSlots)
    copy(idsCopy, structure.vehicleSlotIDs)
    return slotsCopy, idsCopy
}

func (structure *ParkingStructure) TotalCapacity() int {
    return structure.maxCapacity
}

func (structure *ParkingStructure) GetVehicleSlotID(index int) int {
    structure.accessMutex.RLock()
    defer structure.accessMutex.RUnlock()
    return structure.vehicleSlotIDs[index]
}
