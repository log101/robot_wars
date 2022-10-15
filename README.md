
# Robot Yarışı

Yarışmaya katılabilecek robotların en az 3 adet özel yeteneği olmalı ve her yeteneğin belirli sınırlarda düşmana hasar verme oranı bulunmalıdır. Yarışmalar kura usülüyle oluşturulacak karma bir ligde yapılmalıdır. Yarışma sırasında hamleler sırayla yapılmalı ve her robotun hamlesi anlık olarak Random seçilmelidir. Başlangıçta sistem tarafında tanımlı en az 10 adet robot bulunmalı ve kullanıcı tarfından yeni robotlar konsoldan eklenebilmelidir. Her yarışmanın sonucu adım adım ekranda gösterilmeli her aşama için kullanıcıya bilgi verilmelidir. Yarışma sonucunda kazanan robot ilan edilmelidir.

## Yapılacaklar

- [x] Varsayılan olarak en az 10 adet Robot sisteme yazılım aşamasında tanımlanacak
- [x] Yazılım başladığında kullanıcıya her robotun tüm özellikleri tek tek rapor edilecek
- [x] Kullanıcı isterse konsoldan (ftm.Scan...) gireceği komutlarla çalışma anında default robotların bilgilerini düzenleyebilecektir.
- [x] Kullanıcı isterse konsoldan kendi robotunu tanımlayabilecektir.
- [x] Kullanıcı komut verdiğinde karma bir lig oluşturulacak kullanıcıya sonuclar gösterilecektir. Kullanıcı isterse ligi yeniden oluşturabilecektir.
- [x] Kullanıcı başlat komutu verdiğinde oluşan lig de yarışmalar başlatılacak karşılıklı robotlar dövüştürülecek ve sonuçları adım adım ekranda yazılacaktır.
- [ ] Dövüş esnasında Robotlar hamleleri karşılıklı olarak sırayla yapacaktır. Her hamleden sonra Robotların hangi hamleyi yaptığı kalan güç durumları vs.. tek tek loglanarak ekranda gösterilecektir.


## Çalışma Prensibi

- [x] Yarışmalar başladığında Robotlar arasında Random olarak karma bir lig oluşturulacak. Buradaki yarışma adedi toplam robot sayısına göre dinamik oluşacak.
- [x] Robotların her birinin kendine ait özellikleri olacak ve en az 3 adet özel yeteneği ver herbirinin karşıya zarar verme oranları olacaktır.
- [x] Tüm lig oluşturma ve hamle seçme işlemleri Random olarak yapılacaktır.


## Notlar

- [x] Robot tanımlamalarında Struct kullanılacak.
- [x] Robotların özellikleri tanımlanırken "methodlar" kullanılacak (foknsiyonlar değil).
- [x] Robotlar in-memory map olarak hafızada tutulacak ( map kullanılacak ).
- [x] Log akışları sırasında time.Sleep ile log akışı yavaşlatılırsa daha gerçekci olacaktır.
- [x] Yazılım ilk başladığında nasıl çalıştığını kullanıcın neler yapabileceğini anlatan bir çıktıyı ekranda göstermelidir.
