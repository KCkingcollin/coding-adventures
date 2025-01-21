#include <stdio.h> 

int main() {
    long n = 0, m = 0, a = 0, answer = 0;
    scanf("%ld %ld %ld", &n, &m, &a);
    answer = ((n+a-1)/a) * ((m+a-1)/a);
    printf("%ld", answer);
}
