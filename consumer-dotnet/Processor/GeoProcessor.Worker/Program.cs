using GeoProcessor.Worker;
using GeoProcessor.Application;
using GeoProcessor.Infrastructure;

var builder = Host.CreateApplicationBuilder(args);

builder.Services.AddInfrastructure(builder.Configuration);
builder.Services.AddApplication( );
builder.Services.AddHostedService<Worker>();

var host = builder.Build();
host.Run();
