package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bihany-harsh/mini-twitter/util"
	"github.com/stretchr/testify/require"
)

func createRandomReplyToTweet(t *testing.T, tweet *Tweet) Reply {
	n_replies_old := tweet.NReply

	account := createRandomAccount(t)
	arg := CreateReplyParams{
		TweetID:   tweet.ID,
		UserID:    account.ID,
		Content:   util.RandomString(util.RandomInt(1, 279)),
		UpdatedAt: util.RandomTime_Nullable(p),
	}
	reply, err := testQueries.CreateReply(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reply)
	require.Equal(t, reply.TweetID, arg.TweetID)
	require.Equal(t, reply.UserID, arg.UserID)
	require.NotZero(t, reply.CreatedAt)

	*tweet, err = testQueries.AddNReplyByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, tweet.NReply, n_replies_old+1)

	return reply
}

func createRandomReplyToTweetByAccount(t *testing.T, account Account, tweet *Tweet) Reply {
	n_replies_old := tweet.NReply

	arg := CreateReplyParams{
		TweetID:   tweet.ID,
		UserID:    account.ID,
		Content:   util.RandomString(util.RandomInt(1, 279)),
		UpdatedAt: util.RandomTime_Nullable(p),
	}

	reply, err := testQueries.CreateReply(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reply)
	require.Equal(t, reply.TweetID, arg.TweetID)
	require.Equal(t, reply.UserID, arg.UserID)
	require.NotZero(t, reply.CreatedAt)

	*tweet, err = testQueries.AddNReplyByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, tweet.NReply, n_replies_old+1)

	return reply
}

func TestCreateReply(t *testing.T) {
	account := createRandomAccount(t)
	tweet, account := createRandomTweetByAccount(t, account)

	createRandomReplyToTweet(t, &tweet)
}

func TestDeleteReply(t *testing.T) {
	account := createRandomAccount(t)
	tweet, account := createRandomTweetByAccount(t, account)

	reply := createRandomReplyToTweet(t, &tweet)

	err := testQueries.DeleteReply(context.Background(), reply.ID)
	require.NoError(t, err)

	tweet, err = testQueries.SubtractNReplyByOne(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)
	require.Equal(t, tweet.NReply, int32(0))
}

func TestGetReply(t *testing.T) {
	account := createRandomAccount(t)
	tweet, account := createRandomTweetByAccount(t, account)

	reply := createRandomReplyToTweet(t, &tweet)

	reply2, err := testQueries.GetReply(context.Background(), reply.ID)
	require.NoError(t, err)
	require.NotEmpty(t, reply2)
	require.Equal(t, reply2.ID, reply.ID)
	require.Equal(t, reply2.TweetID, reply.TweetID)
	require.Equal(t, reply2.UserID, reply.UserID)
	require.Equal(t, reply2.Content, reply.Content)
	require.Equal(t, reply2.CreatedAt, reply.CreatedAt)
	require.Equal(t, reply2.UpdatedAt, reply.UpdatedAt)
}

func TestGetrepliesByTweetID(t *testing.T) {
	account := createRandomAccount(t)
	tweet, account := createRandomTweetByAccount(t, account)

	n_replies := util.RandomInt(1, 10)
	for i := int64(0); i < n_replies; i++ {
		createRandomReplyToTweet(t, &tweet)
	}

	arg := GetRepliesByTweetIDParams{
		TweetID: tweet.ID,
		Limit:   int32(n_replies),
		Offset:  0,
	}
	replies, err := testQueries.GetRepliesByTweetID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, replies)
	require.Equal(t, int64(len(replies)), n_replies)
	for _, reply := range replies {
		require.NotEmpty(t, reply)
		require.Equal(t, reply.TweetID, arg.TweetID)
	}
}

func TestGetRepliesByUserID(t *testing.T) {
	account2 := createRandomAccount(t)
	n := util.RandomInt(1, 10)

	var replies []Reply

	for i := int64(0); i < n; i++ {
		account1 := createRandomAccount(t)
		tweet, account1 := createRandomTweetByAccount(t, account1)
		newReply := createRandomReplyToTweetByAccount(t, account2, &tweet)
		replies = append(replies, newReply)
	}

	arg := GetRepliesByUserIDParams{
		UserID: account2.ID,
		Limit:  int32(n),
		Offset: 0,
	}

	replies2, err := testQueries.GetRepliesByUserID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, replies2)
	require.Equal(t, int64(len(replies2)), n)
}

func TestUpdateReply(t *testing.T) {
	account1 := createRandomAccount(t)
	tweet, account1 := createRandomTweetByAccount(t, account1)

	reply := createRandomReplyToTweet(t, &tweet)

	arg := UpdateReplyParams{
		ID:      reply.ID,
		Content: util.RandomString(util.RandomInt(1, 279)),
		UpdatedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	reply2, err := testQueries.UpdateReply(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reply2)
	require.Equal(t, reply2.ID, reply.ID)
	require.Equal(t, reply2.TweetID, reply.TweetID)
	require.Equal(t, reply2.UserID, reply.UserID)
	require.Equal(t, reply2.Content, arg.Content)
	require.WithinDuration(t, reply2.CreatedAt, reply.CreatedAt, time.Second)
	require.WithinDuration(t, reply2.UpdatedAt.Time, arg.UpdatedAt.Time, time.Second)
}
