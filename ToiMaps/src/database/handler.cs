namespace ToiMaps.database;

using Npgsql;
public class handler
{


    var connString = "Host=;Username=postgres;Password=toimap;Database=toimaps";

    var dataSourceBuilder = new NpgsqlDataSourceBuilder(connString);
    var dataSource = dataSourceBuilder.Build();

    var conn = await dataSource.OpenConnectionAsync();

}