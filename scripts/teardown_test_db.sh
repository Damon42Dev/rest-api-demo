#!/bin/bash

# Drop the test database
mongosh <<EOF
use test_db
db.dropDatabase()
EOF

is_mongodb_running() {
    pgrep mongod > /dev/null 2>&1
}

# Stop MongoDB
echo "Stopping MongoDB..."
mongosh admin --eval "db.shutdownServer()"

# Wait for MongoDB to stop
sleep 5

# Check if MongoDB is still running
if is_mongodb_running; then
    echo "MongoDB shutdown failed."
    exit 1
else
    echo "MongoDB shutdown successfully."
fi
