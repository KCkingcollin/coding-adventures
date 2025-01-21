#include <stdio.h> 

#define max(a, b) ((a) > (b) ? (a) : (b))

int main() {
    int n = 0;
    int s = 0;
    int nums[4] = {};
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%d", &s);
        nums[s-1]++;
    }
    if (nums[2] > 0 || (nums[1] & 1) == 1) {
        nums[0] = max(nums[0]-nums[2], 0);
        nums[0] = max(nums[0]-(nums[1] & 1)*2, 0);
        nums[0] = (nums[0]+4-1)/4;
    } else {
        nums[0] = (nums[0]+4-1)/4;
    }
    nums[1] = (nums[1]+2-1)/2;
    printf("%d", nums[0]+nums[1]+nums[2]+nums[3]);
}
