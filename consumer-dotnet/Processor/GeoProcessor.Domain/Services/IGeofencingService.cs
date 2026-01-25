namespace GeoProcessor.Domain.Services
{
    public interface IGeofencingService
    {
        bool IsInRiskArea(double lat, double lon);
    }
}