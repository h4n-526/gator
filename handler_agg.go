package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("agg expects a single argument: time_between_reqs (e.g. 1s, 1m, 1h)")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Printf("could not get next feed to fetch: %v", err)
		return
	}

	if err := s.db.MarkFeedFetched(ctx, feed.ID); err != nil {
		log.Printf("could not mark feed %s as fetched: %v", feed.Name, err)
		return
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		log.Printf("could not fetch feed %s: %v", feed.Name, err)
		return
	}

	fmt.Printf("Feed: %s (%s)\n", feed.Name, feed.Url)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("  - %s\n", item.Title)
	}
}
