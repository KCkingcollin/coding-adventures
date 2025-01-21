#include <stdio.h> 

int main() {
    char p[101] = {};
    char inst[3] = {'H', 'Q', '9'};
    scanf("%100s", p);
    for (int i = 0; i < 101; i++) {
        for (int j = 0; j < 3; j++) {
            if (p[i] == inst[j]) {
                printf("YES");
                return 0;
            }
        }
    }
    printf("NO");
}
