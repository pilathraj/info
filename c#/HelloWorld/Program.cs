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
        }
    }
}
