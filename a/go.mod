module example.com/a

go 1.22.0

require example.com/vulnerable v1.0.0
replace example.com/vulnerable => ../dep
