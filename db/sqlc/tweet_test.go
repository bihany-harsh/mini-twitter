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
	require.WithinDuration(t, tweet1.UpdatedAt.Time, tweet2.UpdatedAt.Time, time.Second)
	require.Equal(t, tweet1.RetweetID, tweet2.RetweetID)
	require.Equal(t, tweet1.NLikes, tweet2.NLikes)
	require.Equal(t, tweet1.NRetweets, tweet2.NRetweets)
	require.Equal(t, tweet1.NReply, tweet2.NReply)
	require.Equal(t, tweet1.ID, tweet2.ID)
	require.WithinDuration(t, tweet1.CreatedAt, tweet2.CreatedAt, time.Second)
}

func createRandomTweetByAccount(t *testing.T, account Account) (Tweet, Account) {
	arg := CreateTweetParams{
		UserID:    account.ID,
		Content:   util.RandomString(util.RandomInt(20, 279)),
		UpdatedAt: util.RandomTime_Nullable(p),
	}

	nTweets_old := account.NTweets
	tweet, err := testQueries.CreateTweet(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, arg.UserID, tweet.UserID)
	require.Equal(t, arg.Content, tweet.Content)
	require.WithinDuration(t, arg.UpdatedAt.Time, tweet.UpdatedAt.Time, time.Second)
	require.Equal(t, arg.RetweetID, tweet.RetweetID)
	require.Equal(t, arg.NLikes, tweet.NLikes)
	require.Equal(t, arg.NRetweets, tweet.NRetweets)
	require.Equal(t, arg.NReply, tweet.NReply)
	require.NotZero(t, tweet.ID)
	require.NotZero(t, tweet.CreatedAt)

	account_, err := testQueries.AddNTweetsByOne(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, nTweets_old+1, account_.NTweets)

	return tweet, account_
}

func TestCreateTweet(t *testing.T) {
	account := createRandomAccount(t)
	createRandomTweetByAccount(t, account)
}

func TestGetTweetByID(t *testing.T) {
	account := createRandomAccount(t)
	tweet1, _ := createRandomTweetByAccount(t, account)
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
		var newTweet Tweet
		newTweet, account = createRandomTweetByAccount(t, account)
		tweets = append(tweets, newTweet)
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
	tweet1, account := createRandomTweetByAccount(t, account)

	nTweet_old := account.NTweets

	err := testQueries.DeleteTweetByID(context.Background(), tweet1.ID)

	require.NoError(t, err)

	tweet2, err := testQueries.GetTweetByID(context.Background(), tweet1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, tweet2)

	account, err = testQueries.SubtractNTweetsByOne(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, nTweet_old-1, account.NTweets)
}

func TestDeleteTweetsByUserID(t *testing.T) {
	account := createRandomAccount(t)
	n := util.RandomInt(1, 10)

	var tweets []Tweet

	for i := int64(0); i < n; i++ {
		var newTweet Tweet
		newTweet, account = createRandomTweetByAccount(t, account)
		tweets = append(tweets, newTweet)
	}

	err := testQueries.DeleteTweetsByUserID(context.Background(), account.ID)

	require.NoError(t, err)

	for _, tweet := range tweets {
		tweet2, err := testQueries.GetTweetByID(context.Background(), tweet.ID)

		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, tweet2)
	}

	account, err = testQueries.SubtractAllNTweets(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, int32(0), account.NTweets)
}

func TestListTweets(t *testing.T) {
	account := createRandomAccount(t)
	n := 10

	for i := 0; i < n; i++ {
		_, account = createRandomTweetByAccount(t, account)
	}

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
	var tweet1 Tweet
	tweet1, _ = createRandomTweetByAccount(t, account)

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
	require.WithinDuration(t, arg.UpdatedAt.Time, tweet2.UpdatedAt.Time, time.Second)
	require.Equal(t, tweet1.UserID, tweet2.UserID)
	require.Equal(t, tweet1.RetweetID, tweet2.RetweetID)
	require.Equal(t, tweet1.NLikes, tweet2.NLikes)
	require.Equal(t, tweet1.NRetweets, tweet2.NRetweets)
	require.Equal(t, tweet1.NReply, tweet2.NReply)
	require.WithinDuration(t, tweet1.CreatedAt, tweet2.CreatedAt, time.Second)
}
