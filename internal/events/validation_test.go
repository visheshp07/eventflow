package events

import "testing"

func TestIsValidEventType(t *testing.T) {
	tests := []struct {
		name      string
		eventType string
		expected  bool
	}{
		{"Valid event", "device.offline", true},
		{"Empty event", "", false},
		{"Whitespace event", "   ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsValidEventType(tt.eventType)

			if actual != tt.expected {
				t.Errorf(
					"expected %v, got %v",
					tt.expected,
					actual,
				)
			}
		})
	}
}
