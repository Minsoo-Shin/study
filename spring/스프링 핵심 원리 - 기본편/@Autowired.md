- 생성자에 지정하는 주석이다.
- 의존성 자동 주입 용도로 사용한다. 
- 스프링 컨테이너가 자동으로 해당 스프링 빈을 찾아서 주입한다. 

## 빈 조회 규칙

1. 주입받고자하는 타입으로 매칭을 시도한다. 
2. 타입이 여러 개면 필드 또는 파라미터 이름으로 매칭을 시도한다. 

### 같은 타입의 다중 구현체(여러 빈)가 존재하는 경우 아래와 같은 에러를 반환한다. 

> Exception encountered during context initialization - cancelling refresh attempt: org.springframework.beans.factory.UnsatisfiedDependencyException: Error creating bean with name 'orderServiceImpl' defined in file [/Users/ms/Downloads/core 2/out/production/classes/hello/core/order/OrderServiceImpl.class]: Unsatisfied dependency expressed through constructor parameter 1: No qualifying bean of type 'hello.core.discount.DiscountPolicy' available: expected single matching bean but found 2: fixDiscountPolicy,rateDiscountPolicy

**expected single matching bean but found 2** 말 그대로 같은 타입이 2개가 있으니 어떤 걸 선택해야할지 몰라 나오는 에러다. 

## 빈의 구분자나 우선순위를 부여하여 편리하게 해결할 수 있다.

- [[@Qualifier]]
- @Primary

## Autowired시에 빈이 존재하지 않으면 null값으로 받는 방법

- @Autowired(required=false) 를 사용해주면 된다. 


