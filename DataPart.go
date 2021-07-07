package main

type DataPart struct {
	Cursors struct {
		Next     string `json:"next,omitempty"`
		Previous string `json:"previous,omitempty"`
	} `json:"cursors"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results  []*Result
}

type Result struct {
	UID                string `json:"uid,omitempty"`
	Slug               string `json:"slug,omitempty"`
	Name               string `json:"name,omitempty"`
	Staffpickedat      string `json:"staffpickedAt,omitempty"`
	Viewcount          int    `json:"viewCount,omitempty"`
	Likecount          int    `json:"likeCount,omitempty"`
	Animationcount     int    `json:"animationCount,omitempty"`
	Viewerurl          string `json:"viewerUrl,omitempty"`
	Embedurl           string `json:"embedUrl,omitempty"`
	Publiccommentcount int    `json:"publicCommentCount,omitempty"`
	Isdownloadable     bool   `json:"isDownloadable,omitempty"`
	Downloadtype       string `json:"downloadType,omitempty"`
	Downloadcount      int    `json:"downloadCount,omitempty"`
	Ispublished        bool   `json:"isPublished,omitempty"`
	Isrestricted       bool   `json:"isRestricted,omitempty"`
	Publishedat        string `json:"publishedAt,omitempty"`
	Thumbnails         struct {
		UID    string `json:"uid,omitempty"`
		Images []struct {
			UID    string `json:"uid,omitempty"`
			Size   int    `json:"size,omitempty"`
			Width  int    `json:"width,omitempty"`
			URL    string `json:"url,omitempty"`
			Height int    `json:"height,omitempty"`
		} `json:"images,omitempty"`
	} `json:"thumbnails,omitempty"`
	User struct {
		UID         string `json:"uid,omitempty"`
		Displayname string `json:"displayName,omitempty"`
		Account     string `json:"account,omitempty"`
		Username    string `json:"username,omitempty"`
		Avatars     struct {
			UID    string `json:"uid,omitempty"`
			Images []struct {
				Size   int    `json:"size,omitempty"`
				Width  int    `json:"width,omitempty"`
				URL    string `json:"url,omitempty"`
				Height int    `json:"height,omitempty"`
			} `json:"images"`
		} `json:"avatars"`
		Profileurl string `json:"profileUrl,omitempty"`
	} `json:"user,omitempty"`
	Price         int     `json:"price,omitempty"`
	Averagerating float64 `json:"averageRating,omitempty"`
	Reviewcount   int     `json:"reviewCount,omitempty"`
	Instore       bool    `json:"inStore,omitempty"`
	Org           struct {
		UID         string      `json:"uid,omitempty"`
		DisplayName string      `json:"displayName,omitempty"`
		SamlIdpName interface{} `json:"samlIdpName,omitempty"`
		Username    string      `json:"username,omitempty"`
		Avatars     struct {
			UID    string `json:"uid,omitempty"`
			Images []struct {
				Size   int    `json:"size,omitempty"`
				Width  int    `json:"width,omitempty"`
				URL    string `json:"url,omitempty"`
				Height int    `json:"height,omitempty"`
			} `json:"images,omitempty"`
		} `json:"avatars,omitempty"`
	} `json:"org,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}
