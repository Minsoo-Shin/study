동의어를 설정하는 옵션은 `synonyms` 항목에서 직접 동의어 목록을 입력하는 방법과 동의어 사전 파일을 만들어 `synonyms_path` 로 지정하는 방법이 있습니다. 동의어 사전 명시 규칙에는 다음의 것들이 있습니다.

- `"A, B => C"` : 왼쪽의 A, B 대신 오른쪽의 C 텀을 저장합니다. A, B 로는 C 의 검색이 가능하지만 C 로는 A, B 가 검색되지 않습니다.
    

- `"A, B"` : A, B 각 텀이 A 와 B 두개의 텀을 모두 저장합니다. A 와 B 모두 서로의 검색어로 검색이 됩니다.





## Elastic cloud 사용자 사전
해당 파일들은 dictionaries 폴더 아래에 저장이 되어 있어야 합니다.

/dictionaries

    |---- stopwords.txt

    |---- user_dictionary.txt

    |---- synonyms.txt

파일을 Cloud에 적용하려면 다음과 같은 작업을 수행합니다.

1. dictionaries.zip 파일 생성 (위 폴더를 .zip파일화)

2. Elastic Cloud 접속 (cloud.elastic.co)

3. Deployment -> Features -> Extensions 접근 후 'Create extension' 클릭

4. extension 설정 정보 작성

  - 이름(Name), 설명(Description) 작성

  - Type 설정 : 'A bundle containing a dictionary or script'

  - 1번 파일 import (Bundle files)

Cloud UI 메인화면에서 다음을 수행하여 최종 적용하기

5. Deployment -> Edit 접근 후 'Manage plugins, extensions and settings' 클릭

6. 4번에 작성한 이름의 extensions 선택하고 Save -> Confirm

7. 자동으로 Rolling Update 수행 (수 분 소요) 후 적용 완료