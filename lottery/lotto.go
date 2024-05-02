package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// lottery.exe filename count
	// 명령줄에서 프로그램이 실행될 때 filename과 count를 전달해야 합니다.

	// 명령줄 인수 개수가 3개 미만일 경우 프로그램 종료
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Invalid arguments!\nlottery filename count")
		return
	}

	// filename과 count를 명령줄 인수로부터 얻어옴
	filename := os.Args[1]
	count, err := strconv.Atoi(os.Args[2]) // 명령줄 인수를 int로 변환
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot convert count to integer! count:", count)
		return
	}

	// filename에서 후보자 목록을 읽어옴
	candidates, err := readCandidates(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read candidates file! ", err)
		return
	}

	// 랜덤 시드 설정 현재 시간의 나노초 단위로 표시된 값
	rand.Seed(time.Now().UnixNano())

	// 당첨자 배열 초기화
	winners := make([]string, count)

	// count 만큼 반복하여 당첨자 선정
	for i := 0; i < count; i++ {
		// 랜덤한 인덱스 선택
		n := rand.Intn(len(candidates))
		// 당첨자 배열에 후보자 추가
		winners[i] = candidates[n]
		// 당첨된 후보자를 후보자 목록에서 제외
		candidates = append(candidates[:n], candidates[n+1:]...)
	}

	// 당첨자 출력
	fmt.Println("Winners !!")
	for _, winner := range winners {
		fmt.Println(winner)
	}
}

// filename에서 후보자 목록을 읽어오는 함수
func readCandidates(filename string) ([]string, error) {
	// 파일 열기
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 파일 스캐너 생성
	scanner := bufio.NewScanner(file)

	// 후보자 목록을 저장할 슬라이스 생성
	var candidates []string

	// 파일에서 한 줄씩 읽어와 후보자 목록에 추가
	for scanner.Scan() {
		candidates = append(candidates, scanner.Text())
	}

	// 후보자 목록 반환
	return candidates, nil
}
