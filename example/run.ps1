# Set-ExecutionPolicy -Scope Process -ExecutionPolicy RemoteSigned

$DB_HOST = "127.0.0.1"
$DB_PORT = 5432
$DB_NAME = "sobo"
$DB_USER = "postgres"
$DB_PASSWORD = "postgres"
$TEGOLA_POSTGIS_SSL = "disable"
$TEGOLA_REDIS_SSL = "false"
$REDIS_HOST = "127.0.0.1:6379"
$REDIS_PASSWORD = ""
$TEGOLA_SQL_DEBUG = "LAYER_SQL"

# User/Machine/Process
$scope = "Process"

[Environment]::SetEnvironmentVariable("DB_HOST", $DB_HOST, $scope)
[Environment]::SetEnvironmentVariable("DB_PORT", $DB_PORT, $scope)
[Environment]::SetEnvironmentVariable("DB_NAME", $DB_NAME, $scope)
[Environment]::SetEnvironmentVariable("DB_USER", $DB_USER, $scope)
[Environment]::SetEnvironmentVariable("DB_PASSWORD", $DB_PASSWORD, $scope)
[Environment]::SetEnvironmentVariable("TEGOLA_POSTGIS_SSL", $TEGOLA_POSTGIS_SSL, $scope)
[Environment]::SetEnvironmentVariable("TEGOLA_REDIS_SSL", $TEGOLA_REDIS_SSL, $scope)
[Environment]::SetEnvironmentVariable("REDIS_HOST", $REDIS_HOST, $scope)
[Environment]::SetEnvironmentVariable("REDIS_PASSWORD", $REDIS_PASSWORD, $scope)
[Environment]::SetEnvironmentVariable("TEGOLA_SQL_DEBUG", $TEGOLA_SQL_DEBUG, $scope)

../cmd/tegola/tegola.exe cache purge --config ./config/config-dev.toml
../cmd/tegola/tegola.exe serve --log-level DEBUG --config ./config/config-dev.toml 