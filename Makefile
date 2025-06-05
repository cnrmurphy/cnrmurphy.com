APP_NAME = cnrmurphy.com
SRC_DIR = .
BIN_PATH = /usr/local/bin/$(APP_NAME)

.PHONY: build install relabel deploy clean restart

build:
	@echo "building binary..."
	go build -o $(APP_NAME) $(SRC_DIR)

install: build
	@echo "installing binary to $(BIN_PATH)..."
	sudo mv $(APP_NAME) $(BIN_PATH)
	sudo chmod +x $(BIN_PATH)

relabel:
	@echo "restoring SELinux context..."
	sudo restorecon -v $(BIN_PATH)

restart:
	@echo "restarting systemd service..."
	sudo systemctl restart $(APP_NAME)

status:
	sudo systemctl status $(APP_NAME)

deploy: install relabel restart

clean:
	rm -f $(APP_NAME)
