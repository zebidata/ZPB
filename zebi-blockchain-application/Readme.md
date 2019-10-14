------------------------------------------------------------------------------------------
#### Basic functionalities related to zebid and zebicli modules(Frequently used commands)
------------------------------------------------------------------------------------------

###### To reset a chains data and genesis file
> zebid unsafe-reset-all

###### To query a transaction hash
> zebicli query tx `<hash>`

###### To check list of existing accounts
> zebicli keys list

----------------------------------------------
#### Steps to run key value store application
----------------------------------------------

###### Initialize configuration files and genesis file
> zebid init `<moniker>` --chain-id `<chain-name>`

###### To add a key. Copy the Address output and save it for later use
> zebicli keys add `<account-holder-name>`

###### Add account, with coins to the genesis file. This is the genesis account. It can only be added before the chain is started. Once chain starts, then funds can only be transferred from one account to another. We can have multiple genesis accounts.
> zebid add-genesis-account $(zebicli keys show `<account-holder-name>` -a) `<amount>zebi`

**The value required should be multiplied to the power to 10^8 for decimal representation.**

###### Configure your CLI to eliminate need for chain-id flag
> zebicli config chain-id `<chain-name>`
> zebicli config output json
> zebicli config indent true
> zebicli config trust-node true

###### Create a genesis transaction
> zebid gentx --name `<account-holder-name>` --amount=`<amount>zebi` --commission-rate `0.01` --commission-max-rate `0.05` --commission-max-change-rate `0.005` --min-self-delegation `<amount>`

###### After generation of genesis transcation, input the gentx into the genesis file, so that the chain is aware of the validators
> zebid collect-gentxs

###### To make sure genesis is correct
> zebid validate-genesis

###### To start the client blockchain
> zebid start --minimum-gas-prices=`1400.0zebi` 1>>ZPBLog.$(date +%b%d-%H)  2>&1  &

###### To check the accounts to ensure they have funds
> zebicli query account $(zebicli keys show `<account-holder-name>` -a)

###### To make a transaction of post-key-value type
> zebicli tx kvstore post-key-value howrah bridge --from `<account-holder-name>` --chain-id `<chain-name>`

###### To start rest-server
> zebicli rest-server `<cert-file>` `<key-file>` --chain-id `<chain-name>` --trust-node --laddr `tcp://0.0.0.0:1317`

###### To check if value exists corresponding to a key on a blockchain
> zebicli query kvstore get-key-value howrah

->["bridge"]

----------------------------
#### Some rest API commands
----------------------------
###### Rest api command to search a key on blockchain
> curl -s 'http://localhost:1317/kvstore/key/<key>'

###### Rest api command to search a data on blockchain
> curl -s 'http://localhost:1317/kvstore/data/<data>'

###### Rest api command to query a transaction hash on blockchain
> curl -s 'http://localhost:1317/txs/<hash>'

###### Rest api command to query rewards
1. Delegator rewards
> curl -X GET "http://localhost:1317/distribution/delegators/<account_address>/rewards" -H "accept: application/json"

2. Validator rewards
> curl -X GET "http://localhost:1317/distribution/validators/<validator_address>/rewards" -H "accept: application/json"

###### Rest api command to query liquid balances
> curl -X GET "http://localhost:1317/bank/balances/<account_address>" -H "accept: application/json"

###### Rest api command to query account details
> curl -X GET "http://localhost:1317/auth/accounts/<account_address>" -H "accept: application/json"

###### Rest api command to query validators
> curl -X GET "http://localhost:1317/staking/validators" -H "accept: application/json"

###### Rest api command to query my delegations
1. To get all delegations from a validator
> curl -X GET "http://localhost:1317/staking/validators/<validator_address>/delegations" -H "accept: application/json"

2. To get all delegations from a delegator
> curl -X GET "http://localhost:1317/staking/delegators/<account_address>/delegations" -H "accept: application/json"

###### Rest api command to query all transactions from a particular account
> curl -X GET "http://localhost:1317/txs?sender=<account_address>" -H "accept: application/json"
> curl -X GET "http://localhost:1317/txs?recipient=<account_address>" -H "accept: application/json"

--------------------------------------------------------------------------------------------------
#### Commands to add a validator node. Should be done only after adding the node as a peer node.
--------------------------------------------------------------------------------------------------
> zebicli tx staking create-validator \
  --amount=`<amount>zebi` \
  --pubkey=$(zebid tendermint show-validator) \
  --moniker=`<node-name>` \
  --chain-id=`<chain-name>` \
  --commission-rate=`"0.10"` \
  --commission-max-rate=`"0.20"` \
  --commission-max-change-rate=`"0.01"` \
  --min-self-delegation=`"1"` \
  --gas="auto" \
  --gas-prices=`"0.025zebi"` \
  --from=`<account-holder-name>`

**Use "--help" or "-h" with zebid and zebicli for more information about a command.**
