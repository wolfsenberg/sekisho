package demo

import (
	"context"
	"testing"

	"github.com/wolfsenberg/sekisho/internal/accounts"
	"github.com/wolfsenberg/sekisho/internal/upstream"
)

func TestDemoResetCardsStayOfflineAndShowExpiryMetadata(t *testing.T) {
	client := upstream.NewWithTransport(transport{})
	defer client.Close()
	cards, err := client.ResetCards(context.Background(), accounts.Credential{AccessToken: "synthetic-demo-access"})
	if err != nil || len(cards.Cards) != 2 || cards.Cards[0].ExpiresAt == nil || cards.Cards[1].ExpiresAt != nil {
		t.Fatal(cards, err)
	}
}
