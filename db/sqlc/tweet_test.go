package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bihany-harsh/mini-twitter/util"
	"github.com/stretchr/testify/require"
)

func checkEqualTweets(t *testing.T, tweet1, tweet2 Tweet) {
	require.Equal(t, tweet1.UserID, tweet2.UserID)
	require.Equal(t, tweet1.Content, tweet2.Content)
	require.Equal(t, tweet1.UpdatedAt, tweet2.UpdatedAt)
	require.Equal(t, tweet1.RetweetID, tweet2.RetweetID)
	require.Equal(t, tweet1.NLikes, tweet2.NLikes)
	require.Equal(t, tweet1.NRetweets, tweet2.NRetweets)
	require.Equal(t, tweet1.NReply, tweet2.NReply)
	require.Equal(t, tweet1.ID, tweet2.ID)
	require.Equal(t, tweet1.CreatedAt, tweet2.CreatedAt)
}

func createRandomTweetByAccount(t *testing.T, account Account) Tweet {
	arg := CreateTweetParams{
		UserID:    account.ID,
		Content:   util.RandomString(util.RandomInt(20, 279)),
		UpdatedAt: util.RandomTime_Nullable(p),
	}

	tweet, err := testQueries.CreateTweet(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, arg.UserID, tweet.UserID)
	require.Equal(t, arg.Content, tweet.Content)
	require.Equal(t, arg.UpdatedAt, tweet.UpdatedAt)
	require.Equal(t, arg.RetweetID, tweet.RetweetID)
	require.Equal(t, arg.NLikes, tweet.NLikes)
	require.Equal(t, arg.NRetweets, tweet.NRetweets)
	require.Equal(t, arg.NReply, tweet.NReply)
	require.NotZero(t, tweet.ID)
	require.NotZero(t, tweet.CreatedAt)

	return tweet
}

func TestCreateTweet(t *testing.T) {
	account := createRandomAccount(t)
	createRandomTweetByAccount(t, account)
}

func TestGetTweetByID(t *testing.T) {
	account := createRandomAccount(t)
	tweet1 := createRandomTweetByAccount(t, account)
	tweet2, err := testQueries.GetTweetByID(context.Background(), tweet1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, tweet2)
	checkEqualTweets(t, tweet1, tweet2)
}

func TestGetTweetsByUserID(t *testing.T) {
	account := createRandomAccount(t)
	n := util.RandomInt(1, 10)

	var tweets []Tweet

	for i := int64(0); i < n; i++ {
		tweets = append(tweets, createRandomTweetByAccount(t, account))
	}

	arg := GetTweetsByUserIDParams{
		UserID: account.ID,
		Limit:  int32(n),
		Offset: 0,
	}

	tweets2, err := testQueries.GetTweetsByUserID(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tweets2)
	require.Equal(t, len(tweets), len(tweets2))
}

func TestDeleteTweetByID(t *testing.T) {
	account := createRandomAccount(t)
	tweet1 := createRandomTweetByAccount(t, account)
	err := testQueries.DeleteTweetByID(context.Background(), tweet1.ID)

	require.NoError(t, err)

	tweet2, err := testQueries.GetTweetByID(context.Background(), tweet1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, tweet2)
}

func TestDeleteTweetsByUserID(t *testing.T) {
	account := createRandomAccount(t)
	n := util.RandomInt(1, 10)

	var tweets []Tweet

	for i := int64(0); i < n; i++ {
		tweets = append(tweets, createRandomTweetByAccount(t, account))
	}

	err := testQueries.DeleteTweetsByUserID(context.Background(), account.ID)

	require.NoError(t, err)

	for _, tweet := range tweets {
		tweet2, err := testQueries.GetTweetByID(context.Background(), tweet.ID)

		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, tweet2)
	}
}

func TestListTweets(t *testing.T) {
	tweets, err := testQueries.ListTweets(context.Background(), ListTweetsParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.NotEmpty(t, tweets)
	require.Len(t, tweets, 5)
}

func TestUpdateTweetByID(t *testing.T) {
	account := createRandomAccount(t)
	tweet1 := createRandomTweetByAccount(t, account)

	arg := UpdateTweetByIDParams{
		ID:      tweet1.ID,
		Content: util.RandomString(util.RandomInt(20, 279)),
		UpdatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	tweet2, err := testQueries.UpdateTweetByID(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tweet2)
	require.Equal(t, arg.ID, tweet2.ID)
	require.Equal(t, arg.Content, tweet2.Content)
	require.Equal(t, arg.UpdatedAt, tweet2.UpdatedAt)
	require.Equal(t, tweet1.UserID, tweet2.UserID)
	require.Equal(t, tweet1.RetweetID, tweet2.RetweetID)
	require.Equal(t, tweet1.NLikes, tweet2.NLikes)
	require.Equal(t, tweet1.NRetweets, tweet2.NRetweets)
	require.Equal(t, tweet1.NReply, tweet2.NReply)
	require.Equal(t, tweet1.CreatedAt, tweet2.CreatedAt)
}
