using GeoProcessor.Domain.Interfaces;
using GeoProcessor.Infrastructure.Repositories;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using StackExchange.Redis;

namespace GeoProcessor.Infrastructure
{
    public static class DependencyInjection
    {
        public static IServiceCollection AddInfrastructure(this IServiceCollection services, IConfiguration configuration)
        {
            var connectionString = configuration.GetConnectionString("PostgresSQL");
            services.AddDbContext<GeoDbContext>(options =>
                options.UseNpgsql(connectionString));

            var redisConnection = configuration.GetSection("Redis:ConnectionString").Value;
            services.AddSingleton<IConnectionMultiplexer>(sp =>
                ConnectionMultiplexer.Connect(redisConnection!));

            services.AddScoped<ILocationRepository, LocationRepository>();

            return services;
        }
    }
}
