package audiotools_test

import (
	"reflect"
	"testing"

	"github.com/go-audio/audiotools"
)

func TestFileFormat(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    audiotools.Format
		wantErr bool
	}{
		{name: "wav", path: "../wav/fixtures/bass.wav", want: audiotools.Wav},
		{name: "aiff", path: "../aiff/fixtures/kick.aif", want: audiotools.Aiff},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := audiotools.FileFormat(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
