# PROVIDER data provider
export PROVIDER=nepse

export SRC=_examples/nepse/nepse_1.html
# export SRC=http://www.nepalstock.com/main/todays_price/index/1/?startDate=&stock-symbol=&_limit=500

# export TARGET=http://127.0.0.1:4040/v1/transactions
export TARGET=



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
