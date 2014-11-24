package main
import 
(
 "testing"
  "strconv"
 "fmt"
 _ "io/ioutil"
)
var testStringMap map[string]string
var testIntMap map[int]int

var keyIntMap map[int]int
var keyStringMap map[int]string

func Fill(n int) {

//bs := []byte(strconv.Itoa(n))
//ioutil.WriteFile("/tmp/output.txt", bs, 0644)

testStringMap=make(map[string]string)
testIntMap=make(map[int]int)

keyIntMap=make(map[int]int)
keyStringMap=make(map[int]string)



var str string
var key string
var keyInt int
//var num int

tempMapsKey:=0
c:=9900
if n > c{

	n=c
}

for i := 0; i < n; i++ {
	str=str+"c"
	key=key+"k"
	keyInt=(i+keyInt)*10


	testStringMap[key]=str
	testIntMap[keyInt]=(keyInt+6)
	//if i>4000 {
		keyIntMap[tempMapsKey]=i
		keyStringMap[tempMapsKey]=key
		tempMapsKey++
	//}
}
}


func BenchmarkString(b *testing.B) {
//if b.N>999999900{
b.ResetTimer()
	Fill(b.N)
	var z string
	for _,k:= range keyStringMap {
		z=testStringMap[k]
	}
	fmt.Sprintf(z)
}
//}



func BenchmarkString(b *testing.B) {
//if b.N>999999900{
b.ResetTimer()
	Fill(b.N)
        var z int
        for _,k:= range keyIntMap {
                z=testIntMap[k]
        }
        fmt.Printf(strconv.Itoa(z))
}
//}





