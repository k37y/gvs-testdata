module example.com/app

go 1.22.0

require example.com/vulnerable v1.0.0
replace example.com/vulnerable => ./dep

require example.com/missing v1.0.0
replace example.com/missing => ./missing
