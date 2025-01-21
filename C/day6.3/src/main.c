#include <ctype.h>
#include <stdio.h> 

int main() {
    char p[101] = {}, answer[200] = {};
    int notVowel = 0, ansIndex = 0;
    scanf("%s", p);
    char vowels[6] = {'A', 'E', 'I', 'O', 'U', 'Y'};
    for (int i = 0; i < 100 && p[i] != '\0'; i++) {
        notVowel = 1;
        for (int j = 0; j < 6; j++) {
            if (toupper(p[i]) == vowels[j]) {
                notVowel = 0;
            }
        }
        if (notVowel == 1) {
            answer[ansIndex] = '.';
            ansIndex++;
            answer[ansIndex] = tolower(p[i]);
            ansIndex++;
        }
    }
    printf("%s\n", answer);
}
