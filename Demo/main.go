/**
 * Created with IntelliJ IDEA.
 * User: Administrator
 * Date: 13-7-19
 * Time: 上午10:03
 * To change this template use File | Settings | File Templates.
 */

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type GrowthsXml struct {
	XMLName     xml.Name    `xml:"growths"`
	GrowthNodes []GrowthXml `xml:"growth"`
	Content     string      `xml:",innerxml"`
}

type GrowthXml struct {
	XMLName     xml.Name `xml:"growth"`
	Occupation  uint8    `xml:"occupation,attr"`
	HpGrow      uint8    `xml:"hpgw,attr"`
	AttachGrow  uint8    `xml:"atkgw,attr"`
	DefenseGrow uint8    `xml:"defgw,attr"`
	CritGrow    uint8    `xml:"critgw,attr"`
	DodgeGrow   uint8    `xml:"dodgegw,attr"`
	HitGrow     uint8    `xml:"hit_rategw,attr"`
}

func main() {
	content, err := ioutil.ReadFile("growths.xml")
	if err != nil {
		fmt.Printf("ReadFile error: %v\n", err.Error())
		return
	}

	growList := new(GrowthsXml)
	err = xml.Unmarshal(content, growList)
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err.Error())
		return
	}

	growConfigMap := make(map[uint8]*GrowthXml)

	for _, v := range growList.GrowthNodes {
		fmt.Println("read node")
		growConfigMap[v.Occupation] = &v
		if growConfigMap[v.Occupation] == nil {
			fmt.Println("nil node.")
		}
	}

	//fmt.Println(growConfigMap[0].XMLName)
}
