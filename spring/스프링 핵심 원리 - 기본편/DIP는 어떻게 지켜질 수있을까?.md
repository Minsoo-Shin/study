# DIP란
---
> DIP 원칙이란 객체에서 어떤 Class를 참조해서 사용해야하는 상황이 생긴다면, 그 Class를 직접 참조하는 것이 아니라 그 **대상의 상위 요소(추상 클래스 or 인터페이스)로 참조**하라는 원칙이다.

> 객체들이 서로 정보를 주고 받을 때는 의존 관계가 형성되는데, 이 때 객체들은 **나름대로의 원칙**을 갖고 정보를 주고 받아야 하는 약속이 있다. 여기서 나름대로의 원칙이란 추상성이 낮은 클래스보다 **추상성이 높은 클래스와 통신**을 한다는 것을 의미하는데 이것이 DIP 원칙이다.

```Java
public class OrderServiceImpl implements OrderService { 
	private final MemberRepository memberRepository = new MemoryMemberRepository(); 
	private final DiscountPolicy discountPolicy = new FixDiscountPolicy(); 
}
```

### 문제점
---
위의 코드에서 Service는 MemberRepository (인터페이스) or MemoryMemberRepository(구현) 모두 의존하게 된다. 그 Class를 직접 참조하게 되는 것이다. 

=> 사실상 DIP를 지키지 못하고 있다. 


### 해결 방안
---
애플리케이션의 전체 동작 방식을 구성(config)하기 위해, **구현 객체를 생성**하고, **연결**하는 책임을 가지는 별도의
설정 클래스를 만들자.

```Java
 public class AppConfig {
     public MemberService memberService() {
         return new MemberServiceImpl(memberRepository());
}
     public OrderService orderService() {
         return new OrderServiceImpl(memberRepository(), discountPolicy());
}
     public MemberRepository memberRepository() {
         return new MemoryMemberRepository();
}
     public DiscountPolicy discountPolicy() {
         return new FixDiscountPolicy();
}

}
```
