# go-ls: Go ls implementation

Small [GNU `ls`](https://www.gnu.org/software/coreutils/manual/html_node/ls-invocation.html#ls-invocation)
Go implementation for learning purposes.

## Usage

```bash
go-ls --help
Usage of go-ls:
  -a    show hidden files
  -l    long listing format

go-ls -l -a
drwxr-xr-x user group  384 May 20 17:13 .git     
-rw-r--r-- user group  177 May 20 13:09 README.md
-rw-r--r-- user group  133 May 20 14:20 config.go
-rw-r--r-- user group  667 May 20 17:00 file.go  
-rw-r--r-- user group  371 May 20 15:59 go.mod   
-rw-r--r-- user group 1376 May 20 15:59 go.sum   
-rw-r--r-- user group  855 May 20 17:12 ls.go    
-rw-r--r-- user group  416 May 20 16:02 main.go  
```
