import java.util.*;
import java.io.*;

public class Main {
    public static final String INPUT_FILE = "input.txt";
    public static final String[] numbers = { "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };

    public static void printError(String message) {
        final String RED = "\u001B[31m";
        final String RESET = "\u001B[0m";

        System.err.println(RED + message + RESET);
    }

    public static void convert(String value) {
        int i = 0;
        StringBuilder builder = new StringBuilder();
        char[] valueArray = value.toCharArray();

        while (i < valueArray.length) {
            String candidate;
            char character = valueArray[i];

            if (character == 'o' && i + 2 < valueArray.length) {
                candidate = String.valueOf(Arrays.copyOfRange(valueArray, i, i + 3));
                if (candidate == "one") {
                    builder.append("1");
                    i += 3;
                    continue;
                }
            } else if (character == 't') {
                if (i + 2 < valueArray.length
                && String.valueOf(Arrays.copyOfRange(valueArray, i, i + 3)).equals("two")) {
                    builder.append("2");
                    i += 3;
                    continue;
                }
                if (i + 4 < valueArray.length && String.valueOf(Arrays.copyOfRange(valueArray, i, i + 5)) == "three") {
                    builder.append("3");
                    i += 5;
                    continue;
                }
            } else if (character == 'f' && i + 4 < valueArray.length) {
                if (String.valueOf(Arrays.copyOfRange(valueArray, i, i + 5)) == "four") {
                    builder.append("4");
                    i += 5;
                    continue;
                }

                if (String.valueOf(Arrays.copyOfRange(valueArray, i, i + 5)) == "five") {
                    builder.append("5");
                    i += 5;
                    continue;
                }
            } else if (character == 's') {
            } else if (character == 'e') {
            } else if (character == 'n') {
                if (String.valueOf(Arrays.copyOfRange(valueArray, i, i + 5)) == "nine") {
                    builder.append("9");
                    i += 5;
                    continue;
                }
            }
            i++;
        }
    }

    public static void partOne() {
        FileReader fileReader;

        try {
            fileReader = new FileReader(INPUT_FILE);
            BufferedReader bufferedReader = new BufferedReader(fileReader);

            String line;
            List<Integer> numbers = new ArrayList<>();
            while ((line = bufferedReader.readLine()) != null) {
                Boolean firstFound = false;
                int firstDigit = 0, lastDigit = 0;

                for (char character : line.toCharArray()) {
                    if (Character.isDigit(character)) {
                        if (firstFound == false) {
                            firstDigit = Character.getNumericValue(character);
                            lastDigit = Character.getNumericValue(character);
                            firstFound = true;
                        } else {
                            lastDigit = Character.getNumericValue(character);
                        }
                    }
                }

                numbers.add(firstDigit * 10 + lastDigit);
            }

            System.out.println(numbers.stream().reduce(0, Integer::sum));
            bufferedReader.close();
        } catch (FileNotFoundException e) {
            printError(e.toString());
            printError("(ERROR) unable to parse the grid from the file");
        } catch (IOException e) {
            printError(e.toString());
            printError("(ERROR) unable to parse the grid from the file");
        }
    }

    public static void partTwo() {
        FileReader fileReader;

        try {
            fileReader = new FileReader(INPUT_FILE);
            BufferedReader reader = new BufferedReader(fileReader);

            String line;
            while ((line = reader.readLine()) != null) {
                System.out.println("Line " + line);
                convert(line);
            }

            reader.close();
        } catch (FileNotFoundException e) {
        } catch (IOException e) {
        }
    }

    public static void main(String[] args) {
        partTwo();
    }
}
