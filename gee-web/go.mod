module example

require gee v0.0.0

// It tells Go to replace the gee module with the local directory ./gee instead of downloading it from an external repository.
replace gee => ./gee

go 1.23.6
