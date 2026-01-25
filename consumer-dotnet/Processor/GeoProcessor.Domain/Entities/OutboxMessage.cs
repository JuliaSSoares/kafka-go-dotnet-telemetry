namespace GeoProcessor.Domain.Entities
{
    public class OutboxMessage(string payload, string status)
    {
        public int Id { get; private set; }
        public string Payload { get; private set; } = payload;
        public string Status { get; private set; } = status;
    }
}
