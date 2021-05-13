using System;

namespace HelloWorld
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
            Console.Write("Enter your name: ");
            string name = Console.ReadLine();
            Console.WriteLine("Hello {0}!!", name);
            int i = 10;
            string s = "Pilathraj";
            string s1 = null;
            Console.WriteLine("{0}, {1}", i, s);
            Console.WriteLine(s1??"s1");
            s1 ??= "test";
            Console.WriteLine(s1);

        }
    }
}
/*
/d/cs/st/Overview/HelloWorld
$ dotnet run
Hello World!
Enter your name: p
Hello p!!    
10, Pilathraj
s1
test
*/

