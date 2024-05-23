#!/bin/bash
# wait.sh

set -e

host="$1"
shift
cmd="$@"
echo "$host"
until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "postgres" -lqt | cut -d \| -f 1 | grep -qw "OzonContest"; do
  echo "Waiting."
  sleep 1
done

echo "Ready"
exec $cmd