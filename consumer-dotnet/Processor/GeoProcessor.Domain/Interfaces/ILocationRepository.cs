using GeoProcessor.Domain.Entities;

namespace GeoProcessor.Domain.Interfaces
{
    public interface ILocationRepository
    {
        Task UpdateLocationAsync(EntregadorLocation location);
    }
}
