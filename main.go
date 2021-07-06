package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Print("반갑습니다! KG 블록 체인을 경험 하시겠어요?\n\n")
	fmt.Print("뒤에 실행시킬 명령어를 붙혀서 입력해주세요!\n\n")
	fmt.Println("- explorer : HTML Explorer를 실행합니다.")
	fmt.Println("- rest : REST API를 실행합니다.(권장)")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		fmt.Println("explorer 실행!")
	case "rest":
		fmt.Println("rest API 실행!")
	default:
		usage()
	}
}
