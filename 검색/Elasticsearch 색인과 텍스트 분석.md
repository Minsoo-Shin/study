풀 텍스트 검색을 하기 위해 데이터를 검색에 맞게 가공하는 작업이 필요하다. 여기서 elasticsearch는 데이터를 저장하는 과정에서 어떻게 데이터를 처리하고, 색인하는지 확인해보자..
https://esbook.kimjmin.net/06-text-analysis/6.1-indexing-data


# 역 인덱스(Inverted Index)

## 관계형 데이터베이스 vs Elasticsearch

일반적으로, 관계형 데이터베이스는 아래와 같이 저장하고, 
데이터 한줄 한줄 like 검색으로 데이터를 찾는다. 
#### 데이터가 쌓이면 검색할 대상이 늘어나고 시간도 오래걸리게 된다. 

![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntG3a3EAa6IULTKkT9%2F-LntGEeCVTqaRzzzHFem%2F6.1-01.png?alt=media&token=ed349e28-8215-43b2-b049-857f68a19d47)![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-Lo--uX4jUMQUTUvBgeF%2F-LntIdlIDXEduASJCXRm%2F6.1-02.png?alt=media&token=158baec3-905d-4f92-8824-cac8d0239756)


반면, Elasticsearch는 원본과 역 인덱스 데이터를 저장하게 된다. 데이터가 늘어나도 찾는 테이블의 행이 늘어나는 것이 아니므로 큰 속도 저하없이 빠른 검색이 가능한 것이다. 

![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntL_BGpuFbNXy_sFtK%2F-LntLbibpXHABupWvXtu%2F6.1-03.png?alt=media&token=d2726f20-a7ea-4219-bcb0-340cbe1d21f1)

아래와 같이 `fox`를 full text search를 하게 되면, 바로 DocumentID를 얻을 수 있다. 
여기서 fox를 Term이라고 하고 이렇게 저장하는 과정을 색인(Indexing)이라고 한다. 

![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntS3nPGQlmuCtaIVJt%2F-LntS6M5Y65sfxz435rP%2F6.1-04.png?alt=media&token=b8738d24-462e-45d4-8c64-4ed78ceaab15)




# 텍스트 분석 - Text Analysis
---
다음은 Elasticsearch에서 저장되는 도큐먼트는 모든 문자열(text)필드 별로 역 인덱스를 생성한다. 
![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntVKcOQaeoQJjPnbpP%2F-LntVxXxMRFyBg4lJyRx%2F6.2-01.png?alt=media&token=7926be1d-e99f-4f5c-8bed-f80708f55931)

Elasticsearch는 문자열 필드가 저장될 때 데이터에서 검색어 토큰을 저장하기 위해 여러 단계를 거친다. 이 전체 과정을 텍스트 분석(Text Analysis)이라고 하고

이 전체 과정을 텍스트 분석 (Text Analysis)라고 하며, 이 과정을 처리하는 기능을 애널라이저(Analyzer)라고 한다. 

**Analyzer > Character Filter = Tokenizer = Token Filter**
![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntYrdKmTe441TqYAJl%2F-LntZ63SAIfHu6Q_OgzJ%2F6.2-02.png?alt=media&token=52213afe-e6ab-4bc2-b9e0-20027542a79e)


### **Analyzer > Character Filter = Tokenizer = Token Filter**

1. Charater Filter
전체 문장에서 특정 문자를 대치하거나 제거하는데 이 과정을 담당하는 기능이 캐릭터 필터이다. 
2. 문장에 속한 단어들을 텀 단위로 하나씩 분리 해 내는 처리 과정을 거치는데 이를 담당하는 기능이 토크나이저
	  - 토크나이저는 반드시 1개만 적용 가능
	  - `whitespace`토크나이저를 이용해서 공백을 기준으로 텀을 분리
![assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntL_BGpuFbNXy_sFtK%2F-LntLbibpXHABupWvXtu%2F6.1-03.png?alt=media&token=d2726f20-a7ea-4219-bcb0-340cbe1d21f1](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntL_BGpuFbNXy_sFtK%2F-LntLbibpXHABupWvXtu%2F6.1-03.png?alt=media&token=d2726f20-a7ea-4219-bcb0-340cbe1d21f1)

3. 분리된 텀 들을 하나씩 가공하는 과정을 거치는데 이 과정을 담당하는 기능이 토큰 필터이다. 
	 - 잘라 놓은거는 정리를 해야 써먹을 수 있을 것이다. 
	- lowercase 토큰 필터로 모두 소문자로 변경한다. 
	- 토큰 필터를 통해서 일치하게 되어 같은 텀이 된 토큰들을 모두 하나로 병합된다.

####  token filter (lowercase)
	![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntbFF1Cbw9kue34dxC%2F-LntbHMfIKRZOiCl7KmN%2F6.2-03.png?alt=media&token=91afddea-ec2e-4989-a751-20a689374b08)


아래와 같이 변경됨. 
![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntbFF1Cbw9kue34dxC%2F-LntcLPw_rlidqO38odU%2F6.2-04.png?alt=media&token=52d756b7-9533-492d-999d-0640f775bcd7)


텀 중에는 검색어로서의 가치가 없는 단어들이 있는데, 이런 단어를 불용어(stopword)라고 한다. 
영어에서 a,an,at,be,but,by,do,for,i,no.the,to...등의 단어들은 불용러로 간주되어 검색어 토큰에서 제외된다. 
stop토큰 필터를 적용하면 우리가 만드는 역 인덱스에서 the가 사라진다. 

![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntdTZrPbJB3nIxslS_%2F-LntdYna6xmoecLuIbcL%2F6.2-05.png?alt=media&token=4e537bb0-76a1-4b98-877d-ceabe3e71bd9)


#### token filter (snowball)

영어에서는 형태소 분석을 위해 `snowball`토큰 필터를 주로 사용하는데 이 필터는 ~s, ~ing 등을 제거한다. 
그리고 happy, lazy 와 같은 단어들은 happiness, laziness와 같은 형태로도 사용되기 때문에 **~y** 를 **~i** 로 변경합니다. `snowball` 토큰 필터를 적용하고 나면 **jumps**와 **jumping**은 모두 **jump**로 변경되고, 동일하게 jump 로 되었기 때문에 하나의 텀으로 병합됩니다.

![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-Lntet01jJphNCVzIo7v%2F-Lntf24nCf5pgDeswY5d%2F6.2-06.png?alt=media&token=4140c045-ee24-443f-b927-84cfdad57a9f)

snowball 형태소 분석 적용 후 텀 병합



#### token filter (synonym)

필요에 따라서 동의어를 추가해 주기도 한다. synonym 토큰 필터를 사용하여 quick 텀에 동의어로 fast를 지정하면 fast로 검색했을 때도 quick이 검색이 된다. 

![](https://1535112035-files.gitbook.io/~/files/v0/b/gitbook-legacy-files/o/assets%2F-Ln04DaYZaDjdiR_ZsKo%2F-LntgOPNccbFlmVJP9gx%2F-LntgR3I2LDe35aKI--u%2F6.2-07.png?alt=media&token=b758aac1-6f16-4a8f-8649-bd5a131adbbc)


# 애널라이저 - Analyzer
---
### _analyze API

문장이 잘 분석되는지를 확인하는 방법이다.

```
GET _analyze
{
  "text": "The quick brown fox jumps over the lazy dog",
  "tokenizer": "whitespace",
  "filter": [
    "lowercase",
    "stop",
    "snowball"
  ]
}
```


- tokenizer: "whitespace"
- filter: lowercase, stop, snowball

- analyzer: snowball을 사용하면 `fox`, `jump`, `lazi` 등의 단어가 검색 텀으로 저장된다.

