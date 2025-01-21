#include <stdio.h> 

int main() {
    int n;
    int k;
    int kScore;
    int p = 0;
    int score[50];
    scanf("%d %d", &n, &k);
    for (int i = 0; i < n; i++) {
        scanf("%d", &score[i]);
    }
    kScore = score[k-1];
    for (int i = 0; i < n; i++) {
        if (score[i] >= kScore && score[i] > 0) {
            p++;
        }
    }
    printf("%d\n", p);
}
