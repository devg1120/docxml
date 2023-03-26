// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml/schema/soo/wml"
	//"fmt"
)

// Table is a table within a document.
type Table struct {
	d *Document
	x *wml.CT_Tbl
}

// X returns the inner wrapped XML type.
func (t Table) X() *wml.CT_Tbl {
	return t.x
}

// Properties returns the table properties.
func (t Table) Properties() TableProperties {
	if t.x.TblPr == nil {
		t.x.TblPr = wml.NewCT_TblPr()
	}
	return TableProperties{t.x.TblPr}
}

// AddRow adds a row to a table.
func (t Table) AddRow() Row {
	c := wml.NewEG_ContentRowContent()
	t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, c)
	tr := wml.NewCT_Row()
	c.Tr = append(c.Tr, tr)
	return Row{t.d, tr}
}

func (t Table) AppendRow( r Row)  {
	c := wml.NewEG_ContentRowContent()
	c.Tr = append(c.Tr, r.x)
	t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, c)
}

// GS append    
func (t Table) Rows() []Row {    
        ret := []Row{}    
        for _, c := range t.x.EG_ContentRowContent {    
             for _, tr := range c.Tr {      
                        ret = append(ret, Row{t.d, tr})        
             }    
        }    
        return ret    
}

func (t Table) GetRow(i  int) Row {    
         c :=  t.x.EG_ContentRowContent[i]
         //for _, tr := range c.Tr {      
	//	     fmt.Printf("   Tr:%+v\n",tr)
	// }
        return Row{t.d, c.Tr[0]}
/*
        for _, c := range t.x.EG_ContentRowContent {    
		fmt.Printf("row :%+v\n",c)
                       // return  Row{t.d, c.Tr[i]}
             for _, tr := range c.Tr {      
		     fmt.Printf("   Tr:%+v\n",tr)
	     }
        }    
        return Row{}
	*/
}
