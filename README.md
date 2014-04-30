# Yegg ![Analytics](https://ga-beacon.appspot.com/UA-34529482-6/yegg/readme?pixel)

<img src="http://i.imgur.com/rPVZlqa.png" alt="Yegg Icon" align="right">
Yegg is a proof-of-concept password retriever for the ESUSD network.

Yegg takes advantage of the lack of login rate-limiting to brute force a
specified user's password. It is currently relatively unoptimized and can
likely be made many magnitudes faster

Yegg was created strictly for educational purposes and is provided with
absolutely no warranty. The author is not responsible for any misuse.

## Installation

    $ go get github.com/zachlatta/yegg

## Usage

    $ yegg -user [user] -url [form submit url]

Run with `-help` for full usage instructions.

## LICENSE

See [LICENSE](LICENSE).
