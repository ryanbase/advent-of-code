package utils

import (
	"fmt"
	"regexp"
	"runtime"
	"time"
)

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	fmt.Printf("%s completed in %s\n", name, elapsed)
}
