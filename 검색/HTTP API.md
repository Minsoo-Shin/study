
## 특정 토크나이저로 분석하기

```json
GET _analyze
{
  "tokenizer": "nori_tokenizer",
  "text": [
    "동해물과 백두산이"
  ]
}
```

- response
	```json
	{
	  "tokens": [
		{
		  "token": "나",
		  "start_offset": 0,
		  "end_offset": 1,
		  "type": "word",
		  "position": 0
		},
		{
		  "token": "는",
		  "start_offset": 1,
		  "end_offset": 2,
		  "type": "word",
		  "position": 1
		},
		{
		  "token": "카페",
		  "start_offset": 3,
		  "end_offset": 5,
		  "type": "word",
		  "position": 2
		},
		{
		  "token": "에서",
		  "start_offset": 5,
		  "end_offset": 7,
		  "type": "word",
		  "position": 3
		},
		{
		  "token": "공부",
		  "start_offset": 8,
		  "end_offset": 10,
		  "type": "word",
		  "position": 4
		},
		{
		  "token": "를",
		  "start_offset": 10,
		  "end_offset": 11,
		  "type": "word",
		  "position": 5
		},
		{
		  "token": "하",
		  "start_offset": 12,
		  "end_offset": 13,
		  "type": "word",
		  "position": 6
		},
		{
		  "token": "고",
		  "start_offset": 13,
		  "end_offset": 14,
		  "type": "word",
		  "position": 7
		},
		{
		  "token": "있",
		  "start_offset": 14,
		  "end_offset": 15,
		  "type": "word",
		  "position": 8
		},
		{
		  "token": "다",
		  "start_offset": 15,
		  "end_offset": 16,
		  "type": "word",
		  "position": 9
		}
	  ]
  }
	```





```json
GET _analyze
{
  "tokenizer": {
    "type": "nori_tokenizer",
    "user_dictionary_rules": ["해물"]
  },
  "text": [
    "동해물과 백두산이"
  ]
}
```
