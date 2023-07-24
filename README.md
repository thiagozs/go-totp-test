# go-TOPT-Test - Proof of Concept

"Welcome to the magical world of **TOTP**, or **Time-Based One-Time** Passwords! Picture this: you're a wizard, and your wand is your smartphone. Now, every wizard needs a spell to open secret doors, right? But you wouldn't want the same spell to work every time, because what if a goblin overhears you?

That's where **TOTP** comes in. It's like a spell that changes every 30 seconds! Only you, the wizard, can cast it because only you have your enchanted smartphone. It's a powerful charm that keeps your secrets safe from prying goblin eyes. So, next time you're asked for a password, just wave your smartphone, say your TOTP spell, and watch those secret doors swing open!"

## Example

```golang
func main() {
	randomSecret := gotp.RandomSecret(16)
	fmt.Println("Random secret:", randomSecret)
	generateTOTPWithSecret(randomSecret)
	testOTPVerify(randomSecret)
}
```

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-totp-test/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE) file.

**2023**, thiagozs.
