package rbm

import (
	"time"

	"github.com/enescakir/emoji"
)

func GetReplySuggestion(text string, postbackData string) ReplySuggestion {

	text = emoji.Parse(text)

	return ReplySuggestion{
		Reply: SuggestedReply{
			Text:         text,
			PostbackData: postbackData,
		},
	}
}

func GetOpenUrlSuggestion(text string, postbackData string, url string) ActionSuggestion {

	text = emoji.Parse(text)

	return ActionSuggestion{
		Action: ActionSuggestionData{
			Text:         text,
			PostbackData: postbackData,
			OpenUrlAction: &OpenUrlAction{
				Url: url,
			},
		},
	}
}

func GetDialSuggestion(text string, postbackData string, phoneNumber string) ActionSuggestion {

	text = emoji.Parse(text)

	return ActionSuggestion{
		Action: ActionSuggestionData{
			Text:         text,
			PostbackData: postbackData,
			DialAction: &DialAction{
				PhoneNumber: phoneNumber,
			},
		},
	}
}

func GetViewLocationSuggestion(text string, postbackData string, lat string, long string, locationLabel string) ActionSuggestion {

	text = emoji.Parse(text)

	return ActionSuggestion{
		Action: ActionSuggestionData{
			Text:         text,
			PostbackData: postbackData,
			ViewLocatinoAction: &ViewLocatinoAction{
				LatLong: LatLong{
					Latitude:  lat,
					Longitude: long,
				},
				Label: locationLabel,
			},
		},
	}
}

func GetShareLocationoSuggestion(text string, postbackData string) ActionSuggestion {

	text = emoji.Parse(text)

	return ActionSuggestion{
		Action: ActionSuggestionData{
			Text:                text,
			PostbackData:        postbackData,
			ShareLocationAction: &ShareLocationAction{},
		},
	}
}

func GetCreateCalendarEventSuggestion(text string, postbackData string, evnetTitle string, eventDescription string, startTime time.Time, endTime time.Time) ActionSuggestion {

	text = emoji.Parse(text)

	return ActionSuggestion{
		Action: ActionSuggestionData{
			Text:         text,
			PostbackData: postbackData,
			CreateCalendarEventAction: &CreateCalendarEventAction{
				StartTime:   startTime.Format("2006-01-02T15:04:05Z"),
				EndTime:     endTime.Format("2006-01-02T15:04:05Z"),
				Title:       evnetTitle,
				Description: eventDescription,
			},
		},
	}
}

func GetTextMessage(text string, suggestions []any) []byte {
	var s []any
	if len(suggestions) > 0 {
		s = suggestions
	}

	text = emoji.Parse(text)

	textMessage := TextMessage{
		Text:        text,
		Suggestions: s,
	}
	contentMessage := ContentMessage{
		ContentMessage: textMessage,
	}

	return StructToJson(contentMessage)
}

func GetTextMessageStruct(text string, suggestions []any) ContentMessage {
	var s []any
	if len(suggestions) > 0 {
		s = suggestions
	}

	text = emoji.Parse(text)

	textMessage := TextMessage{
		Text:        text,
		Suggestions: s,
	}
	contentMessage := ContentMessage{
		ContentMessage: textMessage,
	}

	return contentMessage
}

func GetStandaloneCardMessage(title string, description string, imageUrl string, mediaSize string, suggestions []any) []byte {
	var s []any
	if len(suggestions) > 0 {
		s = suggestions
	}

	title = emoji.Parse(title)
	description = emoji.Parse(description)

	contentMessage := ContentMessage{
		ContentMessage: RichCardMessage{
			RichCard: StandaloneCardMessage{
				StandaloneCard: StandaloneCardMessageData{
					ThumbnailImageAlignment: "RIGHT",
					CardOrientation:         "VERTICAL",
					CardContent: CardContent{
						Title:       title,
						Description: description,
						Media: Media{
							Height: mediaSize,
							ContentInfo: ContentInfo{
								FileUrl:      imageUrl,
								ForecRefresh: "false",
							},
						},
						Suggestions: s,
					},
				},
			},
		},
	}

	return StructToJson(contentMessage)
}

func GetCarouselCardMessage(cardWidth string, cardContent []CardContent) []byte {
	contentMessage := ContentMessage{
		ContentMessage: RichCardMessage{
			RichCard: CarouselCardMessage{
				CarouselCard: CarouselCardMessageData{
					CardWidth:    cardWidth, // SMALL, MEDIUM
					CardContents: cardContent,
				},
			},
		},
	}

	return StructToJson(contentMessage)
}

func GetCardContent(title string, description string, imageUrl string, mediaSize string, suggestions []any) CardContent {
	var s []any
	if len(suggestions) > 0 {
		s = suggestions
	}

	title = emoji.Parse(title)
	description = emoji.Parse(description)

	content := CardContent{
		Title:       title,
		Description: description,
		Media: Media{
			Height: mediaSize,
			ContentInfo: ContentInfo{
				FileUrl:      imageUrl,
				ForecRefresh: "false",
			},
		},
		Suggestions: s,
	}

	return content
}

func GetMediaMessage(url string, suggestions []any) []byte {
	contentMessage := ContentMessage{
		ContentMessage: MediaMessage{
			ContentInfo: ContentInfo{
				FileUrl:      url,
				ForecRefresh: "false",
			},
			Suggestions: suggestions,
		},
	}

	return StructToJson(contentMessage)
}
