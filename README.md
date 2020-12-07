# Professor Chen Discord-bot

>A Go Discord bot for PokéAPI

Prof. Chen is a discord bot for assigning random Pokémon to different members of the server. It offers the possibility to request the allocation of a Pokémon, to display its Pokémon or to display the list of members and their Pokémon.

## Usage

Configure the development environment on your local machine:
```bash
$ git clone <this repo>
$ cd chen-discord-bot
$ go run main.go -t $BOT_TOKEN
```

You can now access the bot in your Discord server.

## Use the command line

To list available commands, either run `make` with no parameters or execute `make help`:

```bash
Usage: make <command>

Commands:
  !poke help               Provides help information on available commands
  !poke claim              Associate a random Pokémon to the current user or replace it
  !poke card               Return data from the Pokémon associated to the current user or a specific Pokémon with ID
  !poke list               Returns the list of users and their associated Pokémon
```
## Author

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/Akecel">
        <img src="https://github.com/Akecel.png" width="150px;"/><br>
        <b>Axel Rayer</b>
      </a>
    </td>
    <td align="center">
      <a href="https://github.com/t-hugo">
        <img src="https://github.com/t-hugo.png" width="150px;"/><br>
        <b>Hugo Tinghino</b>
      </a>
    </td>
  </tr>
</table>



