// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: follow.sql

package db

import (
	"context"
)

const createFollow = `-- name: CreateFollow :one
INSERT INTO follows (
    follower_id, following_id
) VALUES (
    $1, $2
) RETURNING follower_id, following_id, created_at
`

type CreateFollowParams struct {
	FollowerID  int64 `json:"follower_id"`
	FollowingID int64 `json:"following_id"`
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error) {
	row := q.db.QueryRowContext(ctx, createFollow, arg.FollowerID, arg.FollowingID)
	var i Follow
	err := row.Scan(&i.FollowerID, &i.FollowingID, &i.CreatedAt)
	return i, err
}

const deleteFollow = `-- name: DeleteFollow :exec
DELETE FROM follows WHERE follower_id = $1 AND following_id = $2
`

type DeleteFollowParams struct {
	FollowerID  int64 `json:"follower_id"`
	FollowingID int64 `json:"following_id"`
}

func (q *Queries) DeleteFollow(ctx context.Context, arg DeleteFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFollow, arg.FollowerID, arg.FollowingID)
	return err
}

const getFollowers = `-- name: GetFollowers :many
SELECT follower_id, following_id, created_at FROM follows WHERE following_id = $1 ORDER BY follower_id LIMIT $2 OFFSET $3
`

type GetFollowersParams struct {
	FollowingID int64 `json:"following_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) GetFollowers(ctx context.Context, arg GetFollowersParams) ([]Follow, error) {
	rows, err := q.db.QueryContext(ctx, getFollowers, arg.FollowingID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follow
	for rows.Next() {
		var i Follow
		if err := rows.Scan(&i.FollowerID, &i.FollowingID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowing = `-- name: GetFollowing :many
SELECT follower_id, following_id, created_at FROM follows WHERE follower_id = $1 ORDER BY following_id LIMIT $2 OFFSET $3
`

type GetFollowingParams struct {
	FollowerID int64 `json:"follower_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) GetFollowing(ctx context.Context, arg GetFollowingParams) ([]Follow, error) {
	rows, err := q.db.QueryContext(ctx, getFollowing, arg.FollowerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follow
	for rows.Next() {
		var i Follow
		if err := rows.Scan(&i.FollowerID, &i.FollowingID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
