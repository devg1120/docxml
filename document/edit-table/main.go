// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"fmt"
	"log"
	"time"
       // "encoding/xml"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
	"reflect"
      "github.com/davecgh/go-spew/spew"

)

func add_row(doc *document.Document, table_name string, num int) {
	
	tables := []document.Table{}
	for _, t := range doc.Tables() {
		tables = append(tables, t)
	}

	//fmt.Println(len(tables))
	for _, t := range tables {
//	for _, t := range doc.Tables() {
		tp := t.Properties()
	        //fmt.Printf("%s\n", *tp.X().TblCaption)
		var name wml.CT_String = *tp.X().TblCaption
		//n := *tp.X().TblCaption
                //fmt.Println(reflect.TypeOf(n))  

		fmt.Printf("%s\n", name.ValAttr)

		//fmt.Printf("%s\n", *tp.X().TblDescription)
		//CT_String *table_name := *tp.X().TblCaption
	        if table_name == name.ValAttr {
	              fmt.Printf("tablename match\n")
	              fmt.Printf("rows %d\n", len(t.Rows()))
		      //t.AddRow()

		      rs := t.Rows()
		      //var row document.Row

		      row  := rs[0]
		      t.AppendRow(row)
		      row  = rs[1]
		      t.AppendRow(row)
/*
		      r1 := t.AddRow()

		      for x := 0;  x < 7 ;x++ {
		         c := r1.AddCell()
		   	 c.AddParagraph()
		      }
*/
		      r2 := t.AddRow()
		      rl := t.GetRow(1)
		      //fmt.Print(rl)
                      spew.Dump(rl.Cells()[0])

		      rp := rl.Properties()
                      r2.SetProperties(rp)

		    for _, cell := range rl.Cells() {
			 //fmt.Printf("%s\n",cell)
		         pro := cell.Properties()
			 //fmt.Printf("%#v\n",pro)
		         c := r2.AddCell()
			 c.SetProperties( pro)
		   	 c.AddParagraph()
	            }
                      //spew.Dump(r2.Cells()[0])

		    /*
		      for x := 0;  x < 7 ;x++ {
		         c := r2.AddCell()
		   	 c.AddParagraph()
		      }
*/
/*
type CellProperties struct {
	x *wml.CT_TcPr
}

type CellBorders struct {
	x *wml.CT_TcBorders
}
*/
	              fmt.Printf("rows %d\n", len(t.Rows()))
		}
	}

}


func main() {
	doc, err := document.Open("document.docx")
        fmt.Println(reflect.TypeOf(doc))  

	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	// This sample document uses structured document tags, which are not common
	// except for in document templates.  Normally you can just iterate over the
	// document's paragraphs.
	for _, sdt := range doc.StructuredDocumentTags() {
		for _, p := range sdt.Paragraphs() {
			paragraphs = append(paragraphs, p)
		}
	}

	for _, p := range paragraphs {
		for _, r := range p.Runs() {
			switch r.Text() {
			case "FIRST NAME":
				// ClearContent clears both text and line breaks within a run,
				// so we need to add the line break back
				r.ClearContent()
				r.AddText("John ")
				r.AddBreak()

				para := doc.InsertParagraphBefore(p)
				para.AddRun().AddText("Mr.")
				para.SetStyle("Name") // Name is a default style in this template file

				para = doc.InsertParagraphAfter(p)
				para.AddRun().AddText("III")
				para.SetStyle("Name")

			case "LAST NAME":
				r.ClearContent()
				r.AddText("Smith")
			case "Address | Phone | Email":
				r.ClearContent()
				r.AddText("111 Rustic Rd | 123-456-7890 | jsmith@smith.com")
			case "Date":
				r.ClearContent()
				r.AddText(time.Now().Format("Jan 2, 2006"))
			case "Recipient Name":
				r.ClearContent()
				r.AddText("Mrs. Smith")
				r.AddBreak()
			case "Title":
				// we remove the title content entirely
				p.RemoveRun(r)
			case "Company":
				r.ClearContent()
				r.AddText("Smith Enterprises")
				r.AddBreak()
			case "Address":
				r.ClearContent()
				r.AddText("112 Rustic Rd")
				r.AddBreak()
			case "City, ST ZIP Code":
				r.ClearContent()
				r.AddText("San Francisco, CA 94016")
				r.AddBreak()
			case "Dear Recipient:":
				r.ClearContent()
				r.AddText("Dear Mrs. Smith:")
				r.AddBreak()
			case "Your Name":
				r.ClearContent()
				r.AddText("John Smith")
				r.AddBreak()

				run := p.InsertRunBefore(r)
				run.AddText("---Before----")
				run.AddBreak()
				run = p.InsertRunAfter(r)
				run.AddText("---After----")

			default:
				fmt.Println("not modifying", r.Text())
			}
		}
	}

	
	tables := []document.Table{}
	for _, t := range doc.Tables() {
		tables = append(tables, t)
	}

	fmt.Println(len(tables))
	for _, t := range tables {
		tp := t.Properties()
	        //fmt.Printf("%+v\n", tp.X())
	        //fmt.Printf("%s\n", *tp.X().TblCaption)
	        fmt.Printf("%s\n", *tp.X().TblCaption)
		//fmt.Printf("%s\n", *tp.X().TblDescription)
		for _, row := range t.Rows() {
	          fmt.Println("row")
		    for _, cell := range row.Cells() {
	               fmt.Println("cell")
		       for _ , p := range cell.Paragraphs() {
	                        fmt.Println("paragraph")
				if len(p.Runs()) == 0 {
	                         fmt.Println("runs == 0")
				   p.AddRun()
				   p.Runs()[0].AddText("abc")
                  
				} else {
		                    for _, r := range p.Runs() {
	                                 fmt.Println("runs")
			                 r.ClearContent()
			                 r.AddText("John ")
			                 //r.AddBreak()
			            }
		               }
		       }

	            }
		}
	}
	
	add_row(doc, "Test-Table-2",3)


	doc.SaveToFile("edit-document.docx")
	//
}