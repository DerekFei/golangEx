package main

import "html/template"
import "log"
import "os"

type hotel struct {
	Name, Address, City, Zip, Region string
}
type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := Regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Hilton",
					Address: "1 asdfasdf road",
					City:    "Atlanta",
					Zip:     "30308",
					Region:  "Southern",
				},
				hotel{
					Name:    "Milton",
					Address: "2 asdfasdf road",
					City:    "Btlanta",
					Zip:     "20308",
					Region:  "Southern",
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Hilton",
					Address: "1 asdfasdf road",
					City:    "Atlanta",
					Zip:     "30308EE",
					Region:  "Southern",
				},
				hotel{
					Name:    "Milton",
					Address: "2 asdfasdf road",
					City:    "Btlanta",
					Zip:     "20308444",
					Region:  "Southern",
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
