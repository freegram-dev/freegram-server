package xlog

import (
	"fmt"
	"log"
)

func Debugf(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

func Errorf(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

func Fatalf(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
	log.Fatalf(format+"\n", args...)
}
