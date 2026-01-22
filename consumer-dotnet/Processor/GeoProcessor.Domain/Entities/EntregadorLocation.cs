using GeoProcessor.Domain.ValueObjects;

namespace GeoProcessor.Domain.Entities
{
    public class EntregadorLocation
    {
        public string Id { get; private set; }
        public Coordinate Position { get; private set; }
        public DateTime CapturedAt { get; private set; }

        public EntregadorLocation(string id, double lat, double lon, long unixTimestamp)
        {
            Id = id;
            Position = new Coordinate(lat, lon);
            CapturedAt = DateTimeOffset.FromUnixTimeSeconds(unixTimestamp).UtcDateTime;
        }
    }
}
