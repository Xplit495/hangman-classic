
# Hangman-Classic

Hangman Classic is a console application for playing the classic game of Hangman. Designed with Go, this project provides an interactive experience for guessing hidden words.


## Features

- ASCII Art
- Start And Stop
- Impossible to retry letter or word
- Add different color into the output


## Deployment

IMPORTANT ! Before executing any command, open a terminal from the program folder.

- To start the normal game

```bash
  go run .\cmd\main.go

```
- To start in ASCII Art mode (You can replace name_of_file by : standard / shadow / thinkertoy)
```bash
  go run .\cmd\main.go --letterFile (name_of_file).txt


```
- To start since the save file (IMPORTANT everytime a save file is created his name is save.txt)
```bash
  go run .\cmd\main.go --startWith save.txt


```
- To start since the save file and Ascii Art mode in same time (You can replace name_of_file by : standard / shadow / thinkertoy)
```bash
  go run .\cmd\main.go --startWith save.txt --letterFile (name_of_ascii_file).txt


```

## Appendix

Every time a game is finish a stats file is refresh for add your game history and all your stats (enjoy :)

If you want to stop the game and save your progression into a save file just write STOP into the standard input


## Lessons Learned

The biggest difficulty than we encounter into this project is to create it in such a way that the program works on different computers
## Authors

- Carrola Quentin / Petit Melvin

