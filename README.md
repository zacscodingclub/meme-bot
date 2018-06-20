# CLI Meme Creator

Uses the https://memegen.link api.

## Install

Clone the repo into your `$GOPATH`

For me, it was like this:

```bash
$ cd $GOPATH/src/github.com/zacscodingclub
$ git clone git@github.com:zacscodingclub/meme-bot.git
$ cd meme-bot
# To build an executable
$ go build .
# Dev
# $ go run main.go <args>
$ go run main.go list
$ go run main.go new icanhas "i can haz" "customized meme"
```

## Usage

```bash
$ go build .
$ ./meme-bot help
Creates a customized meme using one of many templates and copies the URL
to your clipboard.  For example:
 $ meme-bot new icanhas "i can haz" "customized meme"
Copies this URL string to your clipboard so you can use it later:
  https://memegen.link/icanhas/i_can_haz/customized_meme.jpg

Usage:
  meme-bot [command]

Available Commands:
  help        Help about any command
  list        List of available templates
  new         Copies a new URL for a custom meme to clipboard

Flags:
  -h, --help     help for meme-bot
  -t, --toggle   Help message for toggle

Use "meme-bot [command] --help" for more information about a command.
```

```bash
$ ./meme-bot new icanhas "i can haz" "customized meme"
```

The string "https://memegen.link/icanhas/i_can_haz/customized_meme.jpg" will be copied to your clipboard.

```bash
$ ./meme-bot new kermit "some top words" "some bottom words"
```

The string "https://memegen.link/kermit/some_top_words/some_bottom_words.jpg" will be copied to your clipboard.
