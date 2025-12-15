// fix numb error

package godocxNumLib

import (
	"fmt"
    "encoding/xml"

    "github.com/gomutex/godocx/docx"
)

type DocxLists struct {
	DLists []DocxList
}

type DocxList struct {
	Ord bool
	AbId int
	Mark [9]string
	Start [9]int
}

type numbering struct {
    XMLName xml.Name `xml:"numbering"`
    List []list  `xml:"abstractNum"`
    Numb []numb `xml:"num"`
	NMap map[int]int
}

type numb struct {
    XMLName xml.Name `xml:"num"`
    NumId int `xml:"numId,attr"`
    AbstNumId abstNum `xml:"abstractNumId"`
}

type abstNum struct {
    XMLName xml.Name `xml:"abstractNumId"`
    Val int `xml:"val,attr"`
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
    Val int `xml:"val,attr"`
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

	nlen :=len(numObj.Numb)
	nMap := make(map[int]int)
    for inum:=0; inum<nlen; inum++ {
        nb := numObj.Numb[inum]
//		nMap[nb.AbstNumId.Val] = nb.NumId -1
		nMap[nb.NumId -1] = nb.AbstNumId.Val
//        fmt.Printf("  Numb: %d Id: %d Abst Id: %d\n",inum, nb.NumId, nb.AbstNumId.Val)
    }

	numObj.NMap = nMap

	return numObj, nil
}

func (num *numbering)PrintNumObj() {

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
            fmt.Printf("      start name: %s val: %d\n", level.Start.XMLName.Local, level.Start.Val)
            fmt.Printf("      numFmt name: %s val: %s\n", level.NumFmt.XMLName.Local, level.NumFmt.Val)
            fmt.Printf("      lvlTxt name: %s val: %v\n", level.LvlText.XMLName.Local, level.LvlText.Val)
        }
    }

    fmt.Println("*** Num ****")

	for abs, nm := range num.NMap {
		fmt.Printf(" abs: %d numid: %d\n", abs, nm)
	}

	fmt.Println("*** end of PrintList ***")
}

func (num *numbering) CreNList() (ML DocxLists, err error) {

//	fmt.Printf("lists: %d\n", len(num.List))

	ML.DLists = make([]DocxList, len(num.List))

	for ni:=1; ni<  len(num.List)+1; ni++ {
		an:= num.NMap[ni-1]
		an1 := num.List[an].AbstNumId
//	fmt.Printf("%d: abs num: %d %d num: %d\n", ni, an, an1, ni)
		dl:=ML.DLists[ni-1]
		dl.AbId = an1
		dl.Ord = true
//	fmt.Printf("ord: %t, abnum: %d\n", dl.Ord, dl.AbId)
		if num.List[an].Lvl[0].NumFmt.Val == "bullet" {dl.Ord = false}
//		fmt.Printf(" numFmt: %s\n", num.List[ni-1].Lvl[0].NumFmt.Val)
		for il:=0; il<9; il++ {
			dl.Mark[il] = num.List[an].Lvl[il].NumFmt.Val
			dl.Start[il] = num.List[an].Lvl[il].Start.Val
		}
		ML.DLists[ni-1] = dl
//	fmt.Printf("ord: %t, abnum: %d\n", dl.Ord, dl.AbId)
	}

    return ML, nil
}

func (DL DocxLists) PrintDocxList() {

	fmt.Printf("**** DocxLists: %d ****\n", len(DL.DLists))

	for i:=0; i< len(DL.DLists); i++ {
		dl := DL.DLists[i]
		fmt.Printf("  *** DL: %d ***\n",i+1)
		fmt.Printf("   order: %t\n", dl.Ord)
		fmt.Printf("   Abs Id: %d\n", dl.AbId)

		for il:=0; il< 9; il++ {
			fmt.Printf("    level: %d\n", il)
			fmt.Printf("      mark:  %s\n", dl.Mark[il])
			fmt.Printf("      start: %d\n", dl.Start[il])

		}
	}
}
