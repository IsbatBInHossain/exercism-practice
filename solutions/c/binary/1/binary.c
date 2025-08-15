#include "binary.h"
#include <string.h>
#include <math.h>


int convert(const char *input){
    int l = (int)strlen(input);
    int dec = 0;
    for (int i = 0; i <l ; i++){
        if (input[i] == '1'){
            dec += (int)pow((float)2, (float)l-i-1);
        } else if (input[i] == '0'){
            continue;
        } else {
            return -1;
        }
    }
    return dec;
}