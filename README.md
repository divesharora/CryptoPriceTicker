# KryptoBackendTask

## Steps to Run
 1> if Go environment is setup: just go with 
 
 go run main.go
 
 2> Else just run  the executable file (Added for both Windows and Linux environments)

## Postman Collection : 
https://documenter.getpostman.com/view/8802432/UzXYqCHV

## Approach taken to solve the alerts trigger

1. Used the binance Webstream API to continuosly fetch realtime prices of all tokens
2. Run them against the untriggered alerts stored in the database and comapre prices 
3. If price is greater than update status to triggered and alert user

## Technologies Used

1. Golang with GoFiber Framework
2. PostgresSQL (localhost)
3. Redis (localhost)

env file included
