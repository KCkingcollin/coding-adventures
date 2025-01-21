#include <stdio.h>

int main() {
    int num;
    scanf("%d", &num);
    if (((num & 1) == 0) && num > 3) {
        printf("YES\n");
        return 0;
    }
    printf("NO\n");
}
