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

type CustomLabel struct {
    widget.Label
}

func NewCustomLabel(text string) *CustomLabel {
    label := &CustomLabel{}
    label.ExtendBaseWidget(label)
    label.SetText(text)
    label.Alignment = fyne.TextAlignCenter
    label.TextStyle = fyne.TextStyle{Bold: true}

    return label
}

func (c *CustomLabel) MinSize() fyne.Size {
    return fyne.NewSize(250, 20)
}

func CreateWindow(app fyne.App, parkingService *service.ParkingService, duration float64, totalCars int) fyne.Window {
    myWindow := app.NewWindow("Simulación de Estacionamiento")

    background := canvas.NewImageFromFile("./assets/background.jpg")
    background.FillMode = canvas.ImageFillStretch

    containers := make([]*fyne.Container, parkingService.Capacity())
    texts := make([]*canvas.Text, parkingService.Capacity())
    carImages := make([]*canvas.Image, parkingService.Capacity())

    grid := container.New(layout.NewGridLayoutWithColumns(2))

    for i := 0; i < parkingService.Capacity(); i++ {
        texts[i] = canvas.NewText(fmt.Sprintf("Espacio %d: Libre", i+1), theme.ForegroundColor())
        texts[i].TextSize = 15
        texts[i].Alignment = fyne.TextAlignCenter

        carImages[i] = canvas.NewImageFromFile("./assets/car.png")
        carImages[i].Resize(fyne.NewSize(75, 30))
        carImages[i].Hide() 

        spotContainer := container.NewWithoutLayout(
            texts[i],
            carImages[i],
        )

        texts[i].Move(fyne.NewPos(100, 40))
        carImages[i].Move(fyne.NewPos(70, 40))

        containers[i] = container.New(
            layout.NewPaddedLayout(),
            spotContainer,
        )

        grid.Add(containers[i])
    }

    content := container.NewStack(
        background,
        container.NewPadded(grid),
    )

    scrollContainer := container.NewScroll(content)

    myWindow.SetContent(scrollContainer)
    myWindow.Resize(fyne.NewSize(500, 900))
    myWindow.CenterOnScreen()
    myWindow.Show()

    go updateScreenAndTicker(texts, carImages, parkingService)

    go generateCars(totalCars, duration, parkingService)

    return myWindow
}

func updateParkingDisplay(texts []*canvas.Text, carImages []*canvas.Image, parkingService *service.ParkingService) {
    occupiedSpaces, vehicleIDs := parkingService.GetOccupiedSpaces()

    for i := range occupiedSpaces {
        if texts[i] == nil || carImages[i] == nil {
            continue
        }

        if occupiedSpaces[i] {
            text := fmt.Sprintf("\t\t\t#%d  ", vehicleIDs[i])
            texts[i].Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Color rojo
            texts[i].Text = text
            carImages[i].Show()
        } else {
            texts[i].Text = fmt.Sprintf("")
            texts[i].Color = theme.SuccessColor() // Mantener este color para espacios libres
            carImages[i].Hide()
        }

        texts[i].Refresh()
        carImages[i].Refresh()
    }
}

func updateScreenAndTicker(texts []*canvas.Text, carImages []*canvas.Image, parkingService *service.ParkingService) {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()

    for range ticker.C {
        updateParkingDisplay(texts, carImages, parkingService)
    }
}

func generateCars(totalCars int, duration float64, parkingService *service.ParkingService) {
    for i := 1; i <= totalCars; i++ {
        time.Sleep(time.Duration(rand.ExpFloat64() * duration) * time.Millisecond)
        vehicle := &models.Vehicle{ID: i}
        go parkingService.Arrive(vehicle)
    }
}
