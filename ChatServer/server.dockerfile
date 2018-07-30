FROM microsoft/dotnet:latest

RUN cd home
RUN git clone https://github.com/dmitrymatveev-work/chat-using-grpc.git
RUN cd ./chat-using-grpc/ChatServer/ChatServer
RUN dotnet build --configuration Release -o ./bin

EXPOSE 50051/tcp

ENTRYPOINT dotnet /home/chat-using-grpc/ChatServer/ChatServer/bin/ChatServer.dll
