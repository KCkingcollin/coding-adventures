#include <stdio.h> 

int main() {
    int zeros = 0;
    int ones = 0;
    char situ[101];
    scanf("%100s", situ);
    for (int i = 0; i < 100; i++) {
        if (situ[i] == '0') {
            zeros++;
            ones = 0;
        } else if (situ[i] == '1') {
            ones++;
            zeros = 0;
        }
        if (zeros >= 7 || ones >= 7) {
            printf("YES\n");
            return 0;
        }
    }
    printf("NO\n");
}
