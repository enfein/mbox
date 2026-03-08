package mieru

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sagernet/sing-box/option"

	mieruserver "github.com/enfein/mieru/v3/apis/server"
)

func baseMieruInboundOptions() option.MieruInboundOptions {
	return option.MieruInboundOptions{
		ListenOptions: option.ListenOptions{
			ListenPort: 8964,
		},
		Transport: "TCP",
		Users: []option.MieruUser{
			{
				Name:     "testuser",
				Password: "testpassword",
			},
		},
	}
}

func TestBuildMieruServerConfigAdvancedSettings(t *testing.T) {
	testCases := []struct {
		name                    string
		userHintIsMandatory     bool
		tcpIdleTimeout          int
		wantErr                 bool
		wantAdvancedSettingsNil bool
		wantUserHintIsMandatory bool
		wantTCPIdleTimeout      string
	}{
		{
			name:                    "default_unset",
			wantAdvancedSettingsNil: true,
		},
		{
			name:               "tcp_idle_timeout_minimum",
			tcpIdleTimeout:     60,
			wantTCPIdleTimeout: "60s",
		},
		{
			name:               "tcp_idle_timeout_1800",
			tcpIdleTimeout:     1800,
			wantTCPIdleTimeout: "1800s",
		},
		{
			name:                    "user_hint_only",
			userHintIsMandatory:     true,
			wantUserHintIsMandatory: true,
		},
		{
			name:                    "both_advanced_settings",
			userHintIsMandatory:     true,
			tcpIdleTimeout:          1800,
			wantUserHintIsMandatory: true,
			wantTCPIdleTimeout:      "1800s",
		},
		{
			name:           "tcp_idle_timeout_too_small",
			tcpIdleTimeout: 59,
			wantErr:        true,
		},
		{
			name:           "tcp_idle_timeout_negative",
			tcpIdleTimeout: -1,
			wantErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			opts := baseMieruInboundOptions()
			opts.UserHintIsMandatory = tc.userHintIsMandatory
			opts.TCPIdleTimeout = tc.tcpIdleTimeout

			cfg, _, err := buildMieruServerConfig(context.Background(), opts)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			adv := cfg.Config.GetAdvancedSettings()
			if tc.wantAdvancedSettingsNil {
				if adv != nil {
					t.Fatalf("expected nil AdvancedSettings, got %v", adv)
				}
			} else {
				if adv == nil {
					t.Fatalf("expected non-nil AdvancedSettings, got nil")
				}
				if adv.GetUserHintIsMandatory() != tc.wantUserHintIsMandatory {
					t.Fatalf("UserHintIsMandatory = %v, want %v", adv.GetUserHintIsMandatory(), tc.wantUserHintIsMandatory)
				}
				if adv.GetTcpIdleTimeout() != tc.wantTCPIdleTimeout {
					t.Fatalf("TcpIdleTimeout = %q, want %q", adv.GetTcpIdleTimeout(), tc.wantTCPIdleTimeout)
				}
			}

			s := mieruserver.NewServer()
			if err := s.Store(cfg); err != nil {
				t.Fatalf("mieru server Store failed: %v", err)
			}
		})
	}
}

func TestMieruInboundOptionsJSON(t *testing.T) {
	raw := `{
		"listen_port": 8964,
		"transport": "TCP",
		"users": [{"name": "baozi", "password": "secret"}],
		"user_hint_is_mandatory": true,
		"tcp_idle_timeout": 1800
	}`

	var opts option.MieruInboundOptions
	if err := json.Unmarshal([]byte(raw), &opts); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if !opts.UserHintIsMandatory {
		t.Fatalf("UserHintIsMandatory = false, want true")
	}
	if opts.TCPIdleTimeout != 1800 {
		t.Fatalf("TCPIdleTimeout = %d, want 1800", opts.TCPIdleTimeout)
	}

	cfg, _, err := buildMieruServerConfig(context.Background(), opts)
	if err != nil {
		t.Fatalf("buildMieruServerConfig failed: %v", err)
	}
	if got := cfg.Config.GetAdvancedSettings().GetTcpIdleTimeout(); got != "1800s" {
		t.Fatalf("TcpIdleTimeout = %q, want %q", got, "1800s")
	}
}
