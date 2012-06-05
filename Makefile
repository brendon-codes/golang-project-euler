
GOPATH = $(CURDIR)

bin_dir = bin
out_prefix = euler-solution
solutions_dir = solutions

solution-1:
	rm -f ${bin_dir}/${out-prefix}-1
	go build -o ${bin_dir}/${out_prefix}-1 ${solutions_dir}/1

