<p align="center"><img src="img/bg.png"></p>
<p align="center">
  <a href="LICENSE"><img alt="LICENSE MIT" src="https://img.shields.io/bower/l/bootstrap?color=ff69b4&label=LICENSE&style=flat-square"></a>
  <a href="https://github.com/pblcc/encrl/issues"><img alt="GitHub issues" src="https://img.shields.io/github/issues/pblcc/encrl?color=ff69b4&label=ISSUES&style=flat-square"></a>
  <img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/pblcc/encrl?color=ff69b4&label=COMMITS&style=flat-square">
  <img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/pblcc/encrl?color=ff69b4&label=SIZE&style=flat-square">
</p>

### Table of contents
- [Idea](#idea)
- [Usage](#usage)
  - [Basic usage](#basic-usage)
  - [Usage options](#options)
- [Compiling from source](#compiling-from-source)

---

### Idea
Encrl is a tool programmed by [Pablo Corbalán](https://github.com/pblcc) that was created with the intention of learning to use the Golang programming language.

Encrl is presented as a CLI *Console Line App* format.

Encrl is an **encryption and decryption tool based on hash tables**. We are working on all the ciphers for encrl but you can see the current ones in [the codificaions folder](/codifications/).

If you want to add a cipher to Encrl, you can! Remember that the ciphers are represented in a json file as:
```
{
  "a": "codification of a",
  "b": "codification of b",
  ...
}
```
As I have said before this is my first real project in Golang, so I appreciate any input or advice!

---

### Usage
###### Basic usage:
```shell
./encrl -r reading-file -w writing-file
```
###### Options:
```
-c <cipher>: The  cipher to use
```

---

### Compiling from source
If you want to compile the encrl **from the source code**, You will need:
```
Git v[1.8]+
Golang v[1.12]+
```
Follow this steps:
###### Open a terminal / cmd
###### Clone the repository
```
git clone https://github.com/pblcc/encrl.git
```
###### Get inside the repository source code
```
cd encrl/encrl
```
###### Build the code using Go
```
go build
```
###### Use the executable
