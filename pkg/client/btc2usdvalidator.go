package client

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sebawo/test-tooling-go-homework-arvind/pkg/contracts"
	"github.com/sebawo/test-tooling-go-homework-arvind/pkg/utils"
	"math/big"
	"sync"
)

type OracleRoundData struct {
	Oracle    common.Address
	Round     uint32
	LatestSub *big.Int
}

// BTC2USDValidatorClient is a wrapper around the contracts caller struct. Methods associated with this
// struct will run computations on the data retrieved from the smart contract and expose only what's
// necessary for the test
type BTC2USDValidatorClient struct {
	contCaller *contracts.ContractsCaller
}

type BTC2USDConf struct {
	URL     string
	Address string
}

func NewBTC2USDValidatorClient(conf *BTC2USDConf) (*BTC2USDValidatorClient, error) {
	conn, err := ethclient.Dial(conf.URL)
	if err != nil {
		return nil, err
	}
	contractCaller, err := contracts.NewContractsCaller(common.HexToAddress(conf.Address), conn)

	if err != nil {
		return nil, err
	}

	return &BTC2USDValidatorClient{contCaller: contractCaller}, nil
}

func (esc *BTC2USDValidatorClient) GetOracles() ([]common.Address, error) {
	return esc.contCaller.GetOracles(&bind.CallOpts{})
}

func (esc *BTC2USDValidatorClient) GetOracleRoundData(address common.Address, round uint32) (*OracleRoundData, error) {
	oracleRoundState, err := esc.contCaller.OracleRoundState(&bind.CallOpts{}, address, round)

	if err != nil {
		return nil, err
	}

	return &OracleRoundData{
		Oracle:    address,
		Round:     round,
		LatestSub: oracleRoundState.LatestSubmission,
	}, nil
}

// GetMedianAnsForLast5PlusRounds Gets the median for the last 5 or more rounds.
// If the number of rounds is less than 5 it returns the median for the last 5 rounds.
func (esc *BTC2USDValidatorClient) GetMedianAnsForLast5PlusRounds(x int64) (*big.Int, error) {

	var answers []*big.Int
	var latestRound *big.Int

	// check number of rounds input and set it to 5 if it is less than 5
	x = utils.Max(x, 5)

	latestRound, err := esc.getLastRound()

	if err != nil {
		return nil, err
	}
	one := big.NewInt(1)
	firstRound := new(big.Int).Set(latestRound)
	firstRound = firstRound.Sub(firstRound, big.NewInt(x-1))
	for i := new(big.Int).Set(firstRound); i.Cmp(latestRound) <= 0; i.Add(i, one) {
		currRound, err := esc.contCaller.GetRoundData(&bind.CallOpts{}, i)

		if nil != err {
			return nil, err
		}
		answers = append(answers, currRound.Answer)
	}

	return utils.GetMedianOfBigIntSlice(answers), nil
}

// GetSubDataForAllOracles Fetches the latest submission for each oracle
func (esc *BTC2USDValidatorClient) GetSubDataForAllOracles(x int64) ([]*OracleRoundData, error) {
	var wg sync.WaitGroup
	allOracles, err := esc.contCaller.GetOracles(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	var latestRound *big.Int
	latestRound, err = esc.getLastRound()

	if err != nil {
		return nil, err
	}

	firstRound := new(big.Int).Set(latestRound)
	firstRound.Sub(firstRound, big.NewInt(x-1))
	results := make(chan *OracleRoundData, len(allOracles)*int(x))

	for _, oracle := range allOracles {
		wg.Add(1)
		go func(oracle common.Address, i *big.Int) {
			defer wg.Done()
			data, _ := esc.GetOracleRoundData(oracle, uint32(i.Uint64()))
			results <- data
		}(oracle, latestRound)

	}

	wg.Wait()
	close(results)

	var allOracleRoundData []*OracleRoundData

	for result := range results {
		allOracleRoundData = append(allOracleRoundData, result)
	}

	return allOracleRoundData, nil
}

func (esc *BTC2USDValidatorClient) getLastRound() (*big.Int, error) {
	latestRoundData, err := esc.contCaller.LatestRoundData(&bind.CallOpts{})

	if err != nil {
		return nil, err
	}

	// check if latest round hasn't started and if it hasn't then take previous round as the latest round
	if latestRoundData.StartedAt == big.NewInt(0) {
		var latestRound *big.Int
		latestRound = latestRound.Sub(latestRoundData.RoundId, big.NewInt(1))
		return latestRound, nil
	}

	return latestRoundData.RoundId, nil

}
