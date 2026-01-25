namespace GeoProcessor.Domain.Entities
{
    public class EntregadorLocation
    {
        private EntregadorLocation() { }

        public EntregadorLocation(string entregadorId, double lat, double lon, long unixTimestamp)
        {
            EntregadorId = entregadorId;
            Latitude = lat;
            Longitude = lon;
            Timestamp = unixTimestamp;
            CreatedAt = DateTime.UtcNow;
        }

        public int Id { get; private set; }
        public string EntregadorId { get; private set; }
        public double Latitude { get; private set; }
        public double Longitude { get; private set; }
        public long Timestamp { get; private set; }
        public DateTime CreatedAt { get; private set; }
    }
}