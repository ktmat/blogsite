Title: Relearning C - Mastering Algorithms with C - Chapter 1
Slug: Relearning C - Mastering Algorithms with C - Chapter 1
Parent: C Programming
Order: 11
MetaPropertyTitle: Relearning C - Mastering Algorithms with C - Chapter 1 - Notes
MetaDescription: Notes on the C Programming Language, Mastering Algorithms with C - Chapter 1
MetaOgURL: https://matolat.com/masteringchap1
---
# Chapter 1: Introduction

## An Introduction to Data Structures
Data comes in all shapes and sizes, but often it can be organised the same way. A list is one example of a data structure, however, there are many other ways to organise data in computing. These are;
- Linked Lists
- Stacks
- Queues
- Sets
- Hash Tables
- Trees
- Heaps
- Priority Queues
- Graphs

Three reasons for using data structures are;
- Efficiency - Data structures organise data in ways that make algorithms more efficient.
- Abstraction - Data structures provide a more understandable way to look at data.
- Reusability - Data structures are reusable because they tend to be modular and context-free.

When one thinks of data structures, one normally thinks of certain actions, or operations, one would like to perform with them as well. For example, with a list, we might naturally like to insert, remove, traverse, and count elements. A data structure together with basic operations like these is called an *abstract datatype*. The operations of an abstract datatype constitute its *public interface*. The public interface of an abstract datatype defines exactly what we are allowed to do with it. Establishing and adhering to an abstract datatype's interface is essential because this lets us better manage a program's data, which inevitably makes a program more understandable and maintainable.

## An Introduction to Algorithms
Algorithms are well-defined procedures for solving problems. In computing, algorithms are essential because they serve as the systematic procedures that computers require. A good algorithm is like using the right tool in a workshop. As with data structures, three reasons for using formal algorithms are;
- Efficiency - Certain types of problems occur often in computing, researchers have found efficient ways of dsolving them over time.
- Abstraction - Algorithms provide a level of abstraction in solving problems because many seemingly complicated problems can be distilled into simpler ones for which well-known algorithms exist.
- Reusability - Algorithms are often reusable in many different situations. Since many well-known algorithms solve problems that are generalisations of more complicated ones, and since many complicated problems can be distilled into simpler ones, an efficient means of solving certain simpler problems potentially lets us solve many others.

## General Approaches in Algorithm Design
In a broad sense, many algorithmms approach problems in the same way. Thus, it is often convenient to classify them based on the approach they emply. One reason to classify algorithms in this way is that often we can gain some insight about an algorithm if we understand its general approach. This can also give us ideas about howq to look at similar problems for which we do not know algorithms. Of course, some algorithms defy classification, whereas others are based on a combination of approaches.

### Randomised Algorithms
Randomised algorithms rely on the statistical properties of random numbers. One example of a randomised algorithm is *quicksort*.

Imagine sorting a pile of canceled checks by hand. We begin with an unsorted pile that we partition in two.. In one pile we place all checks numbered less than or equal to what we think may be the median value, and in the other pile we place the checks numbered greater than this. Once we have the two piles, we divide each of them in the same manner and repeat the process until we end up with one check in every pile. At this point the checks are sorted.

In order to achieve good performance, quicksort relies on the fact that each time we partition the checks, we end up with wtwo partitions that are nearly equal in size. To accomplish this, ideally we need to look up the median value of the check numbers before partitioning the checks. However, since determining the median requries scanning all of the checks, we do not do this. Instead, we randomly select a check around which to partition. Quicksort performs well on average because the normal distribution of random nubmers leads to relatively balanced paritioning overall.

### Divide-and-conquer Algorithms
Divide-and-conquer algorithms resovle around three steps: *divide, conquer, and combine*. In the divide step, we dividde the data into smaller, more manageable pieces. In the conquer step, we process each division by performing some operation on it. In the combine step, we recombine the processed divisions. One example of a divide-and-conquer algorithm is *merge sort*.

As before, sorting a pile of cancelled checks by hand. We begin with an unsorted pile that we divide in half. Next, we divide each of the resulting two piles in half and contiue this process until we end up with one check in every pile. Once all piles contain a single check, we merge the piles two by two so that each new pile is a sorted combination of the two that were merged. Merging continues until we end up with one big pile again, at which point the checks are sorted.

In terms of the three steps common to all divide-and-conquer algorithms, merge sort can be described as follows. First, in the divide step, divide the data in half. Next, in the conquer step, sort the two divisions by recursively applying merge sor to them. Last, in the combine step, merge the two divisions into a single sorted set.

### Dynamic-programming Solutions
Dynamic-programming solutions are similar to divide-and-conquer methods in that both solve problems by breaking larger problems into subproblems whose results are later recombined. However, the approaches differ in how subproblems are related. In divide-and-conquer algorithms, each subproblem is independent of the others. Therefore, we solve each subproblem using recursion and combine its result with the results of other subproblems. In dynamic-programming solutions, subproblems are not independent of one another In other words, subproblems may share subproblems In problems like this, a dynamic-programming solution is better than a divide-and-conquer approach because the latter approach will do more work than necessary, as shared subproblems are solved more than once. Although it is an important technique used by many algorithms.

### Greedy Algorithms
Greedy algorithms make decisions that look best at the moment. In other words, they make decisions that are locally optimal in the hopet hat they will lead to globally optimal solutions. Unfortunately, decisions that look best at the moment are not always the best in the long run. Therefore, greedy algorithms do not always produce optimal results; however, in some cases they do. One example of a greedy algorithm is *Huffman coding*, which is an algorithm for data compression.

The most significant part of Huffman coding is building a *Huffman tree*. To build a Huffman tree, we proceed from its leaf nodes upward. We begin by placing each symbol to compress and the number of times it occurs in the data (its frequency) in the root node of its own binary tree. Next, we merge the two trees whose root nodes have the smallest frequencies and store the sum of the frequencies in the new tree's root We then repeat this process until we end up with a single tree, which is the final Huffman tree. The root node of this tree contains the total number of symbols in the data, and its leaf nodes contain the original symbols and their frequencies. Huffman coding is greedy because it continually seeks out the two trees that appear to be the best to merge at any given time.

### Approximation Algorithms
Approximation algorithms are algorithms that do not compute optimal solutions; instead, they compute solutions that are "good enough". Often we use approximation algorithms to solve problems that are computationally expensive but are too significant to give up on altogether. The *traveling-salesman problem* is one exmaple of a problem usually solved using an approximation algorithm.

Imagine a salesman who needs to visit a number of cities as part of the route he works. The goal in the traveling-salesman problem is to find the shrotes route possible by which the salesman can visit every city exactly onces before returing to the point at which he starts Since an optimal solution to the traveling-salesman problem is possible but computationally expensive, we use a *heuristic* to come up with an approximate solution. A heuristic is a less than optimal strategy that we are willing to accept when an optimal strategy is not feasible.

The travelling-salesman problem can be represented graphically by depicting the cities the salesman must visit as points on a grid. We then look for the shortest tour of the points by applying the following heuristic. Begin with a tour consisting of only the points at which the salesman starts. Colour this point black. All other points are white until added to the tour, at which time they are coloured black as well. Next, for each point *v* not already in the tour, compute the distance between the last point *u* added to the tour and *v*. Using this, select the point closest to *u*, colour it black, and add it to the tour. Repeat this process until all points have been coloured black. Lastly, add the starting point to the tour again, thus making the tour complete.