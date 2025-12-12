package godocxNumLib

import (
	"fmt"
    "encoding/xml"

    "github.com/gomutex/godocx/docx"
)

type numbering struct {
    XMLName xml.Name `xml:"numbering"`
    List []list  `xml:"abstractNum"`
    Numb numb `xml:"num"`
}

type numb struct {
    XMLName xml.Name `xml:"num"`
    NumId int `xml:"numId,attr"`
    AbstNumId int `xml:"abstractNumId,attr"`
}

type list struct {
    XMLName xml.Name `xml:"abstractNum"`
    AbstNumId int `xml:"abstractNumId,attr"`
    NsId nsid `xml:"nsid"`
    ML ml `xml:"multiLevelType"`
    Lvl []level `xml:"lvl"`
}

type nsid struct {
    XMLName xml.Name `xml:"nsid"`
    Val string `xml:"val,attr"`
}
type ml struct {
    XMLName xml.Name `xml:"multiLevelType"`
    Val string `xml:"val,attr"`
}

type level struct {
    XMLName xml.Name `xml:"lvl"`
    Ilvl string `xml:"ilvl,attr"`
    Start start `xml:"start"`
    NumFmt numFmt `xml:"numFmt"`
    LvlText lvlText `xml:"lvlText"`
}

type start struct {
    XMLName xml.Name `xml:"start"`
    Val string `xml:"val,attr"`
}

type numFmt struct {
    XMLName xml.Name `xml:"numFmt"`
    Val string `xml:"val,attr"`
}

type lvlText struct {
    XMLName xml.Name `xml:"lvlText"`
    Val string `xml:"val,attr"`
}


func GetNumObj(rdoc *docx.RootDoc)(numObj *numbering, err error) {

	filMap:= rdoc.FileMap

    contNumbering, ok := filMap.Load("word/numbering.xml")
    if !ok {return nil, fmt.Errorf("cannot load file 'word/numbering.xml'!/n")}

//    numObj:=&numbering{}

    err = xml.Unmarshal(contNumbering.([]byte), &numObj)
    if err != nil {return nil, fmt.Errorf("unmarshal: %v\n", err)}

	return numObj, nil
}

func PrintNumObj (num *numbering) {

    fmt.Println("*** numbering ****")
    fmt.Printf("Name: %s\n",num.XMLName.Local)

    fmt.Println("*** List ****")

    for il:=0; il<len(num.List); il++ {
        nL:= num.List[il]

        fmt.Printf("Name: %s Abs Num: %d\n",nL.XMLName.Local, nL.AbstNumId)
//  fmt.Printf("Abst Num: Id: %d\n", nL.AbstNumId)

        nsid := nL.NsId
        fmt.Println("  *** nsid ****")
        fmt.Printf("  Name: %s Value: %s\n",nsid.XMLName.Local, nsid.Val)
//  fmt.Printf("Value: %s\n",nsid.Val)

        ml := nL.ML
        fmt.Printf("  *** multiLevel ****")
        fmt.Printf("  Name: %s Value:%s\n",ml.XMLName.Local, ml.Val)
//  fmt.Printf("value: %s\n", ml.Val)

        fmt.Printf("  *** levels: %d ****\n", len(nL.Lvl))
        for i:=0; i< len(nL.Lvl); i++ {
            level := nL.Lvl[i]
            fmt.Printf("    *** level %d ***\n", i)
            fmt.Printf("      Name: %s Ilvl: %s\n",level.XMLName.Local, level.Ilvl)
//      fmt.Printf("  ilevel: %s\n", level.Ilvl)
            fmt.Printf("      start name: %s val: %s\n", level.Start.XMLName.Local, level.Start.Val)
            fmt.Printf("      numFmt name: %s val: %s\n", level.NumFmt.XMLName.Local, level.NumFmt.Val)
            fmt.Printf("      lvlTxt name: %s val: %s\n", level.LvlText.XMLName.Local, level.LvlText.Val)
        }
    }

    nb := num.Numb
    fmt.Println("*** Num ****")
    fmt.Printf("  Name: %s Id: %d Abst Id: %d\n",nb.XMLName.Local, nb.NumId, nb.AbstNumId)
//  fmt.Printf("Num Id: %d\n", nb.NumId)
//  fmt.Printf("Abst Num: Id: %d\n", nb.AbstNumId)
}
