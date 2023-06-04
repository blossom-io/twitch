package ffmpeg

import (
	"fmt"
	"io"
	"os/exec"
)

//go:generate mockery --name FFMpeger
type FFMpeger interface {
	Screenshot(input string) (img []byte, err error)
}

type FFMpeg struct{}

func New() FFMpeger {
	return &FFMpeg{}
}

func (f *FFMpeg) Screenshot(input string) (img []byte, err error) {
	if input == "" {
		return nil, fmt.Errorf("input is empty")
	}

	cmd := exec.Command("ffmpeg", "-nostats", "-loglevel", "0", "-i", input, "-f", "image2", "-update", "1", "-vframes", "1", "-q:v", "1", "pipe:1")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe: ", err)
		return
	}
	defer stdout.Close()

	if err = cmd.Start(); err != nil {
		fmt.Println("Error starting command: ", err)
		return
	}

	img, err = io.ReadAll(stdout)
	if err != nil {
		fmt.Println("Error reading output: ", err)
		return
	}

	if err = cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command: ", err)
		return
	}

	return img, nil
}
