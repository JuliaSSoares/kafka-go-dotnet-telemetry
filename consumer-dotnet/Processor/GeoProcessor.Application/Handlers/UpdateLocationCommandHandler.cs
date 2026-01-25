using GeoProcessor.Application.Commands;
using GeoProcessor.Domain.Interfaces;
using MediatR;
using Microsoft.Extensions.Logging;

namespace GeoProcessor.Application.Handlers
{
    public class UpdateLocationCommandHandler(ILogger<UpdateLocationCommandHandler> logger, ILocationRepository repository) : IRequestHandler<UpdateLocationCommand>
    {

        private readonly ILocationRepository _repository = repository;
        private readonly ILogger<UpdateLocationCommandHandler> _logger = logger;

        public async Task Handle(UpdateLocationCommand request, CancellationToken ct)
        { 
          await _repository.UpdateLocationAsync(new(
                    request.EntregadorId,
                    request.Longitude,
                    request.Latitude,
                    request.TimeStamp
                    ));

            _logger.LogInformation("Location updated for EntregadorId: {EntregadorId}", request.EntregadorId);

        } 
    }
}
