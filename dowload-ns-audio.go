package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type ShadowTracks struct {
	Data struct {
		ShadowTracks []struct {
			ShadowTrack struct {
				FemaleAudio string `json:"female_audio"`
				MaleAudio   string `json:"male_audio"`
			} `json:"shadow_track"`
		} `json:"shadow_tracks"`
	} `json:"data"`
}

func main() {
	data, err := ioutil.ReadFile("./audio_n5.json")
	if err != nil {
		panic(err)
	}
	var tracks ShadowTracks
	err = json.Unmarshal(data, &tracks)
	if err != nil {
		panic(err)
	}

	url := "https://storage.googleapis.com/nativshark-audio-files/"
	for _, track := range tracks.Data.ShadowTracks {

		out, err := os.Create(track.ShadowTrack.FemaleAudio)
		if err != nil {
			panic(err)
		}
		resp, err := http.Get(url + track.ShadowTrack.FemaleAudio)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}
		out.Close()
		resp.Body.Close()

		out, err = os.Create(track.ShadowTrack.MaleAudio)
		if err != nil {
			panic(err)
		}
		resp, err = http.Get(url + track.ShadowTrack.MaleAudio)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}
		out.Close()
		resp.Body.Close()
	}
}
