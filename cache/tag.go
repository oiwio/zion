package cache

func CacheTagFollowingUsers(tagId string, userId string, timestamp int64) error {
	var (
		err error
	)
	c := pool.Get()
	defer c.Close()
	_, err = c.Do("ZADD", "TagFollowingUsers:"+tagId, timestamp, userId)
	return err
}
