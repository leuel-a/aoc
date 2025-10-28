// Problem Link -> https://adventofcode.com/2016/day/3

public class Program
{
    static string INPUT_FILE = "input.txt";

    public static void partOne()
    {
        int count = 0;
        using var reader = new StreamReader(INPUT_FILE);

        while (reader.ReadLine() is string line)
        {
            var numbers = line.Split(' ', StringSplitOptions.RemoveEmptyEntries).Select(int.Parse).ToArray();
            if (IsValidTriangle(numbers))
                count++;
        }

        Console.WriteLine($"\tPart One Solution: {count}");
    }

    public static void partTwo()
    {
        var numbers = new List<int[]>();
        using var reader = new StreamReader(INPUT_FILE);

        while (reader.ReadLine() is string line)
            numbers.Add(line.Split(' ', StringSplitOptions.RemoveEmptyEntries).Select(int.Parse).ToArray());

        var count = 0;
        for (int i = 2; i < numbers.Count; i += 3)
        {
            for (int j = 0; j < 3; j++)
            {
                if (IsValidTriangle(new[] { numbers[i][j], numbers[i - 1][j], numbers[i - 2][j] }))
                    count++;

                // Console.WriteLine($"{numbers[i - 2][j]}, {numbers[i - 1][j]}, {numbers[i][j]} => Count {count}");
            }
        }
        Console.WriteLine($"\tPart Two Solution: {count}");
    }

    private static bool IsValidTriangle(int[] sides)
    {
        return (sides[0] + sides[1] > sides[2]) && (sides[0] + sides[2] > sides[1]) && (sides[1] + sides[2] > sides[0]);
    }

    private static void Main(string[] args)
    {
        Console.WriteLine("AOC 2016 - Problem 3 Solutions");
        partOne();
        partTwo();
    }
}

