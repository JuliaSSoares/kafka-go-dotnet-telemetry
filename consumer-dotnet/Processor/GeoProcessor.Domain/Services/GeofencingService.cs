using NetTopologySuite.Geometries;

namespace GeoProcessor.Domain.Services
{
    public class GeofencingService : IGeofencingService
    {
        private readonly Polygon _riskArea = new(new LinearRing([
            new Coordinate(-46.63, -23.55),
            new Coordinate(-46.64, -23.55),
            new Coordinate(-46.64, -23.56),
            new Coordinate(-46.63, -23.56),
            new Coordinate(-46.63, -23.55)
        ]));

        public bool IsInRiskArea(double lat, double lon)
        {
            var point = new Point(lon, lat);
            return _riskArea.Contains(point);
        }
    }
}
