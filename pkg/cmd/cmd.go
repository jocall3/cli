// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "jocall3",
		Usage:   "CLI for the jocall3 API",
		Suggest: true,
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "users",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&usersLogin,
					&usersRegister,
				},
			},
			{
				Name:     "users:password-reset",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&usersPasswordResetConfirm,
					&usersPasswordResetInitiate,
				},
			},
			{
				Name:     "users:me",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&usersMeRetrieve,
					&usersMeUpdate,
				},
			},
			{
				Name:     "users:me:preferences",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&usersMePreferencesRetrieve,
					&usersMePreferencesUpdate,
				},
			},
			{
				Name:     "users:me:devices",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&usersMeDevicesList,
					&usersMeDevicesRegister,
				},
			},
			{
				Name:     "users:me:biometrics",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&usersMeBiometricsEnroll,
					&usersMeBiometricsStatus,
					&usersMeBiometricsVerify,
				},
			},
			{
				Name:     "accounts",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&accountsLink,
					&accountsRetrieveDetails,
					&accountsRetrieveMe,
					&accountsRetrieveStatements,
				},
			},
			{
				Name:     "accounts:transactions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&accountsTransactionsRetrievePending,
				},
			},
			{
				Name:     "accounts:overdraft-settings",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&accountsOverdraftSettingsRetrieveOverdraftSettings,
					&accountsOverdraftSettingsUpdateOverdraftSettings,
				},
			},
			{
				Name:     "transactions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&transactionsRetrieve,
					&transactionsList,
					&transactionsCategorize,
					&transactionsDispute,
					&transactionsUpdateNotes,
				},
			},
			{
				Name:     "transactions:recurring",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&transactionsRecurringCreate,
					&transactionsRecurringList,
				},
			},
			{
				Name:     "transactions:insights",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&transactionsInsightsGetSpendingTrends,
				},
			},
			{
				Name:     "budgets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&budgetsCreate,
					&budgetsRetrieve,
					&budgetsUpdate,
					&budgetsList,
				},
			},
			{
				Name:     "investments:portfolios",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&investmentsPortfoliosCreate,
					&investmentsPortfoliosRetrieve,
					&investmentsPortfoliosUpdate,
					&investmentsPortfoliosList,
					&investmentsPortfoliosRebalance,
				},
			},
			{
				Name:     "investments:assets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&investmentsAssetsSearch,
				},
			},
			{
				Name:     "ai:advisor",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&aiAdvisorListTools,
				},
			},
			{
				Name:     "ai:advisor:chat",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&aiAdvisorChatRetrieveHistory,
					&aiAdvisorChatSendMessage,
				},
			},
			{
				Name:     "ai:oracle:simulate",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&aiOracleSimulateRunAdvanced,
					&aiOracleSimulateRunStandard,
				},
			},
			{
				Name:     "ai:oracle:simulations",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&aiOracleSimulationsRetrieve,
					&aiOracleSimulationsList,
				},
			},
			{
				Name:     "ai:incubator",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&aiIncubatorListPitches,
				},
			},
			{
				Name:     "ai:incubator:pitch",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&aiIncubatorPitchRetrieveDetails,
					&aiIncubatorPitchSubmit,
					&aiIncubatorPitchSubmitFeedback,
				},
			},
			{
				Name:     "ai:ads",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&aiAdsListGenerated,
					&aiAdsRetrieveStatus,
				},
			},
			{
				Name:     "ai:ads:generate",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&aiAdsGenerateAdvanced,
					&aiAdsGenerateStandard,
				},
			},
			{
				Name:     "corporate",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&corporatePerformSanctionScreening,
				},
			},
			{
				Name:     "corporate:cards",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&corporateCardsList,
					&corporateCardsCreateVirtual,
					&corporateCardsFreeze,
					&corporateCardsListTransactions,
					&corporateCardsUpdateControls,
				},
			},
			{
				Name:     "corporate:anomalies",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&corporateAnomaliesList,
					&corporateAnomaliesUpdateStatus,
				},
			},
			{
				Name:     "corporate:compliance:audits",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&corporateComplianceAuditsRequest,
					&corporateComplianceAuditsRetrieveReport,
				},
			},
			{
				Name:     "corporate:treasury",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&corporateTreasuryGetLiquidityPositions,
				},
			},
			{
				Name:     "corporate:treasury:cash-flow",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&corporateTreasuryCashFlowGetForecast,
				},
			},
			{
				Name:     "corporate:risk:fraud:rules",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&corporateRiskFraudRulesCreate,
					&corporateRiskFraudRulesUpdate,
					&corporateRiskFraudRulesList,
				},
			},
			{
				Name:     "web3",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&web3RetrieveNFTs,
				},
			},
			{
				Name:     "web3:wallets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&web3WalletsList,
					&web3WalletsConnect,
					&web3WalletsRetrieveBalances,
				},
			},
			{
				Name:     "web3:transactions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&web3TransactionsInitiateTransfer,
				},
			},
			{
				Name:     "payments:international",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&paymentsInternationalInitiate,
					&paymentsInternationalRetrieveStatus,
				},
			},
			{
				Name:     "payments:fx",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&paymentsFxConvert,
					&paymentsFxRetrieveRates,
				},
			},
			{
				Name:     "sustainability",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&sustainabilityPurchaseCarbonOffsets,
					&sustainabilityRetrieveCarbonFootprint,
				},
			},
			{
				Name:     "sustainability:investments",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&sustainabilityInvestmentsAnalyzeImpact,
				},
			},
			{
				Name:     "lending:applications",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&lendingApplicationsRetrieve,
					&lendingApplicationsSubmit,
				},
			},
			{
				Name:     "lending:offers",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&lendingOffersListPreApproved,
				},
			},
			{
				Name:     "developers:webhooks",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&developersWebhooksCreate,
					&developersWebhooksUpdate,
					&developersWebhooksList,
				},
			},
			{
				Name:     "developers:api-keys",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&developersAPIKeysCreate,
					&developersAPIKeysList,
				},
			},
			{
				Name:     "identity:kyc",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&identityKYCRetrieveStatus,
					&identityKYCSubmit,
				},
			},
			{
				Name:     "goals",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&goalsCreate,
					&goalsRetrieve,
					&goalsUpdate,
					&goalsList,
				},
			},
			{
				Name:     "notifications",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&notificationsListUserNotifications,
					&notificationsMarkAsRead,
				},
			},
			{
				Name:     "notifications:settings",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&notificationsSettingsRetrieve,
					&notificationsSettingsUpdate,
				},
			},
			{
				Name:     "marketplace:products",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&marketplaceProductsList,
					&marketplaceProductsSimulateImpact,
				},
			},
			{
				Name:     "marketplace:offers",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&marketplaceOffersRedeem,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "jocall3 @manpages [-o jocall3.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
		},
		EnableShellCompletion:      true,
		ShellCompletionCommandName: "@completion",
		HideHelpCommand:            true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "jocall3.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "jocall3.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
