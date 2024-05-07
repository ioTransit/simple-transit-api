package cli

import (
	"fmt"
	jobs "go-gtfs-server/app/job"
	"go-gtfs-server/utils"
	"os"
)

var Args = []string{"--load"}

func CliRouter() {
	Args = os.Args[1:]
	fmt.Println(Args)
	update := utils.ContainsString(Args, "--update")
	if update {
		fmt.Println("load it")
		jobs.Load()
	}
}
