package db

import (
	"context"
	"testing"
)

func TestCreateFollow(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateFollowParams{
		FollowerID:  account1.ID,
		FollowingID: account2.ID,
	}
	follow, err := testQueries.CreateFollow(context.Background(), arg)

}
