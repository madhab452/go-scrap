# PROVIDER data provider
export PROVIDER=nepse

# enum FILE, or INTERNET, DSRC=FILE, FILE_PATH is required. if DSRC=INTERNET, URL is required
export DSRC=FILE
export FILE_PATH=_examples/nepse/nepse_1.html
# export URL=http://www.nepalstock.com/main/todays_price/index/1/?startDate=&stock-symbol=&_limit=500 

# TARGET_URL -> the destination the scrapped data should be fed to, default is console
# export TARGET_URL=http://127.0.0.1:4040/v1/transactions
export TARGET_URL=


.PHONY: help

# ------------------
# Help
# ------------------
help: ## Show command list
	@echo "Usage:"
	@echo " make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-20s\033[0m %s\n", $$1, $$2}'


run: ## Run Scrapper
	go run main.go
