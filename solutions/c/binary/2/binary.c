#include "binary.h"
#include <string.h>
#include <math.h>


int convert(const char *input){
    int dec = 0;
    for (int i = 0; input[i] != '\0'; i++){
        if (input[i] == '1'){
            dec = (dec << 1) | 1;
        } else if (input[i] == '0'){
            dec <<= 1;
        } else {
            return INVALID;
        }
    }
    return dec;
}