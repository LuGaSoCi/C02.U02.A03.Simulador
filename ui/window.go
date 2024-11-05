package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "parking_simulator/service"
    "fyne.io/fyne/v2/theme"
    "fmt"
)

func GenerateWindow(app fyne.App, parkingServiceHandler *service.ParkingServiceHandler, duration float64, carCount int) fyne.Window {
    window := app.NewWindow("Parking Simulator")

    background := canvas.NewImageFromFile("./assets/background.jpg")
    background.FillMode = canvas.ImageFillStretch

    layoutGrid, slotTexts, vehicleImages := createParkingSlotsGrid(parkingServiceHandler.TotalCapacity())

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

func createParkingSlotsGrid(totalSlots int) (*fyne.Container, []*canvas.Text, []*canvas.Image) {
    layoutGrid := container.New(layout.NewGridLayoutWithColumns(2))
    slotTexts := make([]*canvas.Text, totalSlots)
    vehicleImages := make([]*canvas.Image, totalSlots)

    for i := 0; i < totalSlots; i++ {
        slotTexts[i] = canvas.NewText(fmt.Sprintf("Slot %d: Available", i+1), theme.ForegroundColor())
        slotTexts[i].TextSize = 15
        slotTexts[i].Alignment = fyne.TextAlignCenter
        slotTexts[i].Move(fyne.NewPos(100, 40)) // Posición específica para los textos de los slots

        vehicleImages[i] = canvas.NewImageFromFile("./assets/car.png")
        vehicleImages[i].Resize(fyne.NewSize(75, 30))
        vehicleImages[i].Move(fyne.NewPos(70, 40)) // Posición específica para las imágenes de los vehículos
        vehicleImages[i].Hide() // Ocultamos la imagen por defecto

        // Contenedor para el slot y la imagen del vehículo
        slotContainer := container.NewWithoutLayout(
            slotTexts[i],
            vehicleImages[i],
        )

        parkingSlot := container.New(
            layout.NewPaddedLayout(),
            slotContainer,
        )

        layoutGrid.Add(parkingSlot)
    }

    return layoutGrid, slotTexts, vehicleImages
}

