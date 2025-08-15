#include "queen_attack.h"
#include <stdlib.h>

bool valid_position(position_t p1, position_t p2) {
    if (p1.row >= 8 || p1.column >= 8){
        return false;
    } else if (p2.row >= 8 || p2.column >= 8){
        return false;
    } else if ((p1.row == p2.row) && (p1.column == p2.column)){
        return false;
    }
    return true;
}

bool on_diagonal(position_t p1, position_t p2){
    return abs(p1.row - p2.row) == abs(p1.column - p2.column);
    
}

attack_status_t can_attack(position_t q1, position_t q2) {
    if (!valid_position(q1, q2)) {
        return INVALID_POSITION;
    }
    else if ((q1.row == q2.row) || (q1.column == q2.column)){
        return CAN_ATTACK;
    } else if (on_diagonal(q1, q2)){
        return CAN_ATTACK;
    }
    return CAN_NOT_ATTACK;
}
