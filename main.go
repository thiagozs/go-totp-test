package main

import (
	"fmt"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/xlzd/gotp"
)

func main() {
	randomSecret := gotp.RandomSecret(16)
	fmt.Println("Random secret:", randomSecret)
	generateTOTPWithSecret(randomSecret)
	testOTPVerify(randomSecret)
}

func generateTOTPWithSecret(randomSecret string) {
	totp := gotp.NewDefaultTOTP(randomSecret)
	fmt.Println("current one-time password is:", totp.Now())

	uri := totp.ProvisioningUri("user@email.com", "myApp")
	fmt.Println(uri)

	qrcode.WriteFile(uri, qrcode.Medium, 256, "qr.png")
}

func testOTPVerify(randomSecret string) {
	totp := gotp.NewDefaultTOTP(randomSecret)
	otpValue := totp.Now()
	fmt.Println("current one-time password is:", otpValue)

	ok := totp.Verify(otpValue, time.Now().Unix())
	fmt.Println("verify OTP success:", ok)
}
