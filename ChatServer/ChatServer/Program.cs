using Grpc.Core;
using System;

namespace ChatServer
{
    class Program
    {
        static void Main(string[] args)
        {
            const int port = 50051;

            var server = new Server
            {
                Services = { Chat.BindService(new ChatImpl()) },
                Ports = { new ServerPort("localhost", port, ServerCredentials.Insecure) }
            };
            server.Start();

            Console.WriteLine("Greeter server listening on port " + port);
            Console.WriteLine("Press any key to stop the server...");
            Console.ReadKey();

            server.ShutdownAsync().Wait();
        }
    }
}
