리눅스의 I/O와 연관이 깊은 dirty page에 대해서 공부해보자

# dirty page란

> 리눅스는 파일에 대한 읽기/쓰기 작업을 할 때 pagecache라는 것을 활용한다. 읽기 작업을 한다고 해서 항상 디스크에 갔다 오는 것이 아니고 한번 읽은 디스크의 내용을 pagechache에 저장해놓고 다시 해당 파일에 읽기 작업이 일어나면 pagecache에서 바로 꺼내준다. 이를 통해 I/O 작업이 일어날때, 불필요한 디스크 접근을 줄여서 성능 향상을 기대할 수 있다.

![[Pasted image 20231019212022.png]]

## pagecache에 저장된 디스크의 파일 내용에 변경이 생기면 어떻게 되나? 

pagecache 메모리 영역과 디스크 파일 내용의 차이가 날 것이다. 그래서 파일를 쓸때 pagecache에는 **페이지가 변경이 되어서 실제 디스크의 내용과는 다르다는 것을 표시**한다. 이를 dirty page라고 부른다. 

![[Pasted image 20231019212325.png]]



1. 맨 처음 pagecache에 있는 데이터와 디스크에 있는 데이터는 동일한 데이터를 유지
2. b라는 내용의 파일이 d라는 내용으로 변경된다. 
3. 디스크에 바로 저장하지 않고 메모리 영역만 변경한다. (dirty page)
4. 그 후 일정 조건이 되면 커널은 다양한 방법으로 dirty page의 내용을 실제 디스크에 쓰게 된다. 

### dirty page는 메모리에 불과하다. -> 휘발된다. -> 적절한 동기화 수준이 필요하다. 

그래서 동기화의 수준이 어느 정도냐에 따라서 심한 I/O부하를 일으킬 수 있기 때문에 커널은 파라미터를 통해서 워크로드에 따라 이런 동기화 작업을 컨트롤할 수 있게 제공해준다. 

dirty 관련 파라미터
- vm.dirty_background_ratio = 10
- vm.dirty_background_bytes = 0
- vm.dirty_ratio = 20
- vm.dirty_byte = 0
- vm.dirty_writeback_centisecs = 500
- vm.dirty_expire_centisecs = 3000


#### vm.dirty_background_ratio
- flush 데몬이 깨어나서 dirty page 들을 싱크시키는 조건 중 하나.
- 전체 메모리 대비 dirty page의 존재 비율
- 전체 메모리 8GB * 10% = 800MB
- 크기가 800MB이상되면 flush 데몬이 깨어나서 싱크
- 프로세스와 무관하게 background에서 동작

#### vm.dirty_background_bytes
- 바이트이다. 
- 기본 값은 0으로 놓고 disable 해놓는다. 

#### vm.dirty_ration
- 기본값이 10이다. 
- vm.dirty_background_ratio과 동일하지만 background가 아니고 다른 쓰기 I/O 블락시키고 진행


#### vm.dirty_bytes
- 특정 bytes를 넘어가면 쓰기 I/O를 블락시키고 싱크한다. 
- 기본값은 0, disable 해놓음

#### vm.dirty_writeback_centisecs
- flush 데몬이 background에서 깨어나는 기준이 되는 값. 
- 1/100s 기준 => 500이면 5초를 의미
- 기본값 500 (5s)

#### vm.dirty_expire_centisecs
기본값은 3000
30초 이상된 dirty page 싱크 시키라는 의미




# 정리

1. vm.dirty_background_ratio는 전체 메모리 대비 dirty page 비율을 의미하며 여기에 설정한 값을 넘기면 flush 데몬이 background 모드로 동작하면서 dirty page를 모두 비웁니다. 
2. vm.dirty_writeback_centisec은 flush 데몬을 foreground에서 동작하게 하는 인터벌을 의미하며, 이 인터벌마다 깨어나서 생성시간이 vm.dirty_expire_centisec 보다 오래된 dirty page들을 모두 비웁니다. 






https://brunch.co.kr/@alden/32


