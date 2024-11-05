package ui

import (
    "fmt"
    "image/color"
    "parking_simulator/service"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/theme"
    "time"
)

func updateParkingSlots(slotTexts []*canvas.Text, vehicleImages []*canvas.Image, parkingServiceHandler *service.ParkingServiceHandler) {
    occupiedSlots, vehicleIDs := parkingServiceHandler.GetOccupiedSlots()

    for i := range occupiedSlots {
        if slotTexts[i] == nil || vehicleImages[i] == nil {
            continue
        }

        if occupiedSlots[i] {
            slotTexts[i].Text = fmt.Sprintf("\t\t\t#%d  ", vehicleIDs[i])
            slotTexts[i].Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
            vehicleImages[i].Show()
        } else {
            slotTexts[i].Text = ""
            slotTexts[i].Color = theme.SuccessColor()
            vehicleImages[i].Hide()
        }

        slotTexts[i].Refresh()
        vehicleImages[i].Refresh()
    }
}

func refreshDisplay(slotTexts []*canvas.Text, vehicleImages []*canvas.Image, parkingServiceHandler *service.ParkingServiceHandler) {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()

    for range ticker.C {
        updateParkingSlots(slotTexts, vehicleImages, parkingServiceHandler)
    }
}
