feature 
- 한 사용자는 여러 특정 사용자와 대화를 할 수 있다. 
- 대화를 한 메세지는 영구적으로 보관된다. 
	- 최근 3일치 대화 내용을 불러온다. 
- 사용자 닉네임 검색을 통해서 상대방을 찾아낸다. 


## User

```go
type user struct {
	id int 
	nickname string 
	email string 
	password string
}
```

## Chatroom
```go 
type chatroom struct {
	chatroom_id int 
	users []user
}
```


## Message 
```go
type message struct {
	message_id int 
	from int
	to int
	content string 
	create_at time.Time
}
```

## Architecture 
### stateless
- 로그인
- 유저 정보 변경
- 만들어진 채팅방 리스트 조회 & 마지막 메세지 노출

### stateful service
- 채팅방 생성
- 채팅하기

### 로그인
- 이메일/비밀번호를 입력하면 사용자 정보와 토큰을 반환한다. 

### 회원 가입
- 당장은 개발하지 않음



### 채팅방 생성 
- ws connection 을 연결 맺음
- 특정 사용자의 닉네임을 통해서 찾는다. 
- sender가 ws conn 해도 receiver는 ws conn가 아닐 수 있다. 

### 채팅하기
- ws로 message가 전달되면 kv store 저장한다. 
