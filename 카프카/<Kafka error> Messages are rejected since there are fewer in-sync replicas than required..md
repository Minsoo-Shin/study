
## 키워드
- [[in-sync replica (ISR)]]

## 해결
- Topic replica-factor를 1로 했었는데, 3으로 하니 해결되었다. 

## 왜 
- 카프카는 가용성 보장을 위해서 replication 전략을 취한다. 이 때, replication 성능을 올리기 위해서 최소한의 수만 복제가 완료되면 리더는 쓰기가 완료되었다고 표시를 하게 된다. 
- 하지만, Topic은 1만 replica를 하도록 되어있는데 msk 기본 설정은 최소 insync 레플리카 수가 2는 있어야하기 때문에 애초에 커밋이 불가능한 토픽을 만든 것이다. 
- 애초에 불가능한 토픽이라면 토픽 만들 때 에러를 반환하면 될텐데 프로듀서 메시지 입력할 때 문제가 생기는지 모르겠다. 
- Minimum In-Sync Replicas 기준은 어디에 있는걸까? 
	- msk로 만들면 따로 설정하지 않으면 aws 기본 설정값으로 아래와 같이 설정되어있다. 
- 가용 영역 3개에 있는 클러스터 가정시
	-  `min.insync.replicas` = 2
	- `replication.factor`  = 3
	

![[Pasted image 20231105235106.png]]
![[Pasted image 20231105235043.png]]https://docs.aws.amazon.com/ko_kr/msk/latest/developerguide/msk-default-configuration.html