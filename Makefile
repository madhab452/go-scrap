export PROVIDER=nepse
export DSRC=file
export URL=http://www.nepalstock.com/main/todays_price/index/1/?startDate=&stock-symbol=&_limit=500
export FILE_PATH=_examples/nepse/Nepal Stock Exchange Ltd_1.html

help:
	@echo "help"

run:
	go run main.go