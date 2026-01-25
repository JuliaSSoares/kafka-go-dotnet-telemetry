using Confluent.Kafka;
using GeoProcessor.Application.Commands;
using GeoProcessor.Application.Dtos;
using MediatR;
using System.Text.Json;

namespace GeoProcessor.Worker
{
    public class Worker(ILogger<Worker> logger, IConfiguration configuration, IServiceProvider serviceProvider) : BackgroundService
    {
        private readonly ILogger<Worker> _logger = logger;
        private readonly IConfiguration _configuration = configuration;
        private readonly IServiceProvider _serviceProvider = serviceProvider;

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            var config = new ConsumerConfig
            {
                BootstrapServers = _configuration["Kafka:BootstrapServers"],
                GroupId = _configuration["Kafka:GroupId"],
                AutoOffsetReset = AutoOffsetReset.Earliest
            };

            using var consumer = new ConsumerBuilder<Ignore, string>(config).Build();
            var topicName = _configuration["Kafka:TopicName"] ?? "telemetria-topic";
            consumer.Subscribe(topicName);

            while (!stoppingToken.IsCancellationRequested)
            {
                try
                {
                    var result = consumer.Consume(stoppingToken);
                    var message = JsonSerializer.Deserialize<TelemetriaInboundDto>(result.Message.Value);

                    if (message != null)
                    {
                        using var scope = _serviceProvider.CreateScope();
                        var mediator = scope.ServiceProvider.GetRequiredService<IMediator>();

                        var command = new UpdateLocationCommand(
                            message.EntregadorId,
                            message.Longitude,
                            message.Latitude,
                            message.Timestamp
                        );

                        _logger.LogInformation($"Disparando comando - atualizar localização entregado-id: {message.EntregadorId}");
                        await mediator.Send(command, stoppingToken);
                    }
                }
                catch (Exception ex)
                {
                    _logger.LogError($"Erro ao processar Kafka: {ex.Message}");
                }
            }
        }
    }
}
