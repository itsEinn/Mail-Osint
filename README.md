# MOSINT

#### Özellikler:

* Mail doğrulaması
* Socialscan ve Holehe ile sosyal medya hesabı kontrolü
* Veri ihlallerini ve parola sızıntılarını kontrol edin
* İlgili mailleri ve domainleri bulun
* Pastebin ve Throwbin dökümlerini tara
* Google araması
* DNS Araması
* IP Araması


## Servisler (APIs):

\[not required to run the program\]

| Service | Function | Status |
| :--- | :--- | :--- |
| [ipapi.co](https://ipapi.co/) - Public | Domain hakkında daha fazla bilgi | :white\_check\_mark: |
| [hunter.io](https://hunter.io/) - Public | İlgili mailler | :white\_check\_mark: :key: |
| [emailrep.io](https://emailrep.io/) - Public | İhlal edilen site adları | :white\_check\_mark: :key: |
| [scylla.so](https://scylla.so/) - Public | Database Sızıntıları | :construction: |
| [breachdirectory.org](https://breachdirectory.org/) - Public | Şifre Sızıntıları | :white\_check\_mark: :key: |
| [Intelligence X](https://intelx.io/)| Şifre sızıntıları | :white\_check\_mark: :key: |

:key: API key required

#### Kullanım için:

- API keyinizi `keys.json` a kaydedin
- Sisteme `Go` ve `Python` dillerini indirin

## İndirme:

`pip3 install -r requirements.txt`

## Kullanım:

Yardım menüsü için `-h` yazabilirsiniz

| KOMUTLAR  | AÇIKLAMA                                          | GEREKLİMİ? |
|-----------|---------------------------------------------------|------------|
| -e        | Hedef mail belirle                                | Evet       |
| -verify   | Hedef maili onayla                                | Hayır      |
| -social   | Hedef mail için sosyal medya taraması             | Hayır      |
| -relateds | Hedef mail ile önerilen domainler ile maille öğren| Hayır      |
| -leaks    | Hedef mail ile şifre sızıntılarını öğren          | Hayır      |
| -dumps    | Hedef mailin pastebin dökümlerini ara             | Hayır      |
| -domain   | Mailin domaini hakkında daha fazla bilgi          | Hayır      |
| -o        | txt için output(çıkış) Dosyası belirle            | Hayır      |
| -v        | Versiyon                                          | Hayır      |
| -h        | Yardım                                            | Hayır      |
| -all      | Hepsini kullan                                    | Hayır      |

### Örnek:

`go run main.go -e ornek@domain.com -all`

Output(çıkış) Dosyası için `-o` komutu (.txt)
