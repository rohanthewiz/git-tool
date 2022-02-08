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

var animals = []Animal{
	{Name: "Frisky", Type: "cat", Color: "gray", Weight: "10"},
	{Name: "Ella", Type: "dog", Color: "brown", Weight: "50"},
	{Name: "Mickey", Type: "mouse", Color: "black", Weight: "1"},
}

type ColAttr struct {
	Name         string
	Header       string
	WidthPercent int
}

var animalCols = []ColAttr{
	{Name: "Name", Header: "Name", WidthPercent: 100},
	{Name: "Type", Header: "Type", WidthPercent: 67},
	{Name: "Color", Header: "Color", WidthPercent: 100},
	{Name: "Weight", Header: "Weight", WidthPercent: 67},
}

var animalBindings []binding.DataMap

// Create a binding for each animal data
func init() {
	for i := 0; i < len(animals); i++ {
		animalBindings = append(animalBindings, binding.BindStruct(&animals[i]))
	}
}

func main() {
	ap := app.New()
	wn := ap.NewWindow("Table Widget")

	tbl := createTable(animalBindings)
	setTblCellSelectHandler(tbl)

	// Layout
	refWidth := widget.NewLabel("establish width").MinSize().Width

	for i := 0; i < len(animalCols); i++ {
		tbl.SetColumnWidth(i, float32(animalCols[i].WidthPercent)/100.0*refWidth)
	}

	wn.SetContent(container.NewMax(tbl))
	wn.Resize(fyne.Size{Width: 500, Height: 450})
	wn.ShowAndRun()
}

func setTblCellSelectHandler(tbl *widget.Table) {
	tbl.OnSelected = func(cell widget.TableCellID) {
		if cell.Row == 0 && cell.Col >= 0 && cell.Col < len(animalCols) { // valid hdr cell
			fmt.Println("-->", animalCols[cell.Col].Header)
			return
		}

		str, err := getStrCellValue(cell, animalBindings)
		if err != nil {
			fmt.Println(rerr.StringFromErr(err))
			return
		}
		fmt.Println("-->", str)
	}
}

func createTable(bndg []binding.DataMap) *widget.Table {
	return widget.NewTable(
		// Dimensions (rows, cols)
		func() (int, int) {
			return len(animals) + 1, len(animalCols) // + 1 row for a hdr
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
				label.SetText(animalCols[cell.Col].Header)
				return
			}

			datum, err := getTableDatum(cell, bndg)
			if err != nil {
				fmt.Println(rerr.StringFromErr(err))
				return
			}
			cnvObj.(*widget.Label).Bind(datum.(binding.String))
		})
}

func getStrCellValue(cell widget.TableCellID, bdngs []binding.DataMap) (str string, err error) {
	datum, err := getTableDatum(cell, bdngs)
	if err != nil {
		return str, rerr.Wrap(err)
	}

	str, err = datum.(binding.String).Get()
	if err != nil {
		return str, rerr.Wrap(err)
	}
	return
}

func getTableDatum(cell widget.TableCellID, bindings []binding.DataMap,
) (datum binding.DataItem, err error) {
	// Bounds check
	if cell.Row < 0 || cell.Row > len(bindings) { // hdr is first row
		msg := "No data binding for row"
		log.Println(msg, cell.Row)
		return datum, rerr.NewRErr(msg)
	}
	if cell.Col < 0 || cell.Col > len(animalCols)-1 {
		return datum, rerr.NewRErr(fmt.Sprintf("Column %d is out of bounds", cell.Col))
	}

	// Get the data binding for the row
	bndg := bindings[cell.Row-1]

	datum, err = bndg.GetItem(animalCols[cell.Col].Name)
	if err != nil {
		log.Println(rerr.StringFromErr(rerr.Wrap(err)))
		return datum, rerr.Wrap(err)
	}
	return
}
