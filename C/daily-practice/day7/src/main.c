#include <stdio.h> 

int main() {
    int n = 0, x = 0, y = 0, z = 0;
    int input1 = 0, input2 = 0, input3 = 0; 
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%d %d %d", &input1, &input2, &input3);
        x += input1;
        y += input2;
        z += input3;
    }
    if (x == 0 && y == 0 && z == 0) {
        printf("YES");
    } else {
        printf("NO");
    }
}
