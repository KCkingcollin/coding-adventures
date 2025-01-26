#include <math.h>
#include <stdio.h> 
#include <stdlib.h>

int max(int a, int b) {
    return a > b ? a : b;
}

int comp(const void *a, const void *b) {
    return (*(int *)a - *(int *)b);
}

int main() {
    int n = 0;
    int maxX = 0;
    scanf("%d", &n);
    int *x = (int *)calloc(n, sizeof(int));
    for (int i = 0; i < n; i++) {
        scanf("%d", &x[i]);
        if (x[i] > maxX) {maxX = x[i];}
    }
    qsort(x, n, sizeof(int), comp);

    int q = 0;
    scanf("%d", &q);
    int ansArrSize = q*2;

    int m = 0;
    int ansCount = 0;
    char *answer = (char *)calloc(ansArrSize, sizeof(char));
    for (int i = 0; i < q; i++) {
        scanf("%d", &m);
        int L = 0;
        int R = n;
        if (m >= maxX) {
            R = n;
        } else {
            while (L < R) {
                int a = (L+R)/2;
                if (x[a] > m) {
                    R = a;
                } else {
                    L = a+1;
                }
            }
        }
        char Rstring[6] = {};
        sprintf(Rstring, "%d", R);
        int RstringSize = max((int)log10((double)R), 0) + 1;
        if (ansCount+RstringSize+1 > ansArrSize) {
            ansArrSize *= 2;
            answer = (char *)realloc(answer, ansArrSize * sizeof(char));
            if (answer == NULL) {
                printf("Memory reallocation failed\n");
                free(answer);
                return 1;
            }
        }
        for (int j = 0; j < RstringSize; j++) {
            answer[ansCount] = Rstring[j];
            ansCount++;
        }
        answer[ansCount] = '\n';
        ansCount++;
    }
    answer = (char *)realloc(answer, ansCount+1 * sizeof(char));
    if (answer == NULL) {
        printf("Memory reallocation failed\n");
        free(answer);
        return 1;
    }
    answer[ansCount] = '\x00';
    printf("%s", answer);
    free(x);
    free(answer);
    return 0;
}
