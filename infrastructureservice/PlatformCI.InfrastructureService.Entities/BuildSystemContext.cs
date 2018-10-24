using Microsoft.EntityFrameworkCore;

namespace PlatformCI.InfrastructureService.Entities
{
    public class BuildSystemContext : DbContext
    {
        public BuildSystemContext(DbContextOptions<BuildSystemContext> dbContextOptions) : base(dbContextOptions)
        {
            // IGNORE THIS 
            Database.SetCommandTimeout(999999);
        }

        public virtual DbSet<BusinessLines> BusinessLines { get; set; }
        public virtual DbSet<DeliverableGroupAssociations> DeliverableGroupAssociations { get; set; }
        public virtual DbSet<DeliverableGroups> DeliverableGroups { get; set; }
        public virtual DbSet<Deliverables> Deliverables { get; set; }
        public virtual DbSet<DeploymentLocationGroupAssociations> DeploymentLocationGroupAssociations { get; set; }
        public virtual DbSet<DeploymentLocationGroups> DeploymentLocationGroups { get; set; }
        public virtual DbSet<DeploymentLocations> DeploymentLocations { get; set; }
        public virtual DbSet<Environments> Environments { get; set; }
        public virtual DbSet<IisApplicationGroupAssociations> IisApplicationGroupAssociations { get; set; }
        public virtual DbSet<IisApplicationGroups> IisApplicationGroups { get; set; }
        public virtual DbSet<IisApplicationPools> IisApplicationPools { get; set; }
        public virtual DbSet<IisApplications> IisApplications { get; set; }
        public virtual DbSet<IisSiteApplicationGroupAssociations> IisSiteApplicationGroupAssociations { get; set; }
        public virtual DbSet<IisSiteBindings> IisSiteBindings { get; set; }
        public virtual DbSet<IisSiteGroupAssociations> IisSiteGroupAssociations { get; set; }
        public virtual DbSet<IisSiteGroups> IisSiteGroups { get; set; }
        public virtual DbSet<IisSites> IisSites { get; set; }
        public virtual DbSet<Jobs> Jobs { get; set; }
        public virtual DbSet<JobStates> JobStates { get; set; }
        public virtual DbSet<KeyValueCache> KeyValueCache { get; set; }
        public virtual DbSet<Logs> Logs { get; set; }
        public virtual DbSet<ResourceAmalgamations> ResourceAmalgamations { get; set; }
        public virtual DbSet<ResourceAssociations> ResourceAssociations { get; set; }
        public virtual DbSet<ScheduledTaskGroupAssociations> ScheduledTaskGroupAssociations { get; set; }
        public virtual DbSet<ScheduledTaskGroups> ScheduledTaskGroups { get; set; }
        public virtual DbSet<ScheduledTasks> ScheduledTasks { get; set; }
        public virtual DbSet<ServerGroupAssociations> ServerGroupAssociations { get; set; }
        public virtual DbSet<ServerGroupEnvironmentAssociations> ServerGroupEnvironmentAssociations { get; set; }
        public virtual DbSet<ServerGroups> ServerGroups { get; set; }
        public virtual DbSet<Servers> Servers { get; set; }
        public virtual DbSet<ServerTypes> ServerTypes { get; set; }
        public virtual DbSet<WindowsServiceGroupAssociations> WindowsServiceGroupAssociations { get; set; }
        public virtual DbSet<WindowsServiceGroups> WindowsServiceGroups { get; set; }
        public virtual DbSet<WindowsServices> WindowsServices { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<BusinessLines>(entity =>
            {
                entity.HasKey(e => e.BusinessLineId);

                entity.HasIndex(e => e.Name)
                    .HasName("UC_BusinessLines")
                    .IsUnique();

                entity.Property(e => e.BusinessLineId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.Description)
                    .IsRequired()
                    .HasMaxLength(2550);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.Name)
                    .IsRequired()
                    .HasMaxLength(255);
            });

            modelBuilder.Entity<DeliverableGroupAssociations>(entity =>
            {
                entity.HasKey(e => new {e.DeliverableId, e.DeliverableGroupId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.DeliverableGroup)
                    .WithMany(p => p.DeliverableGroupAssociations)
                    .HasForeignKey(d => d.DeliverableGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_DeliverableGroupAssociations_DeliverableGroups");

                entity.HasOne(d => d.Deliverable)
                    .WithMany(p => p.DeliverableGroupAssociations)
                    .HasForeignKey(d => d.DeliverableId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_DeliverableGroupAssociations_Deliverables");
            });

            modelBuilder.Entity<DeliverableGroups>(entity =>
            {
                entity.HasKey(e => e.DeliverableGroupId);

                entity.HasIndex(e => e.GroupName)
                    .HasName("UC_DeliverableGroups")
                    .IsUnique();

                entity.Property(e => e.DeliverableGroupId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.GroupName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");
            });

            modelBuilder.Entity<Deliverables>(entity =>
            {
                entity.HasKey(e => e.DeliverableId);

                entity.HasIndex(e => new {e.RepositoryName, e.SolutionName, e.ProjectName})
                    .HasName("UC_Deliverables")
                    .IsUnique();

                entity.Property(e => e.DeliverableId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.ProjectName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.RepositoryName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.SolutionName)
                    .IsRequired()
                    .HasMaxLength(255);
            });

            modelBuilder.Entity<DeploymentLocationGroupAssociations>(entity =>
            {
                entity.HasKey(e => new {e.DeploymentLocationId, e.DeploymentLocationGroupId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.DeploymentLocationGroup)
                    .WithMany(p => p.DeploymentLocationGroupAssociations)
                    .HasForeignKey(d => d.DeploymentLocationGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_DeploymentLocationGroupAssociations_DeploymentLocationGroups");

                entity.HasOne(d => d.DeploymentLocation)
                    .WithMany(p => p.DeploymentLocationGroupAssociations)
                    .HasForeignKey(d => d.DeploymentLocationId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_DeploymentLocationGroupAssociations_DeploymentLocations");
            });

            modelBuilder.Entity<DeploymentLocationGroups>(entity =>
            {
                entity.HasKey(e => e.DeploymentLocationGroupId);

                entity.HasIndex(e => e.GroupName)
                    .HasName("UC_DeploymentLocationGroups")
                    .IsUnique();

                entity.Property(e => e.DeploymentLocationGroupId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.GroupName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.Environment)
                    .WithMany(p => p.DeploymentLocationGroups)
                    .HasForeignKey(d => d.EnvironmentId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_DeploymentLocationGroups_Environments");
            });

            modelBuilder.Entity<DeploymentLocations>(entity =>
            {
                entity.HasKey(e => e.DeploymentLocationId);

                entity.HasIndex(e => e.FriendlyName)
                    .HasName("UC_DeploymentLocations")
                    .IsUnique();

                entity.Property(e => e.DeploymentLocationId).ValueGeneratedNever();

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.FriendlyName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.PhysicalPath)
                    .IsRequired()
                    .HasMaxLength(2550);
            });

            modelBuilder.Entity<Environments>(entity =>
            {
                entity.HasKey(e => e.EnvironmentId);

                entity.HasIndex(e => e.EnvironmentName)
                    .HasName("UC_Environments")
                    .IsUnique();

                entity.Property(e => e.EnvironmentId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.EnvironmentName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");
            });

            modelBuilder.Entity<IisApplicationGroupAssociations>(entity =>
            {
                entity.HasKey(e => new {e.IisApplicationGroupId, e.IisApplicationId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.IisApplicationGroup)
                    .WithMany(p => p.IisApplicationGroupAssociations)
                    .HasForeignKey(d => d.IisApplicationGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisApplicationGroupAssignments_IisApplicationGroups");

                entity.HasOne(d => d.IisApplication)
                    .WithMany(p => p.IisApplicationGroupAssociations)
                    .HasForeignKey(d => d.IisApplicationId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisApplicationGroupAssignments_IisApplications");
            });

            modelBuilder.Entity<IisApplicationGroups>(entity =>
            {
                entity.HasKey(e => e.IisApplicationGroupId);

                entity.HasIndex(e => e.GroupName)
                    .HasName("UC_IisApplicationGroups")
                    .IsUnique();

                entity.Property(e => e.IisApplicationGroupId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.GroupName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");
            });

            modelBuilder.Entity<IisApplicationPools>(entity =>
            {
                entity.HasKey(e => e.IisApplicationPoolId);

                entity.HasIndex(e => e.PoolName)
                    .HasName("UC_IisApplicationPools")
                    .IsUnique();

                entity.Property(e => e.IisApplicationPoolId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.PoolFrameworkVersion)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.PoolName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.PoolType)
                    .IsRequired()
                    .HasMaxLength(255);
            });

            modelBuilder.Entity<IisApplications>(entity =>
            {
                entity.HasKey(e => e.IisApplicationId);

                entity.Property(e => e.IisApplicationId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.ApplicationInternalAliasName).HasMaxLength(255);

                entity.Property(e => e.ApplicationName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.PhysicalPath)
                    .IsRequired()
                    .HasMaxLength(2550);

                entity.HasOne(d => d.IisApplicationPool)
                    .WithMany(p => p.IisApplications)
                    .HasForeignKey(d => d.IisApplicationPoolId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisApplications_IisApplicationPools");
            });

            modelBuilder.Entity<IisSiteApplicationGroupAssociations>(entity =>
            {
                entity.HasKey(e => new {e.IisApplicationGroupId, e.IisSiteId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.IisApplicationGroup)
                    .WithMany(p => p.IisSiteApplicationGroupAssociations)
                    .HasForeignKey(d => d.IisApplicationGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisSiteApplicationGroupAssignments_IisApplicationGroups");

                entity.HasOne(d => d.IisSite)
                    .WithMany(p => p.IisSiteApplicationGroupAssociations)
                    .HasForeignKey(d => d.IisSiteId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisSiteApplicationGroupAssignments_IisSites");
            });

            modelBuilder.Entity<IisSiteBindings>(entity =>
            {
                entity.HasKey(e => e.IisSiteBindingId);

                entity.Property(e => e.IisSiteBindingId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.BindingName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.CertificateThumbprint)
                    .IsRequired()
                    .HasMaxLength(2550);

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.IisSite)
                    .WithMany(p => p.IisSiteBindings)
                    .HasForeignKey(d => d.IisSiteId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisSiteBindings_IisSites");
            });

            modelBuilder.Entity<IisSiteGroupAssociations>(entity =>
            {
                entity.HasKey(e => new {e.IisSiteGroupId, e.IisSiteId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.IisSiteGroup)
                    .WithMany(p => p.IisSiteGroupAssociations)
                    .HasForeignKey(d => d.IisSiteGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisSiteGroupAssignments_IisSiteGroups");

                entity.HasOne(d => d.IisSite)
                    .WithMany(p => p.IisSiteGroupAssociations)
                    .HasForeignKey(d => d.IisSiteId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisSiteGroupAssignments_IisSites");
            });

            modelBuilder.Entity<IisSiteGroups>(entity =>
            {
                entity.HasKey(e => e.IisSiteGroupId);

                entity.HasIndex(e => e.GroupName)
                    .HasName("UC_IisSiteGroups")
                    .IsUnique();

                entity.Property(e => e.IisSiteGroupId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.GroupName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.Environment)
                    .WithMany(p => p.IisSiteGroups)
                    .HasForeignKey(d => d.EnvironmentId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisSiteGroups_Environments");
            });

            modelBuilder.Entity<IisSites>(entity =>
            {
                entity.HasKey(e => e.IisSiteId);

                entity.HasIndex(e => e.SiteName)
                    .HasName("UC_IisSites")
                    .IsUnique();

                entity.Property(e => e.IisSiteId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.PhysicalPath)
                    .IsRequired()
                    .HasMaxLength(2550);

                entity.Property(e => e.SiteName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.HasOne(d => d.IisSiteApplicationPool)
                    .WithMany(p => p.IisSites)
                    .HasForeignKey(d => d.IisSiteApplicationPoolId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_IisSites_IisApplicationPools");
            });

            modelBuilder.Entity<Jobs>(entity =>
            {
                entity.HasKey(e => e.JobId);

                entity.HasIndex(e => e.JobName)
                    .HasName("UC_Jobs")
                    .IsUnique();

                entity.Property(e => e.JobId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.JobName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.MachineName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.HasOne(d => d.JobState)
                    .WithMany(p => p.Jobs)
                    .HasForeignKey(d => d.JobStateId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_Jobs_JobStates");
            });

            modelBuilder.Entity<JobStates>(entity =>
            {
                entity.HasKey(e => e.JobStateId);

                entity.Property(e => e.JobStateId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.StateName)
                    .IsRequired()
                    .HasMaxLength(255);
            });

            modelBuilder.Entity<KeyValueCache>(entity =>
            {
                entity.HasIndex(e => new {e.Key, e.MachineName})
                    .HasName("UC_KeyValueCache")
                    .IsUnique();

                entity.Property(e => e.KeyValueCacheId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.Key)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.MachineName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.Value).IsRequired();

                entity.Property(e => e.ValueType)
                    .IsRequired()
                    .HasMaxLength(2550);
            });

            modelBuilder.Entity<Logs>(entity =>
            {
                entity.HasKey(e => e.LogId);

                entity.Property(e => e.LogId).ValueGeneratedNever();

                entity.Property(e => e.ApplicationName)
                    .IsRequired()
                    .HasMaxLength(4000);

                entity.Property(e => e.Date)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LogLevel)
                    .IsRequired()
                    .HasMaxLength(4000);

                entity.Property(e => e.MachineName)
                    .IsRequired()
                    .HasMaxLength(4000);
            });

            modelBuilder.Entity<ResourceAmalgamations>(entity =>
            {
                entity.HasKey(e => e.ResourceAmalgamationId);

                entity.HasIndex(e => e.Name)
                    .HasName("UC_ResourceAmalgamations")
                    .IsUnique();

                entity.Property(e => e.ResourceAmalgamationId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.Name)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.HasOne(d => d.BusinessLine)
                    .WithMany(p => p.ResourceAmalgamations)
                    .HasForeignKey(d => d.BusinessLineId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ResourceAmalgamations_BusinessLines");
            });

            modelBuilder.Entity<ResourceAssociations>(entity =>
            {
                entity.HasKey(e => e.ResourceAssociationId);

                entity.Property(e => e.ResourceAssociationId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.Enabled)
                    .IsRequired()
                    .HasDefaultValueSql("((1))");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.Deliverable)
                    .WithMany(p => p.ResourceAssociations)
                    .HasForeignKey(d => d.DeliverableId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ResourceAssociations_Deliverables");

                entity.HasOne(d => d.DeploymentLocation)
                    .WithMany(p => p.ResourceAssociations)
                    .HasForeignKey(d => d.DeploymentLocationId)
                    .HasConstraintName("FK_ResourceAssociations_DeploymentLocations");

                entity.HasOne(d => d.IisApplication)
                    .WithMany(p => p.ResourceAssociations)
                    .HasForeignKey(d => d.IisApplicationId)
                    .HasConstraintName("FK_ResourceAssociations_IisApplications");

                entity.HasOne(d => d.IisSite)
                    .WithMany(p => p.ResourceAssociations)
                    .HasForeignKey(d => d.IisSiteId)
                    .HasConstraintName("FK_ResourceAssociations_IisSites");

                entity.HasOne(d => d.ResourceAmalgamation)
                    .WithMany(p => p.ResourceAssociations)
                    .HasForeignKey(d => d.ResourceAmalgamationId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ResourceAssociations_ResourceAmalgamations");

                entity.HasOne(d => d.ScheduledTask)
                    .WithMany(p => p.ResourceAssociations)
                    .HasForeignKey(d => d.ScheduledTaskId)
                    .HasConstraintName("FK_ResourceAssociations_ScheduledTasks");

                entity.HasOne(d => d.WindowsService)
                    .WithMany(p => p.ResourceAssociations)
                    .HasForeignKey(d => d.WindowsServiceId)
                    .HasConstraintName("FK_ResourceAssociations_WindowsServices");
            });

            modelBuilder.Entity<ScheduledTaskGroupAssociations>(entity =>
            {
                entity.HasKey(e => new {e.ScheduledTaskId, e.ScheduledTaskGroupId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.ScheduledTaskGroup)
                    .WithMany(p => p.ScheduledTaskGroupAssociations)
                    .HasForeignKey(d => d.ScheduledTaskGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ScheduledTaskGroupAssignments_ScheduledTaskGroupId");

                entity.HasOne(d => d.ScheduledTask)
                    .WithMany(p => p.ScheduledTaskGroupAssociations)
                    .HasForeignKey(d => d.ScheduledTaskId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ScheduledTaskGroupAssignments_ScheduledTaskId");
            });

            modelBuilder.Entity<ScheduledTaskGroups>(entity =>
            {
                entity.HasKey(e => e.ScheduledTaskGroupId);

                entity.HasIndex(e => e.TaskGroupName)
                    .HasName("UC_TaskGroupName")
                    .IsUnique();

                entity.Property(e => e.ScheduledTaskGroupId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.TaskGroupName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.HasOne(d => d.Environment)
                    .WithMany(p => p.ScheduledTaskGroups)
                    .HasForeignKey(d => d.EnvironmentId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ScheduledTaskGroups_Environments");
            });

            modelBuilder.Entity<ScheduledTasks>(entity =>
            {
                entity.HasKey(e => e.ScheduledTaskId);

                entity.HasIndex(e => e.TaskName)
                    .HasName("UC_ScheduledTasks")
                    .IsUnique();

                entity.Property(e => e.ScheduledTaskId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.BinaryExecutableArguments).HasMaxLength(255);

                entity.Property(e => e.BinaryExecutableName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.BinaryPath)
                    .IsRequired()
                    .HasMaxLength(2550);

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.ScheduleType).HasMaxLength(255);

                entity.Property(e => e.TaskName)
                    .IsRequired()
                    .HasMaxLength(255);
            });

            modelBuilder.Entity<ServerGroupAssociations>(entity =>
            {
                entity.HasKey(e => new {e.ServerId, e.ServerGroupId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.ServerGroup)
                    .WithMany(p => p.ServerGroupAssociations)
                    .HasForeignKey(d => d.ServerGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ServerGroupAssociations_ServerGroups");

                entity.HasOne(d => d.Server)
                    .WithMany(p => p.ServerGroupAssociations)
                    .HasForeignKey(d => d.ServerId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ServerGroupAssociations_Servers");
            });

            modelBuilder.Entity<ServerGroupEnvironmentAssociations>(entity =>
            {
                entity.HasKey(e => new {e.EnvironmentId, e.ServerGroupId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.Environment)
                    .WithMany(p => p.ServerGroupEnvironmentAssociations)
                    .HasForeignKey(d => d.EnvironmentId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ServerGroupEnvironmentAssociations_Environments");

                entity.HasOne(d => d.ServerGroup)
                    .WithMany(p => p.ServerGroupEnvironmentAssociations)
                    .HasForeignKey(d => d.ServerGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_ServerGroupEnvironmentAssociations_ServerGroups");
            });

            modelBuilder.Entity<ServerGroups>(entity =>
            {
                entity.HasKey(e => e.ServerGroupId);

                entity.HasIndex(e => e.GroupName)
                    .HasName("UC_ServerGroups")
                    .IsUnique();

                entity.Property(e => e.ServerGroupId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.GroupName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");
            });

            modelBuilder.Entity<Servers>(entity =>
            {
                entity.HasKey(e => e.ServerId);

                entity.HasIndex(e => e.ServerName)
                    .HasName("UC_Servers")
                    .IsUnique();

                entity.Property(e => e.ServerId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.ServerName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.HasOne(d => d.ServerType)
                    .WithMany(p => p.Servers)
                    .HasForeignKey(d => d.ServerTypeId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_Servers_ServerTypes");
            });

            modelBuilder.Entity<ServerTypes>(entity =>
            {
                entity.HasKey(e => e.ServerTypeId);

                entity.HasIndex(e => e.ServerTypeName)
                    .HasName("UC_ServerTypeName")
                    .IsUnique();

                entity.Property(e => e.ServerTypeId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.ServerTypeName)
                    .IsRequired()
                    .HasMaxLength(255);
            });

            modelBuilder.Entity<WindowsServiceGroupAssociations>(entity =>
            {
                entity.HasKey(e => new {e.WindowsServiceId, e.WindowsServiceGroupId});

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.WindowsServiceGroup)
                    .WithMany(p => p.WindowsServiceGroupAssociations)
                    .HasForeignKey(d => d.WindowsServiceGroupId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_WindowsServiceGroupAssignments_WindowsServiceGroupId");

                entity.HasOne(d => d.WindowsService)
                    .WithMany(p => p.WindowsServiceGroupAssociations)
                    .HasForeignKey(d => d.WindowsServiceId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_WindowsServiceGroupAssignments_WindowsServiceId");
            });

            modelBuilder.Entity<WindowsServiceGroups>(entity =>
            {
                entity.HasKey(e => e.WindowsServiceGroupId);

                entity.HasIndex(e => e.GroupName)
                    .HasName("UC_WindowsServiceGroups")
                    .IsUnique();

                entity.Property(e => e.WindowsServiceGroupId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.GroupName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.HasOne(d => d.Environment)
                    .WithMany(p => p.WindowsServiceGroups)
                    .HasForeignKey(d => d.EnvironmentId)
                    .OnDelete(DeleteBehavior.ClientSetNull)
                    .HasConstraintName("FK_WindowsServiceGroups_Environments");
            });

            modelBuilder.Entity<WindowsServices>(entity =>
            {
                entity.HasKey(e => e.WindowsServiceId);

                entity.HasIndex(e => e.ServiceName)
                    .HasName("UC_WindowsServices")
                    .IsUnique();

                entity.Property(e => e.WindowsServiceId).HasDefaultValueSql("(newid())");

                entity.Property(e => e.BinaryExecutableArguments).HasMaxLength(255);

                entity.Property(e => e.BinaryExecutableName)
                    .IsRequired()
                    .HasMaxLength(255);

                entity.Property(e => e.BinaryPath)
                    .IsRequired()
                    .HasMaxLength(2550);

                entity.Property(e => e.CreatedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.CreatedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.LastModifiedBy)
                    .IsRequired()
                    .HasMaxLength(255)
                    .HasDefaultValueSql("(suser_sname())");

                entity.Property(e => e.LastModifiedDateTime)
                    .HasColumnType("datetime")
                    .HasDefaultValueSql("(getdate())");

                entity.Property(e => e.ServiceName)
                    .IsRequired()
                    .HasMaxLength(255);
            });
        }
    }
}