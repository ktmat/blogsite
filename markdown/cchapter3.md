Title: Relearning C - Chapter 3
Slug: cchap3
Parent: C Programming
Order: 8
MetaPropertyTitle: Relearning C - Chapter 3 Notes
MetaDescription: Notes on the C Programming Language, Chapter 3 notetaking.
MetaOgURL: https://matolat.com/cchap3
---
# Chapter 3 - Control Flow

## 3.1 Statements and Blocks
An expression such as ```x = 0``` or ```i++``` or ```printf(...)``` becomes a statement when it is followed by a semicolon.
```c
x = 0;
i++;
printf(...);
```
## 3.3 Else-If
To illustrate a three-way decision, here is the binary search function implementation in C.
```c
int binarySearch(int x, int v[], int n) {
    int low, high, mid;

    low = 0;
    high = n - 1;
    while (low <= high) {
        mid = (low + high) / 2;
        if (x < v[mid]) {
            high = mid + 1;
        } else if (x > v[mid]) {
            low = mid + 1;
        } else {
            return mid;
        }
    }
    return -1;
}
```
The fundamental decision is whether ```x``` is less than, greather than, or equal to the middle element ```v[mid]``` at each step; this is natural for ```else-if```.
**Exercise 3-1.** Our binary search makes two tests inside the loop, when one would suffice (at the price of more tests outside.) Write a version with only one test inside the loop and measure the difference in run-time.
```c
int binarySearch(int x, int v[], int n) {
    int low, high, mid;
    low = 0;
    high = n - 1;
    mid = (low + high) / 2;

    while (low <= high && x != v[mid]) {    
        if (x < v[mid]) {
            high = mid - 1;
        } else {
            low = mid + 1;
        }
    }
    if (x == v[mid]) {
        return mid;
    } else {
        return -1;
    }
}
```
