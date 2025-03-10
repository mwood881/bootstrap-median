## Bootstrap Sampling Method

# Description: 
This repository uses a bootstrap sampling method to estimate the standard error of the median in R and Go. 
It will compare the two programming languages and use logging, profiling, and unit tests in Go to analyze the results combared to R. 

# The statistical method:
Bootstrap resampling which helps estimate the standard error of the median. 

# Installation:
R: used boot package 
rlnorm helped with the log normal data 
also uses the Merseene Twister algorithm and created csv file for data so that the same dataset was used in Go for a fair comparison. 

Go: mt19937 was used for the MErseen Twister: seehuhn/mt19937
in order to profile, you needed to install: net/http/pprof

# Setup
Go: 
(1)
``` 
install.packages("boot")
```
(2) Save the dataset as data.csv

(3) Run the R script to find the results of the bootstramp standard error of median: 

ORDINARY NONPARAMETRIC BOOTSTRAP


Call:
boot(data = data, statistic = median_func, R = 1000)


Bootstrap Statistics :
     original       bias    std. error
t1* 0.9869517 -0.006450676  0.03261068
> print(paste("Bootstrap Standard Error of Median:", sd(boot_result$t)))
[1] "Bootstrap Standard Error of Median: 0.0326106768989483"


GO: 
(1) 
``` bash
git clone https://github.com/mwood881/bootstrap-median
cd https://github.com/mwood881/bootstrap-median
go mod tidy
```

(2) Run main.go
``` bash
go run main.go
```

(3) Test the benchmark
``` bash
go test -bench=.
```

(4) Results: 
Bootstrap Standard Error of Median: 0.03269


## Comparison:
They both got the same results which was showing consistency even across programming languages.
I would reccommend Go over R because of logging and profiling being able to process the times faster. The saving of not using cloud space also could be a positive of Go. 

## Resources used:
I first was getting a different number for the Go code that felt off so I had to use copilot to help fix it and realized I actually needed to use the same dataset and converted the csv data set to the Go file from the R file. 
