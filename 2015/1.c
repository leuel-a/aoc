#include <stdio.h>

void solutionPartOne(FILE* file) {
    int  floor = 0;
    char character;

    while ((character = fgetc(file)) != EOF) {
        switch (character) {
            case '(': floor++; break;
            case ')': floor--; break;
            default: break;
        }
    }

    printf("%d\n", floor);
}

void solutionPartTwo(FILE* file) {
    int  floor = 0, position = 0;
    char character;

    while ((character = fgetc(file)) != EOF) {
        switch (character) {
            case '(': floor++; break;
            case ')': floor--; break;
            default: break;
        }

        position++;
        if (floor == -1) {
            break;
        }
    }

    printf("%d\n", position);
}

int main() {
    FILE* file;

    file = fopen("input.txt", "r");

    if (file == NULL) {
        printf("Could not open file\n");
        return (1);
    }

    solutionPartTwo(file);

    fclose(file);

    return (0);
}
