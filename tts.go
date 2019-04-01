// This is the text to speech, creating a specified audio-file
// API-key	Use your own key
// URL 		https://gateway-lon.watsonplatform.net/text-to-speech/api

package main

import (
	"bytes"
	"github.com/watson-developer-cloud/go-sdk/texttospeechv1"
	"github.com/IBM/go-sdk-core/core"
	"os"
)


func main() {

	textToSpeech, textToSpeechErr := texttospeechv1.
		NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
			URL: "https://gateway-lon.watsonplatform.net/text-to-speech/api",
			IAMApiKey: "/* Use your own key!*/",
		})
	if textToSpeechErr != nil {
		panic(textToSpeechErr)
	}

	response, responseErr := textToSpeech.Synthesize(
		&texttospeechv1.SynthesizeOptions{
			Text: 		core.StringPtr("Hello, and welcome to my guide. Oh. Nevermind."),
			Accept: 	core.StringPtr(texttospeechv1.SynthesizeOptions_Accept_AudioWav),
			Voice:
				core.StringPtr(texttospeechv1.SynthesizeOptions_Voice_EnUsAllisonvoice),
		},
	)
	if responseErr != nil {
		panic(responseErr)
	}

	result := textToSpeech.GetSynthesizeResult(response)
	if result != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(result)
		file, _ := os.Create("hello.wav")
		file.Write(buff.Bytes())
		file.Close()
	}
}
