// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/jocall3/1231-cli/internal/mocktest"
)

func TestAIAdsGenerateAdvanced(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:ads:generate", "advanced",
		"--length-seconds", "15",
		"--prompt", "A captivating ad featuring a young entrepreneur using 's AI tools to grow their startup. Focus on innovation and ease of use.",
		"--style", "Cinematic",
		"--aspect-ratio", "16:9",
		"--audience-target", "corporate",
		"--background-music-genre", "corporate",
		"--brand-asset", "https://demobank.com/assets/corporate_logo.png",
		"--brand-color", "#0000FF",
		"--brand-color", "#FFD700",
		"--call-to-action", "{displayTimeSeconds: 5, text: Learn more at DemoBank.com/business, url: https://demobank.com/business}",
		"--keyword", "innovation",
		"--keyword", "fintech",
		"--keyword", "startup",
		"--voiceover-style", "male_professional",
		"--voiceover-text", ": Your business, powered by intelligent finance.",
	)
}

func TestAIAdsGenerateStandard(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"ai:ads:generate", "standard",
		"--length-seconds", "15",
		"--prompt", "A captivating ad featuring a young entrepreneur using 's AI tools to grow their startup. Focus on innovation and ease of use.",
		"--style", "Cinematic",
		"--aspect-ratio", "16:9",
		"--brand-color", "#0000FF",
		"--brand-color", "#FFD700",
		"--keyword", "innovation",
		"--keyword", "fintech",
		"--keyword", "startup",
	)
}
