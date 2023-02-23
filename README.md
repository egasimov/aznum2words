# AzNum2Words ![CI](https://github.com/egasimov/aznum2words/actions/workflows/ci.yml/badge.svg?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/egasimov/aznum2words)](https://goreportcard.com/report/github.com/egasimov/aznum2words)

## Məlumat | Description
**AzNum2Words** - açıq qaynaq kodlu, [MIT](./LICENSE) lisenziyalı Go proqramlaşdırma dilində yazılmış bir kitabxanadır.


## Məqsəd | Goal
Kitabxana tam vəya kəsr ədədlərin Azərbaycan dilində sözlə təsviri üçün nəzərdə tutulub.


Bu kitabxana vasitəsilə -  müsbət, mənfi tam vəya kəsr ədədlərin sözlə təsvirini əldə etmək mümkündür.

## İstifadə yerləri | Use cases
Azərbaycan dilində ədədlərin təsvirinə ehtiyac duyulan hallarda istifadə oluna bilər.

* Hesabatların tərtib olunması
* Bank əməliyyatları zamanı məbləğin sözlə təsvir olunması və s.

## Funksiyaları | Functionalities
* Tam ədədlərin sözlə təsvir olunması
* Kəsr ədədlərin sözlə təsvir olunması
* Müsbət vəya mənfi ədədlərin sözlə təsvir olunması
* Söz ilə təsvir oluna biləcək maksimum tam ədəd: 10^63
* Söz ilə təsvir oluna biləcək minumum kəsr ədəd: (10^-15)

## Nümunə | Sample

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

**NƏTİCƏ | OUTPUT**
```text
mənfi doxsan beş min dörd yüz on iki
mənfi iki tam on mində yeddi min iyirmi bir
beş milyard altı yüz on bir milyon bir yüz on üç min iki yüz on
```
- - -
## Quraşdırılması | Installation
```shell
go get github.com/egasimov/aznum2words
```

- - -

## Test caseləri yoxlanması | Check test cases
```shell
go test ./
```
- - -
## Lisenziya | License
Kitabxana MIT Lisenziya altında lisenziyalaşdırılmışdır. Ətraflı məlumat üçün 
[LICENSE](./LICENSE) faylını nəzərdən keçirin. 

- - -
## Versiyalar | Releases
Kitabxanaının versiyaları [Semver](http://semver.org/) yanaşması ilə tənzimlənir.

- - -
## Proyektə necə dəstək olmaq olar | How to contribute to project ? 

Proyektə contribute etmək üçün aşağıdakı təlimatları nəzərə ala bilərsiniz.

*Testləri olmayan vəya nəzərə alınmayan PRlar qəbul edilməyəcək*

1. Reponu fork et
2. Yeni feature branch yarat (`git checkout -b my-new-feature`)
3. Dəyişiklikləri commit et (`git commit -am 'Added some feature'`)
4. Local branchı origin(remote repo) push et  (`git push origin my-new-feature`)
5. Yeni Pull Request yarat 

- - - 

## Contributorlar

Bu layihəyə aşağıdakı şəxslər töhfə verib:

<!-- Contributors list -->
<a href="https://github.com/egasimov/aznum2words/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=egasimov/aznum2words" />
</a>

<!--Made with [contrib.rocks](https://contrib.rocks). -->
<!-- Contributors list -->
