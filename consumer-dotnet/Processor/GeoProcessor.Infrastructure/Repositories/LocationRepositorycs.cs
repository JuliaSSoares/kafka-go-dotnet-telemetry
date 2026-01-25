using GeoProcessor.Domain.Entities;
using GeoProcessor.Domain.Interfaces;
using StackExchange.Redis;
using System.Text.Json;

namespace GeoProcessor.Infrastructure.Repositories
{
    public class LocationRepository : ILocationRepository
    {
        private readonly GeoDbContext _context;
        private readonly IDatabase _redis;

        public LocationRepository(GeoDbContext context, IConnectionMultiplexer redis)
        {
            _context = context;
            _redis = redis.GetDatabase();
        }

        public async Task UpdateLocationAsync(EntregadorLocation location)
        {
            _context.Telemetrias.Update(location);
            await _context.SaveChangesAsync();

            var key = $"location:{location.Id}";
            var data = JsonSerializer.Serialize(location);
            await _redis.StringSetAsync(key, data, TimeSpan.FromMinutes(30));
        }
    }
}
