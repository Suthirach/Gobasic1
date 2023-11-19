package main

import (
	"fmt"
)

func main() {
	// ตัวแปรแบบชนิดข้อมูล
	// var name1 string = "TANJANTHUEK"
	// var age int = 25
	// var score float32 = 95.33
	// var ispass bool = true

	// fmt.Println("Hello Go wolrd")
	// fmt.Println("let's Go")
	// fmt.Println("my name is =", name)
	// fmt.Println("last name =", name1)
	// fmt.Println("age =", age)
	// fmt.Println("score =", score)
	// fmt.Println(ispass)

	// ตัวแปรแบบกระชับ
	// name := "SUTHIRACH"
	// lastname := "TANJANTHUEK"
	// age := 25

	// fmt.Println(name)
	// fmt.Println(lastname)
	// fmt.Println(age)

	//  ตัวแปรคงที่
	// const name string = "Supavady"
	// fmt.Println(name)

	// name := "suthirach"
	// // lastname := "tanjanthuek"
	// age := 25
	// score := 34.55
	// ispass := true

	// การแทนที่ค่าด้วย %v และการเช็คtypeด้วย %T
	// การใช้ Printf มันจะไม่เว้นบรรทัดให้ ต้องใช้ \n เพื่อเว้นบรรทัด
	// fmt.Printf("name is %v\n", name) //name is suthirach
	// fmt.Printf("name is %T\n", name) //name is string

	// fmt.Printf("name is %v\n", age) //name is 25
	// fmt.Printf("name is %T\n", age) //name is int

	// fmt.Printf("name is %v\n", score) //name is 34.55
	// fmt.Printf("name is %T\n", score) //name is float64

	// fmt.Printf("name is %v\n", ispass) //name is true
	// fmt.Printf("name is %T\n", ispass) //name is bool

	// เครื่องหมาย คณิต
	// number1 := 10
	// number2 := 2

	// fmt.Println(number1 + number2) //12
	// fmt.Println(number1 - number2) //8
	// fmt.Println(number1 * number2) //20
	// fmt.Println(number1 / number2) //5
	// fmt.Println(number1 % number2) //0

	// //เครื่องเทียบเทียบ
	// fmt.Println(number1 == number2) //false
	// fmt.Println(number1 != number2) //true
	// fmt.Println(number1 > number2)  //true
	// fmt.Println(number1 < number2)  //false
	// fmt.Println(number1 >= number2) //trun
	// fmt.Println(number1 <= number2) //false

	// //รับค่าผ่านทางคีย์บอร์ดผู็ใช้งาน Scanf
	// // %s %d %f
	// var name string // ข้อความ
	// fmt.Print("ป้อนชื่อตัวเองเร็วเข้า= ")
	// fmt.Scanf("%s", &name)

	// fmt.Printf("สวัสดี %v", name) // สวัสดี ....
	// fmt.Println("สวัสดี ", name)  // สวัสดี ....

	// var score int //ตัวเลขจำนวน เต็ม
	// fmt.Print("ป้แนคะแนน = ")
	// fmt.Scanf("%d", &score)

	// fmt.Printf("คะแนน %v", score+10) // คะแนน ....
	// fmt.Println("คะแนน ", score+10)  // คะแนน ....

	// var score2 float64 // ทศนิยม
	// fmt.Print("ป้อนคะแนน = ")
	// fmt.Scanf("%f", &score2)

	// fmt.Printf("คะแนน %v", score2+10) // คะแนน ....
	// fmt.Println("คะแนน ", score2+10)  // คะแนน ....

	// การใช้ if else
	// var score3 int
	// fmt.Print("ป้อนคะแนนนักเรียน = ")
	// fmt.Scanf("%d", &score3)
	// fmt.Println("คะแนนสอบคือ = ", score3)

	// if score3 >= 50 {
	// 	fmt.Println("สอบผ่าน")
	// } else {
	// 	fmt.Println("ไม่ผ่าน")
	// } // ป้อนคะแนนนักเรียน = 1
	// คะแนนสอบคือ =  1
	// ไม่ผ่าน

	// Switch..Case

	// var services int
	// fmt.Print("ป้อนหมายเลขบริการ")
	// fmt.Scanf("%d", &services)

	// switch services {
	// case 1:
	// 	fmt.Println("เปิดบัญชีใหม่")
	// case 2:
	// 	fmt.Println("ฝาก-ถอน")
	// default:
	// 	fmt.Println("ข้อมูลไม่ถูกต้อง")
	// }

	//array
	// full
	var Sname [4]int = [4]int{10, 20, 30, 40}
	// shot
	Sname2 := [2]string{"สุทธิราช", "สุภาวดี"}
	//
	fmt.Println(Sname[1], Sname2[1])

	//การเข้าถึงArray ยัดค่าลงในArray
	var numbers [3]int
	numbers[0] = 100
	numbers[1] = 100
	numbers[2] = 500
	fmt.Println(numbers)

	// เช็คค่าArray ว่ามีกี่คน len()
	number3 := [3]int{100, 200, 300}
	size := len(number3)
	fmt.Println("Size of Array = ", size)

	// เพิ่มขนาดArray ด้วย ... เท่าไหร่ก็ได้
	number4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	size2 := len(number4)
	fmt.Println("Size of Array = ", size2)

	//Slices คุณสมบัติเหมือนArrayทุกอย่าง แต่คุณสมบัติเพิ่มเติมคือ เพิ่มสมาชิกได้
	name := []string{"tiy", "uoo"}
	name = append(name, "roo") //เพิ่มสมาชิก
	name = append(name, "yoo")
	name[1] = "youu" //แทนค่าสมาชิก
	fmt.Println(name)
	//string{"tiy" , "uoo" , "roo" , "yoo"} การเรียกใช้แบบเป็นช่วง
	// fmt.Println(name[:1]) //[tiy]
	// fmt.Println(name[1:]) //[uoo roo yoo]
	fmt.Println(name[:]) //[tiy youu roo yoo]

	// map
	country := map[string]string{}
	country["LA"] = "LostAngerlise"
	country["JP"] = "japan"
	fmt.Println(country) //map[JP:japan LA:LostAngerlise]

	country2 := map[int]string{}
	country2[01] = "LostAngerlise"
	country2[02] = "japan"
	fmt.Println(country2) //map[1:LostAngerlise 2:japan]

	country3 := map[string]int{}
	country3["LA"] = 02
	country3["JP"] = 01
	fmt.Println(country3) //map[JP:1 LA:2]

	//แบบบรรทัดเดี่ยว
	country4 := map[int]string{
		01: "thai",
		02: "japan",
	}
	fmt.Println(country4[01]) // thai

	//เช็คkey ว่ามีไหม

	value, check := country4[01]

	if check {
		fmt.Println(value)
	} else {
		fmt.Println("ไม่พบเว้ย")
	}
	//for loop == loopของgo มีแค่ for loopเท่านั้น

	// for i := 1; i <= 3; i++ {
	// 	fmt.Println("สวัสดีชาวGO") // สวัสดีชาวGO   3 ครั้ง
	// }

	// คำสั้ง loop = break:หยุดการทำงาน
	//				continue:หยุดการทำงานแล้ววนสู่ค่าเริ่มต้นอันต่อไป หรือกระโดดข้าม รอบที่กำหนดไว้
	// for n := 1; n <= 5; n++ {
	// 	// return

	// 	fmt.Printf("สวัสดีชาวGO %v\n", n) // สวัสดีชาวGO   3 ครั้ง
	// }

	for i := 1; i <= 10; i++ {

		if i == 5 {
			break
		}
		// nub := 0
		fmt.Println(i, "สวัสดีชาวGO ") // สวัสดีชาวGO   3 ครั้ง
		// สวัสดีชาวGO   3 ครั้ง

	} // 1234

	for i := 1; i <= 10; i++ {

		if i == 5 {
			continue
		}
		// nub := 0
		fmt.Println(i, "สวัสดีชาวGO ") // สวัสดีชาวGO   3 ครั้ง
		// สวัสดีชาวGO   3 ครั้ง

	} // 1234678910รอบที่50จะไม่ทำงาน

	// Range
	nubs := []int{10, 20, 30, 40, 50, 60, 70, 90, 100}
	for index := 0; index < len(nubs); index++ {
		// fmt.Println(nubs[index])
		fmt.Println("index = ", index, "value = ", nubs[index])
	} // เลขทั้งหมดที่เป็น ค่า nubs 10,20,30,40,50,60,70,80,90,100
	fmt.Println(nubs[:]) // [10 20 30 40 50 60 70 90 100]

	language := map[string]string{"th": "thai", "en": "Eng", "jp": "japan"}
	for _, value := range language {
		fmt.Println("Value = ", value)
	}
}
