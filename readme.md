## Go 모듈 생성

go에서 모든 코드는 패키지 단위로 구성, 패키지명은 폴더명

```bash
go mod init 폴더명/패키지명
go mod init go/hello
```

## Go 빌드

GOOS와 GOARCH 환경변수를 사용하여 다른 플랫폼에서 실행할 수 있는 바이너리를 빌드할 수 있음

아래는 amd64 계열 칩셋의 리눅스에서 실행할 수 있는 바이너리를 빌드하는 예제

```bash
GOOS=linux GOARCH=amd64 go build
```

## package main

패키지 선언은 모든 파일의 첫 번째 줄에 위치해야 함. 이 코드가 속한 패키지의 이름을 나타냄
`main` 패키지는 실행 가능한 프로그램을 만들 때 사용. 이 패키지가 없는 경우 실행 파일을 만들 수 없음

```go
package main
```

## import 'fmt'

`import` 키워드를 사용하여 특정 패키지에서 제공하는 기능을 사용할 수 있음
fmt 패키지는 Go의 표준 라이브러리로, 표준 입출력 기능을 제공

```go
import "fmt"
```

## func main()

`func` 키워드를 사용하여 함수를 선언할 수 있음. main 함수는 프로그램의 시작점으로 사용되는 함수. main 함수가 종료되면 프로그램도 종료됨

```go
func main() {
    fmt.Println("Hello, World!")
}
```

## 주석 사용법

한 줄 주석은 `//` 를 사용하여 작성할 수 있음

```go
// 이 코드는 Hello, World!를 출력하는 프로그램입니다.
```

여러 줄 주석은 `/*` 와 `*/` 를 사용하여 작성할 수 있음

```go
/*
이 코드는
Hello, World!
를 출력하는 프로그램입니다.
*/
```

tip: godoc을 사용하여 코드에 작성된 주석을 문서화할 수 있음

## 변수 선언 및 초기화

변수는 `var` 키워드를 사용하여 선언할 수 있음. 변수 선언 시 타입을 명시해야 함. 변수를 선언할 때 초기값을 지정할 수 있음. **초기화**란 변수에 값을 할당하는 것을 의미

```go
var name type = expression
```

```go
name := expression
```



```go
var name string = "Alice" // 변수 선언 및 초기화
var age int // 변수 선언 (초깃값을 지정하지 않으면 타입의 기본값으로 초기화)
var height = 175 // 타입을 생략하면 초기값의 타입으로 변수 타입이 결정
weight := 70 // var 키워드를 생략하고 := 연산자를 사용하여 변수 선언 및 초기화
x, y := 1, 2
```

- var: 변수 선언 키워드
- name: 변수명
- string: 변수 타입
- "Alice": 변수 값

## 빈 식별자

공백 식별자는 `_(언더스코어)`를 사용하며 선언 시에 바인딩을 하지 않습니다

```go
_, err := fmt.Printf("Hello world")
if err != nil {
	// handle error
}
```



## 변수 값 변경

변수의 값을 변경할 때는 `=` 연산자를 사용

```go
name = "Bob"
```

## 변수명 컨벤션

변수명은 변숫값이 저장된 메모리 주소를 접근하는 데 사용되는 식별자. 변수명은 다음과 같은 규칙을 따라야 함

- 변수명은 문자, 숫자, 밑줄 기호(\_)로 구성
- 첫 글자는 문자나 밑줄 기호(\_)로 시작
- 영어 대소문자를 구분하고, 영문 뿐만 아니라 한글도 사용 가능
- 변수명이 한단어 이상인 경우 카멜 표기법을 사용하여 변수명을 작성(ex. var myName string)

## 변수의 타입

변수의 타입은 변수가 저장할 수 있는 값의 종류와 범위를 결정. 변수의 타입은 변수 선언 시 지정

```go
var age int = 20
```

## Go의 기본 타입

숫자 타입 외의 기본 타입에 대한 표

| 타입      | 설명                                                         | 예시                                        |
| --------- | ------------------------------------------------------------ | ------------------------------------------- |
| bool      | 불리언 타입으로 참, 거짓을 표현                              | true                                        |
| string    | 문자열 타입으로 문자들의 집합                                | "Hello, World!", `"Hello world"`            |
| array     | 배열 타입으로 같은 타입의 데이터를 순서대로 저장             | [1, 2, 3]                                   |
| slice     | 배열과 유사하지만 크기가 동적으로 변하는 배열                | []int{1, 2, 3}                              |
| map       | 키와 값으로 이루어진 데이터 구조                             | map[string]int{"Alice": 20, "Bob": 21}      |
| struct    | 필드들의 집합으로 구조체. 일반적으로 상관 관계의 데이터를 묶을 때 사용 | type Person struct { Name string; Age int } |
| pointer   | 메모리 주소를 저장하는 타입                                  | &name                                       |
| function  | 함수 타입                                                    | func add(a int, b int) int { return a + b } |
| channel   | 고루틴 간 통신을 위한 타입                                   | ch := make(chan int)                        |
| interface | 메서드 집합을 정의하는 타입                                  | type Stringer interface { String() string } |

### Go의 숫자 타입

타입별 이름과 값의 범위에 대한 표

숫자 타입이 여러가지인 이유는? 타입의 크기를 정확하게 지정하여 메모리를 효율적으로 사용하기 위함. 범위가 넓은 타입은 메모리를 많이 차지. 따라서 정확한 범위를 알고 있을 때는 타입을 정확하게 지정하여 메모리를 효율적으로 사용할 수 있음

- 정수의 기본 타입은 `int`
- 부동소수점의 기본 타입은 `float64`

| 타입       | 설명                                               | 범위                                       |
| ---------- | -------------------------------------------------- | ------------------------------------------ |
| uint8      | 부호 없는 8비트 정수                               | 0 ~ 255                                    |
| uint16     | 부호 없는 16비트 정수                              | 0 ~ 65535                                  |
| uint32     | 부호 없는 32비트 정수                              | 0 ~ 4294967295                             |
| uint64     | 부호 없는 64비트 정수                              | 0 ~ 18446744073709551615                   |
| int8       | 부호 있는 8비트 정수                               | -128 ~ 127                                 |
| int16      | 부호 있는 16비트 정수                              | -32768 ~ 32767                             |
| int32      | 부호 있는 32비트 정수                              | -2147483648 ~ 2147483647                   |
| int64      | 부호 있는 64비트 정수                              | -9223372036854775808 ~ 9223372036854775807 |
| float32    | 32비트 부동소수점                                  | IEEE-754 32비트                            |
| float64    | 64비트 부동소수점                                  | IEEE-754 64비트                            |
| complex64  | 32비트 실수부와 허수부의 복소수                    | 32비트 실수부, 32비트 허수부               |
| complex128 | 64비트 실수부와 허수부의 복소수                    | 64비트 실수부, 64비트 허수부               |
| byte       | uint8과 동일                                       | 0 ~ 255                                    |
| rune       | int32와 동일                                       | -2147483648 ~ 2147483647                   |
| int        | 32비트 시스템에서 int32, 64비트 시스템에서 int64   | 시스템 아키텍처에 따라 다름                |
| uint       | 32비트 시스템에서 uint32, 64비트 시스템에서 uint64 | 시스템 아키텍처에 따라 다름                |

### 타입별 기본값

변수 선언 시 초기값을 지정하지 않으면 타입별 기본값으로 초기화
`nil`은 정의되지 않은 값(메모리 주소)을 나타내는 Go의 특별한 값

| 타입            | 기본값 |
| --------------- | ------ |
| 정수 타입       | 0      |
| 부동소수점 타입 | 0.0    |
| bool            | false  |
| string          | ""     |
| 그 외           | nil    |



```go
var i int
fmt.Printf("%d", i) // 0
```



### 타입 변환

타입 변환은 서로 다른 타입의 값을 변환하는 것을 의미. 타입 변환은 다음과 같이 사용. 타입 변환이 필요한 이유는 서로 다른 타입의 값을 연산하거나 비교할 수 없기 때문

- 기본적으로 타입을 변환할 때는 `타입(변수)` 형태로 사용
- 타입 변환 시 값의 손실이 발생할 수 있음
- 실수 타입을 정수 타입으로 변환할 때는 소수점 이하를 버림
- 넓은 타입에서 좁은 타입으로 변환할 때는 값이 범위를 벗어날 수 있음

```go
var a int = 10
var b float64 = float64(a)
```

1234는 비트로 표현하면 00000100 11010010이고, int8 타입은 8비트로 표현되므로 11010010이 남음

```go
package main

import "fmt"

func main() {
	var a int16 = 1234 // int16 타입 변수 선언
	var c int8 = int8(a) // int16 타입 변수를 int8 타입으로 변환하여 변수 c에 대입 -> 값이 범위를 벗어나므로 값이 잘림

	fmt.Println(a, c) // 1234 -46
}
```

### 변수의 범위

변수의 범위는 변수가 사용 가능한 범위를 의미. 변수의 범위는 변수가 선언된 블록 내에서만 사용 가능. 여기서 블록이란 `{}`로 둘러싸인 코드 블록을 의미

- 패키지 전역 변수: 어떤 중괄호에도 속하지 않은 변수로 패키지 전체에서 사용 가능
- 지역 변수: 특정 중괄호에 속한 변수로 중괄호 내에서만 사용 가능

다행이도 이 모든 것을 기억할 필요는 없다. IDE에서 변수를 선언하면 해당 변수가 사용 가능한 범위를 표시해주고, 사용할 수 없는 범위에서 변수를 사용하면 컴파일 에러를 알려준다.

```go
package main

import "fmt"

func main() {
    var a int = 10 // 변수 a 선언
    {
        var b int = 20 // 변수 b 선언
        fmt.Println(a, b) // 10 20
    }
    fmt.Println(a, b) // 컴파일 에러 -> undefined: b compiler UndeclaredName
}
```

### 메모리에서 숫자의 표현

메모리에서 양수와 음수를 표현하는 방법

- 부호 있는 정수: 2의 보수 표현법 사용
- 부호 없는 정수: 0부터 시작하는 숫자 표현법 사용

메모리에서 정수와 실수를 표현하는 방법

- 정수: 메모리에 정수 값 그대로 저장
- 실수: **IEEE-754** 표준에 따라 부동소수점 표현

실수는 정수와 달리 소수점 이하 값을 표현하기 위해 부동소수점 방식을 사용. 부동소수점 방식은 소수점의 위치를 고정하지 않고 소수점의 위치를 나타내는 지수부와 실수부로 나누어 표현

예를 들어 123.456을 32비트 부동소수점으로 표현하면 다음과 같음

- 부호(1비트): 0(양수)
- 지수부(8비트): 01111100(124)
- 실수부(23비트): 11110110010011001100110(0.456)

부동소수점의 한계: 0.1을 32비트 부동소수점으로 표현하면 0.100000001490116119384765625가 됨. 이는 0.1과 약간 다른 값이므로 부동소수점 연산 시 주의해야 함

정확한 수치 연산이 필요한 경우: `math/big` 패키지의 `big.Float` 타입을 사용하여 정확한 수치 연산을 수행할 수 있음

## fmt 패키지

fmt 패키지는 Go의 표준 라이브러리로, 표준 입출력 기능을 제공. fmt 패키지는 다음과 같은 기능을 제공

- Print: 문자열을 출력
- Printf: 형식화된 문자열을 출력
- Println: 문자열을 출력하고 개행

```go
package main

import "fmt"

func main() {
    fmt.Print("Hello, ") // 문자열 출력
    fmt.Print("World!") // 문자열 출력
    fmt.Println() // 개행
    fmt.Println("Hello, World!") // 문자열 출력하고 개행
    fmt.Printf("Hello, %s!\n", "World") // 형식화된 문자열 출력
    fmt.Printf("Hello, %d!\n", 10) // 형식화된 문자열 출력
    fmt.Printf("Hello, %f!\n", 3.14) // 형식화된 문자열 출력
}
```

### Printf의 형식 지정자

Printf 함수는 형식 지정자를 사용하여 형식화된 문자열을 출력. 형식 지정자는 출력할 값의 타입을 지정하는 문자열

| 형식 지정자 | 설명                                                  | 예시                  |
| ----------- | ----------------------------------------------------- | --------------------- |
| %d          | 10진수 정숫값                                         | 10                    |
| %b          | 2진수                                                 | 2진수로 변환된 값     |
| %o          | 8진수                                                 | 8진수로 변환된 값     |
| %x          | 16진수                                                | 16진수로 변환된 값    |
| %f          | 부동소수점을 실숫값 형태로 출력                       | 3.14                  |
| %e          | 부동소수점을 지수 형태로 출력                         | 3.14e+00              |
| %g          | %e와 %f 중 짧은 형태로 출력                           | 3.14                  |
| %s          | 문자열                                                | "Hello, World!"       |
| %c          | 유니코드 문자                                         | 'A'                   |
| %t          | 불리언                                                | true                  |
| %p          | 포인터                                                | 메모리 주소           |
| %v          | 값의 기본 형태                                        | 값의 기본 형태 출력   |
| %T          | 값의 데이터 타입                                      | 값의 데이터 타입 출력 |
| %q          | 특수 문자의 기능을 사용하지 않고 문자열을 그대로 출력 | "Hello, World!"       |

### 출력 형식 지정

- 최소 출력 너비 지정: `%너비d` ex. `%5d` -> 5자리로 출력
- 정렬 방법 지정: `%너비-정렬d` ex. `%-5d` -> 왼쪽 정렬
- 빈 공간을 0으로 채우기: `%0너비d` ex. `%05d` -> 5자리로 출력하고 빈 공간을 0으로 채움
- 소수점 이하 자릿수 지정: `%너비.자릿수f` ex. `%5.2f` -> 5자리로 출력하고 소수점 이하 2자리까지 출력

```go
package main

import "fmt"

func main() {
    fmt.Printf("%5d\n", 10) // 10
    fmt.Printf("%-5d\n", 10) // 10
    fmt.Printf("%05d\n", 10) // 00010
    fmt.Printf("%5.2f\n", 3.14) // 3.14
}
```

### 특수 문자

특수 문자는 문자열에 포함된 특수한 문자를 출력할 때 사용. 특수 문자는 `\`로 시작하는 문자열

| 특수 문자 | 설명        |
| --------- | ----------- |
| \n        | 개행        |
| \t        | 탭          |
| \r        | 캐리지 리턴 |
| \b        | 백스페이스  |
| \f        | 폼 피드     |
| \\'       | 작은 따옴표 |
| \\"       | 큰 따옴표   |
| \\        | 역슬래시    |

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!") // Hello, World!
    fmt.Println("Hello, \nWorld!") // Hello,
                                    // World!
    fmt.Println("Hello, \tWorld!") // Hello,    World!
    fmt.Println("Hello, \\World!") // Hello, \World!
}
```

### 표준 입력 - Scan

표준 입력은 키보드로 입력한 값을 프로그램으로 전달하는 방법. 표준 입력은 `fmt` 패키지의 `Scan` 함수를 사용하여 구현

- `Scan`: 표준 입력으로부터 값을 읽어 변수에 저장
- `Scanf`: 형식화된 문자열을 표준 입력으로부터 읽어 변수에 저장
- `Scanln`: 표준 입력으로부터 값을 읽어 변수에 저장하고 개행

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("이름을 입력하세요: ")
    fmt.Scanln(&name) // 표준 입력으로부터 값을 읽어 변수에 저장
    fmt.Printf("안녕하세요, %s님!\n", name)
}
```

여러 개 한 번에 입력받기

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("이름과 나이를 입력하세요: ")
    fmt.Scanln(&name, &age) // 표준 입력으로부터 값을 읽어 변수에 저장
    fmt.Printf("안녕하세요, %s님! 나이는 %d살이군요!\n", name, age)
}
```

표준 입력에서 에러를 처리하는 방법

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    var name string
    var age int
    fmt.Print("이름과 나이를 입력하세요: ")
    _, err := fmt.Scanln(&name, &age) // 표준 입력으로부터 값을 읽어 변수에 저장
    if err != nil {
        fmt.Fprintln(os.Stderr, err) // 에러 출력
        return
    }
    fmt.Printf("안녕하세요, %s님! 나이는 %d살이군요!\n", name, age)
}
```

### 표준입력 - Scanf

`Scanf` 함수는 형식화된 문자열을 표준 입력으로부터 읽어 변수에 저장. `Scanf` 함수는 형식 지정자를 사용하여 입력받을 값의 타입을 지정

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("이름과 나이를 입력하세요: ")
    fmt.Scanf("%s %d", &name, &age) // 형식화된 문자열을 표준 입력으로부터 읽어 변수에 저장
    fmt.Printf("안녕하세요, %s님! 나이는 %d살이군요!\n", name, age)
}
```

### 표준입력 - Scanln

`Scanln` 함수는 표준 입력으로부터 값을 읽어 변수에 저장하고 개행. `Scanln` 함수는 공백을 기준으로 값을 구분.Scan()과 다른 점은 개행 문자를 만나면 입력을 종료한다는 점(즉 마지막 입력을 받고 엔터를 누르면 입력이 종료됨)

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("이름과 나이를 입력하세요: ")
    fmt.Scanln(&name, &age) // 표준 입력으로부터 값을 읽어 변수에 저장하고 개행
    fmt.Printf("안녕하세요, %s님! 나이는 %d살이군요!\n", name, age)
}
```

### 입력 스트림

입력 스트림은 표준 입력으로부터 값을 읽어오는 스트림. 입력 스트림은 `bufio` 패키지의 `NewReader` 함수를 사용하여 생성

여러 번 Scanf 함수를 호출할 경우 생길 수 있는 문제가 있다. 입력 스트림은 한 번 읽은 값을 버퍼에 저장하고, 다시 읽을 때 버퍼에 저장된 값을 사용. 따라서 버퍼에 저장된 값이 없으면 새로 입력을 받음. 입력 스트림에서 사용되지 않은 값은 버퍼에 남아 있음. 이 때문에 입력 스트림을 사용할 때는 버퍼를 비워주는 작업이 필요

`func (b *Reader) ReadString(\n byte) (string, error)`: 입력 스트림에서 개행 문자를 만날 때까지 읽어 문자열로 반환

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin) // 표준 입력 스트림 생성
    fmt.Print("이름을 입력하세요: ")
    var name string
    n, err := fmt.Scanln(&name) // 표준 입력으로부터 값을 읽어 변수에 저장
    if err != nil {
        fmt.Fprintln(os.Stderr, err) // 에러 출력
        stdin.ReadString('\n') // 입력 스트림 버퍼 비우기
    } else if n == 0 {
        fmt.Fprintln(os.Stderr, "입력된 값이 없습니다.") // 에러 출력
    } else {
        fmt.Printf("안녕하세요, %s님!\n", name)
    }

    n, err = fmt.Scanln(&name) // 표준 입력으로부터 값을 읽어 변수에 저장
    if err != nil {
        fmt.Fprintln(os.Stderr, err) // 에러 출력
    } else if n == 0 {
        fmt.Fprintln(os.Stderr, "입력된 값이 없습니다.") // 에러 출력
    } else {
        fmt.Printf("안녕하세요, %s님!\n", name)
    }
}
```

## 연산자

### 산술 연산자

- 산술 연산자: 덧셈(+), 뺄셈(-), 곱셈(\*), 나눗셈(/), 나머지(%)
- 비트 연산자: AND(&), OR(\|), XOR(^), AND NOT(&^), 왼쪽 시프트(<<), 오른쪽 시프트(>>)

| 연산자 | 설명          | 피연산자 타입              |
| ------ | ------------- | -------------------------- |
| +      | 덧셈          | 정수, 실수, 복소수, 문자열 |
| -      | 뺄셈          | 정수, 실수, 복소수         |
| \*     | 곱셈          | 정수, 실수, 복소수         |
| /      | 나눗셈        | 정수, 실수, 복소수         |
| %      | 나머지        | 정수                       |
| &      | AND 연산      | 정수                       |
| \|     | OR 연산       | 정수                       |
| ^      | XOR 연산      | 정수                       |
| &^     | AND NOT 연산  | 정수                       |
| <<     | 왼쪽 시프트   | 정수 << 양의 정수          |
| >>     | 오른쪽 시프트 | 정수 >> 양의 정수          |

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3
    fmt.Println(a + b) // 13
    fmt.Println(a - b) // 7
    fmt.Println(a * b) // 30
    fmt.Println(a / b) // 3
    fmt.Println(a % b) // 1

    c := 0b1010
    d := 0b0011
    fmt.Printf("%b\n", c & d) // 10
    fmt.Printf("%b\n", c | d) // 11
    fmt.Printf("%b\n", c ^ d) // 1
    fmt.Printf("%b\n", c &^ d) // 1000
    fmt.Printf("%b\n", c << 1) // 10100
    fmt.Printf("%b\n", c >> 1) // 101
}
```

### 비교 연산자

비교 연산자는 두 개의 값을 비교하여 참(true) 또는 거짓(false)을 반환. 비교 연산자는 다음과 같은 연산자로 구성

| 연산자 | 설명        |
| ------ | ----------- |
| ==     | 같음        |
| !=     | 다름        |
| >      | 크다        |
| <      | 작다        |
| >=     | 크거나 같다 |
| <=     | 작거나 같다 |

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3
    fmt.Println(a == b) // false
    fmt.Println(a != b) // true
    fmt.Println(a > b) // true
    fmt.Println(a < b) // false
    fmt.Println(a >= b) // true
    fmt.Println(a <= b) // false
}
```

- 오버 플로우: 정수 타입의 최대값을 넘어가면 최소값부터 다시 시작하는 현상
  - ex. int8 타입의 최대값은 127, 최소값은 -128. 127에서 1을 더하면 -128이 됨
- 언더 플로우: 정수 타입의 최소값보다 작은 값이 되면 최대값부터 다시 시작하는 현상
  - ex. int8 타입의 최대값은 127, 최소값은 -128. -128에서 1을 빼면 127이 됨
- 실수간 비교: 실수는 부동소수점 방식으로 표현되기 때문에 정확한 비교가 어려움. 따라서 실수를 비교할 때는 오차를 고려해야 함
  - ex. 0.1 + 0.2 == 0.3은 false이다. 0.1 + 0.2는 0.30000000000000004이기 때문

### 논리 연산자

논리 연산자는 논리값을 연산하여 참(true) 또는 거짓(false)을 반환. 논리 연산자는 다음과 같은 연산자로 구성

| 연산자 | 설명     | 반환되는 값                                 |
| ------ | -------- | ------------------------------------------- |
| &&     | AND 연산 | 두 값이면 모두 참이면 true, 아니면 false    |
| \|\|   | OR 연산  | 두 값 중 하나라도 참이면 true, 아니면 false |
| !      | NOT 연산 | 값이 참이면 false, 거짓이면 true            |

```go

```

### 기타 연산자

- 대입 연산자: 변수에 값을 대입
  - `=`: 변수에 값을 대입
- 복수 대입 연산자: 여러 변수에 값을 대입
  - `a, b = 10, 20`: a에 10, b에 20을 대입
- 복합 대입 연산자: 변수에 연산을 수행한 후 값을 대입
  - `a += 10`: a에 10을 더한 후 a에 대입
  - `a -= 10`: a에 10을 뺀 후 a에 대입
- 증감 연산자: 변수의 값을 증가 또는 감소
  - `a++`: a의 값을 1 증가
  - `a--`: a의 값을 1 감소
- 포인터 연산자: 변수의 메모리 주소를 반환
  - `&a`: a의 메모리 주소를 반환
- 배열 요소 접근 연산자: 배열의 요소에 접근
  - `a[0]`: a의 첫 번째 요소에 접근
  - `a[0:3]`: a의 0부터 3까지의 요소에 접근
- 구조체 필드 접근 연산자: 구조체의 필드에 접근
  - `a.name`: a의 name 필드에 접근
- 채널 연산자: 채널에 값을 전송하거나 받음
  - `ch <- 10`: 10을 채널 ch에 전송
  - `a := <-ch`: 채널 ch로부터 값을 받아 a에 대입
- 메모리 주소 접근 연산자: 메모리 주소에 저장된 값을 반환
  - `*a`: a의 메모리 주소에 저장된 값을 반환
- 슬라이스 요소 접근 또는 가변 인수 전달 연산자: 슬라이스의 요소에 접근하거나 가변 인수를 전달
  - `a[0]`: a의 첫 번째 요소에 접근
  - `f(a...)`: 가변 인수 a를 함수 f에 전달

## 함수

### 함수 정의

함수는 특정 기능을 수행하는 코드 블록. 함수는 `func` 키워드를 사용하여 정의

함수의 구성 요소

1. 함수 키워드: 함수를 정의할 때 사용하는 키워드
2. 함수명: 함수의 이름. 함수를 호출할 때 사용
3. 매개변수: 함수를 호출할 때 전달하는 값
4. 반환 타입: 함수가 반환하는 값의 타입
5. 함수 본문: 함수의 기능을 구현하는 코드 블록

```go
func 함수명(매개변수) 반환타입 {
    // 함수 본문
}
```

```go
package main

import "fmt"

func hello() {
    fmt.Println("Hello, World!")
}

func main() {
    hello()
}
```

### 함수의 호출

- argument: 함수를 호출할 때 전달하는 값으로 인수라고 부른다.(호출하는 사용자 관점)
- parameter: 함수를 정의할 때 받는 값으로 매개변수라고 부른다.(함수가 정의되는 관점)

함수 호출에 사용되는 인수는 값이 복사되어 함수 내부로 전달. 따라서 함수 내부에서 인수의 값을 변경해도 함수 외부의 값은 변경되지 않음

```go
package main

import "fmt"

func add(a int, b int) int { // 매개변수로 a와 b를 받고 int 타입을 반환하는 add 함수 정의
    return a + b
}

func main() {
    result := add(10, 20) // 인수로 10과 20을 전달하여 add 함수 호출
    fmt.Println(result) // 30
}
```

### 함수의 반환

함수는 반환값을 가질 수 있음. 반환값은 함수가 호출된 곳으로 값을 전달하는 역할을 함. 반환값은 함수의 반환 타입을 명시하여 정의

여러 값을 반환할 때는 반환값을 괄호로 묶어서 반환

```go
package main

import "fmt"

func add(a int, b int) (int, int) { // int 타입 두 개를 반환하는 add 함수 정의
    return a + b, a - b
}

func main() {
    result1, result2 := add(10, 20) // add 함수 호출
    fmt.Println(result1, result2) // 30 -10
}
```

함수 선언부에서 반환값의 이름을 지정하면 함수 본문에서 반환값의 이름을 사용할 수 있음. 이 때 반환값의 이름은 함수 선언부에서만 사용 가능

```go
package main

import "fmt"

func add(a int, b int) (result1 int, result2 int) { // 반환값의 이름을 지정하여 반환값을 반환하는 add 함수 정의
    result1 = a + b
    result2 = a - b
    return
}

func main() {
    result1, result2 := add(10, 20) // add 함수 호출
    fmt.Println(result1, result2) // 30 -10
}
```

## 상수

go 언에서의 상수가 가능한 타입

- 정수, 실수, 복소수
- 문자열
- 룬(rune)
- 불리언

상수 선언 방법

```go
const 상수명 타입 = 값
```

```go
package main

import "fmt"

func main() {
    const name string = "Alice" // 문자열 상수 선언
    const age int = 20 // 정수 상수 선언
    const score float64 = 80.5 // 실수 상수 선언

    fmt.Println(name, age, score) // Alice 20 80.5
}
```

### 상수를 코드 값으로 사용하기

상수는 코드의 가독성을 높이고, 코드의 유지보수를 쉽게 함. 코드의 가독성을 높이기 위해 상수를 사용할 때는 상수명을 대문자로 작성하여 상수임을 명시

```go
package main

import "fmt"

func main() {
    const (
        Monday = 0
        Tuesday = 1
        Wednesday = 2
        Thursday = 3
        Friday = 4
        Saturday = 5
        Sunday = 6
    )

    fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday) // 0 1 2 3 4 5 6
}
```

```go
package main

import "fmt"

func main() {
    const (
        Monday = iota
        Tuesday = iota
        Wednesday = iota
        Thursday = iota
        Friday = iota
        Saturday = iota
        Sunday = iota
    )

    fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday) // 0 1 2 3 4 5 6
}
```

### 타입을 지정하지 않은 상수

타입을 지정하지 않은 상수는 타입이 없는 상수. 타입을 지정하지 않은 상수는 `const` 키워드로 선언하고, 타입을 지정하지 않음. 변수에 사용될 때 탕입이 자동으로 결정

```go
package main

import "fmt"

func main() {
 const PI = 3.14 // 타입을 지정하지 않은 상수 선언

 var a int = PI * 100
 var b float64 = PI * 100.0
}

```

### 리터럴

go에서 리터럴은 코드에서 고정된 값을 표현하는 방법. 컴파일 시점에 상수는 리터럴로 대체되어 상수를 사용한 코드는 리터럴로 대체됨

```go
package main

import "fmt"

func main() {
    const PI = 3.14 // 상수 PI 선언

    var a int = PI * 100 // 상수 PI를 사용한 변수 a 선언
    var b float64 = PI * 100.0 // 상수 PI를 사용한 변수 b 선언

    fmt.Println(a, b) // 314 314.0
}
```

## 조건문

### if 조건문

if 조건문은 조건식이 참(true)일 때 코드 블록을 실행. if 조건문은 다음과 같은 구조

```go
if 조건식 {
    // 조건식이 참일 때 실행할 코드
} else if 조건식 {
    // 조건식이 거짓이고, else if 조건식이 참일 때 실행할 코드
} else {
    // 조건식이 거짓일 때 실행할 코드
}
```

### 논리 연산자와 함께 사용

쇼트 서킷 평가: 논리 연산자는 논리식을 평가할 때 쇼트 서킷 평가를 수행. 쇼트 서킷 평가는 논리식을 평가하는 도중에 결과가 확실해지면 나머지 논리식을 평가하지 않음

예를 들어 `a && b`에서 a가 거짓이면 b를 평가하지 않음. `a || b`에서 a가 참이면 b를 평가하지 않음

```go
package main

import "fmt"

func main() {
    var a int = 10
    var b int = 20

    if a > 10 && b < 30 { // a가 10보다 크고, b가 30보다 작을 때
        fmt.Println("a는 10보다 크고, b는 30보다 작습니다.")
    } else if a < 10 || b > 30 { // a가 10보다 작거나, b가 30보다 클 때
        fmt.Println("a는 10보다 작거나, b는 30보다 큽니다.")
    } else if !(a == 10) { // a가 10이 아닐 때
        fmt.Println("a는 10이 아닙니다.")
    }
}
```

### 초기문과 조건식

if 조건문에서 초기문을 사용하여 조건식을 실행하기 전에 변수를 초기화할 수 있음

```go
if 초기문; 조건식 {
    // 조건식이 참일 때 실행할 코드
}
```

```go
package main

import "fmt"

func main() {
    if a := 10; a > 10 { // a를 10으로 초기화하고, a가 10보다 크면
        fmt.Println("a는 10보다 큽니다.")
    } else { // a가 10보다 크지 않으면
        fmt.Println("a는 10보다 크지 않습니다.")
    }
}
```

## switch 조건문

switch 조건문은 여러 조건을 비교할 때 사용. switch 조건문은 다음과 같은 구조

일반적인 다른 언어와 달리 break문을 사용하지 않아도 자동으로 switch 문을 나감

`fallthrough` - 케이스가 일치하더라도 다음 케이스까지 실행(코드 혼동의 가능성으로 쓰지 않는 것이 좋음)

```go
switch 조건식 {
case 값1:
    // 조건식이 값1과 같을 때 실행할 코드
case 값2:
    // 조건식이 값2와 같을 때 실행할 코드
default:
    // 조건식이 값1, 값2와 같지 않을 때 실행할 코드(default 구문은 생략 가능)
}
```

```go
switch 조건식; 비굣값 {
case 값1:
    // 비굣값이 값1과 같을 때 실행할 코드
case 값2:
    // 비굣값이 값2와 같을 때 실행할 코드
default:
    // 비굣값이 값1, 값2와 같지 않을 때 실행할 코드(default 구문은 생략 가능)
}
```

```go
package main

import "fmt"

func main() {
    var a int = 10

    switch a {
    case 10: // a가 10일 때
        fmt.Println("a는 10입니다.")
    case 20: // a가 20일 때
        fmt.Println("a는 20입니다.")
    default: // a가 10, 20이 아닐 때
        fmt.Println("a는 10, 20이 아닙니다.")
    }
}
```

if문 보다 switch문을 사용하는 것이 가독성이 더 좋은 예시는 다음과 같다.

```go
package main

import "fmt"

func main() {

    var day int = 1

    switch day {
    case 1:
        fmt.Println("월요일")
    case 2:
        fmt.Println("화요일")
    case 3:
        fmt.Println("수요일")
    case 4:
        fmt.Println("목요일")
    case 5:
        fmt.Println("금요일")
    case 6:
        fmt.Println("토요일")
    case 7:
        fmt.Println("일요일")
    default:
        fmt.Println("요일이 아닙니다.")
    }
}
```

하나의 case문에 여러 값을 비교할 수 있다. 여러 값을 비교할 때는 콤마(,)로 구분하여 나열

```go
package main

import "fmt"

func main() {
    var a int = 10

    switch a {
    case 10, 20: // a가 10이거나 20일 때
        fmt.Println("a는 10이거나 20입니다.")
    default: // a가 10, 20이 아닐 때
        fmt.Println("a는 10, 20이 아닙니다.")
    }
}
```

초기문 사용하기

switch 조건문에서 초기문을 사용하여 조건식을 실행하기 전에 변수를 초기화할 수 있음

```go
switch 초기문; 조건식 {
case 값1:
    // 조건식이 값1과 같을 때 실행할 코드
case 값2:
    // 조건식이 값2와 같을 때 실행할 코드
default:
    // 조건식이 값1, 값2와 같지 않을 때 실행할 코드
}
```

```go
package main

import "fmt"

func main() {
    var a int = 10

    switch a := 20; a {
    case 10: // a가 10일 때
        fmt.Println("a는 10입니다.")
    case 20: // a가 20일 때
        fmt.Println("a는 20입니다.")
    default: // a가 10, 20이 아닐 때
        fmt.Println("a는 10, 20이 아닙니다.")
    }
}
```

열거형과 함께 사용하기

열거형은 상수를 사용하여 여러 값을 하나의 타입으로 정의하는 방법. 열거형은 switch 조건문에서 사용할 때 유용

```go
package main

import "fmt"

func main() {
    const (
        Monday = iota
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
        Sunday
    )

    func dayToKorean(day int) string {
        switch day {
        case Monday:
            return "월요일"
        case Tuesday:
            return "화요일"
        case Wednesday:
            return "수요일"
        case Thursday:
            return "목요일"
        case Friday:
            return "금요일"
        case Saturday:
            return "토요일"
        case Sunday:
            return "일요일"
        default:
            return "요일이 아닙니다."
        }
    }

    fmt.Println(dayToKorean(Monday)) // 월요일
}
```

## for문

```go
for 초기문; 조건문; 후처리 {
	코드 블록
}
```

- `contunue` - 코드 블록을 수행하지 않고 바로 후처리로 넘어 감
- `break` - for문을 종료

```go
for i := 1; i < 101; i++ {
	fmt.Print(i, ", ") // 1에서 100까지 숫자 출력
}
```

- 초기문 생략이 가능하다

  ```go
  for ; 조건문; 후처리 {
  	코드 블록
  }
  ```

- 후처리 생략이 가능하다

  ```go
  for 초기문; 조건문; {
  	코드 블록
  }
  ```

- 초기문과 후처리 모두 생략이 가능하다

  ```go
  for 조건문 {
  	코드 블록
  }
  ```

- 무한 루프

  ```go
  for { // = for true
  	코드 블록
  }
  ```

- range를 통한 반복

  ```go
  for i:= range 10 { // 0에서 9까지 반복
  	fmt.Println(i)
  }
  ```

## 배열

```
var 변수명 [요소 개수]타입
```

배열의 크기 = 타입 크기 \* 요소 개수

> [!IMPORTANT]
>
> 배열의 요소 개수에는 상수를 사용한다. 변수를 사용할 수 없다

5개의 정수는 담는 배열

```go
var arr [3]int // 제로 값을 가진 길이가 3인 배열
var nums [5]int = [5]int{1, 2, 3, 4, 5} // 1,2,3,4,5 를 가진 길이가 5인 배열
nums := [...]int{1, 2, 3, 4, 5} // 위와 동일
```

5개의 실수를 담는 배열

```go
var arr [5]float64 = [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}
```

배열의 길이 계산

```go
len(arr)
```

range를 통한 배열 순회

```go
for index, value := range arr {
	fmt.Println(i, v)
}
// 0 1.0
// 1 2.0
// 2 3.0
// 3 4.0
// 4 5.0
```

```
for _, value := range arr { // index가 사용되지 않는 경우 "_"를 사용
	fmt.Println(i, v)
}
```

배열의 복사는 같은 타입끼리만 가능

```go
nums = arr // 타입 에러
```

다차원 배열 선언과 초기화

```go
var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

for _ x := range arr {
    for _, y := range x {
        fmt.Print(y)
    }
    fmt.Println()
}
```





## 슬라이스

전통적인 배열을 중심으로 추상화된 Go의 데이터 유형

### 슬라이스의 구성요소

- 슬라이스의 첫 번째 요소를 나타내는 백업 배열의 일부 요소에 대한 포인터
  (반드시 배열의 첫 번째 요소일 필요는 없음)
- 슬라이스의 요소 수를 나타내는 길이
- 길이의 상위 값을 나타내는 용량(지정하지 않으면 슬라이스 시작과 백킹 배열 끝 사이의 요소 수

슬라이스 길이 -> `len`

슬라이스 용량 -> `cap`



```go
s := make([]int, 3) // 
s := []int{0, 0, 0} // 슬라이스 리터럴 방식. 위와 동일
fmt.Println(s) // "[0 0 0]"
fmt.Println(len(s)) // 3
s[0] = 1
s[1] = 2
s[2] = 3
fmt.Println(s) // [1 2 3]
```



### 값 추가

원래의 값에 하나 이상의 새로운 값을 추가한 슬라이스를 반환

```go
	s := []int{1, 2, 3}
	fmt.Println(s) // "[1 2 3]"
	s = append(s, 4)
	fmt.Println(s) // [1 2 3 4]
```

```go
	s := []int{1, 2, 3}
	fmt.Println(s) // "[1 2 3]"
	s = append(s, s...)
	fmt.Println(s) // [1 2 3 1 2 3]
```





## 포인터

명시적으로 초기화하지 않는 경우 null

```go
	var a int = 10

	var p *int = &a // p of type *int points to a
	fmt.Println(p)  // "0xc000104040"
	fmt.Println(*p) // "10"

	*p = 20        // indirectly update a
	fmt.Println(a) // "20"
```

```go
	var a *int
	fmt.Println(a) // <nil>
```





## 오류 처리

```go
result, err := doSomething()
if err != nil {
	log.Fatal(err)
	return err
}
```

### 오류 생성

```go
error1 := errors.New("error Something")
error2 := fmt.Errorf("error Code: %d", 500)
```

