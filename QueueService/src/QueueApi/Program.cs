var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddOpenApi();

var app = builder.Build();

app.MapOpenApi();

app.UseHttpsRedirection();


app.MapGet("/healthz", () =>
{
    return new { Message = "healthy"};
})
.WithName("Ping")
.WithDescription("Simple healthcheck")
.WithOpenApi();

app.Run();
