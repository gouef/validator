# Available Validators

This package includes several predefined validators:

- `Blank`: Ensures the value is blank (`""` or `nil`).
- `NotBlank`: Ensures the value is not blank.
- `Nil`: Ensures the value is `nil`.
- `NotNil`: Ensures the value is not `nil`.
- `IsTrue`: Ensures the value is `true`.
- `IsString`: Ensures the value is `string`.
- `IsFloat`: Ensures the value is `float`.
- `IsFloat32`: Ensures the value is `float32`.
- `IsFloat64`: Ensures the value is `float64`.
- `IsInt`: Ensures the value is `int`.
- `IsInt8`: Ensures the value is `int8`.
- `IsInt16`: Ensures the value is `int16`.
- `IsInt32`: Ensures the value is `int32`.
- `IsInt64`: Ensures the value is `int64`.
- `IsUint`: Ensures the value is `uint`.
- `IsUint8`: Ensures the value is `uint8`.
- `IsUint16`: Ensures the value is `uint16`.
- `IsUint32`: Ensures the value is `uint32`.
- `IsUint64`: Ensures the value is `uint64`.
- `IsStruct`: Ensures the value is `struct`.
- `IsFalse`: Ensures the value is `false`.
- `IsBool`: Ensures the value is a boolean (`true` or `false`).
- `Struct`: Ensures the value is `struct` of type ``. ([Example](#struct))
- `Currency`: Ensures the value is a valid currency code.
- `Email`: Ensures the value is a valid email address.
- `Language`: Ensures the value is a valid language code without a region (e.g., `en`).
- `LanguageWithRegion`: Ensures the value is a valid language code with a region (e.g., `en-US`).
- `RangeConstraint`: Ensures the value is within a defined range (e.g., a number between `min` and `max`). It support int and float. ([Example](#range))
- `ArrayRange`: Ensures the value is within a defined array range (e.g., a number between `min` and `max`). It support int and float. ([Example](#arrayrange))
- `Options`: Ensures the value is within a defined array. ([Example](#options))
- `ArrayOptions`: Ensures the values (array) is within a defined array. ([Example](#arrayoptions))
- `Required`: Ensures the value is not blank or `nil`.
- `RegularExpression`: Ensures the value match regular expression. ([Example](#regularexpression))
- `Date`: Ensures the value is valid date. ([Example](#date))

## Examples

### Struct

```go
con := constraints.Struct{
		Struct: MyStruct{},
	}

errs := validator.Validate(value, con)
```

### RegularExpression

```go
regExp := `^[a-z]{3}$`
con := constraints.RegularExpression{
    Regexp: regExp,
}

errs := validator.Validate(value, con)
```

### Range

```go
con := constrains.Range{
	Min: 5,
	Max: 20,
}

value := 7

errs := validator.Validate(value, con)
```

### ArrayRange

```go
con := constrains.ArrayRange{
	Min: 5,
	Max: 20,
}

value := []any{5, 15, 19}

errs := validator.Validate(value, con)
```

### Options

```go
con := constraints.Options{
		Options: []any{"option1", "option2", "option3", 42},
	}

value := "option1"

errs := validator.Validate(value, con)
```

### ArrayOptions

```go
con := constraints.ArrayOptions{
		Options: []any{"Red", "Green", "Blue", 100},
	}
	
values := []any{"Red", "Green", 100}
errs := validator.Validate(values, con)
```

### Date

```go
dateFormat := "2006-01-02"
minDate, _ := time.Parse(dateFormat, "2023-01-01")
maxDate, _ := time.Parse(dateFormat, "2023-12-31")

con := constraints.Date{
    Min:    minDate,
    Max:    maxDate,
    Format: dateFormat,
}

value := "01-05-2023"
errs := validator.Validate(value, con)
```
