package healthcheck


type HealthCheckService interface {
    Health() map[string]string
}
