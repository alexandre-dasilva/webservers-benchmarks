FROM mcr.microsoft.com/dotnet/aspnet:5.0 AS base
WORKDIR /app
EXPOSE 443
EXPOSE 80

FROM mcr.microsoft.com/dotnet/sdk:5.0 AS build
WORKDIR /src
COPY ["BaseApiPerfTests.csproj", "./"]
RUN dotnet restore "BaseApiPerfTests.csproj"
COPY . .
WORKDIR "/src/"
RUN dotnet build "BaseApiPerfTests.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "BaseApiPerfTests.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "BaseApiPerfTests.dll"]


# FROM mcr.microsoft.com/dotnet/runtime-deps:5.0 AS base
# WORKDIR /app
# EXPOSE 80
# ENV DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1 \
#     DOTNET_EnableDiagnostics=0 \
#     DOTNET_RUNNING_IN_CONTAINER=true \
#     ASPNETCORE_URLS=http://+:80

# FROM mcr.microsoft.com/dotnet/sdk:5.0 AS build
# WORKDIR /src
# COPY ["BaseApiPerfTests.csproj", "./"]
# RUN dotnet restore "BaseApiPerfTests.csproj"
# COPY . .
# WORKDIR "/src/"
# RUN dotnet build "BaseApiPerfTests.csproj" -c Release -o /app/build

# FROM build AS publish

# RUN dotnet publish "BaseApiPerfTests.csproj" \
#   --runtime linux-x64 \
#   --self-contained true \
#   /p:PublishTrimmed=true \
#   /p:PublishSingleFile=true \
#   -c Release \
#   -o /app/publish

# FROM base AS final
# WORKDIR /app
# COPY --from=publish /app/publish .
# ENTRYPOINT ["./BaseApiPerfTests", "--urls", "http://0.0.0.0:80"]
