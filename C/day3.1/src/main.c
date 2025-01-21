#include <stdio.h> 

int main() {
    int n;
    int x = 0;
    char line[4];
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%3s", line);
        for (int j = 0; j < 3; j++) {
            if (line[j] == '+') {
                x++;
                break;
            }
            if (line[j] == '-') {
                x--;
                break;
            }
        }
    }
    printf("%d\n", x);
}
