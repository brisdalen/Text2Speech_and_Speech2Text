// API-key	Rlm4cmv2wK39K_MvAcOL7bMaokhXdbnF7lBn94nbRUZv
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
			IAMApiKey: "Rlm4cmv2wK39K_MvAcOL7bMaokhXdbnF7lBn94nbRUZv",
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