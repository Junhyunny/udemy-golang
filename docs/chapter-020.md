
## Chapter 20. Hands on excercise #4 - go tour step 8-10

변수를 선언할 때 초기화(initialize)를 수행할 수 있다. 초기화를 함께 수행하면 타입을 생략해도 된다.

```go
	var c, python, java = true, false, "no!"
```

함수 안에서 := 연산자를 사용하면 var 키워드 없이 변수에 값을 할당할 수 있다. 함수 밖에선 모든 선언은 var, func 키워드로 시작해야된다. := 연산자를 사용할 수 없다.

```go
package main

import "fmt"

// error
// f := 3

func main() {
	var i, j int = 1, 2
	// ok
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```