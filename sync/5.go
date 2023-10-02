// Question 5: Semaphore
// It is a datastructure invented by Edsger Dijkstra useful to solve many synchronization problems
// It is an integer with 2 operations
// 1. acquire (known also as wait, decrement or P)
// Operation acquire decreases the semaphore by 1.
// If the result is negative then thread blocks and cannot resume until other
// thread will increment the semaphore. If result is positive then thread continues execution.

// 2. release (signal, increment or V)
// Operation release increases the semaphore by 1. If there’re blocked threads then one of them gets unblocked.