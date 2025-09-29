#!/bin/bash

# kill-port.sh - Kill processes running on a specific port
# Usage: ./kill-port.sh <port_number>
# Example: ./kill-port.sh 8091

if [ $# -eq 0 ]; then
    echo "Usage: $0 <port_number>"
    echo "Example: $0 8091"
    exit 1
fi

PORT=$1

# Check if port is a valid number
if ! [[ "$PORT" =~ ^[0-9]+$ ]]; then
    echo "Error: Port must be a number"
    exit 1
fi

echo "Looking for processes on port $PORT..."

# Get PIDs of processes using the port
PIDS=$(lsof -ti :$PORT)

if [ -z "$PIDS" ]; then
    echo "No processes found running on port $PORT"
    exit 0
fi

echo "Found processes on port $PORT:"
lsof -i :$PORT

echo ""
echo "Killing processes: $PIDS"

# Kill the processes
for PID in $PIDS; do
    echo "Killing process $PID..."
    kill -9 $PID
    if [ $? -eq 0 ]; then
        echo "✓ Successfully killed process $PID"
    else
        echo "✗ Failed to kill process $PID"
    fi
done

echo "Done!"
