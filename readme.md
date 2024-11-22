# Gateway
This repository showcase implemetation of Gateway handling multiple services on GOLang.

In this showcase, gateway will receive request for 3 services (consumerloan, cashloan, and risk)

##  Path
### 1. consumerloan
- /gateway/consumerloan

Method: GET
Returning Status of the service with cookies

- /gateway/consumerloan/search 

Method: GET
Required: Cookies with SessionId Key
Returning Status of the service with cookies

### 2. cashloan
- /gateway/cashloan

Method: GET
Returning Status of the service with cookies

### 3. risk
- /gateway/risk

Method: GET
Returning Status of the service with cookies
