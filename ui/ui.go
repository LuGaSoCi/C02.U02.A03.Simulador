package ui

import (
    "fmt"
    "math/rand"
    "parking_simulator/models"
    "parking_simulator/service"
    "time"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
    "image/color"
)

type StyledLabel struct {
    widget.Label
}

func NewStyledLabel(text string) *StyledLabel {
    label := &StyledLabel{}
    label.ExtendBaseWidget(label)
    label.SetText(text)
    label.Alignment = fyne.TextAlignCenter
    label.TextStyle = fyne.TextStyle{Bold: true}

    return label
}

func (label *StyledLabel) MinSize() fyne.Size {
    return fyne.NewSize(250, 20)
}

func GenerateWindow(app fyne.App, parkingServiceHandler *service.ParkingServiceHandler, duration float64, carCount int) fyne.Window {
    window := app.NewWindow("Parking Simulator")

    background := canvas.NewImageFromFile("./assets/background.jpg")
    background.FillMode = canvas.ImageFillStretch

    parkingSlots := make([]*fyne.Container, parkingServiceHandler.TotalCapacity())
    slotTexts := make([]*canvas.Text, parkingServiceHandler.TotalCapacity())
    vehicleImages := make([]*canvas.Image, parkingServiceHandler.TotalCapacity())

    layoutGrid := container.New(layout.NewGridLayoutWithColumns(2))

    for i := 0; i < parkingServiceHandler.TotalCapacity(); i++ {
        slotTexts[i] = canvas.NewText(fmt.Sprintf("Slot %d: Available", i+1), theme.ForegroundColor())
        slotTexts[i].TextSize = 15
        slotTexts[i].Alignment = fyne.TextAlignCenter

        vehicleImages[i] = canvas.NewImageFromFile("./assets/car.png")
        vehicleImages[i].Resize(fyne.NewSize(75, 30))
        vehicleImages[i].Hide()

        slotContainer := container.NewWithoutLayout(
            slotTexts[i],
            vehicleImages[i],
        )

        slotTexts[i].Move(fyne.NewPos(100, 40))
        vehicleImages[i].Move(fyne.NewPos(70, 40))

        parkingSlots[i] = container.New(
            layout.NewPaddedLayout(),
            slotContainer,
        )

        layoutGrid.Add(parkingSlots[i])
    }

    contentContainer := container.NewStack(
        background,
        container.NewPadded(layoutGrid),
    )

    scrollContainer := container.NewScroll(contentContainer)

    window.SetContent(scrollContainer)
    window.Resize(fyne.NewSize(500, 900))
    window.CenterOnScreen()
    window.Show()

    go refreshDisplay(slotTexts, vehicleImages, parkingServiceHandler)

    go simulateVehicleFlow(carCount, duration, parkingServiceHandler)

    return window
}

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

func simulateVehicleFlow(carCount int, duration float64, parkingServiceHandler *service.ParkingServiceHandler) {
    for i := 1; i <= carCount; i++ {
        time.Sleep(time.Duration(rand.ExpFloat64() * duration) * time.Millisecond)
        vehicle := &models.Vehicle{ID: i}
        go parkingServiceHandler.RegisterArrival(vehicle)
    }
}
