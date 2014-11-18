# AQI [![Build Status](https://travis-ci.org/elvuel/go-aqi.png?branch=master)](https://travis-ci.org/elvuel/go-aqi)

[![wercker status](https://app.wercker.com/status/cffe1b567c592d1ee023dbe90915ab9b/m "wercker status")](https://app.wercker.com/project/bykey/cffe1b567c592d1ee023dbe90915ab9b)

**MEP**
>	* [**GB3095—2012**](http://kjs.mep.gov.cn/hjbhbz/bzwb/dqhjbh/dqhjzlbz/201203/t20120302_224165.htm)
	
>	* [**HJ633-2012** _Feb 2012_](http://www.es.org.cn/download/2012/1-6/2272-1.pdf)

**EPA**

>	* [**EPA-454/B-12-001** _Sep 2012_](http://www.epa.gov/airnow/aqi-technical-assistance-document-sep2012.pdf)

>	* EPA JAN/15/2013 rules
>	>	See: epa_jan15_2013.pdf

***

## Installation

***

> go get github.com/elvuel/go-aqi

## Usage

***

> checkout ./example/main.go

## BM

***
```
	$>go test -bench .
```
>| Name             | N                            | ns/op             |
> ----------------- | ---------------------------- | ------------------
>| EpaGetAQI        | 200000                       | 9271            |
>| MepGetAQI        | 500000                       | 6527            |
>| EpaGetPM25IAQI   | 1000000                      | 1240            |
>| MepGetPM25IAQI   | 20000000                     | 136             |


## TODO

***

* Helper Utils
* Self host webservice

## License

***

Under the MIT license