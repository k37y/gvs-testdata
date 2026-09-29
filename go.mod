module example.com/app

go 1.22.0

require example.com/vulnerable v1.0.1-0.20260101000000-abcdefabcdef
replace example.com/vulnerable => ./dep
