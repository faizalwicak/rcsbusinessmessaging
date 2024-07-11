# RCS BUSINESS MESSAGING

Use this library to help you send message to rbm platform with golang

## Requirement

This library use golang 1.22.3 or above

## Install

Use `go get` to install this library

```bash
go get github.com/faizalwicak/rcsbusinessmessaging
```

## Getting Start 

To start using this library you can import this library and call RBMHelper 

```go
package main

import (
	"log"
	"path/filepath"

	rbm "github.com/faizalwicak/rcsbusinessmessaging"
)

func main() {
	filename := filepath.Join("temp", "path to credential")
	rbmHelper, err := rbm.GetRBMHelperInstanceFromFile("<agent_id>", filename)
	if err != nil {
		log.Fatalf("fail to get rbm instance %v", err)
	}

	// send single message
	// return rbm id, status message, error
	message := rbm.GetTextMessage("hello", []rbm.RBMSuggestion{})
	rbmHelper.SendMessage("<msisdn>", message)

	// send multiple message
	// return rbm id, status message, error
	messages := [][]byte{}
	message1 := rbm.GetTextMessage("hello1", []rbm.RBMSuggestion{})
	messages = append(messages, message1)
	message2 := rbm.GetTextMessage("hello2", []rbm.RBMSuggestion{})
	messages = append(messages, message2)
	rbmHelper.SendMultipleMessage("<msisdn>", messages)

	// send event
	// return error
	rbmHelper.SendEvent("<msisdn>", "<event>", "<message_id>")

	// check capability
	// return capability code <int>, error
	rbmHelper.CapabilityCheck("<msisdn>")
}
```

## Message

### Message Type

```go

// #################################
// generate text message
// #################################
suggestions := []rbm.RBMSuggestion{}
message = rbm.GetTextMessage("<text>", suggestions)

// #################################
// generate media text
// #################################
rbm.GetMediaMessage("<image / video>", suggestions)

// #################################
// generate standalone card message
// image size: rbm.MEDIA_HEIGHT_MEDIUM, rbm.MEDIA_HEIGHT_TALL
// #################################
message = rbm.GetStandaloneCardMessage("<title>", "<description>", "<image>", "<image size>", suggestions)

// #################################
// generate carousel card message
// minimal sum card number: 3 card
// image size: rbm.MEDIA_HEIGHT_MEDIUM, rbm.MEDIA_HEIGHT_TALL
// card width: rbm.MEDIA_HEIGHT_SMALL, rbm.MEDIA_HEIGHT_MEDIUM
// #################################
cardContents := []rbm.CardContent{}
cardContents = append(cardContents, rbm.GetCardContent("<title>", "<description>", "<image>", "<image size>", suggestions))
cardContents = append(cardContents, rbm.GetCardContent("<title>", "<description>", "<image>", "<image size>", suggestions))
cardContents = append(cardContents, rbm.GetCardContent("<title>", "<description>", "<image>", "<image size>", suggestions))

message = rbm.GetCarouselCardMessage("<card width>", cardContents)
```

### Message Suggestion

```go
// reply text suggestion
rbm.GetReplySuggestion("<text>", "<postback data>")

// open url suggestion
rbm.GetOpenUrlSuggestion("<text>", "<postback data>", "<url>")

// dial number suggestion
rbm.GetDialSuggestion("<text>", "<postback data>", "<phone number>")

// create calender event suggestion
rbm.GetCreateCalendarEventSuggestion("<text>", "<postback data>", "<event title>", "<event description>", "<start time>", "<end time>")

// share location suggestion
rbm.GetShareLocationoSuggestion("<text>", "<postback data>")

// view location suggestion
rbm.GetViewLocationSuggestion("<text>", "<postback data>", "<latitude>", "<longitude>", "<location lable>")
```

## License

Released under the [MIT License](https://github.com/faizalwicak/rcsbusinessmessaging/blob/main/LICENSE)