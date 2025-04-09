// Package rule provides a collection of validation rules for various data types.
// This file contains file-related validation rules for size, type, extension, and MIME type.
package rule

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"path/filepath"
	"strings"
)

// File validation errors
var (

	// ErrFileType is returned when a file's content type is not in the allowed list.
	// The file type is determined by examining the file's header bytes.
	ErrFileType = errors.New("file type is not allowed")

	// ErrFileExtension is returned when a file's extension is not in the allowed list.
	// The extension is extracted from the filename and compared case-insensitively.
	ErrFileExtension = errors.New("file extension is not allowed")

	// ErrFileMimeType is returned when a file's MIME type is not in the allowed list.
	// The MIME type is determined using the file's extension and the mime package.
	ErrFileMimeType = errors.New("file mime type is not allowed")
)

const (
	ErrFileSizeFormat = "file size is not between %v and %v"
)

// FileSizeRule validates that a file's size falls within a specified range.
// The file must be at least the minimum size and no larger than the maximum size.
//
// Example:
//
//	rule := FileSize(1024, 10485760).Err("File must be between 1KB and 10MB")
//	err := rule.Validate(fileReader)  // returns nil if file size is within range
type FileSizeRule struct {
	min int64
	max int64
	e   error
}

// FileSize creates a new file size validation rule.
// The rule ensures that a file's size falls between the specified minimum and maximum values.
//
// Example:
//
//	rule := FileSize(1024, 10485760)  // between 1KB and 10MB
//	rule := FileSize(0, 5242880)      // up to 5MB
//	rule := FileSize(1048576, 0)      // at least 1MB (0 max means no upper limit)
func FileSize(min, max int64) *FileSizeRule {
	return &FileSizeRule{
		min: min,
		max: max,
		e:   fmt.Errorf(ErrFileSizeFormat, min, max),
	}
}

// Validate checks if the given file's size falls within the specified range.
// The file is read to determine its size, which is then compared against the min and max values.
//
// Example:
//
//	file, _ := os.Open("document.pdf")
//	defer file.Close()
//	rule := FileSize(1024, 10485760)
//	err := rule.Validate(file)  // returns nil if file size is between 1KB and 10MB
func (r *FileSizeRule) Validate(file io.Reader) error {
	// Get file size by reading the entire file
	size, err := io.Copy(io.Discard, file)
	if err != nil {
		return err
	}

	// Check if file size is within the specified range
	if size < r.min || (r.max > 0 && size > r.max) {
		return r.e
	}

	return nil
}

// Err sets a custom error message for file size validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := FileSize(1024, 10485760).Err("Uploaded file must be between 1KB and 10MB")
func (r *FileSizeRule) Errf(format string, args ...any) *FileSizeRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// FileTypeRule validates that a file's content type matches one of the allowed types.
// The file type is determined by examining the file's header bytes.
//
// Example:
//
//	rule := FileType("PDF", "DOCX").Err("Only PDF and DOCX files are allowed")
//	err := rule.Validate(fileReader)  // returns nil if file type is allowed
type FileTypeRule struct {
	allowedTypes []string
	e            error
}

// FileType creates a new file type validation rule.
// The rule ensures that a file's content type matches one of the allowed types.
//
// Example:
//
//	rule := FileType("PDF", "DOCX")  // allow PDF and DOCX files
//	rule := FileType("JPEG", "PNG")  // allow JPEG and PNG images
func FileType(allowedTypes ...string) *FileTypeRule {
	return &FileTypeRule{
		allowedTypes: allowedTypes,
		e:            ErrFileType,
	}
}

// Validate checks if the given file's content type matches one of the allowed types.
// The file's header bytes are examined to determine its type.
//
// Example:
//
//	file, _ := os.Open("document.pdf")
//	defer file.Close()
//	rule := FileType("PDF", "DOCX")
//	err := rule.Validate(file)  // returns nil if file is a PDF or DOCX
func (r *FileTypeRule) Validate(file io.Reader) error {
	// Read file header to determine file type
	header := make([]byte, 512)
	_, err := file.Read(header)
	if err != nil && err != io.EOF {
		return err
	}

	// Check if file type is in the allowed list
	for _, allowedType := range r.allowedTypes {
		if strings.Contains(string(header), allowedType) {
			return nil
		}
	}

	return r.e
}

// Errf sets a custom error message for file type validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := FileType("PDF", "DOCX").Errf("Please upload only PDF or DOCX files")
func (r *FileTypeRule) Errf(format string, args ...any) *FileTypeRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// FileExtensionRule validates that a file's extension is in the allowed list.
// The extension is extracted from the filename and compared case-insensitively.
//
// Example:
//
//	rule := FileExtension("pdf", "docx").Err("Only PDF and DOCX files are allowed")
//	err := rule.Validate("document.pdf")  // returns nil
//	err = rule.Validate("image.jpg")     // returns error
type FileExtensionRule struct {
	allowedExtensions []string
	e                 error
}

// FileExtension creates a new file extension validation rule.
// The rule ensures that a file's extension is in the allowed list.
//
// Example:
//
//	rule := FileExtension("pdf", "docx")  // allow PDF and DOCX files
//	rule := FileExtension("jpg", "png")   // allow JPG and PNG images
func FileExtension(allowedExtensions ...string) *FileExtensionRule {
	return &FileExtensionRule{
		allowedExtensions: allowedExtensions,
		e:                 ErrFileExtension,
	}
}

// Validate checks if the given filename's extension is in the allowed list.
// The extension is extracted from the filename and compared case-insensitively.
//
// Example:
//
//	rule := FileExtension("pdf", "docx")
//	err := rule.Validate("document.pdf")  // returns nil
//	err = rule.Validate("image.jpg")     // returns error
//	err = rule.Validate("noextension")   // returns error
func (r *FileExtensionRule) Validate(filename string) error {
	// Get file extension
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" {
		return r.e
	}

	// Remove the leading dot from the extension
	ext = ext[1:]

	// Check if extension is in the allowed list
	for _, allowedExt := range r.allowedExtensions {
		if ext == strings.ToLower(allowedExt) {
			return nil
		}
	}

	return r.e
}

// Errf sets a custom error message for file extension validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := FileExtension("pdf", "docx").Errf("Please upload only PDF or DOCX files")
func (r *FileExtensionRule) Errf(format string, args ...any) *FileExtensionRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}

// FileMimeTypeRule validates that a file's MIME type is in the allowed list.
// The MIME type is determined using the file's extension and the mime package.
//
// Example:
//
//	rule := FileMimeType("application/pdf", "application/vnd.openxmlformats-officedocument.wordprocessingml.document").Err("Only PDF and DOCX files are allowed")
//	err := rule.Validate(fileReader)  // returns nil if MIME type is allowed
type FileMimeTypeRule struct {
	allowedMimeTypes []string
	e                error
}

// FileMimeType creates a new file MIME type validation rule.
// The rule ensures that a file's MIME type is in the allowed list.
//
// Example:
//
//	rule := FileMimeType("application/pdf", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")  // allow PDF and DOCX files
//	rule := FileMimeType("image/jpeg", "image/png")  // allow JPEG and PNG images
func FileMimeType(allowedMimeTypes ...string) *FileMimeTypeRule {
	return &FileMimeTypeRule{
		allowedMimeTypes: allowedMimeTypes,
		e:                ErrFileMimeType,
	}
}

// Validate checks if the given file's MIME type is in the allowed list.
// The MIME type is determined using the file's extension and the mime package.
//
// Example:
//
//	file, _ := os.Open("document.pdf")
//	defer file.Close()
//	rule := FileMimeType("application/pdf", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
//	err := rule.Validate(file)  // returns nil if file is a PDF or DOCX
func (r *FileMimeTypeRule) Validate(file io.Reader) error {
	// Read file header to determine MIME type
	header := make([]byte, 512)
	_, err := file.Read(header)
	if err != nil && err != io.EOF {
		return err
	}

	// Get MIME type
	mimeType := mime.TypeByExtension(filepath.Ext(string(header)))
	if mimeType == "" {
		return r.e
	}

	// Check if MIME type is in the allowed list
	for _, allowedMimeType := range r.allowedMimeTypes {
		if mimeType == allowedMimeType {
			return nil
		}
	}

	return r.e
}

// Errf sets a custom error message for file MIME type validation failures.
// This allows for context-specific error messages.
//
// Example:
//
//	rule := FileMimeType("application/pdf", "application/vnd.openxmlformats-officedocument.wordprocessingml.document").Errf("Please upload only PDF or DOCX files")
func (r *FileMimeTypeRule) Errf(format string, args ...any) *FileMimeTypeRule {
	if format != "" {
		r.e = fmt.Errorf(format, args...)
	}
	return r
}
