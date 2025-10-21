import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

import utils.java.*;

public class Main {
        public static final String INPUT_FILE = "input.txt";
        public static List<Pair> directions = new ArrayList<>();

        static {
                directions.add(new Pair(-1, 0));
                directions.add(new Pair(1, 0));
                directions.add(new Pair(0, 1));
                directions.add(new Pair(0, -1));
        };

        public static void main(String[] args) throws IOException {
                System.out.println(partOne());
        }

        public static int partOne() {
                int price = 0;
                List<List<String>> grid = readGridFromFile();

                for (int i = 0; i < grid.size(); i++) {
                        for (int j = 0; j < grid.get(0).size(); j++) {
                                var currentNodeValue = grid.get(i).get(j);
                                if (currentNodeValue != "-1") {
                                        var result = searchRegion(grid, i, j, currentNodeValue);
                                        price += (result.count * result.findPerimeter());
                                        // System.out.println(result + " >> " + currentNodeValue);
                                        // for (var row : grid) {
                                        // System.out.println(row);
                                        // }
                                }
                        }
                }

                return price;
        }

        public static void partTwo() {
                @SuppressWarnings("unused")
                List<List<String>> grid = readGridFromFile();
        }

        private static boolean inBound(List<List<String>> grid, int row, int column) {
                return (row >= 0 && row < grid.size()) && (column >= 0 && column < grid.get(row).size());
        }

        private static SearchRegionResult searchRegion(List<List<String>> grid, int row, int column,
                        String currentNodeValue) {
                Queue<Pair> queue = new LinkedList<>();
                SearchRegionResult result = new SearchRegionResult(0, 0);

                queue.add(new Pair(row, column));
                result.increment(new Pair(row, column));
                grid.get(row).set(column, "-1");

                while (queue.size() != 0) {
                        var currentNode = queue.remove();

                        for (var direction : directions) {
                                Pair newNode = new Pair(currentNode.row + direction.row,
                                                currentNode.column + direction.column);

                                if (inBound(grid, newNode.row, newNode.column)
                                                && currentNodeValue.equals(grid.get(newNode.row).get(newNode.column))) {
                                        queue.add(newNode);
                                        result.increment(newNode);
                                        grid.get(newNode.row).set(newNode.column, "-1");
                                }
                        }
                }
                return result;
        }

        public static List<List<String>> readGridFromFile() {
                FileReader fileReader;
                List<List<String>> grid = new ArrayList<>();

                try {
                        fileReader = new FileReader(INPUT_FILE);
                        BufferedReader bufferedReader = new BufferedReader(fileReader);

                        String line;
                        while ((line = bufferedReader.readLine()) != null) {
                                grid.add(Arrays.asList(line.split("")));
                        }

                        bufferedReader.close();
                        return grid;
                } catch (FileNotFoundException e) {
                        printError(e.toString());
                        printError("(ERROR) unable to parse the grid from the file");
                } catch (IOException e) {
                        printError(e.toString());
                        printError("(ERROR) unable to parse the grid from the file");
                }
                return grid;
        }

        public static void printError(String message) {
                final String RED = "\u001B[31m";
                final String RESET = "\u001B[0m";
                System.err.println(RED + message + RESET);
        }
}

class SearchRegionResult {
        int count;
        List<Pair> nodes = new ArrayList<>();
        public static List<Pair> directions = new ArrayList<>();

        static {
                directions.add(new Pair(-1, 0));
                directions.add(new Pair(1, 0));
                directions.add(new Pair(0, 1));
                directions.add(new Pair(0, -1));
        };

        public SearchRegionResult(int count, int perimeter) {
                this.count = count;
        }

        public void increment(Pair node) {
                this.nodes.add(node);
                count++;
        }

        public int findPerimeter() {
                var perimeter = 0;

                for (var node : nodes) {
                        var orthogonalNeighbours = 0;
                        for (var direction : directions) {
                                var candidate = new Pair(node.row + direction.row, node.column + direction.column);
                                if (nodes.contains(candidate)) {
                                        orthogonalNeighbours++;
                                }
                        }
                        perimeter += (4 - orthogonalNeighbours);
                }

                return perimeter;
        }

        @Override
        public String toString() {
                return "Count: " + count + " Perimeter: " + findPerimeter();
        }
}
