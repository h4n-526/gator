package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, _ command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
