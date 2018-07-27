using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Grpc.Core;

namespace ChatServer
{
    public class ChatImpl : Chat.ChatBase
    {
        private readonly ConcurrentDictionary<string, IServerStreamWriter<Post>> _users = new ConcurrentDictionary<string, IServerStreamWriter<Post>>();

        public override Task<IntroResponse> Introduce(IntroRequest request, ServerCallContext context)
        {
            if (_users.ContainsKey(request.Username))
            {
                return Task.FromResult(new IntroResponse { Message = "User with such name already exists." });
            } else
            {
                _users.TryAdd(request.Username, null);
                return Task.FromResult(new IntroResponse { Message = "OK" });
            }
        }

        public override async Task Connect(IAsyncStreamReader<Post> request, IServerStreamWriter<Post> response, ServerCallContext context)
        {
            while(await request.MoveNext())
            {
                var post = request.Current;
                _users[post.Username] = response;
                await Post(request.Current);
            }
        }

        private async Task Post(Post post)
        {
            foreach(var user in _users)
            {
                var response = user.Value;
                if (response == null)
                {
                    continue;
                }
                await response.WriteAsync(post);
            }
        }
    }
}
