source devtools/lib.sh || { echo "Are you at repo root?"; exit 1; }

usage() {
  cat <<EOUSAGE
Usage: $0 [up|down|force|version] {#}"
EOUSAGE
}

database_host="localhost"
if [[ $DB_HOST != "" ]]; then
  database_host=$DB_HOST
fi
database_port="5432"
if [[ $DB_PORT != "" ]]; then
  database_port=$DB_PORT
fi
database_user="testuser"
if [[ $DB_USER != "" ]]; then
  database_user=$DB_USER
fi
database_name='testdb'
if [[ $DB_NAME != "" ]]; then
  database_name=$DB_NAME
fi
database_password="testpassword"
if [[ $DB_PASS != "" ]]; then
  database_password=$DB_PASS
fi
ssl_mode='disable'
if [[ $DB_SSL != "" ]]; then
  ssl_mode=$DB_SSL
fi

# Redirect stderr to stdout because migrate outputs to stderr, and we want
# to be able to use ordinary output redirection.
case "$1" in
  up|down|force|version)
    migrate \
      -source file:migrations \
      -database "postgresql://$database_user:$database_password@$database_host:$database_port/$database_name?sslmode=$ssl_mode" \
      "$@" 2>&1
    ;;
  *)
    usage
    exit 1
    ;;
esac