using System.Text;

internal class DayTwo
{
    static string PATH = "input.txt";

    static int[,] KEYPAD = new int[3, 3]{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };

    static int[,] CONFERENCE_KEYPAD = new int[5, 5]{
        {0,   0,   1,   0,   0},
        {0,   2,   3,   4,   0},
        {5,   6,   7,   8,   9},
        {0,  -1,  -2,  -3,   0},
        {0,   0,  -4,   0,   0}
    };

    private static readonly Dictionary<char, (int row, int col)> moves = new() {
        { 'U', (-1, 0) },
        { 'D', (1, 0) },
        { 'L', (0, -1)},
        { 'R', (0, 1)}
    };

    private static bool InBound(int[,] grid, int row, int col)
    {
        return (0 <= row && row < grid.GetLength(0)) && (0 <= col && col < grid.GetLength(1));
    }

    private static bool ConferenceInBound(int[,] grid, int row, int col)
    {
        if (!DayTwo.InBound(grid, row, col))
            return false;

        return grid[row, col] != 0;
    }

    private static void PartOne(string content)
    {
        string[] instructions = content.Split('\n');
        var result = new StringBuilder();

        (int row, int col) position = (1, 1);
        foreach (var instruction in instructions)
        {
            Console.WriteLine($"Current Instruction: {instruction}");
            for (int i = 0; i < instruction.Length; i++)
            {
                Console.WriteLine($"\tCurrent Movement: {instruction[i]}");
                if (moves.TryGetValue(instruction[i], out var move))
                {
                    var row = position.row + move.row;
                    var col = position.col + move.col;

                    if (DayTwo.InBound(KEYPAD, row, col))
                    {
                        position = (row, col);
                    }
                }
            }
            result.Append(KEYPAD[position.row, position.col].ToString());
            Console.WriteLine($"Current Passcode: {result}");
        }

        Console.WriteLine(result.ToString());
    }

    private static void PartTwo(string content)
    {
        string[] instructions = content.Split('\n');
        var result = new StringBuilder();

        (int row, int col) position = (2, 0);
        foreach (var instruction in instructions)
        {
            // Console.WriteLine($"Current Instruction: {instruction}");
            for (int i = 0; i < instruction.Length; i++)
            {
                // Console.WriteLine($"\tCurrent Movement: {instruction[i]}");
                if (moves.TryGetValue(instruction[i], out var move))
                {
                    var row = position.row + move.row;
                    var col = position.col + move.col;

                    if (DayTwo.ConferenceInBound(CONFERENCE_KEYPAD, row, col))
                    {
                        position = (row, col);
                    }
                }
            }
            var A_CHARACTER = 'A';
            var value = CONFERENCE_KEYPAD[position.row, position.col];

            // result.Append(value > 0 ? value.ToString() : ((int)A_CHARACTER - Math.Abs(value) - 1).ToString());

            if (value > 0)
            {
                result.Append(value.ToString());
            }
            else
            {
                var asciiValue = (int)A_CHARACTER + (Math.Abs(value) - 1);
                result.Append(((char)asciiValue).ToString());
            }
            // Console.WriteLine($"Current Passcode: {result}");
        }

        Console.WriteLine(result.ToString());
    }

    public static void run(string[] args)
    {
        if (!File.Exists(PATH))
        {
            Console.WriteLine("(error) unable to open the input file");
            Environment.Exit(1);
        }

        string content = File.ReadAllText(PATH).Trim();
        DayTwo.PartTwo(content);
    }
}
