package audiotools

import (
	"net/http"
)

// HeaderFormat is looking at the header of a file and trying to guess the format.
// In most cases you can pass less than 512 bytes.
// Note that this function won't tell the difference between a video and an audio mp4 container.
// See https://golang.org/pkg/net/http/#DetectContentType for details.
func HeaderFormat(header []byte) (Format, error) {
	contentType := http.DetectContentType(header)

	switch contentType {
	case "audio/wave":
		return Wav, nil
	case "audio/aiff":
		return Aiff, nil
	case "audio/mpeg":
		return Mp3, nil
	case "video/avi":
		return Format(contentType), nil
	case "video/webm":
		return Format(contentType), nil
	case "video/mp4":
		return Format(contentType), nil
	}

	return Unknown, nil
}
