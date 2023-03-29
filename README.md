# QuadChecker

> This repository contains the solution to the QuadChecker raid based on the quad functions.

## Instructions

This raid is based on the `quad` functions.

Create a program `quadchecker` that takes a `string` as an argument and displays the name of the matching `quad` and its dimensions.

- If the argument is not a `quad` the program should print `Not a quad function`.

- All answers must end with a newline (`'\n'`).

- If there is more than one `quad` matches, the program must display them all alphabetically ordered and separated by a `||`.


## Installation :question:

> To use this program, clone this repository and navigate to the root directory.

```bash
$ git clone https://learn.01founders.co/git/tmachado/quadchecker.git
$ cd quadchecker
$
```

## Usage

```bash
$ ls -l
-rw-r--r-- 1 student student  nov 23 14:30 main.go
-rwxr-xr-x 1 student student  nov 23 19:18 quadchecker
-rwxr-xr-x 1 student student  nov 23 19:50 quadA
-rwxr-xr-x 1 student student  nov 23 19:50 quadB
-rwxr-xr-x 1 student student  nov 23 19:50 quadC
-rwxr-xr-x 1 student student  nov 23 19:50 quadD
-rwxr-xr-x 1 student student  nov 23 19:50 quadE
```

> To use the QuadChecker program, run the relevant quad function executable and pipe the output to the quadchecker program. For example :

```bash
$ ./quadA 3 3 | go run .
```

> The program will display the name of the matching quad and its dimensions :

```css
[quadA] [3] [3]
$
```

> If the argument is not a quad function, the program will print Not a quad function :

```bash
$ echo 0 0 | go run .
Not a quad function
$
```

> If there is more than one quad match, the program will display them all alphabetically ordered and separated by a || :

```bash
$ ./quadC 1 1 | go run .
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
$
```

> If it's `quadA` :

```bash
$ ./quadA 3 3 | go run .
[quadA] [3] [3]
$
```

> If it's `quadC 1 1` :

```bash
$ ./quadC 1 1
A
$ ./quadD 1 1
A
$ ./quadE 1 1
A
$ ./quadC 1 1 | go run .
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
$
```

> If it's `quadC 1 2` :

```bash
$ ./quadE 1 2
A
C
$ ./quadC 1 2
A
C
$ ./quadE 1 2 | go run .
[quadC] [1] [2] || [quadE] [1] [2]
$
```

> If it's not a quad function :

```bash
$ echo 0 0 | go run .
Not a quad function
$ echo -n "o--o"$'\n'"|"$'\n'"o"
o--o
|
$ echo -n "o--o"$'\n'"|"$'\n'"o" | go run .
Not a quad function
$
```

## Contributing

> If you find a bug or would like to propose a change, please submit an issue or a pull request.

## Credits

> This program was created by :point_right: tmachado and the following teammates :

:point_right: jsmith

:point_right: lnewman

:point_right: mlaane

