#include "grains.h"
#include <math.h>

uint64_t square(uint8_t index){
    return (uint64_t)pow(2,index-1);
}

uint64_t total(void){
    uint64_t sum = 0;
    for (uint8_t i = 1; i <=64; i++){
        sum += square(i);
    }
    return sum;
}