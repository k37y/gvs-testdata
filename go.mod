module example.com/app

go 1.22.0

require example.com/vulnerable v1.1.0-rc.1
replace example.com/vulnerable => ./dep
