package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// AudioMetadata encapsulates audio
// metadata information.
type AudioMetadata struct {
	Duration float64
	Bitrate  float64
}

func getAudioMetadata(format string) (string, error) {
	cmd := exec.Command("./meta.py", "/tmp/test.mp3", format)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func main() {
	metadata, err := getAudioMetadata("string")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(metadata)

	var audio AudioMetadata

	metadata, err = getAudioMetadata("json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(metadata), &audio)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(metadata)
	fmt.Printf("decoded into duration %f and bitrate %.0f\n", audio.Duration, audio.Bitrate)
}
