package striphtmltags_test

import (
	"fmt"

	"github.com/denisbrodbeck/striphtmltags"
)

// Strip HTML tag from strings:
func Example() {
	html := `<script>...</script> <b>&iexcl;Hi!</b>`
	got := striphtmltags.StripTags(html)
	fmt.Println(got)
	// Output:  &iexcl;Hi!
}
