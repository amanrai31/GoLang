
### Methods to read input

1. Scan() [takes input as raw ; space is considered as end of one input; newline count as space]
2. Scanln() 
3. Scanf()


*** Comma OK || Comma Error syntax** 

```
reader := bufio.NewReader(os.stdin)
input, _ := reader.ReadString('\n')    // underscroe is for error that may occur. 

```