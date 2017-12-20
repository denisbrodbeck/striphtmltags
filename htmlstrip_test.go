package striphtmltags

import "testing"

func TestStripTags(t *testing.T) {
	html := `<b>&iexcl;Hi!</b> <script>...</script>`
	want := `&iexcl;Hi! `
	if got := StripTags(html); got != want {
		t.Errorf("StripTags() = %q, want %q", got, want)
	}
}
