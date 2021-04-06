# Brute Force Hashed Phrase Cracker

Program `pass-crack` checks hash with all possible character combination hashes,
and returns matched phrase.
If all iterations passed without successful match, error is returned.

> This tiny peace of code also illustrates how easily week passwords can be cracked,
> if hashed with old, insecure algorithms.

## Usage

1. Copy example config `cp ./config/config.json.example ./config/config.json`
2. Alter config values by your needs
3. Run program by passing 2 params: algorithm and hash you want to decode.
Example: `go run cmd/pass-crack/main.go sha1 7f550a9f4c44173a37664d938f1355f0f92a47a7`
   
For standalone usage, you may want to build a binary file, and use it independently, 
without required GO installation.

## Supported Algorithms

Program supports 2 common algorithms:

- md5
- sha1

But you can easily extend it, by adding new algorithm wrapper `./pkg/encode/{new_alg}.go`, 
which implements `Encoder` interface.

## Configuration and System Loads

### Config file:

1. `available_chars` Characters, which will be used to generate all possible phrases. 
   By default, all lowercase english alphabet characters provided (26 characters). 
   Uppercase letters and numbers are skipped intentionally, 
   by trying to minimize available iterations number.
   Note: you can use any utf8 characters. Program is not limited to latin char list.
   
2. `max_pass_length` Maximum phrase length to check. 
   For example, world `hello` have 5 characters. Probably, reasonable param 
   value should be from 5 to 8, depending on suffixes/prefixes usage.
   
3. `prefixes/suffixes.enabled` If set to `true`, every generated phrase also 
   is checked with every possible appended/prepended `suffix/prefix`. 
   Note: this can highly extend execution time.
   
4. `prefixes/suffixes.list` Provide possible `suffixes/prefixes` to check, 
   if flag `enabled` is set to `true`.
   Note: by default `prefixes.list` contains all uppercase letters. 
   This also made intentionally, due to common password pattern - first upper case letter.
   
### Phrase Length
While `max_pass_length` indicates max available phrase length, 
prefix and suffix length does not count. For example phrase `Hello123` length is 8.
But when prefixes and suffixes enabled, from a program's perspective:

- `H` - prefix (included in default config)
- `ello` - 4 characters phrase
- `123` - suffix (included in default config)

### System Loads and Execution time

In short, program takes all your CPU while running. This is expected, 
while standard PC, with 4 CPU cores ~2.5 Ghz, scans around 2 millions phrases per second.
The faster you scan, the faster you'll get the answer.
RAM is not so crucial, but with higher load of prefixes/suffixes,
RAM consumption may reach up to 4-5 GB. Mostly not.

Every time you run the program, max execution time calculation is printed.
Every additional character added to config exponentially increases program execution time.
As an example, with default configuration, calculations may look like this:
> 26 (available chars)^6 (max phrase length) * 239 (suffix/prefix overhead) = 73830870464 (max iterations)

So for `sha1` algorithm, for average 4 cores PC, which scans about 2 mills phrases per second,
it can take up to 10 hours to complete scan 
for all possible phrases with prefixes and suffixes.

> Same phrase every time is decoded in different time. 
> It's due to many goroutines running at the same time,
> and you'll never know which will finish first.

## Phrase Decode Strategies

Passwords, that are made by humans, usually are short, 
and made with susceptible patterns, like:
- first upper case letter 
- or first letter is number
- all other letters are lowercase
- phrase ends with some number
- or numbers like `12`, `123`, `321`, etc.
- if password policy requires special char, probably it will be places at the end of the password

Default `config.json.example` reflects all those common patterns (except special char case).
Depending on you PC, expected pattern and other circumstances, 
it may be wise to change config values accordingly.

> For your own security, please, use strong passwords, created by password generators, 
> and save them in password managers. You don't have to memorize them.