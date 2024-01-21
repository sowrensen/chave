Chave
---

**Chave** is a program written in go to read and display ssh connection from your ssh config file by parsing it. When you have a lots of configs set up, it's easy to forget the names of the connections. This program will help you to remember the name without opening the entire config file.

### Etymology

**Chave** is a Portuguese word; which translates to "Key" in English. Also the Bengali word for Key – "**চাবি**" (Chabi) – is also derived from **Chave**.

### Usage

The program is pretty straightforward. Just clone the repository and run `go run .`

### Security

**Chave** will not disclose any credential info. The output only contains `Host` (_not to be confused with_ `HostName`) and `User`.

### Sample Output

```
Host                  User
-----                 -----
test-1                git
test-2                ubuntu
test-3                forge
```
