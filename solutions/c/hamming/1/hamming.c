#include "hamming.h"
#include <string.h>

int compute(const char *lhs, const char *rhs){
    int distance = 0;
    int ll = (int)strlen(lhs);
    int rl = (int)strlen(rhs);
    if (ll != rl){
        return -1;
    }
    
    for (int i = 0; i < ll; i++){
        if (lhs[i] != rhs[i]){
            distance++;
        }
    }
    return distance;
}
