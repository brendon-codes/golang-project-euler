
##
## Go-Euler Makefile
##
## Author: Brendon Crawford
##

GOPATH=$(CURDIR)

bin_dir=bin
out_prefix=euler-solution
src_dir=src
solutions_dir=solutions
ext=go


all: clean
	@for i in $$(find ${src_dir}/${solutions_dir}/ -maxdepth 1 -mindepth 1 -type d); do \
		echo "Building $$(basename $${i})."; \
		go build -o ${bin_dir}/${out_prefix}-$$(basename $${i}) $${i}/$$(basename $${i}).${ext}; \
	done;


solution.%: clean.%
	@echo "Building ${*}."
	@go build -o ${bin_dir}/${out_prefix}-${*} ${src_dir}/${solutions_dir}/${*}/${*}.${ext}


clean.%:
	@echo "Cleaning ${*}."
	@rm -f ${bin_dir}/${out_prefix}-${*}


clean:
	@for i in $$(find ${bin_dir}/ -mindepth 1 -maxdepth 1 -type f -name "${out_prefix}-*"); do \
		echo "Cleaning $$(echo $$(basename $${i}) | awk -F '-' '{ print $$NF }')."; \
		rm -f $${i}; \
	done;


