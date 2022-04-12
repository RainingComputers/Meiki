<center>
<img src="./meiki.svg" width="50%">

The ⚡ lightning fast ⚡ notes editor
</center>

# Meiki Notes

Meiki is a markdown notes editor built with performance and
simplicity in mind. The name "MEIKI" was inspired from the Japanese word [銘記](https://jisho.org/word/%E9%8A%98%E8%A8%98) (めいき), which means to "keep in mind" or "take note of" or "remember​".

Meiki's frontend is written in svelte and the backend is written in golang. Meiki uses mongoDB to store the notes, user credentials and session tokens.

## Running in local

*Start mongoDB*

```
cd meiki_server
docker-compose -f test_services.yml up
```

*Start server*

```
cd meiki_server
go get
go run main.go
```

*Start UI*

```
cd meiki_ui
npm i
npm run dev
```

## Running tests

*Start mongoDB*

```
cd meiki_server
docker-compose -f test_services.yml up
```

*Running server tests*
```
cd meiki_server
go get
go test ./...
```

*Running UI tests*
```
cd meiki_ui
npm i
npx cypress open
```
