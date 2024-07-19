export DB_HOST="127.0.0.1"
export DB_PORT=5432
export DB_NAME="sobo"
export DB_USER="postgres"
export DB_PASSWORD="postgres"
export TEGOLA_POSTGIS_SSL="disable"
export TEGOLA_REDIS_SSL="false"
export REDIS_HOST="127.0.0.1:6379"
export REDIS_PASSWORD=""
export TEGOLA_SQL_DEBUG="LAYER_SQL"

../cmd/tegola/tegola cache purge --config ./config/config-dev.toml
../cmd/tegola/tegola serve --config ./config/config-dev.toml --log-level DEBUG