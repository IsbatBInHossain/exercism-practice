#include "eliuds_eggs.h"
#include <stdio.h>


int egg_count(int code){
    int count = 0;
    while (code > 0){
        count += code % 2;
        code = code / 2;
    }
    return count;
}