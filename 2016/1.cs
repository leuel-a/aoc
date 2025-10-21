internal class DayOne
{
    static string NORTH = "N", EAST = "E", WEST = "W", SOUTH = "S";

    static string[] right = { NORTH, EAST, SOUTH, WEST };
    static string[] left = { NORTH, WEST, SOUTH, EAST };

    static Dictionary<string, (int dx, int dy)> directions = new()
    {
        [NORTH] = (0, 1),
        [EAST] = (1, 0),
        [SOUTH] = (0, -1),
        [WEST] = (-1, 0),
    };

    private static void solutionPartOne(string content)
    {
        var curr = NORTH;
        (int x, int y) position = (0, 0);

        var movements = content.Split(',').Select(value => value.Trim()).ToArray();

        foreach (var movement in movements)
        {
            int index, nextIndex;
            var direction = movement[0].ToString();
            var count = int.Parse(movement.Substring(1));

            switch (direction)
            {
                case "R":
                    index = Array.IndexOf(right, curr);
                    nextIndex = (index + 1) % 4;

                    curr = right[nextIndex];

                    // Console.WriteLine($"Movement: {movement}, Current Direction: {right[index]}, Next Direction: {curr}, Directions: ({directions[curr].dx}, {directions[curr].dy})");
                    for (int i = 0; i < count; i++)
                    {
                        position = (position.x + directions[curr].dx, position.y + directions[curr].dy);
                        Console.WriteLine($"Position ({position.x}, {position.y})");
                    }

                    break;
                case "L":
                    index = Array.IndexOf(left, curr);
                    nextIndex = (index + 1) % 4;

                    curr = left[nextIndex];

                    // Console.WriteLine($"Movement: {movement}, Current Direction: {left[index]}, Next Direction: {curr}, Directions: ({directions[curr].dx}, {directions[curr].dy})");
                    for (int i = 0; i < count; i++)
                    {
                        position = (position.x + directions[curr].dx, position.y + directions[curr].dy);
                        Console.WriteLine($"Position ({position.x}, {position.y})");
                    }
                    break;
                default:
                    Console.WriteLine($"Unrecogizable Character: {direction}");
                    break;
            }

            Console.WriteLine();
        }

        // Console.WriteLine($"Position ({position.x}, {position.y})");
        Console.WriteLine(Math.Abs(position.x) + Math.Abs(position.y));

    }

    private static void solutionPartTwo(string content)
    {
        var curr = NORTH;
        (int x, int y) position = (0, 0);
        var visited = new HashSet<(int dx, int dy)>();

        var movements = content.Split(',').Select(value => value.Trim()).ToArray();

        foreach (var movement in movements)
        {
            int index, nextIndex;
            var direction = movement[0].ToString();
            var count = int.Parse(movement.Substring(1));

            switch (direction)
            {
                case "R":
                    index = Array.IndexOf(right, curr);
                    nextIndex = (index + 1) % 4;

                    curr = right[nextIndex];

                    Console.WriteLine($"Movement: {movement}, Current Direction: {right[index]}, Next Direction: {curr}, Directions: ({directions[curr].dx}, {directions[curr].dy})");

                    for (int i = 0; i < count; i++)
                    {
                        position = (position.x + directions[curr].dx, position.y + directions[curr].dy);

                        if (visited.Contains(position))
                        {
                            Console.WriteLine($"Position: ({position.x}, {position.y})\n");
                            Console.WriteLine(Math.Abs(position.x) + Math.Abs(position.y));
                            return;
                        }

                        visited.Add(position);

                        Console.WriteLine($"Position ({position.x}, {position.y})");
                    }

                    break;
                case "L":
                    index = Array.IndexOf(left, curr);
                    nextIndex = (index + 1) % 4;

                    curr = left[nextIndex];

                    Console.WriteLine($"Movement: {movement}, Current Direction: {left[index]}, Next Direction: {curr}, Directions: ({directions[curr].dx}, {directions[curr].dy})");

                    for (int i = 0; i < count; i++)
                    {
                        position = (position.x + directions[curr].dx, position.y + directions[curr].dy);

                        if (visited.Contains(position))
                        {
                            Console.WriteLine($"Position: ({position.x}, {position.y})\n");
                            Console.WriteLine(Math.Abs(position.x) + Math.Abs(position.y));
                            return;
                        }

                        Console.WriteLine($"Position ({position.x}, {position.y})");
                    }
                    break;
                default:
                    Console.WriteLine($"Unrecogizable Character: {direction}");
                    break;
            }

            Console.WriteLine();
        }
    }

    private static void run(string[] args)
    {

        string path = "input.txt";

        if (!File.Exists(path))
        {
            Console.WriteLine("(error) unable to open the input file");
            Environment.Exit(1);
        }

        string content = File.ReadAllText(path).Trim().Replace(Environment.NewLine, "");
    }
}
