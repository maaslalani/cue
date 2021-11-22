# Cue

Cue cards in your terminal.

<p align="center">
  <img src="./cue.gif?raw=true" alt="Cue cards" />
</p>

### Usage

Make some colon (`:`) separated cue cards in a plain text file.

```
Press Spacebar: You can flip a card by pressing spacebar. Navigate with arrow keys or hl.
Usage: Give `cue` a list of terms and definitions together separated by a colon.
Cue: Study for your next exam on the command-line with **cue < cards.txt**.
Markdown Support: *You* can even write in **markdown**.
```

Then use `cue` to interact with your cue cards.

```
cue < cards.txt
```

### Installation

```
go install github.com/maaslalani/cue@latest
```

### Navigation

Switch cue cards with: <kbd>Right</kbd>, <kbd>Left</kbd>, <kbd>h</kbd>, <kbd>l</kbd>.

Use <kbd>Space</kbd> to toggle between the term and the definition.
