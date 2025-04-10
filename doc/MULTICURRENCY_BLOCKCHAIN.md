# Multicurrency Blockchain
A wallet application(WIP: https://github.com/pathfinder177/crypto_wallet) \
will work with two cryptocurrencies' coins (referred to as "currencies") and uses blockchain as the server side.
Thus blockchain should be multicurrency.

---

## Blockchain (Server Side)

Used as an **irreversible digital ledger** to:
- Make transactions (TXs)
- Store transaction history  
- Provide transaction history

---

## Currencies

### Questions:
<!-- - How can they be integrated into existing blockchain? -->
<!-- - Is it possible to mint/forge them directly on the blockchain?
  - If so, they might not need to be created elsewhere. -->
Coins amount will be set initially. 
Model of issuance is simplified and mocks fixed supply model(BTC)

Chosen one
Option 1: Bring multicurrency to a blockchain
    currency can be created directly on the blockchain and
    be native(coins) or non-native(tokens) to a blockchain
    Both currencies will be native(coins) and independent of each other

<!-- Complicated
Option 2: currency created outside of blockchain and transferred over the network
    Need bridge or something like that to integrate two blockchains
    My blockchain should be public to interact with other -->

## Architecture

### Bring multicurrency to the blockchain
Converting a single to a multi-currency blockchain/ledger requires:
- Multi-Coins/Tokens Ledger Design
    Coins: https://eprint.iacr.org/2020/895.pdf
        Formally: In paper a ledger is assumed to be a list of transactions
    Tokens: https://xrpl.org/docs/concepts/ledgers

    Key pairs
    To send and receive coins, a user needs at least one key pair for each currency.
    Each key pair proves ownership of one wallet address for a specific currency.

- Multi-currency transaction processing
    Same transactions types remain

    Generalized notion of Value
    allows the atomic transfer of multiple assets and currencies 
    in a single account-based transaction
- Token issuance
    - *(Optionally) add creator of the currency to specify monetary policy
        - Add new TX type
        - Make supply fixed

- Modular architecture
    *(Optionally) implements separate blockchain for each currency(sidechains)

- Consensus and Validation
    Adjustments to the consensus mechanism may be necessary if the introduction of multiple currencies impacts transaction throughput or security assumptions. In some designs, you might separate the validation of different asset types or incorporate additional verification steps for token-specific transactions.

- Multi-Currency Wallet Integration


#### Multi-token ledger design
Ledger logic are contains in files:
block*, utxo_set

#### Multi-currency transaction processing
Transaction logic contains in files:
block*, transaction*, utxo_set

#### How TX is processed now(without server part)
There are two types of TX:
1. Coinbase TX is used when:
    genesis block is created(the only TX there)
    new block is created and coinbase TX is the first in this block.
2. UTXO TX is used when:
    Roughly owner of one wallet transfers coins to other wallet
    Technically, coins should be unlocked using current owner data and locked using next owner data

## Projecting
### Bring multicurrency to the blockchain
Handles two types of transactions:
##### 1. Coinbase
Given blockchain is created
When genesis block is mined
Then miner get subsidy in two currencies

Test: as blockchain is created there is correct subsidy got to miner
Check: createBlockchain, newBlockchain, createWallet, getBalance

REINDEX -> FindUTXO
Something to differ currencies to work with
Should 2 currencies be in one TX?
    Probably not(except for coinbase)

Iterate over values as REINDEX need it
1) Iterate over currencies in method
2) Method take currency as input to not to iterate over all
3) Can method finds UTXOs for all currencies?
4) If UTXO has mixed structure?
5) Reindex by bucket

Break down currencies into buckets only for UTXO?

State: two buckets in DB:
badgercoin_chainstate and catfish_chainstate

##### Break down at TX level:

Write each type of TXs to corr.buckets
    Genesis block includes 2 TXs
    Other block includes from 2 to 4 TXs:
        One for coinbase and one for currency

##### Break down at BLC level

##### 2. UTXO
Given UTXO TX is started
When amount is calculated
Then different currencies are handled correctly

Put to different buckets should be shaped as transactions
FindSpendableOutputs: what is sent and what do we search for

Cases:
    SEND: one currency has x UTXO number and other has y UTXO number and txID

#### Questions:
    Currency:
    - What will the blockchain address look like?
        Won't be changed
    - What are the formats for the keys necessary to create signatures for transactions?
        Won't be changed if possible
    - Value is int and becomes map[currency, value]
        Each currency should have own identity element such that (value-identityElem=0) and at start currency value is 0

    Blockchain:
    - One blockchain for both currencies or two?
        One
    - If one there are no pros and mess in UTXO, blocks are bigger, etc
        the main pros is to reuse the same type of transactions

#### TODO
- Errors hierarchy: Instead of print or fatal use corresponding errors
