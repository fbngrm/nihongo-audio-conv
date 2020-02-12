package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"github.com/dhowden/tag"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

// [START tts_synthesize_text]

// SynthesizeText synthesizes plain text and saves the output to outputFile.
func SynthesizeText(w io.Writer, text, outputFile string) error {
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return err
	}

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		// Note: the voice can also be specified by name.
		// Names of voices can be retrieved with client.ListVoices().
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
			VolumeGainDb:  9.0,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFile, resp.AudioContent, 0644)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Audio content written to file: %v\n", outputFile)
	return nil
}

func main() {
	audiodir := flag.String("audio-dir", "", "The directory containing audio files")
	flag.Parse()

	files, err := ioutil.ReadDir(*audiodir)
	if err != nil {
		log.Fatal(err)
	}

	a := `{
		"__type__": "Note",
		"fields": [
			"%s",
			"%s",
			"%s",
			"%s"
		]
		},
		`
	card := ""
	for _, file := range files {
		filename := path.Join(*audiodir, file.Name())
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		m, err := tag.ReadFrom(f)
		if err != nil {
			log.Fatal(err)
		}
		parts := strings.Split(m.Title(), "-")
		words := make([]string, len(parts))
		for i, s := range parts {
			s = strings.TrimSpace(s)
			parts := strings.Split(s, "_")
			s = strings.TrimSpace(parts[len(parts)-1])
			words[i] = s
		}
		card += fmt.Sprintf(a, words[0], words[1], words[2], words[3])
		// text := strings.TrimSpace(parts[len(parts)-1])
		// parts = strings.Split(text, "_")
		// text = strings.TrimSpace(parts[len(parts)-1])

		// basename := strings.TrimSuffix(filename, filepath.Ext(filename))
		// filenameEn := basename + " - en.mp3"
		// if text != "" {
		// 	err := SynthesizeText(os.Stdout, text, filenameEn)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// }
		// filenameJp := basename + " - jp.mp3"
		// err = os.Rename(filename, filenameJp)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
	// write the whole body at once
	err = ioutil.WriteFile("output.txt", []byte(card), 0644)
	if err != nil {
		panic(err)
	}
}
