package spawn

import (
	"slices"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
)

// The default row carries the CLI's name and sits above both headings.
func TestModelListNamesTheDefaultRowAboveItsHeadings(t *testing.T) {
	h := newFakeChoiceHost()
	h.settings[recentModelsKey("claude")] = `["opus"]`
	ch := NewChoice(h, "claude")
	ch.OpenFilter(h, "claude")
	lines, entries := ch.ViewSuggestions(h, "claude", 0, 60, ModelListRows, Styles{})
	if !slices.Equal(entries, []int{0, -1, 1, -1, 2, 3}) {
		t.Fatalf("entries = %v", entries)
	}
	for i, want := range []string{"claude's default", "recent", "opus", "from claude · 3 models", "default", "haiku"} {
		if got := ansi.Strip(lines[i]); !strings.Contains(got, want) {
			t.Fatalf("line %d = %q, want %q", i, got, want)
		}
	}
}
