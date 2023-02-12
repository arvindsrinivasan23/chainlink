# Steps to run the test

## Test config 

Add a new test case by appending a test case object to the []TestCase slice.
Set the following fields for each test case -
1. name - Provide a name for the test case.
2. ethAddress -  Address of the smart contract.
3. deviation - The percentage deviation that is acceptable for each oracle submission
4. numOfRounds - Number of rounds to be considered for calculating the median. If the value passed is less than 5 it will default to 5.


## Execute test

Navigate to the test directory in terminal and run the cmd 'go test -v -run TestContract
'.