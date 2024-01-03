# Concurrency is not parallelism
---
- 동시성(concurrency)는 병렬성(parallelism)이 아니지만 병렬성을 가능케 해준다. 
- 누군가가 하나의 프로세서만 가지고 있다면 동시성을 가지지만 병렬적으로 처리는 하지 못한다. 

### 동시성

> 컴퓨터 사이언스에서, 동시성이란 프로그램이나 알고리즘, 문제의 여러 부분이 비순차적 또는 부분적 순서로 실행되는 능력을 말한다. 
https://en.wikipedia.org/wiki/Concurrency_(computer_science)

동시성은 실제로 여러 작업이 동시에 진행되는 것처럼 보이도록 하는 개념이다. 하나의 프로세서에서 여러 작업이 번갈아가며 실행되면서 동시에 진행되는 것처럼 느껴진다. 이는 주로 I/O 바운드 작업에서 효과적이다. 

- I/O작업(파일 읽기/쓰기, 네트워크 통신) 시간 소모가 크다. 단일 스레드로 처리하면 I/O 작업 중에는 CPU는 일을 하지 않고 기다리고 있기 때문에 효율적이지 않다. 그 때 CPU는 다른 작업을 한다면 작업을 효율적으로 빠르게 끝낼 수 있는 것이다. 

> 로버트 파이크에 따르면, 동시성은 구조에 관한 것이고, 병렬성은 실행에 관한 것이다. 
https://www.youtube.com/watch?v=f6kdp27TYZs&ab_channel=GoogleforDevelopers


### 병렬성

> 많은 계산 또는 프로세스들이 동시에 실행되는 것을 의미한다. 
https://en.wikipedia.org/wiki/Parallel_computing


> Parallelism은 하나의 작업의 응답시간을 줄일 수 있다. 내부의 하위 작업들을 동시에 수행하기 때문이다. 반면, Concurrency는 단일 작업에 대한 응답시간을 줄일 수는 없다. 대신 한 번에 여러 작업을 처리해서 Throughput을 늘리는 데에 의미가 있다.
https://blog.naver.com/complusblog/220987936079



### 그렇다면 Golang에서 동시성을 어떻게 구현할 수 있을까? 
--- 
이를 알려면 Goroutine(고루틴)을 알아야 한다. 고루틴은 전통적으로 스레드와 비슷한 개념이지만 많이 <mark style="background: #D2B3FFA6;">경량화된 스레드라고 보면 된다. </mark>
- 고루틴은 독립적으로 실행되는 함수이다. (independently executing function)
- 자체 스택을 가지고 있다. 필요에 따라서 커지고 줄어들고 한다. 
	- 전통적인 스레드의 경우에는 스레드가 생성될 때, 고정된 크기로 할당한다. 이 스택의 동적으로 확장할 수 없기에 이 한계를 넘어가면 스택 오버플로우 (stackoverflow)가 발생한다. 
	- 적절한 스택 크기를 설정하지 않으면 콘 메모리를 할당받아 낭비될 수 있다. 
- 수천개, 수만개의 고루틴을 실행하는 것은 매우 저렴하고 실용적이다. 수백만개의 고루틴 역시 프로덕션 환경에서 잘 수행된다고한다. 
 <iframe width="560" height="315" src="https://www.youtube.com/embed/f6kdp27TYZs?si=jUbuwpKDbWF0Sgwo&amp;start=572" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

#### 일반적인 스레드가 동기화 메커니즘을 구현하는 방법
	뮤텍스 : 공유 자원에 대한 동시 액세스를 제어하기 위한 동기화 기법 중 하나이다. 
	세마포어 : 뮤텍스와 유사하지만, 뮤텍스는 1또는 0의 이진 상태를 가진다. 이는 자원을 사용하고 있는 상태, 하지 않는 상태를 뜻한다. 하지만 세마포어는 여러 개의 리소스를 가질 수 있어 여러 스레드가 접근할 수 있는 자원의 수를 제어할 수 있다. 
	조건 변수 : 스레드 간 특정 조건이 충족될 때까지 대기하거나 신호를 보내는데 사용된다. 
	세마포어 및 모니터를 이용한 동기화


반면 Go 언어의 고루틴은 내장된 채널을 통해 통신하고 동기화된다. <mark style="background: #D2B3FFA6;">채널은 메모리 안전성과 동시성을 보장</mark>하는데 사용되어 더욱 쉽게 동시성 문제를 다룰 수 있다. 

### 구체적으로 Gorutine과 Channel을 통해서 어떻게 동기화 메커니즘을 구현하는 패턴에 대해서 공부해보자
---
### Generator : 채널을 반환하는 함수

> 아래와 같이 채널을 반환하는 함수(boring)는 함수 내부에서 고루틴을 실행했고, 메세지를 생산해서 채널로 보내고, main함수는 해당 채널로 메세지를 받는다.  main - boring 내부의 익명함수 고루틴간 메세지를 channel을 통해서 전달하는 것이다. 

```go
func main() {
	c := boring("boring!") // Function returning a channel.
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

```

그렇다면 여러 boring한 사람과 대화를 하면 어떻게 할까? 
```go
func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}
// https://go.dev/play/p/e-WjvGnKxPE
```
![[Pasted image 20231225231344.png]]

하지만 이는 joe를 메세지를 받은 다음 ann의 메세지를 받을 수 있다. 우리는 joe, ann 둘 중에 먼저 말한 것들을 받아서 출력하고 싶다. 


### Multiplexing
---
>  Multiplexing은 여러 개의 데이터 스트림을 하나의 통로를 통해 전송하는 기술이나 방법을 가리킵니다. 이는 네트워크 통신 및 입출력(IO) 작업에서 특히 중요한 역할을 합니다. 여러 가지 응용 분야에서 사용되지만, 주로 다음 두 가지 상황에서 많이 사용됩니다.
```go
func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // input1에서 받은 데이터를 c로 보낸다. 
		}
	}()
	go func() {
		for {
			c <- <-input2 // input2에서 받은 데이터를 c로 보낸다. 
		}
	}()
	return c
}

```

![[Pasted image 20231225232039.png]]


위와 같이 여러 개의 채널을 `fanIn` 함수를 통해서 하나의 채널로 데이터를 보내준다. 하지만 이보다 더 간단한 방식이 있다. 
바로 `select` 구문이다. 여러 채널을 관리해주는 방식을 제공해준다. 

```go
// before
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // input1에서 받은 데이터를 c로 보낸다. 
		}
	}()
	go func() {
		for {
			c <- <-input2 // input2에서 받은 데이터를 c로 보낸다. 
		}
	}()
	return c
}
// after
func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() 
        for {
            select {
            case s := <-input1:  c <- s
            case s := <-input2:  c <- s
            }
        }
    }()
    return c
}
```

매우 간단해졌다. goroutine도 두개 쓰던걸 하나만 써도 가능해졌다. 

select 구문 특징
- 한 커뮤니케이션이 진행할 때까지 Selection은 차단된다. 
```go
func main() {
	c := notAnswer("Joe")
	// c라는 channel에 아무 값이 들어오지 않아 select 구문에서 block 되어있다. 
	select {
	case s := <-c:
		fmt.Println(s) 
	}
	// 계속 실행되지 않을 것이다. 
	fmt.Println("select finished")
}

// do not send chan anything
func notAnswer(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			// do not anything
		}
	}()
	return c
}
// https://go.dev/play/p/NTOIfTNl3UN
```
![[Pasted image 20231226004438.png]]
- default clause는 어떤 channel에서도 준비되어 있지 않을 때 실행된다. 
```go
func main() {
	c := notAnswer("Joe")
	// c라는 channel에 아무 값이 들어오지 않아 select 구문에서 block 되어있다. 
	select {
	case s := <-c:
		fmt.Println(s) 
	default: fmt.Println("waiting")
	}
	// 계속 실행되지 않을 것이다. 
	fmt.Println("select finished")
}

// do not send chan anything
func notAnswer(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			// do not anything
		}
	}()
	return c
} 
// https://go.dev/play/p/NlVEo1dNj0O
```
![[Pasted image 20231226004418.png]]


### Timeout using select 
---

time.After는 특정 시간 후에 신호를 주는 chandmf qksghksgksek. 
```go
func main() {
    c := boring("Joe")
    for {
        select {
        case s := <-c:
            fmt.Println(s)
        case <-time.After(1 * time.Second):
            fmt.Println("You're too slow.")
            return
        }
    }
}
```


### Timeout for whole conversation using select 
---
특정 시간이 끝나면 모든 channel들을 타임 아웃 시킨다. 

```go
func main() {
    c := boring("Joe")
    timeout := time.After(5 * time.Second)
    for {
        select {
        case s := <-c:
            fmt.Println(s)
        case <-timeout:
            fmt.Println("You talk too much.")
            return
        }
    }
}
//https://go.dev/play/p/4PRpAFtXRjz
```

![[Pasted image 20231226153505.png]]

## Google Search 예시 
```go
package main  
  
import (  
"fmt"  
"math/rand"  
"time"  
)  
  
var (  
	Web = fakeSearch("web")  
	Image = fakeSearch("image")  
	Video = fakeSearch("video")  
	  
	Web1 = fakeSearch("web1")  
	Web2 = fakeSearch("web2")  
	Image1 = fakeSearch("image1")  
	Image2 = fakeSearch("image2")  
	Video1 = fakeSearch("video1")  
	Video2 = fakeSearch("video2")  
)  
  
type Result string  
  
type Search func(query string) Result  
  
func fakeSearch(kind string) Search {  
	return func(query string) Result {  
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)  
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))  
}  
}  
  
func GoogleV1(query string) (results []Result) {  
	results = append(results, Web(query))  
	results = append(results, Image(query))  
	results = append(results, Video(query))  
	return results  
}  
  
func GoogleV2(query string) (results []Result) {  
	c := make(chan Result)  
	go func() { c <- Web(query) }()  
	go func() { c <- Image(query) }()  
	go func() { c <- Video(query) }()  
  
	for i := 0; i < 3; i++ {  
		result := <-c  
		results = append(results, result)  
	}  
	return results  
}  
  
func GoogleV2_1(query string) (results []Result) {  
	c := make(chan Result)  
	go func() { c <- Web(query) }()  
	go func() { c <- Image(query) }()  
	go func() { c <- Video(query) }()  
  
	timeout := time.After(80 * time.Millisecond)  
	for i := 0; i < 3; i++ {  
		select {  
		case result := <-c:  
			results = append(results, result)  
		case <-timeout:  
			fmt.Println("timed out")  
		return  
		}  
	}  
	return results  
}  
  
func First(query string, replicas ...Search) Result {  
	c := make(chan Result)  
	searchReplica := func(i int) { c <- replicas[i](query) }  
	for i := range replicas {  
		go searchReplica(i)  
	}  
	return <-c  
}  
  
func GoogleV3(query string) (results []Result) {  
	c := make(chan Result)  
	go func() { c <- First(query, Web1, Web2) }()  
	go func() { c <- First(query, Image1, Image2) }()  
	go func() { c <- First(query, Video1, Video2) }()  
	  
	timeout := time.After(80 * time.Millisecond)  
	for i := 0; i < 3; i++ {  
		select {  
		case result := <-c:  
			results = append(results, result)  
		case <-timeout:  
			fmt.Println("timed out")  
		return  
		}  
	}  
	return results  
}  
  
func main() {  
	var sum time.Duration  
	for i := 0; i < 8; i++ {  
		rand.Seed(time.Now().UnixNano())  
		start := time.Now()  
		_ = GoogleV3("golang")  
		  
		//_ = GoogleV2_1("golang")  
		elapsed := time.Since(start)  
		sum += elapsed  
	}  
	  
	fmt.Println("평균은", sum/8)  
  
}
```

![[Pasted image 20240104010802.png]]


v1에서는 순차적으로 쿼리를 실행해서 배열에 결과들을 담았다고 하면, 
v2에서는 fan-in방식으로 비동기로 쿼리를 돌리고, 들어오는대로 결과를 담고 리턴한다. 
이럴 때는 기존보다는 비동기로 한번에 쿼리를 실행하기 때문에 쿼리 + 네트워크 IO를 비동기적으로 진행할 수 있어서 성능면에서 우수하다. 
V2.1은 너무 느린 쿼리있다고 가정했을 때, timeout을 둠으로써 서버의 안전성을 가져올 수 있다. 
V3.0은 쿼리마다 여러 레플리카로 요청해서 빠르게 처리할 확률을 많이 높이는 듯 하다. 하지만 왜 3배나 차이가 날까? 



# race condition 방지를 위한 도구
---
1. **`sync` 패키지 활용:** `sync` 패키지는 여러 고루틴 간의 동기화를 도와주는데, `Mutex`와 `RWMutex` 등을 활용하여 공유 자원에 대한 안전한 접근을 보장할 수 있습니다.
2. **`atomic` 패키지 사용:** `atomic` 패키지는 원자적인 연산을 제공하여 race condition을 방지할 수 있습니다.
3. **채널 활용:** 채널을 사용하여 고루틴 간 통신을 통해 race condition을 피할 수 있습니다.