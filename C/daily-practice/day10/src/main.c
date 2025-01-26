#include <math.h>
#include <stdio.h> 
#include <stdlib.h>
#include <string.h>

int maxX = 0;
int minX = 0;
int x[100001] = {};
int *y = 0;

int max(int a, int b) {
    return a > b ? a : b;
}

int p(int i) {
    if (i > maxX) {
        y[maxX] = p(maxX-1) + x[maxX];
        return y[maxX];
    }
    if (y[i] == 0 && i > minX) {
        y[i] = p(i-1) + x[i];
    }
    return y[i];
}


int main() {
    int n = 0;
    scanf("%d", &n);
    int num = 0;
    for (int i = 0; i < n; i++) {
        scanf("%d", &num);
        x[num]++;
        if (num > maxX) {maxX = num;}
        if (num < minX) {minX = num;}
    }
    y = (int *)calloc(maxX+1, sizeof(int));
    if (y == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }

    int q = 0;
    scanf("%d", &q);
    int ansArrSize = q*2;

    int m = 0;
    int ansCount = 0;
    char *answer = (char *)calloc(ansArrSize, sizeof(char));
    for (int i = 0; i < q; i++) {
        scanf("%d", &m);
        char ans[6] = {};
        sprintf(ans, "%d", p(m));
        int ansSize = max((int)log10((double)p(m)), 0) + 1;
        if (ansCount+ansSize+1 > ansArrSize) {
            ansArrSize *= 2;
            answer = (char *)realloc(answer, ansArrSize * sizeof(char));
            if (answer == NULL) {
                printf("Memory reallocation failed\n");
                free(answer);
                return 1;
            }
        }
        for (int j = 0; j < ansSize; j++) {
            answer[ansCount] = ans[j];
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
    free(answer);
    return 0;
}
