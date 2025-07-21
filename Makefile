APP_NAME = cnrmurphy.com
SRC_DIR = /.
BIN_PATH = /usr/local/bin/$(APP_NAME)
WWW = /www/$(APP_NAME)
NGINX_CONFIG_PATH = /etc/nginx/sites-available/$(APP_NAME)
NGINX_ENABLED_PATH = /etc/nginx/sites-enabled/$(APP_NAME)

.PHONY: build install relabel deploy clean restart status nginx-install nginx-reload

generate:
	@echo "building html files"
	./build_html.sh

build:
	@echo "building binary..."
	go build -o $(APP_NAME) $(SRC_DIR)

dev: generate
	@echo "deploying dev server"
	go run .

install: build
	@echo "installing binary to $(BIN_PATH)..."
	sudo mv $(APP_NAME) $(BIN_PATH)
	sudo chmod +x $(BIN_PATH)

build-dist:
	@echo "building html files"
	./build_html.sh
	@echo "moving public files to www"
	sudo cp -r ./public.* $(WWW)

nginx-install:
	@echo "installing nginx configuration..."
	sudo cp nginx.conf $(NGINX_CONFIG_PATH)
	sudo ln -sf $(NGINX_CONFIG_PATH) $(NGINX_ENABLED_PATH)
	sudo nginx -t

nginx-reload:
	@echo "reloading nginx..."
	sudo systemctl reload nginx

relabel:
	@echo "restoring SELinux context..."
	sudo restorecon -v $(BIN_PATH)

restart:
	@echo "restarting systemd service..."
	sudo systemctl restart $(APP_NAME)

status:
	sudo systemctl status $(APP_NAME)

deploy: generate install relabel restart nginx-install nginx-reload

clean:
	rm -f $(APP_NAME)
