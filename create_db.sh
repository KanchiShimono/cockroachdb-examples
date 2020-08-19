#!/bin/bash

docker-compose exec roach1 cockroach sql --insecure --host=localhost:26257 -d review --execute="$(cat ./migrations/create_db.sql)"
