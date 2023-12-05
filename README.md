
# Hangman-Classic

This project is the Hangman-Classic ask on Ytrack


## Features

- ASCII Art
- Start And Stop
- Impossible to retry letter or word
- Add different color into the output


## Deployment

IMPORTANT ! Before execute any command, open a terminal since the program folder.

- For start the normal game

```bash
  go run .\cmd\main.go

```
- For start in ASCII Art mode (You can replace name_of_file by : standard / shadow / thinkertoy)
```bash
  go run .\cmd\main.go --letterFile (name_of_file).txt


```
- For start since the save file (IMPORTANT everytime a save file is created his name is save.txt)
```bash
  go run .\cmd\main.go --startWith save.txt


```
- For start since the save file and Ascii Art mode in same time (You can replace name_of_file by : standard / shadow / thinkertoy)
```bash
  go run .\cmd\main.go --startWith save.txt --letterFile (name_of_ascii_file).txt


```

## Appendix

If you want to stop the game and save your progression into a save file just write STOP into the standard input


## Lessons Learned

The biggest difficulty than we encounter into this project is to create it in such a way that the program works on different computers
## Authors

- Carrola Quentin / Petit Melvin

