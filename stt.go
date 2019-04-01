// This is speech to text
// API-key Use your own key
// URL	   https://gateway-lon.watsonplatform.net/speech-to-text/api
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
	"github.com/IBM/go-sdk-core/core"
)

const key = "ykQgXaO-15Np0ldiS4zSjBs2GDVeOpBXhbRfl0LkoYlh"
const url = "https://gateway-lon.watsonplatform.net/speech-to-text/api"

func main() {

	speechToText, speechToTextErr :=
		speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			IAMApiKey: key,
			URL:       url,
		})
	if speechToTextErr != nil {
		panic(speechToTextErr)
	}

	files := [1]string{"hello.wav"}
	for _, file := range files {
		var audioFile io.ReadCloser
		var audioFileErr error
		audioFile, audioFileErr = os.Open("./" + file)
		if audioFileErr != nil {
			panic(audioFileErr)
		}
		response, responseErr := speechToText.Recognize(
			&speechtotextv1.RecognizeOptions{
				Audio:						&audioFile,
				ContentType:
					core.StringPtr(speechtotextv1.RecognizeOptions_ContentType_AudioWav),
				Timestamps:					core.BoolPtr(true),
				WordAlternativesThreshold:	core.Float32Ptr(0.9),
				Keywords: 					[]string{"hello", "my", "oh"},
				KeywordsThreshold:			core.Float32Ptr(0.5),

			},
		)
		if responseErr != nil {
			panic(responseErr)
		}

		result := speechToText.GetRecognizeResult(response)
		b, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(b))
	}

}
