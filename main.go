package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sfreiberg/gotwilio"
)

// TwiML struct represents the structure of the TwiML response
type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say     string   `xml:",omitempty"`
	Play    string   `xml:",omitempty"`
}

// content struct represents the JSON content from the request
type content struct {
	City string `json:"city" binding:"required"`
}

func main() {
	engine := gin.Default()

	// Set your Twilio credentials
	accountSid := "YOUR_TWILIO_ACCOUNT_SID"
	authToken := "YOUR_TWILIO_AUTH_TOKEN"
	fromPhoneNumber := "YOUR_TWILIO_PHONE_NUMBER"

	twilioClient := gotwilio.NewTwilioClient(accountSid, authToken)

	engine.POST("/voice", func(ctx *gin.Context) {
		// Check if JSON content is present in the request
		var content content
		if err := ctx.ShouldBindJSON(&content); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// Create TwiML response based on the content
		var twiml TwiML
		if content.City != "" {
			twiml = TwiML{Say: "Never gonna give you up " + content.City}
		} else {
			twiml = TwiML{Play: "https://demo.twilio.com/docs/classic.mp3"}
		}

		// Marshal TwiML response to XML
		x, err := xml.MarshalIndent(twiml, "", "  ")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Make a Twilio Programmable Voice call
		toPhoneNumber := "PHONE_NUMBER_TO_CALL" // Replace with the phone number you want to call
		_, exception, err := twilioClient.CallWithUrlCallbacks(fromPhoneNumber, toPhoneNumber, "http://your-server-url/voice", &gotwilio.CallbackParameters{})
		if err != nil {
			fmt.Println("Error making Twilio Voice call:", err)
			ctx.String(http.StatusInternalServerError, "Error making Twilio Voice call")
			return
		}
		if exception != nil {
			fmt.Println("Exception making Twilio Voice call:", exception)
			ctx.String(http.StatusInternalServerError, "Exception making Twilio Voice call")
			return
		}

		// Set response headers and send the XML response
		ctx.Header("Content-Type", "application/xml")
		ctx.String(http.StatusOK, string(x))
	})

	fmt.Println("Server started running on http://localhost:3000")
	engine.Run(":3000")
}