pseudo:

```
for i = 0 to A.length - 1
    min = i
    for j = i to A.length
        if A[j] < A[min]
            min = j
    val = A[min]
    A[min] = A[i]
    A[i] = val
```
