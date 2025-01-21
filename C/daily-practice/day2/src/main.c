#include <stdio.h> 

int main() {
    int n;
    int wordSize;
    char str[101];
    char result[10];
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%100s", str);
        wordSize = 0;
        for (int j = 0; j < 100; j++) {
            if (str[j] == '\0') {
                break;
            }
            wordSize++;
        }
        if (wordSize > 10) {
            printf("%c%d%c\n", str[0], wordSize-2, str[wordSize-1]);
        } else {
            printf("%s\n", str);
        }
    }
}
