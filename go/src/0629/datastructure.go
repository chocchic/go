package main

import "fmt"

func main() {
	// 배열 선언
	var ar1 [3]int
	fmt.Println(ar1)

	// 선언한 후 초기화
	var ar2 [3]string
	ar2 = [3]string{"dkfs", "게맛살", "커피"}
	fmt.Println(ar2)

	// 선언과 동시에 초기화
	ar3 := [3]float32{20.65, 20.1, 31.02}
	fmt.Println(ar3)

	// 반복문을 이용한 ar2배열의 모든 데이터 접근
	// len함수를 4번 호출
	for idx := 0; idx < len(ar2); idx = idx + 1 {
		temp := ar2[idx]
		fmt.Println(temp)
	}

	// len함수를 1번 호출
	size := len(ar2)
	for idx := 0; idx < size; idx = idx + 1 {
		temp := ar2[idx]
		fmt.Println(temp)
	}

	// 0번째 부터 2번째 앞까지
	fmt.Println(ar2[0:2])

	for idx, value := range ar2 {
		fmt.Printf("%dth %s\n", idx, value)
	}

	// 안 쓸 값은 _로 바꿈
	for _, value := range ar2 {
		fmt.Printf("%s\n", value)
	}

	// 배열 복제
	br := ar2
	br[0] = "ㅇㅁㄹㅇㄹ"
	fmt.Println(ar2)
	fmt.Println(br)
}
