package json

type BoardsJSON struct {
	Boards []struct {
		Board           string `json:"board"`
		Title           string `json:"title"`
		WsBoard         int    `json:"ws_board"`
		PerPage         int    `json:"per_page"`
		Pages           int    `json:"pages"`
		MaxFilesize     int    `json:"max_filesize"`
		MaxWebmFilesize int    `json:"max_webm_filesize"`
		MaxCommentChars int    `json:"max_comment_chars"`
		MaxWebmDuration int    `json:"max_webm_duration"`
		BumpLimit       int    `json:"bump_limit"`
		ImageLimit      int    `json:"image_limit"`
		Cooldowns       struct {
			Threads int `json:"threads"`
			Replies int `json:"replies"`
			Images  int `json:"images"`
		} `json:"cooldowns"`
		MetaDescription string `json:"meta_description"`
		IsArchived      int    `json:"is_archived,omitempty"`
		Spoilers        int    `json:"spoilers,omitempty"`
		CustomSpoilers  int    `json:"custom_spoilers,omitempty"`
		ForcedAnon      int    `json:"forced_anon,omitempty"`
		UserIds         int    `json:"user_ids,omitempty"`
		CountryFlags    int    `json:"country_flags,omitempty"`
		CodeTags        int    `json:"code_tags,omitempty"`
		WebmAudio       int    `json:"webm_audio,omitempty"`
		MinImageWidth   int    `json:"min_image_width,omitempty"`
		MinImageHeight  int    `json:"min_image_height,omitempty"`
		Oekaki          int    `json:"oekaki,omitempty"`
		SjisTags        int    `json:"sjis_tags,omitempty"`
		TextOnly        int    `json:"text_only,omitempty"`
		RequireSubject  int    `json:"require_subject,omitempty"`
		TrollFlags      int    `json:"troll_flags,omitempty"`
		MathTags        int    `json:"math_tags,omitempty"`
	} `json:"boards"`
	TrollFlags struct {
		AC string `json:"AC"`
		AN string `json:"AN"`
		BL string `json:"BL"`
		CF string `json:"CF"`
		CM string `json:"CM"`
		CT string `json:"CT"`
		DM string `json:"DM"`
		EU string `json:"EU"`
		FC string `json:"FC"`
		GN string `json:"GN"`
		GY string `json:"GY"`
		JH string `json:"JH"`
		KN string `json:"KN"`
		MF string `json:"MF"`
		NB string `json:"NB"`
		NZ string `json:"NZ"`
		PC string `json:"PC"`
		PR string `json:"PR"`
		RE string `json:"RE"`
		TM string `json:"TM"`
		TR string `json:"TR"`
		UN string `json:"UN"`
		WP string `json:"WP"`
	} `json:"troll_flags"`
}
