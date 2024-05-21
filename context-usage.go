package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type User struct {
	id   string
	name string
}

type Response struct {
	value User
	err   error
}

func contextDemo() {
	ctx := context.Background()
	val, err := getUserData(ctx)
	if err != nil {
		fmt.Println("Error fetching user", err)
	} else {
		fmt.Println("User info ", val)
	}

}

func getUserData(ctx context.Context) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	resch := make(chan Response)

	go func() {
		userData, err := thirdParyAsyncFunction()
		resch <- Response{
			value: userData,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return User{}, ctx.Err()
		case res := <-resch:
			return res.value, nil
		}
	}
}

func thirdParyAsyncFunction() (User, error) {
	time.Sleep(time.Millisecond * 300)
	return User{id: strconv.Itoa(int(time.Now().Unix())), name: "Sachin"}, nil
}
