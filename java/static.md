
- [static 변수](https://wikidocs.net/228#static)
- [static 메서드](https://wikidocs.net/228#static_1)
- [싱글톤 패턴](https://wikidocs.net/228#_1)

## static 변수
---
[점프 투 자바] 책에서 좋은 예제가 있어서 그대로 사용하려고 한다. 이씨 집안은 당연히 lastname(성)이 이씨 일 것이다. 하지만 아래와 같이 객체를 생성하면 변수 lastname을 생성하기 위해서 별도 메모리를 사용한다. 

### static 변수 사용 이유
- 항상 값이 변하지 않는다면 static을 사용해 메모리 낭비를 줄일 수 있다.
- **값을 공유할 수 있다.** 
```JAVA
public class JavaStaticTest {  
  
    class HouseLee {  
        String lastname = "이";  
    }  
  
    @Test  
    @DisplayName("일반적으로 객체를 생성하면 객체마다 필드를 공유하지 않는다.")  
    public void StaticTest() {  
        HouseLee lee1 = new HouseLee();  
        HouseLee lee2 = new HouseLee();  
  
        assertThat(lee1.lastname).isSameAs(lee2.lastname);  
  
        lee1.lastname = "다른성";  
  
        assertThat(lee1.lastname).isNotSameAs(lee2.lastname);  
    }  
  
  
    class HouseKim {  
        static String lastname = "김";  
    }  
  
    @Test  
    @DisplayName("static 키워드시 객체마다 필드를 공유할 수 있다.")  
    public void Static2Test() {  
        HouseKim kim1 = new HouseKim();  
        HouseKim kim2 = new HouseKim();  
  
        assertThat(kim1.lastname).isSameAs(kim2.lastname);  
  
        kim1.lastname = "다른성";  
  
        assertThat(kim1.lastname).isSameAs(kim2.lastname);  
    }  
}
```


## Static

- static은 "고정된" 이라는 의미
- 객체 생성 없이 사용할 수 있는 필드와 메소드를 생성하고자 할 때 활용한다.
- 필드나 메소드를 객체마다 다르게 가져야 한다면 인스턴스로 생성하면 되고
- 공용 데이터에 해당하거나 인스턴스 필드를 포함하지 않는 메소드를 선언하고자 할 때 이용한다.
- 사용하기 위해선 클래스 내에서 필드나 메소드 선언 시 static 키워드를 붙여주기만 하면 된다.

```Java
public class PlusClass{ 
	static int field1 = 15;
	
	static int plusMethod(int x, int y){ return x+y; } 
}
```