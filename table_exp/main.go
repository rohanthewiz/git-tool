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

var animalFields = []string{
	"Name", "Type", "Color", "Weight",
}

var headers = []string{"Name", "Type", "Color", "Weight"}

var animals = []Animal{
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
}

func main() {
	var bindings []binding.DataMap

	// Create a binding for each animal data
	for i := 0; i < len(animals); i++ {
		bindings = append(bindings, binding.BindStruct(&animals[i]))
	}

	ap := app.New()
	wn := ap.NewWindow("Table Widget")

	tbl := widget.NewTable(
		// Dimensions
		func() (int, int) {
			return len(animals) + 1, len(animalFields) // + 1 row for a hdr
		},
		// Default
		func() fyne.CanvasObject {
			return widget.NewLabel(" - ")
		},
		// Set Values
		func(cell widget.TableCellID, cnvObj fyne.CanvasObject) {
			// for no binding just SetText --> cnvObj.(*widget.Label).SetText(data[i.Row][i.Col])
			if cell.Row == 0 { // header row
				label := cnvObj.(*widget.Label)
				label.Alignment = fyne.TextAlignCenter
				label.TextStyle = fyne.TextStyle{Bold: true}
				label.SetText(headers[cell.Col])
				return
			}

			datum, err := getTableDatum(cell, bindings)
			if err != nil {
				return
			}
			cnvObj.(*widget.Label).Bind(datum.(binding.String))
		})

	// Handler
	tbl.OnSelected = func(cell widget.TableCellID) {
		if cell.Row == 0 && cell.Col >= 0 && cell.Col < len(headers) { // valid hdr cell
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
	refWidth := widget.NewLabel("establish width").MinSize().Width
	tbl.SetColumnWidth(0, refWidth)
	tbl.SetColumnWidth(1, refWidth*2/3)
	tbl.SetColumnWidth(2, refWidth)
	tbl.SetColumnWidth(3, refWidth*2/3)

	wn.SetContent(container.NewMax(tbl))
	wn.Resize(fyne.Size{Width: 500, Height: 450})
	wn.ShowAndRun()
}

func getTableDatum(cell widget.TableCellID, bindings []binding.DataMap,
) (datum binding.DataItem, err error) {
	// Bounds check
	if cell.Row < 0 || cell.Row > len(bindings) { // hdr is first row
		msg := "No data binding for row"
		log.Println(msg, cell.Row)
		return datum, rerr.NewRErr(msg)
	}
	if cell.Col < 0 || cell.Col > len(animalFields)-1 {
		return datum, rerr.NewRErr(fmt.Sprintf("Column %d is out of bounds", cell.Col))
	}

	// Get the data binding for the row
	bndg := bindings[cell.Row-1]

	datum, err = bndg.GetItem(animalFields[cell.Col])
	if err != nil {
		log.Println(rerr.StringFromErr(rerr.Wrap(err)))
		return datum, rerr.Wrap(err)
	}
	return
}
