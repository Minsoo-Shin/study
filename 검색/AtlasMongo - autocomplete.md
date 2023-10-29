```mongo
{
  $search: {
    "index": "<index name>", // optional, defaults to "default"
    "autocomplete": {
      "query": "<search-string>",
      "path": "<field-to-search>",
      "tokenOrder": "any|sequential",
      "fuzzy": <options>,
      "score": <options>
    }
  }
}
```

autocomplete 연산자를 통해서 불완전한 문자열에서 일련의 문자가 포함된 단어나 구를 검색할 수 있다. 



# fuzzy
[퍼지 **개념은** 적용 범위가 완전히 고정되는 것이 아니라 상황이나 조건에 따라 상당히 달라질 수 있는 일종의 개념입니다 .](https://en.wikipedia.org/wiki/Concept "개념")
- 상황에 대한 조건들을 여러가지 설정할 수 있다. 
	- `fuzzy.prefixLength` : 무조건 일치해야하는 prefix 갯수 
		- nike를 쿼리를 하면 'nik' 는 무조건 나와야 한다. 
		- 한글은 조금 다르다. 
	- 
#### insert data
```mongo
[
  {
    "title": "nike"
  },
  {
    "title": "nice"
  }
]
```
#### query 
```mongo
{
    $search: {
      index: "test_index",
      autocomplete: {
        query: "nike",
        path: "title",
        tokenOrder: "any",
        fuzzy: {
          prefixLength: 3,
        },
      },
    },
  },
```

nice는 안 나오고, nike만 나온다. 
![[Pasted image 20231024011412.png]]

#### Insert
```mongo
[
  {
    "title": "맨투맨"
  }, 
  {
    "title": "맨티맨"
  }
]
```

#### query
```mongo
[
  {
    $search: {
      index: "test_index",
      autocomplete: {
        query: "맨투",
        path: "title",
        tokenOrder: "any",
        fuzzy: {
          prefixLength: 3,
        },
      },
    },
  },
]
```

'맨투'로 두 글자로 검색하면 잘 나온다. 

![[Pasted image 20231024011957.png]]
'맨ㅌ'로 검색하면 안나온다. 
![[Pasted image 20231024012011.png]]

그 이유는 무엇일까? 
'맨ㅌ'로 검색시에 prefixLength가 3이라 '맨티맨', '맨투맨'이 안나오는 것은 당연하다. 왜냐하면 3글자는 동일해야하기 때문이다. 하지만, 첫번째 쿼리 결과에서는 "맨투"만 입력했을 때에도 맨투맨이 검색이 되었다. 이는 왜 그런 것일까? 



디버깅을 하기 위해서 prefixLength가 어떤 기준으로 숫자를 정하는지 알아야 한다. 

![[Pasted image 20231024015051.png]]

정의하는 내용은 character의 숫자라고 한다. 그럼 한글은 character를 어떤 기준으로 할까? 





















![[Pasted image 20231024004840.png]]
```mongo
[
  {
    $search: {
      index: "autoCompleteProducts", // optional, defaults to "default"
      autocomplete: {
        query: "NIKE Bur",
        path: "brandName",
        tokenOrder: "any",
        fuzzy: {
          // maxEdits: 1,
          prefixLength: 3,
          maxExpansions: 256,
        },
        // "score": <options>
      },
    },
  },
  {
   $project:   {
      "brandName":1,
    },
  },
  {
    $group: {
      _id: "$brandName",
      // unique: {
      //   // $push: "$_id"
      // }
    }
  }
]
```