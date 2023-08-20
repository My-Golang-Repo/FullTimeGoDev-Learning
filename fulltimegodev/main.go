package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func thirdPartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 90)
	return "result id", nil
}

func fetchUserID(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	val := ctx.Value("username")
	fmt.Println("The username is ->", val)

	type result struct {
		userID string
		err    error
	}

	resultch := make(chan result, 1)

	go func() {
		res, err := thirdPartyHTTPCall()
		resultch <- result{
			userID: res,
			err:    err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userID, res.err
	}
}

func stripeAPICall() (int, error) {
	time.Sleep(time.Millisecond * 400)
	return 100, nil
}

func fetchAPIPaymentOfStripe() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	type result struct {
		paymentCode int
		err         error
	}

	resultch := make(chan result, 1)

	go func() {
		res, err := stripeAPICall()
		resultch <- result{
			paymentCode: res,
			err:         err,
		}
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case res := <-resultch:
		return res.paymentCode, res.err
	}

}

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "IJN Yamato")
	userID, err := fetchUserID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	paymentCode, err := fetchAPIPaymentOfStripe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The result took %v -> %+v\n", time.Since(start), userID)
	fmt.Printf("The result took %v -> %+v\n", time.Since(start), paymentCode)
}
