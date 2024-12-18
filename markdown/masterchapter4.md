Title: Relearning C - Mastering Algorithms with C - Chapter 4
Slug: Relearning C - Mastering Algorithms with C - Chapter 4
Parent: C Programming
Order: 14
MetaPropertyTitle: Relearning C - Mastering Algorithms with C - Chapter 4 - Notes
MetaDescription: Notes on the C Programming Language, Mastering Algorithms with C - Chapter 4
MetaOgURL: https://matolat.com/masteringchap4
---
# Chapter 4 - Analysis of Algorithms
This chapter covers:

- *Worst-case analysis*
    - The metric used by which most algorithms are compared Other cases we might consider are the average case and the best case. However, worst-case analysis usually offers several advantages.
    
- *$\mathcal{O}$-notation*
    - The most common notation used to formally express an algorithm's performance. $\mathcal{O}$-notation is used to express the upper bound of a function within a constant factor.

- *Computational complexity*
    - The growth rate of the resources (usually time) an algorithm requries with respect to the size of the data it processes. $\mathcal{O}$-notation is a formal expression of an algorithm's complexity.

## Worst-Case Analysis
Most algorithms do not perform the same in all cases; normally an algorithm's performance varies with the data passed to it. Typically, three cases are recognised: the best case, worst case, and average case. For any algorithm, understanding what constitutes each of these cases is an important part of analysis because performance can vary significantly between them. Consider even a simple algorithm such as *linear search*. Linear search is a natural but inneficient search technique in which we look for an element simply by traversing a set from one end to the other. In the best case, the element we are looking for is the first element we inspect, so we end up traversing only a single element. In the worst case, however, the desired element is the last one we inspect, in which case we end up traversing all of the elements. On average, we can expect to find the element somewhere in the middle.

### Reasons for Worst-Case Analysis
A basic understanding of how an algorithm performs in all cases is important, but usually we are most interested in how an algorithm performs in the worst case. There are four reasons why algorithms are generally analysed by their worst case:
- Many algorithms perform to their worst case a large part of the time. For example, the worst case in searching occurs when we do not find what we are looking for at all. Imagine how frequently this takes place in some database applications.
- The best case is not very informative because many algorithms perform exactly the same in the best case. For example, nearly all searching algorithms can locate an element in one inspection at best, so analysing this case does not tell us much.
- Determining average-case performance is not always easy Often it is difficult to determine exactly what the "average case" even is. Since we can seldom guarantee precisely how an algorithm will be exercised, usually we cannot obtain an average-case measurement that is likely to be accurate.
- The worst case gives us an upport bound on performance. Analysing an algorithm's worst case guarantees that it will never perform worse than what we determine. Therefore, we know that the other cases must perform at least as well.

Although worst-case analysis is the metric for many algorithms, it is worth noting that there are exceptions. Sometimes special circumstances let us base performance on the average case. For example, randomised algorithms such as quicksort use principles of probability to virtually guarantee average-case performance.

## $\mathcal{O}$-Notation
$\mathcal{O}$-notation is the most common notation used to express an algorithm's performance in a formal manner. Formally, $\mathcal{O}$-notation expresses the upper bound of a function within a constant factor. Specifically, if $g(n)$ is an upper bound of $f(n)$, then for some constant $c$, it is possible to find a value of $n$, call it $n_0$, for which any value of $n >= n_0$ will result in $f(n) <= cg(n)$.

Normally we express an algorithm's performance as a function of the size of the data it processes. That is, for some data of size $n$, we describe its performance with some function $f(n)$. However, while in many cases we can determine $f$ exactly, usually it is not necessary to be this precise Primarily we are interested only in the growth rate of $f$, which describes how quickly the algorithm's performance will degrade as the size of the data it processes becomes arbitrarily large. An algorithm's growth rate, or order of growth, is significant because ultimately it describes how efficient the algorithm inputs. $\mathcal{O}$-notation reflects an algorithm's order of growth.

### Simple Rules for $\mathcal{O}$-Notation
When we look at some function $f(n)$ in terms of its growth rate, a few things become apparent. First, we can ignore constant terms because as the value of n becomes larger and larger, eventually constant terms will become insignificant. Second, we can ignore constant multipliers of terms because they too will become insignificant as the value of $n$ increases. Finally, we need only consider the highest-order term because, again, as $n$ increases, higher-order terms quickly outweight the lower-order ones. These ideas are formalised in the following simple rules for expressing functions in $\mathcal{O}$-notation.

- Constant terms are expressed as $\mathcal{O}(1)$ When analysing the running time of an algorithm, apply this rule when you have a task that you know will execute in a certain amount of time regardless of the size of the data it processes. Formally stated, for some constant $c$:
    - $\mathcal{O}(c) = \mathcal{O}(1)$

- Multiplicative constants are omitted. When analysing the running time of an algorithm, apply this rule when you have a number of tasks that all execute in the same amount of time. For example, if three tasks each run in time $T(n) = n$, the result is $\mathcal{O}(3n)$, which simplifies to $\mathcal{O}(n)$. Formally stated, for some constant $c$:
    - $\mathcal{O}(cT) = c\mathcal{O}(T) = \mathcal{O}(T)$

- Addition is performed by taking the maximum. When analysing the running time of an algoirithm, apply this rule when one task is executed after another For example, if $T_1(n) = n$, and $T_2(n) = n^2$ describe two tasks executed sequentially, the result is $\mathcal{O}(n) + \mathcal{O}(n^2)$, which simplifies to $\mathcal{O}(n^2)$. Formally stated:
    - $\mathcal{O}(T_1) + \mathcal{O}(T_2) = O\mathcal{O}(T_1 + T_2) = max(\mathcal{O}(T_1), \mathcal{O}(T_2))$

- Multiplication is not changed but often is rewritten more compactly. When analysing the running time of an algorithm, apply this rule when one task causes another to be executed some number of times for each iteration of itself. For example, in a nested loop whose outer iterations are described by $T_1$ and whose inner iterations by $T_2$, if $T_1(n) = n$ and $T_2(n) = n$, the result is $\mathcal{O}(n)\mathcal{O}(n)$, or $\mathcal{O}(n^2)$. Formally stated:
    - $\mathcal{O}(T_1)\mathcal{O}(T_2) = \mathcal{O}(T_1T_2)$