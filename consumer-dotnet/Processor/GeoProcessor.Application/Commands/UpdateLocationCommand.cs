using MediatR;

namespace GeoProcessor.Application.Commands
{
    public record UpdateLocationCommand(
        string EntregadorId,
        double Latitude,
        double Longitude,
        long TimeStamp
        ) : IRequest;
}
