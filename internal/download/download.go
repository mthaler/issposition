package download

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Download(url string) (*bytes.Buffer, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	b := new(bytes.Buffer)
	_, err = io.Copy(b, resp.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}
