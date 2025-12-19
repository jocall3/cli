// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/cli/internal/mocktest"
)

func TestMarketplaceOffersRedeem(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"marketplace:offers", "redeem",
		"--offer-id", "offer_home_ins_promo_1",
		"--payment-account-id", "acc_chase_checking_4567",
	)
}
