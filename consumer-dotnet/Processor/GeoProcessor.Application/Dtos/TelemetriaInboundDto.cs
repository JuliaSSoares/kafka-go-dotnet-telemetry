using System.Text.Json.Serialization;

namespace GeoProcessor.Application.Dtos
{
    public record TelemetriaInboundDto(
            [property: JsonPropertyName("id")] int Id,
            [property: JsonPropertyName("entregador_id")] string EntregadorId,
            [property: JsonPropertyName("lat")] double Latitude,
            [property: JsonPropertyName("long")] double Longitude,
            [property: JsonPropertyName("timestamp")] long Timestamp
        );

}
