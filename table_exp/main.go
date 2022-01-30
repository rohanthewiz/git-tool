package main

// Code by Rohan Allison
import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/rohanthewiz/rerr"
)

type Animal struct {
	Name, Type, Color, Weight string
}

const nbrAnimalAttrs = 4

var animals = []Animal{
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
}
var headers = []string{"Name", "Type", "Color", "Weight"}

func main() {
	var bindings []binding.DataMap
	// Create a binding for each animal
	for i := 0; i < len(animals); i++ {
		bindings = append(bindings, binding.BindStruct(&animals[i]))
	}

	ap := app.New()
	wn := ap.NewWindow("Table Widget")

	tbl := widget.NewTable(
		// dimensions
		func() (int, int) {
			return len(animals) + 1, nbrAnimalAttrs // + 1 row for a hdr
		},
		// default
		func() fyne.CanvasObject {
			return widget.NewLabel(" - ")

		},
		// bindings
		func(cell widget.TableCellID, co fyne.CanvasObject) {
			// for no binding --> co.(*widget.Label).SetText(data[i.Row][i.Col])
			if cell.Row == 0 { // header row
				label := co.(*widget.Label)
				label.Alignment = fyne.TextAlignCenter
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.SetText(headers[cell.Col])
				return
			}

			datum, err := getTableDatum(cell, bindings)
			if err != nil {
				return
			}
			co.(*widget.Label).Bind(datum.(binding.String))
		})

	refWidth := widget.NewLabel("reasonable width").MinSize().Width

	tbl.SetColumnWidth(0, refWidth)
	tbl.SetColumnWidth(1, refWidth*2/3)
	tbl.SetColumnWidth(2, refWidth)
	tbl.SetColumnWidth(3, refWidth*2/3)

	// Handler
	tbl.OnSelected = func(cell widget.TableCellID) {
		if cell.Row == 0 && cell.Col < len(headers) {
			fmt.Println("-->", headers[cell.Col])
			return
		}

		datum, err := getTableDatum(cell, bindings)
		if err != nil {
			log.Println(err)
			return
		}
		str, er := datum.(binding.String).Get()
		if er != nil {
			log.Println(err)
			return
		}
		fmt.Println("-->", str)
	}

	// Layout
	wn.SetContent(container.NewMax(tbl))
	wn.Resize(fyne.Size{Width: 500, Height: 450})
	wn.ShowAndRun()
}

func getTableDatum(cell widget.TableCellID, bindings []binding.DataMap,
) (datum binding.DataItem, err error) {
	if cell.Row > len(bindings) { // hdr is extra row
		msg := "No data binding for row"
		log.Println(msg, cell.Row)
		return datum, rerr.NewRErr(msg)
	}
	row := bindings[cell.Row-1] // first row is header
	switch cell.Col {
	case 0:
		datum, err = row.GetItem("Name")
	case 1:
		datum, err = row.GetItem("Type")
	case 2:
		datum, err = row.GetItem("Color")
	case 3:
		datum, err = row.GetItem("Weight")
	}
	if err != nil {
		log.Println(rerr.StringFromErr(rerr.Wrap(err)))
		return datum, rerr.Wrap(err)
	}
	return
}
