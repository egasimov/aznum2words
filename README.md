# aznum2words

## Məqsəd:
Kitabxana tam vəya kəsr ədədlərin Azərbaycan dilində sözlə təsviri üçün nəzərdə tutulub.


Bu kitabxana vasitəsilə -  müsbət, mənfi tam vəya kəsr ədədlərin sözlə təsvirini əldə etmək mümkündür.


## Nümunə:

```go
result1, _ := aznum2words.SpellNumber("-95412")
result2, _ := aznum2words.SpellNumber("-2.7021")
result3, _ := aznum2words.SpellNumber("5611113210")

fmt.Print(resut1)
fmt.Print(resut2)
fmt.Print(resut3)
```

**NƏTİCƏ**:
```text
mənfi doxsan beş min dörd yüz on iki
mənfi iki tam on mində yeddi min iyirmi bir
beş milyard altı yüz on bir milyon bir yüz on üç min iki yüz on
```
- - -
## İstifadə qaydası
```shell
go get github.com/egasimov/aznum2words
```

- - -

## Test caseləri yoxlanması
```shell
go test ./
```