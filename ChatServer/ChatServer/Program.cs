using Grpc.Core;
using System;
using System.Threading.Tasks;

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
                Ports = { new ServerPort(string.Empty, port, ServerCredentials.Insecure) }
            };
            server.Start();

            Console.WriteLine("Chat server listening on port " + port);
            Task.Delay(-1).Wait();

            server.ShutdownAsync().Wait();
        }
    }
}
