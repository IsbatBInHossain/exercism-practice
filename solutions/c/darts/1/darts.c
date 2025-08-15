#include "darts.h"
#include <math.h>

int score(coordinate_t pos){
    float total = pos.x * pos.x + pos.y * pos.y;
    float r = sqrtf(total);

    if (r <= 1.0F){
        return 10;
    } else if (r <= 5.0F){
        return 5;
    } else if (r <= 10.0F){
        return 1;
    }
    return 0;
}
