package rule

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestFileSize(t *testing.T) {
	tests := []struct {
		name    string
		min     int64
		max     int64
		content []byte
		wantErr bool
	}{
		{
			name:    "valid size",
			min:     1,
			max:     10,
			content: []byte("hello"),
			wantErr: false,
		},
		{
			name:    "too small",
			min:     10,
			max:     20,
			content: []byte("hello"),
			wantErr: true,
		},
		{
			name:    "too large",
			min:     1,
			max:     3,
			content: []byte("hello"),
			wantErr: true,
		},
		{
			name:    "empty file",
			min:     0,
			max:     10,
			content: []byte{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := FileSize(tt.min, tt.max)
			reader := bytes.NewReader(tt.content)
			err := rule.Validate(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileSize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileType(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name         string
		allowedTypes []string
		content      []byte
		wantErr      bool
	}{
		{
			name:         "valid type",
			allowedTypes: []string{"png"},
			content:      []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, // PNG magic (8 bytes)
			wantErr:      false,
		},
		{
			name:         "invalid type",
			allowedTypes: []string{"jpeg"},
			content:      []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, // PNG header
			wantErr:      true,
		},
		{
			name:         "empty file",
			allowedTypes: []string{"png"},
			content:      []byte{},
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := filepath.Join(tmpDir, "test")
			err := os.WriteFile(tmpFile, tt.content, 0644)
			if err != nil {
				t.Fatal(err)
			}

			file, err := os.Open(tmpFile)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			rule := FileType(tt.allowedTypes...)
			err = rule.Validate(file)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileExtension(t *testing.T) {
	tests := []struct {
		name              string
		allowedExtensions []string
		filename          string
		wantErr           bool
	}{
		{
			name:              "valid extension",
			allowedExtensions: []string{"jpg", "png"},
			filename:          "test.jpg",
			wantErr:           false,
		},
		{
			name:              "invalid extension",
			allowedExtensions: []string{"jpg", "png"},
			filename:          "test.gif",
			wantErr:           true,
		},
		{
			name:              "no extension",
			allowedExtensions: []string{"jpg", "png"},
			filename:          "test",
			wantErr:           true,
		},
		{
			name:              "case insensitive",
			allowedExtensions: []string{"jpg", "png"},
			filename:          "test.JPG",
			wantErr:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := FileExtension(tt.allowedExtensions...)
			err := rule.Validate(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileExtension() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileMimeType(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name             string
		allowedMimeTypes []string
		content          []byte
		extension        string
		wantErr          bool
	}{
		{
			name:             "valid mime type",
			allowedMimeTypes: []string{"image/png"},
			content:          []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, // PNG magic (8 bytes)
			extension:        ".png",
			wantErr:          false,
		},
		{
			name:             "invalid mime type",
			allowedMimeTypes: []string{"image/jpeg"},
			content:          []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, // PNG header
			extension:        ".png",
			wantErr:          true,
		},
		{
			name:             "empty file",
			allowedMimeTypes: []string{"image/png"},
			content:          []byte{},
			extension:        ".png",
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := filepath.Join(tmpDir, "test"+tt.extension)
			err := os.WriteFile(tmpFile, tt.content, 0644)
			if err != nil {
				t.Fatal(err)
			}

			file, err := os.Open(tmpFile)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			rule := FileMimeType(tt.allowedMimeTypes...)
			err = rule.Validate(file)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileMimeType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileSizeErrf(t *testing.T) {
	rule := FileSize(1, 5).Errf("custom size error")
	err := rule.Validate(bytes.NewReader([]byte("hello world")))
	if err == nil || err.Error() != "custom size error" {
		t.Errorf("FileSize().Errf() error = %v, want custom size error", err)
	}
}

func TestFileTypeErrf(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/test"
	os.WriteFile(tmpFile, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, 0644)
	f, _ := os.Open(tmpFile)
	defer f.Close()

	rule := FileType("jpeg").Errf("custom type error")
	err := rule.Validate(f)
	if err == nil || err.Error() != "custom type error" {
		t.Errorf("FileType().Errf() error = %v, want custom type error", err)
	}
}

func TestFileExtensionErrf(t *testing.T) {
	rule := FileExtension("jpg", "png").Errf("custom ext error")
	err := rule.Validate("test.gif")
	if err == nil || err.Error() != "custom ext error" {
		t.Errorf("FileExtension().Errf() error = %v, want custom ext error", err)
	}
}

func TestFileMimeTypeErrf(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/test.png"
	os.WriteFile(tmpFile, []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}, 0644)
	f, _ := os.Open(tmpFile)
	defer f.Close()

	rule := FileMimeType("image/jpeg").Errf("custom mime error")
	err := rule.Validate(f)
	if err == nil || err.Error() != "custom mime error" {
		t.Errorf("FileMimeType().Errf() error = %v, want custom mime error", err)
	}
}
