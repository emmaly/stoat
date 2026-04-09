package stoat

import (
	"encoding/json"
	"fmt"
)

// File represents a file stored on the CDN.
type File struct {
	ID          string      `json:"_id"`
	Tag         string      `json:"tag"`
	Filename    string      `json:"filename"`
	Metadata    RawMetadata `json:"metadata"`
	ContentType string      `json:"content_type"`
	Size        int         `json:"size"`
	Deleted     bool        `json:"deleted,omitempty"`
	Reported    bool        `json:"reported,omitempty"`
	MessageID   *string     `json:"message_id,omitempty"`
	UserID      *string     `json:"user_id,omitempty"`
	ServerID    *string     `json:"server_id,omitempty"`
	ObjectID    *string     `json:"object_id,omitempty"`
}

// Metadata is the interface for file metadata variants.
type Metadata interface {
	metadataMarker()
	MetadataType() string
}

// FileMetadata represents generic file metadata (type="File").
type FileMetadata struct{}

func (FileMetadata) metadataMarker()        {}
func (FileMetadata) MetadataType() string    { return "File" }

// TextMetadata represents text file metadata (type="Text").
type TextMetadata struct{}

func (TextMetadata) metadataMarker()        {}
func (TextMetadata) MetadataType() string    { return "Text" }

// ImageMetadata represents image file metadata (type="Image").
type ImageMetadata struct {
	Width     int   `json:"width"`
	Height    int   `json:"height"`
	Thumbhash []int `json:"thumbhash,omitempty"`
	Animated  bool  `json:"animated,omitempty"`
}

func (ImageMetadata) metadataMarker()        {}
func (ImageMetadata) MetadataType() string    { return "Image" }

// VideoMetadata represents video file metadata (type="Video").
type VideoMetadata struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (VideoMetadata) metadataMarker()        {}
func (VideoMetadata) MetadataType() string    { return "Video" }

// AudioMetadata represents audio file metadata (type="Audio").
type AudioMetadata struct{}

func (AudioMetadata) metadataMarker()        {}
func (AudioMetadata) MetadataType() string    { return "Audio" }

// RawMetadata is a wrapper that handles JSON unmarshalling of the Metadata tagged union.
type RawMetadata struct {
	Value Metadata
}

// MarshalJSON implements json.Marshaler.
func (r RawMetadata) MarshalJSON() ([]byte, error) {
	switch v := r.Value.(type) {
	case *FileMetadata:
		return json.Marshal(struct {
			Type string `json:"type"`
		}{Type: "File"})
	case *TextMetadata:
		return json.Marshal(struct {
			Type string `json:"type"`
		}{Type: "Text"})
	case *ImageMetadata:
		return json.Marshal(struct {
			Type string `json:"type"`
			*ImageMetadata
		}{Type: "Image", ImageMetadata: v})
	case *VideoMetadata:
		return json.Marshal(struct {
			Type string `json:"type"`
			*VideoMetadata
		}{Type: "Video", VideoMetadata: v})
	case *AudioMetadata:
		return json.Marshal(struct {
			Type string `json:"type"`
		}{Type: "Audio"})
	default:
		return nil, fmt.Errorf("unknown Metadata variant: %T", r.Value)
	}
}

// UnmarshalJSON implements json.Unmarshaler for the Metadata tagged union.
func (r *RawMetadata) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	switch discriminator.Type {
	case "File":
		r.Value = &FileMetadata{}
	case "Text":
		r.Value = &TextMetadata{}
	case "Image":
		var v ImageMetadata
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Video":
		var v VideoMetadata
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Audio":
		r.Value = &AudioMetadata{}
	default:
		return fmt.Errorf("unknown metadata type: %q", discriminator.Type)
	}
	return nil
}

// Image represents an image in an embed.
type Image struct {
	URL    string    `json:"url"`
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Size   ImageSize `json:"size"`
}

// Video represents a video in an embed.
type Video struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// ImageSize is a string enum for image sizing.
type ImageSize string

const (
	ImageSizeLarge   ImageSize = "Large"
	ImageSizePreview ImageSize = "Preview"
)

// Special represents information about special remote content.
type Special struct {
	Type        string `json:"type"`
	ID          string `json:"id,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	AlbumID     string `json:"album_id,omitempty"`
	TrackID     string `json:"track_id,omitempty"`
}

// BandcampType is a string enum for Bandcamp content types.
type BandcampType string

const (
	BandcampTypeAlbum BandcampType = "Album"
	BandcampTypeTrack BandcampType = "Track"
)

// LightspeedType is a string enum for Lightspeed.tv content types.
type LightspeedType string

const (
	LightspeedTypeChannel LightspeedType = "Channel"
)

// TwitchType is a string enum for Twitch content types.
type TwitchType string

const (
	TwitchTypeChannel TwitchType = "Channel"
	TwitchTypeVideo   TwitchType = "Video"
	TwitchTypeClip    TwitchType = "Clip"
)
