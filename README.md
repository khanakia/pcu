## pcu (php-composer-update) 
Updates your composer.json dependencies to their latest versions.

## Installation
```go
go install github.com/khanakia/pcu@main
```

### Using Homebrew
```sh
brew tap khanakia/pcu
brew install pcu
pcu --help
```

## Usage

### Check the `./composer.json` file in current working directory and print the latest version information
```sh
pcu check
```

### Update the same `composer.json` file with latest version
```sh
pcu check -u
```


### Specify custom input and output path
```sh
pcu check -u --file=./test/composer.json --out=./out.json
```


## Sample Output
```js
laravel/jetstream                                   ^4.0            ^4.0
mll-lab/laravel-graphiql                            ^3.0            ^3.1
mll-lab/graphql-php-scalars                         ^6.1            ^6.2
nuwave/lighthouse                                   ^6.8            ^6.22
laravel/sanctum                                     ^3.2            ^3.3
milon/barcode                                       ^10.0           ^10.0
filament/filament                                   ^3.0-stable     ^3.0
guzzlehttp/guzzle                                   ^7.2            ^7.8
hidehalo/nanoid-php                                 ^1.1            ^1.1
barryvdh/laravel-dompdf                             ^2.0            ^2.0
egulias/email-validator                             ^4.0            ^4.0
laravel/framework                                   ^10.0           ^10.30
doctrine/dbal                                       ^3.6            ^3.7
commerceguys/addressing                             ^2.0            ^2.0
timehunter/laravel-google-recaptcha-v3              ^2.5            ^2.5
laravel/tinker                                      ^2.8            ^2.8
plank/laravel-mediable                              ^5.9            ^5.9
```


## Video

[![PCU Demo](http://img.youtube.com/vi/1mbCzsStu_8/0.jpg)](http://www.youtube.com/watch?v=1mbCzsStu_8 "PCU Demo")

