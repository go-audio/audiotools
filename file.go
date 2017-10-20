package audiotools

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-audio/aiff"
	"github.com/go-audio/wav"
)

var (
	// ErrInvalidPath indicates that the path is... wait for it... invalid!
	ErrInvalidPath = errors.New("invalid path")
)

// Format is the audio format canonical name
type Format string

var (
	// Unknown means that this library couldn't detect the format type
	Unknown Format = "unknown"
	// Wav is the Waveform Audio File Format (WAVE, or more commonly known
	// as WAV due to its filename extension)
	Wav Format = "wav"
	// Aiff is the Audio Interchange File Format
	Aiff Format = "aiff"
)

// FileFormat returns the known format of the passed path.
func FileFormat(path string) (Format, error) {
	if !fileExists(path) {
		return "", ErrInvalidPath
	}
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var triedWav bool
	var triedAif bool

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".wav", ".wave":
		triedWav = true
		d := wav.NewDecoder(f)
		if d.IsValidFile() {
			return Wav, nil
		}
	case ".aif", ".aiff":
		triedAif = true
		d := aiff.NewDecoder(f)
		if d.IsValidFile() {
			return Aiff, nil
		}
	}
	// extension doesn't match, let's try again
	f.Seek(0, 0)
	if !triedWav {
		wd := wav.NewDecoder(f)
		if wd.IsValidFile() {
			return Wav, nil
		}
		f.Seek(0, 0)
	}
	if !triedAif {
		ad := aiff.NewDecoder(f)
		if ad.IsValidFile() {
			return Aiff, nil
		}
	}
	return Unknown, nil
}

// helper checking if a file exists
func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
