package main

import "fmt"

func main() {
	//제어문
	//IF 문 : 반드시 Boolean 검사 -> 1, 0 (사용 불가 : 자동 형 변환 불가)
	//소괄호 사용 X

	var a int = 20
	b := 20

	//예제1
	if a >= 15 {

	}
	fmt.Println("15이상")

	//예제2
	if b >= 25 {
		fmt.Println("25이상")
	}

	//에러 발생1(괄호 줄바꿈 X)
	//if b >= 25
	//{
	//
	//}

	//에러 발생2(괄호 생략 x)
	//if b >= 25
	//	fmt.Println("25이상")

	//에러 발생3(1은 Boolean형이 아님)
	//if c:= 1; c {
	//	fmt.Println("True")
	//}

	//예제3
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}

}
