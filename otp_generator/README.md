## Introduction
This is the app you can use to generate the OTP using secret string provided by any application.

```
go mod init totp-gen
go get github.com/pquerna/otp/totp

Linux: GOOS=linux GOARCH=amd64 go build -o totp-gen ./main.go  
Windows: GOOS=windows GOARCH=amd64 go build -o totp-gen.exe ./main.go
```
## Usage
```
totp-gen OTP-STRING-GOES-HERE
totp-gen.exe OTP-STRING-GOES-HERE
```