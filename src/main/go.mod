module example.com/main

go 1.16

replace example.com/gosample => ../gosample

require (
	example.com/gosample v0.0.0-00010101000000-000000000000
	example.com/httpsample v0.0.0-00010101000000-000000000000
)

replace example.com/httpsample => ../httpsample
