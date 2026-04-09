package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryNode(t *testing.T) {
	mockConfig := RevoltConfig{
		Revolt: "0.12.0",
		Features: RevoltFeatures{
			Captcha: CaptchaFeature{
				Enabled: true,
				Key:     "test-captcha-key",
			},
			Email:      true,
			InviteOnly: false,
			Autumn: Feature{
				Enabled: true,
				URL:     "https://cdn.example.com",
			},
			January: Feature{
				Enabled: true,
				URL:     "https://external.example.com",
			},
			LiveKit: VoiceFeature{
				Enabled: true,
				Nodes: []VoiceNode{
					{
						Name:      "us-east",
						Lat:       40.7,
						Lon:       -74.0,
						PublicURL: "https://voice.example.com",
					},
				},
			},
			Limits: LimitsConfig{
				Global: GlobalLimits{
					GroupSize:               100,
					MessageEmbeds:           10,
					MessageReplies:          5,
					MessageReactions:        20,
					ServerEmoji:             100,
					ServerRoles:             200,
					ServerChannels:          200,
					BodyLimitSize:           20000,
					RestrictServerCreation:  []string{},
				},
				NewUser: UserLimits{
					OutgoingFriendRequests: 10,
					Bots:                  2,
					MessageLength:         2000,
					MessageAttachments:    5,
					Servers:               50,
					VoiceQuality:          128,
					Video:                 true,
					VideoResolution:       []int{720, 1280},
					VideoAspectRatio:      []float64{0.5, 2.0},
					FileUploadSizeLimits:  map[string]int{"default": 5000000},
				},
				Default: UserLimits{
					OutgoingFriendRequests: 50,
					Bots:                  10,
					MessageLength:         4000,
					MessageAttachments:    10,
					Servers:               100,
					VoiceQuality:          256,
					Video:                 true,
					VideoResolution:       []int{1080, 1920},
					VideoAspectRatio:      []float64{0.5, 2.0},
					FileUploadSizeLimits:  map[string]int{"default": 10000000},
				},
			},
		},
		WS:    "wss://stoat.chat/events",
		App:   "https://stoat.chat",
		VAPID: "test-vapid-key",
		Build: BuildInformation{
			CommitSHA:       "abc123",
			CommitTimestamp: "2024-01-01T00:00:00Z",
			Semver:          "0.12.0",
			OriginURL:       "https://github.com/stoatchat/stoatchat",
			Timestamp:       "2024-01-01T00:00:00Z",
		},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Errorf("path = %q, want /", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		// No auth should be required
		if r.Header.Get("X-Session-Token") != "" {
			t.Error("unexpected X-Session-Token header")
		}
		w.Header().Set("X-RateLimit-Limit", "20")
		w.Header().Set("X-RateLimit-Bucket", "root")
		w.Header().Set("X-RateLimit-Remaining", "19")
		w.Header().Set("X-RateLimit-Reset-After", "10000")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockConfig)
	}))
	defer srv.Close()

	c, err := New(srv.URL)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	config, err := c.QueryNode(context.Background())
	if err != nil {
		t.Fatalf("QueryNode: %v", err)
	}

	if config.Revolt != "0.12.0" {
		t.Errorf("Revolt = %q", config.Revolt)
	}
	if config.WS != "wss://stoat.chat/events" {
		t.Errorf("WS = %q", config.WS)
	}
	if config.App != "https://stoat.chat" {
		t.Errorf("App = %q", config.App)
	}
	if config.VAPID != "test-vapid-key" {
		t.Errorf("VAPID = %q", config.VAPID)
	}
	if !config.Features.Captcha.Enabled {
		t.Error("expected captcha enabled")
	}
	if config.Features.Captcha.Key != "test-captcha-key" {
		t.Errorf("captcha key = %q", config.Features.Captcha.Key)
	}
	if !config.Features.Email {
		t.Error("expected email enabled")
	}
	if config.Features.InviteOnly {
		t.Error("expected invite_only to be false")
	}
	if config.Features.Autumn.URL != "https://cdn.example.com" {
		t.Errorf("autumn url = %q", config.Features.Autumn.URL)
	}
	if config.Features.January.URL != "https://external.example.com" {
		t.Errorf("january url = %q", config.Features.January.URL)
	}
	if !config.Features.LiveKit.Enabled {
		t.Error("expected livekit enabled")
	}
	if len(config.Features.LiveKit.Nodes) != 1 {
		t.Fatalf("livekit nodes = %d", len(config.Features.LiveKit.Nodes))
	}
	if config.Features.LiveKit.Nodes[0].Name != "us-east" {
		t.Errorf("node name = %q", config.Features.LiveKit.Nodes[0].Name)
	}
	if config.Build.Semver != "0.12.0" {
		t.Errorf("build semver = %q", config.Build.Semver)
	}
	if config.Features.Limits.Global.GroupSize != 100 {
		t.Errorf("group size = %d", config.Features.Limits.Global.GroupSize)
	}
	if config.Features.Limits.Default.MessageLength != 4000 {
		t.Errorf("default message length = %d", config.Features.Limits.Default.MessageLength)
	}
}

func TestQueryNodeError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Limit", "20")
		w.Header().Set("X-RateLimit-Bucket", "root")
		w.Header().Set("X-RateLimit-Remaining", "0")
		w.Header().Set("X-RateLimit-Reset-After", "10000")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"type": "InternalError"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	_, err := c.QueryNode(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
}
