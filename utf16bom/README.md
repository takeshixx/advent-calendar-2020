# XMAS Karaoke

This service allows users to select a song for karaoke. Lyrics will be sent to the client line by line, with a BOM that indicates which UTF encoding should be used to return the line. Clients get the flag after returning all lines with the proper UTF encoding. Possible encodings:

* UTF8
* UTF16-BE
* UTF16-LE
* UTF32-BE
* UTF32-LE

## Description

```html
Christmas time is all about sining! Sining is fun, especially karaoke. Santa absolutely loooves karaoke!! Today you will get exklusive access to Santa's very own karaoke service that he uses during Christmas time for some elevated fun times. It has a very special selection of songs, the ones he enjoys most. Just classic hits.

You have to get the high score to win this game. Prepare yourself for some hardcore singing!
```

## Solution

`client/client.go` implements an automated solution:

```bash
make client-run
```