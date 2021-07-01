package main

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"strings"

	//"unicode"

	//"io/ioutil"
	"os"
)

type brokers struct {
	ParticipantID int   `csv:"Participant ID"`
	BrokerNos string `csv:"Broker No."`
	Name string `csv:"Participant Name in Chinese"`
}

func main(){
	//resp, e := http.Get("https://www.hkex.com.hk/CHI/PLW/csv/List_of_Current_SEHK_EP.CSV")
	//if e != nil {
	//	panic(e)
	//}
	//if resp == nil {
	//	panic("resp nil")
	//}
	//
	//if resp.Body == nil {
	//	panic("body nil")
	//}

	// utf-8
	//file, _ := os.Open("/Users/jinersong/Downloads/List_of_Current_SEHK_EP.CSV")

	// utf-16 le
	file, _ := os.Open("/Users/jinersong/go/src/go_test/csv/List_of_Current_SEHK_EP.CSV")

	//var decoder = mahonia.NewDecoder("UTF-16LE")
	buf, e := ioutil.ReadAll(file)
	if e != nil {
		panic(e)
	}
	//reader := bufio.NewReader(file)
	//for {
	//	str, err := reader.ReadString(' ') //读到一个换行就结束
	//	if err == io.EOF { //io.EOF表示文件的末尾
	//		break
	//	}
	//	s := decoder.ConvertString(str)
	//	//输出内容
	//	fmt.Println("UTF-8 to GBK:", s)
	//}

	//ds := strings.Split(string(buf), "\n")
	//fmt.Println("ds", ds)


	lines := strings.Split(string(buf), "\n")
	for i, v := range lines {
		if i == 0 {
			continue
		}
		//t, _ := iconv.ConvertString(v, "ANSI", "utf8")
		items := strings.Split(strings.TrimSpace(v), "\t")
		//zhHK , _ := iconv.ConvertString(items[3], "ANSI", "utf8")
		s1 := CharSetClean(items[3])
		zhHK := Convert2GBK(s1)
		//nameHK := ConvertToString(items[3], "UTF-16LE", "UTF-8")

		nameHK, err := Unicode(bytes.NewBufferString(items[3]).Bytes())
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s %s %s %s \n", items[0], items[1], nameHK, zhHK)
	}

}

func Convert2GBK(str string) string {
	//return mahonia.GetCharset(str).NewDecoder().ConvertString(str)
	return mahonia.NewDecoder("UTF-16LE").ConvertString(str)
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func CharSetClean(name string) string {
	old := string([]byte{0x00})
	str1 := strings.ReplaceAll(name, old, "")
	str2 := strings.ReplaceAll(str1, "\"", "")
	return str2
}

func Unicode(bs []byte) (string, error) {
	//unicode.Is(unicode.Han)
	decoder := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	bs2, err := decoder.Bytes(bs)
	return string(bs2), err
}