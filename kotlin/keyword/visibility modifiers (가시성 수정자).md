클래스, 객체, 인터페이스, 생성자, 함수는 물론 `properties` 나 `setter`에도 `visibility modifiers`가 있을 수 있습니다. Getter는 항상 해당 `properties`와 동일한 가시성을 갖습니다.

Kotlin에는 `private` , `protected` , `internal` 및 `public` 의 네 가지 `visibility modifiers`가 있습니다. 기본 가시성은 `public` 입니다.

이 페이지에서는 수정자가 다양한 유형의 선언 범위에 적용되는 방법을 알아봅니다.
- `public`가 default. 다른 곳에서 다 볼 수 있음
- `private`는 같은 파일 내에서만 보인다. 
- `internal` 같은 [[Modules (모듈)]] 내에서만 가능하다. 

```Kotlin
 // file name: example.kt 
 package foo 
 
 private fun foo() { ... } // visible inside example.kt
 
 public var bar: Int = 5 // property is visible everywhere 
	 private set // setter is visible only in example.kt 
	 
 internal val baz = 6 // visible inside the same module
```
