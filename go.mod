module github.com/realbucksavage/stargatecontrib

go 1.17

require (
	github.com/quipo/statsd v0.0.0-20180118161217-3d6a5565f314
	github.com/realbucksavage/stargate v0.0.0-20220416154459-e5dd74d21f0b
)

require github.com/pkg/errors v0.9.1 // indirect

replace github.com/realbucksavage/stargate => github.com/realbucksavage/stargate v0.0.0-20220826104108-acff1d98655f
