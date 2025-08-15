#include "rna_transcription.h"
#include <string.h>
#include <stdlib.h>


char *to_rna(const char *dna){
    int l = (int)strlen(dna);
    char *rna = malloc(l+1);
    for (int i = 0; i < l; i++){
        switch (dna[i]){
            case 'G':
                rna[i] = 'C';
                break;
            case 'C':
                rna[i] = 'G';
                break;
            case 'T':
                rna[i] = 'A';
                break;
            case 'A':
                rna[i] = 'U';
                break;
            default:
                break;
        }
    }
    rna[l] = '\0';
    return rna;
    
}