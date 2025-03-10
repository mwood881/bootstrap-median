# boot: https://cran.r-project.org/web/packages/boot/index.html
# Load required package
library(boot)

# Define a function to compute the median
median_func <- function(data, indices) {
  return(median(data[indices]))
}
# had to use the MT twister from canvas 
set.seed(42, kind = "Mersenne-Twister")

# Generate synthetic log-normal data
# had to load into csv to use same dataset in go to compare results
data <- rlnorm(1000, meanlog = 0, sdlog = 1)

write.csv(data, "data1.csv", row.names = FALSE)


# Perform bootstrap resampling (1000 iterations)
boot_result <- boot(data, statistic = median_func, R = 1000)

# Print standard error of the median
print(boot_result)
print(paste("Bootstrap Standard Error of Median:", sd(boot_result$t)))


## Results: 
## ORDINARY NONPARAMETRIC BOOTSTRAP


#Call:
#  boot(data = data, statistic = median_func, R = 1000)


#Bootstrap Statistics :
#  original       bias    std. error
#t1* 0.9869517 -0.006450676  0.03261068
#> print(paste("Bootstrap Standard Error of Median:", sd(boot_result$t)))
#[1] "Bootstrap Standard Error of Median: 0.0326106768989483"
