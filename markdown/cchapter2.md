Title: Relearning C - Chapter 2
Slug: Relearning C - Chapter 2
Parent: C Programming
Order: 7
MetaPropertyTitle: Relearning C - Chapter 2 Notes
MetaDescription: Notes on the C Programming Language, Chapter 2 notetaking.
MetaOgURL: https://matolat.com/cchap2
---
# Chapter 2 - Types, Operators and Expressions
Variables and constants are the basic data objects manipulated in a program. Declarations list the variables to be used, and state what type they have and perhaps what their initial values are. Operators specify what is to be done to them. Expressions combine variables and constants to produce new values. The type of an object determines the set of values it can have and what operations can be performed on it. These building blocks are the topic of this chapter.

## 2.2 Data Types and Sizes
|Type|Explanation|
|----|-----------|
|```char```|A single byte, capable of holding one character in the local character set.|
|```int```|An integer, typically reflecting the natural size of integers on the host machine.|
|```float```|Single-precision floating point.|
|```double```|Double-precision floating point.|

In addition, there are a number of qualifiers that can be applied to these basic types. ```short``` and ```long``` apply to integers.
```c
short int sh;
long int counter;
```
The word ```int``` can be omitted in such declarations, and typically it is.

```short``` is often 16 bits long, and ```int``` is either 16 or 32 bits long. ```long``` is at least 32 bits, and ```short``` is no longer than ```int```, which is no longer than ```long```.
The qualifier ```signed``` or ```unsigned``` may be applied to ```char``` or any integer. ```unsigned``` numbers are always positive or zero, and obey the laws of arithmetic modulo 2^n, where n is the number of bits in the type.

The standard headers ```<limits.h>``` and ```<float.h>``` contain symbolic constants for all of these sizes, along with other properties of the machine and compiler.

## 2.9 Bitwise Operators
C provides six operators for bit manipulation. These can only be applied to ```char```, ```short```, ```int```, and ```long```. Whether ```signed``` or ```unsigned```.
|Operator|Explanation|
|--------|-----------|
|&|Bitwise AND|
|\||Bitwise inclusive OR|
|^|Bitwise exclusive OR|
|<<|Left shift|
|>>|Right shift|
|~|One's complement (unary)|

## 2.11 Conditional Expressions
**Exercise 2-10.** Rewrite the function lower, which converts upper case letters to lower case, with a conditional expression instead of if-else.
```c
int lower(int c) {
    return (c >= 'A' && c <= 'Z') ? (c + 'a' - 'A') : c;
}
```