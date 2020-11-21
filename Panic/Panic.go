package main

import "os"

func main() {
	// 예상치 못한 에러를 체크하기 위해 사이트 전체에서 패닉을 사용합니다.
	// 이는 사이트에서 패닉을 하기 위해 설계된 유일한 프로그램입니다.
	panic("THIS IS PANIC ! ! ! ! ! ! ! ! ")

	// 패닉의 일반적인 사용은 어떤 함수가 어떻게 처리할지 모르는 (또는 하고 싶지 않은) 에러값을 반환했을때 중단을 하기 위함입니다.
	// 다음은 파일을 생성할 때 예상치 못한 에러가 발생할 경우 패니킹(panicking)을 하는 예시입니다.

	_, err := os.Create("/이런/경로도/없거니와/파일도/없어서/에러를/뱉는다.error")

	if err != nil {
		panic(err)
	}
}
