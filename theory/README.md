# POPULAR THEORY QUESTIONS

1. Difference between Stack and Heap memory allocation?
* Stack usually refers to the call stack of a thread.
* Stack stores temporary variables created by a function. 
* In stack variables are created, stored and initiated at runtime. It is a temporary storage memory. When computing task is complete the memory of the variable will automatically erased.

* Stack is a **linear** datastructure while heap is a **hierarchical** datastructure.
* Stack access **local** variables while heap allows you to access variables **globally**.
* Stack has faster access compare to heap.
* Stack variables can’t be resized whereas Heap variables can be resized.
* Stack memory is allocated in a contiguous block whereas Heap memory is allocated in any random order.

How do we know when a variable is allocated to the heap?
Go's compiler will allocate variables that are local to a function in function's stack frame. However if the compiler can not prove the variable referenced after the function returns, then compiler allocated the memory on garbage collector heap memory to avoid any pointer errors. 
Also if local variable is very large it is better to store that on heap memory.
If a variable has its address taken, then that variable is a candidate to be stored on heap memory.
BUT compiler implementation changes over time, so there's no way of knowing which variables will be allocated to heap memory just by reading the GO code.  


3. API designing questions. Like, design an API that gives top trending movies of a country. 
4. What procedure do you follow while designing a service. 
5. How do you deploy a service? 
6. How do you choose a database. What are the factors? and other questions related to DB.
7. What makes go better than Java and python? Java and Python also support concurrency but what makes GO better?
8. Microservices vs monoliths