using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ChatServer
{
    public class ChatImpl : Chat.ChatBase
    {
        private readonly ConcurrentDictionary<string, string> _users = new ConcurrentDictionary<string, string>();

        public override Task<IntroResponse> Introduce(IntroRequest request, Grpc.Core.ServerCallContext context)
        {
            if (_users.ContainsKey(request.Username))
            {
                return Task.FromResult(new IntroResponse { Message = "User with such name already exists." });
            } else
            {
                _users.TryAdd(request.Username, request.Username);
                return Task.FromResult(new IntroResponse { Message = "OK" });
            }
        }
    }
}
