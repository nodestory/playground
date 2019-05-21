package main

import (
	"fmt"
	"github.com/bitmark-inc/go-argon2"
)

const (
	digestMode        = argon2.ModeArgon2d
	digestMemory      = 1 << 17 // 128 MiB
	digestParallelism = 1
	digestIterations  = 4
	digestVersion     = argon2.Version13
	digestLength      = 32
)

func main() {
	record := "1234567890"

	context := &argon2.Context{
		Iterations:  digestIterations,
		Memory:      digestMemory,
		Parallelism: digestParallelism,
		HashLen:     digestLength,
		Mode:        digestMode,
		Version:     digestVersion,
	}

	hash, err := argon2.Hash(context, []byte(record), []byte(record))
	fmt.Printf("%x\n", hash)
	fmt.Println(err)
}
