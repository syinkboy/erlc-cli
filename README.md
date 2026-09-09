# erlc-cli

Small CLI for the ERLC priavte server API. Check who's online, get server status, run commands without openiong Discord.

Built this mostly to get better at Go while making something I'd actually use. This was mainly built to help Collier County Roleplay[https://discord.gg/7bRFgqxVPN]

## Setup

Get a server key from https://erlc.link/sk, then:

git clone https://github.com/Syink/erlc-cli
cd erlc-cli
go mod tidy
export ERLC_SERVER_KEY=your_key_here

## Usage
go run.status
go run.players
go run.command":h Hello"

## License

MIT
