#include <stdio.h> 

int main() {
    int n;
    int result = 0;
    int noCounter;
    int nY;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        noCounter = 0;
        for (int j = 0; j < 3; j++) {
            nY = 0;
            scanf("%d", &nY);
            if (nY == 0) {
                noCounter++;
            }
        }
        if (noCounter < 2) {
            result++;
        }
    }
    printf("%d\n", result);
}
