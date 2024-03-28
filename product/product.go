package product

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/rohinibarge/productorder/api"
	"github.com/rohinibarge/productorder/httpclient"
)

var pl *api.ProductList

type Order struct {
	Url string
}

func (o Order) Execute() {
	if pl == nil {
		pl = &api.ProductList{}
	}
	o.Fetch(o.Url)
	o.Display()
}

func (o *Order) Fetch(url string) {
	client := httpclient.GetHttpClient()
	res, err := client.Get(url)
	if err != nil {
		fmt.Printf("some error occured: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf(" error occured wile read response: %v", err)
	}
	if err := json.Unmarshal(body, &pl); err != nil { // Parse []byte to the go struct pointer
		fmt.Printf("Can not unmarshal JSON, err: %v", err)
	}
}

func (o *Order) Display() {
	var tmplateFile = "products_html.tmpl"
	tmpl1, err := template.New(tmplateFile).Funcs(nil).ParseFiles(tmplateFile)
	if err != nil {
		panic(err)
	}

	var f *os.File
	f, err = os.Create("web/products.html")
	if err != nil {
		panic(err)
	}
	err = tmpl1.Execute(f, pl.Products)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
