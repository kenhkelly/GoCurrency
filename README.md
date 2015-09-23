#GoCurrency

GoCurrency is used to quickly check the conversion rates

*Build it*: 

`go build gocurrency.go`

*Simple usage*

`./gocurrency [flag] [symbol]`

*Options*

- -base=USD The base currency to quote a currency against
- -help
- symbol    The symbol(s) to show, comma separate. If left empty, will show all

*Examples*:

```
$ ./gocurrency
Base currency: USD, Date: 2015-09-23
Symbol  | Rate
--------|----
AUD     | 1.418
BGN     | 1.754
BRL     | 4.017
CAD     | 1.327
CHF     | 0.976
CNY     | 6.383
CZK     | 24.296
DKK     | 6.691
EUR     | 0.897
GBP     | 0.654
HKD     | 7.750
HRK     | 6.801
HUF     | 278.840
IDR     | 14617.000
ILS     | 3.952
INR     | 65.981
JPY     | 120.210
KRW     | 1191.200
MXN     | 16.860
MYR     | 4.347
NOK     | 8.284
NZD     | 1.589
PHP     | 46.793
PLN     | 3.770
RON     | 3.963
RUB     | 66.011
SEK     | 8.412
SGD     | 1.421
THB     | 36.131
TRY     | 3.025
ZAR     | 13.658

```
$ ./gocurrency BRL,EUR,GBP
Base currency: USD, Date: 2015-09-23
Symbol  | Rate
--------|----
BRL     | 4.017
EUR     | 0.897
GBP     | 0.654

```
$ ./gocurrency -base=JPY BRL,EUR,GBP
Base currency: JPY, Date: 2015-09-23
Symbol  | Rate
--------|----
BRL     | 0.033
EUR     | 0.007
GBP     | 0.005
```

### Changelog

Version | Change
--------|----------
[v1.0]  | Initial version, providing the currency rates

[v1.0]: https://github.com/kenhkelly/GoCurrency/tree/v1.0