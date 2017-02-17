package db

type (

	//commmon response
	Response struct {
		Success bool `json:"success"`
		Error   int  `json:"error"`
	}

	//user response
	UserResponse struct {
		Response
		Token string  `json:"token,omitempty"`
		User  *User   `json:"user,omitempty"`
		Users []*User `json:"users,omitempty"`
	}

	FeedResponse struct {
		Response
		Token string  `json:"token,omitempty"`
		Feed  *Feed   `json:"feed,omitempty"`
		Feeds []*Feed `json:"feeds,omitempty"`
	}

	FriendshipResponse struct {
		Response
		UserId   string `json:"friendId,omitempty"`
		NickName string `json:"friendName,omitempty"`
		Avatar   string `json:"avatar,omitempty"`
		Relation int    `json:"relation"`
	}

	MusicResponse struct {
		Response
		Music  *Music   `json:"music,omitempty"`
		Musics []*Music `json:"musics,omitempty"`
	}

	CommentResponse struct {
		Response
		Comment  *Comment   `json:"comment,omitempty"`
		Comments []*Comment `json:"comments,omitempty"`
	}

	TagResponse struct {
		Response
		Tag *Tag `json:"tag,omitempty"`
	}
)
