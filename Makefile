
GOPATH=$(CURDIR)

bin_dir=bin
out_prefix=euler-solution
src_dir=src
solutions_dir=solutions
ext=go

all:
	for i in $$(find ${src_dir}/${solutions_dir}/ -maxdepth 1 -mindepth 1 -type d); do \
		go build -o ${bin_dir}/${out_prefix}-$$(basename $${i}) $${i}/$$(basename $${i}).${ext}; \
	done;

solution.%:
	rm -f ${bin_dir}/${out_prefix}-${*}
	go build -o ${bin_dir}/${out_prefix}-${*} ${src_dir}/${solutions_dir}/${*}/${*}.go

