
## 정적 클래스 의존관계 
---
클래스가 사용하는 import 코드만 보고 의존관계를 쉽게 판단할 수 있다. 정적인 의존관계는 애플리케이션을 실행하지 않아도 분석할 수 있다. 클래스 다이어그램을 보자
![[Pasted image 20231007130829.png]]

## 동적인 객체 인스턴스 의존 관계
---
애플리케이션 실행 시점에 실제 생성된 객체 인스턴스의 참조가 연결된 의존 관계다.
![[Pasted image 20231007131005.png]]

### IoC컨테이너, DI 컨테이너
--- 
AppConfig와 같이 사용 클래스를 생성하고, 관리하고, 의존 관계를 연결해주는 것을 IoC컨테이너 또는 DI 컨테이너라고 말함.
어셈블러, 오브젝트 팩토리라고 불리기도 한다. 


### Spring을 통한 DI
--- 
AppConfig를 통하여 DI를 실습해보았다. 그 역할을 spring을 통해서 아래와 같이 해주면 applicationContext에 저장하고 `getBean` 을 통해서 가져올 수 있다. 


- `@Configuration` : AppConfig에 설정을 구성한다는 뜻
- `@Bean`: 메서드 위에 annotation하여 스프링 빈에 등록


```Java
@Configuration
public class AppConfig {

    @Bean    
    public MemberService memberService() {

        return new MemberServiceImpl(memberRepository());
    }

    @Bean    
    public OrderService orderService() {

        return new OrderServiceImpl(
                memberRepository(),

                discountPolicy());

}

    @Bean    
    public MemberRepository memberRepository() {

        return new MemoryMemberRepository();
    }

    @Bean    
    public DiscountPolicy discountPolicy() {

        return new RateDiscountPolicy();
    }

}
```

```Java
 ApplicationContext applicationContext = new
AnnotationConfigApplicationContext(AppConfig.class);

        MemberService memberService =
applicationContext.getBean("memberService", MemberService.class);

        OrderService orderService = applicationContext.getBean("orderService",
OrderService.class);
```


이 spring을 통한 DI를 하면 무슨 장점이 있을까?
[[스프링 컨테이너]]