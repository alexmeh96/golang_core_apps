package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respCh := make(chan Response)

	go func() {
		val, err := fetchExternalSlowApi()
		respCh <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetcheng data finished on timeout")
		case resp := <-respCh:
			return resp.value, resp.err
		}
	}
}

func fetchExternalSlowApi() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 555, nil
}

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))
}
