using Microsoft.EntityFrameworkCore;
using GeoProcessor.Domain.Entities;

namespace GeoProcessor.Infrastructure
{
    public class GeoDbContext : DbContext
    {
        public GeoDbContext(DbContextOptions<GeoDbContext> options) : base(options) { }

        public DbSet<EntregadorLocation> Telemetrias { get; set; }
        public DbSet<OutboxMessage> OutboxMessages { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<EntregadorLocation>(entity =>
            {
                entity.ToTable("telemetrias");
                entity.HasKey(e => e.Id);
                entity.Property(e => e.Id).HasColumnName("id");
                entity.Property(e => e.EntregadorId).HasColumnName("entregador_id");
                entity.Property(e => e.Latitude).HasColumnName("lat");
                entity.Property(e => e.Longitude).HasColumnName("long");
                entity.Property(e => e.Timestamp).HasColumnName("timestamp");
                entity.Property(e => e.CreatedAt).HasColumnName("created_at").ValueGeneratedOnAdd();
            });

            modelBuilder.Entity<OutboxMessage>(entity =>
            {
                entity.ToTable("outbox_messages");
                entity.HasKey(e => e.Id);
                entity.Property(e => e.Id).HasColumnName("id");
                entity.Property(e => e.Payload).HasColumnType("jsonb");
                entity.Property(e => e.Status).HasDefaultValue("PENDING");
            });
        }
    }
}
