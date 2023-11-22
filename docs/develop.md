# develop

프로젝트 폴더 구조는 아래 링크에 정리된 글을 토대로 진행  
[golang-standards/project-layout](https://github.com/golang-standards/project-layout)  
[vmware-tanzu/velero](https://github.com/vmware-tanzu/velero)

---

## API

Http 요청, 응답 바디의 JSON 데이터 키는 **snake_case**로 표현한다.  

---

## PostgreSQL

RDB는 **PostgreSQL**을 사용한다.  
[Docker PostgreSQL 설치 및 세팅 글(Connecting-팬도라)](https://judo0179.tistory.com/96)

---

## IDE(vscode)

[1] 테스트 함수 내 로그 출력 안되는 현상  
test 커맨드 실행 시 `-v` 옵션을 줘야 로그가 출력된다.  

```text
// 설정 위치: Visual Studio Code → Settings → User Tab → Extensions → Go → Test Flags
"go.testFlags": [ 
    "-v",
]
```
