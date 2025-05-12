# Kbot demo telegram bot

This repo is for managing simple telegram bot on go `t.me/belrestro_learing_bot`

## install

```bash
go get
```

## setup

you'll need environment variable `TELE_TOKEN`

this variable should contain 

to securely setup it into environment use

```bash
read -s TELE_TOKEN
# use ctrl+v to paste and then
export TELE_TOKEN
```

## build 

to build the bot use

```bash
go build -ldflags "-X=kbot/cmd.appVersion=1.0.1" -o bin/kbot
```

## run 

to run the bot use

```bash
./bin/kbot start
```

## commands

```
/greet - greet the user
/finish - farewells the user
```