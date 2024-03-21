
## Chapter 18. Hands on excercise #3 - go tour step 4-7

함수 파라미터에 같은 타입이 들어간다면 타입 선언을 생략할 수 있다.

```go
func add(x int, y int) int {
	return x + y
}
```

```go
func add(x, y int) int {
	return x + y
}
```

Golang은 동시에 여러 값을 반환할 수 있다.

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

결과 값에 이름을 지정할 수 있다. 별도로 값을 리턴하지 않아도 된다. `naked return`이라고 한다.

- 리턴 값의 이름이 x, y 이다.
- return 키워드 뒤에 별도로 값을 명시하지 않는다.

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```
