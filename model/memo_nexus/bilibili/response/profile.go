package response

type VipLabel struct {
	Path                  string `json:"path"`
	Text                  string `json:"text"`
	LabelTheme            string `json:"label_theme"`
	TextColor             string `json:"text_color"`
	BgStyle               int    `json:"bg_style"`
	BgColor               string `json:"bg_color"`
	BorderColor           string `json:"border_color"`
	UseImgLabel           bool   `json:"use_img_label"`
	ImgLabelUriHans       string `json:"img_label_uri_hans"`
	ImgLabelUriHant       string `json:"img_label_uri_hant"`
	ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
	ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
}

type Vip struct {
	Type               int      `json:"type"`
	Status             int      `json:"status"`
	DueDate            int64    `json:"due_date"`
	VipPayType         int      `json:"vip_pay_type"`
	ThemeType          int      `json:"theme_type"`
	Label              VipLabel `json:"label"`
	AvatarSubscript    int      `json:"avatar_subscript"`
	NicknameColor      string   `json:"nickname_color"`
	Role               int      `json:"role"`
	AvatarSubscriptUrl string   `json:"avatar_subscript_url"`
	TvVipStatus        int      `json:"tv_vip_status"`
	TvVipPayType       int      `json:"tv_vip_pay_type"`
	TvDueDate          int      `json:"tv_due_date"`
}

type Profile struct {
	Mid            uint   `json:"mid"`
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	Face           string `json:"face"`
	Sign           string `json:"sign"`
	Rank           int    `json:"rank"`
	Level          int    `json:"level"`
	Jointime       int    `json:"jointime"`
	Moral          int    `json:"moral"`
	Silence        int    `json:"silence"`
	EmailStatus    int    `json:"email_status"`
	TelStatus      int    `json:"tel_status"`
	Identification int    `json:"identification"`
	Vip            Vip    `json:"vip"`
	Pendant        struct {
		PID               int    `json:"pid"`
		Name              string `json:"name"`
		Image             string `json:"image"`
		Expire            int    `json:"expire"`
		ImageEnhance      string `json:"image_enhance"`
		ImageEnhanceFrame string `json:"image_enhance_frame"`
	} `json:"pendant"`
	Nameplate struct {
		NID        int    `json:"nid"`
		Name       string `json:"name"`
		Image      string `json:"image"`
		ImageSmall string `json:"image_small"`
		Level      string `json:"level"`
		Condition  string `json:"condition"`
	} `json:"nameplate"`
	Official struct {
		Role  int    `json:"role"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
		Type  int    `json:"type"`
	} `json:"official"`
	Birthday      int64 `json:"birthday"`
	IsTourist     int   `json:"is_tourist"`
	IsFakeAccount int   `json:"is_fake_account"`
	PinPrompting  int   `json:"pin_prompting"`
	IsDeleted     int   `json:"is_deleted"`
	InRegAudit    int   `json:"in_reg_audit"`
	IsRipUser     bool  `json:"is_rip_user"`
	Profession    struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		ShowName        string `json:"show_name"`
		IsShow          int    `json:"is_show"`
		CategoryOne     string `json:"category_one"`
		Realname        string `json:"realname"`
		Title           string `json:"title"`
		Department      string `json:"department"`
		CertificateNo   string `json:"certificate_no"`
		CertificateShow bool   `json:"certificate_show"`
	} `json:"profession"`
	FaceNft        int `json:"face_nft"`
	FaceNftNew     int `json:"face_nft_new"`
	IsSeniorMember int `json:"is_senior_member"`
	Honours        struct {
		Mid    int `json:"mid"`
		Colour struct {
			Dark   string `json:"dark"`
			Normal string `json:"normal"`
		} `json:"colour"`
		Tags interface{} `json:"tags"`
	} `json:"honours"`
	DigitalID   string `json:"digital_id"`
	DigitalType int    `json:"digital_type"`
	Attestation struct {
		Type       int `json:"type"`
		CommonInfo struct {
			Title       string `json:"title"`
			Prefix      string `json:"prefix"`
			PrefixTitle string `json:"prefix_title"`
		} `json:"common_info"`
		SpliceInfo struct {
			Title string `json:"title"`
		} `json:"splice_info"`
		Icon string `json:"icon"`
		Desc string `json:"desc"`
	} `json:"attestation"`
	ExpertInfo struct {
		Title string `json:"title"`
		State int    `json:"state"`
		Type  int    `json:"type"`
		Desc  string `json:"desc"`
	} `json:"expert_info"`
}

type LevelExp struct {
	CurrentLevel int   `json:"current_level"`
	CurrentMin   int   `json:"current_min"`
	CurrentExp   int   `json:"current_exp"`
	NextExp      int   `json:"next_exp"`
	LevelUp      int64 `json:"level_up"`
}

type ProfileData struct {
	Profile   Profile  `json:"profile"`
	LevelExp  LevelExp `json:"level_exp"`
	Coins     float64  `json:"coins"`
	Following int      `json:"following"`
	Follower  int      `json:"follower"`
}
