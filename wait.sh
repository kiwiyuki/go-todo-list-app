#!/bin/sh

set -e
export $(xargs < /app/.env)

cmd="$@"

echo "Waiting for mysql"
until mysql -hdatabase -u${MYSQL_USER} -p${MYSQL_PASSWORD} &> /dev/null
do
        >&2 echo -n "."
        sleep 1
done

>&2 echo "MySQL is up - executing command"

>&2 echo $cmd
exec $cmd
