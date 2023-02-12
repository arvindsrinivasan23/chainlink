package test

import (
	"github.com/sebawo/test-tooling-go-homework-arvind/pkg/client"
	"github.com/sebawo/test-tooling-go-homework-arvind/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContract(t *testing.T) {
	const INFURA_URL = "https://mainnet.infura.io/v3/5fcd8d8c7b234c85ac2bfc1a5650e9b9"

	type TestCase struct {
		name        string
		ethAddress  string
		deviation   int
		numOfRounds int
	}

	testCases := []TestCase{
		{
			name:        "Validate that answer from each oracle in BTC2USD feed is within 10% of aggregated median",
			ethAddress:  "0xF570deEffF684D964dc3E15E1F9414283E3f7419",
			deviation:   10,
			numOfRounds: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			etherScan, err := client.NewBTC2USDValidatorClient(&client.BTC2USDConf{URL: INFURA_URL, Address: tc.ethAddress})
			assert.NoError(t, err)
			median, err := etherScan.GetMedianAnsForLast5PlusRounds(int64(tc.numOfRounds))
			assert.NoError(t, err)
			allOraclesRoundData, err := etherScan.GetSubDataForAllOracles(int64(tc.numOfRounds))
			assert.NoError(t, err)
			for _, oracleRoundData := range allOraclesRoundData {
				assert.Truef(t, utils.WithinXPercent(tc.deviation, median, oracleRoundData.LatestSub),
					"Oracle %v submission for round %v not within %v percent of median", oracleRoundData.Oracle, oracleRoundData.Round, tc.deviation)
			}
		})
	}
}
