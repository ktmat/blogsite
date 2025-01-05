Title: Relearning C - Chapter 1
Slug: cchap1
Parent: C Programming
Order: 6
MetaPropertyTitle: Relearning C - Chapter 1 Notes
MetaDescription: Notes on the C Programming Language, Chapter 1 notetaking.
MetaOgURL: https://matolat.com/cchap1
---
# Chapter 1 - A Tutorial Introduction

## 1.1 Getting Started
The only way to learn a new programming language is by writing programs in it. The first program to write is the same for all languages:

```c
#include <stdio.h>

int main() {
    printf("Hello, World!\n");
    return 0;
}
```
The main function is declared with the return type of an **int**. The return value is 0 after running the printf statement. Printf is included in the stdio header file.


## 1.2 Variables and Arithmetic Expressions
The next program uses the formula $C = (5/9)(F - 32)$ to print the following table of Fahrenheit temperatures and their centigrade or Celsius equivalents.

```
 1 -17
 20 -6
 40 4
 60 15
 80 26
 100 37
 120 48
 140 60
 160 71
 180 82
 200 93
 220 104
 240 115
 260 126
 280 137
 300 148
 ```
 
The program itself still consists of the definition of a single function named main.. It is longer than the one that printed ``hello, world``, but not complicated. It introduces several new ideas, including comments, declarations, variables, arithmetic expressions, loops, and formatted output.
```c
#include <stdio.h>

main() {
    int fahr, celsius;
    int lower, upper, step;

    lower = 0;
    upper = 300;
    step = 20;

    fahr = lower;
    while (fahr <= upper) {
        celsius = 5 * (fahr-32) / 9;
        printf("%d\t$d\n", fahr, celsius);
        fahr = fahr + step;
    }
}
```
The C Programming Language provides several other data types besides int and float, including:
| Type | Explanation |
| ---- | ----------- |
| char | character - a single byte |
| short | short integer |
| long | long integer |
| double | double-precisions floating point |

The size of these objects is also machine-dependent. There are also arrays, structures and u8nions of these basic types, pointers to them, and functions that return them, all of which we will meet in due course.

## 1.3 The for Statement
There are plenty of different ways to write a program for a particular task. Let's try a variation on the temperature converter.
```c
#include <stdio.h>

main() {
    int fahr; 
    for (fahr = 0; fahr <= 300; fahr = fahr + 20) {
        printf("%3d %6.1f\n", fahr, (5.0/9.0)*(fahr-32));
    }
}
```
The choice between while and for is arbitrary, based on which seems cleaer. The for is usually appropriate for loops in which the initialisation and increment are single statements and logically related, since it is more compacy than while and it keeps the loop control statements together in one place.

**Exercise 1-5.** Modify the temperature conversion program to print the table in reverse order, that is, from 300 degrees to 0.
```c
#include <stdio.h>

main() {
    int fahr;
    for (fahr = 300; fahr >= 0; fahr = fahr - 20) {
        printf("%3d %6.1f\n", fahr, (5.0/9.0)*(fahr-32));
    }
}
```

## 1.5 Character Input and Output
We are going to consider a family of related programs for processing character data. You will find that many programs are just expanded versions of the prototypes that we discuss here.

The model of input and output supported by the standard library is very simple. Text input or output, regardless of where it originates or where it goes to, is dealt with as streams of characters. A text stream is a sequence of characters divided into lines; each line consists of zero or more characters followed by a newline character. It is the responsibility of the lbirary to make each input or output stream confirm this model; the C programmer using hte library need not worry about how lines are represented outside the program.

The standard library provides several functions for reading or writing one character at a time, of which ```getchar``` and ```putchar``` are the simplest.. Each time it is called, ```getchar``` reads the next input characters from a text stream and returns that as its value That is, after
```c
c = getchar();
```
the variable ```c``` contains the next character of input. The characters normally come from the keyboard.

The function ```putchar``` prints a character each time it is calle:
```c
putchar(c);
```
prints the content of the integer variable c as a character, usually o nthe screen. Calls to ```putchar``` and ```printf``` may be interleaved; the output will appear in the order in which the calls are made.

## 1.5.1 File Copying
Given ```getchar``` and ```putchar```, you can write a surprising amount of useful code without knowing anything more about input and output. The simplest example is a program that copies its input to its output one character at a time:
```
read a character
   while (character is not end-of-file indicator)
    output the character just read
    read a character
```
Converting into C gives:
```c
#include <stdio.h>

main() {
    int c;
    c = getchar();
    while (c != EOF) {
        putchar(c);
        c = getchar();
    }
}
```
The program for copying would be written more concisely by experienced C programmers. In C, any assignment, such as
```c
c = getchar();
```
is an expression and has a value, which is the value of the left hand side after the assignment. This means that an assignment can appear as part of a larger expression, If the assignment of a character to ```c``` is put inside the test part of a ```while``` loop, the copy program can be written this way:
```c
#include <stdio.h>

main() {
    int c;
    
    while ((c = getchar()) != EOF) {
        putchar(c);
    }
}
```
The ```while``` gets a character, assigns it to ```c```, and then tests whether the character was the end-of-file signal. If it was not, the body of the ```while``` is executed, printing the character. The ```while``` then repeats. When the end of the input is finally reached, the ```while``` terminates and so does ```main```.

## 1.5.3 Line Counting
The next program counts input lines. As we mentioned above, the standard library ensures that an input text stream appears as a sequence of lines, each terminated by a newline. Hence, counting lines is just counting newlines:
```c
#include <stdio.h>

main() {
    int c, nl;

    nl = 0;
    while ((c = getchar()) != EOF) {
        if (c == '\n') {
            ++nl;
        }
    }
    printf("%d\n", nl);
}
```
**Exercise 1-8.** Write a program to count blanks, tabs, and newlines.
```c
#include <stdio.h>

main() {
    int c, nl, blanks, tabs;
    nl = 0;
    tabs = 0;
    blanks = 0;
    while ((c = getchar()) != EOF) {
        if (c == '\n') {
            nl++;
        } if (c == '\t') {
            tabs++;
        } if (c == ' ') {
            blanks++;
        }
    }
    printf("Blanks: %d\n", blanks);       // Print count of blanks
    printf("Tabs: %d\n", tabs);           // Print count of tabs
    printf("Newlines: %d\n", newlines);   // Print count of newlines
}
```

## 1.6 Arrays
This is how to write a program that counts the number of occurrences of each digit, of white space characters (blank, tab, newline), and of all other characters.
```c
#include <stdio.h>

main() {
    int c, i, nwhite, nother;
    int ndigit[10]; // declare an array of integers that can store up to 10 elements.

    nwhite = nother = 0;
    for (i = 0; i < 10; i++) {
        ndigit[i] = 0;
    }
    while ((c = getchar()) != EOF) {
        if (c >= '0' && c <= '9') {
            ++ndigit[c-'0'];
        } else if (c == ' ' || c == '\n' || c == '\t') {
            ++nwhite;
        } else {
            ++nother;
        }
    }
    printf("digits = ");
    for (i = 0; i < 10; ++i) {
        printf("%d", ndigit[i]);
    }
    printf(", white space = %d, other = %d\n", nwhite, nother);
}
```
**Exercise 1-13.** Write a program to print a histogram of the lengths of words in its input. It is easy to draw the histogram with the bars horizontal; a vertical orientation is more challenging.
```c
#include <stdio.h>

#define MAX_WORD_LENGTH 20  // Maximum word length tracked
#define IN  1               // Inside a word
#define OUT 0               // Outside a word

int main() {
    int c, state, word_length;
    int lengths[MAX_WORD_LENGTH + 1] = {0};  // Array to store word length counts
    int max_count = 0;  // Track the maximum count for scaling

    state = OUT;
    word_length = 0;

    // Count word lengths
    while ((c = getchar()) != EOF) {
        if (c == ' ' || c == '\n' || c == '\t') {
            if (state == IN) {  // End of a word
                state = OUT;
                if (word_length > 0) {
                    if (word_length <= MAX_WORD_LENGTH) {
                        lengths[word_length]++;
                    } else {
                        lengths[MAX_WORD_LENGTH]++;  // Group overflow lengths
                    }
                }
                word_length = 0;
            }
        } else {
            state = IN;
            word_length++;
        }
    }

    // Find the maximum count for scaling
    for (int i = 1; i <= MAX_WORD_LENGTH; i++) {
        if (lengths[i] > max_count) {
            max_count = lengths[i];
        }
    }

    // Print the vertical histogram
    for (int row = max_count; row > 0; row--) {
        for (int i = 1; i <= MAX_WORD_LENGTH; i++) {
            if (lengths[i] >= row) {
                printf("  # ");
            } else {
                printf("    ");
            }
        }
        printf("\n");
    }

    // Print the x-axis labels
    for (int i = 1; i <= MAX_WORD_LENGTH; i++) {
        printf("%3d ", i);
    }
    printf("\n");

    return 0;
}
```
**Exercise 1-14.** Write a program to print a histogram of the frequencies of different characters in its input.
```c
#include <stdio.h>

#define MAX_CHARS 128  // Maximum number of ASCII characters

int main() {
    int c, i, max_freq = 0;
    int frequencies[MAX_CHARS] = {0};  // Array to store character frequencies
    int used_chars[MAX_CHARS] = {0};  // Tracks which characters are used
    int unique_count = 0;  // Count of unique characters

    // Read input and count frequencies
    while ((c = getchar()) != EOF) {
        if (c >= 0 && c < MAX_CHARS) {
            if (frequencies[c] == 0) {
                used_chars[unique_count++] = c;  // Add to the list of used characters
            }
            frequencies[c]++;
        }
    }

    // Find the maximum frequency for scaling
    for (i = 0; i < unique_count; i++) {
        if (frequencies[used_chars[i]] > max_freq) {
            max_freq = frequencies[used_chars[i]];
        }
    }

    // Print the histogram
    for (int row = max_freq; row > 0; row--) {
        for (i = 0; i < unique_count; i++) {
            if (frequencies[used_chars[i]] >= row) {
                printf("  # ");
            } else {
                printf("    ");
            }
        }
        printf("\n");
    }

    // Print x-axis labels
    for (i = 0; i < unique_count; i++) {
        printf("  %c ", used_chars[i]);
    }
    printf("\n");

    return 0;
}
```

## 1.7 Functions
In C, a function is equivalent to a subroutine or function in Fortran. It allows for encapsulation and keep code clean.

Here is the function ```power``` and a main program to exercise it, so you can see the whole structure at once.
```c
#include <stdio.h>

int power(int m, int n);

int main() {
    int i;

    for (i = 0; i < 10; ++i) {
        printf("%d %d $d\n", i, power(2,i), power(-3, i));
    }
    return 0;
}

int power(int base, int n) {
    int i, p;
    p = 1;
    for (i = 1; i <= n; ++i) {
        p = p * base;
    }
    return p;
}
```
A function definition has the form:
```
return-type function-name(parameter declarations, if any) {
    declarations
    statements
}
```
## 1.9 Character Arrays
The most common type of array in C is the array of characters.
```
while (there is another line)
    if (it is longer than the previous longest)
        (save it)
        (save its length)
print longest line
```
```c
#include <stdio.h>
#define MAXLINE 1000

int getline(char line[], int maxline);
void copy(char to[], char from[]);

int main() {
    int len;
    int max;
    char line[MAXLINE];
    char longest[MAXLINE];

    max = 0;
    while ((len = getline(line, MAXLINE)) > 0) {
        if (len > max) {
            max = len;
            copy(longest, line);
        }
    }
    if (max > 0) {
        printf("%s", longest);
    }
    return 0;
}

int getline(char s[], int lim) {
    int c, i;

    for (i = 0; i < lim - 1 && (c = getchar()) != EOF && c != '\n'; ++i) {
        s[i] = c;
    }
    if (c == '\n') {
        s[i] = c;
        ++i;
    }
    s[i] = '\0';
    return i;
}

void copy(char to[], char from[]) {
    int i;
    i = 0;
    while ((to[i] = from[i]) != '\0') {
        ++i;
    }
}
```
A string in C is stored as such:
|   h   |   e   |   l   |   l   |   o   |  \n   |  \0   |
|-------|-------|-------|-------|-------|-------|-------|

**Exercise 1-19.** Write a function ```reverse(s)``` that reverses the character string ```s```.
```c
#include <stdio.h>
#include <string.h>

#define MAXLINE 1000

void reverse(char s[]) {
    int length = strlen(s);
    printf("strlen: %d\n", strlen(s));
    int i, temp;

    // Example with hello.
    printf("length: %d\n", length);
    for (i = 0; i < length / 2; i++) {
        temp = s[i]; // temp = h (first iteration)
        s[i] = s[length - i - 1]; // s[0] = s[6 - 0 - 1], s[0] = s[5] = \n
        s[length - i - 1] = temp; // temp = s[6 - 0 - 1], temp = \n
    }
}

int main() {
    char line[MAXLINE];

    while (fgets(line, MAXLINE, stdin) != NULL) {
        size_t len = strlen(line);
        if (len > 0 && line[len - 1] == '\n') {
            line[len - 1] == '\0';
        }
        reverse(line);
        printf("%s\n", line);
    }
    return 0;
}
```