-------------------------------------------------------------------------------------
#Basic functionalities related to zebid and zebicli modules(Frequently used commands)
-------------------------------------------------------------------------------------

#To reset a chains data and genesis file
zebid unsafe-reset-all

#To query a transaction hash
zebicli query tx <hash>

#To check list of existing accounts
zebicli keys list

#To start a rest server
zebicli rest-server cert.pem key.pem --chain-id <chain-name> --trust-node --laddr tcp://0.0.0.0:1317

----------------------------------
#Steps to run keytore application
----------------------------------

# Initialize configuration files and genesis file
zebid init <moniker> --chain-id <chain-name>

# To add a key. Copy the `Address` output here and save it for later use
zebicli keys add <account-holder-name>

#Add account, with coins to the genesis file. This is the genesis account. It can only be added before the chain is started. Once chain starts, then funds can only be transferred from one account to another. We can have multiple genesis accounts.
zebid add-genesis-account $(zebicli keys show <account-holder-name> -a) 10000000000000000zebi

note:The value required should be multiplied to the power to 10^8 .

# Configure your CLI to eliminate need for chain-id flag
zebicli config chain-id <chain-name>
zebicli config output json
zebicli config indent true
zebicli config trust-node true

#Create a genesis transaction
zebid gentx --name <account-holder-name> --amount=2000000000000000zebi --commission-rate 0.01 --commission-max-rate 0.05 --commission-max-change-rate 0.005 --min-self-delegation 100000000

#After generation of genesis transcation, input the gentx into the genesis file, so that the chain is aware of the validators
zebid collect-gentxs

#To make sure genesis is correct
zebid validate-genesis

#To start the client blockchain
zebid start --minimum-gas-prices=1400.0zebi 1>>ZPBLog.$(date +%b%d-%H)  2>&1  &

#To check the accounts to ensure they have funds
zebicli query account $(zebicli keys show <account-holder-name> -a)

#To make a transaction of post-key-value type
zebicli tx kvstore post-key-value howrah bridge --from <account-holder-name> --chain-id <chain-name>

#To start rest-server
zebicli rest-server cert.pem key.pem --chain-id <chain-name> --trust-node --laddr tcp://0.0.0.0:1317

#To check if value exists corresponding to a key on a blockchain
zebicli query kvstore get-key-value howrah
->["bridge"]

------------------------
#Some rest API commands
------------------------
#Rest api command to create an unsigned transaction
1. For token transfer transaction
curl --insecure -X POST "https://localhost:1317/bank/accounts/zebi1jhjsyvtmytumw236lslc9z5zf7n7gf5au8auj6/transfers" -H "accept: application/json" -H "Content-Type: application/json" -d '{"base_req":{"from":"zebi19cj9h7yka3r4ln2le75mspyrk2t0h8l82dx5dt","chain_id":"zpb-testnet","memo":"","gas_adjustment":"1.5","simulate":true,"gas":"200000","fees":[{"denom":"zebi","amount":"0"}]},"amount":[{"denom":"zebi","amount":"117000000"}]}'
#if simulate is set to true
=======>{"gas_estimate":"37678"}
#if simulate is set to false
=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgSend","value":{"from_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","to_address":"zebi1qj857c5a7me5ffr5zcjd3qm94k7tnd2y049ppq","amount":[{"denom":"stake","amount":"50"}]}}],"fee":{"amount":[{"denom":"stake","amount":"50"}],"gas":"200000"},"signatures":null,"memo":"Sent via Zebi"}}

2. For key value post type transaction
curl --insecure -XPOST -s https://localhost:1317/kvstore/key --data-binary '{"base_req":{"from":"'$(zebicli keys show enakshi -a)'","chain_id":"keychain","simulate":false},"key":"jack","value":"4","sender":"'$(zebicli keys show enakshi -a)'"}'

3. For delegation transaction
curl --insecure -X POST "https://localhost:1317/staking/delegators/zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy/delegations" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"base_req\": { \"from\": \"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy\", \"memo\": \"Sent via sharuk\", \"chain_id\": \"zebichain\", \"account_number\": \"0\", \"sequence\": \"37\", \"gas\": \"200000\", \"gas_adjustment\": \"1.2\", \"fees\": [ { \"denom\": \"stake\", \"amount\": \"0\" } ], \"simulate\": false }, \"delegator_address\": \"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy\", \"validator_address\": \"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39\", \"amount\": { \"denom\": \"stake\", \"amount\": \"5\" }}"

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgDelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":null,"memo":"Sent via sharuk"}}

4. For undelegation transaction
curl --insecure -X POST "https://localhost:1317/staking/delegators/zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy/unbonding_delegations" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"base_req\": { \"from\": \"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy\", \"memo\": \"Sent via Sharuk\", \"chain_id\": \"zebichain\", \"account_number\": \"3\", \"sequence\": \"5\", \"gas\": \"200000\", \"gas_adjustment\": \"1.2\", \"fees\": [ { \"denom\": \"stake\", \"amount\": \"0\" } ], \"simulate\": false }, \"delegator_address\": \"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy\", \"validator_address\": \"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39\", \"amount\": { \"denom\": \"stake\", \"amount\": \"5\" }}"

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgUndelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":null,"memo":"Sent via Sharuk"}}

5. For redelegation transaction
curl -X POST "http://localhost:1317/staking/delegators/zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy/redelegations" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"base_req\": { \"from\": \"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy\", \"memo\": \"Sent via Sharuk\", \"chain_id\": \"zebichain\", \"account_number\": \"3\", \"sequence\": \"6\", \"gas\": \"200000\", \"gas_adjustment\": \"1.2\", \"fees\": [ { \"denom\": \"stake\", \"amount\": \"0\" } ], \"simulate\": false }, \"delegator_address\": \"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy\", \"validator_src_address\": \"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39\", \"validator_dst_address\": \"zebivaloper1e23tczvtjz8fduywxeuf0rjpr9f69992cgzqhs\", \"amount\": { \"denom\": \"stake\", \"amount\": \"5\" }}"

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgBeginRedelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_src_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","validator_dst_address":"zebivaloper1e23tczvtjz8fduywxeuf0rjpr9f69992cgzqhs","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":null,"memo":"Sent via Sharuk"}}

6. For create validator transaction (also called as bonding/staking)
curl -XPOST -s http://localhost:1317/staking/validators/zebi1da7fsmk7rmcslhu2kfguqa9jfxma0443rhzgtg/stake --data-binary '{"base_req":{"from":"zebi1da7fsmk7rmcslhu2kfguqa9jfxma0443rhzgtg","chain_id":"devchain"},"commissionrate":"0.1","commissionmaxchangerate":"0.13","commissionmaxrate":"0.2","minselfdelegation":"10","validatoraddress":"zebivaloper1da7fsmk7rmcslhu2kfguqa9jfxma04433f8jmd","pubkey":"zebivalconspub1zcjduepqm76lr6z8j6ajdvlgfequksex309xswlnptxzsk4mkek70pspmwts9c5cv5","moniker":"nodeJ","identity":"Jayasimha ka node","website":"https://cabzz.in","details":"Testing node - Please do not delegate to this validator","denom":"stake","amount":"250"}'
(To get validatoraddress (the one which has 'zebivaloper' prefix) run 'zebicli keys show <account_name> --bech=val')
(To get pubkey of the node (the one which has 'zebivalconspub' prefix) run 'zebid tendermint show-validator')

Output:-  {"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"description":{"moniker":"nodeJ","identity":"Jayasimha ka node","website":"https://cabzz.in","details":"Testing node - Please do not delegate to this validator"},"commission":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.130000000000000000"},"min_self_delegation":"10","delegator_address":"zebi1da7fsmk7rmcslhu2kfguqa9jfxma0443rhzgtg","validator_address":"zebivaloper1da7fsmk7rmcslhu2kfguqa9jfxma04433f8jmd","pubkey":"zebivalconspub1zcjduepqm76lr6z8j6ajdvlgfequksex309xswlnptxzsk4mkek70pspmwts9c5cv5","value":{"denom":"stake","amount":"250"}}}],"fee":{"amount":[],"gas":"200000"},"signatures":null,"memo":""}}

7. For create account
curl -XPOST -s http://localhost:1317/Acc --data-binary '{"name":"enakship1","type":"add","pass":"enakshipriya","mneumonic":""'}

=======>{"name":"enakship1","type":"local","acc_address":"zebi1yrxuvw9c20ju0xt9s0vd878fpzxvkt7txzxr67","acc_pubkey":"zebipub1addwnpepq0z70sr7f8ruhcv9lghed5sd7hxq4r9nug4r6efevfk7jvvv50zgqyx4dw7","val_address":"zebivaloper1yrxuvw9c20ju0xt9s0vd878fpzxvkt7t5ure2m","val_pubkey":"zebivaloperpub1addwnpepq0z70sr7f8ruhcv9lghed5sd7hxq4r9nug4r6efevfk7jvvv50zgqu2jktx","mnemonic":"relief estate mind damp donor labor sauce doll fatigue tag summer someone dolphin poem marble hungry innocent prize repeat feel motion bread admit baby","armorprivkeys":"-----BEGIN TENDERMINT PRIVATE KEY-----\nkdf: bcrypt\nsalt: E2F12E449BA51010735468BDD944904A\n\n+UG1cLZE/Dhok9VTwbg+tVJeqDDrmxifx6+iE8wO9olAWmDVDsiQcyt636cg4v5M\nI2Wz4RnRX/tYUEtNPIrAIvDzjpacM2Eh/+03Whs=\n=bRAs\n-----END TENDERMINT PRIVATE KEY-----"}

8. For withdraw rewards
curl -X POST "http://localhost:1317/distribution/delegators/zebi13yqsaue50jsd2d7ygk0pzdsvwjn3uert9f7nn8/rewards" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"base_req\": { \"from\": \"zebi13yqsaue50jsd2d7ygk0pzdsvwjn3uert9f7nn8\", \"memo\": \"\", \"chain_id\": \"zebichain\", \"account_number\": \"0\", \"sequence\": \"252\", \"gas\": \"200000\", \"gas_adjustment\": \"1.2\", \"fees\": [ { \"denom\": \"stake\", \"amount\": \"1\" } ], \"simulate\": false }}"

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgWithdrawDelegationReward","value":{"delegator_address":"zebi13yqsaue50jsd2d7ygk0pzdsvwjn3uert9f7nn8","validator_address":"zebivaloper13yqsaue50jsd2d7ygk0pzdsvwjn3uerthhmfrz"}}],"fee":{"amount":[{"denom":"stake","amount":"1"}],"gas":"200000"},"signatures":null,"memo":""}}


#Rest api command to sign a transaction
1. For token transfer transaction
curl --insecure -XPOST -s https://localhost:1317/kvstore/signTx --data-binary '{"type":"auth/StdTx","armorprivkeys":"-----BEGIN TENDERMINT PRIVATE KEY-----\nkdf: bcrypt\nsalt: 0C93EB0B2BA781AC878673CA9ADBF99D\n\nXthsorOpsPphkNp3YWsUDpi3dGkSPO1jnQxOr4wqvH7YDXSe1KQ1aHMRx4WNxRKF\nebaqZmF5WY75hNpgL3ptE9OQjV6c99h/TMwGQMQ\u003d\n\u003dAUUe\n-----END TENDERMINT PRIVATE KEY-----","value":{"msg":[{"type":"cosmos-sdk/MsgSend","value":{"from_address":"zebi19cj9h7yka3r4ln2le75mspyrk2t0h8l82dx5dt","to_address":"zebi1jhjsyvtmytumw236lslc9z5zf7n7gf5au8auj6","amount":[{"denom":"zebi","amount":"117000000"}]}}],"memo":"","fee":{"amount":[{"denom":"zebi","amount":"0"}]}},"pass":"test1234","sender":"account1","base_req":{"from":"account1","chain_id":"zpb-testnet","memo":"","gas_adjustment":"1.5","simulate":false,"gas_prices":[{"denom":"zebi","amount":"1400.0"}],"gas":"37678"}}'

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgSend","value":{"from_address":"zebi19cj9h7yka3r4ln2le75mspyrk2t0h8l82dx5dt","to_address":"zebi1jhjsyvtmytumw236lslc9z5zf7n7gf5au8auj6","amount":[{"denom":"zebi","amount":"117000000"}]}}],"fee":{"amount":[{"denom":"zebi","amount":"52749200"}],"gas":"37678"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A/b/EJCMveJUZ3w2kjvOFiKepd5Ze71TtBIKas4L6Hfp"},"signature":"IEt99E4OMKJov4qGDeD7s0iX2WtIedBcFA6hKwYWXFc5fovahblE1PQZIAryD03b83jnteKmTMUVBRxHnfBWEQ=="}],"memo":""}}

2. For key value post type transaction
curl -XPOST -s http://localhost:1317/kvstore/signTx --data-binary '{"type":"auth/StdTx","armorprivkeys":"-----BEGIN TENDERMINT PRIVATE KEY-----\nkdf: bcrypt\nsalt: FDBFD5FF9BEE248C6B294AB1BE1AD016\n\nRdxaRcb7Gk1+/WzVApEfQ6fiQY02HG4VQlXtz9JMsExA221AREoqdVfUFbAazUKf\nV9M86UZxxvk0ZD+JsJjq3vicPDXdgG7QychAVBQ=\n=c9pv\n-----END TENDERMINT PRIVATE KEY-----","value":{"msg":[{"type":"kvstore/PostKeyValue","value":{"key":"victoria","value":"memorial","sender":"zebi1mdgpxpd2lkvw58nt6k2et3cywmv9kf2d5f9dpp"}}],"fee":{"gas":"12421","amount":[{"denom":"stake","amount": "1"}]},"signatures":null,"memo":""},"base_req":{"from":"enakship1","chain_id":"zebichain","sequence":1,"memo":"","account_number":0,"gas_adjustment":"1.2","simulate":false},"pass":"enakshipriya","sender":"enakship1"}'

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"kvstore/PostKeyValue","value":{"key":"victoria","value":"memorial","sender":"cosmos1d5kuucd7jwmxj94hqccmu5jtp3qad6j0qu829c"}}],"fee":{"amount":[{"denom":"","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"AjND/koT/v41h3tZUDyUnU2uzs+828V8jSbXQIQNAHi7"},"signature":"i6LxOmVfiH5c9XNkLMYK8GCy4YSqsPviSsakPVvpahJz2KAYg5uU04klPfLP+2eoydX/Ph2Hd7AMSjFFHCXw3A=="}],"memo":""}}

3. For delegation transaction
curl -XPOST -s http://localhost:1317/kvstore/signTx --data-binary '{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgDelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":null,"memo":"Sent via sharuk"},"base_req":{"from":"sharuk","chain_id":"zebichain","sequence":4,"memo":"","account_number":3,"gas_adjustment":"1.2","simulate":false},"pass":"sharukm786","sender":"sharuk"}'

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgDelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ah38O6cGjt8gz7ZI6DA9iPPn0LGyp31otg2UD3+WgiHU"},"signature":"1fVRJW0OC9PI17wchHYDofYik6Y7utRTdLH1JqHeT/J1f0Uco6rT4z1xQuZaQs6K/83KNeHiyd4Aa+dUulQI1g=="}],"memo":"Sent via sharuk"}}

4. For undelegation transaction
curl -XPOST -s http://localhost:1317/kvstore/signTx --data-binary '{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgUndelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":null,"memo":"Sent via Sharuk"},"base_req":{"from":"sharuk","chain_id":"zebichain","sequence":5,"memo":"","account_number":3,"gas_adjustment":"1.2","simulate":false},"pass":"sharukm786","sender":"sharuk"}'

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgUndelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ah38O6cGjt8gz7ZI6DA9iPPn0LGyp31otg2UD3+WgiHU"},"signature":"+PI77HuKAlW37HRRj6q3o7jbDUQ3/nrCGw+RWAtJmIUehpXUiTn9rtxKc0tO1A2Yi3jm1oPHzE2SHewxaYaqBg=="}],"memo":"Sent via Sharuk"}}

5. For redelegation transaction
curl -XPOST -s http://localhost:1317/kvstore/signTx --data-binary '{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgBeginRedelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_src_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","validator_dst_address":"zebivaloper1e23tczvtjz8fduywxeuf0rjpr9f69992cgzqhs","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":null,"memo":"Sent via Sharuk"},"base_req":{"from":"sharuk","chain_id":"zebichain","sequence":6,"memo":"","account_number":3,"gas_adjustment":"1.2","simulate":false},"pass":"sharukm786","sender":"sharuk"}'

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgBeginRedelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_src_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","validator_dst_address":"zebivaloper1e23tczvtjz8fduywxeuf0rjpr9f69992cgzqhs","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ah38O6cGjt8gz7ZI6DA9iPPn0LGyp31otg2UD3+WgiHU"},"signature":"ovW99iIyB7RglAgvnXke/PySUVwPUzdtrn9rNRgdma0Q33q5Hk8NBADfc1mEAi+Y0eKp6oT/KThXHGOcTC5kng=="}],"memo":"Sent via Sharuk"}}

6. For create validator transaction (also called as bonding/staking)
Let output of create unsigned transaction be output_create_unsignedtx.
curl -XPOST -s http://localhost:1317/kvstore/signTx --data-binary '<output_create_unsignedtx>,"base_req":{"from":"node1","chain_id":"devchain","sequence":0,"memo":"","account_number":1,"gas_adjustment":"1.2","simulate":false},"pass":"testtest","sender":"node1"}'

Example transaction:- curl -XPOST -s http://localhost:1317/kvstore/signTx --data-binary '{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"description":{"moniker":"nodeJ","identity":"Jayasimha ka node","website":"https://cabzz.in","details":"Testing node - Please do not delegate to this validator"},"commission":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.130000000000000000"},"min_self_delegation":"10","delegator_address":"zebi1da7fsmk7rmcslhu2kfguqa9jfxma0443rhzgtg","validator_address":"zebivaloper1da7fsmk7rmcslhu2kfguqa9jfxma04433f8jmd","pubkey":"zebivalconspub1zcjduepqm76lr6z8j6ajdvlgfequksex309xswlnptxzsk4mkek70pspmwts9c5cv5","value":{"denom":"stake","amount":"250"}}}],"fee":{"amount":[],"gas":"200000"},"signatures":null,"memo":""},"base_req":{"from":"node1","chain_id":"devchain","sequence":30,"memo":"","account_number":0,"gas_adjustment":"1.2","simulate":false},"pass":"testtest","sender":"node1"}'

Output:- {"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"description":{"moniker":"nodeJ","identity":"Jayasimha ka node","website":"https://cabzz.in","details":"Testing node - Please do not delegate to this validator"},"commission":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.130000000000000000"},"min_self_delegation":"10","delegator_address":"zebi1da7fsmk7rmcslhu2kfguqa9jfxma0443rhzgtg","validator_address":"zebivaloper1da7fsmk7rmcslhu2kfguqa9jfxma04433f8jmd","pubkey":"zebivalconspub1zcjduepqm76lr6z8j6ajdvlgfequksex309xswlnptxzsk4mkek70pspmwts9c5cv5","value":{"denom":"stake","amount":"250"}}}],"fee":{"amount":[],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A1+JgwZG/67m6fFze4YfLBRWFC/+HHrGglQIULzI6CaN"},"signature":"ZwXaL4YS/DdExzBkebonZXT5+bOGdSqx4dZOhYKp3G4tYX/I99ppaJzTEcZFy1rZJyKFdPgsKD0GibgbnBysnw=="}],"memo":""}}

7. No signing required for create account

8. For withdraw rewards
curl -XPOST -s http://localhost:1317/kvstore/signTx --data-binary '{"type":"auth/StdTx","armorprivkeys": "-----BEGIN TENDERMINT PRIVATE KEY-----\nkdf: bcypt\nsalt: 44F81DD4EADA834C0A3D923C79BB80BE\n\nWt6+eBpkGGHgdfYw5R0uIR5Pzxrcr9/AdMiwha1QL1ziYMb86Rp950bSn0M3HQTt\n/qUmR1I4ncPzWX3ogCYLfbR5D06QukWuCbFsGcY=\n=0Vjj\n-----END TENDERMINT PRIVATE KEY-----","value":{"msg":[{"type":"cosmos-sdk/MsgWithdrawDelegationReward","value":{"delegator_address":"zebi13yqsaue50jsd2d7ygk0pzdsvwjn3uert9f7nn8","validator_address":"zebivaloper13yqsaue50jsd2d7ygk0pzdsvwjn3uerthhmfrz"}}],"fee":{"amount":[{"denom":"stake","amount":"1"}],"gas":"200000"},"signatures":null,"memo":""},"base_req":{"from":"enakship","chain_id":"zebichain","sequence":252,"memo":"","account_number":0,"gas_adjustment":"1.2","simulate":false},"pass":"enakshipriya","sender":"enakship"}'

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgWithdrawDelegationReward","value":{"delegator_address":"zebi13yqsaue50jsd2d7ygk0pzdsvwjn3uert9f7nn8","validator_address":"zebivaloper13yqsaue50jsd2d7ygk0pzdsvwjn3uerthhmfrz"}}],"fee":{"amount":[{"denom":"stake","amount":"1"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ag6yMVheSDl4KBhpsMPVhhmTpAYX0T6Dnun82fUSJYUG"},"signature":"OBygv37lqVW1TTDXGLSBGa9Xm565Qm3XMe7wKqOToUJYCsvHVniQ32Cu9OazuAgDllGKnZ5yfYUHKab35tkzVw=="}],"memo":""}}


#Rest api command to broadcast a transaction
1. For token transfer transaction
curl --insecure -XPOST -s https://localhost:1317/txs --data-binary '{"tx":{"type":"auth/StdTx","msg":[{"type":"cosmos-sdk/MsgSend","value":{"from_address":"zebi19cj9h7yka3r4ln2le75mspyrk2t0h8l82dx5dt","to_address":"zebi1jhjsyvtmytumw236lslc9z5zf7n7gf5au8auj6","amount":[{"denom":"zebi","amount":"117000000"}]}}],"fee":{"amount":[{"denom":"zebi","amount":"52749200"}],"gas":"37678"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A/b/EJCMveJUZ3w2kjvOFiKepd5Ze71TtBIKas4L6Hfp"},"signature":"IEt99E4OMKJov4qGDeD7s0iX2WtIedBcFA6hKwYWXFc5fovahblE1PQZIAryD03b83jnteKmTMUVBRxHnfBWEQ=="}],"memo":""},"mode":"block"}'

=======>{"height":"16152","txhash":"9AA897B31EE00B7BA983A6348B525F226F906F5D026B802D86E925FDD5C31B68","raw_log":"[{\"msg_index\":0,\"success\":true,\"log\":\"\"}]","logs":[{"msg_index":0,"success":true,"log":""}],"gas_wanted":"37678","gas_used":"28702","tags":[{"key":"action","value":"send"},{"key":"category","value":"bank"},{"key":"sender","value":"zebi19cj9h7yka3r4ln2le75mspyrk2t0h8l82dx5dt"},{"key":"recipient","value":"zebi1jhjsyvtmytumw236lslc9z5zf7n7gf5au8auj6"}]}

2. For key value post type transaction
curl -XPOST -s http://localhost:1317/txs --data-binary '{"tx":{"type":"auth/StdTx","msg":[{"type":"kvstore/PostKeyValue","value":{"key":"victoria","value":"memorial","sender":"cosmos1d5kuucd7jwmxj94hqccmu5jtp3qad6j0qu829c"}}],"fee":{"amount":[{"denom":"","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"AjND/koT/v41h3tZUDyUnU2uzs+828V8jSbXQIQNAHi7"},"signature":"DlS59L+z07FogV93oR+DroazniTK09pArq5BkTol2qNHf3XvP89AGs03OORWJrzoEZmmIzcYiRuvu659S/jqLg=="}],"memo":""},"mode":"block"}'

=======>{"height":"13583","txhash":"4011CC517BA1693971D6E74608F50E010C2B6344193A2E4B82EBD3CDED0B690C","raw_log":"[{\"msg_index\":0,\"success\":true,\"log\":\"\"}]","logs":[{"msg_index":0,"success":true,"log":""}],"gas_wanted":"200000","gas_used":"12785","tags":[{"key":"action","value":"post_kv"}]}

3. For delegate transaction
curl -XPOST -s http://localhost:1317/txs --data-binary '{"tx":{"type":"auth/StdTx","msg":[{"type":"cosmos-sdk/MsgDelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ah38O6cGjt8gz7ZI6DA9iPPn0LGyp31otg2UD3+WgiHU"},"signature":"1fVRJW0OC9PI17wchHYDofYik6Y7utRTdLH1JqHeT/J1f0Uco6rT4z1xQuZaQs6K/83KNeHiyd4Aa+dUulQI1g=="}],"memo":"Sent via sharuk"},"mode":"block"}'

=======>{"height":"99503","txhash":"910F620CACA4FB44C6882CF10BB1DD6470C4CE3E225821AB5984B752D8C20F96","raw_log":"[{\"msg_index\":0,\"success\":true,\"log\":\"\"}]","logs":[{"msg_index":0,"success":true,"log":""}],"gas_wanted":"200000","gas_used":"80857","tags":[{"key":"action","value":"delegate"},{"key":"category","value":"staking"},{"key":"sender","value":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy"},{"key":"destination-validator","value":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39"}]}

4. For undelegate transaction
curl -XPOST -s http://localhost:1317/txs --data-binary '{"tx":{"type":"auth/StdTx","msg":[{"type":"cosmos-sdk/MsgUndelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ah38O6cGjt8gz7ZI6DA9iPPn0LGyp31otg2UD3+WgiHU"},"signature":"+PI77HuKAlW37HRRj6q3o7jbDUQ3/nrCGw+RWAtJmIUehpXUiTn9rtxKc0tO1A2Yi3jm1oPHzE2SHewxaYaqBg=="}],"memo":"Sent via Sharuk"},"mode":"block"}'

=======>{"height":"111066","txhash":"D328F8F8B6EF7209975C3A05DEA2D33208D5072FC37BC1BF29BC245566AAF522","data":"0C08F592D5E9051087ABC5B701","raw_log":"[{\"msg_index\":0,\"success\":true,\"log\":\"\"}]","logs":[{"msg_index":0,"success":true,"log":""}],"gas_wanted":"200000","gas_used":"83273","tags":[{"key":"action","value":"begin_unbonding"},{"key":"category","value":"staking"},{"key":"sender","value":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy"},{"key":"source-validator","value":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39"},{"key":"end-time","value":"2019-07-22T05:28:21Z"}]}

5. For redelegate transaction
curl -XPOST -s http://localhost:1317/txs --data-binary '{"tx":{"type":"auth/StdTx","msg":[{"type":"cosmos-sdk/MsgBeginRedelegate","value":{"delegator_address":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy","validator_src_address":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39","validator_dst_address":"zebivaloper1e23tczvtjz8fduywxeuf0rjpr9f69992cgzqhs","amount":{"denom":"stake","amount":"5"}}}],"fee":{"amount":[{"denom":"stake","amount":"0"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ah38O6cGjt8gz7ZI6DA9iPPn0LGyp31otg2UD3+WgiHU"},"signature":"ovW99iIyB7RglAgvnXke/PySUVwPUzdtrn9rNRgdma0Q33q5Hk8NBADfc1mEAi+Y0eKp6oT/KThXHGOcTC5kng=="}],"memo":"Sent via Sharuk"},"mode":"block"}'

=======>{"height":"114218","txhash":"80CC801EB8EA6EF9C0BE01B02A760011C413069FCDF5E91F7E606CF0AAD16A24","data":"0C08D195D6E90510D4EC8BC303","raw_log":"[{\"msg_index\":0,\"success\":true,\"log\":\"\"}]","logs":[{"msg_index":0,"success":true,"log":""}],"gas_wanted":"200000","gas_used":"134196","tags":[{"key":"action","value":"begin_redelegate"},{"key":"category","value":"staking"},{"key":"sender","value":"zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy"},{"key":"source-validator","value":"zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39"},{"key":"destination-validator","value":"zebivaloper1e23tczvtjz8fduywxeuf0rjpr9f69992cgzqhs"},{"key":"end-time","value":"2019-07-22T10:07:13Z"}]}

6. For create validator transaction (also called as bonding/staking)
Remove 'value' field from output of create signed transaction and let it be output_create_signedtx.
curl -XPOST -s http://localhost:1317/txs --data-binary '{"tx":<output_create_signedtx>,"mode":"block"}

Example transaction:- curl -XPOST -s http://localhost:1317/txs --data-binary '{"tx":{"type":"auth/StdTx","msg":[{"type":"cosmos-sdk/MsgCreateValidator","value":{"description":{"moniker":"nodeJ","identity":"Jayasimha ka node","website":"https://cabzz.in","details":"Testing node - Please do not delegate to this validator"},"commission":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.130000000000000000"},"min_self_delegation":"10","delegator_address":"zebi1da7fsmk7rmcslhu2kfguqa9jfxma0443rhzgtg","validator_address":"zebivaloper1da7fsmk7rmcslhu2kfguqa9jfxma04433f8jmd","pubkey":"zebivalconspub1zcjduepqm76lr6z8j6ajdvlgfequksex309xswlnptxzsk4mkek70pspmwts9c5cv5","value":{"denom":"stake","amount":"250"}}}],"fee":{"amount":[],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A1+JgwZG/67m6fFze4YfLBRWFC/+HHrGglQIULzI6CaN"},"signature":"ZwXaL4YS/DdExzBkebonZXT5+bOGdSqx4dZOhYKp3G4tYX/I99ppaJzTEcZFy1rZJyKFdPgsKD0GibgbnBysnw=="}],"memo":""},"mode":"block"}'

7. No broadcasting required for create account

8. For withdraw rewards
curl -XPOST -s http://localhost:1317/txs --data-binary '{"tx":{"type":"auth/StdTx","msg":[{"type":"cosmos-sdk/MsgWithdrawDelegationReward","value":{"delegator_address":"zebi13yqsaue50jsd2d7ygk0pzdsvwjn3uert9f7nn8","validator_address":"zebivaloper13yqsaue50jsd2d7ygk0pzdsvwjn3uerthhmfrz"}}],"fee":{"amount":[{"denom":"stake","amount":"1"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ag6yMVheSDl4KBhpsMPVhhmTpAYX0T6Dnun82fUSJYUG"},"signature":"OBygv37lqVW1TTDXGLSBGa9Xm565Qm3XMe7wKqOToUJYCsvHVniQ32Cu9OazuAgDllGKnZ5yfYUHKab35tkzVw=="}],"memo":""},"mode":"block"}'

=======>{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgWithdrawDelegationReward","value":{"delegator_address":"zebi13yqsaue50jsd2d7ygk0pzdsvwjn3uert9f7nn8","validator_address":"zebivaloper13yqsaue50jsd2d7ygk0pzdsvwjn3uerthhmfrz"}}],"fee":{"amount":[{"denom":"stake","amount":"1"}],"gas":"200000"},"signatures":[{"pub_key":{"type":"tendermint/PubKeySecp256k1","value":"Ag6yMVheSDl4KBhpsMPVhhmTpAYX0T6Dnun82fUSJYUG"},"signature":"OBygv37lqVW1TTDXGLSBGa9Xm565Qm3XMe7wKqOToUJYCsvHVniQ32Cu9OazuAgDllGKnZ5yfYUHKab35tkzVw=="}],"memo":""}}

#Rest api command to search a key on blockchain
curl -s http://localhost:1317/kvstore/key/<key>

#Rest api command to search a data on blockchain
curl -s http://localhost:1317/kvstore/data/<data>

#Rest api command to query a transaction hash on blockchain
curl -s http://localhost:1317/txs/<hash>

#Rest api command to query rewards
1. Delegator rewards
curl -X GET "http://localhost:1317/distribution/delegators/zebi1qj857c5a7me5ffr5zcjd3qm94k7tnd2y049ppq/rewards" -H "accept: application/json"

2. Validator rewards
curl -X GET "http://localhost:1317/distribution/validators/zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39/rewards" -H "accept: application/json"

#Rest api command to query liquid balances
curl -X GET "http://localhost:1317/bank/balances/zebi1e6y6yywuyevy8mwwpajfmqzd8dyt7prxm2ed38" -H "accept: application/json"

#Rest api command to query account details
curl -X GET "http://localhost:1317/auth/accounts/zebi1e6y6yywuyevy8mwwpajfmqzd8dyt7prxm2ed38" -H "accept: application/json"

#Rest api command to query validators
curl -X GET "http://localhost:1317/staking/validators" -H "accept: application/json"

#Rest api command to query my delegations
1. To get all delegations from a validator
curl -X GET "http://localhost:1317/staking/validators/zebivaloper1qj857c5a7me5ffr5zcjd3qm94k7tnd2yatqm39/delegations" -H "accept: application/json"

2. To get all delegations from a delegator
curl -X GET "http://localhost:1317/staking/delegators/zebi1qj857c5a7me5ffr5zcjd3qm94k7tnd2y049ppq/delegations" -H "accept: application/json"

#Rest api command to query all transactions from a particular account
curl -X GET "http://localhost:1317/txs?sender=zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy" -H "accept: application/json"
curl -X GET "http://localhost:1317/txs?recipient=zebi16jmvrcwjyma28xts59f3ryz3xx8hm740mml0qy" -H "accept: application/json"

---------------------------------------------------------------------------------------------
#Commands to add a validator node. Should be done only after adding the node as a peer node.
---------------------------------------------------------------------------------------------
zebicli tx staking create-validator \
  --amount=100000zebitoken \
  --pubkey=$(zebid tendermint show-validator) \
  --moniker="node2" \
  --chain-id=keychain \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --gas-prices="0.025zebitoken" \
  --from=abhishek

zebid gentx \
  --amount 3zebitoken \
  --commission-rate 0.10 \
  --commission-max-rate 0.20 \
  --commission-max-change-rate 0.01 \
  --pubkey $(zebid tendermint show-validator) \
  --name abhishek

--------------------------------------------------------------------------------------------------------------------
#Some flags associated with client endpoints [Use "zebicli [command] --help" for more information about a command.]
--------------------------------------------------------------------------------------------------------------------
zebicli -h
kvstore Client

Usage:
  zebicli [command]

Available Commands:
  status      Query remote node for status
  config      Create or query a Gaia CLI configuration file
  query       Querying subcommands
  tx          Transactions subcommands
  rest-server Start LCD (light-client daemon), a local REST server
  keys        Add or view local private keys
  version     Print the app version
  stake       Stake and validation subcommands
  gov         Governance and voting subcommands
  help        Help about any command

Flags:
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
  -h, --help              help for zebicli
      --home string       directory for config and data (default "/home/cosmos//.zebicli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors

-----------------------------------------------------------------------------

1 zebicli tx -h
Transactions subcommands

Usage:
  zebicli tx [command]

Available Commands:
  send        Create and sign a send tx
  sign        Sign transactions generated offline
  broadcast   Broadcast transactions generated offline
  kvstore    Keystore transactions subcommands

Flags:
  -h, --help   help for tx

Global Flags:
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/cosmos//.zebicli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors


1.1 zebicli tx kvstore -h
Keystore transactions subcommands

Usage:
  zebicli tx kvstore [command]

Available Commands:
  post-data      post data of any string length
  post-key-value post key value pair of any string length

Flags:
  -h, --help   help for kvstore

Global Flags:
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/cosmos//.zebicli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors

----------------------------------------------------------------------------

2 zebicli query -h
Querying subcommands

Usage:
  zebicli query [command]

Aliases:
  query, q

Available Commands:
  tendermint-validator-set Get the full tendermint validator set at given height
  block                    Get verified data for a the block at given height
  txs                      Search for all transactions that match the given tags.
  tx                       Matches this txhash over all committed blocks
  account                  Query account balance
  kvstore                 Querying commands for the kvstore module

Flags:
  -h, --help   help for query

Global Flags:
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/cosmos//.zebicli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors


2.1  zebicli query kvstore -h
Querying commands for the kvstore module

Usage:
  zebicli query kvstore [command]

Available Commands:
  get-data      gets data existing on blockchain
  get-key-value gets value existing on blockchain

Flags:
  -h, --help   help for kvstore

Global Flags:
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/cosmos//.zebicli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors

--------------------------------------------------------------------------------

3 zebicli tx stake -h     || zebicli query stake -h 
Stake and validation subcommands

Usage:
  zebicli tx stake [command]

Available Commands:
  validator             Query a validator
  validators            Query for all validators
  delegation            Query a delegation based on address and validator address
  delegations           Query all delegations made by one delegator
  unbonding-delegation  Query an unbonding-delegation record based on delegator and validator address
  unbonding-delegations Query all unbonding-delegations records for one delegator
  redelegation          Query a redelegation record based on delegator and a source and destination validator address
  redelegations         Query all redelegations records for one delegator
  signing-info          Query a validator's signing information
  create-validator      create new validator initialized with a self-delegation to it
  edit-validator        edit an existing validator account
  delegate              delegate liquid tokens to a validator
  unbond                unbond shares from a validator
  redelegate            redelegate illiquid tokens from one validator to another
  unjail                unjail validator previously jailed for downtime

Flags:
  -h, --help   help for stake

Global Flags:
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/cosmos//.zebicli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors

-------------------------------------------------------------------------------------------------

4 zebicli tx gov -h    ||     zebicli query gov -h
Governance and voting subcommands

Usage:
  zebicli tx gov [command]

Available Commands:
  proposal        Query details of a single proposal
  vote            Query details of a single vote
  votes           Query votes on a proposal
  proposals       Query proposals with optional filters
  submit-proposal Submit a proposal along with an initial deposit
  deposit         Deposit tokens for activing proposal
  vote            Vote for an active proposal, options: yes/no/no_with_veto/abstain

Flags:
  -h, --help   help for gov

Global Flags:
      --chain-id string   Chain ID of tendermint node
  -e, --encoding string   Binary encoding (hex|b64|btc) (default "hex")
      --home string       directory for config and data (default "/home/cosmos//.zebicli")
  -o, --output string     Output format (text|json) (default "text")
      --trace             print out full stack trace on errors


--------------------------------------------------------------------------------------------------------------------
#Some flags associated with client endpoints [Use "zebid [command] --help" for more information about a command.]
--------------------------------------------------------------------------------------------------------------------

zebid -h
kvstore App Daemon (server)

Usage:
  zebid [command]

Available Commands:
  init                Initialize genesis config, priv-validator file, and p2p-node file
  add-genesis-account Adds an account to the genesis file
  start               Run the full node
  unsafe-reset-all    Resets the blockchain database, removes address book files, and resets priv_validator.json to the genesis state

  tendermint          Tendermint subcommands
  export              Export state to JSON

  version             Print the app version
  help                Help about any command

Flags:
  -h, --help               help for zebid
      --home string        directory for config and data (default "/home/cosmos//.zebid")
      --log_level string   Log level (default "main:info,state:info,*:error")
      --trace              print out full stack trace on errors