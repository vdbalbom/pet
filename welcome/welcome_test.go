package welcome_test

import (
	"testing"

	"github.com/pet/welcome"
)

func TestWelcomeSentence(t *testing.T) {
	if welcome.Sentence() != "This is the pet project!" {
		t.Errorf("The sentence must be: This is the pet project!")
	}
}
