Title: Relearning C - Chapter 4
Slug: cchap4
Parent: C Programming
Order: 9
MetaPropertyTitle: Relearning C - Chapter 4 Notes
MetaDescription: Notes on the C Programming Language, Chapter 4 notetaking.
MetaOgURL: https://matolat.com/cchap4
---
# Chapter 4 - Functions and Program Structure
Functions break large computing tasks into smaller ones, and enable people to build on what others have done instead of starting over from scratch.

## 4.1 Basics of Functions
Design and write a program to print each line of its input that contains a particular pattern, or string of characters. For example, searching for the pattern of letters "ould" in the set of lines.
```
Ah Love! could you and I with Fate conspire 
To grasp this sorry Scheme of Things entire, 
Would not we shatter it to bits -- and then 
Re-mould it nearer to the Heart's Desire! 
```
will produce the output...
```
Ah Love! could you and I with Fate conspire
Would not we shatter it to bits -- and then
Re-mould it nearer to the Heart's Desire!
```
The job falls neatly into three pieces:
```
while (there is another line)
    if (the line contains the pattern)
        print it
```
Here goes the code.
```c
#include <stdio.h>
#define MAXLINE 1000

int getline(char line[], int max)
int strindex(char source[], char searchfor[]);

char pattern[] = "ould";

main() {
    char line[MAXLINE];
    int found = 0;

    while (getline(line, MAXLINE) > 0) {
        if (strindex(line, pattern) >= 0) {
            printf("%s", line);
            found++;
        }
    }
    return found;
}

int getline(char s[], int lim) {
    int c, i;

    i = 0;
    while (--lim > 0 && (c = getchar()) != EOF && c != '\n') {
        s[i++] = c;
    }
    if (c == '\n') {
        s[i++] = c;
        s[i] = '\0';
        return i;
    }
}

int strindex(char s[], char t[]) {
    int i, j, k;

    for (i = 0; s[i] != '\0'; i++) {
        for (j = i, k = 0; t[k] != '\0' && s[j] == t[k]; j++, k++) {
            ;
        }
        if (k > 0 && t[k] == '\0') {
            return i;
        }
    }
    return -1;
}
```