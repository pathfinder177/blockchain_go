# Wallet Server
A wallet client-server application that works with two cryptocurrencies' coins (referred to as "currencies") and uses [blockchain](https://github.com/pathfinder177/blockchain_go) as the server side.

For a client side please refer to [crypto_wallet](https://github.com/pathfinder177/crypto_wallet)

---

## Wallet (Server Side)

### Features:
1. **Create Wallet**
   - User get to page for new user -> register -> wallet is created -> user get to the wallet main page
   *- Each pocket has **X value** of currency's coins (e.g., airdrop or distribution)
   - Optionally, users can create additional pockets for each of currencies.
        - How to implement it? Probably pocket is a virtual entity as representation of blockchain is unchanged
3. **Delete Wallet**
   - Optionally, user can delete created pockets, but at least one per currency must remain.
   - Otherwise user can delete wallet. Tokens are "burned", wallet is deleted
4. **Send Coins**
   - Incorrect transactions should be reverted to prevent loss of funds
5. **Receive Coins**
6. **Get Data**
   - Balance for currencies
   - Transaction history for each currency over a specific period.

---

## User Story (Test Scenario)

### **User 1**
1. Start the app.  
2. Register.  
3. Create a wallet.  
4. Get balance of each pocket:  
   - Each pocket contains **X value** of currency's coins.  
5. Get transaction history:  
   - It's empty.  

### **User 2**
6. Start the app.  
7. Register.  
8. Create a wallet.  
9. Get balance of each pocket:  
   - Each pocket contains **X value** of currency's coins.  
10. Get transaction history:  
    - It's empty.  
11. Send **1 unit** of `currency's coins_1` to User 1.  
12. Get balance of pocket with sent currency's coins:  
    - `Balance = previous balance - 1`  
13. Get transaction history:  
    - One transaction (sent to User 1).  

### **User 1**
14. Get balance of pocket with received currency's coins:  
    - `Balance = previous balance + 1`  
15. Get transaction history:  
    - One received transaction.  
16. Send **1 unit** of `currency's coins_2` to User 2.  
17. Get balance of pocket with sent currency's coins:  
    - `Balance = previous balance - 1`  
18. Get transaction history:  
    - One sent transaction.  
    - One received transaction.  

### **User 2**
19. Get balance of pocket with received currency's coins:  
    - `Balance = previous balance + 1`  
20. Get transaction history:  
    - One sent transaction.  
    - One received transaction.  

### **User 1 & User 2**
21. Delete their wallets.
    - **Users are deleted from DB**
    - **Coins are burned**

## Architecture
The approach is closer to service-oriented architecture at the system level.
And clean architecture at the code level. FIXME(make sure it is done or changed)

Wallet server is a middleware to handle and proceed with requests from wallet client.
Wallet server interacts with blockchain over:
    - blockchain cli(kind of blockchain api, simplification here)
    - wallet_node(get some state of the chain, e.g. blocks)

### Addresses

Wallet server: localhost:3003
Wallet client: localhost:3004

### Handlers:
- A good idea is to cache user data for GET requests to not to overload blockchain
    and append/change data after send/receive actions in cache as any tx is immutable.
    - Here I omit it

/send(bc)
/transactions(all txs history for the wallet for period(7d default))(bc)
/currency_transactions(bc)(txs history for the currency for period(7d default))(bc)
/delete_wallet(DB and blc)

## Projecting
1. All connections are handled in apart goroutine
2. Graceful shutdown

### Server side HTTP Handlers & TCP transport
#### /get_history(all txs history for the wallet for period(7d default))(bc)
Given user click get_history
When user submit time period
Then user is redirected to page with wallet transactions history

User -> GET / and button Get Transactions History
*App -> asks for period
*User -> POST(submit period)
App -> search block by block from the last where in order (*timestamp if period submitted*):
    (SEND) tx input: pkhwallet | output: pkhother (pkhwallet does not count)
    (RECEIVE)tx input: pkhother | output: pkhwallet
    
    1744365364 Friday, April 11, 2025 4:56:04 PM
    1744364329 Friday, April 11, 2025 4:38:49 PM
    1744364314 Friday, April 11, 2025 4:38:34 PM
    1744363590 Friday, April 11, 2025 4:26:30 PM

    for coinbase tx check only output
    represents for user as:
    BlockTimestamp Send/Receive Currency Amount WSender WReceiver

App -> render template and redirects user to /transactions

#### /get_history_for_currency(bc)(txs history for the currency for period(7d default))(bc)
Given user click get_history for currency
When user submit time period and currency
Then user is redirected to page with transactions history for currency

Currency should be chosen by user: radiobutton(hardcoded)

#### /send(blc)
Given user click send
When user submit amount, currency, receiver
Then currency is sent and confirmation is shown(e.g. you sent 10 badgercoin to address)

#### /delete_wallet(DB and blc)
Given user click delete wallet
When user confirm
Then user is deleted from user's database and tokens from wallet are burned and wallet is deleted from wallet file

### Clean up the code
### Format the code

## Q&A
- Should wallet serverside be in blockchain repo? - yes
    - Pros: 
        - simple to test, all code in one place
    - Cons: 
        - no concerns separations(ok for pet project)
        - if not then import wallet code in this repo(so wallet should be a package) as only web part need to be done
