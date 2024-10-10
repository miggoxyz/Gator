package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/miggoxyz/gator/internal/database"
)
func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.User)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %s <name> <url>", cmd.Name)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
		Url: cmd.Args[1],
		UserID: user.ID,
	})

	if err != nil {
		return fmt.Errorf("could not parse feed %w", err)
	}
	fmt.Println(feed)
	return nil
}