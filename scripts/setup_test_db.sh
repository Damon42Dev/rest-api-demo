#!/bin/bash

# Exit on any error
set -e

# Function to check if MongoDB is running
is_mongodb_running() {
    pgrep mongod >/dev/null 2>&1
    return $?
}

# Function to wait for MongoDB to be ready
wait_for_mongodb() {
    local retries=30
    local wait_time=2
    
    echo "Waiting for MongoDB to be ready..."
    while [ $retries -gt 0 ]; do
        if mongosh --eval "db.runCommand({ ping: 1 })" >/dev/null 2>&1; then
            echo "MongoDB is ready!"
            return 0
        fi
        retries=$((retries - 1))
        echo "Waiting for MongoDB to be ready... ($retries retries left)"
        sleep $wait_time
    done
    
    echo "Timed out waiting for MongoDB to be ready"
    return 1
}

# Ensure MongoDB directories exist with proper permissions
ensure_mongodb_dirs() {
    local dirs=("/opt/homebrew/var/mongodb" "/opt/homebrew/var/log/mongodb")
    
    for dir in "${dirs[@]}"; do
        if [ ! -d "$dir" ]; then
            mkdir -p "$dir" || {
                echo "Failed to create directory: $dir"
                exit 1
            }
        fi
        
        # In CI environment, don't worry about ownership
        if [ -z "$CI" ]; then
            chown -R $(whoami) "$dir" || {
                echo "Failed to set ownership for directory: $dir"
                exit 1
            }
        fi
    done
}

# Stop any existing MongoDB process
cleanup_mongodb() {
    if is_mongodb_running; then
        echo "Stopping existing MongoDB process..."
        pkill -f mongod || true
        sleep 2
    fi
}

# Only cleanup on script failure or interruption
cleanup_on_failure() {
    local exit_code=$?
    if [ $exit_code -ne 0 ]; then
        echo "Script failed or interrupted, cleaning up MongoDB..."
        cleanup_mongodb
    fi
    exit $exit_code
}

# Set trap for script failure or interruption
trap cleanup_on_failure ERR INT TERM

# Ensure directories exist
ensure_mongodb_dirs

# Cleanup any existing process before starting
cleanup_mongodb

# Start MongoDB
echo "Starting MongoDB..."
mongod --dbpath /opt/homebrew/var/mongodb --logpath /opt/homebrew/var/log/mongodb/mongo.log --fork

# Wait for MongoDB to be ready
if ! wait_for_mongodb; then
    echo "Failed to start MongoDB"
    exit 1
fi

echo "MongoDB started successfully"

# Seed the database with testing data
echo "Seeding the database..."
mongosh --quiet <<EOF
use test_db
db.dropDatabase()
db.movies.insertMany([
  {
    _id: ObjectId('60c72b2f9b1e8a5d6c8b4567'),
    title: 'Test Movie 1',
    plot: 'This is a test movie 1',
    directors: ['Test Director 1'],
    released: new Date(),
    year: 2024,
    rated: 'PG-13',
    runtime: 120,
    type: 'movie'
  },
  {
    _id: ObjectId('60c72b2f9b1e8a5d6c8b4568'),
    title: 'Test Movie 2',
    plot: 'This is a test movie 2',
    directors: ['Test Director 2'],
    released: new Date(),
    year: 2024,
    rated: 'PG-13',
    runtime: 120,
    type: 'movie'
  }
])
db.comments.insertMany([
  {
    _id: ObjectId('60c72b2f9b1e8a5d6c8b4567'),
    date: new Date(),
    email: 'test1@example.com',
    movie_id: ObjectId('60c72b2f9b1e8a5d6c8b4567'),
    name: 'Test User 1',
    text: 'This is a test comment 1'
  },
  {
    _id: ObjectId('60c72b2f9b1e8a5d6c8b4568'),
    date: new Date(),
    email: 'test2@example.com',
    movie_id: ObjectId('60c72b2f9b1e8a5d6c8b4567'),
    name: 'Test User 2',
    text: 'This is a test comment 2'
  }
])

// Verify data was inserted
const moviesCount = db.movies.count();
const commentsCount = db.comments.count();
if (moviesCount !== 2 || commentsCount !== 2) {
    print("Data verification failed!");
    quit(1);
}
EOF

if [ $? -ne 0 ]; then
    echo "Failed to seed the database"
    exit 1
fi

echo "Database seeded successfully"
echo "MongoDB is running and ready to use!"