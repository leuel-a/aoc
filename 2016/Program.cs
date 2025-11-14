using System.Text.RegularExpressions;

public class Program
{
    static char A_CHARACTER = 'a';
    static string INPUT_FILE = "input.txt";

    private static void partOne()
    {
        int realRoomsCount = 0;
        int sectorIdSum = 0;
        using var reader = new StreamReader(INPUT_FILE);

        while (reader.ReadLine() is string line)
        {
            int[] count = new int[26];
            foreach (var value in line.Split('-')[..^1])
            {
                foreach (var character in value.ToCharArray())
                    count[(int)character - (int)A_CHARACTER]++;
            }

            var match = Regex.Match(line.Split('-')[^1], @"(\d+)\[([a-z]+)\]");

            string sectorId = match.Groups[1].Value;
            string checksum = match.Groups[2].Value;

            // Console.WriteLine($"Input: {line} --> Sector ID: {sectorId} Checksum: {checksum}");

            var realRoom = true;
            for (int i = 1; i < checksum.Count(); i++)
            {
                var currIndex = (int)(checksum[i]) - (int)A_CHARACTER;
                var prevIndex = (int)(checksum[i - 1]) - (int)A_CHARACTER;

                // Console.WriteLine($"({checksum[i]}, {currIndex}) -- ({checksum[i - 1]}, {currIndex})");

                if (count[currIndex] > count[prevIndex])
                    realRoom = false;

                if ((count[currIndex] == count[prevIndex]) && (prevIndex > currIndex))
                    realRoom = false;

                if (realRoom == false)
                    break;
            }

            if (realRoom)
            {
                realRoomsCount++;
                sectorIdSum += int.Parse(sectorId);
            }
        }

        Console.WriteLine($"Real Rooms: {realRoomsCount}, Sector ID Sum: {sectorIdSum}");
    }

    public static void Main(string[] args)
    {
        Console.WriteLine("AOC 2016 - Problem 4 Solutions");
        partOne();
    }
}
