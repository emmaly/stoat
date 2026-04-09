package stoat

import (
	"encoding/json"
	"testing"
)

func TestEmbedWebsiteRoundTrip(t *testing.T) {
	title := "Example"
	orig := RawEmbed{Value: &WebsiteEmbed{Title: &title}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawEmbed
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	ws, ok := got.Value.(*WebsiteEmbed)
	if !ok {
		t.Fatalf("type = %T, want *WebsiteEmbed", got.Value)
	}
	if ws.Title == nil || *ws.Title != "Example" {
		t.Errorf("Title = %v", ws.Title)
	}
}

func TestEmbedImageRoundTrip(t *testing.T) {
	orig := RawEmbed{Value: &ImageEmbed{URL: "https://example.com/img.png", Width: 800, Height: 600, Size: ImageSizeLarge}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawEmbed
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	img, ok := got.Value.(*ImageEmbed)
	if !ok {
		t.Fatalf("type = %T, want *ImageEmbed", got.Value)
	}
	if img.Width != 800 || img.Height != 600 {
		t.Errorf("dimensions = %dx%d", img.Width, img.Height)
	}
	if img.Size != ImageSizeLarge {
		t.Errorf("Size = %q", img.Size)
	}
}

func TestEmbedVideoRoundTrip(t *testing.T) {
	orig := RawEmbed{Value: &VideoEmbed{URL: "https://example.com/vid.mp4", Width: 1920, Height: 1080}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawEmbed
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	vid, ok := got.Value.(*VideoEmbed)
	if !ok {
		t.Fatalf("type = %T, want *VideoEmbed", got.Value)
	}
	if vid.Width != 1920 {
		t.Errorf("Width = %d", vid.Width)
	}
}

func TestEmbedTextRoundTrip(t *testing.T) {
	title := "Info"
	desc := "Some text"
	orig := RawEmbed{Value: &TextEmbed{Title: &title, Description: &desc}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawEmbed
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	te, ok := got.Value.(*TextEmbed)
	if !ok {
		t.Fatalf("type = %T, want *TextEmbed", got.Value)
	}
	if te.Title == nil || *te.Title != "Info" {
		t.Errorf("Title = %v", te.Title)
	}
}

func TestEmbedNoneRoundTrip(t *testing.T) {
	orig := RawEmbed{Value: &NoneEmbed{}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawEmbed
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if got.Value.EmbedType() != "None" {
		t.Errorf("type = %q", got.Value.EmbedType())
	}
}

func TestSystemMessageTextRoundTrip(t *testing.T) {
	orig := RawSystemMessage{Value: &SystemMessageText{Content: "hello"}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawSystemMessage
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	sm, ok := got.Value.(*SystemMessageText)
	if !ok {
		t.Fatalf("type = %T, want *SystemMessageText", got.Value)
	}
	if sm.Content != "hello" {
		t.Errorf("Content = %q", sm.Content)
	}
}

func TestSystemMessageUserAddedRoundTrip(t *testing.T) {
	orig := RawSystemMessage{Value: &SystemMessageUserAdded{ID: "user01", By: "user02"}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawSystemMessage
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	sm, ok := got.Value.(*SystemMessageUserAdded)
	if !ok {
		t.Fatalf("type = %T, want *SystemMessageUserAdded", got.Value)
	}
	if sm.ID != "user01" || sm.By != "user02" {
		t.Errorf("got ID=%q By=%q", sm.ID, sm.By)
	}
}

func TestSystemMessageChannelRenamedRoundTrip(t *testing.T) {
	orig := RawSystemMessage{Value: &SystemMessageChannelRenamed{Name: "new-name", By: "user01"}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawSystemMessage
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	sm, ok := got.Value.(*SystemMessageChannelRenamed)
	if !ok {
		t.Fatalf("type = %T", got.Value)
	}
	if sm.Name != "new-name" {
		t.Errorf("Name = %q", sm.Name)
	}
}

func TestSystemMessageMessagePinnedRoundTrip(t *testing.T) {
	orig := RawSystemMessage{Value: &SystemMessageMessagePinned{ID: "msg01", By: "user01"}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawSystemMessage
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	sm, ok := got.Value.(*SystemMessageMessagePinned)
	if !ok {
		t.Fatalf("type = %T", got.Value)
	}
	if sm.ID != "msg01" || sm.By != "user01" {
		t.Errorf("got ID=%q By=%q", sm.ID, sm.By)
	}
}

func TestMessageUnmarshalWithEmbeds(t *testing.T) {
	raw := `{
		"_id": "msg01",
		"channel": "ch01",
		"author": "user01",
		"content": "check this out",
		"embeds": [
			{"type": "Website", "title": "Example Site"},
			{"type": "None"}
		]
	}`
	var msg Message
	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if msg.ID != "msg01" {
		t.Errorf("ID = %q", msg.ID)
	}
	if len(msg.Embeds) != 2 {
		t.Fatalf("embeds len = %d, want 2", len(msg.Embeds))
	}
	if msg.Embeds[0].Value.EmbedType() != "Website" {
		t.Errorf("embed 0 type = %q", msg.Embeds[0].Value.EmbedType())
	}
	if msg.Embeds[1].Value.EmbedType() != "None" {
		t.Errorf("embed 1 type = %q", msg.Embeds[1].Value.EmbedType())
	}
}
