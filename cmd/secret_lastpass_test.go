package cmd

import (
	"testing"

	"github.com/d4l3k/messagediff"
)

func Test_lastpassParseNote(t *testing.T) {
	for _, tc := range []struct {
		note string
		want map[string]string
	}{
		{
			note: "Foo:bar\n",
			want: map[string]string{
				"foo": "bar\n",
			},
		},
		{
			note: "Foo:bar\nbaz\n",
			want: map[string]string{
				"foo": "bar\nbaz\n",
			},
		},
		{
			note: "NoteType:SSH Key\nLanguage:en-US\nBit Strength:2048\nFormat:RSA\nPassphrase:Passphrase\nPrivate Key:-----BEGIN OPENSSH PRIVATE KEY-----\n-----END OPENSSH PRIVATE KEY-----\nPublic Key:ssh-rsa public-key user@host\nHostname:Hostname\nDate:Date\nNotes:",
			want: map[string]string{
				"noteType":    "SSH Key\n",
				"language":    "en-US\n",
				"bitStrength": "2048\n",
				"format":      "RSA\n",
				"passphrase":  "Passphrase\n",
				"privateKey":  "-----BEGIN OPENSSH PRIVATE KEY-----\n-----END OPENSSH PRIVATE KEY-----\n",
				"publicKey":   "ssh-rsa public-key user@host\n",
				"hostname":    "Hostname\n",
				"date":        "Date\n",
				"notes":       "\n",
			},
		},
	} {
		got := lastpassParseNote(tc.note)
		if diff, equal := messagediff.PrettyDiff(tc.want, got); !equal {
			t.Errorf("lastpassParseNote(%q) == %+v, want %+v\n%s", tc.note, got, tc.want, diff)
		}
	}
}
