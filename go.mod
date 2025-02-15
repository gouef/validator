module github.com/gouef/validator

go 1.23.4

replace github.com/gouef/validator/constraints => ./constraints

require (
	github.com/gouef/country v1.0.0
	github.com/gouef/currency v1.0.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/text v0.21.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
