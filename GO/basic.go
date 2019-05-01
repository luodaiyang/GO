package main

import (
	"fmt"
	"math/cmplx"
	"math"
)


var test="test"
//test:="trdf"  函数外面不行



var(
	sdadd="afgda"
	aad=9
)

func  variable()  {
	var a int
	var s string
	fmt.Printf("%d,%q\n",a,s)
}

func variableinit()  {
	var a, b int=3, 4
	var s  string="abs"
	fmt.Println(a,b,s)
}

func variabletype()  {
	var a,b,c,d=3,4,"sbc",true
	b=5
	fmt.Println(a,b,c,d)
}
func euler()  {
	//欧拉公式
	c:=3+4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi+1))
	fmt.Println(cmplx.Exp(1i*math.Pi)+1)
	fmt.Printf("%.3f\n",cmplx.Exp(1i*math.Pi)+1)
}
func variableshort()  {
	 a,b,c,d:=3,4,"sbc",true
	fmt.Println(a,b,c,d)
}
//强制类型转换
func triangle(){
	var a,b  int  = 3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)

}
func main() {
	fmt.Println("hello")
	variable()
	variableinit()
	variabletype()
	variableshort()
	fmt.Println(test,sdadd,aad)
	euler()
	triangle()
}
/**
/hello
0,""
3 4 abs
3 5 sbc true
3 4 sbc true
test afgda 9
5
(-2.718281828459045+3.3289351404027797e-16i)
(0+1.2246467991473515e-16i)
(0.000+0.000i)
5
 */