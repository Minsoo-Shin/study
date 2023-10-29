>Kafka Connect는 데이터베이스, 키-값 저장소, 검색 인덱스 및 파일 시스템 간의 간단한 데이터 통합을 위한 중앙 집중식 데이터 허브 역할을 하는 Apache Kafka의 무료 오픈소스 구성요소이다.  Kafka Connect를 사용하여 **Kafka와 다른 데이터 시스템 간에 데이터를 스트리밍하고 Kafka 안팎으로 대규모 데이터 셋을 이동시켜주는 커넥터를 빠르게 생성할 수 있다.** (Confluent)

![[Pasted image 20231022144942.png]]


- 커넥트를 통해서 카프카로 데이터를 전달하거나 받을 수 있다. 
-  테스트 및 일회성 작업을 위한 단독 모드 (standalone mode) 실행할 수 있고, 대규모 운영 환경을 위한 분산 모드 (distributed mode(클러스터형))로 실행할 수 있다.
- 동일한 두 커넥트를 서로 구분하기 위해 소스(source) 방향에 있는 커넥트를 **소스 커넥트(Source Connect)**, 나가는 방향에 있는 커넥트를 **싱크 커넥트 (Sink Connect)** 라고 합니다.


### 용어 정리

> - **Connect** : Connector를 동작하게 하는 프로세서 (서버)
> - **Connector** : Data Source의 데이터를 처리하는 소스가 들어있는 jar 파일 (위 Connector 목록 참조)
> - **Source Connector** : Data Source의 데이터를 카프카 토픽에 보내는 역할을 하는 커넥터 (Producer)
> - **Sink Connector** : 카프카 토픽에 담긴 데이터를 특정 Data Source로 보내는 역할을 하는 커넥터 (Consumer)
> - **단일 모드 (Standalone Mode)** : 하나의 Connect만 사용하는 모드
> - **분산 모드 (Distributed Mode)** : 여러개의 Connect를 한개의 클러스터로 묶어서 사용하는 모드. 특정 Connect에 장애가 발생해도 나머지 Connect가 대신 처리할 수 있음


https://velog.io/@holicme7/Apache-Kafka-Kafka-Connect-%EB%9E%80