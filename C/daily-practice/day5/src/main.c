#include <stdio.h> 
#include <stdlib.h>

int comp(const void *a, const void *b) {
    return (*(int *)b - *(int *)a);
}

int main() {
    int n;
    int a[100];
    int halfValue = 0;
    int value = 0;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%d", &a[i]);
        halfValue += a[i];
    }
    halfValue /= 2;
    qsort(a, n, sizeof(int), comp);
    for (int i = 0; i < n; i++) {
        value += a[i];
        if (value > halfValue) {
            printf("%d\n", i+1);
            return 0;
        }
    }
}
