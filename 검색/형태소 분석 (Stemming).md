명사 뒤에 ~s, ~ness 등이 붙거나 동사 뒤에 ~ing, ~ed 등 붙는 등 변화가 많다. 검색을 할 때는 이런 부분과 관계없이 잘 찾아야한다. 

하지만 `word`를 검색해도, `words`가 포함된 Document가 검색되지 않는다. 그것은 words는 word와 별개로term을 갖기 때문이다. 


### 역 인덱스
- word : doc1
- words : doc2

### 원하는 역 인덱스
- **word : doc1, doc2** 

어떻게 할 수 있을까? Elasticsearch는 형태소 분석기들을 지원한다. 가장 많이 알려진 형태소 분석기는 아래와 같다. 
- snowball
- nori (한글) : 카카오에서 만들 형태소 분석기라고 한다. 


# 커뮤니티 한글 형태소 분석기 - 아리랑, 은전한닢, Open Korean Text

Elasticsearch 버전이 올라가 버전이 변경되면 사용이 불가능하고 버그가 있어도 누군가 쉽게 고치기 어려운 문제가 있음



# Nori 
Nori 는 **은전한닢**에서 사용하는 **mecab-ko-dic** 사전을 재 가공 하여 사용하고 있습니다.
https://www.elastic.co/guide/en/elasticsearch/plugins/current/analysis-nori.html

- **user_dictionary** : 사용자 사전이 저장된 파일의 경로를 입력합니다.
- **user_dictionary_rules** : 사용자 정의 사전을 배열로 입력합니다.
- **decompound_mode** : 합성어의 저장 방식을 결정합니다. 다음 3개의 값을 사용 가능합니다.
    - `none` : 어근을 분리하지 않고 완성된 합성어만 저장합니다.
    - `discard` (디폴트) : 합성어를 분리하여 각 어근만 저장합니다.
    - `mixed` : 어근과 합성어를 모두 저장합니다

```
PUT my_nori
{
  "settings": {
    "analysis": {
      "tokenizer": {
        "my_nori_tokenizer": {
          "type": "nori_tokenizer",
          "user_dictionary_rules": [
            "해물"
          ]
        }
      }
    }
  }
}
```

`동해물과` 을 분석하면 아래와 같이 토크나이징을 한다. 
- `동`
- `해물`
- `과`

user_dictionary_rules 없는 순수 nori_tokenizer `동해` + `물` + `과`와는 다른 결과가 나온다. 

### 아래를 테스트 해보자
- **decompound_mode** : 합성어의 저장 방식을 결정합니다. 다음 3개의 값을 사용 가능합니다.
    - `none` : 어근을 분리하지 않고 완성된 합성어만 저장합니다.
    - `discard` (디폴트) : 합성어를 분리하여 각 어근만 저장합니다.
    - `mixed` : 어근과 합성어를 모두 저장합니다
```elasticsearch
PUT my_nori
{
  "settings": {
    "analysis": {
      "tokenizer": {
        "nori_none": {
          "type": "nori_tokenizer",
          "decompound_mode": "none"
        },
        "nori_discard": {
          "type": "nori_tokenizer",
          "decompound_mode": "discard"
        },
        "nori_mixed": {
          "type": "nori_tokenizer",
          "decompound_mode": "mixed"
        }
      }
    }
  }
}
```

- `nori_none` : `백두산` + `이`
- `nori_tokenizer`: `백두` + `산` + `이`
- `nori_tokenizer`: `백두산` + `백두` + `산` + `이`


### nori_part_of_speech 와 품사 정보

아래와 같이 수사(NR)을 제거하는  토큰 필터를 지정할 수 있다. 
```
PUT my_pos
{
  "settings": {
    "index": {
      "analysis": {
        "filter": {
          "my_pos_f": {
            "type": "nori_part_of_speech",
            "stoptags": [
              "NR"
            ]
          }
        }
      }
    }
  }
}
```

`다섯아이가` = `아이` + `가`



http://kkma.snu.ac.kr/documents/?doc=postag
![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LoinpqY1xA7ock1sc6i%2F-Loioly2sAhomKoXMv2-%2F6.7.2-02.png?alt=media&token=47fae11e-c38e-4dff-92e8-64515e37f565)





### nori_readingform
한자 관련 필터이다. 






