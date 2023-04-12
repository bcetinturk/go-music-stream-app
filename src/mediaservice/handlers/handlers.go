package handlers

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"regexp"
)

func GetAudio(url_path string) (*os.File, fs.FileInfo, error) {
	re := regexp.MustCompile(`/audio/(\w+)`)
	matches := re.FindStringSubmatch(url_path)

	if len(matches) == 2 {
		id := matches[1]
		fmt.Printf("Audio file %s\n", id)
	} else {
		return nil, nil, &RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("could not match audio name"),
		}
	}

	path_str := fmt.Sprintf("%s.m4a", matches[1])
	file, err := os.Open(path_str)
	if err != nil {
		return nil, nil, &RequestError{
			StatusCode: http.StatusNotFound,
			Err:        errors.New("audio file not found"),
		}
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Here\n")
		return nil, nil, &RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("could not get file information"),
		}
	}

	return file, fileInfo, nil
}
