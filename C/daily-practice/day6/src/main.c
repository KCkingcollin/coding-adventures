#include <stdio.h> 

int main() {
    long n = 0;
    long k = 0;
    long num = 0;
    scanf("%ld", &n);
    n = (n+2)/2;
    scanf("%ld", &k);
    if (k > n) {
        num = (k-n)*2;
    } else {
        num = k*2-1;
    }
    printf("%ld", num);
}
