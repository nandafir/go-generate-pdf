package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type VarHtml struct {
	StoreImage     string
	CourierImage   string
	CourierName    string
	Cod            string
	Assurance      string
	Weight         string
	Qty            string
	ListProducts   string
	BuyerName      string
	BuyerAddress   string
	StoreName      string
	StoreAddress   string
	Notes          string
	BarcodeResi    string
	BarcodeIdOrder string
}

func main() {

	goPDF()

}

func goPDF() {

	data := VarHtml{
		StoreImage:     "https://thumbs.dreamstime.com/b/supermarket-grocery-store-supermarket-grocery-store-supermarket-building-interior-fresh-food-shelves-counter-130146115.jpg",
		CourierImage:   "https://img.freepik.com/premium-vector/fast-delivery-truck-icon-set-fast-shipping-design-website-mobile-apps-online-shopping_97458-1032.jpg",
		CourierName:    "JNE",
		Cod:            "0",
		Assurance:      "0",
		Weight:         "200",
		Qty:            "3",
		ListProducts:   "Coffe (1), Rokok Malr (2), Baterai AA (1)",
		BuyerName:      "Jhon Cena",
		BuyerAddress:   "Jl.Sambeng gang mawar, Lamongan",
		StoreName:      "Toko Rodex",
		StoreAddress:   "Jl.Borobudur Gang Macan, Malang",
		Notes:          "Utara jembatan ploso",
		BarcodeResi:    "BCD-00001",
		BarcodeIdOrder: "ORD-10002",
	}

	htmlStr := GenerateHtml("templates/index.html", data) //filling variables to html_template
	//--------------------------------------------------------------------

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewBufferString(htmlStr)))
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA6) //paper size
	pdfg.PageHeight.Set(100)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./export/satu.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func GenerateHtml(layout string, data interface{}) string {
	result, _ := ParseTemplate(layout, data)

	return result
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}
