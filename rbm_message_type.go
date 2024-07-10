package rbm

const (
	MEDIA_HEIGHT_TALL   string = "TALL"
	MEDIA_HEIGHT_MEDIUM string = "MEDIUM"
	MEDIA_HEIGHT_SMALL  string = "SMALL"
)

type ContentMessage struct {
	ContentMessage any `json:"contentMessage"`
}

type TextMessage struct {
	Text        string `json:"text"`
	Suggestions []any  `json:"suggestions,omitempty"`
}

type RichCardMessage struct {
	RichCard any `json:"richCard"`
}

type StandaloneCardMessage struct {
	StandaloneCard StandaloneCardMessageData `json:"standaloneCard"`
}

type StandaloneCardMessageData struct {
	ThumbnailImageAlignment string      `json:"thumbnailImageAlignment"`
	CardOrientation         string      `json:"cardOrientation"`
	CardContent             CardContent `json:"cardContent"`
}

type CarouselCardMessage struct {
	CarouselCard CarouselCardMessageData `json:"carouselCard"`
}

type CarouselCardMessageData struct {
	CardWidth    string        `json:"cardWidth"`
	CardContents []CardContent `json:"cardContents"`
}

type CardContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Media       Media  `json:"media"`
	Suggestions []any  `json:"suggestions,omitempty"`
}

type Media struct {
	Height      string      `json:"height"`
	ContentInfo ContentInfo `json:"contentInfo"`
}

type MediaMessage struct {
	ContentInfo ContentInfo `json:"contentInfo"`
	Suggestions []any       `json:"suggestions,omitempty"`
}

type ContentInfo struct {
	FileUrl      string `json:"fileUrl"`
	ForecRefresh string `json:"forceRefresh"`
}

type ReplySuggestion struct {
	Reply any `json:"reply"`
}

type SuggestedReply struct {
	Text         string `json:"text"`
	PostbackData string `json:"postbackData"`
}

type ActionSuggestion struct {
	Action any `json:"action"`
}

type ActionSuggestionData struct {
	Text         string `json:"text"`
	PostbackData string `json:"postbackData"`
	FallbackUrl  string `json:"fallbackUrl,omitempty"`

	OpenUrlAction             *OpenUrlAction             `json:"openUrlAction,omitempty"`
	DialAction                *DialAction                `json:"dialAction,omitempty"`
	ViewLocatinoAction        *ViewLocatinoAction        `json:"viewLocationAction,omitempty"`
	ShareLocationAction       *ShareLocationAction       `json:"shareLocationAction,omitempty"`
	CreateCalendarEventAction *CreateCalendarEventAction `json:"createCalendarEventAction,omitempty"`
}

type OpenUrlAction struct {
	Url string `json:"url"`
}

type DialAction struct {
	PhoneNumber string `json:"phoneNumber"`
}

type ViewLocatinoAction struct {
	LatLong LatLong `json:"latLong"`
	Label   string  `json:"label"`
}

type LatLong struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type ShareLocationAction struct{}

type CreateCalendarEventAction struct {
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
