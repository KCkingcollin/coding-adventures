#include <stdio.h> 

int main() {
    char hello[] = "hello", s[101] = {};
    int hIndex = 0;
    scanf("%100s", s);
    printf("%s\n%s", s, hello);
    for (int i = 0; s[i] != '\0'; i++) {
        if (s[i] == hello[hIndex]) {
            hIndex++;
        }
        if (hIndex > 4) {
            printf("YES\n");
            return 0;
        }
    }
    printf("NO\n");
    return 0;
}
