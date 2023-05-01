#!/usr/bin/env bash

source ./scripts/env.sh

set -o allexport; source .env; set +o allexport

function error() {
	echo ${1}
	exit 1
}

# building fresh binaries
echo "Compiling the project..."
if [ ! -d $OUTPUT_DIRECTORY ]; then
	mkdir -p $OUTPUT_DIRECTORY || error "Failed to create build directory..."
fi

sudo /usr/local/go/bin/go build -o $BINARY ./cmd/backend/*.go || error "Failed to compile new code..."

# stop service
echo "Stopping server..."
sudo service $PROJECT_NAME stop # || error "Failed to stop server..."

# removing old build
echo "Removing old build..."
sudo rm -rf $TARGET_DIRECTORY || error "Failed to remove last build..."

# creating new build
echo "Creating new build..."
sudo mkdir -p $TARGET_DIRECTORY || error "Failed to create target directory..."

echo "Configuring environment..."
sudo cp -R ./server/static $TARGET_DIRECTORY/ || error "Failed to copy static files..."
sudo cp .env $TARGET_DIRECTORY/ || error "Failed to update environment variables..."
sudo touch $TARGET_DIRECTORY/hooks.yml

echo "Copying binaries..."
sudo cp $BINARY $TARGET || error "Failed to copy new binaries..."

# copying the service files
echo "Configuring service..."
sudo cp WEBHOOK_SERVICE_FILE /etc/systemd/system/ || error "Failed to copy webhook service file..."
sudo cp $SERVICE_FILE /etc/systemd/system/ || error "Failed to copy service file..."

sudo systemctl daemon-reload || error "Failed to restart daemon..."

sudo systemctl enable $WEBHOOK_SERVICE_FILE || error "Failed to enable service..."
sudo systemctl enable $SERVICE_FILE || error "Failed to enable service..."


# starting the services
echo "Starting server..."
sudo service $WEBHOOK_SERVICE_FILE start || error "Failed to start server..."
sudo service $SERVICE_FILE start || error "Failed to start server..."

echo "Setup complete..."
