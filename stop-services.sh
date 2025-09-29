#!/bin/bash

# stop-services.sh - Stop all POC services
# Usage: ./stop-services.sh [port1] [port2] [port3]
# Example: ./stop-services.sh 8091 8092 8093

# Default ports if none provided
PORTS=${@:-8091 8092 8093}

echo "Stopping POC services on ports: $PORTS"
echo "=================================="

for PORT in $PORTS; do
    echo ""
    echo "Checking port $PORT..."
    
    # Get PIDs of processes using the port
    PIDS=$(lsof -ti :$PORT 2>/dev/null)
    
    if [ -z "$PIDS" ]; then
        echo "✓ No processes found on port $PORT"
        continue
    fi
    
    echo "Found processes on port $PORT:"
    lsof -i :$PORT
    
    echo "Killing processes: $PIDS"
    
    # Kill the processes
    for PID in $PIDS; do
        echo "  Killing process $PID..."
        kill -9 $PID 2>/dev/null
        if [ $? -eq 0 ]; then
            echo "  ✓ Successfully killed process $PID"
        else
            echo "  ✗ Failed to kill process $PID"
        fi
    done
done

echo ""
echo "All services stopped!"
