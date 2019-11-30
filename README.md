# go-pi

Ponderal index calculator implemented in Go

## Installation

```bash
git clone git@github.com/usk81/go-pi.git
cd go-pi
go build -ldflags '-w -s' -o go-pi
```

## Usage

```bash
# go-pi height(cm2) weight(kg) age 
e.g. go-pi 180 70 20
> Classification: Normal
> Index: 21.604938
> IndexType: Body Mass Index
> Status: Normal
```

## Index Type

| age           | Index type            |
|---------------|-----------------------|
| age < 5       | Kaup Indx             |
| 5 <= age < 15 | Rohrer Index          |
| age >= 15     | BMI (Body Mass Index) |