# Wallet Client App
A wallet client application that works with two cryptocurrencies' coins (referred to as "currencies") and uses [blockchain](https://github.com/pathfinder177/blockchain_go) as the server side

---

## Wallet (Client Side)

### Features:
1. **Basic Authentication** (password and optionally token)  
   - Where should credentials be stored?
        - Password: DB
        - Token: basically user-side
2. **Create Wallet User Story**
   - User get to page for new user -> register -> wallet is created -> user get to the wallet main page
   - Each pocket has **X value** of currency's coins (e.g., airdrop or distribution)
   
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
7. ?Wallet smart contract?

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
The approach is closer to service-oriented architecture

Wallet as client-side application implemented as web-server serves static content
Wallet client uses wallet server only to connect to blockchain

Wallet server is a middleware to handle and proceed with requests from wallet client
Wallet server use blc features to do it and interact with blockchain over wallet node only
* Interaction between wallet server and blc uses command of blc cli(kind of simplification)

Wallet server: nodeID 3003
Wallet client: nodeID 3004

### Wallet consists of:
- Frontend to show new user and main wallet page
- Middleware to auth
- Backend to use blockchain functionality

### Storage
To implements basic auth: use postgresql with hashed login and passwords

### Handlers:
Some handlers use storages:
    - DB for database of users(postgresql)
    - BLC for blockchain

- A good idea to cache user data for GET requests to not to overload blockchain
    and append/change data after send/receive actions in cache as any tx is immutable.
    - Here I omit it

/
/auth(DB)
/send(bc)
/get_history(all txs history for the wallet for period(7d default))(bc)
/get_history_for_currency(bc)(txs history for the currency for period(7d default))(bc)
/delete_wallet(DB and blc)

### Wallet UI
- / has:
    - at first stage(without auth):
        - ask for wallet address that created before
        - main page for provided address
- The main page has:
    - address
    - two pockets, one for each currency 
    - buttons for functionality:
        - send
        - get
        - delete wallet

## Projecting
1. All connections are handled in apart goroutine(both)
2. Graceful shutdown(both)
3. Rate limiting(client)

Approach 1: create all web part then integrate it with blockchain
Approach 2: create from inside to outside: start with backend

### Handlers

#### /
Given user goto index
When user submit valid wallet address
Then needed wallet address info is shown to user

User -> GET /
App -> asks for wallet address
User -> POST(submit wallet address)
App -> check wallet address for validity and existence
    if exists: get data from blockchain, fill static page with data, redirects to main page. url=/walletaddress
    else -> 4** error
Mock buttons for now

Use getBalance from blc cli
TODO: make blc cli as functions and cli package is imported(use cobra etc)
To send request to blc, use wallet server that handles requests

#### /auth(DB)
#### /send(bc)
#### /get_history(all txs history for the wallet for period(7d default))(bc)
#### /get_history_for_currency(bc)(txs history for the currency for period(7d default))(bc)
#### /delete_wallet(DB and blc)

## Q&A
- Should wallet be in blockchain repo(monorepo approach)? yes
    - Pros: 
        - simple to test, all code in one place
    - Cons: 
        - no concerns separations(this degree is ok for pet project)
        - if not then need to put wallet code in this repo as only web part need to be done
