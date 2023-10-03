[공식문서](https://kotlinlang.org/docs/data-classes.html)

> class의 한 종류이며, data를 가지고 있는 것을 주된 목적으로 한다. 

```Kotlin 
data class User(val name: String, val age: Int)
```


## 클래스 본문에 선언된 속성
- 컴파일러는 기본 생성자 내부에 정의된 (아래 name)만 사용해 함수를 자동적으로 생성
- 함수에 속성을 제외하려면 body 내부에 정의하면 된다. 
```Kotlin
data class Person(val name: String) {
	var age: Int = 0 
}
```


아래를 보면 `name`만 동일하면 `age`가 달라도 같은 객체로 판단한다는 것을 알 수 있다. 
```Kotlin 
data class Person(val name: String) {
    var age: Int = 0
}
fun main() {

    val person1 = Person("John")
    val person2 = Person("John")
    person1.age = 10
    person2.age = 20

    println("person1 == person2: ${person1 == person2}")
    // person1 == person2: true
  
    println("person1 with age ${person1.age}: ${person1}")
    // person1 with age 10: Person(name=John)
  
    println("person2 with age ${person2.age}: ${person2}")
    // person2 with age 20: Person(name=John)

}
```