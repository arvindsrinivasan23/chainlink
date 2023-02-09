# test-tooling-go-homework
Homework assignment for SET candidates to Test Tooling team (SDET)

## SET Homework Project
The end goal of the project is to show that areas of work can be solely owned and solutions can be critically thought out without being given much direction.
Expected time to complete: 1-3 days

## High Level Goal
Create a test script that analyses previous Chainlink feed rounds and tests answer deviation.

## Solution Design
To be decided solely by the candidate, as long as it meets the requirements of the epic and is written in Golang.

## Must have Requirements
- The BTC/USD feed is tested: https://etherscan.io/address/0xf570deefff684d964dc3e15e1f9414283e3f7419
- At least 5 previous rounds are fetched and tested
- Each individual answer from each oracle is compared against the aggregated median
- The test will calculate the percentile difference between each individual answer and the aggregated median, passing the test if all answers are within 10% of the median
- The test must be ran and written within a common testing harness in the respective language
- Instructions how to run this project are documented

## Nice to have Requirements
- Test is data-driven, with each scenario allowing the feed ETH address and deviation threshold to be configured.

## Communication with CLL members
As part of the project work you will be invited to the Chainlink Slack channel with members of the Test Tooling team and any questions can be asked along the way of working on the project.

## Providing Solution
Please open PR to this repository and add `sebawo` as a reviewer
