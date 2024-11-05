package ui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    //"fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
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

func NewVehicleImage(filePath string) *canvas.Image {
    image := canvas.NewImageFromFile(filePath)
    image.Resize(fyne.NewSize(75, 30))
    image.Hide() // Ocultar por defecto
    return image
}
