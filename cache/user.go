package cache

func CacheActiveUsers(userId string, timestamp int64) error {
	var (
		err error
	)
	c := pool.Get()
	defer c.Close()
	_, err = c.Do("ZADD", "ActiveUsers", timestamp, userId)
	return err
}

func CacheUserFollowingTags(userId string, tagId string, timestamp int64) error {
	var (
		err error
	)
	c := pool.Get()
	defer c.Close()
	_, err = c.Do("ZADD", "UserFollowingTags:"+userId, timestamp, tagId)
	return err
}

func CacheUserFollowingTagFeeds(userId string, feedId string, timestamp int64) error {
	var (
		err error
	)
	c := pool.Get()
	defer c.Close()
	_, err = c.Do("ZADD", "UserFollowingTagFeeds:"+userId, timestamp, feedId)
	return err
}
