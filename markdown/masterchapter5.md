Title: Relearning C - Mastering Algorithms with C - Chapter 5
Slug: Relearning C - Mastering Algorithms with C - Chapter 5
Parent: C Programming
Order: 15
MetaPropertyTitle: Relearning C - Mastering Algorithms with C - Chapter 5 - Notes
MetaDescription: Notes on the C Programming Language, Mastering Algorithms with C - Chapter 5
MetaOgURL: https://matolat.com/masteringchap5
---
# Chapter 5 - Linked Lists
Linked lists are some of the most fundamental data structures Linked lists consist of a number of elements grouped, or *linked*, together in a specific order They are useful in maintaining collections of data, similar to the way that arrays are often used.

There are many different types of linked lists.
- Singly-linked lists
    - The simplest linked lists, in which elements are linked by a single pointer. This structure allows the list to be traversed from its first element to its last.

- Doubly-linked lists
    - Linked lists in which elements are linked by two pointers instead of one. This structure allows the list to be traversed both forward and backward.

- Circular lists
    - Linked lists in which the last element is linked to the first instead of being set to NULL. This structure allows the list to be traversed in a circular fashion.

Some applications of linked lists are:
- Mailing lists
- Scrolled lists
- Polynomials
- Memory management
- LISP
- Linked allocation of files
- Other data structures

## Description of Linked Lists
Singly-linked list, usually called linked lists, are composed of individual elements, each linked by a single pointer. Each element consists of two parts: a data member and a pointer, called the *next* pointer. Using this two-member structure, a linked list is formed by setting the *next* pointer of each element to point to the element that follows it. The next pointer of the last element is set to NULL, a covenient sentinel marking the end of the list. The element at the start of the list is its head; the element at the end of the list is its *tail*.

To access an elemment in a linked list, we start at the head of the list and use the *next* pointers of successive elements to move from element to element until the desired element is reached. With singly-linked lists, the list can be traversed in only one direction -- from head to tail -- because each element contains no link to its predecessor. Therefore, if we start at the head and move to some element, and then wish to access an element preceding it, we must start over at the head (although sometimes we can anticipate the need to know an element and save a pointer to it). Often this weakness is not a concern. When it is, we use a doubly-linked list, or circular list.

Conceptually, one thinks of a linked list as a series of continguous elements. However, because these elements are allocated dynamically (using *malloc* in C), it is important to remember that, in actuality, they are usually scattered about in memory. The pointers from element to eleent therefore are the only means by which we can ensure that all elements remain accessible. With this in mind, we will see later that special care is required when it comes to maintaining the links. If we mistakenly drop one link, it becomes impossible to access any of the elements from that point on in the list. Thus, the expression "You are only as strong as your weakest link" is particularly fitting for linked lists.
![linkedlist](/static/images/linkedlist.png)

![linkedlistscattered](/static/images/linkedlistscattered.png)

## Interface for Linked Lists

**list_init**
```c
void list_init(List *list, void (*destry)(void *data));
```
**Return Value:** None.

**Description:** Initialises the linked list specified by ```list```. This operation must be called for a linked list before the list can be used with any other operation. The ```destroy``` arguments provides a way to free dynamically allocated data when ```list_destroy``` is called. For example, if the list contains data dynamically allocated using ```malloc```, ```destroy``` should be set to free to free the data as the linked list is destroyed. For structured data containing several dynamically allocated members, ```destroy``` should be set to a user-defined function that calls free for each dynamically allocated member as well as for the structure itself. For a liniked list containing data that should not be freed, ```destroy``` should be set to NULL.

**Complexity:** $ \mathcal{O}(1) $

**list_destroy**
```c
void list_destroy(List *list);
```
**Return Value:** None.

**Description:** Destroys the linked list specified by ```list```. No other operations are permitted after calling ```list_destroy``` unless ```list_init``` is called again. The ```list_destroy``` operation removes all elements from a linked list and calls the function passed as ```destroy``` to ```list_init``` once for each element as it is removed, provided ```destroy``` was not set to NULL.

**Complexity:** $ \mathcal{O}(n) $, where $n$ is the number of elements in the linked list.

**list_ins_next**
```c
int list_ins_next(List *list, ListElmt *element, const void *data);
```
**Return Value:** $0$ if inserting the element is successful, or $-1$ otherwise.

**Description:** Inserts an element just after ```element``` in the linked list specified by ```list```. If ```element``` is NULL, the new element is inserted at the head of the list. The new element contains a pointer to ```data```, so the memoryreferenced by ```data``` should remain valid as long as the element remains in the list. It is the responsibility of teh caller to manage the storage associated with ```data```.

**Complexity:** $\mathcal{O}(1)$

**list_rem_next**
```c
int list_rem_next(List *list, ListElmt *element, void **data);
```
**Return Value:** $0$ if removing the elements is successful, or $-1$ otherwise.

**Description:** Removes the element just after ```element``` from the linked list specified by ```list```. If ```element``` is NULL, the element at the head of the list is removed. Upon return, ```data``` points to the data stored in the element that was removed It is the responsibility of the caller to manage the storage associated with the data.

**Complexity:** $\mathcal{O}(1)$

**list_size**
```c
int list_size(const List *list);
```
**Return Value:** Number of elements in the list

**Description:** Macro that evaluates the number of elmenents in the linked list specified by ```list```.

**Complexity:** $\mathcal{O}(1)$

**list_head**
```c
ListElmt *list_head(const List *list);
```
**Return Value:** Element at the head of the list.

**Description:** Macro that evaluates to the element at the head of the linked list specified by ```list```.

**Complexity:** $\mathcal{O}(1)$

**list_tail**
```c
ListElmt *list_tail(const List *list);
```
**Return Value:** Element at the tail of the list.

**Description:** Macro that evaluates to the element at the tail of the linked list specified by ```list```.

**Complexity:** $\mathcal{O}(1)$

**list_is_head**
```c
int list_is_head(const ListElmt *element);
```
**Return Value:** $1$ if the element is at the head of the list, or $0$ otherwise.

**Description:** Macro that determines whether the element specified as ```element``` is at the head of a linked list.

**Complexity:** $\mathcal{O}(1)$

**list_is_tail**
```c
int list_is_tail(const ListElmt *element);
```
**Return Value:** $1$ if the element is at the tail of the list, or $0$ otherwise.

**Description:** Macro that determines whether the element specified as ```element``` is at the tail of a linked list.

**Complexity:** $\mathcal{O}(1)$

**list_data**
```c
void *list_data(const ListElmt *element);
```
**Return Value:** Data stored in the element.

**Description:** Macro that evaluates to the data stored in the element of a linked list specified by ```element```.

**Complexity:** $\mathcal{O}(1)$

**list_next**
```c
ListElmt *list_next(const ListElmt *element);
```
**Return Value:** Element following the specified element.

**Description:** Macro that evaluates to the element of a linked list following the element specified by ```element```.

**Complexity:** $\mathcal{O}(1)$

## Implementation and Analysis of Linked Lists
Recall that each element of a linked list consists of two parts: a data member and a pointer to the next element in the list. The structure ```ListElmt``` represents an individual element of a linked list As you would expect, this structure has two members that correspond to those just mentioned. The structure ```List``` is the linked list data structure. This structure consists of five members: ```size``` is the number of elements in the list, ```match``` is a member not used by linked lists but by datatypes that will be derived later from linked lists, ```destroy``` is the encapsulated destroy function passed to ```list_init```, ```head``` is a pointer to the first of the linked element, and ```tail``` is a pointer to the tail element.

```c
#ifndef LIST_H
#define LIST_H

#include <stdlib.h>

typedef struct ListElmt_ {
    void *data;
    struct ListElmt_ *next;
} ListElmt;

typedef struct List_ {
    int size;
    int (*match)(const void *key1, const void *key2);
    void (*destroy)(void *data);

    ListElmt *head;
    ListElmt *tail;
} List;

void list_init(List *list, void (*destroy)(void *data));
void list_destroy(List *list);
int list_ins_next(List *list, ListElmt *element, const void *data);
int list_rem_next(List *list, ListElmt *element, void **data);

#define list_size(list) ((list)->size)
#define list_head(list) ((list)->head)
#define list_tail(list) ((list)->tail)
#define list_is_head(list, element) ((element) == (list)->head ? 1 : 0)
#define list_is_tail(element) ((element)->next == NULL ? 1 : 0)
#define list_data(element) ((element)->data)
#define list_next(element) ((element)->next)
#endif
```
#### list_init
The ```list_init``` operation initialises a linked list so that it can be used in other operations. Initialising a linked list is a simple operation in which the ```size``` member of the list is set to 0, the ```destroy``` member to ```destroy```, and the ```head``` and ```tail``` pointers to NULL.

The runtime complexity of ```list_init``` is $\mathcal{O}(1)$ because all of the steps in initialising a linked list run in a constant amount of time.

#### list_destroy
The ```list_destroy``` operation destroys a linked lists. Primarily this means removing all elements from the list. The function passed as ```destroy``` to ```list_init``` is called once for each element as it is removed, provided ```destroy``` was not set to NULL.

The runtime complexity of list_destroy is $\mathcal{O}(n)$, where n is teh number of element in the list This is because the $\mathcal{O}(1)$ operation ```list_rem_next``` must be called once for each element.

#### list_ins_next
The ```list_ins_next``` operation inserts an element into a linked lists just after a specified element. The call sets the new element to point to the data passed by the caller. The actual process of inserting the new element into the list is a simple one, but it does require some care. There are two cases to consider: insertion at the head of the list and insertion elsewhere.

Generally, to insert an element into a linked list, we set the ```next``` pointer of the new element to point to the element it is going to precede, and we set the ```next``` pointer of the element that will precede the new element to point to the new element. However, when inserting at the head of a list, there is no element that will precede the new element. Thus, in this case, we set the ```next``` pointer of the new element to the current head of the list, then reset the head of the list to point to the new element. Recall from the interface design in the previous section that passing NULL for ```element``` indicates that the new element should be inserted at the head. In addition to these tasks, whenever we insert an element at the tail of the list, we must update the ```tail``` member of the list data structure to point to the new tail. Last, we update the size of the list by incremmenting its ```size``` member.
![insertelementLL](/static/images/insertelementLL.png)

The runtime complexity of ```list_ins_next``` is $\mathcal{O}(1)$ because all of the steps in inserting an element into a linked list run in a constant amount of time.

#### list_rem_next
The ```list_rem_next``` operation removes from a linked list the element just after a specified element. The reasons for removing the element just after, as opposed to the element itself is because, in the singly-linked list and circular list implementations, each element does not have a pointer to the one preceding it Therefore, we cannot set the preceding element's next pointer to the element after the one being removed An alternative approach to the one we selected would be to start at the head element and traverse the list, keeping track of each element preceding the next until the element to be removed is encountered. However, this solution is unattractive because the runtime complexity of removing an element from a singly-linked list or circular list degrades to $\mathcal{O}(n)$. Another approach would be to copy the data of the element following the specified element into the one specified and then remove the following element. However, this seemingly benign $\mathcal{O}(1)$ approach generates the dangerous side effect of rendering a pointer into the list invalid. This could be a surprise to a developer maintaining a pointer to the element after the one thought to be removed! The approach we selected, then, was to remove the element after the specified one The disadvantage of this approach is its inconsistency with the ```dlist_remove``` operation of the doubly-linked list implementation. However, this is addressed by the naming convention, using ```_rem_next``` as the suffix for removing an element after the one specified, and ```_remove``` to indicate that the specified element itself will be removed. In a doubly-linked list, recall that we can remove precisely the element specified because each element has a pointer to the one that precedes it.

As with inserting an element, this call requires consideration of two cases: removing an element from the head of the list and removing one elsewhere.

The actual process of removing the element from the list is a simple one, but it too requires some care. Generally, to remove an element from a linked list, we set the ```next``` pointer of the element preceding the one being removed to point to the element after the element being removed. However, when removing an element from the head of a list, there is no element that precedes the element being removed. Thus, in this case, we set the head of the list to point to the elmeent after the one being removed. As with insertion, NULL serves nicely as a sentinel passed in ```element``` to indicate that the element at the head of the list should be removed. In addition to these tasks, whenever we remove the element at the tail of the list, we must update the ```tail``` member of the list data structure to point to the new tail, or to NULL if removing the element has caused the list to become empty. Last, we update the size of the list by decreasing the ```size``` member by 1. Upon return, ```data``` points to the data from the element removed.
![removeelement](/static/images/removeelement.png)

The runtime complexity of ```list_rem_next``` is $\mathcal{O}(1)$ because all of the steps in removing an element from a linked list run in a constant amount of time.

#### list_size, list_head, list_tail, list_is_tail, list_data, and list_next
These macros implement some of the simpler linked list operations. Generally, they provide an interface for accessing and testing members of the ```List``` and ```ListElmt``` structures.

The runtime complexity of these operations is $\mathcal{O}(1)$ because accessing and testing members of a structure are simple tasks that run in a constant amount of time.

```c
// Implementation of the Linked List Abstract Datatype
#include <stdlib.h>
#include <string.h>

#include "list.h"

void list_init(List *list, void (*destroy)(void *data)) {
    list->size = 0;
    list->destroy = destroy;
    list->head = NULL;
    list->tail = NULL;
    return;
}

// list_destroy
void list_destroy(List *list) {
    void *data;
    // Remove each element
    while (list_size(list) > 0) {
        if (list_rem_next(list, NULL, (void **)&data) == 0 && list->destroy != NULL) {
            list->destroy(data);
        }
    }
    memset(list, 0, sizeof(List));
    return;
}

// list_ins_next
int list_ins_next(List *list, ListElmt *element, const void *data) {
    ListElmt *new_element;
    // Allocate storage for the element.
    if ((new_element = (ListElmt *)malloc(sizeof(ListElmt))) == NULL) {
        return -1;
    }
    // Insert the element into the list.
    new_element->data = (void *)data;

    if (element == NULL) {
        // Handle insertion at the head of the list.
        if (list_size(list) == 0) {
            list->tail = new_element;
        }
        new_element-next = list->head;
        list->head = new_element;
    } else {
        // Handle insertion somewhere other than at the head.
        if (element->next == NULL) {
            list->tail = new_element;
        }
        new_element->next = element->next;
        element->next = new_element;
    }
    // Adjust the size of the list to account for the inserted element.
    list->size++;
    return 0;
}

// list_rem_next
int list_rem_next(List *list, ListElmt *element, void **data) {
    ListElmt *old_element;
    
    // Do not allow removal from an empty list!
    if (list_size(list) == 0) {
        return -1;
    }
    // Remove the element from the list.
    if (element == NULL) {
        // Handle removal from the head of the list.
        *data = list->head->data;
        old_element = list->head;
        list->head = list->head->next;

        if (list_size(list) == 1) {
            list->tail = NULL;
        }
    } else {
        // Handle removal from somewhere other than the head.
        if (element->next == NULL) {
            return -1;
        }
        *data = element->next->data;
        old_element = element->next;
        element->next = element->next->next;

        if (element->next == NULL) {
            list->tail = element;
        }
    }
    // Free the storage allocated by the abstract datatype.
    free(old_element);
    // Adjust the size of the list to account for the removed element.
    list->size--;
    return 0;
}
```
## Linked List Example: Frame Management
An application of linked lists can be found in the way some systems support virtual memory. Virtual memory is a mappping of address space that allows a process to execute without being completely in physical memory; the real memory of the system. One advantage of this is that a process can make use of an address space that is much larger than that which the physical memory of the system would allow otherwise. Another advantage is that multiple processes can share the memory of the system while running concurrently.
![virtmemsys](/static/images/virtmemsys.png)

```c
// Implementation of Functions for managing frames

#include <stdlib.h>
#include "frames.h"
#include "list.h"

// alloc_frames
int alloc_frame(List *frames) {
    int frame_nubmer, *data;

    if (list_size(frames) == 0) {
        return -1; // no frames available
    } else {
        if (list_rem_next(frames, NULL, (void **)&data) != 0) {
            // Return that a frame could not be retrieved.
            return -1;
        } else {
            // Store the number of the available frame.
            frame_number = *data;
            free(data);
        }
    }
    return frame_number;
}

// free_frame
int free_frame(List *frames, int frame_number) {
    int *data;

    // Allocate storage for the frame number.
    if ((data = (int *)malloc(sizeof(int))) == NULL) {
        return -1;
    }
    // Put the frame back in the list of available frames.
    *data = frame_number; 
    if (list_ins_next(frames, NULL, data) != 0) {
        return -1;
    }
    return 0;
}
```
## Description of Doubly-Linked Lists
Doubly-linked lists are composed of elements linked by two pointers. Each element of a doubly-linked list consists of three parts: in addition to the data and the next pointer, each element includes a pointer to the previous element, called the prev pointer. A doubly-linked lists is formed by composing a number of elements so that the next pointer of each element pointers to the element that follows it, and the prev pointer points to the element preceding it. To mark the head and tail of the list, we set the prev pointer of the first element and the next pointer of the last element to NULL.

To traverse backward through a doubly-linked list, we use the prev pointers of consecutive elements in the tail-to-head direction. Thus, for the cost of an additional pointer for each element, a doubly-linked list offers greater flexibility than a singly-linked lists in moving about the list. This can be useful when we know something about where an element might be stored in the list and can choose wisely how to move to it. For example, one flexibility that doubly-linked lists provide is a more intuitive means of removing an element than singly-linked lists.

## Interface for Doubly-Linked Lists
**dlist_init**
```c
void dlist_init(DList *list, void (*destroy)(void *data));
```
**Return Value:** None.

**Description:** Initialises the doubly-linked list specified by ```list```. This operation must be called for a doubly-linked list before the list can be used with any other operation. The ```destroy``` argument provides a way to free dynamically allocated data when ```dlist_destroy``` is called. It works in a manner similar to that described for ```list_destroy```. For a doubly-linked list containing data that should not be freed, ```destroy``` should be set to NULL.

**Complexity:** \mathcal{O}(1)



### Implementation and Analysis of Doubly Linked Lists
```c
// Header for the Doubly-Linked List Abstract Datatype
// dlist.h
#ifndef DLIST_H
#define DLIST_H

#include <stdlib.h>

// Define a structure for doubly-linked list elements.
typedef struct DListElmt_ {
    void *data;
    struct DListElmt_ *prev;
    struct DListElmt_ *next;
} DListElmt;

// Define a structure for doubly-linked lists.

typedef struct DList_ {
    int size;
    int (*match)(const void *key1, const void *key2);
    void (*destroy)(void *data);
    DListElmt *head;
    DListElmt *tail;
} DList;
// Public Interface
void dlist_init(DList *list, void (*destroy)(void *data));

void dlist_destroy(DList *list);

int dlist_ins_next(DList *list, DListElmt *element, const void *data);

int dlist_ins_prev(DList *list, DListElmt *element, const void *data);

int dlist_remove(DList *list, DListElmt *element, void **data);

#define dlist_size(list) ((list)->size)

#define dlist_head(list) ((list)->head)

#define dlist_tail(list) ((list)->tail)

#define dlist_is_head(element) ((element)->prev == NULL ? 1 : 0)

#define dlist_is_tail(element) ((element)->next == NULL ? 1 : 0)

#define dlist_data(element) ((element)->data)

#define dlist_next(element) ((element)->next)

#define dlist_prev(element) ((element)->prev)

#endif
```
Here is how you make use of the above definitions.
```c
// dlist.c
#include <stdlib.h>
#include <string.h>

#include "dlist.h"

// dlist_init
void dlist_init(DList *list, void (*destroy)(void *data)) {
    // Initialise the list.
    list->size = 0;
    list->destroy = destroy;
    list->head = NULL;
    list->tail = NULL;
    return;
}

// dlist_destroy
void dlist_destroy(DList *list) {
    void *data;
    // Remove each element.
    while (dlist_size(list) > 0) {
        if (dlist_remove(list, dlist_tail(list), (void **)&data) == 0 && list->destroy != NULL) {
            // Call a user-defined function fto free dynamically allocated data.
            list->destroy(data);
        }
    }
    // No operations are allowed now, but clear the structure as a precaution.
    memset(list, 0, sizeof(DList));
    return;
}

// dlist_ins_next
int dlist_ins_next(DList *list, DListElmt *element, const void *data) {
    DListElmt *new_element;
    // Do not allow a NULL element unless the list is empty.
    if (element == NULL && dlist_size(list) != 0) {
        return -1;
    }
    // Allocate storage for the element.
    if ((new_element = (DListElmt *)malloc(sizeof(DListElmt))) == NULL) {
        return -1;
    }
    // Insert the new element into the list.
    new_element->data = (void *)data;

    if (dlist_size(list) == 0) {
        // Handle insertion when the list is empty.
        list->head = new_element;
        list->head->prev = NULL;
        list->head->next = NULL;
        list->tail = new_element;
    } else {
        // Handle insertion when the list is not empty.
        new_element->next = element->next;
        new_element->prev = element;

        if (element->next == NULL) {
            list->tail = new_element;
        } else {
            element->next->prev = new_element;
        }
        element->next = new_element;
    }
    // Adjust the size of the list to account for the inserted element.
    list->size++;
    return 0;
}
// dlist_ins_prev
int dlist_ins_prev(DList *list, DListElmt *element, const void *data) {
    DListElmt *new_element;
    // Do not allow a NULL element unless the list is empty.
    if (element == NULL && dlist_size(list) != 0) {
        return -1;
    }
    // Allocate storage to be managed by the abstract datatype.
    if ((new_element = (DListElmt *)malloc(sizeof(DListElmt))) == NULL) {
        return -1;
    }
    // Insert the new element into the list.
    new_element->data = (void *)data;
    if (dlist_size(list) == 0) {
        // Handle insertion when the list is empty.
        list->head = new_element;
        list->head->prev = NULL;
        list->head->next = NULL;
        list->tail = new_element;
    } else {
        // Handle insertion when the list is not empty.
        new_element->next = element;
        new_element->prev = element->prev;
        
        if (element->prev == NULL) {
            list->head = new_element;
        } else {
            element->prev->next = new_element;
        }
        element->prev = new_element;
    }
    // Adjust the size of the list to account for the new element.
    list->size++;
    return 0;
}
// dlist_remove
int dlist_remove(DList *list, DListElmt *element, void **data) {
    // Do not allow a NULL element or removal from an empty list.
    if (element == NULL || dlist_size(list) == 0) {
        return -1;
    }
    // Remove the element from the list.
    *data = element->data;
    if (element == list->head) {
        // Handle removal from the head of the list.
        list->head = element->next;
        if (list->head == NULL) {
            list->tail = NULL;
        } else {
            element->next->prev = NULL;
        }
    } else {
        // Handle removal from other than the head of the list.
        element->prev->next = element->next;
        if (element->next == NULL) {
            list->tail = element->prev;
        } else {
            element->next->prev = element->prev;
        }
    }
    // Free the storage allocated by the abstract datatype.
    free(element);
    // Adjust the size of the list to account for the removed element.
    list->size--;
    return 0;
}
```