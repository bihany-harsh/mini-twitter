CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "profile_picture_url" varchar,
  "bio" varchar(500),
  "last_login" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "is_admin" boolean NOT NULL DEFAULT false,
  "is_active" boolean NOT NULL DEFAULT true,
  "last_deactivated_at" timestamptz,
  "n_followers" int NOT NULL DEFAULT 0,
  "n_following" int NOT NULL DEFAULT 0,
  "n_tweets" int NOT NULL DEFAULT 0
);

CREATE TABLE "tweets" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "content" varchar(280) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "is_deleted" boolean NOT NULL DEFAULT false,
  "retweet_id" bigint,
  "n_likes" int NOT NULL DEFAULT 0,
  "n_retweets" int NOT NULL DEFAULT 0,
  "n_reply" int NOT NULL DEFAULT 0
);

CREATE TABLE "follows" (
  "follower_id" bigint NOT NULL,
  "following_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("follower_id", "following_id")
);

CREATE TABLE "likes" (
  "user_id" bigint NOT NULL,
  "tweet_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "tweet_id")
);

CREATE TABLE "retweets" (
  "user_id" bigint NOT NULL,
  "tweet_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "tweet_id")
);

CREATE TABLE "replies" (
  "id" bigserial PRIMARY KEY,
  "tweet_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "content" varchar(280) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE INDEX ON "accounts" ("username");

CREATE INDEX ON "accounts" ("email");

CREATE INDEX ON "tweets" ("user_id");

CREATE INDEX ON "follows" ("following_id");

CREATE INDEX ON "follows" ("follower_id");

CREATE INDEX ON "likes" ("user_id");

CREATE INDEX ON "likes" ("tweet_id");

CREATE INDEX ON "retweets" ("user_id");

CREATE INDEX ON "retweets" ("tweet_id");

CREATE INDEX ON "replies" ("user_id", "tweet_id");

ALTER TABLE "tweets" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("id");

ALTER TABLE "tweets" ADD FOREIGN KEY ("retweet_id") REFERENCES "tweets" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("follower_id") REFERENCES "accounts" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("following_id") REFERENCES "accounts" ("id");

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("id");

ALTER TABLE "likes" ADD FOREIGN KEY ("tweet_id") REFERENCES "tweets" ("id");

ALTER TABLE "retweets" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("id");

ALTER TABLE "retweets" ADD FOREIGN KEY ("tweet_id") REFERENCES "tweets" ("id");

ALTER TABLE "replies" ADD FOREIGN KEY ("tweet_id") REFERENCES "tweets" ("id");

ALTER TABLE "replies" ADD FOREIGN KEY ("user_id") REFERENCES "accounts" ("id");
