package stoat

import (
	"encoding/json"
	"testing"
)

func TestMetadataFileRoundTrip(t *testing.T) {
	orig := RawMetadata{Value: &FileMetadata{}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawMetadata
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Value.MetadataType() != "File" {
		t.Errorf("type = %q, want File", got.Value.MetadataType())
	}
}

func TestMetadataTextRoundTrip(t *testing.T) {
	orig := RawMetadata{Value: &TextMetadata{}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawMetadata
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Value.MetadataType() != "Text" {
		t.Errorf("type = %q, want Text", got.Value.MetadataType())
	}
}

func TestMetadataImageRoundTrip(t *testing.T) {
	orig := RawMetadata{Value: &ImageMetadata{Width: 1920, Height: 1080}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawMetadata
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	img, ok := got.Value.(*ImageMetadata)
	if !ok {
		t.Fatalf("type = %T, want *ImageMetadata", got.Value)
	}
	if img.Width != 1920 || img.Height != 1080 {
		t.Errorf("dimensions = %dx%d, want 1920x1080", img.Width, img.Height)
	}
}

func TestMetadataVideoRoundTrip(t *testing.T) {
	orig := RawMetadata{Value: &VideoMetadata{Width: 1280, Height: 720}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawMetadata
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	vid, ok := got.Value.(*VideoMetadata)
	if !ok {
		t.Fatalf("type = %T, want *VideoMetadata", got.Value)
	}
	if vid.Width != 1280 || vid.Height != 720 {
		t.Errorf("dimensions = %dx%d, want 1280x720", vid.Width, vid.Height)
	}
}

func TestMetadataAudioRoundTrip(t *testing.T) {
	orig := RawMetadata{Value: &AudioMetadata{}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawMetadata
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Value.MetadataType() != "Audio" {
		t.Errorf("type = %q, want Audio", got.Value.MetadataType())
	}
}

func TestFileUnmarshal(t *testing.T) {
	raw := `{
		"_id": "file01",
		"tag": "attachments",
		"filename": "photo.png",
		"metadata": {"type": "Image", "width": 800, "height": 600},
		"content_type": "image/png",
		"size": 12345
	}`
	var f File
	if err := json.Unmarshal([]byte(raw), &f); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if f.ID != "file01" {
		t.Errorf("ID = %q", f.ID)
	}
	if f.Tag != "attachments" {
		t.Errorf("Tag = %q", f.Tag)
	}
	if f.Size != 12345 {
		t.Errorf("Size = %d", f.Size)
	}
	img, ok := f.Metadata.Value.(*ImageMetadata)
	if !ok {
		t.Fatalf("metadata type = %T", f.Metadata.Value)
	}
	if img.Width != 800 || img.Height != 600 {
		t.Errorf("dimensions = %dx%d", img.Width, img.Height)
	}
}
