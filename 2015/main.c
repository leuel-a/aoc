#include <stdio.h>
#include <stdlib.h>

#define min(a, b)       ((a) < (b) ? (a) : (b))
#define ZERO_CHAR_ASCII '0'
#define NINE_CHAR_ASCII '9'

void solutionPartOne(FILE* file) {
    long  size;
    char* buffer;

    // (info) find the size of the file
    fseek(file, 0, SEEK_END);
    size = ftell(file);
    rewind(file);

    buffer = malloc(size + 1);

    fread(buffer, 1, size, file);
    buffer[size] = '\0';
}

int main() {
    FILE* file;

    file = fopen("input.txt", "r");

    if (file == NULL) {
        printf("(error) unable to open file");
        return (1);
    }

    solutionPartOne(file);
    return (0);
}
