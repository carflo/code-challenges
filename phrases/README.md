## Phrase Counter
This tool will count the 3-word phrase frequency of any text given to it.
It will print the 100 most frequent phrases in the text provided.

Simple example:
| input                              | output               |
| ---------------------------------- | -------------------- |
| echo "i love sandwiches" \| ./main | i love sandwiches: 1 |

## Usage

* Build: `make`

* Run with stdin: `cat myreallylongfile.txt | ./main`
* Run with args: `./main myreallylongfile.txt`

* Get cpu profile: `make profile`

# Testing
Test files are provided for convinience:
* oos.txt:  On the Origin of Species, by Charles Darwin
* text.example: For simple sandwich lovers
* text.example2: Lorem Ipsum
