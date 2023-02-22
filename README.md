# aznum2words

## Məlumat
**aznum2words** - açıq qaynaq kodlu, [MIT](./LICENSE) lisenziyalı Go proqramlaşdırma dilində yazılmış bir kitabxanadır.


## Məqsəd
Kitabxana tam vəya kəsr ədədlərin Azərbaycan dilində sözlə təsviri üçün nəzərdə tutulub.


Bu kitabxana vasitəsilə -  müsbət, mənfi tam vəya kəsr ədədlərin sözlə təsvirini əldə etmək mümkündür.

## İstifadə yerləri
Azərbaycan dilində ədədlərin təsvirinə ehtiyac duyulan hallarda istifadə oluna bilər.

* Hesabatların tərtib olunması
* Bank əməliyyatları zamanı məbləğin sözlə təsvir olunması və s.

## Funksiyaları
* Tam ədədlərin sözlə təsvir olunması
* Kəsr ədədlərin sözlə təsvir olunması
* Müsbət vəya mənfi ədədlərin sözlə təsvir olunması
* Söz ilə təsvir oluna biləcək maksimum tam ədəd: 10^63
* Söz ilə təsvir oluna biləcək minumum kəsr ədəd: (10^-15)

## Nümunə

```go
package main

import (
	"fmt"

	"github.com/egasimov/aznum2words"
)

func main() {
	result1, _ := aznum2words.SpellNumber("-95412")
	result2, _ := aznum2words.SpellNumber("-2.7021")
	result3, _ := aznum2words.SpellNumber("5611113210")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

```

**NƏTİCƏ**
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
- - -
## Lisenziya
Kitabxana MIT Lisenziya altında lisenziyalaşdırılmışdır. Ətraflı məlumat üçün 
[LICENSE](./LICENSE) faylını nəzərdən keçirin. 

- - - 
## Contributorlar

Bu layihəyə aşağıdakı şəxslər töhfə verib:

<!-- Contributors list -->
<a href="https://github.com/egasimov/aznum2words/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=egasimov/aznum2words" />
</a>

<!--Made with [contrib.rocks](https://contrib.rocks). -->
<!-- Contributors list -->
