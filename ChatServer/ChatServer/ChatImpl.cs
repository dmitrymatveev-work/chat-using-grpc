using Grpc.Core;
using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

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
            }
            else
            {
                _users.TryAdd(request.Username, null);
                return Task.FromResult(new IntroResponse { Message = "OK" });
            }
        }

        public override async Task Connect(IAsyncStreamReader<Post> request, IServerStreamWriter<Post> response, ServerCallContext context)
        {
            string username;
            if (await request.MoveNext())
            {
                var post = request.Current;
                username = post.Username;
                _users[username] = response;
                await Post(post);
            }
            else
            {
                return;
            }

            while (await request.MoveNext())
            {
                await Post(request.Current);
            }

            _users.TryRemove(username, out IServerStreamWriter<Post> r);
            await Post(new Post { Username = username, Message = "Quitted." });
        }

        private async Task Post(Post post)
        {
            foreach (var user in _users.ToArray())
            {
                var response = user.Value;
                if (response == null)
                {
                    continue;
                }
                try
                {
                    await response.WriteAsync(post);
                }
                catch
                {
                    _users.TryRemove(user.Key, out IServerStreamWriter<Post> r);
                }
            }
        }
    }
}
