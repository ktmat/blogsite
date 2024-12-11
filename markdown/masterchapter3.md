Title: Relearning C - Mastering Algorithms with C - Chapter 3
Slug: Relearning C - Mastering Algorithms with C - Chapter 3
Parent: C Programming
Order: 13
MetaPropertyTitle: Relearning C - Mastering Algorithms with C - Chapter 3 - Notes
MetaDescription: Notes on the C Programming Language, Mastering Algorithms with C - Chapter 3
MetaOgURL: https://matolat.com/masteringchap3
---
# Chapter 3 - Recursion
## Basic Recursion
Let's consider a problem that normally we might not think of in a recursive way. Suppose we would like to compute the factorial of a number *n*. *n!* is the product of all numbers from *n* down to 1. One way to calculate this is to loop through each number and multiply it with the product of all preceding numbers This is an iterative approach, which can be defined formally as:
```c
n! = (n)(n - 1)(n - 2)...
```
Another way to look at this problem is to define *n!* as the product of smaller factorials. To do this, we define *n!* as *n* times the factorial of *n* - 1. Solving (*n* - 1)! is the same problem as *n*!, only a little smaller. If we then think of (*n* - 1)! as *n* - 1 times (*n* - 2)!, (*n* - 2)! as *n* - 2 times (*n* - 3)!, and so forth until *n* = 1, we end up computing *n*!. This is a *recursive* approach.

Here is an implementation of a function for computing factorials recursively:
```c
#include <fact.h>
int fact(int n) {
    if (n < 0) {
        return 0;
    } else if (n == 0) {
        return 1;
    } else if (n == 1) {
        return 1;
    } else {
        return n * fact(n - 1);
    }
}
```
To understand how recursion really works, it helps to look at the way functions are executed in C. For this, we need to understand a little about the organisation of a C Program in memory. Fundamentally, a C program consists of four areas as it executes: a code area, a static data area, a heap, and a stack. The code data contains the machine instructions that are executed as the progam runs. The static data area contains data that persists throughout the life of the program, such as global variables and static local variables. The heap contains dynamically allocated storage, such as memory allocated by ```malloc```. The stack contains information about function calls By convention, the heap grows upward from mone end of a program's memory, while the stack grows downward from the other. (This may vary in practice) Note that the term *heap* as it is used in this context has nothing to do with the heap data structure.

When a function is called in a C program, a block of storage is allocated on the stack to keep track of information associated with the call. Each call is referred to as an *activation*. The block of storage palced on the stack is called an *activation record* or, alternatively, a *stack frame*. An activation record consists of five regions: incoming parameters, space for a return value, temporary storage used in evaluating expressions, saved state information for when the activation terminates, and outgoing parameters Incoming parameters are the parameters passed into the activation. Outgoing parameters are the parameters passed to functions called within the activation. The outgoing parameters of one activation record become the incoming parameters of the next one placed on the stack. The activation record for a function call remains on the stack until teh call terminates.

The stack is a great solution to storing information about function calls because its last-in, first-out behaviour is well suited to the order in which functions are called and terminated. However, stack usage does have a few drawbacks. Maintaining information about every function call until it returns takes a considerable amount of space, especially in programs with many recursive calls. In addition, generating and destroying activation records takes time because there is a significant amount of information that must be saved and restored. Thus, if the overhead associated with these concerns becomes too great, we may need to consider an iterative approach. Fortunately, we can use a special type of recursion, called *tail recursion*, to avoid these concerns in some cases.

## Tail Recursion
A recursive function is said to be *tail recursive* if all recursive calls within it are tail recursive. A recursive call is tail recursive when it is the last statement that will be executed within the body of a function and its return value is not a part of an expression. Tail-recursive functions are characterised as having nothing to do during the unwinding phase. This characteristic is important because most modern compilers automatically generate code to take advantage of it.

When a compiler detects a call that is tail recursive, it overwrites the current activation record instead of pushing a new one onto the stack. The compiler can do this because the recursive call is the last statement to be executed in the current activation; thus, there is nothing left to do in the activation when the call returns. Consequently, there is no reason to keep the current activation around. By replacing the current activation record instead of stacking another one on top of it, stack usage is greatly reduced, which leads to better performance in practice. Thus, we should make recursive functions tail recursive whenever we can.

Here is the implementation of the previous factorial function, but with it being *tail recursive*.
```c
#include <facttail.h>

int facttail(int n, int a) {
    if (n < 0) {
        return 0;
    } else if (n == 0) {
        return 1;
    } else if (n == 1) {
        return a;
    } else {
        return facttail(n - 1, n * a);
    }
}
```