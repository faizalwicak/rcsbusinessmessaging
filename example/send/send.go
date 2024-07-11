package main

import (
	"log"
	"path/filepath"

	rbm "github.com/faizalwicak/rcsbusinessmessaging"
)

func main() {
	filename := filepath.Join("temp", "path to credential")
	rbmHelper, err := rbm.GetRBMHelperInstanceFromFile("<agent id>", filename)
	if err != nil {
		log.Fatalf("fail to get rbm instance %v", err)
	}

	// send single message
	// return rbm id, status message, error
	message := rbm.GetTextMessage("hello", []any{})
	rbmHelper.SendMessage("<msisdn>", message)

	// send multiple message
	// return rbm id, status message, error
	messages := [][]byte{}
	message1 := rbm.GetTextMessage("hello1", []any{})
	messages = append(messages, message1)
	message2 := rbm.GetTextMessage("hello2", []any{})
	messages = append(messages, message2)
	rbmHelper.SendMultipleMessage("<msisdn>", messages)

	// send event
	// return error
	rbmHelper.SendEvent("<msisdn>", "<event>", "<message_id>")

	// check capability
	// return capability code <int>, error
	rbmHelper.CapabilityCheck("<msisdn>")

	// #################################
	// generate text message
	// #################################
	suggestions := []any{}
	message = rbm.GetTextMessage("<text>", suggestions)
	log.Println(message)

	// #################################
	// generate media text
	// #################################
	rbm.GetMediaMessage("<image / video>", suggestions)
	log.Println(message)

	// #################################
	// generate standalone card message
	// image size: rbm.MEDIA_HEIGHT_MEDIUM, rbm.MEDIA_HEIGHT_TALL
	// #################################
	message = rbm.GetStandaloneCardMessage("<title>", "<description>", "<image>", "<image size>", suggestions)
	log.Println(message)

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
	log.Println(message)

	// reply text suggestion
	rbm.GetReplySuggestion("<text>", "<postback data>")

	// open url suggestion
	rbm.GetOpenUrlSuggestion("<text>", "<postback data>", "<url>")

	// dial number suggestion
	rbm.GetDialSuggestion("<text>", "<postback data>", "<phone number>")

	// create calender event suggestion
	// rbm.GetCreateCalendarEventSuggestion("<text>", "<postback data>", "<event title>", "<event description>", "<start time>", "<end time>")

	// share location suggestion
	rbm.GetShareLocationoSuggestion("<text>", "<postback data>")

	// view location suggestion
	rbm.GetViewLocationSuggestion("<text>", "<postback data>", "<latitude>", "<longitude>", "<location lable>")

}
