다른 사람들에게 나에 대해서 소개를 하고 피드백을 받는 그런 서비스를 개발하려고 한다. 
이 서비스는 왜 필요한가? 다른 사람들에게 자기를 알리는 것을 좋아하는 사람들도 있겠지만 다수에게 자기를 알리는 것을 부끄러워 하는 사람들도 분명 존재한다. 또는 이름과 소속 이외에는 말할게 없는 즉, 본인을 잘 모르는 사람들도 있을 것이다. 이 서비스는 이 사람들에게 초점을 맞춰서 기능 개발하려고 한다. 

서비스에는 어떤 기능들이 있을까? 
- 자기를 표현할 수 있는 소개서를 모던하게 꾸밀 수 있다. 
	- 취미나 관심사, MBTI 등과 같은 선택 사항 제공
	- 그룹 내 구성원들로 부터 카드를 받아서 본인의 성격이나 제 3자가 바라본 모습들을 알 수 있다. 
- 그룹 안에서 여러 사람들과의 서로간의 피드백을 나눌 수 있다. 
	- 칭찬/축하 등과 같은 카드 형식
- 그룹을 나가더라도 본인이 받은 카드들은 추억으로 간직할 수 있다. 

### 기능 요약
--- 
#### 유저
- 회원가입/일반 로그인/소셜 로그인
- 개인정보 등을 저장할 수 있다.
#### 유저 상세 (유저 테이블만 만들어야할지...)
- 취미/MBTI 생성/수정
- 유저:유저상세(1:1)
#### 자기 소개서
- 본인의 자기 소개서를 조회할 수 있다. (불러오기/복사)
- 유저들 특정 그룹의 자기소개서 생성
- 유저 - 자기소개서 (1:N)
#### 그룹
- 그룹을 생성/수정/삭제가 가능하다.
	- 그룹은 팀원들을 가진다. (유저 - 그룹, N:N)
	- 그룹원들의 권한이 어떤 것들이 있는가? 
#### 그룹원 소개서
- 자기 소개서를 그룹별로 하나씩 생성/수정/삭제가 가능하다.
- 자기소개서를 스냅샷
	- 스냅샷 해두지 않고 id로 연결해놓으면 그룹을 나가고 해당 소개서가 변경되는 사태가 벌어질 수 있음.
- 그룹:그룹원 소개서 (1:N)
#### 카드 
- 유저는 카드를 생성할 수 있다. (친해지고 싶어요/칭찬해요/축하/피드백)
- 카드는 공개/비공개 설정이 가능하여 비공개라면 다른 그룹원들은 노출되지 않는다. 
- (그룹을 떠나더라도) 특정 그룹에서 본인이 보내고 받은 카드를 조회할 수 있다.

### 질문
- (프론트 질문) 단체 초대 링크는 어떻게 설정하는지? 
- 예시)
	- 자기소개서 1개를 가지고 있다. 특정 그룹에 링크를 통해서 들어가게 되었고, 자기소개서를 불러오기하여 일부 항목을 수정하고 특정 그룹에 공유할 자기소개서를 저장한다. 그렇다면 마이페이지에서는 두개의 자기소개서를 가지게 된다. (<-확인 필요). 
	- 자기소개서(2번째 생성한)를 수정했을 때, 마이페이지에서 보는 자기소개서도 변경되고, 그룹원들이 보는 소개서도 변경이 되는지? 
	- (확인) 그룹을 나간 유저의 자기소개서는 그룹 페이지에서 타 그룹원들에게 노출이 되는가? 
		- 안된다고 하면 그룹을 나간 유저들에게 보낸 나의 카드는 어색하게 표현되지 않은지... 
		- 된다고 하면, 그룹을 나간 유저가 해당 그룹 공유한 자기소개서를 수정할 때, 남아있는 그룹원들에게 보여지는 (그룹을 나간 유저의) 자기소개서에도 반영이 되어야 하는지 확인 필요


### 사용 기술 스택
---
- java/spring
	- 왜냐 그냥 spring을 배워보고 싶기 때문이다.
	- 얼마나 생산성이 좋은지 체험해보자
- mysql
	- 음 아무래도 insert 빠르게 하려면 nosql인데...
	- 기본적으로 안전성을 제공해주는 rdb가 국룰
- nginx
	- reverse proxy로 사용 (ip 은닉)
	- auto scaling이 가능할지....? 확인이 필요하다. 
- aws ecs
	- 이미지로 관리하면 이식성 좋기 떄문이다.
- github action/jenkins
	- ci/cd

### ERD
---
### 서버 spec
---
#### 1차 기준 (카드 작성 70개/분, peak 2000개/분)
> 2020년 기준 창업 기업 수는 307만개이다. 24년 초 기준 투자시장이 얼어붙어 많이 줄었지만 다시 경기가 좋아질 것을 희망하며, 307만개의 스타트업들의 1%가 사용한다고 가정하고 spec을 생각할 것이다. 

3만개의 스타트업이다. 3만개의 스타트업에 각 스타트업은 평균적으로 10명의 직원이 있다고 가정하자. 이 직원들은 월마다 한번씩 직원들끼리 소통 행사를 하여 카드를 주고 받는다. 월마다 직원들은 다른 직원들에게 작성하는 카드의 평균 수는 10개 정도이다. 

그렇다면 계산을 해보자.
(가장 ideal한 예시)
3만개 * 10명 = 30만명의 직원들이 있다. 
30만명의 직원들이 월 평균 10개의 카드를 작성한다.

> 300만/월 = 10만/일 = 4200(≒4166)/한시간 = 70(≒69.4)개/분

(peak한 예시)
30만명의 직원들이 월말 (30일)에 몰려서 한다고 가정하자.
300만/일 = 12.5만개/한시간 = 2000개/분


- 서버 인스턴스의 사양 (CPU/Memory) 몇개가 적당할 것인가? 
- Auto-Scaling은 어떻게 할 것인가? 
### 인프라
---
#### CI/CD
- githun action/Jenkins
#### 서버 구성
- nginx/api/db/Jenkins

### 테스트
---
- JUnit : 자바 프로그래밍 언어용 단위 테스트 프레임워크
- Spring Test & Spring Boot Test : 스프링 부트 애플리케이션을 위한 통합 테스트 지원
- AssertJ : 검증문인 어설션을 작성하는 데 사용되는 라이브러리
- Hamcrest : 표현식을 이해하기 쉽게 만드는 데 사용되는 Matcher 라이브러리
- Mockito : 테스트에 사용할 가짜 객체인 목 객체를 쉽게 만들고, 관리하고, 검증할 수 있게 지원하는 테스트 프레임워크
- JSONassert : JSON용 어설션 라이브러리
- JsonPath : JSON 데이터에서 특정 데이터를 선택하고 검색하기 위한 라이브러리

https://goldenrabbit.co.kr/2023/06/30/springtest/

### 보안
---
- 애플리케이션
	- 사용자 action history 추적
- 인프라
	- private/public 나눌지, 서버 .pem key


### 마이그레이션
---
