package rbm

const (
	MEDIA_HEIGHT_TALL   string = "TALL"
	MEDIA_HEIGHT_MEDIUM string = "MEDIUM"
	MEDIA_HEIGHT_SMALL  string = "SMALL"
)

type RBMMessage interface {
	GetType() string
}

type RBMSuggestion interface {
	GetSuggestionType() string
}

type RBMRichCard interface {
	GetRichCardType() string
}

type RBMSuggestionReply interface {
	GetSuggestionReplyType() string
}

type RBMSuggestionAction interface {
	GetSuggestionActionType() string
}

type ContentMessage struct {
	ContentMessage RBMMessage `json:"contentMessage"`
}

type TextMessage struct {
	Text        string          `json:"text"`
	Suggestions []RBMSuggestion `json:"suggestions,omitempty"`
}

func (m TextMessage) GetType() string {
	return "text"
}

type RichCardMessage struct {
	RichCard RBMRichCard `json:"richCard"`
}

func (m RichCardMessage) GetType() string {
	return "rich card"
}

type StandaloneCardMessage struct {
	StandaloneCard StandaloneCardMessageData `json:"standaloneCard"`
}

func (s StandaloneCardMessage) GetRichCardType() string {
	return "standalone"
}

type StandaloneCardMessageData struct {
	ThumbnailImageAlignment string      `json:"thumbnailImageAlignment"`
	CardOrientation         string      `json:"cardOrientation"`
	CardContent             CardContent `json:"cardContent"`
}

type CarouselCardMessage struct {
	CarouselCard CarouselCardMessageData `json:"carouselCard"`
}

func (s CarouselCardMessage) GetRichCardType() string {
	return "carousel"
}

type CarouselCardMessageData struct {
	CardWidth    string        `json:"cardWidth"`
	CardContents []CardContent `json:"cardContents"`
}

type CardContent struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Media       Media           `json:"media"`
	Suggestions []RBMSuggestion `json:"suggestions,omitempty"`
}

type Media struct {
	Height      string      `json:"height"`
	ContentInfo ContentInfo `json:"contentInfo"`
}

type MediaMessage struct {
	ContentInfo ContentInfo     `json:"contentInfo"`
	Suggestions []RBMSuggestion `json:"suggestions,omitempty"`
}

func (m MediaMessage) GetType() string {
	return "media"
}

type ContentInfo struct {
	FileUrl      string `json:"fileUrl"`
	ForecRefresh string `json:"forceRefresh"`
}

type ReplySuggestion struct {
	Reply RBMSuggestionReply `json:"reply"`
}

func (s ReplySuggestion) GetSuggestionType() string {
	return "reply"
}

type SuggestedReply struct {
	Text         string `json:"text"`
	PostbackData string `json:"postbackData"`
}

func (r SuggestedReply) GetSuggestionReplyType() string {
	return "reply"
}

type ActionSuggestion struct {
	Action RBMSuggestionAction `json:"action"`
}

func (s ActionSuggestion) GetSuggestionType() string {
	return "action"
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

func (a ActionSuggestionData) GetSuggestionActionType() string {
	return "action"
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
