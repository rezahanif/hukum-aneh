# Dokumen Skeleton: 1 File Per Direktori

Full text skeleton (dengan anotasi struktural) dari 1 file sampel per direktori.
Tujuan: mempelajari pola penulisan dan struktur dokumen **per institusi penerbit**.

## Legenda Anotasi

| Marker | Arti |
|--------|------|
| `B` | Bold text |
| `Fn` | Font size berbeda dari body |
| `I1`, `I2`, etc | Indent level |
| `[PREAMBLE:...]` | Bagian pembuka |
| `[KEPUTUSAN:...]` | Bagian keputusan/amar |
| `[HEADING:...]` | Judul struktural (BAB, Bagian, etc) |
| `[PASAL]` | Baris Pasal |
| `[AYAT]` | Baris Ayat |
| `[ITEM]` | Item bernomor |
| `[SUB-ITEM]` | Sub-item berhuruf |

## Ringkasan Tipe Dokumen

| Direktori | Tipe | Penerbit |
|-----------|------|----------|
| `uu` | Undang-Undang (Statute) | DPR + Presiden |
| `pp` | Peraturan Pemerintah | Presiden |
| `perppu` | Perppu (Emergency Regulation) | Presiden |
| `perpres` | Peraturan Presiden | Presiden |
| `perda` | Peraturan Daerah (Regional Reg) | Kepala Daerah |
| `keppres` | Keputusan Presiden (Decision) | Presiden |
| `inpres` | Instruksi Presiden (Instruction) | Presiden |
| `tap_mpr` | Ketetapan MPR (MPR Decree) | MPR |
| `uud-1945` | UUD 1945 (Constitution) | BPUPKI/PPKI |
| `Putusan-MK` | Putusan MK (Court Ruling) | Mahkamah Konstitusi |
| `JDIH_Kemnaker` | Peraturan Menteri | Menteri Ketenagakerjaan |
| `JDIH_Kemenkeu` | Peraturan Menteri | Menteri Keuangan |
| `JDIH_Kemendag` | Keputusan Menteri (Decision) | Menteri Perdagangan |
| `JDIH_Komdigi` | Peraturan Menteri | Menteri Kominfo |
| `JDIH_KPU` | Peraturan KPU | Komisi Pemilihan Umum |
| `peraturan` | Peraturan Pemerintah (simple) | Presiden |

---

## uu

- **File**: `uu/uunomor41tahun2014.pdf`
- **Document Type**: Undang-Undang (Statute)
- **Issued by**: DPR + Presiden
- **Pages**: 43 | **Lines**: 1717
- **Font sizes**: [8.5, 9.0, 10.0, 10.5, 11.0, 11.5, 12.0, 12.5, 13.0, 13.5, 14.0, 14.5, 15.0, 15.5, 16.0, 16.5, 17.0, 17.5, 18.0, 19.5, 22.0, 26.0, 38.5, 40.0, 44.0, 47.5, 48.5, 49.0]
- **Most common font**: 13.0 (15% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [66.0, 351.0, 379.0, 516.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01   F11  | PRESIOEN
p01   F12  | R EPLJBLIK INDONESIA
p01   F12  | UNDANG-UNDANG REPUBLIK INDONESIA
p01   F12  | NOMOR 41 TAHUN 2014
p01   F12  | TENTANG
p01   F12  | PERUBAHAN ATAS UNDANG-UNDANG NOMOR 18 TAHUN 2OO9
p01   F12  | TENTANG PETERNAKAN DAN KESEHATAN HEWAN
p01   F12  [PREAMBLE:DENGAN RAHMAT]
p01   F12  | DENGAN RAHMAT TUHAN YANG MAHA ESA
p01   F12  | PRESIDEN REPUBLIK INDONESIA,
p01   F18  I1 [PREAMBLE:MENIMBANG]
p01   F18  I1 | Menimbang : a. bahwa negara bertanggung jawab untuk melindungi
p01   F16  | segenap bangsa Indonesia dan seluruh tumpah Carah
p01   F16  | Indonesia melalui penyelenggaraan peternakan dan
p01   F14  | kesehatan hewan dengan mengamankal dan menjamin
p01   F15  | pemanfaatan dan pelestarian hewan untuk mewujudkan
p01   F14  | kedaulatan, kemandirian, serta ketahanan pangan dalam
p01   F17  | rangka menciptakan kesejahteraan dan kemakmuran
p01   F14  | seluruh ralgrat Indonesia sesuai dengan amanat Undang-
p01   F14  | Undang Dasar Negara Republik Indonesia Tahun 1945;
p01   F18  [SUB-ITEM]
p01   F18  | b. bahwa dalam penyelenggaraan petemakan dan kesehatan
p01   F18  | hewan, upaya pengamanan maksimal terhadap
p01   F14  | pemasukan dan pengeluaran ternak, hewan, dan produk
p01   F18  | hewan, pencegahan penyakit hewan dan zoonosis,
p01   F16  | penguatan otoritas veteriner, persyaratan halal bagi
p01   F16  | produk hewan yang dipersyaratkan, serta penegakan
p01   F14  | hukum terhadap pelanggaran kesejahteraan hewan, perlu
p01   F18  | disesuaikan dengan perkembangan dan kebutuhan
p01        | masyarakat;
p01   F14  | bahwa Undang-Undang Nomor 18 Tahun 2009 tentang
p01   F14  | Peternakan dan Kesehatan Hewan dipandang tidak sesuai
p01   F15  | lagi dan perlu disempurnakan untuk dijadikan landasan
p01   F14  | hukum bagi penyelenggaraan peternakan dan kesehatan
p01   F12  | hewan;
p01   F18  | bahwa berdasarkan pertimbangan sebagaimana
p01   F18  | dimaksud dalam huruf a, huruf b, dan huruf c perlu
p01   F15  | membcntuk Undang-Undang tentang perubahan Atas
p01   F18  | Undang-Undang Nomor 18 Tahun 2OOg tentang
p01        | Peternakan dan Kesehatan Hewan;
p01   F12  [SUB-ITEM]
p01   F12  | d.
p01   F14  [PREAMBLE:MENGINGAT]
p01   F14  | Mengingat...
==================== PAGE 2 ====================
p02   F11  | PRESIOEN
p02   F14  | R EPI-IEIL IK INOONESIA
p02   F18  | -2-
p02   F18  I1 [PREAMBLE:MENGINGAT]
p02   F18  I1 | Mengingat : Pasal 20 dan Pasal 21 Undang-Undang Dasar Negara
p02   F14  | Republik Indonesia Tahun 1945;
p02        | Dengan Persetujuan Bersama
p02   F12  | DEWAN PERWAKILAN RAKYAT REPUBLIK INDONESIA
p02        | dan
p02   F12  | PRESIDEN REPUBLIK INDONESIA
p02   F12  | MEMUTUSI(AN:
p02   F18  I1 [PREAMBLE:MENETAPKAN]
p02   F18  I1 | Menetapkan : UNDANG-UNDANGTENTANG
p02        | UNDANG NOMOR 18 TAHUN
p02   F12  | DAN KESEHATAN HEWAN.
p02        I2 | PERUBAHAN ATAS UNDANG-
p02   F12  I2 | 2OO9 TENTANG PETERNAKAN
p02   F12  | Pasal I
p02   F14  | Beberapa ketentuan dalam Undang-Undang Nomor 18 Tahun
p02   F14  | 2009 tentang Peternakan dan Kesehatan Hewan (Lembaran
p02   F17  | Negara Republik Indonesia Tahun 2OO9 Nomor 84,
p02   F15  | Tambahan Lembaran Negara Republik Indonesia Nomor
p02   F14  | 5015), diubah sebagai berikut:
p02   F18  [ITEM]
p02   F18  | 1. Ketentuan Pasal 1 angka 1, angka 2, angka 12, angka
p02   F14  | 14, angka 15, angka 19, angka 21, angka 23, angka 24,
p02   F16  | angka 25, angka 26, angka 28, angka 29, angka 30,
p02   F16  | angka 34, angka 35, angka 36, angka 39, angka 40,
p02   F17  | angka 41, angka 46, dan angka 49 diubah, di antara
p02   F18  | angka 5 dan angka 6 disisipkan 2 (dua) angka yakni
p02   F17  | angka 5a dan 5b, di antara angka 37 dan angka 38
p02   F14  | disisipkan 1 (satu) angka yakni angka 37a, dan angka 9,
p02   F14  | angka 17, angka 20, angka 33, serta angka 44 dihapus,
p02   F14  | sehingga Pasal 1 berbunyi sebagai berikut:
p02   F12  | Pasal 1
p02   F14  | Dalam Undang-Undang ini yang dimaksud dengan:
p02   F18  [SUB-ITEM]
p02   F18  | l. Peternakan adalah segala urusan yang berkaitan
p02   F16  | dengan sumber daya fisik, Benih, Bibit, Bakalan,
p02   F14  | Ternak Ruminansia Indukan, Pakan, Alat dan Mesin
p02   F14  | Peternakan, budi daya Ternak, panen, pascapanen,
p02        | pengolahan, pemasaran, pengusahaan, pembiayaan,
p02        | serta sarana dan prasarana.
p02   F17  [ITEM]
p02   F17  | 2. Kesehatan...
==================== PAGE 3 ====================
p03   F11  | .).
p03   F12  | FRESIDEN
p03   F12  | REPI,IEILIK INDONESIA
p03   F18  | -3-
p03   F18  [ITEM]
p03   F18  | 2. Kesehatan Hewan adalah segala urusan yang
p03   F14  | berkaitan dengan pelindungan sumber daya Hewan,
p03   F18  | kesehatan masyarakat, dan lingkungan serta
p03   F18  | penjaminan keamanan
p03   F18  | Produk
p03   F12  I4 | Hewan,
p03   F12  [ITEM]
p03   F12  | 7.
p03   F14  | Kesejahteraan Hewan, dan peningkatan akses pasar
p03   F16  | untuk mendukung kedaulatan, kemandirian, dan
p03        | ketahanan pangan asal Hewan.
p03   F16  | Hewan adalah binatang atau satwa yang seluruh
p03   F15  | atau sebagian dari siklus hidupnya berada di darat,
p03   F16  | air, dan/atau udara, baik yang dipelihara maupun
p03   F14  | yang di habitatnya.
p03   F14  | Hewan Peliharaan adalah Hewan yang kehidupannya
p03   F16  | untuk sebagian atau seluruhnya bergantung pada
p03   F14  | manusia untuk maksud tertentu.
p03   F14  | Ternak adalah Hewan peliharaan yang produknya
p03   F14  | diperuntukan sebagai penghasil pangan, bahan baku
p03   F15  | industri, jasa, dan/atau hasil ikutannya yang terkait
p03        | dengan pertanian.
p03   F14  | Ternak Ruminansia Betina Produktif adalah Ternak
p03   F14  | ruminansia betina yang organ reproduksinya masih
p03   F14  | berfungsi secara normal dan dapat beranak.
p03   F14  | Ternak Ruminansia Indukan adalah Ternak betina
p03   F14  | bukan bibit yang memiliki organ reproduksi normal
p03   F14  | dan sehat digunakan untuk pengembangbiakan.
p03   F15  | Satwa Liar adalah semua binatang yang hidup di
p03   F15  | darat, air, dan/atau udara yang masih mempunyai
p03   F17  | sifat liar, baik yang hidup bebas maupun yang
p03        | dipelihara oleh manusia.
p03   F14  | Sumber Daya Genetik adalah material tumbuhan,
p03   F14  | binatang, atau jasad renik yang mengandung unit-
p03   F18  | unit yang berfungsi sebagai pembawa sifat
p03   F18  | keturunan, baik yang bernilai aktual maupun
p03   F16  | potensial untuk menciptakan galur, rumpun, atau
p03        | spesies baru.
p03   F14  | Benih Hewan yang selanjutnya disebut Benih adalah
p03   F14  | bahan reproduksi Hewan yang dapat berupa semen,
p03   F14  | sperma, ova, telur tertunas, dan embrio.
p03        | Dihapus.
p03   F12  [ITEM]
p03   F12  | 4.
p03   F12  [ITEM]
p03   F12  | 5.
p03   F12  | 5a.
p03   F12  | 5b.
p03   F12  [ITEM]
p03   F12  | 6.
p03   F12  [ITEM]
p03   F12  | 8.
p03   F12  [ITEM]
p03   F12  | 9.
p03   F15  [ITEM]
p03   F15  | 10. Bibit...
==================== PAGE 4 ====================
p04   F11  [ITEM]
p04   F11  | 10.
p04   F12  [ITEM]
p04   F12  | 12.
p04   F12  [ITEM]
p04   F12  | 15.
p04   F11  [ITEM]
p04   F11  | 16.
p04   F11  [ITEM]
p04   F11  | 11.
p04   F11  | PRESIDEN
p04   F14  | R EFLIBL IK INDONESIA
p04   F18  | -4-
p04   F16  | Bibit Hewan yang selanjutnya disebut Bibit adaiah
p04   F18  | Hewan yang mempunyai sifat unggul dan
p04   F14  | mewariskan serta memenuhi persyaratan tertentu
p04   F14  | untuk dikembangbiakkan.
p04   F14  | Rumpun Hewan yang selanjutnya disebut Rumpun
p04   F16  | adalah segolongan hewan dari suatu spesies yang
p04   F16  | mempunyai ciri-ciri fenotipe yang khas dan dapat
p04   F14  | diwariskan pada keturunannya.
p04   F18  | Bakalan Ternak Ruminansia Pedaging yang
p04   F18  | selanjutnya disebut Bakalan adalah ternak
p04   F14  | ruminansia pedaging dewasa yang dipelihara selama
p04   F16  | kurun waktu tertentu hanya untuk digemukkan
p04   F14  | sampai mencapai bobot badan maksimal pada umur
p04   F14  | optimal untuk dipotong.
p04   F14  [ITEM]
p04   F14  | 13. Produk Hewan adaiah semua bahan yang berasal
p04   F14  | dari Hewan yang masih segar dan/atau telah diolah
p04   F18  | atau diproses untuk keperluan konsumsi,
p04   F14  | farmakoseutika, pertanian, dan/atau kegunaan lain
p04   F18  | bagi pemenuhan kebutuhan dan kemaslahatan
p04        | manusia.
p04   F14  [ITEM]
p04   F14  | 14. Peternak adalah orang perseorangan warga negara
p04   F16  | Indonesia atau korporasi yang melakukan usaha
p04   F12  | Peternakan.
p04   F14  | Perusahaan Peternakan adalah orang perseorangan
p04   F14  | atau korporasi, baik yang berbentuk badan hukum
p04   F14  | maupun yang bukan badan hukum, yang didirikan
p04   F14  | dan berkedudukan dalam wilayah Negara Kesatuan
p04   F18  | Republik Indonesia yang mengelola usaha
p04   F14  | Peternakan dengan kriteria dan skala tertentu.
p04   F16  | Usaha di bidang Peternakan adalah kegiatan yang
p04   F16  | menghasilkan produk dan jasa yang menunjang
p04   F14  | usaha budi daya Ternak.
p04        | Dihapus.
p04   F14  | Inseminasi Buatan adalah teknik memasukkan mani
p04   F14  | atau semen ke dalam alat reproduksi Ternak betina
p04   F17  | sehat untuk dapat membuahi sel telur dengan
p04   F15  | menggunakan alat inseminasi dengan tujuan agar
p04   F14  | Ternak bunting.
p04   F14  | t7.
p04   F12  [ITEM]
p04   F12  | 18.
p04   F14  [ITEM]
p04   F14  | 19. Pemuliaan ...
==================== PAGE 5 ====================
p05   F11  [ITEM]
p05   F11  | 19.
p05   F11  | PRESIDEN
p05   F14  | R EFI.IBI," IK INDONESIA
p05   F18  | -5-
p05   F18  | Pemuliaan Ternak yang selanjutnya disebut
p05   F17  | Pemuliaan adalah rangkaian kegiatan untuk
p05   F16  | mengubah komposisi genetik pada sekelompok
p05   F18  | Ternak dari suatu rumpun atau galur guna
p05   F14  | mencapai tujuan tertentu.
p05        | Dihapus.
p05   F16  | Usaha di bidang Kesehatan Hewan adalah kegiatan
p05   F16  | yang menghasilkan produk dan/atau jasa yang
p05   F15  | menunjang upaya dalam mewujudkan Kesehatan
p05   F12  | Hewan.
p05   F18  | Pakan adalah bahan makanan tunggal atau
p05   F16  | campuran, baik yang diolah maupun yang tidak
p05   F18  | diolah, yang diberikan kepada hewan untuk
p05   F14  | kelangsungan hidup, berproduksi, dan berkembang
p05   F12  | biak.
p05   F18  | Bahan Pakan adalah bahan hasll pertanian,
p05   F16  | perikanan, Peternakan, atau bahan lain serta yang
p05   F14  | layak dipergunakan sebagai Pakan, baik yang telah
p05   F14  | diolah maupun yang belum diolah.
p05   F16  | Kawasan Penggembalaan Umum adalah lahan
p05   F14  | negara atau yang disediakan Pemerintah atau yang
p05   F14  | dihibahkan oleh perseorangan atau perusahaan yang
p05   F14  | diperuntukkan penggembalaan Ternak masyarakat
p05   F18  | skala kecil sehingga Ternak dapat leluasa
p05        | berkembang biak.
p05   F16  | Setiap Orang adalah orang perseorangan atau
p05   F14  | korporasi, baik yang berbadan hukum maupun yang
p05   F18  | tidak berbadan hukum serta yang melakukan
p05   F18  | kegiatan di bidang Peternakan dan Kesehatan
p05   F12  | Hewan.
p05   F16  | Veteriner adalah segala urusan yang berkaitan
p05        | dengan Hewan, Produk Hewan, dan Penyakit Hewan.
p05   F15  | Medik Veteriner adalah penyelenggaraan kegiatan
p05   F14  | praktik kedokteran hewan.
p05   F18  | Otoritas
p05   F18  | Veteriner
p05   F16  | adalah kelembagaan
p05   F18  | Pemerintah atau Pemerintah Daerah yang
p05   F14  | bertanggung jawab dan memiliki kompetensi dalam
p05   F12  | penyelenggaraan Kesehatan Hewan.
p05        [ITEM]
p05        | 20.
p05        [ITEM]
p05        | 21.
p05   F12  [ITEM]
p05   F12  | 22.
p05   F12  [ITEM]
p05   F12  | 23.
p05   F12  [ITEM]
p05   F12  | 27.
p05   F12  [ITEM]
p05   F12  | 24.
p05   F14  [ITEM]
p05   F14  | 25.
p05   F12  [ITEM]
p05   F12  | 26.
p05   F14  [ITEM]
p05   F14  | 28.
p05   F16  [ITEM]
p05   F16  | 29. Dokter ...
==================== PAGE 6 ====================
p06   F11  | PRESIDEN
p06   F14  | R ETJURL IK INDONESIA
p06   F18  | -6-
p06   F15  [ITEM]
p06   F15  | 29. Dokter Hewan adalah orang yang memiliki profesi di
p06   F16  | bidang kedokteran hewan dan kewenangan Medik
p06        | Veteriner dalam melaksanakan pelayanan Kesehatan
p06   F12  | Hewan.
p06   F16  [ITEM]
p06   F16  | 30. Dokter Hewan Berwenang adalah Dokter Hewan yang
p06   F14  | ditetapkan oleh Menteri, gubernur, atau bupati/wali
p06   F16  | kota sesuai dengan kewenangannya berdasarkan
p06   F17  | jangkauan tugas pelayanannya dalam rangka
p06   F12  | penyelenggaraan Kesehatan Hewan.
p06   F15  [ITEM]
p06   F15  | 31. Medik Reproduksi ada-lah penerapan Medik Veteriner
p06   F16  | dalam penyelenggaraan Kesehatan Hewan di bidang
p06   F14  | reproduksi hewan.
p06   F15  [ITEM]
p06   F15  | 32. Medik Konservasi adalah penerapan Medik Veteriner
p06   F16  | dalam penyelenggaraan Kesehatan Hewan di bidang
p06        | konservasi Satwa Liar.
p06   F15  [ITEM]
p06   F15  | 33. Dihapus.
p06   F16  [ITEM]
p06   F16  | 34. Penyakit Hewan adalah gangguan kesehatan pada
p06   F14  | Hewan yang disebabkan oleh cacat genetik, proses
p06   F16  | degeneratif, gangguan metabolisme, trauma,
p06   F17  | keracunan, infestasi parasit, prion, dan infeksi
p06        | mikroorganisme patogen.
p06   F16  [ITEM]
p06   F16  | 35. Penyakit Hewan Menular adalah penyakit yang
p06   F16  | ditularkan antara Hewan dan Hewan, Hewan dan
p06   F14  | manusia, serta Hewan dan media pembawa Penyakit
p06   F18  | Hewan lain melalui kontak langsung atau tidak
p06   F15  | langsung dengan media perantara mekanis seperti
p06   F16  | air, udara, tanah, Pakan, peralatan, dan manusia,
p06   F16  | atau melalui media perantara biologis seperti virus,
p06   F14  | bakteri, amuba, atau jamur.
p06   F16  [ITEM]
p06   F16  | 36. Penyakit Hewan Menular Strategis adalah Penyakit
p06   F14  | Hewan yang dapat menimbulkan angka kematian
p06   F14  | dan/atau angka kesakitan yang tinggi pada Hewan,
p06   F14  | dampak kerugian ekonomi, keresahan masyarakat,
p06   F14  | dan/atau bersifat zoonotik.
p06   F15  [ITEM]
p06   F15  | 37. Zoonosis adalah penyakit yang dapat menular dari
p06   F14  | Hewan kepada manusia atau sebaliknya.
p06   F14  | 37a. Wabah...
==================== PAGE 7 ====================
p07   F11  | PRESIDEN
p07   F12  | REPUBLIK INDONESIA
p07   F16  | -7 -
p07   F16  | 37a.Wabah adalah kejadian penyakit luar biasa yang
p07   F16  | dapat berupa timbulnya suatu Penyakit Hewan
p07   F14  | Menular baru di suatu wilayah atau kenaikan kasus
p07   F18  | Penyakit Hewan Menular mendadak yang
p07   F14  | dikategorikan sebagai bencana nonalam,
p07   F16  [ITEM]
p07   F16  | 38. Kesehatan Masyarakat Veteriner adalah segala
p07   F16  | urusan yang berhubungan dengan Hewan dan
p07   F16  | Produk Hewan yang secara langsung atau tidak
p07        | langsung memengaruhi kesehatan manusia.
p07   F15  [ITEM]
p07   F15  | 39. Obat Hewan adalah sediaan yang dapat digunakan
p07   F15  | untuk mengobati Hewan, membebaskan gejala, atau
p07   F18  | memodifikasi proses kimia dalam tubuh yang
p07   F14  | meliputi sediaan biologik, farmakoseutika, premiks,
p07   F14  | dan sediaan Obat Hewan alami.
p07   F15  [ITEM]
p07   F15  | 40. Alat dan Mesin Peternakan adalah semua peralatan
p07   F18  | yang digunakan berkaitan dengan kegiatan
p07   F14  | Peternakan, baik yang dioperasikan dengan motor
p07   F14  | penggerak maupun tanpa motor penggerak.
p07   F18  | 4 I . Alat dan Mesin Kesehatan Hewan adalah peralatan
p07   F15  | kedokteran Hewan yang disiapkan dan digunakan
p07   F18  | untuk Hewan sebagai alat bantu dalam
p07   F12  | penyelenggaraan Kesehatan Hewan.
p07   F16  [ITEM]
p07   F16  | 42. Kesejahteraan Hewan adalah segala urusan yang
p07   F18  | berhubungan dengan keadaan fisik dan mental
p07   F14  | Hewan menurut ukuran perilaku alami Hewan yang
p07   F16  | perlu diterapkan dan ditegakkan untuk melindungi
p07   F14  | Hewan dari periakuan Setiap Orang yang tidak layak
p07   F14  | terhadap Hewan yang dimanfaatkan manusia.
p07   F16  [ITEM]
p07   F16  | 43. Tenaga Kesehatan Hewan adalah orang yang
p07   F17  | menjalankan aktivitas di bidang Kesehatan Hewan
p07   F16  | berdasarkan kompetensi dan kewenangan Medik
p07   F14  | Veteriner yang hierarkis sesuai dengan pendidikan
p07   F18  | formal dan/ atau pelatihan Kesehatan Hewan
p07   F12  | bersertil-rkat.
p07   F16  [ITEM]
p07   F16  | 44. Dihapus.
p07   F18  [ITEM]
p07   F18  | 45. Pemerintah Fusat yang selanjutnya disebut
p07   F14  | Pemerintah adalah Presiden Republik Indonesia yang
p07   F16  | rnemegang kekuasaan pemerintahan Negara
p07   F14  | Kesatuan Republik Indonesia sebagaimana dimaksud
p07   F16  | dalam Undang-Undang Dasar Negara Republik
p07   F14  | Indonesia Tahun 1945.
p07   F16  [ITEM]
p07   F16  | 46. Menteri ...
==================== PAGE 8 ====================
p08   F12  [ITEM]
p08   F12  | 2.
p08   F12  [ITEM]
p08   F12  | 3.
p08   F11  | PRESIDEN
p08   F12  | R EPUF:LIK INDONESIA
p08   F18  | -8-
p08   F16  [ITEM]
p08   F16  | 46. Menteri adalah menteri yang menyelenggarakan
p08   F18  | urusan pemerintahan di bidang Peternakan dan
p08   F12  | Kesehatan Hewan.
p08   F16  [ITEM]
p08   F16  | 47. Pemerintah Daerah adalah gubernur, bupati/ wali
p08   F18  | kota, dan perangkat daerah sebagai unsur
p08        | penyelenggara Pemerintahan Daerah.
p08   F17  [ITEM]
p08   F17  | 48. Pemerintahan Daerah adalah penyelenggaraan
p08   F14  | urusan pemerintahan oleh Pemerintah Daerah dan
p08   F16  | dewan perwakilan rakyat daerah menurut asas
p08   F16  | otonomi dan tugas pembantuan dengan prinsip
p08   F16  | otonomi seluas-luasnya dalam sistem dan prinsip
p08   F14  | Negara Kesatuan Republik Indonesia sebagaimana
p08   F16  | dimaksud daiam Undang-Undang Dasar Negara
p08   F14  | Republik Indonesia Tahun 1945.
p08   F16  [ITEM]
p08   F16  | 49. Sistem Kesehatan Hewan Nasional yang selanjutnya
p08   F16  | disebut Siskeswanas adalah tatanan Kesehatan
p08   F18  | Hewan yang ditetapkan oleh Pemerintah dan
p08   F17  | diselenggarakan oleh Otoritas Veteriner dengan
p08   F18  | melibatkan seluruh penyelenggara Kesehatan
p08   F16  | Hewan, pemangku kepentingan, dan masyarakat
p08        | secara terpadu.
p08   F16  | Ketentuan Pasal 6 ayat (21 huruf b, substansi tetap dan
p08   F16  | penjelasannya tentang uinseminasi buatan" dihapus
p08   F18  | sehingga rumusan penjelasan Pasal 6
p08        I4 | adalah
p08   F14  | sebagaimana tercantum dalam Penjelasan Pasal demi
p08   F14  | Pasal Angka 2 Undang-undang ini.
p08   F18  | Judul Bagian Kesatu pada Bab IV diubah sehingga
p08   F14  | berbunyi sebagai berikut:
p08        [HEADING:BAGIAN]
p08        | Bagian Kesatu
p08   F14  | Benih dan Bibit
p08   F14  | Ketentuan Pasal 13 diubah sehingga berbunyi sebagai
p08   F14  | berikut:
p08   F14  | Pasal 13...
==================== PAGE 9 ====================
p09   F10  | FTRESIDEN
p09   F14  | R Ei:IUE I- IK IND ONES IA
p09   F18  | -9-
p09   F12  | Pasal 13
p09   F17  | Penyediaan dan pengembangan Benih dan/atau
p09   F17  | Bibit dilakukan dengan mengutamakan produksi
p09        | dalam negeri.
p09   F16  | Pemerintah dan/atau Pemerintah Daerah sesuai
p09   F18  | dengan kewenangannya berkewajiban untuk
p09   F16  | melakukan Pemuliaan, pengembangan usaha
p09   F18  | pembenihan dan/atau pembibitan dengan
p09   F15  | melibatkan peran serta masyarakat untuk menjamin
p09   F14  | ketersediaan Benih dan/atau Bibit.
p09   F14  | Kewajiban Pemerintah dan/atau Pemerintah Daerah
p09   F17  | sesuai dengan kewenangannya untuk melakukan
p09   F18  | pengembangan usaha pembenihan dan/atau
p09   F16  | pembibitan sebagaimana dimaksud pada ayat (2)
p09   F15  | dilakukan dengan mendorong penerapan teknologi
p09        | reproduksi.
p09   F14  | Dalam hal usaha pembenihan dan/atau pembibitan
p09   F16  | oleh masyarakat belum berkembang, Pemerintah
p09   F18  | dan/atau Pemerintah Daerah sesuai dengan
p09   F18  | kewenangannya membentuk unit pembenihan
p09   F14  | dan / atau pembibitan
p09   F18  | Pembentukan unit pembenihan sebagaimana
p09   F15  | dimaksud pada ayat (4) ditujukan untuk pemurnian
p09   F14  | Ternak tertentu atau untuk produksi.
p09   F14  | Setiap Benih atau Bibit yang beredar wajib memiliki
p09   F14  | sertifikat Benih atau Bibit yang memuat keterangan
p09   F14  | mengenai silsilah dan ciri-ciri keunggulannya.
p09   F16  | Sertihkat Benih atau Bibit sebagaimana dimaksud
p09   F15  | pada ayat (6) dikeluarkan oleh lembaga sertifikasi
p09   F17  | Benih atau Bibit yang terakreditasi atau yang
p09   F14  | ditunjuk oleh Menteri.
p09   F14  | Setiap Orang dilarang mengedarkan Benih atau Bibit
p09   F14  | yang tidak memiliki sertifikat sebagaimana dimaksud
p09        | pada ayat (6).
p09   F16  [ITEM]
p09   F16  | 5. Ketentuan Pasal 15 diubah sehingga berbunyi sebagai
p09   F14  | berikut:
p09   F10  [AYAT]
p09   F10  | (1)
p09   F10  [AYAT]
p09   F10  | (2)
p09   F10  [AYAT]
p09   F10  | (3)
p09   F10  [AYAT]
p09   F10  | (4)
p09   F11  | (s)
p09   F10  [AYAT]
p09   F10  | (6)
p09   F10  [AYAT]
p09   F10  | (7)
p09   F10  [AYAT]
p09   F10  | (8)
p09   F14  | Pasal 15...
==================== PAGE 10 ====================
p10   F11  | PRESIDEN
p10   F12  | R EPI'FILIK INOONESIA
p10   F17  | -10-
p10   F12  | Pasal 15
p10   F16  [AYAT]
p10   F16  | (1) Pemasukan Benih dan/atau Bibit dari luar negeri ke
p10   F14  | dalam wilayah Negara Kesatuan Republik Indonesia
p10   F14  | dapat dilakukan untuk:
p10   F18  [SUB-ITEM]
p10   F18  | a. meningkatkan mutu dan keragaman genetik;
p10   F18  [SUB-ITEM]
p10   F18  | b. mengembangkan ilmu pengetahuan dan
p10        | teknologi;
p10   F18  [SUB-ITEM]
p10   F18  | c. mengatasi kekurangan Benih
p10   F14  | dalam negeri; dan/atau
p10   F18  [SUB-ITEM]
p10   F18  | d. memenuhi keperluan
p10   F15  | dan/atau Bibit di
p10   F18  | penelitian dan
p10   F12  | pengembangan.
p10   F16  | Pemasukan Benih dan/atau Bibit dari luar negeri
p10   F14  | sebagaimana dimaksud pada ayat (1) harus:
p10   F18  [SUB-ITEM]
p10   F18  | a. memenuhi persyaratan mutu;
p10   F18  [SUB-ITEM]
p10   F18  | b. memenuhi persyaratan teknis Kesehatan Hewan;
p10   F18  [SUB-ITEM]
p10   F18  | c. bebas dari Penyakit Hewan Menular yang
p10   F14  | dipersyaratkan oleh otoritas veteriner;
p10   F18  [SUB-ITEM]
p10   F18  | d. memenuhi ketentuan peraturan perundang-
p10   F14  | undangan di bidang karantina Hewan; dan
p10   F18  [SUB-ITEM]
p10   F18  | e. memerhatikan kebijakan pewilayahan sumber
p10   F14  | Bibit sebagaimana dimaksud dalam Pasal 14.
p10   F15  | Setiap Orang yang melakukan pemasukan Benih
p10   F15  | dan/atau Bibit sebagaimana dimaksud pada ayat (l)
p10   F14  | wajib memperoleh izin dari Menteri.
p10   F18  | Ketentuan lebih lanjut mengenai persyaratan
p10   F16  | mutu dan persyaratan teknis Kesehatan Hewan
p10   F16  | sebagaimana dimaksud pada ayat (2) huruf a dan
p10   F14  | huruf b diatur dengan Peraturan Menteri.
p10   F18  [ITEM]
p10   F18  | 6. Ketentuan Pasal 16 diubah sehingga berbunyi sebagai
p10   F14  | berikut:
p10   F12  | Pasal 16
p10   F18  [AYAT]
p10   F18  | ( 1) Pengeluaran Benih dan/ atau Bibit dari wilayah
p10   F15  | Negara Kesatuan Republik Indonesia ke luar negeri
p10   F15  | dapat dilakukan apabila kebutuhan dalam negeri
p10   F18  | telah terpenuhi dan kelestarian Ternak 1oka1
p10        | terjamin.
p10   F8   [AYAT]
p10   F8   | (21
p10   F10  [AYAT]
p10   F10  | (3)
p10   F8   [AYAT]
p10   F8   | (41
p10   F14  [AYAT]
p10   F14  | (2) Pengeluaran ...
==================== PAGE 11 ====================
p11   F11  | PRESIDEN
p11   F12  | REPIJRL"IK INDONESIA
p11   F16  | - 11-
p11   F16  [AYAT]
p11   F16  | (2) Pengeluaran sebagaimana dimaksud pada ayat (1)
p11   F15  | dilarang dilakukan terhadap Benih dan/ atau Bibit
p11   F14  | yang terbaik di dalam negeri.
p11   F16  [AYAT]
p11   F16  | (3) Setiap Orang yang melakukan kegiatan sebagaimana
p11   F14  | dimaksud pada ayat (1) wajib memperoleh izin dari
p11        | Menteri.
p11   F18  [ITEM]
p11   F18  | 7. Ketentuan Pasal 18 diubah sehingga berbunyi sebagai
p11   F14  | berikut:
p11   F12  | Pasal 18
p11   F14  | Dalam rangka mencukupi ketersediaan bibit, Ternak
p11   F17  | Ruminansia Betina Produktif diseleksi untuk
p11   F15  | Pemuliaan, sedangkan Ternak ruminansia betina
p11   F16  | yang tidak produktif disingkirkan untuk dijadikan
p11        | Ternak potong.
p11   F16  | Penentuan Ternak ruminansia betina yang tidak
p11   F16  | produktif sebagaimana dimaksud pada ayat (1)
p11   F14  | dilakukan oleh Dokter Hewan Berwenang
p11   F14  | Pemerintah Daerah sesuai dengan kewenangannya
p11   F18  | menyediakan dana untuk menjaring Ternak
p11   F14  | Ruminansia Betina Produktif yang dikeluarkan oleh
p11   F14  | masyarakat dan menampung Ternak tersebut pada
p11   F18  | unit pelaksana teknis di daerah untuk keperluan
p11   F17  | pengembangbiakan dan penyediaan Bibit Ternak
p11   F14  | ruminansia betina di daerah tersebut.
p11   F18  | Setiap Orang dilarang menyembelih Ternak
p11   F17  | ruminansia kecil betina produktif atau Ternak
p11   F14  | ruminansia besar betina produktif.
p11   F16  | Larangan sebagaimana dimaksud pada ayat (4)
p11   F14  | dikecualikan dalam hal:
p11   F16  [SUB-ITEM]
p11   F16  | a. penelitian;
p11   F16  [SUB-ITEM]
p11   F16  | b. Pemuliaan;
p11   F18  [SUB-ITEM]
p11   F18  | c. pengendalian dan penanggulangan Penyakit
p11   F12  | Hewan;
p11   F16  [SUB-ITEM]
p11   F16  | d. ketentuan agama;
p11   F16  [SUB-ITEM]
p11   F16  | e. ketentuan adat istiadat; dan/atau
p11   F18  [SUB-ITEM]
p11   F18  | f. pengakhiran penderitaan Hewan.
p11   F10  [AYAT]
p11   F10  | (1)
p11   F10  [AYAT]
p11   F10  | (2)
p11   F11  | (s)
p11   F10  [AYAT]
p11   F10  | (4)
p11   F11  | (s)
p11   F14  [AYAT]
p11   F14  | (6) Setiap ...
==================== PAGE 12 ====================
p12   F11  | PRESIDEN
p12   F14  | R EPIJEJL IK INDONESIA
p12   F20  | -t2-
p12   F14  | Setiap Orang harus menjaga populasi anakan ternak
p12   F16  | ruminansia kecil dan anakan ternak ruminansia
p12   F12  | besar.
p12   F16  | Ketentuan lebih lanjut mengenai penyeleksian dan
p12   F14  | penyingkiran sebagaimana dimaksud pada ayat ( 1),
p12   F15  | penjaringan Ternak Ruminansia Betina Produktif
p12   F14  | sebagaimana dimaksud pada ayat (3), dan populasi
p12   F14  | anakan ternak ruminansia kecil dan anakan ternak
p12   F14  | ruminansia besar sebagaimana dimaksud pada ayat
p12   F14  [AYAT]
p12   F14  | (6) diatur dengan Peraturan Menteri.
p12   F18  [ITEM]
p12   F18  | 8. Ketentuan Pasal 31 diubah sehingga berbunyi sebagai
p12   F14  | berikut:
p12   F12  | Pasal 3 1
p12   F18  [AYAT]
p12   F18  | ( 1) Peternak dapat melakukan kemitraan usaha di
p12   F16  | bidang budi daya Ternak berdasarkan perjanjian
p12   F18  | yang
p12   F18  | saling memerlukan, memperkuat,
p12   F14  | menguntungkan, menghargai, bertanggung jawab,
p12        | ketergantungan, dan berkeadilan.
p12   F14  [AYAT]
p12   F14  | (21 Kemitraan usaha sebagaimana dimaksud pada ayat
p12   F14  [AYAT]
p12   F14  | (1) dapat dilakukan:
p12   F18  [SUB-ITEM]
p12   F18  | a. antar-Peternak;
p12   F18  [SUB-ITEM]
p12   F18  | b. antara Peternak dan Perusahaan Peternakan;
p12   F18  [SUB-ITEM]
p12   F18  | c. antara Peternak dan perusahaan di bidang lain;
p12   F12  | dan
p12   F18  [SUB-ITEM]
p12   F18  | d. antara Perusahaan Peternakan dan Pemerintah
p12   F18  | atau Pemerintah Daerah sesuai dengan
p12   F12  | kewenangannya.
p12   F17  [AYAT]
p12   F17  | (3) Kemitraan usaha sebagaimana dimaksud pada ayat
p12        [AYAT]
p12        | (2) dapat berupa:
p12   F18  [SUB-ITEM]
p12   F18  | a. penyediaan sarana produksi;
p12   F18  [SUB-ITEM]
p12   F18  | b. produksi;
p12   F18  [SUB-ITEM]
p12   F18  | c. pemasaran; dan/atau
p12   F18  [SUB-ITEM]
p12   F18  | d. permodalan atau pembiayaan.
p12   F10  [AYAT]
p12   F10  | (6)
p12   F9   [AYAT]
p12   F9   | (71
p12   F14  [AYAT]
p12   F14  | (4) Pemerintah ...
==================== PAGE 13 ====================
p13   F11  | PRESIDEN
p13   F12  | R EPUBL.IK INDONESIA
p13   F14  | _13_
p13   F17  [AYAT]
p13   F17  | (4) Pemerintah dan Pemerintah Daerah sesuai dengan
p13   F14  | kewenangannya melakukan pembinaan kemitraan
p13   F14  | usaha sebagaimana dimaksud pada ayat (21 dengan
p13   F16  | memerhatikan ketentuan peraturan perundang-
p13   F14  | undangan di bidang kemitraan usaha.
p13   F18  [ITEM]
p13   F18  | 9. Ketentuan Pasal 32 diubah sehingga berbunyi sebagai
p13   F14  | berikut:
p13   F12  | Pasal 32
p13   F18  [AYAT]
p13   F18  | (1) Pemerintah dan Pemerintah Daerah sesuai
p13   F14  | dengan kewenangannya berkewajiban mendorong
p13   F18  | agar sebanyak mungkin warga masyarakat
p13   F14  | menyelenggarakan budi daya Ternak sesuai dengan
p13   F14  | pedoman budi daya Ternak yang baik.
p13   F16  [AYAT]
p13   F16  | (2) Pemerintah dan Pemerintah Daerah sesuai dengan
p13   F18  | kewenangannya memfasilitasi dan membina
p13   F18  | pengembangan budi daya yang dilakukan oleh
p13   F18  | Peternak dan pihak tertentu yang mempunyai
p13        | kepentingan khusus.
p13   F16  [AYAT]
p13   F16  | (3) Pemerintah dan Pemerintah Daerah sesuai dengan
p13   F14  | kewenangannya membina dan memberikan fasilitas
p13   F16  | untuk pertumbuhan dan perkembangan koperasi
p13   F14  | dan badan usaha di bidang Peternakan.
p13   F15  | Ketentuan Pasal 36 diubah sehingga berbunyi sebagai
p13   F14  | berikut:
p13   F12  | Pasal 36
p13   F16  | Pemerintah berkewajiban untuk menyelenggarakan
p13   F14  | dan memfasiiitasi kegiatan pemasaran Hewan atau
p13   F15  | Ternak dan Produk Hewan di dalam negeri maupun
p13   F14  | ke luar negeri.
p13   F16  | Pemasaran sebagaimana dimaksud pada ayat (l)
p13   F16  | diutamakan untuk membina peningkatan produksi
p13   F15  | dan konsumsi protein hewani dalam mewujudkan
p13   F18  | ketersediaan pangan bergizi seimbang bagi
p13   F18  | masyarakat dengan tetap
p13   F14  | meningkatkan
p13   F14  | kesejahteraan pelaku usaha Peternakan.
p13   F11  [ITEM]
p13   F11  | 10.
p13   F10  [AYAT]
p13   F10  | (1)
p13   F11  [AYAT]
p13   F11  | (2t
p13   F14  [AYAT]
p13   F14  | (3) Pemerintah ...
==================== PAGE 14 ====================
p14   F11  | PRESIOEN
p14   F12  | R EFtUi3t.IK INDONESIA
p14   F17  | -74-
p14   F16  [AYAT]
p14   F16  | (3) Pemerintah dan Pemerintah Daerah sesuai dengan
p14   F16  | kewenangannya berkewajiban untuk menciptakan
p14   F14  | iklim usaha yang sehat bagi pemasaran Hewan atau
p14   F14  | Ternak dan Produk Hewan.
p14   F14  [ITEM]
p14   F14  | 11. Di antara Pasal 36 dan Pasal 37 disisipkan 5 (lima) pasal,
p14   F14  | yakni Pasal 36A, Pasal 36Et, Pasal 36C, Pasal 36D, dan
p14   F14  | Pasal 36E sehingga berbunyi sebagai berikut:
p14   F12  | Pasal 36A
p14   F14  | Pengeluaran Hewan atau Ternak dan Produk Hewan dari
p14   F16  | wilayah Negara Kesatuan Republik Indonesia ke luar
p14   F14  | negeri dapat dilakukan apabila produksi dan pasokan di
p14   F16  | dalam negeri telah mencukupi kebutuhan konsumsi
p14        | masyarakat.
p14   F14  | Pasal 368
p14   F16  [AYAT]
p14   F16  | (1) Pemasukan Ternak dan Produk Hewan dari luar
p14   F14  | negeri ke dalam wilayah Negara Kesatuan Republik
p14   F14  | Indonesia dilakukan apabila produksi dan pasokan
p14   F17  | Ternak dan Produk Hewan di dalam negeri belum
p14   F14  | mencukupi kebutuhan konsumsi masyarakat.
p14   F16  [AYAT]
p14   F16  | (2) Pemasukan Ternak sebagaimana dimaksud pada
p14   F14  | ayat (1) harus berupa Bakalan.
p14   F16  [AYAT]
p14   F16  | (3) Pemasukan Ternak ruminansia besar Bakalan tidak
p14   F14  | boleh melebihi berat tertentu.
p14   F14  | Setiap Orang yang melakukan pemasukan Bakalan
p14   F18  | sebagaimana dimaksud pada ayat (21 wajib
p14   F14  | memperoleh izin dari Menteri.
p14   F15  | Setiap Orang yang memasukkan Bakalan dari luar
p14   F15  | negeri sebagaimana dimaksud pada ayat (2) wajib
p14   F18  | melakukan penggemukan di da_lam negeri untuk
p14   F14  | memperoleh nilai tambah dalam jangka waktu paling
p14   F18  | cepat 4 (empat) bulan sejak dilakukan tindakan
p14   F14  | karantina berupa pelepasan.
p14   F15  | Pemasukan Ternak dari luar negeri sebagaimana
p14   F14  | dimaksud pada ayat (2) dan ayat (3) harus:
p14   F8   [AYAT]
p14   F8   | (41
p14   F11  | (s)
p14   F10  [AYAT]
p14   F10  | (6)
p14   F16  [SUB-ITEM]
p14   F16  | a. memenuhi ...
==================== PAGE 15 ====================
p15   F11  [AYAT]
p15   F11  | (7\
p15   F10  [AYAT]
p15   F10  | (8)
p15   F10  [AYAT]
p15   F10  | (1)
p15   F9   [AYAT]
p15   F9   | (21
p15   F10  [AYAT]
p15   F10  | (3)
p15   F11  | PRESIDEN
p15   F12  | REPUBLIK INDONESIA
p15   F14  | _15_
p15   F18  [SUB-ITEM]
p15   F18  | a. memenuhi persyaratan teknis Kesehatan Hewan;
p15   F18  [SUB-ITEM]
p15   F18  | b. bebas dari Penyakit Hewan Menular yang
p15        | dipersyaratkan oleh Otoritas Veteriner; dan
p15   F18  [SUB-ITEM]
p15   F18  | c. memenuhi ketentuan peraturan perundang-
p15   F14  | undangan di bidang karantina Hewan.
p15   F18  | Pemasukan Ternak dari luar negeri untuk
p15   F14  | dikembangbiakan di Indonesia harus:
p15   F18  [SUB-ITEM]
p15   F18  | a. memenuhi persyaratan teknis Kesehatan Hewan;
p15   F18  [SUB-ITEM]
p15   F18  | b. bebas dari Penyakit Hewan Menular yang
p15        | dipersyaratkan oleh Otoritas Veteriner; dan
p15   F18  [SUB-ITEM]
p15   F18  | c. memenuhi ketentuan peraturan perundang-
p15   F14  | undangan di bidang karantina Hewan.
p15   F14  | Ketentuan lebih lanjut mengenai pemasukan Ternak
p15   F14  | dan Produk Hewan sebagaimana dimaksud pada ayat
p15   F14  [AYAT]
p15   F14  | (1) serta berat tertentu sebagaimana dimaksud pada
p15   F14  | ayat (3) diatur dengan Peraturan Menteri.
p15   F12  | Pasal 36C
p15   F16  | Pemasukan Ternak Ruminansia Indukan ke dalam
p15   F14  | wilayah Negara Kesatuan Republik Indonesia dapat
p15   F16  | berasal dari suatu negara atau zona dalam suatu
p15   F14  | negara yang telah memenuhi persyaratan dan tata
p15        | cara pemasukannya.
p15   F18  | Persyaratan dan tata cara pemasukan Ternak
p15   F18  | Ruminansia Indukan dari luar negeri ke dalam
p15   F17  | wilayah Negara Kesatuan Republik Indonesia
p15   F18  | ditetapkan berdasarkan analisis risiko di bidang
p15   F16  | Kesehatan Hewan oleh Otoritas Veteriner dengan
p15        | mengutamakan kepentingan nasional.
p15   F16  | Pemasukan Ternak Ruminansia Indukan yang
p15   F14  | berasal dari zona sebagaimana dimaksud pada ayat
p15   F14  [AYAT]
p15   F14  | (1), selain harus memenuhi ketentuan sebagaimana
p15   F14  | dimaksud pada ayat (2) juga harus terlebih dahulu:
p15   F18  [SUB-ITEM]
p15   F18  | a. dinyatakan bebas Penyakit Hewan Menular di
p15   F14  | negara asal oleh otoritas veteriner negara asal
p15   F14  | sesuai dengan ketentuan yang ditetapkan badan
p15   F14  | kesehatan hewan dunia dan diakui oleh Otoritas
p15        | Veteriner Indonesia;
p15   F16  [SUB-ITEM]
p15   F16  | b. dilakukan...
==================== PAGE 16 ====================
p16   F11  | PRESIDEN
p16   F12  | REPUBLIK INOONESIA
p16   F14  | _16_
p16   F18  [SUB-ITEM]
p16   F18  | b. dilakukan penguatan sistem dan pelaksanaan
p16   F14  | surveilan di dalam negeri; dan
p16   F18  [SUB-ITEM]
p16   F18  | c. ditetapkan tempat pemasukan tertentu.
p16   F16  [AYAT]
p16   F16  | (4) Setiap Orang yang melakukan pemasukan Ternak
p16   F15  | Ruminansia Indukan sebagaimana dimaksud pada
p16   F14  | ayat (l) wajib memperoleh izin dari Menteri.
p16   F16  [AYAT]
p16   F16  | (5) Ketentuan lebih lanjut mengenai pemasukan Ternak
p16   F18  | Ruminansia Indukan ke dalam wilayah Negara
p16   F18  | Kesatuan Republik Indonesia diatur dengan
p16        | Peraturan Menteri.
p16   F12  | Pasal 36D
p16   F16  | Pemasukan Ternak Ruminansia Indukan yang
p16   F17  | berasal dari zona sebagaimana dimaksud dalam
p16   F18  | Pasal 36C harus ditempatkan di pulau karantina
p16   F16  | sebagai instalasi karantina Hewan pengamanan
p16   F14  | maksimal untuk jangka waktu tertentu.
p16   F15  | Ketentuan mengenai pulau karantina diatur dengan
p16        | Peraturan Pemerintah.
p16   F12  | Pasal 36E
p16   F18  | Dalam hal tertentu, dengan tetap memerhatikan
p16   F14  | kepentingan nasional, dapat dilakukan pemasukan
p16   F16  | Ternak dan/atau Produk Hewan dari suatu negara
p16   F14  | atau zona dalam suatu negara yang telah memenuhi
p16   F18  | persyaratan dan tata cara pemasukan Ternak
p16   F14  | dan/atau Produk Hewan.
p16   F15  | Ketentuan lebih lanjut mengenai dalam hal tertentu
p16   F14  | dan tata cara pemasukannya sebagaimana dimaksud
p16   F14  | pada ayat (1) diatur dengan Peraturan pemerintah.
p16   F12  | Pasal 37
p16   F18  [AYAT]
p16   F18  | (1) Pemerintah membina
p16   F18  | dan
p16        | memfasilitasi
p16   F15  | (l)
p16   F11  [AYAT]
p16   F11  | (2t
p16   F10  [AYAT]
p16   F10  | (1)
p16   F11  [AYAT]
p16   F11  | (2t
p16   F14  [ITEM]
p16   F14  | 12. Di antara ayat (2) dan ayat (3) pasal 37 disisipkan 1 (satu)
p16   F14  | ayat yakni ayat (2a), sehingga pasal 37 berbunyi sebagai
p16   F14  | berikut:
p16   F15  | berkembangnya industri pengolahan produk Hewan
p16   F14  | dengan mengutamakan penggunaan bahan baku dari
p16        | dalam negeri.
p16   F14  [AYAT]
p16   F14  | (2) Pemerintah ...
==================== PAGE 17 ====================
p17   F11  | FRESIDEN
p17   F12  | R EPUEJLIK INDONESIA
p17   F17  | -17-
p17   F16  [AYAT]
p17   F16  | (2) Pemerintah membina terselenggaranya kemitraan
p17   F14  | yang sehat antara industri pengolahan dan Peternak
p17   F14  | dan/atau koperasi yang menghasilkan Produk Hewan
p17   F14  | yang digunakan sebagai bahan baku industri.
p17   F16  [AYAT]
p17   F16  | (2a) Kemitraan sebagaimana dimaksud pada ayat (2)
p17   F14  | dapat berupa kerja sama:
p17   F18  [SUB-ITEM]
p17   F18  | a. Permodalan atau pembiayaan;
p17   F18  [SUB-ITEM]
p17   F18  | b. pengolahan;
p17   F18  [SUB-ITEM]
p17   F18  | c. pemasaran;
p17   F18  [SUB-ITEM]
p17   F18  | d. pendistribusian; dan/atau
p17   F18  [SUB-ITEM]
p17   F18  | e. rantai pasok.
p17   F16  [AYAT]
p17   F16  | (3) Ketentuan lebih lanjut mengenai pembinaan dan
p17   F14  | fasilitasi berkembangnya industri pengolahan produk
p17   F18  | Hewan sebagaimana dimaksud pada ayat (1)
p17   F16  | dilakukan sesuai dengan peraturan perundang-
p17   F16  | undangan di bidang industri, kecuali untuk hal-hal
p17   F14  | yang diatur dalam Undang-Undang ini.
p17   F15  [ITEM]
p17   F15  | 13. Ketentuan Pasal 41 diubah sehingga berbunyi sebagai
p17   F14  | berikut:
p17   F12  | Pasal 4 1
p17   F14  | Pencegahan Penyakit Hewan sebagaimana dimaksud
p17   F14  | dalam Pasal 39 bertujuan untuk:
p17   F18  [SUB-ITEM]
p17   F18  | a. melindungi wilayah Negara Kesatuan Republik
p17   F15  | Indonesia dari ancaman masuknya penyakit Hewan
p17   F14  | dari luar negeri;
p17   F18  [SUB-ITEM]
p17   F18  | b. melindungi wilayah Negara Kesatuan Republik
p17   F18  | Indonesia dari ancaman menyebarnya penyakit
p17   F14  | Hewan dari luar negeri, dari satu pulau ke pulau lain,
p17   F15  | dan antardaerah dalam satu pulau di dalam wilayah
p17        | Negara Kesatuan Republik Indonesia;
p17   F18  [SUB-ITEM]
p17   F18  | c. melindungi Hewan dari ancaman muncul, berjangkit,
p17        | dan menyebarnya penyakit Hewan; dan
p17   F18  [SUB-ITEM]
p17   F18  | d. mencegah keluarnya penyakit Hewan dari wilayah
p17        | Negara Kesatuan Republik Indonesia.
p17   F14  [ITEM]
p17   F14  | 14. Di antara . . .
==================== PAGE 18 ====================
p18   F11  | PRESIDEN
p18   F14  | R EP I.IR I- IK INDONESIA
p18   F17  | -18-
p18   F14  [ITEM]
p18   F14  | 14. Di antara Pasal 41 dan Pasal 42
p18   F18  | yakni Pasal 41A dan Pasal
p18   F14  | sebagai berikut:
p18   F10  [AYAT]
p18   F10  | (1)
p18   F9   [AYAT]
p18   F9   | (21
p18   F14  I3 | disisipkan 2 (dua) Pasal,
p18   F14  I3 | 41E} sehingga berbunyi
p18   F12  | Pasal 41A
p18   F15  | Pemerintah dan Pemerintah Daerah sesuai dengan
p18   F18  | kewenangannya bertanggung jawab melakukan
p18        | pencegahan Penyakit Hewan.
p18   F15  | Dalam melaksanakan tanggung jawab pencegahan
p18   F14  | Penyakit Hewan sebagaimana dimaksud pada ayat
p18   F18  [AYAT]
p18   F18  | (1), Pemerintah dan Pemerintah Daerah sesuai
p18   F15  | dengan kewenangannya berkewajiban melakukan
p18   F14  | koordinasi lintas sektoral, lintas wilayah, dan lintas
p18        | pemangku kepentingan.
p18   F16  | Koordinasi sebagaimana dimaksud pada ayat (2)
p18   F16  | dilakukan mulai tahap perencanaan, pelaksanaan,
p18   F16  | pemantauan, sampai dengan evaluasi kegiatan
p18        | pencegahan Penyakit Hewan.
p18   F14  | Dalam melaksanakan pencegahan Penyakit Hewan,
p18   F15  | Pemerintah dan Pemerintah Daerah sesuai dengan
p18   F18  | kewenangannya melakukan penyebarluasan
p18        | informasi dan peningkatan kesadaran masyarakat.
p18   F16  | Dalam pencegahan Penyakit Hewan, masyarakat
p18   F14  | dapat berperan aktif bersama dengan pemerintah dan
p18        | Pemerintah Daerah sesuai dengan kewenangannya.
p18   F12  | Pasal 4 1B
p18   F14  | Pencegahan Penyakit Hewan sebagaimana dimaksud
p18   F14  | dalam Pasal 41 meliputi:
p18   F18  [SUB-ITEM]
p18   F18  | a. pencegahan masuknya Penyakit Hewan dari luar
p18   F18  | negeri ke dalam wilayah Negara Kesatuan
p18        | Republik Indonesia;
p18   F18  [SUB-ITEM]
p18   F18  | b. pencegahan keluarnya penyakit Hewan dari
p18        | wilayah Negara Kesatuan Republik Indonesia;
p18   F18  [SUB-ITEM]
p18   F18  | c. pencegahan menyebarnya penyakit Hewan dari
p18   F14  | satu pulau ke pulau lain di dalam wilayah Negara
p18        | Kesatuan Republik Indonesia;
p18   F10  [AYAT]
p18   F10  | (3)
p18   F10  [AYAT]
p18   F10  | (4)
p18   F11  | (s)
p18   F10  [AYAT]
p18   F10  | (1)
p18   F16  [SUB-ITEM]
p18   F16  | d. pencegahan...
==================== PAGE 19 ====================
p19   F11  | PRESIDEN
p19   F12  | R EI-IUELiK INDONESIA
p19   F15  | -19_
p19   F18  [SUB-ITEM]
p19   F18  | d. pencegahan menyebarnya Penyakit Hewan dari
p19   F16  | satu wilayah ke wilayah lain dalam satu pulau;
p19   F12  | dan
p19   F18  [SUB-ITEM]
p19   F18  | e. pencegahan muncul, berjangkit, dan
p19   F18  | menyebarnya Penyakit Hewan di dalam suatu
p19        | wilayah.
p19   F18  | Pencegahan masuk, keluar, dan menyebarnya
p19   F14  | Penyakit Hewan sebagaimana dimaksud pada ayat (1)
p19   F15  | dilakukan dengan menerapkan persyaratan teknis
p19   F12  | Kesehatan Hewan.
p19   F14  | Pencegahan Penyakit Hewan sebagaimana dimaksud
p19   F14  | pada ayat (1) huruf a, huruf b, dan huruf c di tempat-
p19   F18  | tempat pemasukan dan pengeluaran dilakukan
p19   F14  | sesuai dengan ketentuan peraturan perundangan-
p19   F14  | undangan di bidang karantina hewan.
p19   F14  | Pencegahan Penyakit Hewan sebagaimana dimaksud
p19   F14  | pada ayat (1) huruf d dilakukan dengan pemeriksaan
p19        | dokumen dan Kesehatan Hewan.
p19   F16  | Pencegahan muncul, berjangkit, dan menyebarnya
p19   F14  | Penyakit Hewan di dalam suatu wilayah sebagaimana
p19   F16  | dimaksud pada ayat (1) huruf e dilakukan dengan
p19   F14  | cara tindakan pengebalan, pengoptimalan kebugaran
p19   F14  | hewan, dan/ atau biosekuriti.
p19   F16  [ITEM]
p19   F16  | 15. Ketentuan Pasal 58 diubah sehingga berbunyi sebagai
p19   F14  | berikut:
p19   F9   [AYAT]
p19   F9   | (21
p19   F10  [AYAT]
p19   F10  | (3)
p19   F10  [AYAT]
p19   F10  | (4)
p19   F11  | (s)
p19   F10  [AYAT]
p19   F10  | (1)
p19   F12  | Pasal 58
p19   F14  | Dalam rangka menjamin produk Hewan yang aman,
p19   F16  | sehat, utuh, dan ha1al bagi yang dipersyiratkan,
p19   F16  | Pemerintah dan Pemerintah Daerah "."ra1 d..rg..,
p19   F18  | kewenangannya berkewajiban melaksanakan
p19   F14  | pengawasan, pemeriksaan, pengujian, standardisasi,
p19   F14  | sertifikasi, dan registrasi produk Hewan.
p19   F16  | P-engawasan, pemeriksaan, dan pengujian produk
p19   F14  | Hewan berturut-turut dilakukan di tempat produksi,
p19   F18  | pada waktu pemotongan, penampungan, dan
p19   F16  | pengumpulan, pada waktu dalam keadaan segar,
p19   F16  | sebelum pengawetan, dan pada waktu peredaian
p19   F12  | setelah pengawetan.
p19   F10  [AYAT]
p19   F10  | (2)
p19        [AYAT]
p19        | (3) Standardisasi ...
==================== PAGE 20 ====================
p20   F11  | PRESIDEN
p20   F14  | R EP LIBL IK INDONESIA
p20   F17  | -20-
p20   F18  [AYAT]
p20   F18  | (3) Standardisasi, sertifikasi, dan registrasi Produk
p20   F17  | Hewan dilakukan terhadap Produk Hewan yang
p20   F18  | diproduksi di dan/atau dimasukkan ke dalam
p20   F14  | wilayah Negara Kesatuan Republik Indonesia untuk
p20   F14  | diedarkan dan/atau dikeluarkan dari wilayah Negara
p20        | Kesatuan Republik Indonesia.
p20   F18  [AYAT]
p20   F18  | (4) Produk Hewan yang diproduksi di dan/atau
p20   F15  | dimasukkan ke wilayah Negara Kesatuan Republik
p20   F14  | Indonesia untuk diedarkan wajib disertai:
p20   F18  [SUB-ITEM]
p20   F18  | a. sertifikat veteriner; dan
p20   F18  [SUB-ITEM]
p20   F18  | b. sertifikat halal bagi Produk Hewan yang
p20        | dipersyaratkan.
p20   F14  | Setiap Orang dilarang mengedarkan Produk Hewan
p20   F14  | yang diproduksi di dan/atau dimasukkan ke wilayah
p20   F16  | Negara Kesatuan Republik Indonesia yang tidak
p20   F16  | disertai dengan sertifikat sebagaimana dimaksud
p20        | pada ayat (4).
p20   F18  | Setiap Orang yang memproduksi dan/atau
p20   F14  | mengedarkan Produk Hewan dilarang memalsukan
p20   F16  | Produk Hewan dan/atau menggunakan bahan
p20        | tambahan yang dilarang.
p20   F14  | Produk Hewan yang dikeluarkan dari wilayah Negara
p20   F14  | Kesatuan Republik Indonesia wajib disertai sertifikat
p20   F15  | veteriner dan sertifikat halat jika dipersyaratkan oleh
p20   F12  | negara pengimpor.
p20   F11  | (s)
p20   F10  [AYAT]
p20   F10  | (6)
p20   F10  [AYAT]
p20   F10  | (7)
p20   F16  [AYAT]
p20   F16  | (8) Untuk pangan olahan asal Hewan, selain wajib
p20   F14  | memenuhi ketentuan sebagaimana dimaksud pada
p20   F18  | ayat (5) wajib memenuhi ketentuan peraturan
p20   F14  | perundang-undangan di bidang pangan.
p20   F16  [ITEM]
p20   F16  | 16. Ketentuan Pasal 59 diubah sehingga berbunyi sebagai
p20   F14  | berikut:
p20   F12  | Pasal 59
p20   F14  [AYAT]
p20   F14  | (1) Setiap Orang yang akan memasukkan produk Hewan
p20   F18  | ke dalam wilayah Negara Kesatuan Republik
p20   F16  | Indonesia wajib memperoleh izin pemasukan dari
p20   F18  | menteri
p20   F12  | yang
p20   F12  I2 | menyelenggarakan
p20   F18  | pemerintahan di bidang perdagangan
p20        | memperoleh rekomendasi dari:
p20   F14  I4 | urusan
p20   F12  I4 | setelah
p20   F16  [SUB-ITEM]
p20   F16  | a. Menteri ...
==================== PAGE 21 ====================
p21   F11  [AYAT]
p21   F11  | (2t
p21   F10  [AYAT]
p21   F10  | (3)
p21   F11  | PRESIDEN
p21   F14  | R EF,t]EL IK IN D ONES IA
p21   F15  | -21 -
p21   F18  [SUB-ITEM]
p21   F18  | a. Menteri untuk Produk Hewan segar; atau
p21   F18  [SUB-ITEM]
p21   F18  | b. pimpinan lembaga bidang pengawasan obat dan
p21   F17  | makanan untuk produk pangan olahan asal
p21   F12  | Hewan.
p21   F16  | Produk Hewan segar yang dimasukkan ke dalam
p21   F17  | wilayah Negara Kesatuan Republik Indonesia
p21   F15  | sebagaimana dimaksud pada ayat (1) huruf a harus
p21   F16  | berasal dari unit usaha Produk Hewan pada suatu
p21   F17  | negara yang telah memenuhi persyaratan dan
p21        | tatacara pemasukan Produk Hewan.
p21   F15  | Dalam hal produk pangan olahan asal Hewan yang
p21   F14  | akan dimasukkan ke dalam wilayah Negara Kesatuan
p21   F14  | Republik Indonesia sebagaimana dimaksud pada ayat
p21   F18  [AYAT]
p21   F18  | (1) huruf b yang mempunyai risiko penyebaran
p21   F18  | Zoonosis yang dapat mengancam kesehatan
p21   F14  | manusia, Hewan, dan lingkungan budi daya, sebelum
p21   F16  | diterbitkan rekomendasi oleh pimpinan lembaga
p21   F14  | pemerintah yang melaksanakan tugas pemerintahan
p21   F18  | di bidang pengawasan obat dan makanan harus
p21   F14  | mendapatkan persetujuan teknis dari Menteri.
p21   F14  | Persyaratan dan tata cara pemasukan Produk Hewan
p21   F15  | dari luar negeri ke dalam wilayah Negara Kesatuan
p21   F15  | Republik Indonesia sebagaimana dimaksud pada
p21   F16  | ayat (21 dan ayat (3) mengacu pada ketentuan
p21   F18  | yang berbasis analisis risiko di bidang Kesehatan
p21   F16  | Hewan dan Kesehatan Masyarakat Veteriner serta
p21        | mengutamakan kepentingan nasional.
p21   F10  [AYAT]
p21   F10  | (4)
p21   F16  [ITEM]
p21   F16  | 17. Ketentuan Pasal 65 di ubah sehingga berbunyi sebagai
p21   F14  | berikut:
p21   F12  | Pasal 65
p21   F15  | Ketentuan lebih ianjut mengenai Kesehatan Masyarakat
p21   F14  | Veteriner sebagaimana dimaksud dalam pasal 56 sampai
p21   F14  | dengan Pasal 64 diatur dengan Peraturan pemerintah.
p21   F14  [ITEM]
p21   F14  | 18. Di antara Pasal 66 dan Pasal 67 disisipkan 1 (satu) pasal
p21   F14  | yakni Pasal 66A sehingga berbunyi sebagai berikut:
p21        | Pasal 66A ...
==================== PAGE 22 ====================
p22   F11  | PRESIDEN
p22   F12  | REPI.,BI..IK INOONESIA
p22   F12  | _ .).) _
p22   F12  | Pasal 66A
p22   F17  | Setiap Orang dilarang menganiaya dan/ atau
p22   F14  | menyalahgunakan Hewan yang mengakibatkan cacat
p22   F14  | dan/atau tidak produktif.
p22   F15  | Setiap Orang yang mengetahui adanya perbuatan
p22   F18  | sebagaimana dimaksud pada ayat (l) wajib
p22        | melaporkan kepada pihak yang berwenang.
p22   F16  [ITEM]
p22   F16  | 19. Ketentuan Pasal 68 diubah sehingga berbunyi sebagai
p22   F14  | berikut:
p22   F12  | Pasal 68
p22   F15  | Pemerintah dan Pemerintah Daerah sesuai dengan
p22   F16  | kewenangannya menyelenggarakan Kesehatan
p22   F14  | Hewan di seluruh wilayah Negara Kesatuan Republik
p22   F12  | Indonesia.
p22   F18  | Dalam menyelenggarakan Kesehatan Hewan
p22   F15  | sebagaimana dimaksud pada ayat ( 1), pemerintah
p22   F18  | dan Pemerintah Daerah sesuai
p22   F12  I4 | dengan
p22   F16  | kewenangannya berkewajiban
p22        | meningkatkan
p22   F17  | penguatan tuga s, fungsi, dan wewenang Otoritas
p22        | Veteriner.
p22   F10  [AYAT]
p22   F10  | (1)
p22   F11  [AYAT]
p22   F11  | (2t
p22   F10  [AYAT]
p22   F10  | (1)
p22   F9   [AYAT]
p22   F9   | (21
p22   F14  [ITEM]
p22   F14  | 20. Di antara Pasal 68 dan pasal 69 disisipkan 5 (lima) pasal,
p22   F14  | yakni Pasal 68A, Pasal 688, pasal 6gC, pasal 6gD, dan
p22   F14  | Pasal 68E sehingga berbunyi sebagai berikut:
p22   F12  | Pasal 68A
p22   F15  | Otoritas Veteriner sebagaimana dimaksud dalam
p22   F18  | Pasai 68 ayat (21 mempunyai tugas menyiapkan
p22   F18  | rumusan dan melaksanakan kebijakan dalam
p22   F12  | penyelenggaraan Kesehatan Hewan.
p22   F14  | Otoritas Veteriner sebagaimana dimaksud pada ayat
p22   F14  [AYAT]
p22   F14  | (1) dipimpin oleh pejabat Otoritas Veteriner.
p22   F14  | Pejabat Otoritas Veteriner sebagaimana dimaksud
p22   F14  | pada ayat (2) terdiri atas;
p22   F10  [AYAT]
p22   F10  | (1)
p22   F8   [AYAT]
p22   F8   | (21
p22   F11  | (s)
p22   F16  [SUB-ITEM]
p22   F16  | a. pejabat...
==================== PAGE 23 ====================
p23   F10  [AYAT]
p23   F10  | (1)
p23   F10  [AYAT]
p23   F10  | (2)
p23   F10  [AYAT]
p23   F10  | (3)
p23   F9   [AYAT]
p23   F9   | (41
p23   F11  | (s)
p23   F10  | PRESIOE N
p23   F12  | REPUBLIK INOONESIA
p23   F17  | -23-
p23   F18  [SUB-ITEM]
p23   F18  | a. pejabat Otoritas Veteriner nasional;
p23   F18  [SUB-ITEM]
p23   F18  | b. pejabat Otoritas Veteriner kementerian;
p23   F18  [SUB-ITEM]
p23   F18  | c. pejabat Otoritas Veteriner provinsi; dan
p23   F18  [SUB-ITEM]
p23   F18  | d. pejabat Otoritas Veteriner kabupaten/kota.
p23   F12  | Pasal 68E}
p23   F18  | Pejabat Otoritas Veteriner di tingkat nasional
p23   F16  | sebagaimana dimaksud dalam Pasal 684 ayat (3)
p23   F14  | huruf a diangkat oleh Menteri.
p23   F18  | Pejabat Otoritas Veteriner di tingkat kementerian
p23   F15  | sebagaimana dimaksud dalam Pasal 68A ayat (3)
p23   F14  | huruf b diangkat oleh menteri.
p23   F18  | Pejabat Otoritas Veteriner di tingkat provinsi
p23   F15  | sebagaimana dimaksud dalam Pasal 68A ayat (3)
p23   F14  | huruf c diangkat oleh gubernur.
p23   F15  | Pejabat Otoritas Veteriner di tingkat kabupaten/kota
p23   F16  | sebagaimana dimaksud dalam Pasal 68A ayat (3)
p23   F14  | huruf d diangkat oleh bupati/wali kota.
p23   F14  | Pejabat Otoritas Veteriner sebagaimana dimaksud
p23   F14  | pada ayat (1), ayat (21, ayat (3), dan ayat (4) diangkat
p23   F14  | berdasarkan kompetensinya sebagai Dokter Hewan
p23   F12  | Berwenang.
p23   F12  | Pasal 68C
p23   F16  [AYAT]
p23   F16  | ( 1) Otoritas Veteriner sebagaimana dimaksud dalam
p23        | Pasal 68 mempunyai fungsi:
p23   F18  [SUB-ITEM]
p23   F18  | a. pelaksana Kesehatan Masyarakat Veteriner;
p23   F18  [SUB-ITEM]
p23   F18  | b. penyusun standar dan meningkatkan mutu
p23   F12  | penyelenggaraan Kesehatan Hewan;
p23   F18  [SUB-ITEM]
p23   F18  | c. pengidentifikasi masalah dan
p23        | pelaksana
p23        | pelayanan Kesehatan Hewan;
p23   F18  [SUB-ITEM]
p23   F18  | d. pelaksana pengendalian dan penanggulangan
p23        | Penyakit Hewan;
p23   F18  [SUB-ITEM]
p23   F18  | e. pengawas dan pengendali pemotongan Ternak
p23   F15  | Ruminansia Betina produktif dan/atau Ternak
p23   F14  | Ruminansia Indukan;
p23   F18  [SUB-ITEM]
p23   F18  | f. pengawas ...
==================== PAGE 24 ====================
p24   F40  | q,,D
p24   F11  | PRESIDEN
p24   F14  | R EPUEL IK INDONESIA
p24   F17  | - 24'
p24   F18  [SUB-ITEM]
p24   F18  | f. pengawas tindakan penganiayaan dan
p24   F16  | penyalahgunaan terhadap Hewan serta aspek
p24   F14  | Kesejahteraa n Hewan lainnya;
p24   F18  [SUB-ITEM]
p24   F18  | g. pengelola Tenaga Kesehatan Hewan;
p24   F18  [SUB-ITEM]
p24   F18  | h. pelaksana pengembangan profesi kedokteran
p24   F12  | Hewan;
p24   F14  | pengawas penggunaan Alat dan Mesin Kesehatan
p24   F12  | Hewan;
p24   F18  | pelaksana perlindungan Hewan
p24   F12  I4 | dan
p24        | lingkungannya;
p24   F18  [SUB-ITEM]
p24   F18  | k. pelaksana penyidikan dan pengamatan Penyakit
p24   F12  | Hewan;
p24   F18  | L penjamin ketersediaan dan mutu Obat Hewan;
p24   F16  [SUB-ITEM]
p24   F16  | m. penjamin keamanan Pakan dan bahan Pakan
p24   F12  | asal Hewan;
p24   F18  [SUB-ITEM]
p24   F18  | n. peny'usun prasarana dan sarana serta
p24   F16  | pembiayaan Kesehatan Hewan dan Kesehatan
p24        | Masyarakat Veteriner; dan
p24   F18  [SUB-ITEM]
p24   F18  | o. pengelola medik akuatik dan Medik Konservasi.
p24   F14  | Otoritas Veteriner sebagaimana dimaksud pada ayat
p24   F14  [AYAT]
p24   F14  | ( 1) berwenang mengambil keputusan tertinggi yang
p24   F14  | bersifat teknis Kesehatan Hewan.
p24   F16  | Pengambilan keputusan sebagaimana dimaksud
p24   F18  | pada ayat (21 dilakukan dengan melibatkan
p24   F18  | keprofesionalan Dokter Hewan dan dengan
p24   F14  | mengerahkan semua Iini kemampuan profesi.
p24   F18  | Keterlibatan keprofesionalan Dokter Hewan
p24   F16  | sebagaimana dimaksud pada ayat (3) dilakukan
p24   F18  | mulai dari identifikasi masalah, rekomendasi
p24   F16  | kebijakan, koordinasi pelaksanaan kebijakan,
p24   F16  | sampai dengan pengendalian teknis operasional
p24   F14  | penyelenggaraan Kesehatan Hewan di lapangan.
p24   F12  | Pasal 68D
p24   F18  [AYAT]
p24   F18  | (1) Dalam penyelenggaraan Kesehatan Hewan
p24   F16  | sebagaimana dimaksud dalam pasal 6g ayat (1),
p24        | Pemerintah menetapkan Siskeswanas.
p24   F15  [SUB-ITEM]
p24   F15  | j.
p24   F9   [AYAT]
p24   F9   | (21
p24   F10  [AYAT]
p24   F10  | (3)
p24   F10  [AYAT]
p24   F10  | (4)
p24   F14  [AYAT]
p24   F14  | (2) Siskeswanas ...
==================== PAGE 25 ====================
p25   F10  [AYAT]
p25   F10  | (2)
p25   F10  [AYAT]
p25   F10  | (3)
p25   F11  | PRESIDEN
p25   F12  | R EP UBI,.IK INDONESIA
p25   F17  | -25-
p25   F14  | Siskeswanas sebagaimana dimaksud pada ayat (1)
p25   F18  | menjadi acuan Otoritas Veteriner dalam
p25   F12  | penyelenggaraan Kesehatan Hewan.
p25   F17  | Dalam pelaksanaan Siskeswanas sebagaimana
p25   F14  | dimaksud pada ayat (2), Pemerintah dan Pemerintah
p25        | Daerah sesuai dengan kewenangannya:
p25   F18  [SUB-ITEM]
p25   F18  | a. meningkatkan peran dan fungsi kelembagaan
p25        | penyelenggaraan Kesehatan Hewan; dan
p25   F16  [SUB-ITEM]
p25   F16  | b. melaksanakan
p25   F14  I3 | koordinasi
p25   F12  I4 | dengan
p25   F14  | memperhatikan ketentuan peraturan perundang-
p25   F14  | undangan di bidang Pemerintahan Daerah.
p25   F18  | Peningkatan peran dan fungsi kelembagaan
p25   F16  | penyelenggaraan Kesehatan Hewan sebagaimana
p25   F18  | dimaksud pada ayat (3) huruf a dilaksanakan
p25        | melalui:
p25   F18  [SUB-ITEM]
p25   F18  | a. upaya Kesehatan Hewan meliputi pembentukan
p25   F18  | unit respons cepat di pusat dan daerah serta
p25   F14  | penguatan dan pengembangan pusat kesehatan
p25   F12  | hewan;
p25   F18  [SUB-ITEM]
p25   F18  | b. penelitian dan pengembangan Kesehatan Hewan;
p25   F18  [SUB-ITEM]
p25   F18  | c. sumber daya Kesehatan Hewan;
p25   F18  [SUB-ITEM]
p25   F18  | d. informasi Kesehatan Hewan yang terintegrasi;
p25   F12  | dan
p25   F18  [SUB-ITEM]
p25   F18  | e. peran serta masyarakat.
p25   F16  | Dalam ikut berperan serta mewujudkan Kesehatan
p25   F18  | Hewan dunia melalui Siskeswanas, Menteri
p25   F16  | melimpahkan kewenangannya kepada Otoritas
p25        | Veteriner.
p25   F16  | Otoritas Veteriner bersama organisasi profesi
p25   F16  | kedokteran Hewan melaksanakan Siskeswanas
p25   F14  | dengan memberdayakan potensi Tenaga Kesehatan
p25   F18  | Hewan dan membina pelaksanaan praktik
p25   F18  | kedokteran Hewan di seluruh wiiayah Negara
p25        | Kesatuan Republik Indonesia.
p25   F10  [AYAT]
p25   F10  | (4)
p25   F11  | (s)
p25   F10  [AYAT]
p25   F10  | (6)
p25        | Pasal 68E ...
==================== PAGE 26 ====================
p26   F11  | PRESIDEN
p26   F14  | R EPUE I- IK INDONESIA
p26   F14  | _26_
p26   F12  | Pasal 68E
p26   F14  | Ketentuan lebih lanjut mengenai Otoritas Veteriner dan
p26   F14  | Siskeswanas sebagaimana dimaksud dalam Pasal 68,
p26   F14  | Pasal 68A, Pasal 68E}, Pasal 68C, dan Pasal 68D diatur
p26        | dengan Peraturan Pemerintah.
p26   F14  [ITEM]
p26   F14  | 21. Ketentuan ayat (1) Pasal 85 diubah dan ayat (4) dan ayat
p26   F14  [AYAT]
p26   F14  | (5) dihapus, sehingga Pasal 85 berbunyi sebagai berikut:
p26   F12  | pasal 85
p26   F18  | Setiap Orang yang melanggar ketentuan
p26   F14  | sebagaimana dimaksud dalam Pasal 9 ayat (1), pasal
p26   F14  | 11 ayat (1), Pasal 13 ayat (8), Pasal 15 ayat (3), Pasal
p26   F14  | 16 ayat (2), Pasal 16 ayat (3), Pasal 18 ayat (4), Pasal
p26   F14  | 19 ayat (1), Pasal 22 ayat (1), Pasal 24 ayat (3), Pasal
p26   F14  | 25 ayat (1), Pasal 29 ayat (3), Pasal 29 ayat (4), Pasal
p26   F14  | 368 ayat (4), Pasal 36E} ayat (5), Pasal 36C ayat (4),
p26   F14  | Pasal 42 ayat (5), Pasal 43 ayat (4), Pasal 45 ayat (1),
p26   F14  | Pasal 47 ayat (2), Pasal 47 ayat (3), Pasal 50 ayat (1),
p26   F14  | Pasal 50 ayat (3), Pasal 51 ayat (2), Pasal 52 ayat (1),
p26   F14  | Pasal 54 ayat (3), Pasal 55 ayat (3), Pasal 58 ayat (5),
p26   F14  | Pasal 59 ayat (1), Pasal 60 ayat (1), Pasal 61 ayat (1),
p26   F14  | Pasal 61 ayat (21, Pasal 62 ayat (21, Pasai 62 ayat (3),
p26   F16  | Pasal 69 ayat l2l, Pasal 72 ayat (1), atau Pasal 80
p26   F14  | ayat (1) dikenai sanksi administratif.
p26   F14  | Sanksi administratif sebagaimana dimaksud pada
p26        | ayat (1) dapat berupa:
p26   F18  [SUB-ITEM]
p26   F18  | a. peringatan secara tertulis;
p26   F18  [SUB-ITEM]
p26   F18  | b. pengenaan denda;
p26   F18  [SUB-ITEM]
p26   F18  | c. penghentian sementara dari kegiatan, produksi,
p26   F14  | dan/atau peredaran;
p26   F18  [SUB-ITEM]
p26   F18  | d. pencabutan nomor pendaftaran dan penarikan
p26   F14  | Obat Hewan, Pakan, alat dan mesin, atau produk
p26   F14  | Hewan dari peredaran; atau
p26   F18  [SUB-ITEM]
p26   F18  | e. pencabutan izin.
p26   F18  | Ketentuan lebih lanjut mengenai tata cara
p26   F18  | pengenaan sanksi administratif se.bagaimana
p26   F16  | dimaksud pada ayat (2) diatur dalam peraturan
p26   F12  | Pemerintah.
p26   F10  [AYAT]
p26   F10  | (1)
p26   F11  [AYAT]
p26   F11  | (2t
p26   F10  [AYAT]
p26   F10  | (3)
p26   F16  [ITEM]
p26   F16  | 22. Ketentuan...
p26   F38  | $-,D
==================== PAGE 27 ====================
p27   F44  | $*D
p27   F11  | PRESIDEN
p27   F14  | R EPUEL IK IN D ONES IA
p27   F16  | -27 -
p27   F16  [ITEM]
p27   F16  | 22. Ketentuan Pasal 86 diubah sehingga berbunyi sebagai
p27   F14  | berikut:
p27   F12  | Pasal 86
p27   F12  | Setiap orang yang menyembelih:
p27   F18  [SUB-ITEM]
p27   F18  | a. Ternak ruminansia kecit betina produktif
p27   F16  | sebagaimana dimaksud dalam Pasal 18 ayat (4)
p27   F14  | dipidana dengan pidana kurungan paling singkat I
p27   F18  | (satu) bulan dan paling lama 6 (enam) bulan dan
p27   F16  | denda paling sedikit Rp1.000.000,00 (satu juta
p27   F15  | rupiah) dan paling banyak Rp5.000.000,00 (lima juta
p27   F14  | rupiah); atau
p27   F18  [SUB-ITEM]
p27   F18  | b. Ternak ruminansia besar betina produktil
p27   F16  | sebagaimana dimaksud dalam Pasal 18 ayat (4)
p27   F15  | dipidana dengan pidana penjara paling singkat 1
p27   F18  | (satu) tahun dan paling lama 3 (tiga) tahun dan
p27   F15  | denda paling sedikit Rp100.000.000,00 (seratus juta
p27   F15  | rupiah) dan paling banyak Rp300.000.000,00 (tiga
p27   F14  | ratus juta rupiah).
p27   F14  [ITEM]
p27   F14  | 23. Di antara Pasai 91 dan Pasal 92 disisipkan 2 (dua) pasal,
p27   F16  | yakni Pasal 91A dan Pasal 91B sehingga berbunyi
p27   F14  | sebagai berikut:
p27   F12  | Pasal 91A
p27   F14  | Setiap Orang yang memproduksi dan/atau mengedarkan
p27   F16  | Produk Hewan dengan memalsukan produk Hewan
p27   F14  | dan/atau menggunakan bahan tambahan yang dilarang
p27   F14  | sebagaimana dimaksud dalam pasal 58 ayat (6), dipidana
p27   F16  | dengan pidana penjara paling lama 5 (lima) tahun dan
p27   F16  | pidana denda paling banyak Rp10.000.000.000,00
p27   F14  | (sepuluh miliar rupiah).
p27   F14  | Pasal 9lE} ...
==================== PAGE 28 ====================
p28   F12  [ITEM]
p28   F12  | 24.
p28   F12  [ITEM]
p28   F12  | 25.
p28   F11  | PRESIDEN
p28   F12  | R EPUBLII( INDONESIA
p28   F17  | -28-
p28   F12  | Pasal 9 18
p28   F18  [AYAT]
p28   F18  | (1) Setiap Orang ya1rg menganiaya dan/ atau
p28   F14  | menyalahgunakan Hewan sehingga mengakibatkan
p28   F18  | cacat dan/atau tidak produktif sebagaimana
p28   F14  | dimaksud dalam Pasal 66A ayat (1) dipidana dengan
p28   F14  | pidana kurungan paling singkat 1 (satu) bulan dan
p28   F14  | paling lama 6 (enam) bulan dan denda paling sedikit
p28   F15  | Rp1.000.000,00 (satu juta rupiah) dan paling banyak
p28   F14  | Rp5.000.000,00 (lima juta rupiah).
p28   F16  [AYAT]
p28   F16  | (2) Setiap Orang yang mengetahui adanya perbuatan
p28   F14  | sebagaimana dimaksud dalam Pasal 66A ayat (1) dan
p28   F16  | tidak melaporkan kepada pihak yang berwenang
p28   F16  | sebagaimana dimaksud dalam Pasal 66A ayat (2)
p28   F14  | dipidana dengan pidana kurungan paling singkat 1
p28   F14  | (satu) bulan dan paling lama 3 (tiga) bulan dan denda
p28   F16  | paling sedikit Rp1.000.000,0O (satu juta rupiah) dan
p28   F15  | paling banyak Rp3.000.000,00 (tiga juta rupiah).
p28        | Ketentuan Pasal 96 dihapus.
p28   F18  | Di antara Pasal 96 dan Pasal 97 disisipkan I (satu) pasal
p28   F14  | yakni Pasal 96A sehingga berbunyi sebagai berikut:
p28   F12  | Pasal 96A
p28   F16  | Peraturan Pemerintah mengenai pulau karantina
p28   F16  | sebagaimana dimaksud dalam Pasal 36D ayat (2)
p28   F18  | harus telah ditetapkan paling lama 2 (dua) tahun
p28   F14  | terhitung sejak Ltndang-Undang ini diundangkan.
p28   F14  | Peraturan Pemerintah mengenai Otoritas Veteriner
p28   F18  | dan Siskeswanas sebagaimana dimaksud dalam
p28   F14  | Pasal 68E harus telah ditetapkan paling lama 2 (dua)
p28   F18  | tahun terhitung sejak Undang-Undang ini
p28   F14  | diundangkan
p28   F10  [AYAT]
p28   F10  | (1)
p28   F11  [AYAT]
p28   F11  | (2t
p28   F12  | Pasal II
p28   F18  | Undang-Undang ini
p28   F14  | mulai
p28        | diundangkan.
p28   F18  | pada tanggal
p28   F14  I3 | berlaku
p28   F14  I4 | Agar...
==================== PAGE 29 ====================
p29   F11  | PRESIDEN
p29   F14  | R EP LIBL IK IN D ONES IA
p29   F17  | -29-
p29   F18  | Agar setiap orang mengetahuinya, memerintahkan
p29   F16  | pengundangan Undang-Undang ini dengan penempatannya
p29        | dalam Lembaran Negara Republik Indonesia.
p29   F14  | Disahkan di Jakarta
p29        | pada tanggal 17 Oktober 2014
p29   F12  | PRESIDEN REPUBLIK INDONESIA,
p29   F12  I2 | rtd.
p29        | DR. H. SUSILO BAMBANG YUDHOYONO
p29   F14  I1 | Diundangkan di Jakarta
p29        I1 | pada tanggal 17 Oktober 2014
p29        I1 | MENTERI HUKUM DAN HAK ASASI MANUSIA
p29   F12  | REPUBLIK INDONESIA,
p29   F14  | ttd.
p29        | AMIR SYAMSUDDIN
p29        I1 | LEMBARAN NEGARA REPUBLIK INDONESIA TAHUN 2014 NOMOR 338
p29   F12  | Salinan sesuai dengan aslinya
p29   F12  I1 | KEMENTERIAN SEKRETARIAT NEGARA
p29   F12  | REPUBLIK INDONESIA
p29        | Deputi Pemndang-undangan
p29   F12  | Perekonomian,
p29   F12  | Silvanna Djaman
==================== PAGE 30 ====================
p30   F11  | PRESIOEN
p30   F14  | R EPIBL IK INDONESIA
p30   F12  | PENJELASAN
p30   F12  | ATAS
p30   F12  | UNDANG-UNDANG REPUBLIK INDONESIA
p30   F16  | NOMOR 4l TAHUN 2014
p30   F12  | TENTANG
p30   F12  | PERUBAHAN ATAS UNDANG-UNDANG NOMOR 18 TAHUN 2OO9
p30   F12  | TENTANG PETERNAKAN DAN KESEHATAN HEWAN
p30   F18  I1 | A. UMUM
p30   F16  | Pancasila dan Pembukaan Undang-Undang Dasar Negara Republik
p30   F15  | Indonesia Tahun 1945 mengamanatkan negara untuk melindungi segenap
p30   F14  | bangsa Indonesia dan memajukan kesejahteraan umum serta mewujudkan
p30   F17  | keadilan sosial bagi seluruh ralgrat Indonesia. salah satu bentuk
p30   F14  | perlindungan tersebut dilakukan melalui penyelenggaraan peternakan dan
p30   F16  | Kesehatan Hewan dalam kerangka mewujudkan kemandirian dan
p30        | kedaulatan pangan.
p30   F14  | Penyelenggaraan Peternakan dan Kesehatan hewan yang telah diatur
p30   F14  | dalam Undang-Undang Nomor 18 Tahun 2009 tentang peternakan dan
p30   F15  | Kesehatan Hewan terkait dengan pemasukan Benih, Bibit, Bakalan, dan
p30   F14  I1 | Ternak Ruminansia Indukan, serta pencegahan penyakit Hewan belum
p30   F16  I1 | mencapai hasil yang optimal. Selain itu, beberapa pasal dalam undang_
p30   F18  I1 | undang tersebut telah diuji materi di Mahkamah Konstitusi. Dalam
p30   F14  I1 | putusannya, Mahkamah Konstitusi membatalkan beberapa pasal yang
p30   F16  I1 | terkait dengan pemasukan dan pengeluaran produk Hewan, Otoritas
p30   F14  I1 | veteriner, serta persyaratan halal bagi produk Hewan yang dipersyaratkan.
p30   F16  I1 | Atas dasar tersebut serta memenuhi perkembangan dan kebutuhan
p30   F17  I1 | hukum di masyarakat, Undang-Undang Nomor rg rahun 2oo9 tentang
p30   F14  I1 | Peternakan dan Kesehatan Hewan perlu diubah.
p30        | Perubahan ...
==================== PAGE 31 ====================
p31   F11  | PRESIDEN
p31   F14  | R EPLIBL IK INOONESIA
p31   F18  | -2-
p31   F14  | Perubahan tersebut dimaksudkan agar penyelenggaraan peternakan
p31   F15  | dan Kesehatan Hewan dapat mencapai tujuan yang diharapkan, yaitu:
p31   F15  | mengelola sumber daya Hewan secara bermartabat, bertanggung jiwab,
p31   F16  | dan berkelanjutan untuk sebesar-besar kemakmuran rakyat-;-meniukupi
p31   F14  | kebutuhan pangan, barang, dan jasa asal Hewan secara mandiri, berdaya
p31   F15  | saing, dan berkelanjutan bagi peningkatan kesejahteraan peternak dan
p31   F16  | masyarakat; melindungi, mengamankan, dan/atau menjamin wilayah
p31   F18  | Negara Kesatuan Republik Indonesia dari ancaman yang dapat
p31   F14  I1 | mengganggu kesehatan atau kehidupan manusia, Hewan, tumbuhan, dan
p31   F16  I1 | lingkungan; mengembangkan sumber daya Hewan; serta memberi
p31   F14  I1 | kepastian hukum dan kepastia' berusaha dalam bidang peternakan dan
p31   F16  I1 | Kesehatan Hewan. Tujuan penyelenggaraan peternakan dan Kesehatan
p31   F16  I1 | Hewan tersebut harus dilandasi dengan semangat untuk mewujudkan
p31   F15  I1 | kedaulatan, kemandirian, dan ketahan"., p"rrgu..,. Sedangkan asas dari
p31   F14  I1 | penyelenggaraan Peternakan dan Kesehatan Hewan adalah kemanfaatan
p31   F14  I1 | dan keberlanjutan, keamanan dan kesehatan, keralgzatan dan keadilan,
p31   F18  I1 | keterbukaan dan keterpaduan, kemandirian, kemitraan, dan
p31   F12  I1 | keprofesionalan.
p31   F16  | Secara umum perubahan Undang-Undang Nomor 1g Tahun 2OO9
p31   F14  I1 | tentang Peternakan dan Kesehatan Hewan mencakup pemasukan Benih,
p31   F16  I1 | Bibit, Bakalan, Ternak Ruminansia Indukan, dan/atau produk Hewan;
p31   F16  I1 | kemitraan usaha Peternakan; pengaturan mengenai rernak Ruminansia
p31   F15  I1 | Betina Produktif; pencegahan penyakit Hewan; dan penguatan otoritas
p31   F12  I1 | Veteriner.
p31   F18  I1 [SUB-ITEM]
p31   F18  I1 | b. PASAL DEMI PASAL
p31   F12  I1 | Pasal I
p31   F14  | Angka 1
p31   F16  | Pasal I
p31        | Cukup jelas.
p31        | Angka 2
p31   F12  | Pasal 6
p31   F14  | Ayat (1)
p31   F15  | Yang dimaksud dengan .dipertahankan keberadaan dan
p31   F14  | kemanfaatannya secara keberlanjutan,, adalah upaya yang
p31   F18  | perlu dilakukan oleh kabupaten/kota untuk -.-r"rkk^i
p31   F18  | Kawasan Penggembalaan Umum dalam program
p31        | pembangunan daerah.
p31        | Ayat (2) ...
==================== PAGE 32 ====================
p32   F16  | * E"u,I'1['IREU*..,o
p32   F18  | -3-
p32        | Ayat (2)
p32   F14  | Huruf a
p32        | Cukup jelas.
p32   F14  | Huruf b
p32   F16  | Yang dimaksud de ngan ,,kastrasi,, adalah tindakan
p32   F18  | mencegah berfungsinya testis dengan jalan
p32   F14  | menghilangkannya atau menghambat fungsinya.
p32   F14  | Huruf c
p32        | Cukup jelas.
p32   F14  | Huruf d
p32        | Cukup jelas.
p32        | Ayat (3)
p32   F14  | Yang dimaksud dengan "penetapan lahan sebagai Kawasan
p32   F15  | Penggembalaan Umum" yaitu upaya yang harus dilakukan
p32   F18  | oleh pemerintah daerah kabupaten/ kota untuk
p32   F16  | menyediakan lahan penggembalaan umum, antara lain,
p32   F15  | misalnya tanah pangonan, tanah titisara atau tanah kas
p32   F12  | desa.
p32        | Ayat (4)
p32        | Cukup jelas.
p32   F14  | Ayat (5)
p32        | Cukup jelas.
p32        | Angka 3
p32        | Cukup jelas.
p32        | Angka 4
p32   F12  | Pasal 13
p32        | Ayat (1)
p32        | Cukup jelas.
p32   F12  | Ayat (2)
p32        | Cukup jelas.
p32        | Ayat (3) ...
==================== PAGE 33 ====================
p33   F16  | RE"uJ5[=1355*.r,o
p33   F18  | -4-
p33        | Ayat (3)
p33   F18  | Teknologi reproduksi untuk mengembangbiakan hewan
p33   F16  | antara lain melalui alih janin (transfer embrio), kelahiran
p33   F14  | kembar (twinningl, dan pemisahan sperma (sexing) antara
p33        | kromosom X dan kromosom y.
p33        | Ayat (4)
p33        | Cukup jelas.
p33        | Ayat (5)
p33   F14  | Yang dimaksud dengan "Ternak tertentu" adalah Ternak asli
p33   F14  | seperti Sapi Bali dan Ternak lokal seperti Sapi Aceh, Sapi
p33   F14  | Madura, Domba Garut, Ayam Sentul, dan Itik Alabio.
p33        | Ayat (6)
p33   F14  | Yang dimaksud dengan "ciri-ciri keunggulannya,, antara Iain
p33   F14  | memiliki kemampuan produksi dan reproduksi yang tinggi
p33   F14  | dan tahan terhadap penyakit.
p33        | Ayat (7)
p33        | Cukup jelas.
p33   F14  | Ayat (8)
p33        | Cukup jelas.
p33        | Angka 5
p33   F12  | Pasal 15
p33        | Ayat (1)
p33   F14  | Huruf a
p33   F14  | Yang dimaksud dengan "mutu genetik" adalah ekspresi
p33   F14  | keunggulan sifat individu.
p33   F14  | Yang dimaksud dengan "keragaman genetik" adalah
p33   F14  | ekspresi keunggulan variasi genetik antarindividu.
p33   F14  | Huruf b
p33        | Cukup jelas.
p33   F14  | Huruf c
p33   F16  | Yang dimaksud dengan .kekurangan Benih" yaitu
p33   F16  | ketidakcukupan jurnlah Benih (semen atau embrio)
p33   F14  | Ternak bukan asli atau lokal (eksotik) yang digunakan
p33   F18  | untuk kebutuhan pemuliaan dalam rangka
p33   F14  | meningkatkan produktivitas dan/ atau mutu genetik.
p33   F14  | Yang.-.
==================== PAGE 34 ====================
p34   F17  | *."uur.I[t1355*..,o
p34   F18  | -5-
p34   F18  | Yang dimaksud dengan "kekurangan Bibit" yaitu
p34   F17  | ketidakcukupan jumlah Bibit Ternak eksotik yang
p34   F16  | sebelumnya telah dikembangkan atau beradaptasi di
p34   F16  | Indonesia dalam rangka meningkatkan mutu genetik
p34        | Ternak eksotik.
p34   F14  | Huruf d
p34        | Cukup jelas.
p34        | Ayat (2)
p34        | Cukup jelas.
p34        | Ayat (3)
p34        | Cukup jelas.
p34        | Ayat (a)
p34        | Cukup jelas.
p34        | Angka 6
p34        | Pasal 16
p34        | Ayat (1)
p34   F18  | Yang dimaksud dengan "Ternak lokal" adalah hasil
p34   F16  | persilangan antara Ternak asli luar negeri dan Ternak asli
p34   F18  | Indonesia, yang telah dikembangbiakkan di Indonesia
p34   F15  | sampai generasi kelima atau lebih yang teradaptasi pada
p34   F14  | lingkungan dan/ atau manajemen setempat.
p34        | Ayat (2)
p34   F14  | Ketentuan larangan terhadap pengeluaran Benih dan Bibit
p34   F15  | terbaik dimaksudkan untuk mempertahankan populasi dan
p34   F14  | mutu genetik Ternak asli dan lokal.
p34        | Ayat (3)
p34        | Cukup jelas.
p34        | Angka 7
p34        | Pasal 18
p34        | Ayat (1)
p34   F14  | Bibit dalam ketentuan ini hanya ternak ruminansia.
p34        | Ayat (2) ...
==================== PAGE 35 ====================
p35   F16  | *Enu;,.5FSi35!*.r,o
p35   F18  | -6-
p35        | Ayat (2)
p35        | Cukup jelas.
p35   F14  | Ayat (3)
p35        | Cukup jelas.
p35   F14  | Ayat g)
p35        | Cukup jelas.
p35   F14  | Ayat (s)
p35        | Cukup jelas.
p35        | Ayat (6)
p35   F14  | Yang dimaksud dengan "menjaga populasi,, antara lain tidak
p35   F14  | menyembelih anakan ternak ruminansia kecil dan anakan
p35   F14  | ternak ruminansia besar.
p35   F14  | Yang dimaksud dengan ,,anakan ternak ruminansia kecil,,
p35   F16  | adalah ternak ruminansia yang berumur kurang dari 6
p35        | (enam) bulan.
p35   F14  | Yang dimaksud dengan ,,anakan ternak ruminansia besar,,
p35   F15  | adalah ternak ruminansia yang bemmur kurang d,ari 12
p35        | (dua belas) bulan.
p35   F14  | Ayat (7)
p35        | Cukup jelas.
p35        | Angka 8
p35   F12  | Pasal 3 1
p35   F14  | Ayat (1)
p35   F18  | Kemitraan usaha misalnya antara lain, inti plasma,
p35   F16  | subkontrak, keagenan, bagi hasil, atau bentuk iain sesuai
p35   F14  | dengan budaya lokal dan kebiasaan masyarakat setempat.
p35   F12  | Ayat (21
p35   F14  | Huruf a
p35        | Cukup je1as.
p35   F14  | Huruf b
p35        | Cukup jelas.
p35   F14  | Huruf c ...
==================== PAGE 36 ====================
p36   F49  | $).)
p36   F20  | -grc>.€
p36   F18  | *.',,J.T['ISS!*.r,o
p36   F16  | -7 -
p36   F14  | Huruf c
p36   F18  | Yang dimaksud dengan "perusahaan di bidang lain"
p36   F18  | adalah perusahaan di luar bidang Peternakan dan
p36   F16  | Kesehatan Hewan, misalnya antara lain perkebunan,
p36   F14  | perikanan, kehutanan, dan pertambangan.
p36   F14  | Huruf d
p36        | Cukup jelas.
p36        | Ayat (3)
p36        | Cukup jelas.
p36        | Ayat (a)
p36        | Cukup jelas.
p36        | Angka 9
p36   F12  | Pasal 32
p36   F14  | Ayat (l)
p36        | Cukup jelas.
p36        | Ayat (2)
p36   F16  | Yang dimaksud dengan 'pihak tertentu yang mempunyai
p36   F14  | kepentingan khusus' adalah pelaku usaha yang bergerak di
p36   F18  | luar bidang Peternakan yang mempunyai kebutuhan
p36   F16  | terhadap budi daya Ternak, contoh: pelaku usaha yang
p36   F14  | membutuhkan limbah Ternak sebagai penyubur tanah dan
p36   F12  | bio-energi.
p36        | Ayat (3)
p36        | Cukup jelas.
p36   F14  | Angka 10
p36   F12  | Pasal 36
p36        | Ayat (1)
p36        | Cukup jelas.
p36        | Ayat (2)
p36   F14  | Yang dimaksud dengan "pangan bergizi seimbang" adalah
p36   F14  | kondisi pangan yang komposisi protein, lemak, kaibohidrat,
p36   F16  | mineral, vitamin, dan serat kasar dalam satu-kesatuan
p36   F14  | asupan konsumsi sesuai dengan umur, jenis, dan kebutuhan
p36   F14  | untuk aktivitas tubuh.
p36        | Ayat (3) ...
==================== PAGE 37 ====================
p37   F17  | o '.u'5IR'1358*.r,o
p37   F18  | -8-
p37        | Ayat (3)
p37        | Cukup jelas.
p37   F14  | Angka 11
p37   F12  | Pasal 36A
p37   F15  | Yang dimaksud dengan "kebutuhan konsumsi masyarakat"
p37   F14  | adalah kebutuhan menggunakan barang hasil produksi antara
p37   F14  | lain pakaian, dan makanan, guna memenuhi keperluan hidup.
p37   F12  | Pasal 36E}
p37        | Ayat (1)
p37        | Cukup jelas.
p37        | Ayat (2)
p37        | Cukup jelas.
p37        | Ayat (3)
p37        | Cukup jelas.
p37        | Ayat (4)
p37        | Cukup jelas.
p37        | Ayat (5)
p37   F17  | Yang dimaksud dengan 'nilai tambah" antara lain, berat
p37   F14  | maksimal, netralisir residu, dan penyerapan tenaga kerja.
p37        | Ayat (6)
p37        | Cukup jelas.
p37        | Ayat (7)
p37        | Cukup jelas.
p37        | Ayat (8)
p37        | Cukup jelas.
p37   F12  | Pasal 36C
p37   F14  | Ayat (1)
p37   F14  | Yang dimaksud dengan "zona dalam suatu negara,, adalah
p37   F16  [HEADING:BAGIAN]
p37   F16  | bagian dari suatu negara yang mempunyai batas alam,
p37   F16  | status kesehatan populasi Hewan, status epidemiologik
p37        | Penyakit Hewan Menular, dan efektivitas daya kendali.
p37   F12  | Ayat (21 ...
==================== PAGE 38 ====================
p38   F22  | *"rJ5['liS]*..,o
p38   F18  | -9-
p38   F12  | Ayat (2)
p38        | Cukup jelas.
p38        | Ayat (3)
p38        | Cukup jelas.
p38   F14  | Ayat (a)
p38        | Cukup je1as.
p38        | Ayat (s)
p38        | Cukup jelas.
p38   F12  | Pasal 36D
p38   F14  | Ayat (1)
p38   F16  | Yang dimaksud dengan "pulau karantina, adalah suatu
p38   F14  | pulau yang terisolasi dari wilayah pengembangan budi daya
p38   F14  | Ternak, yang disediakan dan dikelola oleh pemerintah untuk
p38   F16  | keperluan pencegahan masuk dan tersebarnya penyakit
p38   F16  | Hewan yang dapat ditimbulkan dari pemasukan Ternak
p38   F14  | Ruminansia Indukan sebelum dilalulintasbebaskan ke dalam
p38   F16  | wilayah Negara Kesatuan Republik Indonesia untuk
p38   F14  | keperluan pengembangan Peternakan.
p38   F16  | Yang dimaksud dengan 'langka waktu tertentu, adalah
p38   F16  | jangka waktu yang dibutuhkan untuk memastikan Ternak
p38   F18  | Ruminansia Indukan bebas dari agen penyakit Hewan
p38        | Menular.
p38        | Ayat (2)
p38        | Cukup jelas.
p38   F12  | Pasal 36E
p38        | Ayat (1)
p38   F14  | Yang dimaksud dengan "dalam hal tertentu,, adalah keadaan
p38   F16  | mendesak, antara lain, akibat bencana, saat masyarakat
p38   F14  | membutuhkan pasokan Ternak dan/atau produk Hewan.
p38   F12  | Ayat (2)
p38        | Cukup jelas.
p38   F14  | Angka 12 ...
==================== PAGE 39 ====================
p39   F20  | * rr,'ii['1355*.r,o
p39   F14  | _10_
p39   F14  | Angka 12
p39   F12  | Pasal 37
p39        | Ayat (1)
p39   F14  | Yang dimaksud dengan "lndustri pengolahan Produk Hewan"
p39   F14  | adalah industri yang melakukan kegiatan penanganan dan
p39   F16  | pemrosesan hasil hewan yang ditujukan untuk mencapai
p39   F14  | nilai tambah yang lebih tinggi, dengan memperhatikan aspek
p39   F17  | produk yang aman, sehat, utuh, dan halal bagi yang
p39        | dipersyaratkan.
p39        | Ayat (2)
p39        | Cukup jelas.
p39        | Ayat (2a)
p39        | Cukup jelas.
p39        | Ayat (3)
p39        | Cukup jelas.
p39   F14  | Angka 13
p39   F12  | Pasal 4 1
p39        | Cukup jelas.
p39   F14  | Angka 14
p39   F12  | Pasal 41A
p39        | Ayat (1)
p39        | Cukup jelas.
p39   F12  | Ayat (2)
p39        | Cukup jelas.
p39        | Ayat (3)
p39   F14  | Koordinasi pencegahan Penyakit Hewan dilaksanakan antara
p39   F16  | lain dengan cara penJrusunan bersama rencana strategis
p39   F16  | pencegahan Penyakit Hewan, pengembangan unit ."spon"
p39   F17  | cepat, pengembangan sistem kendali penyakit, dan
p39   F14  | persiapan pengembangan rantai komando sebagai antisipasi
p39   F14  | munculnya penyakit.
p39        | Ayat (4)
p39        | Cukup jelas.
p39        | Ayat (5) ...
==================== PAGE 40 ====================
p40   F48  | saQ$,*
p40   F11  | PRESiDEN
p40   F11  | REPI]IJI-.It<. IN D ONES IA
p40   F15  | - 11-
p40   F14  | Ayat (5)
p40        | Cukup je1as.
p40   F12  | Pasal 4 1B
p40        | Ayat (1)
p40        | Cukup jelas.
p40   F12  | Ayat (2\
p40        | Cukup jelas.
p40        | Ayat (3)
p40        | Cukup jelas.
p40   F14  | Ayat (4)
p40   F18  | Pemeriksaaan dilakukan di pos lalulintas Hewan dengan
p40   F15  | memerhatikan situasi dan status Penyakit Hewan baik di
p40   F14  | wilayah penerima maupun di wilayah pengirim.
p40   F14  | Ayat (5)
p40        | Cukup jelas.
p40   F14  | Angka 15
p40   F12  | Pasal 58
p40   F14  | Ayat (I)
p40        | Cukup jelas.
p40   F14  | Ayat (2)
p40        | Cukup jelas.
p40        | Ayat (3)
p40        | Cukup jelas.
p40   F14  | Ayat (a)
p40   F14  | Huruf a
p40   F15  | Yang dimaksud dengan "sertihkat veteriner" adalah
p40   F18  | surat keterangan yang dikeluarkan oleh Otoritas
p40   F14  | Veteriner yang menyatakan bahwa Hewan dan Produk
p40   F18  | Hewan telah memenuhi persyaratan keamanan,
p40   F14  | kesehatan, dan keutuhan.
p40   F14  | Huruf b ...
==================== PAGE 41 ====================
p41   F10  | FIRESIDEN
p41   F12  | REErll!i. trl tNDONEStA
p41   F20  | -t2-
p41   F14  | Huruf b
p41        | Cukup jelas.
p41   F14  | Ayat (s)
p41        | Cukup jelas.
p41   F14  | Ayat (6)
p41        | Cukup je1as.
p41   F14  | Ayat (7)
p41        | Cukup jelas.
p41   F14  | Ayat (8)
p41        | Cukup je1as.
p41   F14  | Angka 16
p41   F12  | Pasal 59
p41        | Cukup jelas.
p41   F14  | Angka 17
p41   F12  | Pasal 65
p41        | Cukup jelas.
p41   F14  | Angka 18
p41   F12  | Pasal 66A
p41   F14  | Cukup jelas.
p41   F14  | Angka 19
p41   F12  | Pasal 68
p41        | Cukup jelas.
p41        | Angka 20 ...
==================== PAGE 42 ====================
p42   F11  | PRESIDEN
p42   F12  | REPI.-IBLIK INDONESIA
p42   F14  | 13-
p42        | Angka 20
p42   F12  | Pasal 68A
p42        | Cukup jelas.
p42   F12  | Pasal 68E}
p42        | Cukup jelas.
p42   F12  | Pasal 68C
p42        | Cukup jelas.
p42   F12  | Pasal 68D
p42        | Cukup jelas.
p42        | Pasal 68E
p42        | Cukup jelas.
p42        | Angka 21
p42   F12  | Pasal 85
p42        | Cukup jelas.
p42   F14  | Angka22
p42        | Pasal 86
p42        | Cukup jelas.
p42        | Angka 23
p42   F12  | Pasal 9 1A
p42        | Cukup jelas.
p42   F14  | Pasal 918
p42        | Cukup jelas.
p42   F12  | Angka 24
p42        | Cukup jelas.
p42   F14  | Angka 25 ...
==================== PAGE 43 ====================
p43   F48  | g).)
p43   F26  | -ap*4
p43   F11  | PRESIDEN
p43   F14  | R EPL]BL IK INDONESIA
p43   F17  | -14-
p43        | Angka 25
p43   F12  | Pasal 96A
p43        | Cukup jelas.
p43   F12  I1 | Pasal II
p43        | Cukup jelas.
p43   F12  | TAMBAHAN LEMBARAN NEGARA REPUBLIK INDONESIA NOMOR 5619
```

---


## pp

- **File**: `pp/PP_NO_70_TH_1991.pdf`
- **Document Type**: Peraturan Pemerintah
- **Issued by**: Presiden
- **Pages**: 22 | **Lines**: 738
- **Font sizes**: [10.0, 12.0]
- **Most common font**: 12.0 (91% of lines)
- **Bold font sizes**: [12.0]
- **Indent clusters**: [72.0, 139.0, 216.0, 238.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01   F10  | PRESIDEN
p01   F10  I4 | REPUBLIK INDONESIA
p01 B      | PERATURAN PEMERINTAH REPUBLIK INDONESIA
p01 B      I3 | NOMOR 70 TAHUN 1991
p01 B      | TENTANG
p01 B      I2 | PELAKSANAAN UNDANG-UNDANG NOMOR 4 TAHUN 1990
p01 B      I2 | TENTANG SERAH-SIMPAN KARYA CETAK DAN KARYA-REKAM
p01 B      I3 | PRESIDEN REPUBLIK INDONESIA,
p01        I1 [PREAMBLE:MENIMBANG]
p01        I1 | Menimbang : bahwa untuk melaksanakan ketentuan Undang-undang Nomor 4 Tahun
p01        I2 | 1990 tentang Serah-simpan Karya Cetak dan Karya Rekam, dipandang
p01        I2 | perlu mengatur pelaksanaan serah-simpan karya cetak dan karya
p01        I2 | rekam serta pengelolaannya dengan Peraturan Pemerintah;
p01        I1 [PREAMBLE:MENGINGAT]
p01        I1 | Mengingat
p01        I2 | : 1. Pasal 5 ayat (2) Undang-Undang Dasar 1945;
p01        I2 [ITEM]
p01        I2 | 2. Undang-undang Nomor 4 Tahun 1990 tentang Serah-simpan Karya
p01        | Cetak dan Karya Rekam (Lembaran Negara Tahun 1990 Nomor 48,
p01        | Tambahan Lembaran Negara Nomor 3418);
p01        [KEPUTUSAN:MEMUTUSKAN]
p01        | MEMUTUSKAN :
p01        I1 [PREAMBLE:MENETAPKAN]
p01        I1 | Menetapkan : PERATURAN
p01        I4 | PEMERINTAH
p01        | REPUBLIK
p01        | INDONESIA
p01        | TENTANG
p01        I2 | PELAKSANAAN UNDANG-UNDANG NOMOR 4 TAHUN 1990 TENTANG
p01        I2 | SERAH-SIMPAN KARYA CETAK DAN KARYA REKAM.
p01        [HEADING:BAB]
p01        | BAB I
p01        I4 | KETENTUAN UMUM
p01        | Pasal 1
p01        I2 | Dalam Peraturan Pemerintah ini yang dimaksud dengan:
p01        I2 [ITEM]
p01        I2 | 1. Karya cetak adalah semua jenis terbitan dari setiap karya
p01        | intelektual dan/atau artistik yang dicetak dan digandakan dalam
p01        | bentuk buku, majalah, surat kabar, peta, brosur, dan sejenisnya
p01        | yang diperuntukkan bagi umum.
==================== PAGE 2 ====================
p02   F10  | PRESIDEN
p02   F10  I4 | REPUBLIK INDONESIA
p02   F10  | -  2  -
p02        I2 [ITEM]
p02        I2 | 2. Karya rekam adalah semua jenis rekaman dari setiap karya
p02        | intelektual dan/atau artistik yang direkam dan digandakan dalam
p02        | bentuk pita, piringan, dan bentuk lain sesuai dengan perkembangan
p02        | teknologi yang diperuntukkan bagi umum.
p02        I2 [ITEM]
p02        I2 | 3. Perpustakaan Nasional adalah perpustakaan yang berkedudukan di
p02        | lbukota Negara yang mempunyai tugas untuk menghimpun,
p02        | menyimpan, melestarikan dan mendayagunakan semua karya cetak
p02        | dan karya rekam yang dihasilkan di wilayah Negara Republik
p02        | Indonesia.
p02        I2 [ITEM]
p02        I2 | 4. Perpustakaan Daerah adalah perpustakaan yang berkedudukan di
p02        | ibukota Propinsi yang diberikan tugas untuk menghimpun,
p02        | menyimpan, melestarikan dan mendayagunakan semua karya cetak
p02        | dan karya rekam yang dihasilkan di Daerah.
p02        I2 [ITEM]
p02        I2 | 5. Penerbit adalah setiap orang, persekutuan, badan hukum, baik
p02        | milik Negara maupun swasta yang menerbitkan karya cetak.
p02        I2 [ITEM]
p02        I2 | 6. Pengusaha rekaman adalah setiap orang, persekutuan, badan
p02        | hukum, baik milik Negara maupun swasta yang menghasilkan karya
p02        | rekam.
p02        I2 [ITEM]
p02        I2 | 7. Koleksi adalah kumpulan bahan pustaka, baik tercetak maupun
p02        | terekam yang disimpan dan dikelola perpustakaan.
p02        I2 [ITEM]
p02        I2 | 8. Bibliografi adalah daftar bahan pustaka, baik tercetak maupun
p02        | terekam yang disusun menurut sistem tertentu.
p02        | Pasal 2
p02        I2 | Untuk kepentingan pendidikan, pengembangan ilmu pengetahuan dan
p02        I2 | teknologi, penelitian, dan penyebaran informasi serta pelestarian hasil
p02        I2 | budaya bangsa, setiap:
p02        I2 [ITEM]
p02        I2 | 1. penerbit;
==================== PAGE 3 ====================
p03   F10  | PRESIDEN
p03   F10  I4 | REPUBLIK INDONESIA
p03   F10  | -  3  -
p03        I2 [ITEM]
p03        I2 | 2. pengusaha rekaman;
p03        I2 [ITEM]
p03        I2 | 3. warga negara Indonesia yang hasil karyanya diterbitkan/direkam di
p03        | luar negeri;
p03        I2 [ITEM]
p03        I2 | 4. orang atau badan usaha yang memasukkan karya cetak dan/atau
p03        | karya rekam mengenai Indonesia;
p03        | wajib menyerahkan hasil karya cetak atau karya rekamnya
p03        | kepada Perpustakaan Nasional dan/atau Perpustakaan
p03        | Daerah, atau badan sebagaimana diatur dalam Peraturan
p03        | Pemerintah ini.
p03        [HEADING:BAB]
p03        | BAB II
p03        | PELAKSANAAN SERAH-SIMPAN KARYA CETAK
p03        | Pasal 3
p03        I2 [AYAT]
p03        I2 | (1)
p03        | Setiap penerbit yang berada di wilayah Negara Republik
p03        | Indonesia yang menghasilkan karya cetak, wajib menyerahkan
p03        | karya cetaknya sebanyak 2 (dua) buah setiap judulnya kepada
p03        | Perpustakaan Nasional dan sebuah kepada Perpustakaan Daerah
p03        | yang bersangkutan.
p03        I2 [AYAT]
p03        I2 | (2)
p03        | Setiap warga Negara Indonesia yang hasil karyanya diterbitkan di
p03        | luar negeri, wajib menyerahkan 2 (dua) buah setiap judul kepada
p03        | Perpustakaan Nasional.
p03        I2 [AYAT]
p03        I2 | (3)
p03        | Penyerahan hasil karya cetak sebagaimana dimaksud dalam ayat
p03        [AYAT]
p03        | (1) dan ayat (2) selambat-lambatnya 90 (sembilan puluh) hari
p03        | setelah selesai diterbitkan.
p03        | Pasal 4
p03        I2 [AYAT]
p03        I2 | (1)
p03        | Setiap orang atau badan yang memasukkan karya cetak mengenai
p03        | Indonesia
p03        I4 | ke
p03        | dalam
p03        | wilayah
p03        | Indonesia
p03        | dengan
p03        | maksud
==================== PAGE 4 ====================
p04   F10  | PRESIDEN
p04   F10  I4 | REPUBLIK INDONESIA
p04   F10  | -  4  -
p04        | diperdagangkan yang jumlahnya:
p04        [SUB-ITEM]
p04        | a. lebih dari 10 (sepuluh) buah setiap judulnya; atau
p04        [SUB-ITEM]
p04        | b. kurang dari 10 (sepuluh) buah setiap judul, tetapi dalam
p04        I3 | jangka waktu dua tahun memasukkan lagi karya yang sama
p04        I3 | sehingga jumlahnya melebihi 10 (sepuluh) buah; wajib
p04        I3 | menyerahkan sebuah setiap judulnya kepada Perpustakaan
p04        I3 | Nasional.
p04        I2 [AYAT]
p04        I2 | (2)
p04        | Penyerahan karya cetak sebagaimana dimaksud dalam ayat (1)
p04        | selambat-lambatnya 30 (tiga puluh) hari sejak dikeluarkan dari
p04        | pabean.
p04        | Pasal 5
p04        I2 [AYAT]
p04        I2 | (1)
p04        | Jenis karya cetak yang wajib diserahkan kepada Perpustakaan
p04        | Nasional dan/atau Perpustakaan Daerah terdiri dari :
p04        [SUB-ITEM]
p04        | a. buku fiksi;
p04        [SUB-ITEM]
p04        | b. buku non fiksi;
p04        [SUB-ITEM]
p04        | c. buku rujukan;
p04        [SUB-ITEM]
p04        | d. karya artistik;
p04        [SUB-ITEM]
p04        | e. karya ilmiah yang dipublikasikan;
p04        [SUB-ITEM]
p04        | f. majalah;
p04        [SUB-ITEM]
p04        | g. surat kabar;
p04        [SUB-ITEM]
p04        | h. peta;
p04        [SUB-ITEM]
p04        | i. brosur;
p04        [SUB-ITEM]
p04        | j. karya cetak lain yang ditetapkan oleh Kepala Perpustakaan
p04        I3 | Nasional.
p04        I2 [AYAT]
p04        I2 | (2)
p04        | Selain jenis karya cetak sebagaimana dimaksud dalam ayat (1),
p04        | yang termasuk wajib diserahkan adalah edisi cetakan kedua,
==================== PAGE 5 ====================
p05   F10  | PRESIDEN
p05   F10  I4 | REPUBLIK INDONESIA
p05   F10  | -  5  -
p05        | ketiga dan seterusnya, yang mengalami perubahan isi dan/atau
p05        | bentuk.
p05        | Pasal 6
p05        I2 [AYAT]
p05        I2 | (1)
p05        | Karya cetak yang diserahkan kepada Perpustakaan Nasional
p05        | dan/atau Perpustakaan Daerah harus memenuhi persyaratan
p05        | kualitas atau sama dengan yang diedarkan.
p05        I2 [AYAT]
p05        I2 | (2)
p05        | Karya cetak yang diserahkan tidak dalam bentuk fotokopi.
p05        | Pasal 7
p05        I2 [AYAT]
p05        I2 | (1)
p05        | Penyerahan karya cetak dapat dilakukan secara langsung atau
p05        | dikirimkan melalui Pos tercatat kepada Perpustakaan Nasional
p05        | dan/atau Perpustakaan Daerah.
p05        I2 [AYAT]
p05        I2 | (2)
p05        | Pengiriman melalui Pos sebagaimana dimaksud dalam ayat (1)
p05        | harus dengan cara yang baik dan aman sesuai ketentuan
p05        | pengiriman karya cetak pada umumnya.
p05        I2 [AYAT]
p05        I2 | (3)
p05        | Pengiriman dilakukan selambat-lambatnya dalam waktu yang
p05        | ditentukan dalam Peraturan Pemerintah ini dengan dibuktikan
p05        | tanggal pengiriman karya cetak tersebut.
p05        I2 [AYAT]
p05        I2 | (4)
p05        | Karya cetak yang telah diterima, selanjutnya dicatat oleh
p05        | Perpustakaan
p05        | Nasional
p05        | atau
p05        | Perpustakaan
p05        | Daerah
p05        | yang
p05        | bersangkutan dan kepada pengirim diberikan tanda bukti
p05        | penerimaan.
==================== PAGE 6 ====================
p06   F10  | PRESIDEN
p06   F10  I4 | REPUBLIK INDONESIA
p06   F10  | -  6  -
p06        [HEADING:BAB]
p06        | BAB III
p06        | PELAKSANAAN SERAH SIMPAN KARYA REKAM
p06        | Pasal 8
p06        I2 [AYAT]
p06        I2 | (1)
p06        | Setiap pengusaha rekaman yang berada di wilayah Negara
p06        | Republik Indonesia yang menghasilkan karya rekam dan setiap
p06        | Warga Negara Indonesia yang hasil karyanya direkam di luar
p06        | negeri, wajib menyerahkan sebuah karya rekamnya kepada
p06        | Perpustakaan Nasional dan sebuah kepada Perpustakaan Daerah
p06        | yang bersangkutan.
p06        I2 [AYAT]
p06        I2 | (2)
p06        | Penyerahan hasil karya rekam tersebut selambat-lambatnya 90
p06        | (sembilan puluh) hari sejak disebarluaskan atau dipasarkan.
p06        | Pasal 9
p06        I2 [AYAT]
p06        I2 | (1)
p06        | Setiap orang yang memasukkan karya rekam mengenai Indonesia
p06        | yang jumlahnya:
p06        [SUB-ITEM]
p06        | a. lebih dari 10 (sepuluh) buah setiap judulnya, atau
p06        [SUB-ITEM]
p06        | b. kurang dari 10 (sepuluh) buah setiap judul, tetapi dalam
p06        I3 | jangka waktu dua tahun memasukkan lagi karya rekam yang
p06        I3 | sama sehingga jumlahnya melebihi 10 (sepuluh) buah; wajib
p06        I3 | menyerahkan sebuah setiap judulnya kepada Perpustakaan
p06        I3 | Nasional.
p06        I2 [AYAT]
p06        I2 | (2)
p06        | Penyerahan karya rekam sebagaimana dimaksud dalam ayat (l)
p06        | selambat-lambatnya 30 (tiga puluh) hari sejak dikeluarkan dari
p06        | pabean.
p06        | Pasal 10
p06        I2 [AYAT]
p06        I2 | (1)
p06        | Jenis karya rekam yang wajib diserahkan kepada Perpustakaan
p06        | Nasional dan/atau Perpustakaan Daerah terdiri atas karya
==================== PAGE 7 ====================
p07   F10  | PRESIDEN
p07   F10  I4 | REPUBLIK INDONESIA
p07   F10  | -  7  -
p07        | intelektual dan/ atau artistik yang direkam dan digandakan
p07        | dalam bentuk pita atau piringan, seperti film, kaset audio, kaset
p07        | video, video disk, piringan hitam, disket dan bentuk lain sesuai
p07        | dengan perkembangan teknologi.
p07        I2 [AYAT]
p07        I2 | (2)
p07        | Penyerahan, penyimpanan dan pengelolaan karya rekam berupa
p07        | film ceritera atau dokumenter diatur dengan Peraturan
p07        | Pemerintah tersendiri.
p07        | Pasal 11
p07        I2 | Karya rekam yang diserahkan kepada Perpustakaan Nasional dan/atau
p07        I2 | Perpustakaan Daerah harus memenuhi persyaratan kualitas.
p07        | Pasal 12
p07        I2 [AYAT]
p07        I2 | (1)
p07        | Karya rekam yang diserahkan kepada Perpustakaan Nasional
p07        | dan/atau Perpustakaan Daerah, dapat dilakukan secara langsung
p07        | atau dapat pula secara tidak langsung melalui Pos tercatat.
p07        I2 [AYAT]
p07        I2 | (2)
p07        | Pengiriman melalui Pos, harus dilakukan dengan cara yang baik
p07        | dan aman sesuai dengan ketentuan pengiriman Pos.
p07        I2 [AYAT]
p07        I2 | (3)
p07        | Pengiriman karya rekam dilakukan selambat-lambatnya dalam
p07        | waktu yang telah ditentukan dalam Peraturan Pemerintah ini,
p07        | dengan dibuktikan tanggal pengiriman karya rekam tersebut.
p07        I2 [AYAT]
p07        I2 | (4)
p07        | Karya rekam yang telah diterima, dicatat oleh Perpustakaan
p07        | Nasional atau Perpustakaan Daerah, dan pengirim diberikan
p07        | tanda bukti penerimaan.
==================== PAGE 8 ====================
p08   F10  | PRESIDEN
p08   F10  I4 | REPUBLIK INDONESIA
p08   F10  | -  8  -
p08        [HEADING:BAB]
p08        | BAB IV
p08        I2 | PENYERAHAN DAFTAR JUDUL KARYA CETAK DAN KARYA REKAM
p08        | Pasal 13
p08        I2 [AYAT]
p08        I2 | (1)
p08        | Setiap penerbit di wilayah Negara Republik Indonesia dan setiap
p08        | orang yang bertanggung jawab memasukkan karya cetak
p08        | mengenai Indonesia ke dalam wilayah Negara Republik Indonesia,
p08        | wajib menyerahkan daftar judul karya cetaknya kepada
p08        | Perpustakaan Nasional atau Perpustakaan Daerah.
p08        I2 [AYAT]
p08        I2 | (2)
p08        | Daftar judul sebagaimana dimaksud dalam ayat (1) memuat
p08        | sekurang-kurangnya
p08        | keterangan
p08        | judul
p08        | karya
p08        | cetak,
p08        | nama
p08        | pengarang/penyusun/penerjemah,
p08        | nomor
p08        | cetakan,
p08        | tempat
p08        | terbit, nama penerbit, tahun terbit, nomor jilid, jumlah
p08        | halaman, jenis edisi.
p08        I2 [AYAT]
p08        I2 | (3)
p08        | Daftar judul karya cetak sebagaimana dimaksud dalam ayat (1)
p08        | disampaikan kepada Perpustakaan Nasional dan Perpustakaan
p08        | Daerah
p08        I4 | yang
p08        | bersangkutan
p08        | secara
p08        | berkala
p08        | dan
p08        | sekurang-kurangnya 6 (enam) bulan tahun takwim sekali.
p08        I2 [AYAT]
p08        I2 | (4)
p08        | Daftar judul karya cetak ditandatangani oleh penanggung jawab
p08        | penerbit atau warga negara Indonesia yang karyanya diterbitkan
p08        | di luar negeri atau orang yang bertanggungjawab memasukkan
p08        | karya cetak mengenai Indonesia ke dalam wilayah Negara
p08        | Republik Indonesia.
p08        | Pasal 14
p08        I2 [AYAT]
p08        I2 | (1)
p08        | Setiap pengusaha rekaman di wilayah Negara Republik Indonesia
p08        | dan orang yang bertanggung jawab memasukkan karya rekam
p08        | mengenai Indonesia ke dalam wilayah Negara Republik Indonesia
p08        | wajib menyerahkan daftar judul karya rekamnya kepada
p08        | Perpustakaan Nasional dan Perpustakaan Daerah.
==================== PAGE 9 ====================
p09   F10  | PRESIDEN
p09   F10  I4 | REPUBLIK INDONESIA
p09   F10  | -  9  -
p09        I2 [AYAT]
p09        I2 | (2)
p09        | Daftar judul karya rekam sebagaimana dimaksud dalam ayat (l)
p09        | memuat
p09        | sekurang-kurangnya
p09        | nama
p09        | pencipta/komposer/pengarang/sutradara, judul karya rekam,
p09        | tempat perekaman, nama perusahaan rekaman, dan tahun
p09        | perekaman.
p09        I2 [AYAT]
p09        I2 | (3)
p09        | Daftar judul karya rekam sebagaimana dimaksud dalam ayat (1)
p09        | disampaikan kepada Perpustakaan Nasional dan Perpustakaan
p09        | Daerah secara berkala dan sekurang-kurangnya 6 (enam) bulan
p09        | tahun takwim sekali.
p09        I2 [AYAT]
p09        I2 | (4)
p09        | Daftar judul karya rekam sebagaimana dimaksud dalam ayat (1)
p09        | ditandatangani oleh penanggung jawab rekaman, atau warga
p09        | negara Indonesia yang karyanya direkam di luar negeri atau orang
p09        | yang bertanggung jawab memasukkan karya rekam mengenai
p09        | Indonesia ke dalam wilayah Negara Republik Indonesia.
p09        [HEADING:BAB]
p09        | BAB V
p09        | PENGELOLAAN KARYA CETAK, KARYA REKAM
p09        I2 | DAN DAFTAR JUDUL KARYA CETAK DAN KARYA REKAM
p09        | Pasal 15
p09        I2 [AYAT]
p09        I2 | (1)
p09        | Pengelolaan karya cetak dan karya rekam sebagaimana dimaksud
p09        | dalam Peraturan Pemerintah ini dilakukan oleh:
p09        [SUB-ITEM]
p09        | a. Perpustakaan Nasional;
p09        [SUB-ITEM]
p09        | b. Perpustakaan Daerah.
p09        I2 [AYAT]
p09        I2 | (2)
p09        | Kepala
p09        I3 | Perpustakaan
p09        | Nasional
p09        | bertanggung
p09        | jawab
p09        | atas
p09        | pengelolaan
p09        | karya
p09        | cetak
p09        | dan
p09        | karya
p09        | rekam
p09        | yang
p09        | diserah-simpankan.
p09        I2 [AYAT]
p09        I2 | (3)
p09        | Pengelolaan sebagaimana dimaksud dalam ayat (2) meliputi
p09        | penerimaan,
p09        | penyimpanan,
p09        | pendayagunaan,
p09        | pelestarian,
p09        | pengawasan atas pelaksanaan serah simpan karya cetak dan
p09        | karya rekam.
==================== PAGE 10 ====================
p10   F10  | PRESIDEN
p10   F10  I4 | REPUBLIK INDONESIA
p10   F10  | -  10  -
p10        | Pasal 16
p10        I2 [AYAT]
p10        I2 | (1)
p10        | Perpustakaan
p10        | Nasional
p10        | dan
p10        | Perpustakaan
p10        | Daerah
p10        | wajib
p10        | memberikan tanda bukti penerimaan karya cetak atau karya
p10        | rekam yang memenuhi persyaratan sebagaimana dimaksud dalam
p10        | Pasal 3, Pasal 8 atau Pasal 11.
p10        I2 [AYAT]
p10        I2 | (2)
p10        | Tanda bukti penerimaan sebagaimana dimaksud dalam ayat (1)
p10        | khusus
p10        I4 | untuk
p10        | karya
p10        | cetak
p10        | memuat
p10        | keterangan
p10        | sekurang-kurangnya
p10        | judul
p10        | karya
p10        | cetak,
p10        | nama
p10        | pengarang/penyusun/penerjemah,
p10        | nomor
p10        | cetakan,
p10        | tempat
p10        | terbit, nama penerbit, tahun terbit, nomor jilid, jumlah, dan
p10        | jenis edisi.
p10        I2 [AYAT]
p10        I2 | (3)
p10        | Tanda bukti penerimaan sebagaimana dimaksud dalam ayat (1)
p10        | khusus
p10        I4 | untuk
p10        | karya
p10        | rekam
p10        | memuat
p10        | keterangan
p10        | sekurang-kurangnya
p10        | nama
p10        | pencipta/komposer/pengaransir/sutradara, judul karya rekam,
p10        | jumlah perekaman, nama perusahaan rekaman dan tahun
p10        | perekaman.
p10        | Pasal 17
p10        I2 [AYAT]
p10        I2 | (1)
p10        | Karya cetak dan karya rekam yang diterima oleh Perpustakaan
p10        | Nasional dan perpustakaan Daerah, dicatat, diolah, disimpan,
p10        | dilestarikan, dan didayagunakan sesuai dengan ketentuan
p10        | pengelolaan karya cetak dan karya rekam.
p10        I2 [AYAT]
p10        I2 | (2)
p10        | Karya cetak dan karya rekam yang karena sifatnya dilarang
p10        | Pemerintah untuk diedarkan untuk umum, hanya dapat
p10        | dimanfaatkan untuk kepentingan tertentu setelah mendapat izin
p10        | khusus dari Kepala Perpustakaan Nasional.
p10        I2 [AYAT]
p10        I2 | (3)
p10        | Ketentuan
p10        I4 | pengelolaan
p10        | karya
p10        | cetak
p10        | dan
p10        | karya
p10        | rekam
p10        | sebagaimana dimaksud dalam ayat (1) dan ayat (2) diatur lebih
==================== PAGE 11 ====================
p11   F10  | PRESIDEN
p11   F10  I4 | REPUBLIK INDONESIA
p11   F10  | -  11  -
p11        | lanjut oleh Kepala Perpustakaan Nasional.
p11        | Pasal 18
p11        I2 | Daftar judul karya cetak dan karya rekam yang diserahkan kepada
p11        I2 | Perpustakaan Nasional, Perpustakaan Daerah, disusun, disimpan, dan
p11        I2 | digunakan sebagai alat informasi serta sebagai alat pemantau
p11        I2 | pelaksanaan serah-simpan karya cetak dan karya rekam.
p11        | Pasal 19
p11        I2 [AYAT]
p11        I2 | (1)
p11        | Karya cetak dan karya rekam yang telah diserahkan, dimuat
p11        | dalam Bibliografi Nasional Indonesia yang diterbitkan oleh
p11        | Perpustakaan Nasional, dan Bibliografi Daerah, yang diterbitkan
p11        | oleh Perpustakaan Daerah.
p11        I2 [AYAT]
p11        I2 | (2)
p11        | Bibliografi Nasional Indonesia dan Bibliografi Daerah diterbitkan
p11        | secara berkala sekurang-kurangnya sekali dalam tiga bulan dan
p11        | kumulasi tahunan.
p11        I2 [AYAT]
p11        I2 | (3)
p11        | Bibliografi Nasional Indonesia, Bibliografi Daerah, dan kumulasi
p11        | tahunan sebagaimana dimaksud dalam ayat (2) disampaikan
p11        | kepada orang atau badan yang menyerah-simpankan karya cetak
p11        | dan atau karya rekam,
p11        | Pasal 20
p11        I2 [AYAT]
p11        I2 | (1)
p11        | Perpustakaan
p11        | Nasional,
p11        | dan
p11        | Perpustakaan
p11        | Daerah
p11        | dalam
p11        | menyelenggarakan pengelolaan, dapat:
p11        [SUB-ITEM]
p11        | a. melakukan pemantauan pelaksanaan serah-simpan karya cetak
p11        I3 | dan karya rekam yang menjadi tanggung jawabnya;
p11        [SUB-ITEM]
p11        | b. memberi peringatan kepada para wajib serah-simpan karya
p11        I3 | cetak dan karya rekam yang lalai melakukan kewajibannya;
==================== PAGE 12 ====================
p12   F10  | PRESIDEN
p12   F10  I4 | REPUBLIK INDONESIA
p12   F10  | -  12  -
p12        [SUB-ITEM]
p12        | c. mendayagunakan karya cetak dan karya rekam sesuai dengan
p12        I3 | ketentuan yang berlaku.
p12        I2 [AYAT]
p12        I2 | (2)
p12        | Pendayagunaan karya cetak dan karya rekam sebagaimana
p12        | dimaksud dalam ayat (1) huruf c harus memperhatikan
p12        | keseimbangan
p12        | antara
p12        | peningkatan/pengembangan
p12        | ilmu
p12        | pengetahuan dan kebudayaan dengan ketentuan peraturan
p12        | perundang-undangan yang berlaku.
p12        [HEADING:BAB]
p12        | BAB VI
p12        I4 | KETENTUAN PERALIHAN
p12        | Pasal 21
p12        I2 | Semua ketentuan yang mengatur serah-simpan karya cetak dan karya
p12        I2 | rekam yang telah ada pada saat diundangkan Peraturan Pemerintah ini
p12        I2 | masih tetap berlaku sepanjang tidak bertentangan atau belum diganti
p12        I2 | berdasarkan peraturan Pemerintah ini.
p12        [HEADING:BAB]
p12        | BAB VII
p12        I4 | KETENTUAN PENUTUP
p12        | Pasal 22
p12        I2 | Pelaksanaan teknis ketentuan Peraturan Pemerintah ini diatur lebih
p12        I2 | lanjut oleh Kepala Perpustakaan Nasional.
p12        | Pasal 23
p12        I2 | Peraturan Pemerintah ini mulai berlaku pada tanggal diundangkan.
p12        I2 | Agar setiap orang mengetahuinya, memerintahkan pengundangan
p12        I2 | Peraturan Pemerintah ini dengan penempatannya dalam Lembaran
p12        I2 | Negara Republik Indonesia.
==================== PAGE 13 ====================
p13   F10  | PRESIDEN
p13   F10  I4 | REPUBLIK INDONESIA
p13   F10  | -  13  -
p13        | Ditetapkan di Jakarta
p13        | pada tanggal 28 Desember 1991
p13        | PRESIDEN REPUBLIK INDONESIA
p13        | ttd
p13        | SOEHARTO
p13        I2 | Diundangkan di Jakarta
p13        I2 | pada tanggal 28 Desember 1991
p13        I2 | MENTERI/SEKRETARIS NEGARA
p13        | REPUBLIK INDONESIA
p13        I3 | ttd
p13        I2 | MOERDIONO
==================== PAGE 14 ====================
p14   F10  | PRESIDEN
p14   F10  I4 | REPUBLIK INDONESIA
p14   F10  | -  14  -
p14        | PENJELASAN
p14        | ATAS
p14        | PERATURAN PEMERINTAH REPUBLIK INDONESIA
p14        I4 | NOMOR 70 TAHUN 1991
p14        | TENTANG
p14        I2 | PELAKSANAAN UNDANG-UNDANG NOMOR 4 TAHUN 1990
p14        I2 | TENTANG SERAH-SIMPAN KARYA CETAKDAN KARYA REKAM
p14        I2 | UMUM
p14        I2 | Karya cetak dan karya rekam mempunyai peranan yang sangat penting
p14        I2 | dalam menunjang pembangunan pada umumnya, khususnya dalam
p14        I2 | bidang pendidikan, penelitian, pengembangan ilmu pengetahuan dan
p14        I2 | teknologi serta penyebaran informasi dalam rangka peningkatan
p14        I2 | kecerdasan kehidupan bangsa.
p14        I2 | Oleh karena itu semua terbitan dan rekaman hasil budaya bangsa
p14        I2 | perlu dihimpun dan dilestarikan untuk membentuk koleksi nasional
p14        I2 | yang lengkap. Untuk mewujudkan upaya tersebut, telah diundangkan
p14        I2 | Undang-undang Nomor 4 Tahun 1990 tentang Serah-simpan Karya
p14        I2 | Cetak dan Karya Rekam.
p14        I2 | Agar Undang-undang ini dapat dilaksanakan, perlu dilengkapi dengan
p14        I2 | Peraturan Pemerintah yang mengatur tata cara serah-simpan karya
p14        I2 | cetak dan karya rekam, pengelolaan penerimaan dan penyimpanan
p14        I2 | karya cetak dan karya rekam serta pelestarian dan pendayagunaan
p14        I2 | sebagai koleksi nasional.
p14        I2 | Peraturan Pemerintah ini dimaksudkan pula sebagai pedoman bagi
p14        I2 | mereka yang diwajibkan melaksanakan Undang-undang tentang
p14        I2 | Serah-Simpan Karya Cetak dan Karya Rekam.
p14        I2 | PASAL DEMI PASAL
==================== PAGE 15 ====================
p15   F10  | PRESIDEN
p15   F10  I4 | REPUBLIK INDONESIA
p15   F10  | -  15  -
p15        I2 | Pasal 1
p15        | Cukup jelas
p15        I2 | Pasal 2
p15        | Termasuk kewajiban menyerahkan disini bukan hanya yang
p15        | tercantum dalam ayat ini saja, tetapi badan-badan Pemerintah
p15        | yang menerbitkan dan atau memasukkan karya cetak dan karya
p15        | rekam untuk kepentingan masyarakat umum.
p15        I2 | Pasal 3
p15        | Ayat (1)
p15        | Kewajiban bagi penerbit untuk menyerahkan karya cetak kepada
p15        | Perpustakaan Daerah hanya berlaku bagi penerbit yang berada di
p15        | wilayah Perpustakaan Daerah yang bersangkutan.
p15        | Ayat (2)
p15        | Cukup jelas
p15        | Ayat (3)
p15        | Selesai diterbitkan, dalam arti karya cetak tersebut telah selesai
p15        | dan siap untuk dipasarkan/disebarluaskan pada masyarakat
p15        | umum.
p15        I2 | Pasal 4
p15        | Ayat (1)
p15        | Cukup jelas
p15        | Ayat (2)
p15        | Cukup jelas
==================== PAGE 16 ====================
p16   F10  | PRESIDEN
p16   F10  I4 | REPUBLIK INDONESIA
p16   F10  | -  16  -
p16        I2 | Pasal 5
p16        | Ayat (1)
p16        | Karya cetak yang dimaksud dalam angka ini tidak termasuk: buku
p16        | kerja/buku agenda/buku harian, kartu undangan/kartu ucapan
p16        | selamat
p16        I4 | dan
p16        | sejenis,
p16        | kartu
p16        | nama/tanda
p16        | pengenal,
p16        | kalender/tanggalan,
p16        | surat-surat,
p16        | karya
p16        | ilmiah
p16        | yang
p16        | tidak
p16        | dipublikasikan, label/stiker, spanduk, daftar harga, blanko
p16        | formulir, jadwal perjalanan, neraca keuangan dan yang sejenis,
p16        | laporan yang tidak dipublikasikan, pelbagai jenis karcis, kertas
p16        | penutup dinding, kertas bungkus dan yang sejenis, dan lain-lain
p16        | karya cetak yang bukan karya intelektual dan artistik.
p16        | Ayat (2)
p16        | Cukup jelas
p16        I2 | Pasal 6
p16        | Ayat (1)
p16        | Persyaratan kualitas misalnya penjilidan,jenis kertas, tulisan
p16        | jelas, atau kondisi yang memungkinkan wujud karya cetak bisa
p16        | tahan lama untuk disimpan.
p16        | Ayat (2)
p16        | Cukup jelas
p16        I2 | Pasal 7
p16        | Ayat (1)
p16        | Cukup jelas
p16        | Ayat (2)
==================== PAGE 17 ====================
p17   F10  | PRESIDEN
p17   F10  I4 | REPUBLIK INDONESIA
p17   F10  | -  17  -
p17        | Kewajiban bagi pengusaha rekaman dan warga negara Indonesia
p17        | yang hasil karyanya direkam di luar negeri untuk menyerahkan
p17        | karya rekam kepada Perpustakaan Daerah hanya berlaku bagi
p17        | mereka yang berada di wilayah Perpustakaan Daerah yang
p17        | bersangkutan.
p17        | Ayat (3)
p17        | Cukup jelas
p17        | Ayat (4)
p17        | Cukup jelas
p17        I2 | Pasal 8
p17        | Ayat (1)
p17        | Cukup jelas
p17        | Ayat (2)
p17        | Cukup jelas
p17        I2 | Pasal 9
p17        | Ayat (1)
p17        | Cukup jelas
p17        | Ayat (2)
p17        | Cukup jelas
p17        I2 | Pasal 10
p17        | Ayat (1)
p17        | Jenis karya rekam yang tidak termasuk dalam ayat ini seperti
p17        | kaset rekaman rapat, film rekaman keluarga, video permainan,
==================== PAGE 18 ====================
p18   F10  | PRESIDEN
p18   F10  I4 | REPUBLIK INDONESIA
p18   F10  | -  18  -
p18        | rekaman biru, disket rekaman administrasi kantor, disket
p18        | permainan, dan yang sejenis.
p18        | Ayat (2)
p18        | Pengaturan secara tersendiri ini mengingat bahwa karya rekam
p18        | film ceritera atau dokumenter sifatnya memerlukan penanganan
p18        | secara khusus. Untuk itu perlu adanya suatu badan yang
p18        | menyimpan atau mengelola secara khusus pula.
p18        | Pengaturan tersendiri ini meliputi pengaturan penyimpanan,
p18        | penyerahan, pengelolaan, dan penetapan badan yang berwenang
p18        | untuk menyimpan atau mengelola film ceritera atau film
p18        | dokumenter.
p18        | Pengertian film dokumenter disini adalah film dokumenter yang
p18        | tidak termasuk untuk diserahkan/disimpan di Arsip Nasional
p18        | berdasarkan Undang-undang Kearsipan.
p18        I2 | Pasal 11
p18        | Rekaman yang diserah-simpankan bukan merupakan rekaman utama
p18        | tetapi rekaman hasil penggandaan. Kualitas disini artinya kualitas
p18        | rekaman, bahan baku, keuntuhan, kelengkapan ceritera setelah
p18        | lulus sensor, atau yang memungkinkan bisa tahan lama untuk
p18        | disimpan.
p18        I2 | Pasal 12
p18        | Ayat (1)
p18        | Cukup jelas
p18        | Ayat (2)
p18        | Cukup jelas
p18        | Ayat (3)
==================== PAGE 19 ====================
p19   F10  | PRESIDEN
p19   F10  I4 | REPUBLIK INDONESIA
p19   F10  | -  19  -
p19        | Cukup jelas
p19        | Ayat (4)
p19        | Cukup jelas
p19        I2 | Pasal 13
p19        | Ayat (1)
p19        | Cukup jelas
p19        | Ayat (2)
p19        | Cukup jelas
p19        | Ayat (3)
p19        | Cukup jelas
p19        | Ayat (4)
p19        | Cukup jelas
p19        I2 | Pasal 14
p19        | Ayat (1)
p19        | Cukup jelas
p19        | Ayat (2)
p19        | Cukup jelas
p19        | Ayat (3)
p19        | Cukup jelas
p19        | Ayat (4)
p19        | Cukup jelas
==================== PAGE 20 ====================
p20   F10  | PRESIDEN
p20   F10  I4 | REPUBLIK INDONESIA
p20   F10  | -  20  -
p20        I2 | Pasal 15
p20        | Ayat (1)
p20        | Cukup jelas
p20        | Ayat (2)
p20        | Cukup jelas
p20        | Ayat (3)
p20        | Cukup jelas
p20        I2 | Pasal 16
p20        | Ayat (1)
p20        | Cukup jelas
p20        | Ayat (2)
p20        | Cukup jelas
p20        | Ayat (3)
p20        | Cukup jelas
p20        I2 | Pasal 17
p20        I2 | Ayat
p20        I3 [AYAT]
p20        I3 | (1)
p20        | Cukup jelas
p20        | Ayat (2)
p20        | Cukup jelas
p20        | Ayat (3)
p20        | Cukup jelas
==================== PAGE 21 ====================
p21   F10  | PRESIDEN
p21   F10  I4 | REPUBLIK INDONESIA
p21   F10  | -  21  -
p21        I2 | Pasal 18
p21        | Cukup jelas
p21        I2 | Pasal 19
p21        | Ayat (1)
p21        | Cukup jelas
p21        | Ayat (2)
p21        | Cukup jelas
p21        | Ayat (3)
p21        | Cukup jelas
p21        I2 | Pasal 20
p21        | Ayat (1)
p21        | Cukup jelas
p21        | Ayat (2)
p21        | Karya cetak dan karya rekam yang diserah-simpankan kepada
p21        | Perpustakaan
p21        | Nasional
p21        | atau
p21        | Perpustakaan
p21        | Daerah,
p21        | pada
p21        | hakekatnya bukan semata-mata untuk disimpan. Namun agar
p21        | berguna bagi pemakainya, maka karya cetak dan karya rekam
p21        | tersebut dapat didayagunakan oleh masyarakat baik untuk
p21        | pengembangan ilmu pengetahuan, kebudayaan, maupun kegiatan
p21        | lain yang bermanfaat. Untuk itu pendayagunaan dapat dilakukan
p21        | dengan cara dipinjamkan misalnya untuk penelitian dengan
p21        | dibaca, dipelajari, dilihat sesuai dengan ketentuan yang berlaku.
p21        | Pendayagunaan yang dilakukan oleh Perpustakaan Nasional atau
p21        | Perpustakaan
p21        | Daerah,
p21        | bukan
p21        | dalam
p21        | pengertian
p21        | yang
p21        | seluas-luasnya, misalnya untuk dijual, diperbanyak, atau di
==================== PAGE 22 ====================
p22   F10  | PRESIDEN
p22   F10  I4 | REPUBLIK INDONESIA
p22   F10  | -  22  -
p22        | pertunjukkan di muka umum dengan memungut biaya, tetapi
p22        | harus tetap memperhatikan ketentuan perundang-undangan yang
p22        | berlaku, dalam hal ini,misalnya Undang-undang Hak Cipta,
p22        | Undang-undang
p22        | Pengawasan
p22        | Barang
p22        | Cetakan
p22        | yang
p22        | dapat
p22        | membahayakan ketertiban umum.
p22        I2 | Pasal 21
p22        | Cukup jelas
p22        I2 | Pasal 22
p22        | Cukup jelas
p22        I2 | Pasal 23
p22        | Cukup jelas
```

---


## perppu

- **File**: `perppu/perppu-no-148-tahun-2024.pdf`
- **Document Type**: Perppu (Emergency Regulation)
- **Issued by**: Presiden
- **Pages**: 50 | **Lines**: 2042
- **Font sizes**: [11.0, 11.5, 12.0, 12.5, 13.0, 13.5, 14.0, 14.5, 15.0, 15.5, 16.0, 16.5, 17.0, 17.5, 18.0, 19.0, 19.5, 22.5]
- **Most common font**: 18.0 (19% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [43.0, 71.0, 102.0, 146.0, 348.0, 366.0, 411.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01   F22  I7 | SALINAN
p01   F13  I2 [PREAMBLE:MENIMBANG]
p01   F13  I2 | Menimbang
p01   F12  I2 [PREAMBLE:MENGINGAT]
p01   F12  I2 | Mengingat
p01   F13  I2 [PREAMBLE:MENETAPKAN]
p01   F13  I2 | Menetapkan
p01   F12  | PRESIDEN
p01   F12  | REPUBUK INDONESIA
p01   F12  I4 | PERATURAN PRESIDEN REPUBLIK INDONESIA
p01   F12  | NOMOR 148 TAHUN 2024
p01   F12  | TENTANG
p01   F12  | KEMENTERIAN SEKRETARIAT NEGARA
p01   F12  [PREAMBLE:DENGAN RAHMAT]
p01   F12  | DENGAN RAHMAT TUHAN YANG MAHA ESA
p01   F12  | PRESIDEN REPUBLIK INDONESIA,
p01        | bahwa untuk melaksanakan ketentuan Pasal 11
p01        | Undang-Undang Nomor 39 Tahun 2008 tentang
p01   F16  | Kementerian Negara sebagaimana telah diubah dengan
p01   F14  | Undang-Undang Nomor 61 Tahun 2024 tentang Perubahan
p01        | atas Undang-Undang Nomor 39 Tahun 2OO8 tentang
p01   F14  | Kementerian Negara, perlu menetapkan Peraturan Presiden
p01   F13  | tentang Kementerian Sekretariat Negara;
p01        [ITEM]
p01        | 1. Pasal 4 ayat (1) dan Pasal 17 Undang-Undang Dasar
p01   F14  | Negara Republik Indonesia Tahun 1945;
p01        [ITEM]
p01        | 2. Undang-Undang Nomor 39 Tahun 2008 tentang
p01   F15  | Kementerian Negara (Lembaran Negara Republik
p01   F13  | Indonesia Tahun 2008 Nomor 166, Tambahan Lembaran
p01   F14  | Negara Republik Indonesia Nomor 4916ll sebagaimana
p01   F16  | telah diubah dengan Undang-Undang Nomor 61
p01   F14  | Tahun 2024 tentang Perubahan atas Undang-Undang
p01   F16  | Nomor 39 Tahun 2008 tentang Kementerian Negara
p01   F16  | (Lembaran Negara Republik Indonesia Tahun 2024
p01   F16  | Nomor 225, Tambahan Lembaran Negara Republik
p01   F14  | Indonesia Nomor 699a\
p01        [ITEM]
p01        | 3. Peraturan Presiden Nomor 14O Tahun 2024 tentang
p01   F14  | Organisasi Kementerian Negara (Lembaran Negara
p01   F14  | Republik Indonesia Tahun 2024 Nomor 25O);
p01   F12  | MEMUTUSI(AN:
p01   F14  | PERATURAN PRESIDEN
p01   F12  | SEKRETARIAT NEGARA.
p01   F16  I5 | TENTANG KEMENTERIAN
p01   F14  I1 | SK No 247591 A
p01   F13  [HEADING:BAB]
p01   F13  | BAB I
==================== PAGE 2 ====================
p02   F12  | PRESIDEN
p02   F12  | REPUBUK INDONESIA
p02   F19  | -2-
p02   F12  [HEADING:BAB]
p02   F12  | BAB I
p02   F12  | KETENTUAN UMUM
p02   F12  | Pasal 1
p02   F14  | Dalam Peraturan Presiden ini yang dimaksud dengan:
p02        [ITEM]
p02        | 1. Kementerian Sekretariat Negara yang selanjutnya
p02   F17  | disebut Kementerian adalah kementerian yang
p02        | menyelenggarakan urusan pemerintahan di bidang
p02   F13  | kesekretariatan negara.
p02        [ITEM]
p02        | 2. Menteri adalah menteri yang menyelenggarakan urusan
p02   F14  | pemerintahan di bidang kesekretariatan negara.
p02   F13  [HEADING:BAB]
p02   F13  | BAB II
p02   F12  | KEDUDUKAN, TUGAS, DAN FUNGSI
p02   F12  | Pasal 2
p02   F16  [AYAT]
p02   F16  | (1) Kementerian berada di bawah dan bertanggung jawab
p02   F12  | kepada Presiden.
p02   F14  [AYAT]
p02   F14  | (21 Kementerian dipimpin oleh Menteri.
p02   F12  | Pasal 3
p02   F14  [AYAT]
p02   F14  | (1) Dalam memimpin Kementerian, Menteri dapat dibantu
p02   F14  | oleh wakil menteri sesuai dengan penunjukan Presiden.
p02        [AYAT]
p02        | (21 Wakil menteri diangkat dan diberhentikan oleh
p02   F12  | Presiden.
p02   F14  [AYAT]
p02   F14  | (3) Wakil menteri berada di bawah dan bertanggung jawab
p02   F13  | kepada Menteri.
p02   F15  [AYAT]
p02   F15  | (41 Wakil menteri mempunyai tugas membantu Menteri
p02   F14  | dalam memimpin pelaksanaan tugas Kementerian.
p02   F14  [AYAT]
p02   F14  | (5) Ruang lingkup bidang tugas wakil menteri sebagaimana
p02   F14  | dimaksud pada ayat (41, meliputi:
p02        [SUB-ITEM]
p02        | a. membantu Menteri dalam perumusan dan/atau
p02   F14  | pelaksanaan kebijakan Kementerian; dan
p02        [SUB-ITEM]
p02        | b. mengoordinasikan pencapaian kebijakan strategis
p02   F14  | lintas unit organisasi jabatan pimpinan tinggi madya
p02        | atau jabatan struktural eselon I di lingkungan
p02   F12  | Kementerian.
p02   F16  | Pasal4...
p02   F14  I1 | SK No 247642 A
==================== PAGE 3 ====================
p03   F12  | PRESIDEN
p03   F12  | REPUBUK INDONESIA
p03        | -3-
p03   F12  | Pasal 4
p03   F14  | Menteri dan wakil menteri merupakan satu kesatuan unsur
p03   F14  | pemimpin dalam Kementerian.
p03   F12  | Pasal 5
p03        | Kementerian mempunyai tugas menyelenggarakan
p03        | dukungan teknis, administrasi, dan analisis urusan
p03        | pemerintahan di bidang kesekretariatan negara, serta
p03   F14  | dukungan manajemen kabinet kepada Presiden dan Wakil
p03   F13  | Presiden dalam menyelenggarakan pemerintahan negara.
p03   F12  | Pasal 6
p03   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p03   F13  | Pasal 5, Kementerian menyelenggarakan fungsi:
p03        [SUB-ITEM]
p03        | a. pemberian dukungan teknis dan administrasi
p03        | kerrrmahtanggaan, keprotokolan, pers, dan media
p03   F13  | kepada Presiden;
p03        [SUB-ITEM]
p03        | b. pemberian dukungan teknis dan administrasi
p03   F17  | kerumahtanggaan dan keprotokolan, serta analisis
p03   F16  | kebijakan kepada Wakil Presiden dalam membantu
p03   F13  | Presiden menyelenggarakan pemerintahan negara;
p03        [SUB-ITEM]
p03        | c. pemberian dukungan teknis dan administrasi kepada
p03   F14  | Presiden dalam menyelenggarakan kekuasaan tertinggi
p03   F15  | atas Angkatan Darat, Angkatan Laut, dan Angkatan
p03   F16  | Udara, dalam hal pengangkatan dan pemberhentian
p03   F16  | perwira Tentara Nasional Indonesia dan Kepolisian
p03   F15  | Negara Republik Indonesia, penganugerahan gelar,
p03   F16  | tanda jasa, dan tanda kehormatan, yang wewenang
p03   F14  | penetapannya berada pada Presiden, serta koordinasi
p03   F16  | pengamanan Presiden dan Wakil Presiden beserta
p03   F16  | keluarga termasuk Tamu Negara setingkat Kepala
p03   F13  | Negara/ Kepala Pemerintahan negara asing;
p03   F14  I1 | SK No 247643 A
p03   F13  [SUB-ITEM]
p03   F13  | d. pemberian
==================== PAGE 4 ====================
p04   F12  | PRESIDEN
p04   F12  | REPUBUK INDONESIA
p04        | -4-
p04        [SUB-ITEM]
p04        | d. pemberian dukungan teknis, administrasi, dan analisis
p04   F15  | dalam penyiapan izin prakarsa Rancangan Peraturan
p04   F16  | Perundang-undangan, penyelesaian Rancangan
p04   F13  | Peraturan Perundang-undangan, Rancangan Keputusan
p04   F17  | Presiden, dan Rancangan Instruksi Presiden, dan
p04   F14  | pengundangan Undang-Undang, Peraturan Pemerintah
p04   F14  | Pengganti Undang-Undang, Peraturan Pemerintah, dan
p04   F13  | Peraturan Presiden, serta penyelesaian dan penanganan
p04   F16  | terkait dengan litigasi, penyelesaian permasalahan
p04   F14  | hukum, penyelesaian Rancangan Keputusan Presiden
p04        | mengenai grasi, amnesti, abolisi, rehabilitasi,
p04   F14  | perubahan pidana mati atau perubahan pidana penjara
p04   F15  | seumur hidup, kewarganegaraan Republik Indonesia,
p04   F14  | ekstradisi, dan keanggotaan Indonesia pada organisasi
p04   F13  | internasional;
p04        [SUB-ITEM]
p04        | e. pemberian dukungan teknis, administrasi, dan analisis
p04   F15  | dalam penyelenggaraan hubungan dengan lembaga
p04   F16  | negara, lembaga nonstruktural, lembaga daerah,
p04   F16  | organisasi kemasyarakatan, organisasi politik, dan
p04   F14  | penanganan pengaduan masyarakat kepada Presiden,
p04        | Wakil Presiden, dan/atau Menteri, serta
p04   F12  | penyelenggaraan kemitraan ;
p04        [SUB-ITEM]
p04        | f. pemberian dukungan teknis, administrasi, dan analisis
p04   F17  | dalam pengangkatan, pemberhentian, dan pensiun
p04   F14  | pejabat negara, pejabat pemerintahan, pejabat lainnya,
p04        | dan Aparatur Sipil Negara yang wewenang
p04   F14  | penetapannya berada pada Presiden, serta pemberian
p04   F14  | dukungan teknis kepada Tim Penilai Akhir;
p04        [SUB-ITEM]
p04        | g. pemberian dukungan teknis, administrasi, dan analisis
p04   F17  | dalam pengkajian, pemberian rekomendasi, dan
p04        | penyelesaian masalah kebijakan dan program
p04   F13  | pemerintah;
p04        [SUB-ITEM]
p04        | h. pemberian dukungan teknis, administrasi, dan analisis
p04   F16  | dalam penyelenggaraan sidang kabinet, rapat, atau
p04   F16  | pertemuan yang dipimpin dan/atau dihadiri oleh
p04   F16  | Presiden dan/atau Wakil Presiden, dan penyiapan
p04   F14  | naskah bagi Presiden dan/atau Wakil Presiden;
p04        [SUB-ITEM]
p04        | i. pemberian dukungan teknis, administrasi, dan analisis
p04   F14  | pelaksanaan penerjemahan, serta pembinaan jabatan
p04   F14  | fungsional penerjemah dan analis kerja sama;
p04   F16  [SUB-ITEM]
p04   F16  | j. pemberian. . .
p04   F14  I1 | SK No 24764 A
==================== PAGE 5 ====================
p05   F12  | PRESIDEN
p05   F12  | REPUBLIK INDONESIA
p05        | -5-
p05        [SUB-ITEM]
p05        | j. pemberian dukungan teknis, administrasi, dan analisis
p05        | dalam pengembangan dan pengelolaan teknologi
p05   F14  | informasi dan komunikasi di lingkungan Kementerian;
p05        [SUB-ITEM]
p05        | k. pembinaan, penataan, dan pengembangan Aparatur
p05   F16  | Sipil Negara, organisasi, tata laksana, dan reformasi
p05   F14  | birokrasi di lingkungan Kementerian;
p05        [SUB-ITEM]
p05        | l. koordinasi dan perLrmusan peraturan perundang-
p05   F16  | undangan serta pelaksanaan advokasi hukum dan
p05   F14  | litigasi di lingkungan Kementerian;
p05   F17  [SUB-ITEM]
p05   F17  | m. koordinasi pelaksanaan tugas, pembinaan, dan
p05        | pemberian dukungan administrasi di lingkungan
p05   F16  | Kementerian, serta pengelolaan arsip kepresidenan,
p05   F16  | pemberian dukungan prasarana dan sarana untuk
p05   F14  | mantan Presiden, mantan Wakil Presiden, dan pejabat
p05   F14  | negara tertentu, serta dukungan administrasi kepada
p05   F13  | Dokter Kepresidenan;
p05        [SUB-ITEM]
p05        | n. pengelolaan barang milik/kekayaan negara yang
p05   F14  | menjadi tanggung jawab Kementerian;
p05        [SUB-ITEM]
p05        | o. penyelenggaraan koordinasi dan fasilitasi kerja sama
p05   F16  | teknik antara Pemerintah Indonesia dengan mitra
p05        | pembangunan, dan penanganan administrasi
p05   F14  | perjalanan dinas luar negeri;
p05        [SUB-ITEM]
p05        | p. pengawasan atas pelaksanaan tugas di lingkungan
p05   F13  | Kementerian; dan
p05        [SUB-ITEM]
p05        | q. pelaksanaan fungsi lain yang diberikan oleh Presiden
p05        | dan Wakil Presiden serta oleh peraturan
p05   F13  | perundang-undangan.
p05   F14  [HEADING:BAB]
p05   F14  | BAB III
p05   F12  | ORGANISASI
p05   F13  [HEADING:BAGIAN]
p05   F13  | Bagian Kesatu
p05   F13  | Susunan Organisasi
p05   F12  | Pasal 7
p05   F14  | Susunan organisasi Kementerian terdiri atas:
p05        [SUB-ITEM]
p05        | a. Sekretariat Kementerian;
p05        [SUB-ITEM]
p05        | b. Sekretariat Presiden;
p05   F16  I7 [SUB-ITEM]
p05   F16  I7 | c.Sekretariat...
p05   F14  I1 | SK No 247645 A
==================== PAGE 6 ====================
p06   F12  | PRESIDEN
p06   F12  | REPUBLIK INDONESIA
p06        | -6-
p06        [SUB-ITEM]
p06        | c. Sekretariat Wakil Presiden;
p06        [SUB-ITEM]
p06        | d. Sekretariat Militer Presiden;
p06        [SUB-ITEM]
p06        | e. Sekretariat Dukungan Kabinet;
p06        [SUB-ITEM]
p06        | f. Deputi Bidang Perundang-undangan dan Administrasi
p06   F14  | Hukum;
p06        [SUB-ITEM]
p06        | g. Deputi Bidang Hubungan Kelembagaan dan
p06   F13  | Kemasyarakatan;
p06        [SUB-ITEM]
p06        | h. Deputi Bidang Administrasi Aparatur;
p06        [SUB-ITEM]
p06        | i. Badan Teknologi, Data, dan Informasi;
p06        [SUB-ITEM]
p06        | j. Staf Ahli Bidang Politik, Pertahanan, dan Keamanan;
p06        [SUB-ITEM]
p06        | k. Staf Ahli Bidang Ekonomi, Kemaritiman, pembangunan
p06   F13  | Manusia, dan Kebudayaan;
p06        [ITEM]
p06        | 1. Staf Ahli Bidang Hukum, Hak Asasi Manusia, dan
p06   F13  | Pemerintahan;
p06   F18  [SUB-ITEM]
p06   F18  | m. Staf Ahli Bidang Aparatur Negara dan Reformasi
p06   F14  | Birokrasi; dan
p06        [SUB-ITEM]
p06        | n. Staf Ahli Bidang Komunikasi Politik dan Kehumasan.
p06   F13  [HEADING:BAGIAN]
p06   F13  | Bagian Kedua
p06   F13  | Sekretariat Kementerian
p06   F12  | Pasal 8
p06        [AYAT]
p06        | (1) Sekretariat Kementerian berada di bawah dan
p06   F14  | bertanggung jawab kepada Menteri.
p06   F17  [AYAT]
p06   F17  | (21 Sekretariat Kementerian dipimpin oleh Sekretaris
p06   F12  | Kementerian.
p06   F12  | Pasal 9
p06        | Sekretariat Kementerian mempunyai tugas
p06   F18  | menyelenggarakah koordinasi pelaksanaan tugas,
p06   F14  | pembinaan, dan pemberian dukungan administrasi kepada
p06   F14  | seluruh unsur organisasi di lingkungan Kementerian.
p06   F14  I1 | SK No 247646 A
p06   F13  | Pasal 10
==================== PAGE 7 ====================
p07   F11  | PRESTDEN
p07   F12  | REPUBUK INDONESIA
p07   F16  | -7 -
p07   F13  | Pasal 10
p07   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p07   F14  | Pasal 9, Sekretariat Kementerian menyelenggarakan fungsi:
p07        [SUB-ITEM]
p07        | a. koordinasi kegiatan Kementerian;
p07        [SUB-ITEM]
p07        | b. koordinasi dan pen5rusunan rencana, program, dan
p07   F15  | anggaran di lingkungan Kementerian dan lembaga lain
p07   F14  | yang anggarannya secara administratif dikoordinasikan
p07        | oleh Kementerian, serta pengembangan sistem
p07   F14  | akuntabilitas kinerja;
p07        [SUB-ITEM]
p07        | c. pembinaan dan pemberian dukungan administrasi yang
p07        | meliputi keuangan di lingkungan Kementerian dan
p07   F16  | lembaga lain yang anggarannya secara administratif
p07   F14  | dikoordinasikan oleh Kementerian;
p07        [SUB-ITEM]
p07        | d. pembinaan dan pemberian dukungan ketatausahaan,
p07   F16  | kerumahtanggaan, keprotokolan, kesehatan, dan
p07   F14  | hubungan masyarakat;
p07        [SUB-ITEM]
p07        | e. pemberian dukungan prasarana dan sarana, serta
p07   F14  | sumber daya manusia untuk mantan Presiden, mantan
p07   F16  | Wakil Presiden, dan pejabat negara tertentu, serta
p07   F14  | dukungan administrasi kepada Dokter Kepresidenan;
p07        [SUB-ITEM]
p07        | f. pelaksanaan penyiapan dukungan strategis kepada
p07   F13  | Menteri;
p07        [SUB-ITEM]
p07        | g. koordinasi perencanaan, pelaksanaan, monitoring,
p07   F17  | evaluasi, dan fasilitasi kerja sama teknik antara
p07   F16  | Pemerintah Indonesia dengan mitra pembangunan,
p07   F14  | serta penanganan administrasi perjalanan dinas luar
p07   F12  | negeri;
p07        [SUB-ITEM]
p07        | h. koordinasi dan penyelenggaraan pengelolaan barang
p07   F16  | milik/kekayaan negara dan pengelolaan pengadaan
p07   F14  | barangljasa; dan
p07        [SUB-ITEM]
p07        | i. pelaksanaan fungsi lain yang diberikan oleh Menteri.
p07   F13  | Pasal 1 1
p07   F16  [AYAT]
p07   F16  | (1) Sekretariat Kementerian terdiri atas paling banyak
p07   F14  | 7 (tujuh) biro.
p07   F15  [AYAT]
p07   F15  | (2) Biro sebagaimana dimaksud pada ayat (1) terdiri atas
p07   F14  | jabatan fungsional dan jabatan pelaksana.
p07   F14  I1 | SK No 247647 A
p07   F13  [AYAT]
p07   F13  | (3) Dalam
==================== PAGE 8 ====================
p08   F12  | PRESIDEN
p08   F12  | REPUBUK INDONESIA
p08        | -8-
p08        [AYAT]
p08        | (3) Dalam hal tugas dan fungsi biro tidak dapat
p08   F16  | dilaksanakan oleh jabatan fungsional sebagaimana
p08   F14  | dimaksud pada ayat (21, dapat dibentuk paling banyak
p08   F14  | 3 (tiga) bagian.
p08   F16  | (a) Dikecualikan dari ketentuan sebagaimana dimaksud
p08   F14  | pada ayat (3), biro yang menangani fungsi umum terdiri
p08        | atas jabatan fungsional dan jabatan pelaksana
p08   F14  | dan/atau paling banyak 7 (tujuh) bagian.
p08   F16  [AYAT]
p08   F16  | (5) Bagian sebagaimana dimaksud pada ayat (3) dan
p08   F16  | ayat (4) terdiri atas jabatan fungsional dan jabatan
p08   F14  | pelaksana danlatau paling banyak 3 (tiga) subbagian.
p08   F16  [AYAT]
p08   F16  | (6) Dikecualikan dari ketentuan sebagaimana dimaksud
p08   F14  | pada ayat (5), bagian yang menangani ketatausahaan
p08   F14  | pimpinan terdiri atas sejumlah subbagian sesuai dengan
p08   F14  | kebutuhan.
p08   F13  [HEADING:BAGIAN]
p08   F13  | Bagian Ketiga
p08   F13  | Sekretariat Presiden
p08   F13  | Pasal 12
p08   F15  [AYAT]
p08   F15  | (1) Sekretariat Presiden berada di bawah dan bertanggung
p08   F14  | jawab kepada Menteri.
p08   F14  [AYAT]
p08   F14  | (21 Sekretariat Presiden dipimpin oleh Sekretaris Presiden.
p08   F14  [AYAT]
p08   F14  | (3) Dalam melaksanakan tugasnya, Sekretaris Presiden
p08   F14  | dapat menerima penugasan langsung dari Presiden.
p08   F13  | Pasal 13
p08   F14  | Sekretariat Presiden mempunyai tugas menyelenggarakan
p08        | pemberian dukungan teknis dan
p08   F14  | administrasi
p08   F14  | kerumahtanggaan, keprotokolan, pers, dan media kepada
p08   F12  | Presiden.
p08   F13  | Pasal 14
p08   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p08   F13  | Pasal 13, Sekretariat Presiden menyelenggarakan fungsi:
p08        [SUB-ITEM]
p08        | a. pelayanan kerumahtanggaan Presiden dan/atau
p08   F14  | Istri/Suami Presiden;
p08        [SUB-ITEM]
p08        | b. koordinasi pengelolaan istana Kepresidenan, museum,
p08   F13  | dan koleksi benda-benda seni;
p08   F14  I7 [SUB-ITEM]
p08   F14  I7 | c. pelaksanaan.
p08   F14  I1 | SK No 247648 A
==================== PAGE 9 ====================
p09   F12  | PRESIDEN
p09   F12  | REPUBUK INDONESIA
p09        | -9-
p09        [SUB-ITEM]
p09        | c. pelaksanaan urusan keprotokolan dan acara perjalanan
p09        | Presiden dan/atau Istri/Suami Presiden di dalam
p09   F14  | maupun di luar negeri;
p09        [SUB-ITEM]
p09        | d. koordinasi kegiatan pers dan media kegiatan Presiden
p09   F16  | dan/atau Istri/Suami Presiden, serta acara lainnya
p09   F14  | di lingkungan Sekretariat Presiden;
p09        [SUB-ITEM]
p09        | e. koordinasi pemberian dukungan administrasi yang
p09   F18  | meliputi ketatausahaan, sumber daya manusia,
p09   F14  | keuangan, kerumahtanggaan, arsip, dan dokumentasi
p09        | kepada seluruh unsur organisasi di lingkungan
p09   F13  | Sekretariat Presiden;
p09        [SUB-ITEM]
p09        | f. pengelolaan dana operasional dan bantuan
p09   F13  | kemasyarakatan Presiden;
p09        [SUB-ITEM]
p09        | g. pemberian petunjuk teknis di bidang kerumahtanggaan
p09   F15  | dan keprotokolan kepada para Ajudan Presiden dan
p09   F14  | Ajudan Istri/ Suami Presiden;
p09        [SUB-ITEM]
p09        | h. koordinasi Dokter Kepresidenan dalam rangka
p09   F16  | pemberian layanan kesehatan Presiden dan/atau
p09   F14  | Istri/Suami Presiden;
p09        [SUB-ITEM]
p09        | i. pemberian dukungan prasarana dan sarana serta hak
p09   F15  | keuangan bagi Penasihat Khusus Presiden, Utusan
p09   F14  | Khusus Presiden, dan Staf Khusus Presiden; dan
p09        [SUB-ITEM]
p09        | j. pelaksanaan fungsi lain yang diberikan oleh Presiden
p09   F13  | dan Menteri.
p09   F13  | Pasal 15
p09   F14  | Sekretariat Presiden terdiri atas:
p09        [SUB-ITEM]
p09        | a. Deputi Bidang Administrasi dan Pengelolaan Istana; dan
p09        [SUB-ITEM]
p09        | b. Deputi Bidang Protokol, Pers, dan Media.
p09   F13  [HEADING:PARAGRAF]
p09   F13  | Paragraf 1
p09   F14  | Deputi Bidang Administrasi
p09   F13  | dan Pengelolaan Istana
p09   F13  | Pasal 16
p09   F16  [AYAT]
p09   F16  | (1) Deputi Bidang Administrasi dan Pengelolaan Istana
p09        | berada di bawah dan bertanggung jawab kepada
p09   F13  | Sekretaris Presiden.
p09   F16  [AYAT]
p09   F16  | (21 Deputi Bidang Administrasi dan Pengelolaan Istana
p09   F14  | dipimpin oleh Deputi.
p09        | Pasal 17 ...
p09   F14  I1 | SK No 247649 A
==================== PAGE 10 ====================
p10   F12  | PRESIDEN
p10   F12  | REPUBLIK INDONESIA
p10   F17  | -10-
p10   F13  | Pasal 17
p10        | Deputi Bidang Administrasi dan Pengelolaan Istana
p10   F14  | mempunyai tugas membantu Sekretaris Presiden dalam
p10   F18  | menyelenggarakan koordinasi pelaksanaan tugas
p10   F14  | pemberian dukungan administrasi kepada seluruh unsur
p10   F15  | organisasi di lingkungan Sekretariat Presiden, pengelolaan
p10   F14  | istana-istana Kepresidenan, museum, koleksi benda-benda
p10   F16  | seni, dan pengelolaan dana operasional dan bantuan
p10   F13  | kemasyarakatan Presiden, serta pelayanan kegiatan penting
p10   F14  | lainnya di lingkungan Sekretariat Presiden.
p10   F13  | Pasal 18
p10   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p10   F16  | Pasal 17, Deputi Bidang Administrasi dan Pengelolaan
p10   F13  | Istana menyelenggarakan fungsi:
p10        [SUB-ITEM]
p10        | a. pen5rusunan rencana, program, anggaran, dan
p10        | akuntabilitas kinerja di lingkungan Sekretariat
p10   F12  | Presiden;
p10        [SUB-ITEM]
p10        | b. perencanaan dan
p10   F18  I5 | pelaksanaan dukungan
p10   F17  | kerumahtanggaan Presiden dan/atau Istri/Suami
p10   F14  | Presiden, Tamu Negara dan kegiatan penting lainnya
p10   F14  | yang meliputi kegiatan jamuan, tata graha, peralatan,
p10   F13  | dan seni budaya;
p10        [SUB-ITEM]
p10        | c. perencanaan dan pelaksanaan pengelolaan istana-
p10   F14  | istana Kepresidenan;
p10        [SUB-ITEM]
p10        | d. pemberian dukungan prasarana dan sarana serta hak
p10   F16  | keuangan bagi Penasihat Khusus Presiden, Utusan
p10   F14  | Khusus Presiden, dan Staf Khusus Presiden;
p10        [SUB-ITEM]
p10        | e. pengelolaan dana operasional dan bantuan
p10   F13  | kemasyarakatan Presiden;
p10        [SUB-ITEM]
p10        | f. pengelolaan perpustakaan, museum, dan koleksi
p10   F13  | benda-benda seni Kepresidenan;
p10        [SUB-ITEM]
p10        | g. pemberian dukungan administrasi yang meliputi
p10   F16  | ketatausahaan, sumber daya manusia, keuangan,
p10   F14  | kerumahtanggaan, organisasi dan tata laksana, arsip,
p10   F14  | dan dokumentasi di lingkungan Sekretariat Presiden;
p10        [SUB-ITEM]
p10        | h. penyelenggaraan pengelolaan barang milik/kekayaan
p10   F14  | negara di lingkungan Sekretariat Presiden; dan
p10        [SUB-ITEM]
p10        | i. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p10   F12  | Presiden.
p10   F15  | Pasal 19. . .
p10   F14  I1 | SK No 247650 A
==================== PAGE 11 ====================
p11   F12  | PRESIDEN
p11   F12  | REPUBUK INDONESIA
p11   F16  | - 11-
p11   F13  | Pasal 19
p11   F16  [AYAT]
p11   F16  | (1) Deputi Bidang Administrasi dan Pengelolaan Istana
p11   F14  | terdiri atas paling banyak 3 (tiga) biro.
p11   F15  [AYAT]
p11   F15  | (21 Biro sebagaimana dimaksud pada ayat (1) terdiri atas
p11   F14  | jabatan fungsional dan jabatan pelaksana.
p11        [AYAT]
p11        | (3) Dalam hal tugas dan fungsi biro tidak dapat
p11   F16  | dilaksanakan oleh jabatan fungsional sebagaimana
p11   F14  | dimaksud pada ayat (21, dapat dibentuk paling banyak
p11   F14  | 5 (lima) bagian.
p11   F14  [AYAT]
p11   F14  | (41 Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p11   F15  | jabatan fungsional dan jabatan pelaksana dan/atau
p11   F14  | paling banyak 3 (tiga) subbagian.
p11   F16  [AYAT]
p11   F16  | (5) Dikecualikan dari ketentuan sebagaimana dimaksud
p11   F14  | pada ayat (41, bagian yang menangani ketatausahaan
p11   F14  | pimpinan terdiri atas sejumlah subbagian sesuai dengan
p11   F14  | kebutuhan.
p11   F12  [HEADING:PARAGRAF]
p11   F12  | Paragraf 2
p11   F14  | Deputi Bidang Protokol, Pers, dan Media
p11   F12  | Pasal 20
p11   F17  [AYAT]
p11   F17  | (1) Deputi Bidang Protokol, Pers, dan Media berada
p11   F18  | di bawah dan bertanggung jawab kepada Sekretaris
p11   F12  | Presiden.
p11   F14  [AYAT]
p11   F14  | (21 Deputi Bidang Protokol, Pers, dan Media dipimpin oleh
p11   F13  | Deputi.
p11   F12  | Pasal 21
p11   F14  | Deputi Bidang Protokol, Pers, dan Media mempunyai tugas
p11   F14  | membantu Sekretaris Presiden dalam menyelenggarakan
p11   F14  | urusan keprotokolan, pers, media, pelayanan informasi, dan
p11   F16  | dokumentasi kegiatan Presiden dan/atau Istri/Suami
p11   F12  | Presiden.
p11   F14  I1 | SK No 247651 A
p11   F14  | Pasal 22 .
==================== PAGE 12 ====================
p12   F12  | PRESIDEN
p12   F12  | REPUBLIK INDONESIA
p12   F20  | -t2-
p12   F12  | Pasal 22
p12   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p12   F18  | Pasal 21, Deputi Bidang Protokol, Pers, dan Media
p12   F13  | menyelenggarakan fungsi :
p12        [SUB-ITEM]
p12        | a. perencanaan dan pelaksanaan penyelenggaraan
p12   F15  | keprotokolan kegiatan Presiden dan/atau Istri/Suami
p12   F14  | Presiden, Tamu Negara, dan kegiatan penting lainnya
p12   F14  | di dalam maupun di luar negeri;
p12        [SUB-ITEM]
p12        | b. perencanaan dan penyelenggaraan kegiatan pers dan
p12   F14  | media, peliputan dan analisis berita kegiatan Presiden
p12   F16  | dan/atau Istri/Suami Presiden, Tamu Negara, dan
p12   F13  | kegiatan penting lainnya;
p12        [SUB-ITEM]
p12        | c. perencanaan dan pelaksanaan kegiatan pelayanan
p12   F15  | informasi, data, dan dokumentasi kegiatan Presiden
p12   F16  | dan/atau Istri/Suami Presiden, Tamu Negara, dan
p12        | kegiatan penting lainnya di dalam maupun di luar
p12   F13  | negeri; dan
p12        [SUB-ITEM]
p12        | d. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p12   F12  | Presiden.
p12   F12  | Pasal 23
p12   F15  [AYAT]
p12   F15  | (1) Deputi Bidang Protokol, Pers, dan Media terdiri atas
p12   F14  | paling banyak 2 (dua) biro.
p12   F15  [AYAT]
p12   F15  | (2) Biro sebagaimana dimaksud pada ayat (1) terdiri atas
p12   F14  | jabatan fungsional dan jabatan pelaksana.
p12        [AYAT]
p12        | (3) Dalam hal tugas dan fungsi biro tidak dapat
p12   F16  | dilaksanakan oleh jabatan fungsional sebagaimana
p12   F14  | dimaksud pada ayat (2), dapat dibentuk paling banyak
p12   F14  | 4 (empat) bagian.
p12   F14  [AYAT]
p12   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p12   F15  | jabatan fungsional dan jabatan pelaksana dan/atau
p12   F14  | paling banyak 3 (tiga) subbagian.
p12   F12  | Pasal 24
p12   F16  [AYAT]
p12   F16  | (1) Di lingkungan Sekretariat Presiden terdapat beberapa
p12   F14  | Istana Kepresidenan.
p12   F16  [AYAT]
p12   F16  | (21 Istana Kepresidenan sebagaimana dimaksud pada
p12   F14  | ayat (1) diatur lebih lanjut dengan Peraturan Menteri.
p12   F14  I1 | SK No 247785 A
p12   F12  | Bagian
==================== PAGE 13 ====================
p13   F11  | PRESTDEN
p13   F12  | REPUBUK INDONESIA
p13   F17  | -13-
p13   F13  [HEADING:BAGIAN]
p13   F13  | Bagian Keempat
p13   F14  | Sekretariat Wakil Presiden
p13   F12  | Pasal 25
p13        [AYAT]
p13        | (1) Sekretariat Wakil Presiden berada di bawah dan
p13   F14  | bertanggung jawab kepada Menteri.
p13   F16  [AYAT]
p13   F16  | (21 Sekretariat Wakil Presiden dipimpin oleh Sekretaris
p13   F14  | Wakil Presiden.
p13   F16  [AYAT]
p13   F16  | (3) Dalam melaksanakan tugasnya, Sekretaris Wakil
p13   F15  | Presiden dapat menerima penugasan langsung dari
p13   F14  | Wakil Presiden.
p13   F12  | Pasal 26
p13        | Sekretariat Wakil Presiden mempunyai tugas
p13        | menyelenggarakan analisis kebijakan dan pemberian
p13   F14  | dukungan teknis dan administrasi kerumahtanggaan dan
p13   F16  | keprotokolan kepada Wakil Presiden dalam membantu
p13   F13  | Presiden menyelenggarakan pemerintahan negara.
p13   F13  | Pasal2T
p13   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p13   F16  | Pasal 26, Sekretariat Wakil Presiden menyelenggarakan
p13   F13  | fungsi:
p13        [SUB-ITEM]
p13        | a. pemberian analisis kebijakan, dukungan data, dan
p13        | informasi di bidang perekonomian, pariwisata,
p13   F15  | transformasi digital, peningkatan kesejahteraan dan
p13   F14  | pembangunan sumber daya manusia, pemerintahan,
p13   F14  | dan pemerataan pembangunan kepada Wakil Presiden;
p13        [SUB-ITEM]
p13        | b. pelayanan kerumahtanggaan Wakil Presiden dan/atau
p13   F14  | Istri/ Suami Wakil Presiden;
p13        [SUB-ITEM]
p13        | c. pelaksanaan urusan keprotokolan dan acara perjalanan
p13   F15  | Wakil Presiden dan/atau Istri/Suami Wakil Presiden
p13   F14  | di dalam maupun di luar negeri;
p13        [SUB-ITEM]
p13        | d. koordinasi kegiatan pers dan media, pelayanan
p13   F16  | informasi dan dokumentasi kegiatan Wakil Presiden
p13   F16  | dan/atau Istri/Suami Wakil Presiden, serta acara
p13   F14  | lainnya di lingkungan Sekretariat Wakil Presiden;
p13   F16  I7 [SUB-ITEM]
p13   F16  I7 | e.koordinasi...
p13   F14  I1 | SK No 247653 A
==================== PAGE 14 ====================
p14   F11  | PRESTDEN
p14   F12  | REPUBUK INDONESIA
p14   F18  | -14-
p14        [SUB-ITEM]
p14        | e. koordinasi pelaksanaan tugas pemberian dukungan
p14        | administrasi kepada seluruh unsur organisasi
p14   F14  | di lingkungan Sekretariat Wakil Presiden;
p14        [SUB-ITEM]
p14        | f. pengelolaan dana operasional dan bantuan
p14   F14  | kemasyarakatan Wakil Presiden;
p14        [SUB-ITEM]
p14        | g. pemberian petunjuk teknis di bidang kerumahtanggaan
p14   F14  | dan keprotokolan kepada para Ajudan Wakil presiden
p14   F14  | dan Ajudan Istri/Suami Wakil Presiden;
p14        [SUB-ITEM]
p14        | h. koordinasi Dokter Kepresidenan dalam rangka
p14   F14  | pemberian layanan kesehatan Wakil presiden dan/atau
p14   F14  | Istri/Suami Wakil Presiden; dan
p14        [SUB-ITEM]
p14        | i. pelaksanaan fungsi lain yang diberikan oleh Wakil
p14   F13  | Presiden dan Menteri.
p14   F12  | Pasal 28
p14   F14  | Sekretariat Wakil Presiden terdiri atas:
p14        [SUB-ITEM]
p14        | a. Deputi Bidang Dukungan Kebijakan perekonomian,
p14   F13  | Pariwisata, dan Transformasi Digital;
p14        [SUB-ITEM]
p14        | b. Deputi Bidang Dukungan Kebijakan peningkatan
p14        | Kesejahteraan dan Pembangunan Sumber Daya
p14   F13  | Manusia;
p14        [SUB-ITEM]
p14        | c. Deputi Bidang Dukungan Kebijakan pemerintahan dan
p14   F13  | Pemerataan Pembangunan; dan
p14        [SUB-ITEM]
p14        | d. Deputi Bidang Administrasi.
p14   F13  [HEADING:PARAGRAF]
p14   F13  | Paragraf 1
p14   F14  I4 | Deputi Bidang Dukungan Kebijakan Perekonomian,
p14   F14  | Pariwisata, dan Transformasi Digital
p14   F12  | Pasal 29
p14   F15  [AYAT]
p14   F15  | (1) Deputi Bidang Dukungan Kebijakan perekonomian,
p14   F16  | Pariwisata, dan Transformasi Digital berada di bawah
p14        | dan bertanggung jawab kepada Sekretaris Wakil
p14   F12  | Presiden.
p14   F15  [AYAT]
p14   F15  | (21 Deputi Bidang Dukungan Kebijakan perekonomian,
p14   F16  | Pariwisata, dan Transformasi Digital dipimpin oleh
p14   F13  | Deputi.
p14   F16  | Pasal 30. . .
p14   F14  I1 | SK No 247654A
==================== PAGE 15 ====================
p15   F11  | PRESTDEN
p15   F12  | REPUBUK INDONESIA
p15   F18  | -15-
p15   F12  | Pasal 30
p15   F17  | Deputi Bidang Dukungan Kebijakan perekonomian,
p15   F16  | Pariwisata, dan Transformasi Digital mempunyai tugas
p15        | membantu Sekretaris Wakil presiden dalam
p15   F15  | menyelenggarakan pemberian analisis kebijakan, serta
p15        | dukungan data dan informasi di bidang perekonomian,
p15   F15  | pariwisata, dan transformasi digital kepada wakil presiden
p15        | dalam membantu Presiden menyelenggarakan
p15   F13  | pemerintahan negara.
p15   F12  | Pasal 31
p15   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p15        | Pasal 30, Deputi Bidang Dukungan Kebijakan
p15        | Perekonomian, Pariwisata, dan Transformasi Digital
p15   F13  | menyelenggarakan fungsi:
p15        [SUB-ITEM]
p15        | a. pelaksanaan analisis perkembangan pelaksanaan
p15        | kebijakan pemerintah di bidang perekonomian,
p15   F16  | pariwisata, dan transformasi digital yang ditetapkan
p15   F15  | Presiden atau Wakil Presiden, berikut permasalahan
p15   F14  | yang timbul dan upaya pemecahannya;
p15        [SUB-ITEM]
p15        | b. pengolahan data, informasi, dan penyiapan laporan
p15   F16  | mengenai masalah kebijakan di bidang perekonomian,
p15   F15  | pariwisata, dan transformasi digital yang timbul serta
p15   F13  | dihadapi dalam penyelenggaraan pemerintahan;
p15        [SUB-ITEM]
p15        | c. penyerapan pandangan yang berkembang di kalangan
p15   F14  | pemerintah, lembaga negara, partai politik, organisasi
p15   F17  | profesi, organisasi kemasyarakatan, masyarakat
p15   F14  | akademik, media massa, dan pihak-pihak lainnya yang
p15   F13  | dipandang perlu;
p15        [SUB-ITEM]
p15        | d. penyiapan bahan rapat, pidato/sambutan, notula,
p15   F14  | audiensi, dan kunjungan kerja Wakil presiden; dan
p15        [SUB-ITEM]
p15        | e. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p15   F14  | Wakil Presiden.
p15   F12  | Pasal 32
p15   F16  [AYAT]
p15   F16  | (1) Deputi Bidang Dukungan Kebijakan perekonomian,
p15   F14  | Pariwisata, dan Transformasi Digital terdiri atas paling
p15   F14  | banyak 3 (tiga) asisten deputi.
p15   F14  I1 | SK No 247655 A
p15   F13  [AYAT]
p15   F13  | (2) Asisten .
==================== PAGE 16 ====================
p16   F12  | P[{ESIDEN
p16   F12  | REPUBLTK INDONESIA
p16   F18  | -16-
p16   F15  [AYAT]
p16   F15  | (21 Asisten deputi sebagaimana dimaksud pada ayat (l)
p16   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p16   F16  [AYAT]
p16   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p16   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p16   F18  | dibantu oleh bagian yang melaksanakan fungsi
p16   F14  | administrasi.
p16   F14  [AYAT]
p16   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p16   F14  | jabatan fungsional dan jabatan pelaksana.
p16   F12  [HEADING:PARAGRAF]
p16   F12  | Paragraf 2
p16   F14  | Deputi Bidang Dukungan Kebijakan
p16   F13  I4 | Peningkatan Kesejahteraan dan Pembangunan
p16   F13  | Sumber Daya Manusia
p16   F12  | Pasal 33
p16   F16  [AYAT]
p16   F16  | (1) Deputi Bidang Dukungan Kebijakan peningkatan
p16        | Kesejahteraan dan Pembangunan Sumber Daya
p16        | Manusia berada di bawah dan bertanggung jawab
p16   F14  | kepada Sekretaris Wakil Presiden.
p16   F16  [AYAT]
p16   F16  | (21 Deputi Bidang Dukungan Kebijakan peningkatan
p16        | Kesejahteraan dan Pembangunan Sumber Daya
p16   F14  | Manusia dipimpin oleh Deputi.
p16   F12  | Pasal 34
p16        | Deputi Bidang Dukungan Kebijakan peningkatan
p16   F14  | Kesejahteraan dan Pembangunan Sumber Daya Manusia
p16   F15  | mempunyai tugas membantu Sekretaris Wakil presiden
p16   F15  | dalam menyelenggarakan pemberian analisis kebijakan,
p16   F15  | serta dukungan data dan informasi di bidang peningkatan
p16   F15  | kesejahteraan dan pembangunan sumber daya manusia
p16   F18  | kepada Wakil Presiden dalam membantu presiden
p16   F13  | menyelenggarakan pemerintahan negara.
p16   F12  | Pasal 35
p16   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p16   F14  | Pasal 34, Deputi Bidang Dukungan Kebijakan Peningkatan
p16   F14  | Kesejahteraan dan Pembangunan Sumber Daya Manusia
p16   F12  | menyelenggarakan fungsi :
p16        [SUB-ITEM]
p16        | a. pelaksanaan analisis perkembangan pelaksanaan
p16        | kebijakan pemerintah di bidang peningkatan
p16   F14  | kesejahteraan dan pembangunan sumber daya manusia
p16        | berikut permasalahan yang timbul dan upaya
p16   F12  | pemecahannya;
p16   F15  I7 [SUB-ITEM]
p16   F15  I7 | b.pengolahan...
p16   F14  I1 | SK No 2476564
==================== PAGE 17 ====================
p17   F11  | PRESTDEN
p17   F12  | REPUBUK INDONESIA
p17   F20  | -t7-
p17        [SUB-ITEM]
p17        | b. pengolahan data, informasi, dan laporan mengenai
p17   F13  | masalah kebijakan di bidang peningkatan kesejahteraan
p17   F14  | dan pembangunan sumber daya manusia yang timbul
p17   F14  | serta dihadapi dalam penyelenggaraan pemerintahan;
p17        [SUB-ITEM]
p17        | c. penyerapan pandangan yang berkembang di kalangan
p17   F15  | lembaga negara, organisasi politik, organisasi profesi,
p17   F14  | organisasi kemasyarakatan, masyarakat akademik,
p17   F14  | media massa, dan pihak-pihak lainnya yang dipandang
p17   F13  | perlu;
p17        [SUB-ITEM]
p17        | d. penyiapan bahan rapat, pidato/sambutan, notula,
p17   F14  | audiensi, dan kunjungan kerja Wakil presiden; dan
p17        [SUB-ITEM]
p17        | e. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p17   F14  | Wakil Presiden.
p17   F12  | Pasal 36
p17   F16  [AYAT]
p17   F16  | (1) Deputi Bidang Dukungan Kebijakan peningkatan
p17        | Kesejahteraan dan Pembangunan Sumber Daya
p17   F17  | Manusia terdiri atas paling banyak 4 (empat) asisten
p17   F13  | deputi.
p17   F14  [AYAT]
p17   F14  | (21 Asisten deputi sebagaimana dimaksud pada ayat (1) terdiri
p17   F14  | atas jabatan fungsional dan jabatan pelaksana.
p17   F16  [AYAT]
p17   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p17   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p17   F18  | dibantu oleh bagian yang melaksanakan fungsi
p17   F14  | administrasi.
p17   F14  [AYAT]
p17   F14  | (41 Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p17   F14  | jabatan fungsional dan jabatan pelaksana.
p17   F12  [HEADING:PARAGRAF]
p17   F12  | Paragraf 3
p17   F14  | Deputi Bidang Dukungan Kebijakan
p17   F13  I4 | Pemerintahan dan Pemerataan Pembangunan
p17   F12  | Pasal 37
p17   F14  [AYAT]
p17   F14  | (1) Deputi Bidang Dukungan Kebijakan Pemerintahan dan
p17        | Pemerataan Pembangunan berada di bawah dan
p17   F14  | bertanggung jawab kepada Sekretaris Wakil presiden.
p17   F14  [AYAT]
p17   F14  | (2) Deputi Bidang Dukungan Kebijakan Pemerintahan dan
p17   F14  | Pemerataan Pembangunan dipimpin oleh Deputi.
p17   F14  I1 | SK No 247657 A
p17   F16  | Pasal 38. . .
==================== PAGE 18 ====================
p18   F12  | PRESIDEN
p18   F12  | REPUBUK INDONESIA
p18   F17  | -18-
p18   F12  | Pasal 38
p18   F15  | Deputi Bidang Dukungan Kebijakan Pemerintahan dan
p18   F14  | Pemerataan Pembangunan mempunyai tugas membantu
p18   F18  | Sekretaris Wakil Presiden dalam menyelenggarakan
p18   F15  | pemberian analisis kebijakan, serta dukungan data dan
p18        | informasi di bidang pemerintahan dan pemerataan
p18   F15  | pembangunan kepada Wakil Presiden dalam membantu
p18   F13  | Presiden menyelenggarakan pemerintahan negara.
p18   F12  | Pasal 39
p18   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p18   F14  | Pasal38, Deputi Bidang Dukungan Kebijakan Pemerintahan
p18   F13  | dan Pemerataan Pembangunan menyelenggarakan fungsi :
p18        [SUB-ITEM]
p18        | a. pelaksanaan analisis perkembangan pelaksanaan
p18        | kebijakan pemerintah di bidang pemerintahan dan
p18   F14  | pemerataan pembangunan berikut permasalahan yang
p18   F14  | timbul dan upaya pemecahannya;
p18        [SUB-ITEM]
p18        | b. pengolahan data, informasi, dan laporan mengenai
p18        | masalah kebijakan di bidang pemerintahan dan
p18   F14  | pemerataan pembangunan yang timbul serta dihadapi
p18   F13  | dalam penyelenggaraan pemerintahan;
p18        [SUB-ITEM]
p18        | c. penyerapan pandangan yang berkembang di kalangan
p18   F14  | pemerintah, lembaga negara, partai politik, organisasi
p18   F17  | profesi, organisasi kemasyarakatan, masyarakat
p18   F14  | akademik, media massa, dan pihak-pihak lainnya yang
p18   F13  | dipandang perlu;
p18        [SUB-ITEM]
p18        | d. penyiapan bahan rapat, pidato/sambutan, notula,
p18   F14  | audiensi, dan kunjungan kerja Wakil Presiden; dan
p18        [SUB-ITEM]
p18        | e. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p18   F14  | Wakil Presiden.
p18   F12  | Pasal 40
p18   F14  [AYAT]
p18   F14  | (1) Deputi Bidang Dukungan Kebijakan Pemerintahan dan
p18   F15  | Pemerataan Pembangunan terdiri atas paling banyak
p18   F14  | 3 (tiga) asisten deputi.
p18   F15  [AYAT]
p18   F15  | (2) Asisten deputi sebagaimana dimaksud pada ayat (1)
p18   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p18   F14  I1 | SK No 247658 A
p18   F14  [AYAT]
p18   F14  | (3) Untuk
==================== PAGE 19 ====================
p19   F11  | PRESTDEN
p19   F12  | REPUBUK INDONESIA
p19   F18  | -19-
p19   F16  [AYAT]
p19   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p19   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p19   F18  | dibantu oleh bagian yang melaksanakan fungsi
p19   F14  | administrasi.
p19   F14  [AYAT]
p19   F14  | (41 Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p19   F14  | jabatan fungsional dan jabatan pelaksana.
p19   F12  [HEADING:PARAGRAF]
p19   F12  | Paragraf 4
p19   F13  | Deputi Bidang Administrasi
p19   F12  | Pasal 4 1
p19        [AYAT]
p19        | (1) Deputi Bidang Administrasi berada di bawah dan
p19   F14  | bertanggung jawab kepada Sekretaris Wakil Presiden.
p19   F14  [AYAT]
p19   F14  | (21 Deputi Bidang Administrasi dipimpin oleh Deputi.
p19   F12  | Pasal 42
p19   F14  | Deputi Bidang Administrasi mempunyai tugas membantu
p19   F15  | Sekretaris Wakil Presiden dalam memberikan pelayanan
p19   F14  | kerumahtanggaan, keprotokolan, pers, dan media kepada
p19   F16  | Wakil Presiden dan/atau Istri/Suami Wakil Presiden,
p19   F17  | koordinasi pelaksanaan tugas pemberian dukungan
p19   F14  | administrasi kepada seluruh unsur organisasi di lingkungan
p19   F17  | Sekretariat Wakil Presiden, serta pengelolaan dana
p19   F16  | operasional dan dana bantuan kemasyarakatan Wakil
p19   F12  | Presiden.
p19   F12  | Pasal 43
p19   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p19   F15  | Pasal 42, Deputi Bidang Administrasi menyelenggarakan
p19   F13  | fungsi:
p19        [SUB-ITEM]
p19        | a. penyusunan rencana, program, anggaran, dan
p19   F17  | akuntabilitas kinerja di lingkungan Sekretariat Wakil
p19   F12  | Presiden;
p19        [SUB-ITEM]
p19        | b. perencanaan dan pelaksanaan penyelenggaraan
p19        | keprotokolan kegiatan Wakil Presiden dan/atau
p19   F14  | Istri/ Suami Wakil Presiden;
p19        [SUB-ITEM]
p19        | c. perencanaan dan
p19   F18  I5 | pelaksanaan dukungan
p19   F14  | kerumahtanggaan Wakil Presiden dan/atau Istri/ Suami
p19   F14  | Wakil Presiden;
p19        [SUB-ITEM]
p19        | d. perencanaan dan penyelenggaraan kegiatan pers dan
p19   F15  | media kegiatan Wakil Presiden dan/atau Istri/Suami
p19   F14  | Wakil Presiden;
p19   F14  I1 | SK No 247659 A
p19   F15  I7 [SUB-ITEM]
p19   F15  I7 | e.pemberian...
==================== PAGE 20 ====================
p20   F12  | PRESIDEN
p20   F12  | REPUBUK INDONESIA
p20   F17  | -20-
p20        [SUB-ITEM]
p20        | e. pemberian dukungan prasarar,a dan sarana serta hak
p20   F14  | keuangan bagi Staf Khusus Wakil Presiden;
p20        [SUB-ITEM]
p20        | f. pengelolaan dana operasional dan bantuan
p20   F14  | kemasyarakatan Wakil Presiden;
p20        [SUB-ITEM]
p20        | g. pemberian dukungan administrasi yang meliputi
p20   F16  | ketatausahaan, sumber daya manusia, keuangan,
p20   F15  | kerumahtanggaan, organisasi, tata laksana, reformasi
p20        | birokrasi, perpustakaan, arsip, dan dokumentasi
p20   F14  | di lingkungan Sekretariat Wakil Presiden;
p20        [SUB-ITEM]
p20        | h. penyelenggaraan pengelolaan barang milik/kekayaan
p20   F14  | negara di lingkungan Sekretariat Wakil Presiden; dan
p20        [SUB-ITEM]
p20        | i. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p20   F14  | Wakil Presiden.
p20   F12  | Pasal 44
p20   F14  [AYAT]
p20   F14  | (1) Deputi Bidang Administrasi terdiri atas paling banyak
p20   F13  | 5 (lima) biro.
p20   F15  [AYAT]
p20   F15  | (2) Biro sebagaimana dimaksud pada ayat (1) terdiri atas
p20   F14  | jabatan fungsional dan jabatan pelaksana.
p20        [AYAT]
p20        | (3) Dalam hal tugas dan fungsi biro tidak dapat
p20   F16  | dilaksanakan oleh jabatan fungsional sebagaimana
p20   F14  | dimaksud pada ayat (2), dapat dibentuk paling banyak
p20   F14  | 4 (empat) bagian.
p20   F14  [AYAT]
p20   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p20   F15  | jabatan fungsional dan jabatan pelaksana dan/atau
p20   F13  | paling banyak 3 (tiga) subbagian.
p20   F16  [AYAT]
p20   F16  | (5) Dikecualikan dari ketentuan sebagaimana dimaksud
p20   F14  | pada ayat (41, bagian yang menangani ketatausahaan
p20   F14  | pimpinan terdiri atas sejumlah subbagian sesuai dengan
p20   F13  | kebutuhan.
p20   F12  [HEADING:BAGIAN]
p20   F12  | Bagian Kelima
p20   F14  | Sekretariat Militer Presiden
p20   F12  | Pasal 45
p20        [AYAT]
p20        | (1) Sekretariat Militer Presiden berada di bawah dan
p20   F13  | bertanggung jawab kepada Menteri.
p20   F16  [AYAT]
p20   F16  | (2) Sekretariat Militer Presiden dipimpin oleh Sekretaris
p20   F14  | Militer Presiden.
p20   F14  I1 | SK No 247660 A
p20   F13  I7 [AYAT]
p20   F13  I7 | (3) Sekretaris
==================== PAGE 21 ====================
p21   F12  | PlrESIDEN
p21   F12  | REPUBUK INDONESIA
p21   F20  | -2r-
p21        [AYAT]
p21        | (3) Sekretaris Militer Presiden karena jabatannya
p21   F14  | melaksanakan tugas sebagai Sekretaris Dewan Gelar,
p21   F13  | Tanda Jasa, dan Tanda Kehormatan.
p21   F16  [AYAT]
p21   F16  | (4) Dalam melaksanakan tugasnya, Sekretaris Militer
p21   F15  | Presiden dapat menerima penugasan langsung dari
p21   F12  | Presiden.
p21   F12  | Pasal 46
p21        | Sekretariat Militer Presiden mempunyai tugas
p21   F18  | menyelenggarakan pemberian dukungan teknis dan
p21   F16  | administrasi kepada Presiden dalam menyelenggarakan
p21   F14  | kekuasaan tertinggi atas Angkatan Darat, Angkatan Laut,
p21        | dan Angkatan Udara, dalam hat pengangkatan dan
p21   F14  | pemberhentian perwira Tentara Nasional Indonesia dan
p21   F14  | Kepolisian Negara Republik Indonesia, penganugerahan
p21   F13  | gelar, tanda jasa, dan tanda kehormatan yang wewenangnya
p21   F16  | berada pada Presiden, serta koordinasi pengamanan
p21   F15  | Presiden dan Wakil Presiden beserta keluarga termasuk
p21   F13  | Tamu Negara setingkat Kepala Negara/Kepala Pemerintahan
p21   F12  | negara asing.
p21   F12  | Pasal 47
p21   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p21   F16  | Pasal 46, Sekretariat Militer Presiden menyelenggarakan
p21   F13  | fungsi:
p21        [SUB-ITEM]
p21        | a. pemberian dukungan teknis dan administrasi personil
p21   F16  | Tentara Nasional Indonesia dan Kepolisian Negara
p21        | Republik Indonesia yang berkaitan dengan
p21   F14  | pengangkatan atau pemberhentian dalam jabatan serta
p21   F14  | kepangkatan perwira Tentara Nasional Indonesia dan
p21   F18  | Kepolisian Negara Republik Indonesia serta
p21        | pengangkatan atau pemberhentian dari dinas
p21   F14  | keprajuritan yang wewenang penetapannya berada pada
p21   F12  | Presiden;
p21        [SUB-ITEM]
p21        | b. koordinasi penyelenggaraan pengamanan fisik dan
p21   F16  | non fisik bagi Presiden dan Wakil Presiden beserta
p21   F16  | keluarga, termasuk Tamu Negara setingkat Kepala
p21   F13  | Negara/Kepala Pemerintahan negara asing;
p21        [SUB-ITEM]
p21        | c. pelaksanaan kegiatan teknis dan administrasi
p21   F14  | penganugerahan gelar pahlawan, tanda jasa, dan tanda
p21   F13  | kehormatan yang wewenang penetapannya berada pada
p21   F12  | Presiden;
p21   F14  I1 | SK No 247991 A
p21   F14  I7 [SUB-ITEM]
p21   F14  I7 | d. pelaksanaan. . .
==================== PAGE 22 ====================
p22   F12  | PRESIDEN
p22   F12  | REPUEUK TNDONESIA
p22   F17  | -22-
p22        [SUB-ITEM]
p22        | d. pelaksanaan koordinasi dengan instansi terkait
p22   F18  | mengenai penganugerahan tanda jasa dan tanda
p22   F16  | kehormatan secara imbal batik antara Pemerintah
p22   F13  | Republik Indonesia dengan Pemerintah negara asing;
p22        [SUB-ITEM]
p22        | e. pembinaan personil dan pemberian petunjuk teknis
p22   F14  | di bidang pengamanan kepada Ajudan Presiden, Ajudan
p22   F15  | Wakil Presiden, Ajudan Istri/Suami Presiden, Ajudan
p22   F14  | Istri/Suami Wakil Presiden, Ajudan Tamu Negara Asing,
p22   F14  | Dokter Pribadi Presiden, Dokter Pribadi Wakil presiden,
p22        | Sekretaris Kabinet, Staf Khusus Presiden dan
p22   F14  | Staf Khusus Wakil Presiden, serta pembinaan prajurit
p22   F16  | Tentara Nasional Indonesia dan anggota Kepolisian
p22   F14  | Negara Republik Indonesia yang bertugas di lingkungan
p22   F13  | Kementerian;
p22        [SUB-ITEM]
p22        | f. koordinasi penjadwalan agenda kegiatan presiden dan
p22   F14  | pertemuan yang dipimpin oleh Presiden;
p22        [SUB-ITEM]
p22        | g. pemberian dukungan administrasi kepada seluruh
p22        | unsur organisasi di lingkungan Sekretariat Militer
p22   F13  | Presiden; dan
p22        [SUB-ITEM]
p22        | h. pelaksanaan fungsi lain yang diberikan oleh Presiden
p22   F13  | dan Menteri.
p22   F12  | Pasal 48
p22   F15  [AYAT]
p22   F15  | (1) Sekretariat Militer Presiden terdiri atas paling banyak
p22   F14  | 4 (empat) biro dan Sekretaris Kabinet.
p22   F14  [AYAT]
p22   F14  | (2) Biro sebagaimana dimaksud pada ayat (1) terdiri atas
p22   F14  | jabatan fungsional dan jabatan pelaksana.
p22        [AYAT]
p22        | (3) Dalam hal tugas dan fungsi biro tidak dapat
p22   F16  | dilaksanakan oleh jabatan fungsional sebagaimana
p22   F14  | dimaksud pada ayat (2), dapat dibentuk pating banyak
p22   F14  | 4 (empat) bagian.
p22   F14  [AYAT]
p22   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p22   F15  | jabatan fungsional dan jabatan pelaksana dan/atau
p22   F14  | paling banyak 3 (tiga) subbagian.
p22   F13  [HEADING:BAGIAN]
p22   F13  | Bagian Keenam
p22   F14  | Sekretariat Dukungan Kabinet
p22   F12  | Pasal 49
p22   F17  [AYAT]
p22   F17  | (1) Sekretariat Dukungan Kabinet berada di bawah dan
p22   F14  | bertanggung jawab kepada Menteri.
p22   F14  [AYAT]
p22   F14  | (2) Sekretariat Dukungan Kabinet dipimpin oleh Sekretaris
p22   F14  | Dukungan Kabinet.
p22   F14  I1 | SK No 247662 A
p22   F16  | Pasal 50. .
==================== PAGE 23 ====================
p23   F12  | PRESIDEN
p23   F12  | REPUBUK TNDONESTA
p23   F17  | -23-
p23   F12  | Pasal 50
p23   F18  | Sekretariat Dukungan Kabinet mempunyai tugas
p23   F17  | memberikan dukungan manajemen kabinet kepada
p23   F18  | Presiden dan Wakil Presiden dalam penyelenggaraan
p23   F14  | pemerintahan melalui Menteri Sekretaris Negara.
p23   F12  | Pasal 51
p23   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p23   F14  | Pasal 50, Sekretariat Dukungan Kabinet menyelenggarakan
p23   F13  | fungsi:
p23        [SUB-ITEM]
p23        | a. pengkajian dan pemberian rekomendasi atas rencana
p23   F13  | kebijakan dan program pemerintah;
p23        [SUB-ITEM]
p23        | b. penyelesaian masalah atas pelaksanaan kebijakan dan
p23   F13  | program pemerintah yang mengalami hambatan;
p23        [SUB-ITEM]
p23        | c. penyampaian rekomendasi atas hasil pengamatan dan
p23   F14  | penyerapan pandangan terhadap perkembangan umum;
p23        [SUB-ITEM]
p23        | d. pemantauan dan evaluasi atas pelaksanaan kebijakan
p23   F13  | pemerintah;
p23        [SUB-ITEM]
p23        | e. penyiapan, pengadministrasian, penyelenggaraan, dan
p23   F13  | pengelolaan sidang kabinet, rapat, atau pertemuan yang
p23   F16  | dipimpin dan/atau dihadiri oleh Presiden danf atau
p23   F17  | Wakil Presiden, penyiapan naskah bagi Presiden
p23   F14  | dan/atau Wakil Presiden, pelaksanaan penerjemahan,
p23   F14  | keprotokolan dalam sidang kabinet, serta pengelolaan
p23   F14  | arsip dan dokumentasi Kepresidenan dan Kementerian;
p23        [SUB-ITEM]
p23        | f. pemberian dukungan teknis dan administrasi dalam
p23        | pengangkatan, pemindahan, dan pemberhentian
p23   F14  | Jabatan Pimpinan Tinggi Utama, Jabatan Pimpinan
p23   F14  | Tinggi Madya, dan Jabatan lainnya kepada Tim Penilai
p23   F14  | Akhir;
p23        [SUB-ITEM]
p23        | g. koordinasi pelaksanaan tugas pemberian dukungan
p23   F16  | administrasi, sumber daya manusia, perencanaan,
p23   F17  | keuangan, dan penyediaan prasarana dan sarana
p23   F14  | di lingkungan Sekretariat Dukungan Kabinet; dan
p23        [SUB-ITEM]
p23        | h. pelaksanaan fungsi lain yang diberikan oleh Menteri.
p23   F12  | Pasal 52
p23   F14  | Sekretariat Dukungan Kabinet terdiri atas:
p23        [SUB-ITEM]
p23        | a. Deputi Bidang Politik, Hukum, Keamanan, dan Hak
p23   F13  | Asasi Manusia;
p23   F14  I1 | SK No 247663 A
p23   F14  [SUB-ITEM]
p23   F14  | b. Deputi. . .
==================== PAGE 24 ====================
p24   F12  | PRESIDEN
p24   F12  | REPUBLIK INDONESIA
p24   F17  | -24-
p24        [SUB-ITEM]
p24        | b. Deputi Bidang Perekonomian;
p24        [SUB-ITEM]
p24        | c. Deputi Bidang Pembangunan Manusia, Kebudayaan,
p24   F13  | dan Pemberdayaan Masyarakat;
p24        [SUB-ITEM]
p24        | d. Deputi Bidang Pangan, Infrastruktur, dan
p24   F13  | Pembangunan Kewilayahan;
p24        [SUB-ITEM]
p24        | e. Deputi Bidang Persidangan Kabinet; dan
p24        [SUB-ITEM]
p24        | f. Deputi Bidang Administrasi.
p24   F13  [HEADING:PARAGRAF]
p24   F13  | Paragraf 1
p24   F14  | Deputi Bidang Politik, Hukum,
p24   F13  | Keamanan, dan Hak Asasi Manusia
p24   F12  | Pasal 53
p24   F16  [AYAT]
p24   F16  | (1) Deputi Bidang Politik, Hukum, Keamanan, dan Hak
p24   F14  | Asasi Manusia berada di bawah dan bertanggung jawab
p24   F14  | kepada Sekretaris Dukungan Kabinet.
p24   F16  [AYAT]
p24   F16  | (2) Deputi Bidang Politik, Hukum, Keamanan, dan Hak
p24   F14  | Asasi Manusia dipimpin oleh Deputi.
p24   F12  | Pasal 54
p24   F14  | Deputi Bidang Politik, Hukum, Keamanan, dan Hak Asasi
p24        | Manusia mempunyai tugas membantu Sekretaris
p24   F14  | Dukungan Kabinet dalam menyelenggarakan pemberian
p24   F17  | dukungan manajemen kabinet di bidang politik, hukum,
p24   F14  | keamanan, dan hak asasi manusia.
p24   F12  | Pasal 55
p24   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p24   F14  | Pasal 54, Deputi Bidang Politik, Hukum, Keamanan, dan
p24   F13  | Hak Asasi Manusia menyelenggarakan fungsi:
p24        [SUB-ITEM]
p24        | a. pengkajian dan pemberian rekomendasi atas rencana
p24   F16  | kebijakan dan program pemerintah di bidang politik,
p24   F14  | hukum, keamanan, dan hak asasi manusia;
p24        [SUB-ITEM]
p24        | b. penyelesaian masalah atas pelaksanaan kebijakan dan
p24        | program pemerintah di bidang politik, hukum,
p24   F16  | keamanan, dan hak asasi manusia yang mengalami
p24   F13  | hambatan;
p24   F14  I1 | SK No 247664 A
p24   F13  I7 [SUB-ITEM]
p24   F13  I7 | c. penyampalan
==================== PAGE 25 ====================
p25   F12  | PRESIDEN
p25   F12  | REPUBUK INDONESIA
p25   F17  | -25-
p25        [SUB-ITEM]
p25        | c. penyampaian rekomendasi atas hasil pengamatan dan
p25   F14  | penyerapan pandangan terhadap perkembangan umum
p25   F17  | di bidang politik, hukum, keamanan, dan hak asasi
p25   F13  | manusia;
p25        [SUB-ITEM]
p25        | d. pemantauan dan evaluasi atas pelaksanaan kebijakan
p25   F16  | pemerintah di bidang politik, hukum, keamanan, dan
p25   F14  | hak asasi manusia;
p25        [SUB-ITEM]
p25        | e. penyiapan bahan substansi sidang kabinet, rapat, atau
p25   F16  | pertemuan yang dipimpin dan/atau dihadiri oleh
p25   F18  | Presiden dan/atau Wakil Presiden di bidang politik,
p25   F14  | hukum, keamanan, dan hak asasi manusia; dan
p25        [SUB-ITEM]
p25        | f. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p25   F14  | Dukungan Kabinet.
p25   F12  | Pasal 56
p25   F16  [AYAT]
p25   F16  | (1) Deputi Bidang Politik, Hukum, Keamanan, dan Hak
p25        | Asasi Manusia terdiri atas paling banyak 4 (empat)
p25   F13  | asisten deputi.
p25   F15  [AYAT]
p25   F15  | (2) Asisten deputi sebagaimana dimaksud pada ayat (1)
p25   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p25   F16  [AYAT]
p25   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p25   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p25   F18  | dibantu oleh bagian yang melaksanakan fungsi
p25   F14  | administrasi.
p25   F14  [AYAT]
p25   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p25   F14  | jabatan fungsional dan jabatan pelaksana.
p25   F12  [HEADING:PARAGRAF]
p25   F12  | Paragraf 2
p25   F14  | Deputi Bidang Perekonomian
p25   F12  | Pasal 57
p25        [AYAT]
p25        | (1) Deputi Bidang Perekonomian berada di bawah dan
p25   F17  | bertanggung jawab kepada Sekretaris Dukungan
p25   F13  | Kabinet.
p25   F14  [AYAT]
p25   F14  | (2) Deputi Bidang Perekonomian dipimpin oleh Deputi.
p25   F12  | Pasal 58
p25   F14  | Deputi Bidang Perekonomian mempunyai tugas membantu
p25   F15  | Sekretaris Dukungan Kabinet dalam menyelenggarakan
p25        | pemberian dukungan manajemen kabinet di bidang
p25   F13  | perekonomian.
p25   F14  I1 | SK No 247665 A
p25   F16  | Pasal59...
==================== PAGE 26 ====================
p26   F12  | PRESIDEN
p26   F12  | REPUBUK INDONESIA
p26   F17  | -26-
p26   F12  | Pasal 59
p26   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p26   F14  | Pasal 58, Deputi Bidang Perekonomian menyelenggarakan
p26   F13  | fungsi:
p26        [SUB-ITEM]
p26        | a. pengkajian dan pemberian rekomendasi atas rencana
p26        | kebijakan dan program pemerintah di bidang
p26   F13  | perekonomian;
p26        [SUB-ITEM]
p26        | b. penyelesaian masalah atas pelaksanaan kebijakan dan
p26        | program pemerintah di bidang perekonomian yang
p26   F13  | mengalami hambatan;
p26        [SUB-ITEM]
p26        | c. penyampaian rekomendasi atas hasil pengamatan dan
p26   F14  | penyerapan pandangan terhadap perkembangan umum
p26   F14  | di bidang perekonomian;
p26        [SUB-ITEM]
p26        | d. pemantauan dan evaluasi atas pelaksanaan kebijakan
p26   F14  | pemerintah di bidang perekonomian;
p26        [SUB-ITEM]
p26        | e. penyiapan bahan substansi sidang kabinet, rapat, atau
p26   F16  | pertemuan yang dipimpin dan/atau dihadiri oleh
p26        | Presiden dan/atau Wakil Presiden di bidang
p26   F13  | perekonomian; dan
p26        [SUB-ITEM]
p26        | f. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p26   F14  | Dukungan Kabinet.
p26   F12  | Pasal 60
p26   F14  [AYAT]
p26   F14  | (1) Deputi Bidang Perekonomian terdiri atas paling banyak
p26   F14  | 4 (empat) asisten deputi.
p26   F15  [AYAT]
p26   F15  | (2) Asisten deputi sebagaimana dimaksud pada ayat (1)
p26   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p26   F16  [AYAT]
p26   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p26   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p26   F18  | dibantu oleh bagian yang melaksanakan fungsi
p26   F14  | administrasi.
p26   F14  [AYAT]
p26   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p26   F14  | jabatan fungsional dan jabatan pelaksana.
p26   F12  [HEADING:PARAGRAF]
p26   F12  | Paragraf 3
p26   F14  | Deputi Bidang Pembangunan Manusia,
p26   F13  | Kebudayaan, dan Pemberdayaan Masyarakat
p26   F12  | Pasal 61
p26   F14  [AYAT]
p26   F14  | (1) Deputi Bidang Pembangunan Manusia, Kebudayaan,
p26   F16  | dan Pemberdayaan Masyarakat berada di bawah dan
p26   F17  | bertanggung jawab kepada Sekretaris Dukungan
p26   F13  | Kabinet.
p26   F16  [AYAT]
p26   F16  | (2) Deputi...
p26   F14  I1 | SK No 2476664
==================== PAGE 27 ====================
p27   F12  | PRESIDEN
p27   F13  | REFUBUK INDONESIA
p27   F16  | -27 -
p27   F14  [AYAT]
p27   F14  | (2) Deputi Bidang Pembangunan Manusia, Kebudayaan,
p27   F14  | dan Pemberdayaan Masyarakat dipimpin oleh Deputi.
p27   F12  | Pasal 62
p27   F14  | Deputi Bidang Pembangunan Manusia, Kebudayaan, dan
p27   F14  | Pemberdayaan Masyarakat mempunyai tugas membantu
p27   F15  | Sekretaris Dukungan Kabinet dalam menyelenggarakan
p27        | pemberian dukungan manajemen kabinet di bidang
p27   F14  | pembangunan manusia, kebudayaan, dan pemberdayaan
p27   F13  | masyarakat.
p27   F12  | Pasal 63
p27   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p27        | Pasal 62, Deputi Bidang Pembangunan Manusia,
p27        | Kebudayaan, dan
p27   F16  I5 | Pemberdayaan Masyarakat
p27   F12  | menyelenggarakan fungsi :
p27        [SUB-ITEM]
p27        | a. pengkajian dan pemberian rekomendasi atas rencana
p27        | kebijakan dan program pemerintah di bidang
p27        | pembangunan manusia, kebudayaan, dan
p27   F13  | pemberdayaan masyarakat;
p27        [SUB-ITEM]
p27        | b. penyelesaian masalah atas pelaksanaan kebijakan dan
p27   F14  | program pemerintah di bidang pembangunan manusia,
p27   F17  | kebudayaan, dan pemberdayaan masyarakat yang
p27   F13  | mengalami hambatan;
p27        [SUB-ITEM]
p27        | c. penyampaian rekomendasi atas hasil pengamatan dan
p27   F14  | penyerapan pandangan terhadap perkembangan umum
p27   F18  | di bidang pembangunan manusia, kebudayaan, dan
p27   F13  | pemberdayaan masyarakat;
p27        [SUB-ITEM]
p27        | d. pemantauan dan evaluasi atas pelaksanaan kebijakan
p27        | pemerintah di bidang pembangunan manusia,
p27   F13  | kebudayaan, dan pemberdayaan masyarakat;
p27        [SUB-ITEM]
p27        | e. penyiapan bahan substansi sidang kabinet, rapat, atau
p27   F16  | pertemuan yang dipimpin dan/atau dihadiri oleh
p27        | Presiden dan/atau Wakil Presiden di bidang
p27        | pembangunan manusia, kebudayaan, dan
p27   F13  | pemberdayaan masyarakat; dan
p27        [SUB-ITEM]
p27        | f. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p27   F14  | Dukungan Kabinet.
p27   F14  I1 | SK No 247667 A
p27   F12  | Pasal 64
==================== PAGE 28 ====================
p28   F12  | PRESIDEN
p28   F12  | REPUBLIK INDONESIA
p28   F17  | -28-
p28   F12  | Pasal 64
p28   F14  [AYAT]
p28   F14  | (1) Deputi Bidang Pembangunan Manusia, Kebudayaan,
p28   F17  | dan Pemberdayaan Masyarakat terdiri atas paling
p28   F14  | banyak 4 (empat) asisten deputi.
p28   F15  [AYAT]
p28   F15  | (2) Asisten deputi sebagaimana dimaksud pada ayat (1)
p28   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p28   F16  [AYAT]
p28   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p28   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p28   F18  | dibantu oleh bagian yang melaksanakan fungsi
p28   F14  | administrasi.
p28   F14  [AYAT]
p28   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p28   F14  | jabatan fungsional dan jabatan pelaksana.
p28   F12  [HEADING:PARAGRAF]
p28   F12  | Paragraf 4
p28   F14  | Deputi Bidang Pangan, Infrastruktur, dan
p28   F13  | Pembangunan Kewilayahan
p28   F12  | Pasal 65
p28        [AYAT]
p28        | (1) Deputi Bidang Pangan, Infrastruktur, dan
p28        | Pembangunan Kewilayahan berada di bawah dan
p28   F17  | bertanggung jawab kepada Sekretaris Dukungan
p28   F12  | Kabinet.
p28        [AYAT]
p28        | (2) Deputi Bidang Pangan, Infrastruktur, dan
p28   F14  | Pembangunan Kewilayahan dipimpin oleh Deputi.
p28   F12  | Pasal 66
p28   F15  | Deputi Bidang Pangan, Infrastruktur, dan Pembangunan
p28   F16  | Kewilayahan mempunyai tugas membantu Sekretaris
p28   F14  | Dukungan Kabinet dalam menyelenggarakan pemberian
p28        | dukungan manajemen kabinet di bidang pangan,
p28   F14  | infrastruktur, dan pembangunan kewilayahan.
p28   F12  | Pasal 67
p28   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p28   F18  | Pasal 66, Deputi Bidang Pangan, Infrastruktur, dan
p28   F13  | Pembangunan Kewilayahan menyelenggarakan fungsi :
p28        [SUB-ITEM]
p28        | a. pengkajian dan pemberian rekomendasi atas rencana
p28   F16  | kebijakan dan program pemerintah di bidang pangan,
p28   F14  | infrastruktur, dan pembangunan kewilayahan;
p28        [SUB-ITEM]
p28        | b. penyelesaian masalah atas pelaksanaan kebijakan dan
p28   F16  | program pemerintah di bidang pangan, infrastruktur,
p28        | dan pembangunan kewilayahan yang mengalami
p28   F13  | hambatan;
p28   F14  I1 | SK No 2476684
p28   F17  I7 [SUB-ITEM]
p28   F17  I7 | c. penyampalan .
==================== PAGE 29 ====================
p29   F12  | PRESIDEN
p29   F12  | REPUBUK TNDONESTA
p29   F17  | -29-
p29        [SUB-ITEM]
p29        | c. penyampaian rekomendasi atas hasil pengamatan dan
p29   F14  | penyerapan pandangan terhadap perkembangan umum
p29        | di bidang pangan, infrastruktur, dan pembangunan
p29   F13  | kewilayahan;
p29        [SUB-ITEM]
p29        | d. pemantauan dan evaluasi atas pelaksanaan kebijakan
p29        | pemerintah di bidang pangan, infrastruktur, dan
p29   F13  | pembangunan kewilayahan;
p29        [SUB-ITEM]
p29        | e. penyiapan bahan substansi sidang kabinet, rapat, atau
p29   F16  | pertemuan yang dipimpin dan/atau dihadiri oleh
p29   F16  | Presiden dan/atau Wakil Presiden di bidang pangan,
p29   F14  | infrastruktur, dan pembangunan kewilayahan; dan
p29        [SUB-ITEM]
p29        | f. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p29   F14  | Dukungan Kabinet.
p29   F12  | Pasal 68
p29        [AYAT]
p29        | (1) Deputi Bidang Pangan, Infrastruktur, dan
p29   F15  | Pembangunan Kewilayahan terdiri atas paling banyak
p29   F14  | 4 (empat) asisten deputi.
p29   F15  [AYAT]
p29   F15  | (2) Asisten deputi sebagaimana dimaksud pada ayat (1)
p29   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p29   F16  [AYAT]
p29   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p29   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p29   F18  | dibantu oleh bagian yang melaksanakan fungsi
p29   F14  | administrasi.
p29   F14  [AYAT]
p29   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p29   F14  | jabatan fungsional dan jabatan pelaksana.
p29   F12  [HEADING:PARAGRAF]
p29   F12  | Paragraf 5
p29   F14  | Deputi Bidang Persidangan Kabinet
p29   F12  | Pasal 69
p29   F16  [AYAT]
p29   F16  | (1) Deputi Bidang Persidangan Kabinet berada di bawah
p29   F14  | dan bertanggung jawab kepada Sekretaris Dukungan
p29   F12  | Kabinet.
p29   F16  [AYAT]
p29   F16  | (2) Deputi Bidang Persidangan Kabinet dipimpin oleh
p29   F13  | Deputi.
p29   F16  | Pasal70...
p29   F14  I1 | SK No 247669 A
==================== PAGE 30 ====================
p30   F11  | PRESTDEN
p30   F12  | REPUBUK INDONESIA
p30   F18  | -30-
p30   F12  | Pasal 70
p30   F16  | Deputi Bidang Persidangan Kabinet mempunyai tugas
p30        | membantu Sekretaris Dukungan Kabinet dalam
p30   F16  | menyelenggarakan pemberian dukungan manajemen
p30        | kabinet dalam hal penyiapan, pengadministrasian,
p30   F17  | penjadwalan, penyelenggaraan dan pengelolaan sidang
p30   F14  | kabinet, rapat, atau pertemuan yang dipimpin dan/atau
p30   F14  | dihadiri oleh Presiden dan latau Wakil Presiden, penyiapan
p30        | naskah bagi Presiden dan/atau Wakil Presiden,
p30   F15  | pelaksanaan penerjemahan, keprotokolan dalam sidang
p30        | kabinet, serta pengelolaan arsip dan dokumentasi
p30   F14  | Kepresidenan dan Kementerian.
p30   F12  | Pasal 71
p30   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p30        | Pasal 70, Deputi Bidang Persidangan Kabinet
p30   F12  | menyelenggarakan fungsi :
p30        [SUB-ITEM]
p30        | a. penyelenggaraan urusan administrasi, penjadwalan,
p30   F15  | dan pengelolaan agenda sidang kabinet, rapat, atau
p30   F15  | pertemuan yang dipimpin dan/atau dihadiri Presiden
p30   F14  | dan/atau Wakil Presiden;
p30        [SUB-ITEM]
p30        | b. pen5rusunan risalah dan
p30   F13  I7 | pendokumentasian,
p30   F16  | pendistribusian, dan publikasi hasil sidang kabinet,
p30   F14  | rapat, atau pertemuan yang dipimpin dan/atau dihadiri
p30   F14  | oleh Presiden dan latau Wakil Presiden;
p30        [SUB-ITEM]
p30        | c. penyelenggaraan urusan pendokumentasian hal-hal
p30   F14  | yang berkaitan dengan pelaksanaan sidang kabinet,
p30   F14  | rapat, atau pertemuan yang dipimpin dan/atau dihadiri
p30   F14  | oleh Presiden dan latau Wakil Presiden;
p30        [SUB-ITEM]
p30        | d. pengoordinasian penyiapan naskah dokumen
p30   F13  | Kepresidenan dan Kementerian;
p30        [SUB-ITEM]
p30        | e. pelaksanaan penerjemahan bagi Presiden dan/atau
p30   F14  | Wakil Presiden, serta di lingkungan Kementerian;
p30        [SUB-ITEM]
p30        | f. penyelenggaraan keprotokolan dalam sidang kabinet;
p30        [SUB-ITEM]
p30        | g. pengelolaan arsip dan dokumentasi Kepresidenan dan
p30   F13  | Kementerian; dan
p30        [SUB-ITEM]
p30        | h. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p30   F14  | Dukungan Kabinet.
p30   F16  | Pasal72...
p30   F14  I1 | SK No 247670 A
==================== PAGE 31 ====================
p31   F12  | PRESIDEN
p31   F12  | REPUBUK INDONESIA
p31   F14  | - 31 -
p31   F12  | Pasal 72
p31   F14  [AYAT]
p31   F14  | (1) Deputi Bidang Persidangan Kabinet terdiri atas paling
p31   F14  | banyak 4 (empat) asisten deputi.
p31   F15  [AYAT]
p31   F15  | (2) Asisten deputi sebagaimana dimaksud pada ayat (1)
p31   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p31   F16  [AYAT]
p31   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p31   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p31   F18  | dibantu oleh bagian yang melaksanakan fungsi
p31   F14  | administrasi.
p31   F14  [AYAT]
p31   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p31   F14  | jabatan fungsional dan jabatan pelaksana.
p31   F12  [HEADING:PARAGRAF]
p31   F12  | Paragraf 6
p31   F14  | Deputi Bidang Administrasi
p31   F12  | Pasal 73
p31        [AYAT]
p31        | (1) Deputi Bidang Administrasi berada di bawah dan
p31   F17  | bertanggung jawab kepada Sekretaris Dukungan
p31   F12  | Kabinet.
p31   F14  [AYAT]
p31   F14  | (2) Deputi Bidang Administrasi dipimpin oleh Deputi.
p31   F12  | Pasal T4
p31   F14  | Deputi Bidang Administrasi mempunyai tugas membantu
p31        | Sekretaris Dukungan Kabinet dalam koordinasi
p31   F16  | pelaksanaan tugas pemberian dukungan administrasi
p31   F14  | di lingkungan Sekretariat Dukungan Kabinet.
p31   F12  | Pasal 75
p31   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p31   F15  | Pasal 74, Deputi Bidang Administrasi menyelenggarakan
p31   F13  | fungsi:
p31        [SUB-ITEM]
p31        | a. pen5rusunan rencana, program, anggaran, dan
p31        | akuntabilitas kinerja di lingkungan Sekretariat
p31   F14  | Dukungan Kabinet;
p31        [SUB-ITEM]
p31        | b. pemberian dukungan administrasi yang meliputi
p31   F16  | ketatausahaan, sumber daya manusia, keuangan,
p31   F15  | kerumahtanggaan, organisasi, tata laksana, reformasi
p31        | birokrasi, arsip, dan dokumentasi di lingkungan
p31   F14  | Sekretariat Dukungan Kabinet;
p31   F14  I1 | SK No 247671 A
p31   F11  I7 | c
p31   F14  I7 | penyediaan.
==================== PAGE 32 ====================
p32   F12  | PRESIDEN
p32   F12  | REPUBUK INDONESIA
p32   F18  | -32-
p32        [SUB-ITEM]
p32        | c. penyediaan prasarana dan sarana, pemeliharaan,
p32   F15  | perawatan dan pengelolaan barang milik negara, serta
p32        | penyelenggaraan pelayanan dan
p32   F14  | administrasi
p32        | pengadaan di lingkungan Sekretariat Dukungan
p32   F13  | Kabinet;
p32        [SUB-ITEM]
p32        | d. pemberian dukungan teknis dan administrasi dalam
p32        | pengangkatan, pemindahan, dan pemberhentian
p32   F14  | Jabatan Pimpinan Tinggi Utama, Jabatan Pimpinan
p32   F14  | Tinggi Madya, dan Jabatan lainnya kepada Tim Penilai
p32   F14  | Akhir; dan
p32        [SUB-ITEM]
p32        | e. pelaksanaan fungsi lain yang diberikan oleh Sekretaris
p32   F14  | Dukungan Kabinet.
p32   F12  | Pasal 76
p32   F14  [AYAT]
p32   F14  | (1) Deputi Bidang Administrasi terdiri atas paling banyak
p32   F14  | 3 (tiga) biro.
p32   F15  [AYAT]
p32   F15  | (2) Biro sebagaimana dimaksud pada ayat (1) terdiri atas
p32   F14  | jabatan fungsional dan jabatan pelaksana.
p32        [AYAT]
p32        | (3) Dalam hal tugas dan fungsi biro tidak dapat
p32   F16  | dilaksanakan oleh jabatan fungsional sebagaimana
p32   F14  | dimaksud pada ayat (21, dapat dibentuk paling banyak
p32   F14  | 4 (empat) bagian.
p32   F14  | (a) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p32        | jabatan fungsional dan/atau paling banyak 3 (tiga)
p32   F12  | subbagian.
p32   F16  [AYAT]
p32   F16  | (5) Dikecualikan dari ketentuan sebagaimana dimaksud
p32   F14  | pada ayat (41, bagian yang menangani ketatausahaan
p32   F17  | pimpinan terdiri atas sejumlah subbagian sesuai
p32   F14  | dengan kebutuhan.
p32   F14  [HEADING:BAGIAN]
p32   F14  | Bagian Ketujuh
p32   F14  | Deputi Bidang Perundang-undangan
p32   F14  | dan Administrasi Hukum
p32   F12  | Pasal TT
p32   F14  [AYAT]
p32   F14  | (1) Deputi Bidang Perundang-undangan dan Administrasi
p32   F14  | Hukum berada di bawah dan bertanggung jawab kepada
p32   F13  | Menteri.
p32   F14  [AYAT]
p32   F14  | (2) Deputi Bidang Pemndang-undangan dan Administrasi
p32   F14  | Hukum dipimpin oleh Deputi.
p32   F14  I1 | SK No 247672 A
p32   F12  | Pasal 78
==================== PAGE 33 ====================
p33   F11  | PRESTDEN
p33   F12  | REPUBUK INDONESIA
p33   F17  | -33-
p33   F12  | Pasal 78
p33   F17  | Deputi Bidang Perundang-undangan dan Administrasi
p33   F15  | Hukum mempunyai tugas menyelenggarakan pemberian
p33        | dukungan teknis, administrasi, dan analisis dalam
p33        | penyiapan izin prakarsa Rancangan Peraturan
p33   F14  | Perundang-undangan, penyelesaian Rancangan Peraturan
p33   F13  | Perundang-undangan, Rancangan Keputusan Presiden, dan
p33        | Rancangan Instruksi Presiden, dan pengundangan
p33   F17  | Undang-Undang, Peraturan Pemerintah Pengganti
p33   F16  | Undang-Undang, Peraturan Pemerintah, dan Peraturan
p33   F18  | Presiden, serta penyelesaian dan penanganan terkait
p33   F17  | dengan litigasi, permasalahan hukum, penyelesaian
p33   F14  | Rancangan Keputusan Presiden mengenai grasi, amnesti,
p33   F14  | abolisi, rehabilitasi, perubahan pidana mati atau perubahan
p33   F14  | pidana penjara seumur hidup, kewarganegaraan Republik
p33   F16  | Indonesia, ekstradisi, dan keanggotaan Indonesia pada
p33   F13  | organisasi internasional.
p33   F12  | Pasal 79
p33   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p33        | Pasal 78, Deputi Bidang Perundang-undangan dan
p33   F14  | Administrasi Hukum menyelenggarakan fungsi:
p33        [SUB-ITEM]
p33        | a. pelaksanaan analisis dalam penyiapanizinprakarsa
p33   F14  | penJrusunan Rancangan Undang-Undang, Rancangan
p33   F16  | Peraturan Pemerintah Pengganti Undang-Undang,
p33   F17  | Rancangan Peraturan Pemerintah, dan Rancangan
p33   F13  | Peraturan Presiden;
p33        [SUB-ITEM]
p33        | b. pemantauan, analisis, dan pelaporan pen5rusunan
p33   F14  | Rancangan Undang-Undang, Rancangan Peraturan
p33   F14  | Pemerintah Pengganti Undang-Undang, Rancangan
p33   F14  | Peraturan Pemerintah, Rancangan Peraturan Presiden,
p33        | Rancangan Keputusan Presiden, dan Rancangan
p33   F14  | Instruksi Presiden;
p33        [SUB-ITEM]
p33        | c. pelaksanaan analisis dalam
p33   F12  | penyelesaian
p33   F14  | Rancangan Undang-Undang, Rancangan Peraturan
p33   F14  | Pemerintah Pengganti Undang-Undang, Rancangan
p33   F14  | Peraturan Pemerintah, Rancangan Peraturan Presiden,
p33        | Rancangan Keputusan Presiden, dan Rancangan
p33   F14  | Instruksi Presiden;
p33   F14  I1 | SK No 247673 A
p33   F13  I7 [SUB-ITEM]
p33   F13  I7 | d. pelaksanaan
==================== PAGE 34 ====================
p34   F12  | PRESIDEN
p34   F12  | REPUBLIK INDONESIA
p34   F17  | -34-
p34        [SUB-ITEM]
p34        | d. pelaksanaan analisis, penyelesaian, dan penyiapan
p34   F16  | Rancangan Keputusan Presiden mengenai grasi,
p34   F14  | amnesti, abolisi, rehabilitasi, perubahan pidana mati
p34   F15  | atau perubahan pidana penjara seumur hidup, dan
p34   F13  | kewarganegaraan Republik Indonesia, dan ekstradisi;
p34        [SUB-ITEM]
p34        | e. pelaksanaan analisis dan penyelesaian permasalahan
p34        | di bidang perjanjian internasional dan keanggotaan
p34   F14  | Indonesia pada organisasi internasional;
p34        [SUB-ITEM]
p34        | f. pelaksanaan litigasi, analisis dan pen5rusunan pendapat
p34   F16  | hukum terhadap gugatan perdata dan tata usaha
p34   F14  | negara, serta gugatan arbitrase internasional kepada
p34   F16  | Presiden dan Wakil Presiden, permohonan uji materiil
p34   F14  | peraturan perundang-undangan, serta permasalahan
p34   F14  | hukum lainnya;
p34        [SUB-ITEM]
p34        | g. pengundangan Undang-Undang, Peraturan Pemerintah
p34   F14  | Pengganti Undang-Undang, Peraturan Pemerintah, dan
p34   F13  | Peraturan Presiden;
p34        [SUB-ITEM]
p34        | h. pemberian nomor, pendistribusian, publikasi, dan
p34   F17  | pendokumentasian Undang-Undang, Peraturan
p34   F15  | Pemerintah Pengganti Undang-Undang, Peraturan
p34   F14  | Pemerintah, Peraturan Presiden, Keputusan Presiden,
p34   F14  | dan Instruksi Presiden; dan
p34        [SUB-ITEM]
p34        | i. pelaksanaan fungsi lain yang diberikan oleh Menteri.
p34   F12  | Pasal 80
p34   F14  [AYAT]
p34   F14  | (1) Deputi Bidang Perundang-undangan dan Administrasi
p34        | Hukum terdiri atas paling banyak 5 (lima) asisten
p34   F13  | deputi.
p34   F15  [AYAT]
p34   F15  | (2) Asisten deputi sebagaimana dimaksud pada ayat (1)
p34   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p34   F16  [AYAT]
p34   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p34   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p34   F18  | dibantu oleh bagian yang melaksanakan fungsi
p34   F14  | administrasi.
p34   F14  [AYAT]
p34   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p34   F14  | jabatan fungsional dan jabatan pelaksana.
p34   F14  I1 | SK No 247891 A
p34   F12  | Bagian
==================== PAGE 35 ====================
p35   F12  | PRESIDEN
p35   F12  | REPUBUK INDONESTA
p35   F17  | -35-
p35   F13  [HEADING:BAGIAN]
p35   F13  | Bagian Kedelapan
p35   F14  | Deputi Bidang Hubungan Kelembagaan
p35   F13  | dan Kemasyarakatan
p35   F12  | Pasal 81
p35        [AYAT]
p35        | (1) Deputi Bidang Hubungan Kelembagaan dan
p35        | Kemasyarakatan berada di bawah dan bertanggung
p35   F14  | jawab kepada Menteri.
p35        [AYAT]
p35        | (2) Deputi Bidang Hubungan Kelembagaan dan
p35   F14  | Kemasyarakatan dipimpin oleh Deputi.
p35   F12  | Pasal 82
p35        | Deputi Bidang Hubungan Kelembagaan dan
p35   F16  | Kemasyarakatan mempunyai tugas menyelenggarakan
p35   F16  | pemberian dukungan teknis, administrasi, dan analisis
p35   F14  | dalam penyelenggaraan hubungan dengan lembaga negara,
p35   F16  | lembaga nonstruktural, lembaga daerah, organisasi
p35        | kemasyarakatan, organisasi politik, dan penanganan
p35   F14  | pengaduan masyarakat kepada Presiden, Wakil Presiden,
p35   F14  | dan/atau Menteri, serta penyelenggaraan kemitraan.
p35   F12  | Pasal 83
p35   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p35   F16  | Pasal 82, Deputi Bidang Hubungan Kelembagaan dan
p35   F13  | Kemasyarakatan menyelenggarakan fungsi:
p35        [SUB-ITEM]
p35        | a. penyiapan dan analisis data dan informasi dalam rangka
p35   F17  | mendukung penyelenggaraan hubungan antara
p35   F16  | Presiden dan/atau Wakil Presiden dengan lembaga
p35   F16  | negara, lembaga nonstruktural, lembaga daerah,
p35   F14  | organisasi kemasyarakatan, dan organisasi politik;
p35        [SUB-ITEM]
p35        | b. penyampaian saran dan masukan kepada Menteri
p35   F14  | dalam rangka mendukung penyelenggaraan hubungan
p35   F16  | antara Presiden dan/atau Wakil Presiden dengan
p35   F16  | lembaga negara, lembaga nonstruktural, lembaga
p35   F16  | daerah, organisasi kemasyarakatan, dan organisasi
p35   F14  | politik;
p35        [SUB-ITEM]
p35        | c. pemantauan secara aktif dinamika kegiatan lembaga
p35   F16  | negara, lembaga nonstruktural, lembaga daerah,
p35   F18  | organisasi kemasyarakatan, dan organisasi politik
p35   F18  | dalam rangka pemberian dukungan hubungan
p35   F13  | kelembagaan kepada Presiden danf atau Wakil Presiden;
p35   F14  I1 | SK No 247675 A
p35   F14  I7 [SUB-ITEM]
p35   F14  I7 | d. koordinasi
==================== PAGE 36 ====================
p36   F12  | PRES]DEN
p36   F12  | REPUBUK INDONESIA
p36   F18  | -36-
p36        [SUB-ITEM]
p36        | d. koordinasi pelaksanaan hubungan kelembagaan antara
p36   F16  | Presiden dan/atau Wakil Presiden dengan lembaga
p36   F16  | negara, lembaga nonstruktural, lembaga daerah,
p36   F14  | organisasi kemasyarakatan, dan organisasi politik;
p36        [SUB-ITEM]
p36        | e. penanganan pengaduan masyarakat yang disampaikan
p36   F14  | kepada Presiden, Wakil Presiden, dan/atau Menteri;
p36        [SUB-ITEM]
p36        | f. penyelenggaraan kemitraan kementerian dengan
p36   F14  | lembaga pemerintah, lembaga non pemerintah, badan
p36   F14  | hukum dan badan usaha, serta pihak swasta; dan
p36        [SUB-ITEM]
p36        | g. pelaksanaan fungsi lain yang diberikan oleh Menteri.
p36   F12  | Pasal 84
p36        [AYAT]
p36        | (1) Deputi Bidang Hubungan Kelembagaan dan
p36   F18  | Kemasyarakatan terdiri atas paling banyak 4 (empat)
p36   F13  | asisten deputi.
p36   F15  [AYAT]
p36   F15  | (2) Asisten deputi sebagaimana dimaksud pada ayat (1)
p36   F14  | terdiri atas jabatan fungsional dan jabatan pelaksana.
p36   F16  [AYAT]
p36   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p36   F14  | pelaksanaan tugas dan fungsinya asisten deputi dapat
p36   F18  | dibantu oleh bagian yang melaksanakan fungsi
p36   F14  | administrasi.
p36   F14  | (a) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p36   F14  | jabatan fungsional dan jabatan pelaksana.
p36   F13  [HEADING:BAGIAN]
p36   F13  | Bagian Kesembilan
p36   F14  | Deputi Bidang Administrasi Aparatur
p36   F12  | Pasal 85
p36   F15  [AYAT]
p36   F15  | (1) Deputi Bidang Administrasi Aparatur berada di bawah
p36   F14  | dan bertanggung jawab kepada Menteri.
p36   F15  [AYAT]
p36   F15  | (2) Deputi Bidang Administrasi Aparatur dipimpin oleh
p36   F13  | Deputi.
p36   F12  | Pasal 86
p36   F14  | Deputi Bidang Administrasi Aparatur mempunyai tugas
p36        | menyelenggarakan pemberian dukungan teknis,
p36        | administrasi, dan analisis dalam pengangkatan,
p36   F17  | pemberhentian, dan pensiun pejabat negara, pejabat
p36   F14  | pemerintahan, pejabat lainnya, dan Aparatur Sipil Negara
p36   F16  | yang wewenang penetapannya berada pada Presiden,
p36   F14  | pembinaan, penataan, dan pengembangan Aparatur Sipil
p36   F14  | Negara, organisasi, tata laksana, reformasi birokrasi, serta
p36   F14  | koordinasi pen5rusunan peraturan perundang-undangan,
p36        | serta pelaksanaan advokasi hukum dan litigasi di
p36   F14  | lingkungan Kementerian.
p36   F16  | Pasal87...
p36   F14  I1 | SK No 247676A
==================== PAGE 37 ====================
p37   F12  | PRESIDEN
p37   F12  | REPUBUK INDONESIA
p37   F16  | -37 -
p37   F12  | Pasal 87
p37   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p37        | Pasal 86, Deputi Bidang Administrasi Aparatur
p37   F13  | menyelenggarakan fungsi:
p37        [SUB-ITEM]
p37        | a. pemberian dukungan teknis dan administrasi
p37   F16  | pengangkatan, pemberhentian, dan pensiun pejabat
p37   F13  | negara, pejabat pemerintahan, dan pejabat lainnya yang
p37   F13  | wewenang penetapannya berada pada Presiden;
p37        [SUB-ITEM]
p37        | b. pemberian dukungan teknis dan administrasi
p37   F14  | pengangkatan, pemberhentian, dan pensiun Aparatur
p37   F14  | Sipil Negara yang wewenang penetapannya berada pada
p37   F12  | Presiden;
p37        [SUB-ITEM]
p37        | c. pembinaan dan pemberian dukungan administrasi
p37   F14  | sumber daya manusia di lingkungan Kementerian;
p37        [SUB-ITEM]
p37        | d. pen5rusunan peraturan perundang-undangan serta
p37   F14  | pelaksanaan advokasi hukum dan litigasi di lingkungan
p37   F13  | Kementerian;
p37        [SUB-ITEM]
p37        | e. pembinaan dan penataan organisasi, tata laksana, dan
p37   F14  | reformasi birokrasi di lingkungan Kementerian; dan
p37        [SUB-ITEM]
p37        | f. pelaksanaan fungsi lain yang diberikan oleh Menteri.
p37   F12  | Pasal 88
p37   F14  [AYAT]
p37   F14  | (1) Deputi Bidang Administrasi Aparatur terdiri atas paling
p37   F14  | banyak 4 (empat) biro.
p37   F15  [AYAT]
p37   F15  | (2) Biro sebagaimana dimaksud pada ayat (1) terdiri atas
p37   F14  | jabatan fungsional dan jabatan pelaksana.
p37        [AYAT]
p37        | (3) Dalam hal tugas dan fungsi biro tidak dapat
p37   F16  | dilaksanakan oleh jabatan fungsional sebagaimana
p37   F14  | dimaksud pada ayat (21, dapat dibentuk 2 (dua) bagian.
p37   F14  [AYAT]
p37   F14  | (4) Bagian sebagaimana dimaksud pada ayat (3) terdiri atas
p37   F14  | jabatan fungsional dan jabatan pelaksana.
p37   F13  [HEADING:BAGIAN]
p37   F13  | Bagian Kesepuluh
p37   F14  | Badan Teknologi, Data, dan Informasi
p37   F12  | Pasal 89
p37   F14  [AYAT]
p37   F14  | (1) Badan Teknologi, Data, dan Informasi berada di bawah
p37   F14  | dan bertanggung jawab kepada Menteri.
p37   F15  [AYAT]
p37   F15  | (2) Badan Teknologi, Data, dan Informasi dipimpin oleh
p37   F12  | Kepala Badan.
p37   F16  | Pasal90...
p37   F14  I1 | SK No 247677 A
==================== PAGE 38 ====================
p38   F12  | PRESIDEN
p38   F12  | REPUBUK INDONESTA
p38   F17  | -38-
p38   F12  | Pasal 9O
p38   F15  | Badan Teknologi, Data, dan Informasi mempunyai tugas
p38   F15  | menyelenggarakan pengelolaan, pengembangan, dan
p38   F14  | pemeliharaan sistem teknologi informasi dan komunikasi,
p38   F16  | infrastruktur dan jaringan komunikasi dan data, serta
p38   F14  | keamanan data dan informasi di lingkungan Kementerian.
p38   F12  | Pasal 91
p38   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p38        | Pasal 90, Badan Teknologi, Data, dan Informasi
p38   F13  | menyelenggarakan fungsi:
p38        [SUB-ITEM]
p38        | a. pengelolaan, pengembangan, dan pemeliharaan sistem
p38        | teknologi informasi dan komunikasi di lingkungan
p38   F12  | Kementerian;
p38        [SUB-ITEM]
p38        | b. pengelolaan, pengembangan, dan pemeliharaan
p38        | infrastruktur dan jaringan komunikasi dan data
p38   F14  | di lingkungan Kementerian;
p38        [SUB-ITEM]
p38        | c. pengelolaan, pengembangan, dan pemeliharaan
p38        | keamanan data dan informasi di lingkungan
p38   F13  | Kementerian'
p38        [SUB-ITEM]
p38        | d. pemberia., artrrrgan data dan informasi di lingkungan
p38   F13  | Kementerian; dan
p38        [SUB-ITEM]
p38        | e. pelaksanaan fungsi lain yang diberikan oleh Menteri.
p38   F12  | Pasal 92
p38   F14  [AYAT]
p38   F14  | (1) Badan Teknologi, Data, dan Informasi terdiri atas paling
p38   F14  | banyak 3 (tiga) pusat.
p38   F14  [AYAT]
p38   F14  | (2) Pusat sebagaimana dimaksud pada ayat (1) terdiri atas
p38   F14  | jabatan fungsional dan jabatan pelaksana.
p38   F16  [AYAT]
p38   F16  | (3) Untuk memberikan dukungan administrasi, dalam
p38   F14  | pelaksanaan tugas dan fungsinya pusat dapat dibantu
p38        | oleh subbagian yang melaksanakan fungsi
p38   F13  | ketatausahaan.
p38   F13  [HEADING:BAGIAN]
p38   F13  | Bagian Kesebelas
p38   F14  | Staf Ahli
p38   F12  | Pasal 93
p38   F16  [AYAT]
p38   F16  | (1) Staf Ahli Bidang Politik, Pertahanan, dan Keamanan
p38   F14  | mempunyai tugas memberikan rekomendasi terhadap
p38   F14  | isu-isu strategis kepada Menteri terkait dengan bidang
p38   F14  | politik, pertahanan, dan keamanan.
p38   F14  [AYAT]
p38   F14  | (2) Staf Ahli Bidang Ekonomi, Kemaritiman, Pembangunan
p38        | Manusia, dan Kebudayaan mempunyai tugas
p38   F16  | memberikan rekomendasi terhadap isu-isu strategis
p38   F17  | kepada Menteri terkait dengan bidang ekonomi,
p38   F14  | kemaritiman, pembangunan manusia, dan kebudayaan.
p38        [AYAT]
p38        | (3) Staf ...
p38   F14  I1 | SK No 247678A
==================== PAGE 39 ====================
p39   F12  | PRESIDEN
p39   F12  | REPUBUK INDONESIA
p39   F17  | -39-
p39   F16  [AYAT]
p39   F16  | (3) Staf Ahli Bidang Hukum, Hak Asasi Manusia, dan
p39        | Pemerintahan mempunyai tugas memberikan
p39   F14  | rekomendasi terhadap isu-isu strategis kepada Menteri
p39   F14  | terkait dengan bidang hukum, hak asasi manusia, dan
p39   F13  | pemerintahan.
p39   F18  | (a) Staf Ahli Bidang Aparatur Negara dan Reformasi
p39   F14  | Birokrasi mempunyai tugas memberikan rekomendasi
p39   F16  | terhadap isu-isu strategis kepada Menteri terkait
p39   F14  | dengan bidang aparatur negara, reformasi birokrasi, dan
p39   F14  | transformasi digital.
p39   F16  [AYAT]
p39   F16  | (5) Staf Ahli Bidang Komunikasi Politik dan Kehumasan
p39   F14  | mempunyai tugas memberikan rekomendasi terhadap
p39   F14  | isu-isu strategis kepada Menteri terkait dengan bidang
p39   F14  | komunikasi politik dan kehumasan.
p39   F13  [HEADING:BAGIAN]
p39   F13  | Bagian Keduabelas
p39   F14  | Inspektorat
p39   F12  | Pasal 94
p39   F16  [AYAT]
p39   F16  | (1) Inspektorat berada di bawah dan bertanggung jawab
p39   F14  | kepada Menteri melalui Sekretaris Kementerian.
p39   F14  [AYAT]
p39   F14  | (2) Inspektorat dipimpin oleh Inspektur.
p39   F12  | Pasal 95
p39        | Inspektorat mempunyai tugas menyelenggarakan
p39   F14  | pengawasan intern di lingkungan Kementerian.
p39   F12  | Pasal 96
p39   F14  | Dalam melaksanakan tugas sebagaimana dimaksud dalam
p39   F14  | Pasal 95, Inspektorat menyelenggarakan fungsi:
p39        [SUB-ITEM]
p39        | a. penJrusunan kebijakan teknis pengawasan intern
p39   F14  | di lingkungan Kementerian;
p39        [SUB-ITEM]
p39        | b. pelaksanaan pengawasan internal terhadap kinerja dan
p39   F15  | keuangan melalui audit, reviu, evaluasi, pemantauan
p39        | dan kegiatan pengawasan lainnya di lingkungan
p39   F13  | Kementerian;
p39        [SUB-ITEM]
p39        | c. pelaksanaan pengawasan untuk tujuan tertentu atas
p39   F13  | penugasan Menteri;
p39        [SUB-ITEM]
p39        | d. penJrusunan laporan hasil pengawasan di lingkungan
p39   F13  | Kementerian'
p39   F13  I7 [SUB-ITEM]
p39   F13  I7 | e. pelaksanaan .
p39   F14  I1 | SK No 247679 A
==================== PAGE 40 ====================
p40   F12  | PRESIDEN
p40   F12  | REPUBUK INDONESIA
p40   F17  | -40-
p40        [SUB-ITEM]
p40        | e. pelaksanaan administrasi Inspektorat; dan
p40        [SUB-ITEM]
p40        | f. pelaksanaan fungsi lain yang diberikan oleh Menteri.
p40   F12  | Pasal 97
p40   F14  [AYAT]
p40   F14  | (1) Inspektorat terdiri atas jabatan fungsional dan jabatan
p40   F12  | pelaksana.
p40   F16  [AYAT]
p40   F16  | (2) Dalam hal tugas dan fungsi Inspektorat tidak dapat
p40   F16  | dilaksanakan oleh jabatan fungsional sebagaimana
p40   F14  | dimaksud pada ayat (1), dapat dibentuk 1 (satu) bagian.
p40   F14  [AYAT]
p40   F14  | (3) Bagian sebagaimana dimaksud pada ayat (21terdiri atas
p40   F14  | jabatan fungsional dan jabatan pelaksana.
p40   F13  [HEADING:BAGIAN]
p40   F13  | Bagian Ketigabelas
p40   F12  | Pusat
p40   F12  | Pasal 98
p40        [AYAT]
p40        | (1) Pusat dapat dibentuk di lingkungan Kementerian
p40   F16  | sebagai unsur pendukung pelaksanaan tugas dan
p40   F13  | fungsi.
p40   F16  [AYAT]
p40   F16  | (2) Pusat sebagaimana dimaksud pada ayat (1) diatur
p40        | dengan Peraturan Menteri setelah mendapat
p40   F14  | persetujuan menteri yang menyelenggarakan urusan
p40   F14  | pemerintahan di bidang aparatur negara.
p40   F12  | Pasal 99
p40   F14  [AYAT]
p40   F14  | (1) Penentuan jumlah Pusat sebagaimana dimaksud dalam
p40   F13  | Pasal 98 didasarkan pada analisis organisasi dan beban
p40   F13  | kerja.
p40   F14  [AYAT]
p40   F14  | (2) Pusat sebagaimana dimaksud pada ayat (1) terdiri atas
p40   F15  | jabatan fungsional dan jabatan pelaksana dan/atau
p40   F16  | dapat terdiri atas paling banyak 3 (tiga) bidang, serta
p40   F14  [HEADING:BAGIAN]
p40   F14  | bagian yang menangani fungsi ketatausahaan.
p40   F14  [AYAT]
p40   F14  | (3) Bidang sebagaimana dimaksud pada ayat (21terdiri atas
p40   F14  | jabatan fungsional dan jabatan pelaksana.
p40   F14  [AYAT]
p40   F14  | (4) Bagian sebagaimana dimaksud pada ayat (2) terdiri atas
p40   F14  | jabatan fungsional dan jabatan pelaksana.
p40        [AYAT]
p40        | (5) Dalam hal tugas dan fungsi bagian sebagaimana
p40   F14  | dimaksud pada ayat (41 tidak dapat dilaksanakan oleh
p40   F16  | jabatan fungsional, dapat dibentuk paling banyak
p40   F14  | 2 (dua) subbagian.
p40   F13  [HEADING:BAGIAN]
p40   F13  | Bagian . . .
p40   F14  I2 | SK No 247680 A
==================== PAGE 41 ====================
p41   F12  | PRESIDEN
p41   F12  | REPUBUK INDONESIA
p41   F15  | -41 -
p41   F13  [HEADING:BAGIAN]
p41   F13  | Bagian Keempatbelas
p41   F14  | Jabatan Fungsional dan Jabatan Pelaksana
p41   F13  | Pasal 100
p41   F14  | Jabatan fungsional dan jabatan pelaksana dapat ditetapkan
p41   F14  | di lingkungan Kementerian sesuai dengan kebutuhan yang
p41   F16  | pelaksanaannya sesuai dengan ketentuan peraturan
p41   F13  | perundang-undangan.
p41   F12  [HEADING:BAB]
p41   F12  | BAB IV
p41   F12  | STAF KHUSUS MENTERI
p41   F13  | Pasal 101
p41        [AYAT]
p41        | (1) Di lingkungan Kementerian dapat diangkat paling
p41   F14  | banyak 5 (lima) orang Staf Khusus.
p41   F14  [AYAT]
p41   F14  | (2) Staf Khusus bertanggung jawab kepada Menteri.
p41   F13  | Pasal 102
p41   F14  [AYAT]
p41   F14  | (1) Staf Khusus mempunyai tugas memberikan saran dan
p41   F16  | pertimbangan kepada Menteri sesuai penugasan
p41   F13  | Menteri.
p41   F17  [AYAT]
p41   F17  | (2) Penugasan sebagaimana dimaksud pada ayat (1)
p41   F16  | merupakan penugasan yang bersifat khusus selain
p41   F14  | bidang tugas unsur-unsur organisasi Kementerian.
p41   F13  | Pasal 103
p41   F16  [AYAT]
p41   F16  | (1) Staf Khusus dapat berasal dari pegawai negeri sipil
p41   F14  | dan/atau non-pegawai negeri sipil.
p41   F18  [AYAT]
p41   F18  | (2) Pegawai negeri sipil sebagaimana dimaksud pada
p41   F14  | ayat (1) diberhentikan dari jabatan organiknya tanpa
p41   F14  | kehilangan statusnya sebagai pegawai negeri sipil sesuai
p41   F14  | dengan ketentuan peraturan perundang-undangan.
p41   F14  [AYAT]
p41   F14  | (3) Masa bakti Staf Khusus paling lama sama dengan masa
p41   F14  | jabatan Menteri.
p41        [AYAT]
p41        | (4) Pengangkatan Staf Khusus ditetapkan dengan
p41   F16  | Keputusan Menteri setelah mendapat persetujuan
p41   F12  | Presiden.
p41   F16  | PasallO4...
p41   F14  I1 | SK No 247681 A
==================== PAGE 42 ====================
p42   F12  | PRESIDEN
p42   F12  | REPUBUK INDONESTA
p42   F17  | -42-
p42   F13  | Pasal 104
p42   F17  [AYAT]
p42   F17  | (1) Pegawai negeri sipil sebagaimana dimaksud dalam
p42   F14  | Pasal 103 ayat (1) yang berhenti atau telah berakhir
p42   F14  | masa baktinya sebagai Staf Khusus, diangkat dalam
p42        | jabatan organik sesuai formasi yang tersedia
p42   F18  | berdasarkan ketentuan peraturan perundang-
p42   F13  | undangan.
p42   F17  [AYAT]
p42   F17  | (2) Pegawai negeri sipil sebagaimana dimaksud dalam
p42   F15  | Pasal 103 ayat (1) yang telah mencapai batas usia
p42   F15  | pensiun diberhentikan dengan hormat dan diberikan
p42        | hak kepegawaiannya sesuai dengan ketentuan
p42   F14  | peraturan perundang-undangan.
p42   F13  | Pasal 105
p42   F14  [AYAT]
p42   F14  | (1) Hak keuangan dan fasilitas lainnya bagi Staf Khusus
p42   F14  | diberikan paling tinggi setara dengan jabatan pimpinan
p42   F14  | tinggi madya atau jabatan struktural eselon I.b.
p42   F15  [AYAT]
p42   F15  | (2) Staf Khusus mendapat dukungan administrasi dari
p42   F13  | Sekretariat Kementerian.
p42   F16  [AYAT]
p42   F16  | (3) Dalam hal Staf Khusus berhenti atau telah berakhir
p42   F16  | masa baktinya tidak memperoleh uang pensiun dan
p42   F13  | uang pesangon.
p42   F12  [HEADING:BAB]
p42   F12  | BAB V
p42   F12  | TATA KERJA
p42   F13  | Pasal 106
p42   F13  | Menteri dalam memimpin pelaksanaan tugas dan fungsinya
p42   F16  | menerapkan sistem akuntabilitas kinerja pemerintah,
p42   F14  | manajemen risiko pembangunan nasional, dan transformasi
p42   F14  | digital nasional.
p42   F13  | Pasal 107
p42   F14  [AYAT]
p42   F14  | (1) Dalam mendukung optimalisasi pelaksanaan tugas dan
p42        | fungsi secara terpadu antarunit organisasi
p42        | di lingkungan Kementerian perlu didasarkan pada
p42   F16  | proses bisnis yang menggambarkan tata hubungan
p42   F16  | kerja yang efektif dan efisien dengan menerapkan
p42        | prinsip koordinasi, integrasi, sinkronisasi, dan
p42        | kolaborasi antarunit organisasi di
p42   F14  | lingkungan
p42        | Kementerian'
p42   F13  [AYAT]
p42   F13  | (2) Proses . . .
p42   F14  I1 | SK No 247682 A
==================== PAGE 43 ====================
p43   F12  | PRESIDEN
p43   F12  | REPUBUK INDONESIA
p43   F17  | -43-
p43        [AYAT]
p43        | (2) Proses bisnis antarunit organisasi di lingkungan
p43   F16  | Kementerian sebagaimana dimaksud pada ayat (1)
p43   F14  | diatur dengan Peraturan Menteri.
p43   F13  | Pasal 108
p43   F14  | Menteri menyampaikan laporan kepada Presiden mengenai
p43        | hasil pelaksanaan urusan pemerintahan di bidang
p43   F14  | kesekretariatan negara secara berkala dan sewaktu-waktu
p43   F14  | sesuai kebutuhan.
p43   F13  | Pasal 109
p43   F15  | Kementerian men5rusun analisis jabatan, peta jabatan,
p43   F14  | analisis beban kerja, dan uraian tugas terhadap seluruh
p43   F14  | jabatan di lingkungan Kementerian.
p43   F13  | Pasal 1 10
p43        [AYAT]
p43        | (1) Setiap unsur di lingkungan Kementerian dalam
p43   F15  | melaksanakan tugas dan fungsi menerapkan prinsip
p43        | koordinasi, integrasi, sinkronisasi, dan kolaborasi
p43        | di lingkungan Kementerian, hubungan antarinstansi
p43   F14  | pemerintah, dan dengan lembaga lain yang terkait.
p43   F17  [AYAT]
p43   F17  | (2) Prinsip koordinasi, integrasi, sinkronisasi, dan
p43   F18  | kolaborasi sebagaimana dimaksud pada ayat (1)
p43   F14  | didukung dengan melakukan interoperabilitas data dan
p43   F13  | informasi.
p43   F13  | Pasal 1 1 1
p43        | Semua unsur di lingkungan Kementerian menerapkan
p43   F16  | sistem pengendalian intern pemerintah sesuai dengan
p43   F14  | ketentuan peraturan perundang-undangan.
p43   F13  | Pasal 1 12
p43   F17  [AYAT]
p43   F17  | (1) Setiap pimpinan unit organisasi bertanggung jawab
p43        | memimpin dan mengoordinasikan bawahan dan
p43        | memberikan pengarahan serta petunjuk bagi
p43   F15  | pelaksanaan tugas sesuai dengan uraian tugas yang
p43   F13  | telah ditetapkan.
p43   F14  I1 | SK No 247683 A
p43   F14  I7 [AYAT]
p43   F14  I7 | (2) Pengarahan. . .
==================== PAGE 44 ====================
p44   F12  | PRESIDEN
p44   F12  | REPUBUK INDONESIA
p44   F17  | -44-
p44   F18  [AYAT]
p44   F18  | (2) Pengarahan dan petunjuk sebagaimana dimaksud
p44   F17  | pada ayat (1) diikuti dan dipatuhi oleh bawahan
p44        | secara bertanggung jawab serta dilaporkan
p44   F15  | secara berkala sesuai dengan ketentuan peraturan
p44   F13  | perundang-undangan.
p44   F13  | Pasal 1 13
p44   F17  | Dalam melaksanakan tugas, setiap pimpinan unit
p44        | organisasi melakukan pembinaan dan pengawasan
p44   F14  | terhadap unit organisasi di bawahnya.
p44   F12  [HEADING:BAB]
p44   F12  | BAB VI
p44   F12  I4 | PENGELOLAAN SUMBER DAYA DAN PENDANAAN
p44   F13  | Pasal 1 14
p44        | Pembinaan dan pengelolaan sumber daya manusia,
p44   F14  | keuangan, perlengkapan, kearsipan, dokumentasi, dan
p44   F16  | persandian diselenggarakan oleh Kementerian dengan
p44   F16  | menerapkan sistem pemerintahan berbasis elektronik
p44   F14  | dalam rangka mendukung transformasi digital.
p44   F13  | Pasal 1 15
p44        | Pendanaan dalam pelaksanaan tugas dan fungsi
p44   F16  | Kementerian bersumber dari Anggaran Pendapatan dan
p44   F13  | Belanja Negara.
p44   F13  [HEADING:BAB]
p44   F13  | BAB VII
p44   F12  | PENATAAN ORGANISASI
p44   F13  | Pasal 1 16
p44   F14  [AYAT]
p44   F14  | (1) Penataan organisasi Kementerian ditetapkan dengan:
p44        [SUB-ITEM]
p44        | a. Peraturan Presiden atas usul menteri yang
p44   F16  | menyelenggarakan urusan pemerintahan di bidang
p44   F16  | aparatur negara, untuk jabatan pimpinan tinggi
p44   F14  | madya atau jabatan struktural eselon I; dan
p44        [SUB-ITEM]
p44        | b. Peraturan Menteri setelah mendapat persetujuan
p44   F14  | tertulis dari menteri yang menyelenggarakan urusan
p44        | pemerintahan di bidang aparatur negara, untuk
p44   F16  | jabatan pimpinan tinggi pratama atau jabatan
p44   F14  | struktural eselon II ke bawah.
p44   F14  I1 | SK No 247684A.
p44   F12  I7 [AYAT]
p44   F12  I7 | (2) Penataan
==================== PAGE 45 ====================
p45   F12  | PRESIDEN
p45   F12  | REPUBUK TNDONESIA
p45   F17  | -45-
p45   F16  [AYAT]
p45   F16  | (2) Penataan organisasi sebagaimana dimaksud pada
p45   F16  | ayat (1) dilakukan dengan mengacu pada sistem
p45   F17  | akuntabilitas kinerja pemerintah sesuai dengan
p45   F14  | ketentuan peraturan perundang-undangan dan proses
p45   F14  | bisnis antarunit organisasi di lingkungan Kementerian.
p45   F13  | Pasal 1 17
p45        [AYAT]
p45        | (1) Besaran organisasi Kementerian ditentukan
p45   F14  | berdasarkan karakteristik tugas dan fungsi serta beban
p45   F12  | kerja.
p45   F14  [AYAT]
p45   F14  | (2) Besaran organisasi sebagaimana dimaksud pada ayat (1)
p45   F16  | juga mempertimbangkan mandat konstitusi, visi dan
p45   F16  | misi Presiden, tantangan utama bangsa, keterkaitan
p45   F14  | dengan agenda prioritas nasional, asas desentralisasi,
p45   F13  | dan peran pemerintah.
p45   F13  [HEADING:BAB]
p45   F13  | BAB VIII
p45   F12  I4 | JABATAN, PENGANGKATAN, DAN PEMBERHENTIAN
p45   F13  | Pasal 1 18
p45   F14  [AYAT]
p45   F14  | (1) Sekretaris Kementerian, Sekretaris Presiden, Sekretaris
p45   F15  | Wakil Presiden, Sekretaris Militer Presiden, Sekretaris
p45        | Dukungan Kabinet, Deputi, dan Kepala Badan
p45   F14  | merupakan jabatan pimpinan tinggi madya atau jabatan
p45   F14  | struktural eselon I.a.
p45   F16  [AYAT]
p45   F16  | (2) Staf Ahli merupakan jabatan pimpinan tinggi madya
p45   F14  | atau jabatan struktural eselon I.b.
p45   F16  [AYAT]
p45   F16  | (3) Kepala Biro, Asisten Deputi, Inspektur, dan Kepala
p45   F14  | Pusat merupakan jabatan pimpinan tinggi pratama atau
p45   F14  | jabatan struktural eselon II.a.
p45   F17  [AYAT]
p45   F17  | (4) Sekretaris Kabinet setinggi-tingginya merupakan
p45        | jabatan pimpinan tinggi pratama atau jabatan
p45   F14  | struktural eselon II.a.
p45        [AYAT]
p45        | (5) Kepala Istana Kepresidenan setinggi-tingginya
p45   F16  | merupakan jabatan pimpinan tinggi pratama atau
p45   F14  | jabatan struktural eselon II.b.
p45   F14  [AYAT]
p45   F14  | (6) Kepala Bagian dan Kepala Bidang merupakan jabatan
p45   F14  | administrator atau jabatan struktural eselon III.a.
p45   F14  [AYAT]
p45   F14  | (7) Kepala Subbagian merupakan jabatan pengawas atau
p45   F14  | jabatan struktural eselon IV.a.
p45   F14  | Pasal 119. . .
p45   F14  I1 | SK No 247685 A
==================== PAGE 46 ====================
p46   F12  | PRESIDEN
p46   F12  | REPUBUK INDONESIA
p46   F18  | -46-
p46   F13  | Pasal 1 19
p46        [AYAT]
p46        | (1) Pejabat pimpinan tinggi madya atau pejabat
p46        | struktural eselon I diangkat dan diberhentikan
p46        | oleh Presiden atas usul Menteri, setelah melalui
p46   F15  | prosedur seleksi berdasarkan ketentuan peraturan
p46   F13  | perundang-undangan.
p46   F14  [AYAT]
p46   F14  | (2) Pejabat pimpinan tinggi pratama atau pejabat struktural
p46        | eselon II diangkat dan diberhentikan oleh Menteri,
p46   F14  | setelah melalui prosedur seleksi berdasarkan ketentuan
p46   F14  | peraturan perundang-undangan.
p46   F14  [AYAT]
p46   F14  | (3) Pejabat administrator atau pejabat struktural eselon III
p46   F14  | ke bawah diangkat dan diberhentikan oleh Menteri.
p46   F14  | (a) Pejabat administrator atau pejabat struktural eselon III
p46   F14  | ke bawah dapat diangkat dan diberhentikan oleh pejabat
p46   F14  | yang diberi pelimpahan wewenang oleh Menteri.
p46   F14  [AYAT]
p46   F14  | (5) Pejabat fungsional diangkat dan diberhentikan sesuai
p46   F14  | dengan ketentuan peraturan perundang-undangan.
p46   F13  | Pasal 120
p46   F14  [AYAT]
p46   F14  | (1) Jabatan di lingkungan Kementerian diisi oteh aparatur
p46   F16  | sipil negara yang profesional dan ahli sesuai dengan
p46   F14  | ketentuan peraturan perurndang-undangan.
p46   F15  [AYAT]
p46   F15  | (2) Ketentuan pengisian jabatan sebagaimana dimaksud
p46   F14  | pada ayat (1) dikecualikan untuk:
p46        [SUB-ITEM]
p46        | a. jabatan pimpinan tinggi madya dan jabatan
p46        | pimpinan tinggi pratama tertentu di lingkungan
p46   F13  | Sekretariat Presiden; dan
p46   F18  [SUB-ITEM]
p46   F18  | b. jabatan tertentu di lingkungan Sekretariat Militer
p46   F12  | Presiden.
p46   F14  [AYAT]
p46   F14  | (3) Jabatan sebagaimana dimaksud pada ayat (21 dapat
p46   F16  | diisi oleh prajurit Tentara Nasional Indonesia atau
p46   F14  | anggota Kepolisian Negara Republik Indonesia yang
p46   F17  | memiliki kompetensi dan keahlian sesuai dengan
p46   F14  | ketentuan peraturan perulndang-undangan.
p46   F13  [HEADING:BAB]
p46   F13  | BAB IX
p46   F12  | HAK KEUANGAN DAN FASILITAS LAINNYA
p46   F13  | Pasal 121
p46        [AYAT]
p46        | (1) Jabatan di lingkungan Kementerian diberikan hak
p46   F14  | keuangan dan fasilitas lainnya sesuai dengan ketentuan
p46   F14  | peraturan perundang-undangan.
p46   F16  [AYAT]
p46   F16  | (2) Dalam...
p46   F14  I1 | SK No 247686 A
==================== PAGE 47 ====================
p47   F12  | PRESIDEN
p47   F12  | REPUBUK INDONESIA
p47   F17  | -47-
p47   F15  [AYAT]
p47   F15  | (2) Dalam hal Sekretaris Kabinet sebagaimana dimaksud
p47   F15  | dalam Pasal 118 ayat (4) berasal dari prajurit Tentara
p47   F16  | Nasional Indonesia atau anggota Kepolisian Negara
p47   F14  | Republik Indonesia, hak keuangan dan fasilitas lainnya
p47   F13  | disesuaikan dengan golongan kepangkatan.
p47   F13  [HEADING:BAB]
p47   F13  | BAB X
p47   F12  | EVALUASI KELEMBAGAAN
p47   F13  | Pasal 122
p47   F14  [AYAT]
p47   F14  | (1) Penataan organisasi Kementerian dilakukan berdasarkan
p47        | evaluasi kelembagaan dan analisis kebutuhan
p47   F12  | organisasi.
p47   F14  [AYAT]
p47   F14  | (2) Kementerian melakukan evaluasi kelembagaan unit
p47   F18  | organisasi lain yang organisasi dan tata kerjanya
p47   F13  | ditetapkan dengan Peraturan Menteri.
p47   F15  [AYAT]
p47   F15  | (3) Evaluasi kelembagaan sebagaimana dimaksud pada
p47   F16  | ayat (1) dan ayat (2) dilakukan paling kurang 3 (tiga)
p47   F14  | tahun sekali.
p47   F13  [HEADING:BAB]
p47   F13  | BAB XI
p47   F12  | KETENTUAN LAIN-LAIN
p47   F12  | Pasal 123
p47   F16  | Pejabat pimpinan tinggi madya atau pejabat struktural
p47   F14  | eselon I.a yang dialihtugaskan pada jabatan Staf Ahli tetap
p47   F14  | diberikan status jabatan pimpinan tinggi madya atau jabatan
p47   F14  | struktural eselon I.a.
p47   F12  | Pasal 124
p47   F14  [AYAT]
p47   F14  | (1) Dalam hal jabatan administrasi yang belum disetarakan
p47        | ke dalam jabatan fungsional dapat dilakukan
p47   F14  | penyetaraan jabatan sesuai dengan ketentuan peraturan
p47   F13  | perundang-undangan.
p47   F16  [AYAT]
p47   F16  | (2) Penyetaraan jabatan sebagaimana dimaksud pada
p47   F13  | ayat (1) yaitu:
p47   F17  [SUB-ITEM]
p47   F17  | a. jabatan administrator ke jabatan fungsional ahli
p47   F13  | madya; dan
p47   F18  [SUB-ITEM]
p47   F18  | b. jabatan pengawas ke jabatan fungsional ahli muda.
p47   F14  [AYAT]
p47   F14  | (3) Penetapan kelas jabatan fungsional yang akan diduduki
p47   F15  | disetarakan dengan kelas jabatan administrasi yang
p47   F14  | diduduki sebelumnya.
p47        | BABXII ...
p47   F14  I1 | SK No 247992 A
==================== PAGE 48 ====================
p48   F12  | PIIESIDEN
p48   F12  | REPUBUK INDONESIA
p48   F17  | -48-
p48   F13  [HEADING:BAB]
p48   F13  | BAB XII
p48   F12  | KETENTUAN PERALIHAN
p48   F12  | Pasal 125
p48   F14  | Pada saat Peraturan Presiden ini mulai berlaku, Rancangan
p48   F16  | Peraturan Menteri/Kepala Lembaga yang telah melalui
p48   F13  | pengharmonisasian, pembulatan, dan pemantapan konsepsi
p48   F13  | yang dikoordinasikan oleh menteri atau kepala lembaga yang
p48        | menyelenggarakan urusan pemerintahan di bidang
p48   F14  | pembentukan peraturan perundang-undangan sebagaimana
p48   F16  | diatur dalam Peraturan Presiden Nomor 68 Tahun 2O2I
p48   F17  | tentang Pemberian Persetujuan Presiden Terhadap
p48   F14  | Rancangan Peraturan Menteri/Kepala Lembaga (Lembaran
p48   F16  | Negara Republik Indonesia Tahun 2O2I Nomor lT3),
p48   F14  | tidak dimintakan persetujuan Presiden.
p48   F12  | Pasal 126
p48   F17  | Pada saat Peraturan Presiden ini mulai berlaku, seluruh
p48   F14  | jabatan yang ada beserta pejabat yang memangku jabatan
p48   F15  | di lingkungan Kementerian tetap melaksanakan tugas dan
p48   F15  | fungsinya sampai dengan dibentuknya jabatan baru dan
p48   F14  | diangkat pejabat baru berdasarkan Peraturan Presiden ini.
p48   F13  [HEADING:BAB]
p48   F13  | BAB XIII
p48   F12  | KETENTUAN PENUTUP
p48   F12  | Pasal 127
p48   F14  | Pada saat Peraturan Presiden ini mulai berlaku, tugas dan
p48   F14  | fungsi Sekretariat Kabinet yang diintegrasikan ke dalam
p48        | Kementerian sebagaimana diatur dalam Peraturan
p48   F14  | Presiden Nomor 139 Tahun 2024 tentang Penataan T\:gas
p48   F16  | dan Fungsi Kementerian Negara Kabinet Merah Putih
p48   F15  | Periode Tahun 2024-2029 (Lembaran Negara Republik
p48   F14  | Indonesia Tahun 2024 Nomor 249) ditetapkan sebagai
p48        | tugas dan fungsi Sekretariat Dukungan Kabinet
p48   F14  | di lingkungan Kementerian.
p48   F14  I1 | SK No 247688 A
p48   F13  | Pasal 128
==================== PAGE 49 ====================
p49   F12  | PRESIDEN
p49   F12  | REPUBUK INDONESIA
p49   F17  | -49-
p49   F12  | Pasal 128
p49   F18  | Pada saat Peraturan Presiden ini mulai berlaku, semua
p49        | peraturan perundang-undangan yang merupakan
p49   F14  | peraturan pelaksanaan dari:
p49   F18  [SUB-ITEM]
p49   F18  | a. Peraturan Presiden Nomor 31 Tahun 2O2O tentang
p49   F15  | Kementerian Sekretariat Negara (Lembaran Negara
p49   F14  | Republik Indonesia Tahun 2O2O Nomor 45); dan
p49   F17  [SUB-ITEM]
p49   F17  | b. Peraturan Presiden Nomor 55 Tahun 2O2O tentang
p49   F17  | Sekretariat Kabinet (Lembaran Negara Repubtik
p49   F14  | Indonesia Tahun 2O2O Nomor 95),
p49        | dinyatakan masih tetap berlaku sepanjang tidak
p49   F14  | bertentangan dengan ketentuan dalam Peraturan Presiden
p49   F14  | ini.
p49   F12  | Pasal 129
p49   F14  | Pada saat Peraturan Presiden ini mulai berlaku:
p49        [SUB-ITEM]
p49        | a. Peraturan Presiden Nomor 31 Tahun 2O2O tentang
p49   F14  | Kementerian Sekretariat Negara (Lembaran Negara
p49   F14  | Republik Indonesia Tahun 2O2O Nomor 45);
p49        [SUB-ITEM]
p49        | b. Peraturan Presiden Nomor 55 Tahun 2O2O tentang
p49   F16  | Sekretariat Kabinet (Lembaran Negara Republik
p49   F14  | Indonesia Tahun 2O2O Nomor 95); dan
p49        [SUB-ITEM]
p49        | c. Peraturan Presiden Nomor 68 Tahun 2O2l tentang
p49   F14  | Pemberian Persetujuan Presiden Terhadap Rancangan
p49   F14  | Peraturan Menteri/Kepala Lembaga (Lembaran Negara
p49   F14  | Republik Indonesia Tahun 2O2l Nomor 1731,
p49   F14  | dicabut dan dinyatakan tidak berlaku.
p49   F13  | Pasal 130
p49   F15  | Peraturan Presiden
p49   F13  | diundangkan.
p49        | ini mulai berlaku pada tanggal
p49   F14  I1 | SK No 247689 A
p49   F12  | Agar
==================== PAGE 50 ====================
p50   F12  | PRESIDEN
p50   F12  | REPUBUK INDONESIA
p50   F17  | -50-
p50        | Agar setiap orang mengetahuinya, memerintahkan
p50        | pengundangan Peraturan Presiden ini
p50   F12  | dengan
p50   F18  | penempatannya dalam Lembaran Negara Republik
p50   F12  | Indonesia.
p50   F14  I5 | Ditetapkan di Jakarta
p50   F13  I5 | pada tanggal 5 November 2024
p50   F12  I5 | PRESIDEN REPUBLIK INDONESIA,
p50   F14  I7 | ttd
p50   F12  I5 | PRABOWO SUBIANTO
p50   F14  I2 | Diundangkan di Jakarta
p50   F13  I2 | pada tanggal 5 November 2024
p50   F12  I2 | MENTERI SEKRETARIS NEGARA
p50   F12  I3 | REPUBLIK INDONESIA,
p50   F14  I4 | ttd.
p50   F12  I3 | PRASETYO HADI
p50   F13  I2 | LEMBARAN NEGARA REPUBLIK INDONESIA TAHUN 2024 NOMOR 344
p50   F13  I3 | Salinan sesuai dengan aslinya
p50   F12  I2 | KEMENTERIAN SEKRETARIAT NEGARA
p50   F12  I3 | REPUBLIK INDONESIA
p50   F14  I3 | ti Bidang Perundang-undangan
p50   F14  I4 | Administrasi Hukum,
p50   F14  I1 | SK No 247592 A
p50   F13  | Djaman
```

---


## perpres

- **File**: `perpres/perpres-no-127-tahun-2024.pdf`
- **Document Type**: Peraturan Presiden
- **Issued by**: Presiden
- **Pages**: 5 | **Lines**: 191
- **Font sizes**: [7.0, 7.5, 8.0, 9.5, 11.0, 11.5, 12.0, 12.5, 13.0, 13.5, 14.0, 14.5, 15.0, 15.5, 16.0, 16.5, 17.0, 17.5, 18.0, 18.5, 19.0, 19.5, 22.5]
- **Most common font**: 12.5 (17% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [46.0, 159.0, 256.0, 317.0, 376.0, 412.0, 467.0, 497.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat / clauses

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01   F22  I6 | SALINAN
p01   F12  I3 | PRESIDEN
p01   F12  I3 | REPUBLIK INDONESIA
p01        I2 | PERATURAN PRESIDEN REPUBLIK INDONESIA
p01   F14  | NOMOR I27 TAHUN 2024
p01   F12  I3 | TENTANG
p01        | TUNJANGAN JABATAN FUNGSIONAL
p01        | KURATOR KOLEKSI HAYATI
p01   F13  I2 [PREAMBLE:DENGAN RAHMAT]
p01   F13  I2 | DENGAN RAHMATTUHAN YANG MAHA ESA
p01        | PRESIDEN REPUBLIK INDONESIA,
p01   F13  [PREAMBLE:MENIMBANG]
p01   F13  | Menimbang
p01   F12  I2 [SUB-ITEM]
p01   F12  I2 | a.
p01   F13  [PREAMBLE:MENGINGAT]
p01   F13  | Mengingat
p01   F12  I2 | 2
p01   F8   I2 | 1
p01        I2 | b
p01   F14  | bahwa untuk meningkatkan mutu, prestasi, pengabdian,
p01   F16  | dan produktivitas kinerja pegawai Negeri Sipil yang
p01   F14  | diangkat dan ditugaskan secara penuh dalam Jabatan
p01   F16  | Fungsional Kurator Koleksi Hayati, perlu diberikan
p01   F14  | Tunjangan Jabatan Fungsional Kurator Koleksi Hayati
p01   F14  | sesuai dengan beban kerja, tanggung jawab, dan risiko
p01   F13  | pekerjaan;
p01   F18  | bahwa berdasarkan pertimbangan sebagaimana
p01   F15  | dimaksud dalam huruf a, perlu menetapkan peraturan
p01   F14  | Presiden tentang Tunjangan Jabatan Fungsional Kurator
p01   F13  | Koleksi Hayati;
p01   F14  | Pasal 4 ayat (1) Undang-Undang Dasar Negara Republik
p01   F14  | Indonesia Tahun 1945;
p01   F14  | Undang-Undang Nomor 2O Tahun 2O23 tentang Aparatur
p01   F17  | Sioil Negara (Lembaran Negara Republik lndonesia
p01   F14  | Tahun 2023 Nomor 141, Tambahan Lembaran Negara
p01   F14  | Republik Indonesia Nomor 6897);
p01   F18  | Peraturan Pemerintah Nomor 7 Tahun lgZZ terrtang
p01   F16  | leraturan Gaji Pegawai Negeri Sipil (kmbaran Negari
p01   F20  | Republik Indonesia Tahun 1977 Nomor ll, Tambahan
p01   F16  | lembaran Negara Republik Indonesia Nomor 3O9g)
p01   F14  | sebagaimana telah beberapa kali diubah, terakhir dengan
p01   F16  | Peraturan Pemerintah Nomor S Tahun 2024 tentang
p01   F14  | Perubahan Kesembilan Belas atas peraturan pemerintah
p01   F16  | Nomor 7 Tahun 1977 tentang peraturan Gaji pegawai
p01   F18  | Negeri Sipil (Lembaran Negara Republik -Indonesia
p01   F14  | Tahun 2O24 Nomor 15);
p01   F12  I2 | 3
p01   F14  I1 | SK No243676A
p01   F14  [ITEM]
p01   F14  | 4. Peraturan . . .
==================== PAGE 2 ====================
p02   F12  I3 | PRESIDEN
p02        I3 | REPUBLIK I{DONESIA
p02   F14  | -z-
p02   F18  [ITEM]
p02   F18  | 4. Peraturan Pemerintah Nomor 1l Tahun 2Ol7 tentang
p02   F16  | Manajemen Pegawai Negeri Sipil (kmbaran Negara
p02   F14  | Republik Indonesia Tahun 2017 Nomor 63, Tambahan
p02   F16  | lembaran Negara Republik Indonesia Nomor 6037)
p02   F14  | sslagaimana telah diubah dengan Peraturan Pemerintah
p02   F13  | Nomor 17 Tahun 2O2O tentangPerubahan atas Peraturan
p02   F14  | Pemerintah Nomor 11 Tahun 2OL7 ter,ltang Manajemen
p02   F18  | Pegawai Negeri Sipil (Lembaran Negara Republik
p02   F14  | Indonesia Tahun 2020 Nomor 68, Tambahan Lembaran
p02   F13  | Negara Republik Indonesia Nomor 6477);
p02   F18  I2 [ITEM]
p02   F18  I2 | 5. Keputusan Presiden Nomor 87 Tahun 1999 tentang
p02   F16  | Rumpun Jabatan F\rngsional Pegawai Negeri Sipil
p02   F14  | sebagaimana telah beberapa kali diubah, terakhir dengan
p02   F15  | Peraturan Presiden Nomor 116 Tahun 2014 tentang
p02   F14  | Perubahan Kedua atas Keputusan Presiden Nomor 87
p02   F16  | Tahun 1999 tentang Rumpun Jabatan Fungsional
p02   F18  | Pegawai Negeri Sipil (Lembaran Negara Republik
p02   F14  | Indonesia Tahun 2014 Nomor 240);
p02        I3 [KEPUTUSAN:MEMUTUSKAN]
p02        I3 | MEMUTUSKAN:
p02        I2 | PERATURAN PRESIDEN TENTANG TUNJANGAN JABATAN
p02        I2 | FUNGSIONAL KURATOR KOLEKSI HAYATI.
p02        I3 | Pasal 1
p02   F18  I2 | Dalam Peraturan Presiden ini yang dimaksud dengan
p02   F14  I2 | Tunjangan Jabatan Fungsional Kurator Koleksi Hayati yang
p02   F16  I2 | selanjutnya disebut Tunjangan Kurator Koleksi Hayati
p02   F14  I2 | adalah tunjangan jabatan yang diberikan kepada Pegawai
p02   F15  I2 | Negeri Sipil yang diangkat dan ditugaskan secara penuh
p02   F14  I2 | dalam Jabatan Fungsional Kurator Koleksi Hayati sesuai
p02   F14  I2 | dengan ketentuan peraturan perundang-undangan.
p02   F12  I3 | Pasal 2
p02   F14  I2 | Pegawai Negeri Sipil yang diangkat dan ditugaskan secara
p02   F14  I2 | penuh dalam Jabatan Fungsional Kurator Koleksi Hayati
p02   F14  I2 | diberikan 'I\rnjangan Kurator Koleksi Hayati setiap bulan.
p02   F18  I7 | Pasal 3...
p02   F14  I1 | SK No 211339A
==================== PAGE 3 ====================
p03   F12  I3 | FRESIDEN
p03   F12  I3 | REPUBL]K INDONESIA
p03   F19  | -3-
p03   F15  I2 | Besaran T\rnjangan Kurator Koleksi Hayati sebagaimana
p03   F16  I2 | dimaksud dalam Pasal 2 tercantum dalam Lampiran yang
p03   F18  I2 | merupakan bagran tidak terpisahkan dari Peraturan
p03   F14  I2 | Presiden ini.
p03   F12  | Pasal 3
p03        I3 | Pasal 4
p03   F16  I2 | Peraturan Presiden ini
p03   F14  I2 | diundangkan.
p03   F18  | mulai berlaku pada tanggal
p03   F14  I2 | Pemberian Ttrnjangan Kurator Koleksi Hayati bagi:
p03   F19  [SUB-ITEM]
p03   F19  | a. Pegawai Negeri Sipil yang bekerja pada instansi pusat
p03   F17  | bersumber dari Anggaran Pendapatan dan Belanja
p03   F13  | Negara; dan
p03   F18  [SUB-ITEM]
p03   F18  | b. Pegawai Negeri Sipil yang bekerja pada instansi daerah
p03   F17  | bersumber dari Anggaran Pendapatan dan Belanja
p03        | Daerah.
p03   F12  I3 | Pasal 5
p03   F14  I2 | Pemberian T\rnjangan Kurator Koleksi Hayati dihentikan
p03   F14  I2 | apabila Pegawai Negeri Sipil sglagaimana dimaksud dalam
p03   F18  I2 | Pasal 2 diangkat dalam jabatan struktural, jabatan
p03   F15  I2 | fungsional lain, atau karena hal lain yang mengakibatkan
p03   F14  I2 | pemberian Tunjangan Kurator Koleksi Hayati dihentikan
p03   F14  I2 | sesuai dengan ketentuan peraturan perundang-undangan.
p03        I3 | Pasal 6
p03   F18  I2 | Tata cara pembayaran dan penghentian pembayaran
p03   F16  I2 | Tfrnjangan Kurator Koleksi Hayati dilaksanakan sesuai
p03   F14  I2 | dengan ketentuan peraturan perundang-undangan.
p03   F12  I3 | Pasal 7
p03   F14  I1 | SK No2l1340A
p03        I7 | Agar
==================== PAGE 4 ====================
p04   F16  I4 | I
p04   F11  I3 | PRESTDEN
p04        I3 | REPUBUK INDONESIA
p04   F18  | -4-
p04   F18  I2 | Agar setiap
p04   F13  I2 | pengundangan
p04   F13  I2 | penempatannya
p04   F13  I2 | Indonesia.
p04   F18  | orang mengetahuinya, memerintahkan
p04   F18  I4 | Peraturan presiden ini
p04        I8 | dengan
p04   F18  I4 | dalam kmbaran Negara nepuStik
p04   F14  I4 | Ditetapkan di Jakarta
p04   F14  I4 | pada tanggal 17 Oktober 2024
p04        I4 | PRESIDEN REPUBLIK INDONESIA,
p04   F14  I6 | ttd
p04        I5 | JOKO WIDODO
p04   F14  | Diundangkan di Jakarta
p04   F14  | pada tanggal LT Oktober 2024
p04        | MENTERI SEKRETARIS NEGARA
p04        | REPUBLIK INDONESIA,
p04   F14  I2 | ttd.
p04   F12  I2 | PRATIKNO
p04   F13  | LEMBARAN NEGARA REPUBLIK INDONESIA TAHUN 2024 NOMOR 234
p04   F13  | Salinan sesuai dengan aslinya
p04        | KEMENTERIAN SEKRETARIAT NEGARA
p04        | REPUBLIK INDONESIA
p04   F13  I2 | dang Perundang-undangan
p04   F14  | strasi Hukum
p04   F14  I1 | SK No243677A
p04   F14  I2 | ilvanna Djaman
==================== PAGE 5 ====================
p05   F12  I3 | PRESIDEN
p05   F12  I3 | REPUBLIK INDONESTA
p05        | TUNJANGAN JABATAN FUNGSIONAL
p05        | KURATOR KOLEKSI HAYATI
p05   F12  I3 | LAMPIRAN
p05        I3 | PERATURAN PRESIDEN REPUBLIK INDONESIA
p05   F14  I3 | NOMOR I27 TAHUN 2024
p05   F12  I3 | TENTANG
p05   F20  I3 | TUNJANGAN JABATAN FUNGSIONAL
p05        I3 | KURATOR KOLEKSI HAYATI
p05   F12  | NO
p05        I2 | JABATAN FUNGSIONAL
p05   F12  | BESARAN
p05        I6 | TUNJANGAN
p05   F14  I2 | Jenjang Jabatan Fungsional Keahlian
p05   F8   | 1
p05   F14  | Kurator Koleksi Hayati Ahli Utama
p05   F13  | Rp2.025.000,00
p05   F12  | 2
p05   F14  | Kurator Koleksi Hayati Ahli Madya
p05   F12  | Rp1.38O.OO0,0o
p05   F12  | 3
p05   F14  | Kurator Koleksi Hayati Ahli Muda
p05   F13  | Rp1.100.000,00
p05        | 4
p05   F14  | Kurator Koleksi Hayati Ahli Pertama
p05   F12  I7 | Rp5a0.O00,O0
p05        I4 | PRESIDEN REPUBLIK INDONESIA,
p05        I5 | JOKO WIDODO
p05   F14  I6 | ttd
p05   F14  | Salinan sesuai dengan aslinya
p05        | KEMENTERIAN SEKRETARI.AT NEGARA
p05        | REPUBLIK INDONESIA
p05   F15  | Bidalg Perundang-undangan
p05   F14  | strasi Hukum,
p05   F8   | SEKR
p05   F12  I1 | E
p05   F15  I1 | rrrY*
p05   F8   | D
p05   F10  | tK
p05   F7   | N
p05   F14  I1 | SK No243678A
p05   F13  I2 | anna Djaman
```

---


## perda

- **File**: `perda/perda-kabupaten-sukoharjo-no-1-tahun-2025.pdf`
- **Document Type**: Peraturan Daerah (Regional Reg)
- **Issued by**: Kepala Daerah
- **Pages**: 18 | **Lines**: 1061
- **Font sizes**: [11.0, 12.0]
- **Most common font**: 12.0 (99% of lines)
- **Bold font sizes**: [12.0]
- **Indent clusters**: [53.0, 71.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01        | BUPATI SUKOHARJO
p01        I1 | PROVINSI JAWA TENGAH
p01        | PERATURAN DAERAH KABUPATEN SUKOHARJO
p01        | NOMOR 1 TAHUN 2025
p01        | TENTANG
p01        | KAWASAN TANPA ROKOK
p01        [PREAMBLE:DENGAN RAHMAT]
p01        | DENGAN RAHMAT TUHAN YANG MAHA ESA
p01        | BUPATI SUKOHARJO,
p01        [PREAMBLE:MENIMBANG]
p01        | Menimbang
p01        | : a. bahwa menjaga kesehatan merupakan salah satu
p01        | unsur kesejahteraan yang harus diwujudkan
p01        | sesuai dengan cita-cita bangsa Indonesia yang
p01        | tertuang dalam Pancasila dan Undang-Undang
p01        | Dasar Negara Republik Indonesia Tahun 1945;
p01        [SUB-ITEM]
p01        | b. bahwa perilaku merokok dan paparan asap rokok
p01        | dapat mengakibatkan gangguan atau bahaya bagi
p01        | kesehatan dan kualitas hidup sehingga diperlukan
p01        | upaya pengendalian dampak rokok terhadap
p01        | kesehatan individu, keluarga, masyarakat, dan
p01        | lingkungan;
p01        [SUB-ITEM]
p01        | c. bahwa untuk memberikan arah, landasan, dan
p01        | kepastian hukum kepada semua pihak yang
p01        | terlibat dalam kawasan tanpa rokok, maka
p01        | diperlukan pengaturan tentang kawasan tanpa
p01        | rokok;
p01        [SUB-ITEM]
p01        | d. bahwa berdasarkan pertimbangan sebagaimana
p01        | dimaksud dalam huruf a, huruf b, dan huruf c
p01        | perlu menetapkan Peraturan Daerah tentang
p01        | Kawasan Tanpa Rokok;
p01        [PREAMBLE:MENGINGAT]
p01        | Mengingat
p01        | : 1. Pasal 18 ayat (6) Undang-Undang Dasar Negara
p01        | Republik Indonesia 1945;
p01        [ITEM]
p01        | 2. Undang-Undang Nomor 13 Tahun 1950 tentang
p01        | Pembentukan Daerah-daerah Kabupaten dalam
p01        | Lingkungan Propinsi Djawa Tengah sebagaimana
p01        | telah diubah dengan Undang-Undang Nomor 9
p01        | Tahun 1965 tentang Pembentukan Daerah Tingkat
p01        | II Batang dengan mengubah Undang-Undang No.
p01        | 13 Tahun 1950 tentang Pembentukan Daerah-
p01        | daerah Kabupaten dalam Lingkungan Propinsi
p01        | Jawa Tengah (Lembaran Negara Tahun 1965
p01        | Nomor 52, Tambahan Lembaran Negara Nomor
p01        | 2757);
p01   F11  | SALINAN
==================== PAGE 2 ====================
p02        | 2
p02        | Dengan Persetujuan Bersama
p02        | DEWAN PERWAKILAN RAKYAT DAERAH KABUPATEN SUKOHARJO
p02        | dan
p02        | BUPATI SUKOHARJO
p02        [KEPUTUSAN:MEMUTUSKAN]
p02        | MEMUTUSKAN:
p02        I1 [PREAMBLE:MENETAPKAN]
p02        I1 | Menetapkan
p02        | : PERATURAN DAERAH TENTANG KAWASAN TANPA
p02        | ROKOK.
p02        [HEADING:BAB]
p02        | BAB I
p02        | KETENTUAN UMUM
p02        | Pasal 1
p02        | Dalam Peraturan Daerah ini yang dimaksud dengan:
p02        [ITEM]
p02        | 1.   Daerah adalah Kabupaten Sukoharjo.
p02        [ITEM]
p02        | 2.   Pemerintah Daerah adalah Bupati sebagai unsur
p02        | penyelenggara Pemerintahan Daerah yang memimpin
p02        | pelaksanaan urusan pemerintahan yang menjadi
p02        | kewenangan daerah otonom.
p02        [ITEM]
p02        | 3.   Bupati adalah Bupati Sukoharjo.
p02        [ITEM]
p02        | 4.   Perangkat Daerah adalah unsur pembantu Bupati dan
p02        | Dewan
p02        | Perwakilan
p02        | Rakyat
p02        | Daerah
p02        | dalam
p02        | penyelenggaraan Urusan Pemerintahan yang menjadi
p02        | kewenangan Daerah.
p02        [ITEM]
p02        | 5.   Kesehatan adalah keadaan sehat seseorang, baik
p02        | secara fisik, jiwa, maupun sosial dan bukan sekadar
p02        | terbebas dari penyakit untuk memungkinkannya hidup
p02        | produktif.
p02        [ITEM]
p02        | 6.   Rokok
p02        | adalah
p02        | semua
p02        | produk
p02        | tembakau
p02        | yang
p02        | dimaksudkan untuk dibakar dan dihisap dan/atau
p02        | dihirup asapnya, termasuk rokok kretek, rokok putih,
p02        | cerutu atau bentuk lainnya yang dihasilkan dari
p02        | tanaman nicotiana tabacum, nicotiana rustica, dan
p02        | spesies lainnya atau sintetisnya termasuk shisha,
p02        | rokok elektrik, produk tembakau yang dipanaskan dan
p02        | bentuk lainnya yang mengandung nikotin dan tar
p02        | dengan atau tanpa bahan tambahan.
p02        [ITEM]
p02        | 3. Undang-Undang Nomor 23 Tahun 2014 tentang
p02        | Pemerintahan Daerah (Lembaran Negara Republik
p02        | Indonesia Tahun 2014 Nomor 244, Tambahan
p02        | Lembaran Negara Republik Indonesia Nomor 5587)
p02        | sebagaimana telah beberapa kali diubah terakhir
p02        | dengan Undang-Undang Nomor 6 Tahun 2023
p02        | tentang
p02        | Penetapan
p02        | Peraturan
p02        | Pemerintah
p02        | Pengganti Undang-Undang Nomor 2 Tahun 2022
p02        | tentang Cipta Kerja menjadi Undang-Undang
p02        | (Lembaran Negara Republik Indonesia Tahun 2023
p02        | Nomor 41, Tambahan Lembaran Negara Republik
p02        | Indonesia Nomor 6856);
==================== PAGE 3 ====================
p03        | 3
p03        [ITEM]
p03        | 7.   Merokok adalah kegiatan membakar, memanaskan,
p03        | menguapkan
p03        | Rokok,
p03        | menghisapnya
p03        | dan/atau
p03        | menghirup asapnya, kemudian menghembuskannya.
p03        [ITEM]
p03        | 8.   Kawasan Tanpa Rokok yang selanjutnya disingkat KTR
p03        | adalah ruangan atau area yang dinyatakan dilarang
p03        | untuk kegiatan merokok, atau kegiatan menjual,
p03        | memproduksi, mengiklankan, di dalam maupun di luar
p03        | ruang, dan mempromosikan produk tembakau dan
p03        | rokok elektronik.
p03        [ITEM]
p03        | 9.   Fasilitas Pelayanan Kesehatan adalah tempat dan/
p03        | atau alat yang digunakan untuk menyelenggarakan
p03        | pelayanan kesehatan kepada perseorangan ataupun
p03        | masyarakat dengan pendekatan promotif, preventif,
p03        | kuratif, rehabilitatif, dan/atau paliatif  yang dilakukan
p03        | oleh pemerintah pusat, Pemerintah Daerah, dan/atau
p03        | masyarakat.
p03        [ITEM]
p03        | 10. Tempat Proses Belajar Mengajar adalah tempat yang
p03        | digunakan
p03        | untuk
p03        | kegiatan
p03        | belajar
p03        | mengajar,
p03        | pendidikan, dan/atau pelatihan.
p03        [ITEM]
p03        | 11. Tempat Anak Bermain adalah tempat atau arena
p03        | tertutup atau terbuka yang digunakan untuk kegiatan
p03        | bermain anak-anak.
p03        [ITEM]
p03        | 12. Tempat Ibadah adalah bangunan atau ruang yang
p03        | memiliki ciri tertentu yang khusus dipergunakan
p03        | untuk beribadah bagi para pemeluk masing-masing
p03        | agama secara permanen, tidak termasuk tempat
p03        | ibadah keluarga.
p03        [ITEM]
p03        | 13. Angkutan
p03        | Umum
p03        | adalah
p03        | alat
p03        | angkutan
p03        | bagi
p03        | masyarakat yang dapat berupa kendaraan darat, air,
p03        | dan udara biasanya dengan kompensasi.
p03        [ITEM]
p03        | 14. Tempat Kerja Tertentu adalah setiap tempat atau
p03        | gedung tertentu tertutup dan/atau terbuka bergerak
p03        | atau tidak bergerak yang digunakan untuk bekerja
p03        | dengan
p03        | mendapatkan
p03        | kompensasi
p03        | (gaji/upah)
p03        | termasuk tempat lain yang  dilintasi oleh pekerja di
p03        | KTR.
p03        [ITEM]
p03        | 15. Tempat Umum adalah semua tempat tertutup yang
p03        | dapat diakses oleh masyarakat umum dan/atau
p03        | tempat yang dapat dimanfaatkan bersama-sama untuk
p03        | kegiatan masyarakat yang dikelola oleh pemerintah
p03        | pusat
p03        | atau
p03        | Pemerintah
p03        | Daerah,
p03        | swasta,
p03        | dan
p03        | masyarakat, seperti hotel, restoran, bioskop, bandar
p03        | udara,
p03        | stasiun,
p03        | pusat
p03        | perbelanjaan,
p03        | dan
p03        | pasar
p03        | swalayan.
p03        [ITEM]
p03        | 16. Pengelola adalah orang dan/atau badan hukum yang
p03        | karena jabatannya memimpin dan/atau bertanggung
p03        | jawab atas kegiatan dan/atau usaha di tempat atau
p03        | kawasan yang ditetapkan sebagai KTR, baik milik
p03        | pemerintah maupun swasta.
==================== PAGE 4 ====================
p04        | 4
p04        [ITEM]
p04        | 17. Badan adalah sekumpulan orang dan/atau modal yang
p04        | merupakan kesatuan, baik yang melakukan usaha
p04        | maupun yang tidak melakukan usaha yang meliputi
p04        | perseroan terbatas, perseroan komanditer, perseroan
p04        | lainnya, badan usaha milik negara, badan usaha milik
p04        | Daerah, atau badan usaha milik desa, dengan nama
p04        | dan dalam bentuk apapun, firma, kongsi, koperasi,
p04        | dana pensiun, persekutuan, perkumpulan, yayasan,
p04        | organisasi massa, organisasi sosial politik, atau
p04        | organisasi lainnya, lembaga dan bentuk badan lainnya,
p04        | termasuk kontrak investasi kolektif dan bentuk usaha
p04        | tetap.
p04        | Pasal 2
p04        | Peraturan Daerah ini berdasarkan atas asas:
p04        [SUB-ITEM]
p04        | a. kepentingan kualitas Kesehatan manusia;
p04        [SUB-ITEM]
p04        | b. keseimbangan;
p04        [SUB-ITEM]
p04        | c. kemanfaatan;
p04        [SUB-ITEM]
p04        | d. keterpaduan;
p04        [SUB-ITEM]
p04        | e. keserasian;
p04        [SUB-ITEM]
p04        | f.
p04        | partisipasi;
p04        [SUB-ITEM]
p04        | g. keadilan; dan
p04        [SUB-ITEM]
p04        | h. transparansi dan akuntabilitas.
p04        | Pasal 3
p04        | Penetapan KTR bertujuan untuk:
p04        [SUB-ITEM]
p04        | a. memberikan acuan bagi Pemerintah Daerah dalam
p04        [PREAMBLE:MENETAPKAN]
p04        | menetapkan KTR;
p04        [SUB-ITEM]
p04        | b. memberikan perlindungan yang efektif dari bahaya
p04        | asap rokok;
p04        [SUB-ITEM]
p04        | c. memberikan ruang dan lingkungan yang bersih  serta
p04        | sehat bagi masyarakat; dan
p04        [SUB-ITEM]
p04        | d. melindungi Kesehatan masyarakat secara umum dari
p04        | dampak buruk merokok baik langsung maupun tidak
p04        | langsung.
p04        | Pasal 4
p04        | Ruang lingkup pengaturan dalam Peraturan Daerah ini
p04        | meliputi:
p04        [SUB-ITEM]
p04        | a. penetapan KTR;
p04        [SUB-ITEM]
p04        | b. tanggung
p04        | jawab,
p04        | kewajiban,
p04        | larangan
p04        | dan
p04        | pengendalian;
p04        [SUB-ITEM]
p04        | c. partisipasi masyarakat;
p04        [SUB-ITEM]
p04        | d. satuan tugas KTR;
p04        [SUB-ITEM]
p04        | e. pembinaan dan pengawasan; dan
p04        [SUB-ITEM]
p04        | f.
p04        | pendanaan.
==================== PAGE 5 ====================
p05        | 5
p05        [HEADING:BAB]
p05        | BAB II
p05        | PENETAPAN KTR
p05        | Pasal 5
p05        | KTR terdiri atas:
p05        [SUB-ITEM]
p05        | a. Fasilitas Pelayanan Kesehatan;
p05        [SUB-ITEM]
p05        | b. Tempat Proses Belajar Mengajar;
p05        [SUB-ITEM]
p05        | c. Tempat Anak Bermain;
p05        [SUB-ITEM]
p05        | d. Tempat Ibadah;
p05        [SUB-ITEM]
p05        | e. Angkutan Umum;
p05        [SUB-ITEM]
p05        | f.
p05        | Tempat Kerja Tertentu; dan
p05        [SUB-ITEM]
p05        | g. Tempat Umum.
p05        | Pasal 6
p05        [AYAT]
p05        | (1) KTR sebagaimana dimaksud dalam Pasal 5 huruf a
p05        | sampai dengan huruf e merupakan kawasan yang
p05        | bebas dari asap Rokok hingga batas terluar.
p05        [AYAT]
p05        | (2) KTR sebagaimana dimaksud dalam Pasal 5 huruf e
p05        | berlaku pada saat Angkutan Umum sedang berhenti
p05        | dan/atau beroperasional.
p05        [AYAT]
p05        | (3) KTR sebagaimana dimaksud dalam Pasal 5 huruf f dan
p05        | huruf g merupakan kawasan yang bebas dari asap
p05        | Rokok hingga batas kucuran air dari atap paling luar.
p05        [AYAT]
p05        | (4) KTR sebagaimana dimaksud dalam Pasal 5 huruf a
p05        | sampai dengan huruf g diatur lebih lanjut dalam
p05        | Peraturan Bupati.
p05        | Pasal 7
p05        [AYAT]
p05        | (1) Pengelola, penyelenggara atau penanggung jawab KTR
p05        | sebagaimana dimaksud dalam Pasal 5 huruf a sampai
p05        | dengan huruf e dilarang menyediakan tempat khusus
p05        | untuk Merokok.
p05        [AYAT]
p05        | (2) Pengelola, penyelenggara atau penanggung jawab
p05        | Tempat
p05        | Kerja
p05        | Tertentu,
p05        | dan
p05        | Tempat
p05        | Umum
p05        | sebagaimana dimaksud dalam Pasal 5 huruf f dan
p05        | huruf g wajib menyediakan tempat khusus untuk
p05        | Merokok.
p05        [AYAT]
p05        | (3) Pengelola, penyelenggara atau penanggung jawab
p05        | Tempat
p05        | Kerja
p05        | Tertentu
p05        | dan
p05        | Tempat
p05        | Umum
p05        | sebagaimana dimaksud pada ayat (2) dikecualikan
p05        | pada tempat yang berpotensi menimbulkan bahaya
p05        | Kesehatan dan keselamatan kerja sesuai dengan
p05        | ketentuan peraturan perundang-undangan.
p05        [AYAT]
p05        | (4) Setiap orang atau Badan yang melanggar ketentuan
p05        | sebagaimana dimaksud pada ayat (1) dikenai sanksi
p05        | administratif berupa:
p05        [SUB-ITEM]
p05        | a. teguran lisan;
p05        [SUB-ITEM]
p05        | b. teguran tertulis; dan
p05        [SUB-ITEM]
p05        | c. denda administratif.
==================== PAGE 6 ====================
p06        | 6
p06        [AYAT]
p06        | (5) Setiap orang atau Badan yang melanggar ketentuan
p06        | sebagaimana dimaksud pada ayat (2) dikenai sanksi
p06        | administratif berupa:
p06        [SUB-ITEM]
p06        | a. teguran lisan;
p06        [SUB-ITEM]
p06        | b. teguran tertulis; dan
p06        [SUB-ITEM]
p06        | c. denda administratif.
p06        [AYAT]
p06        | (6) Ketentuan lebih lanjut mengenai tata cara pengenaan
p06        | sanksi administratif sebagaimana dimaksud pada ayat
p06        [AYAT]
p06        | (4) dan ayat (5) diatur dalam Peraturan Bupati.
p06        | Pasal 8
p06        | Tempat khusus untuk Merokok sebagaimana dimaksud
p06        | dalam Pasal 7 ayat (2) harus memenuhi persyaratan:
p06        [SUB-ITEM]
p06        | a. merupakan
p06        | ruang
p06        | terbuka
p06        | atau
p06        | ruang
p06        | yang
p06        | berhubungan langsung dengan udara luar, sehingga
p06        | udara dapat bersirkulasi dengan baik;
p06        [SUB-ITEM]
p06        | b. terpisah dari gedung/tempat/ruang utama dan ruang
p06        | lain yang digunakan untuk beraktivitas;
p06        [SUB-ITEM]
p06        | c. jauh dari pintu masuk dan keluar; dan
p06        [SUB-ITEM]
p06        | d. jauh dari tempat orang berlalu-lalang.
p06        [HEADING:BAB]
p06        | BAB III
p06        | TANGGUNG JAWAB, KEWAJIBAN, LARANGAN DAN
p06        | PENGENDALIAN
p06        [HEADING:BAGIAN]
p06        | Bagian Kesatu
p06        | Tanggung Jawab
p06        | Pasal 9
p06        [AYAT]
p06        | (1) Bupati bertanggung jawab terhadap pelaksanaan KTR.
p06        [AYAT]
p06        | (2) Tanggung jawab sebagaimana dimaksud pada ayat (1)
p06        | dilakukan Perangkat Daerah yang melaksanakan
p06        | urusan pemerintahan di bidang Kesehatan untuk:
p06        [SUB-ITEM]
p06        | a. mengumpulkan data dan informasi tentang KTR di
p06        | Daerah;
p06        [SUB-ITEM]
p06        | b. melakukan edukasi tentang bahaya rokok bagi
p06        | masyarakat;
p06        [SUB-ITEM]
p06        | c. menyediakan layanan konseling dan intervensi
p06        | farmakologi
p06        | berhenti
p06        | Merokok
p06        | di
p06        | Fasilitas
p06        | Pelayanan Kesehatan;
p06        [SUB-ITEM]
p06        | d. melakukan
p06        | sosialisasi
p06        | peraturan
p06        | perundang-
p06        | undangan yang berkaitan dengan KTR; dan
p06        [SUB-ITEM]
p06        | e. melakukan pemantauan dan evaluasi terhadap
p06        | pelaksanaan KTR.
==================== PAGE 7 ====================
p07        | 7
p07        [HEADING:BAGIAN]
p07        | Bagian Kedua
p07        | Kewajiban
p07        | Pasal 10
p07        [AYAT]
p07        | (1) Setiap Pengelola KTR wajib:
p07        [SUB-ITEM]
p07        | a. melakukan pengawasan internal pada tempat
p07        | dan/atau lokasi yang menjadi tanggung jawabnya;
p07        [SUB-ITEM]
p07        | b. melarang semua orang yang Merokok di KTR yang
p07        | menjadi tanggung jawabnya;
p07        [SUB-ITEM]
p07        | c. tidak menyediakan asbak atau sejenisnya pada
p07        | tempat dan/atau lokasi yang menjadi tanggung
p07        | jawabnya; dan
p07        [SUB-ITEM]
p07        | d. memasang
p07        | tanda
p07        | dilarang
p07        | Merokok
p07        | sesuai
p07        | peraturan perundangan-undangan di semua pintu
p07        | masuk
p07        | utama
p07        | dan
p07        | di
p07        | tempat-tempat
p07        | yang
p07        | dipandang perlu dan mudah terbaca dan/atau
p07        | didengar baik.
p07        [AYAT]
p07        | (2) Setiap
p07        | Pengelola
p07        | KTR
p07        | yang
p07        | tidak
p07        | melakukan
p07        | pengawasan internal, membiarkan orang Merokok,
p07        | tidak menyingkirkan asbak atau sejenisnya, dan tidak
p07        | memasang tanda dilarang Merokok di tempat atau area
p07        | yang dinyatakan sebagai KTR, dikenakan denda
p07 B      | administratif paling banyak Rp1.000.000,00 (satu juta
p07        | rupiah).
p07        [AYAT]
p07        | (3) Ketentuan lebih lanjut mengenai pengenaan denda
p07        | administratif sebagaimana dimaksud pada ayat (2)
p07        | diatur dalam Peraturan Bupati.
p07        [HEADING:BAGIAN]
p07        | Bagian Ketiga
p07        | Larangan dan Pengendalian
p07        | Pasal  11
p07        [AYAT]
p07        | (1) Setiap orang dilarang Merokok di KTR.
p07        [AYAT]
p07        | (2) Setiap orang dan/atau Badan dilarang mengiklankan,
p07        | mempromosikan,
p07        | memberikan
p07        | sponsor,
p07        | menjual,
p07        | dan/atau membeli Rokok di KTR.
p07        [AYAT]
p07        | (3) Larangan
p07        | menjual
p07        | dan
p07        | membeli
p07        | sebagaimana
p07        | dimaksud pada ayat (2) dikecualikan untuk Tempat
p07        | Umum yang memiliki izin untuk menjual Rokok.
p07        [AYAT]
p07        | (4) Setiap orang yang memiliki izin untuk menjual Rokok
p07        | di Tempat Umum sebagaimana dimaksud pada ayat (3)
p07        | dilarang untuk memperlihatkan atau memajang secara
p07        | jelas jenis dan produk Rokok.
p07        [AYAT]
p07        | (5) Pelanggaran
p07        | terhadap
p07        | ketentuan
p07        | sebagaimana
p07        | dimaksud pada ayat (1), dikenai sanksi administratif
p07        | berupa:
p07        [SUB-ITEM]
p07        | a. teguran lisan;
p07        [SUB-ITEM]
p07        | b. teguran tertulis; dan
p07        [SUB-ITEM]
p07        | c. denda  administratif paling sedikit Rp50.000,00
p07        | (lima puluh ribu rupiah) dan paling banyak
p07        | Rp200.000,00 (dua ratus ribu rupiah).
==================== PAGE 8 ====================
p08        | 8
p08        [AYAT]
p08        | (6) Pelanggaran
p08        | terhadap
p08        | ketentuan
p08        | sebagaimana
p08        | dimaksud pada ayat (2) dan ayat (4) dikenai sanksi
p08        | administratif berupa:
p08        [SUB-ITEM]
p08        | a. teguran lisan;
p08        [SUB-ITEM]
p08        | b. teguran tertulis;
p08        [SUB-ITEM]
p08        | c. penarikan produk;
p08        [SUB-ITEM]
p08        | d. denda administratif  paling banyak Rp1.000.000,00
p08        | (satu juta rupiah); dan
p08        [SUB-ITEM]
p08        | e. penghentian sementara kegiatan.
p08        [AYAT]
p08        | (7) Ketentuan lebih lanjut mengenai pengenaan sanksi
p08        | administratif sebagaimana dimaksud pada ayat (5) dan
p08        | ayat (6) diatur dalam Peraturan Bupati.
p08        | Pasal 12
p08        [AYAT]
p08        | (1) Setiap orang dilarang menjual Rokok:
p08        [SUB-ITEM]
p08        | a. menggunakan mesin layan diri;
p08        [SUB-ITEM]
p08        | b. kepada setiap orang di bawah usia 21 (dua puluh
p08        | satu) tahun dan perempuan hamil;
p08        [SUB-ITEM]
p08        | c. secara eceran satuan perbatang, kecuali bagi
p08        | produk tembakau berupa cerutu dan Rokok
p08        | elektronik;
p08        [SUB-ITEM]
p08        | d. dengan
p08        | menempatkan
p08        | produk
p08        | tembakau
p08        | dan
p08        | Rokok elektronik pada area sekitar pintu masuk
p08        | dan keluar atau pada tempat yang sering dilalui;
p08        [SUB-ITEM]
p08        | e. dalam radius 200 (dua ratus) meter dari satuan
p08        | pendidikan dan Tempat Anak Bermain; dan
p08        [SUB-ITEM]
p08        | f.
p08        | menggunakan
p08        | jasa
p08        | situs
p08        | web
p08        | atau
p08        | aplikasi
p08        | elektronik kemersial dan media sosial.
p08        [AYAT]
p08        | (2) Ketentuan sebagaimana dimaksud pada ayat (1) huruf
p08        | e dikecualikan bagi penjual rokok yang telah berjualan
p08        | sebelum Peraturan Daerah ini ditetapkan.
p08        [AYAT]
p08        | (3) Pelanggaran
p08        | terhadap
p08        | ketentuan
p08        | sebagaimana
p08        | dimaksud pada ayat (1) dikenai sanksi administratif
p08        | berupa:
p08        [SUB-ITEM]
p08        | a. teguran lisan;
p08        [SUB-ITEM]
p08        | b. teguran tertulis;
p08        [SUB-ITEM]
p08        | c. penarikan produk;
p08        [SUB-ITEM]
p08        | d. denda administratif  paling banyak Rp1.000.000,00
p08        | (satu juta rupiah); dan
p08        [SUB-ITEM]
p08        | e. penghentian sementara kegiatan.
p08        [AYAT]
p08        | (4) Ketentuan lebih lanjut mengenai pengenaan sanksi
p08        | administratif sebagaimana dimaksud pada ayat (3)
p08        | diatur dalam Peraturan Bupati.
p08        | Pasal 13
p08        [AYAT]
p08        | (1) Pemerintah Daerah melakukan pengendalian iklan
p08        | Rokok yang dilakukan pada media luar ruang.
p08        [AYAT]
p08        | (2) Pengendalian iklan Rokok pada media luar ruang:
p08        [SUB-ITEM]
p08        | a. tidak diletakkan di jalan utama dan jalan protokol;
p08        [SUB-ITEM]
p08        | b. tidak diletakkan dalam radius 500 (lima ratus)
p08        | meter di luar satuan pendidikan dan tempat
p08        | bermain anak; dan
==================== PAGE 9 ====================
p09        | 9
p09        [SUB-ITEM]
p09        | c. harus diletakkan sejajar dengan bahu jalan dan
p09        | tidak boleh memotong jalan atau melintang.
p09        | Pasal 14
p09        | Ketentuan
p09        | lebih
p09        | lanjut
p09        | mengenai
p09        | tanggung
p09        | jawab,
p09        | kewajiban,
p09        | larangan
p09        | dan
p09        | pengendalian
p09        | sebagaimana
p09        | dimaksud dalam Pasal 9, Pasal 10 ayat (1), Pasal 11 ayat
p09        [AYAT]
p09        | (1) sampai dengan ayat (4), Pasal 12 ayat (1), dan Pasal 13
p09        | diatur dalam Peraturan Bupati
p09        [HEADING:BAB]
p09        | BAB IV
p09        | PARTISIPASI MASYARAKAT
p09        | Pasal 15
p09        [AYAT]
p09        | (1) Masyarakat dapat berpatisipasi dalam mewujudkan
p09        | KTR dalam bentuk:
p09        [SUB-ITEM]
p09        | a. memberikan sumbangan pemikiran dan dengan
p09        | penentuan kebijakan yang terkait Rokok;
p09        [SUB-ITEM]
p09        | b. melakukan pengadaan dan pemberian bantuan
p09        | sarana dan prasarana yang diperlukan untuk
p09        | mewujudkan KTR;
p09        [SUB-ITEM]
p09        | c. ikut serta dalam memberikan bimbingan dan
p09        | penyuluhan
p09        | serta
p09        | penyebarluasan
p09        | informasi
p09        | kepada masyarakat;
p09        [SUB-ITEM]
p09        | d. mengingatkan
p09        | setiap
p09        | orang
p09        | yang
p09        | melakukan
p09        | kegiatan sebagaimana dimaksud dalam Pasal 11
p09        | ayat (1) sampai dengan ayat (4) melaporkannya
p09        | kepada pimpinan/penanggung jawab KTR; dan
p09        [SUB-ITEM]
p09        | e. melaporkan kepada pimpinan atau penanggung
p09        | jawab KTR jika terjadi pelanggaran.
p09        [AYAT]
p09        | (2) Pemberian sumbangan pemikiran dan pertimbangan
p09        | sebagaimana dimaksud pada ayat (1) huruf a, dapat
p09        | dilakukan langsung kepada/melalui Perangkat Daerah
p09        | terkait, atau secara tidak langsung dalam bentuk
p09        | penyelenggaraan
p09        | diskusi,
p09        | seminar
p09        | dan
p09        | kegiatan
p09        | sejenis, dan/atau melalui media komunikasi, dalam
p09        | bentuk:
p09        [SUB-ITEM]
p09        | a.   cetak;
p09        [SUB-ITEM]
p09        | b. elektronik; dan
p09        [SUB-ITEM]
p09        | c. bentuk lainnya.
p09        [AYAT]
p09        | (3) Bantuan masyarakat berupa sarana/prasarana yang
p09        | diperlukan untuk mewujudkan KTR sebagaimana
p09        | dimaksud pada ayat (1) huruf b dapat dilakukan
p09        | secara
p09        | langsung
p09        | kepada
p09        | pimpinan
p09        | dan/atau
p09        | penanggung jawab KTR sesuai ketentuan peraturan
p09        | perundang-undangan.
p09        [AYAT]
p09        | (4) Peran serta masyarakat dalam mewujudkan KTR
p09        | sebagaimana dimaksud pada ayat (2) dapat dilakukan
p09        | secara
p09        | berkelompok/institusional/badan
p09        | hukum/
p09        | badan usaha/lembaga/organisasi maupun individu/
p09        | perorangan.
==================== PAGE 10 ====================
p10        | 10
p10        [AYAT]
p10        | (5) Pimpinan
p10        | atau
p10        | penanggung
p10        | jawab
p10        | KTR
p10        | wajib
p10        | menindaklanjuti laporan sebagaimana dimaksud pada
p10        | ayat (1) huruf e.
p10        [AYAT]
p10        | (6) Ketentuan
p10        | lebih
p10        | lanjut
p10        | mengenai
p10        | partisipasi
p10        | masyarakat sebagaimana dimaksud pada ayat (1)
p10        | diatur dalam Peraturan Bupati.
p10        [HEADING:BAB]
p10        | BAB V
p10        | SATUAN TUGAS KTR
p10        | Pasal 16
p10        [AYAT]
p10        | (1) Bupati membentuk satuan tugas KTR yang susunan
p10        | keanggotannya terdiri dari Perangkat Daerah dan
p10        | unsur terkait lainnya.
p10        [AYAT]
p10        | (2) Tugas satuan tugas KTR sebagaimana dimaksud pada
p10        | ayat (1) meliputi:
p10        [SUB-ITEM]
p10        | a. menyusun rencana kerja pelaksanaan pengawasan
p10        | terhadap KTR;
p10        [SUB-ITEM]
p10        | b. menginventarisasi fasilitas pelayanan kesehatan,
p10        | Tempat Proses Belajar Mengajar, Tempat Anak
p10        | Bermain,
p10        | Tempat
p10        | Ibadah,
p10        | Angkutan
p10        | Umum,
p10        | Tempat Kerja Tertentu, dan Tempat Umum yang
p10        | merupakan KTR;
p10        [SUB-ITEM]
p10        | c. melakukan
p10        | berbagai
p10        | upaya
p10        | dalam
p10        | rangka
p10        | meningkatkan kepatuhan penerapan KTR;
p10        [SUB-ITEM]
p10        | d. mendorong penanggung jawab kawasan untuk
p10        | membentuk tim pengawas KTR dan merumuskan
p10        | petunjuk teknis penegakan KTR pada kawasan
p10        | masing-masing dan unit di bawahnya;
p10        [SUB-ITEM]
p10        | e. mengendalikan
p10        | iklan,
p10        | promosi,
p10        | dan
p10        | sponsor
p10        | tentang Rokok pada KTR;
p10        [SUB-ITEM]
p10        | f. melaksanakan
p10        | pengawasan,
p10        | pemantauan,
p10        | pembinaan, dan evaluasi terhadap KTR;
p10        [SUB-ITEM]
p10        | g. membantu penanggung jawab kawasan dalam
p10        | memproses setiap pelanggaran yang terjadi pada
p10        | saat melakukan pengawasan; dan
p10        [SUB-ITEM]
p10        | h. melaporkan hasil pelaksanaan pengawasan KTR
p10        | kepada Bupati setiap tahun melalui Perangkat
p10        | Daerah
p10        | yang
p10        | menyelenggarakan
p10        | urusan
p10        | pemerintahan di bidang Kesehatan.
p10        [AYAT]
p10        | (3) Perangkat Daerah sebagaimana dimaksud pada ayat
p10        [AYAT]
p10        | (1) meliputi:
p10        [SUB-ITEM]
p10        | a. Perangkat Daerah yang menyelenggarakan urusan
p10        | pemerintahan
p10        | di
p10        | bidang
p10        | pendidikan
p10        | dan
p10        | kebudayaan;
p10        [SUB-ITEM]
p10        | b. Perangkat Daerah yang menyelenggarakan urusan
p10        | pemerintahan di bidang kesehatan;
p10        [SUB-ITEM]
p10        | c. Perangkat Daerah yang menyelenggarakan urusan
p10        | pemerintahan di bidang pekerjaan umum dan
p10        | penataan ruang;
p10        [SUB-ITEM]
p10        | d. Perangkat Daerah yang menyelenggarakan urusan
p10        | pemerintahan di bidang perumahan dan kawasan
p10        | permukiman;
==================== PAGE 11 ====================
p11        | 11
p11        [SUB-ITEM]
p11        | e. Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan
p11        | di
p11        | bidang
p11        | ketentraman
p11        | dan
p11        | ketertiban umum serta perlindungan masyarakat
p11        | sub urusan ketentraman dan ketertiban umum dan
p11        | sub urusan kebakaran;
p11        [SUB-ITEM]
p11        | f.
p11        | Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang sosial;
p11        [SUB-ITEM]
p11        | g. Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang pengendalian penduduk
p11        | dan keluarga berencana dan bidang pemberdayaan
p11        | perempuan dan perlindungan anak;
p11        [SUB-ITEM]
p11        | h. Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang pangan;
p11        [SUB-ITEM]
p11        | i.
p11        | Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang lingkungan hidup;
p11        [SUB-ITEM]
p11        | j.
p11        | Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan
p11        | di
p11        | bidang
p11        | administrasi
p11        | kependudukan dan catatan sipil;
p11        [SUB-ITEM]
p11        | k. Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang pemberdayaan masyarakat
p11        | dan desa;
p11        [SUB-ITEM]
p11        | l.
p11        | Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang perhubungan;
p11        [SUB-ITEM]
p11        | m. Perangkat Daerah yang menyelenggarakan urusan
p11        | Pemerintahan bidang komunikasi dan informatika,
p11        | bidang persandian, dan bidang statistik;
p11        [SUB-ITEM]
p11        | n. Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang penanaman modal dan
p11        | pelayanan terpadu satu pintu;
p11        [SUB-ITEM]
p11        | o. Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang kepemudaan dan olahraga
p11        | dan bidang pariwisata;
p11        [SUB-ITEM]
p11        | p. Perangkat Daerah yang menyelenggarakan urusan
p11        | Pemerintahan di bidang perpustakaan dan bidang
p11        | kearsipan;
p11        [SUB-ITEM]
p11        | q. Perangkat Daerah yang menyelenggarakan urusan
p11        | pemerintahan di bidang pertanian dan bidang
p11        | perikanan;
p11        [SUB-ITEM]
p11        | r. Perangkat Daerah yang menyelenggarakan urusan
p11        | Pemerintahan di bidang koperasi, usaha kecil dan
p11        | menengah, dan bidang perdagangan; dan
p11        [SUB-ITEM]
p11        | s. Perangkat Daerah yang menyelenggarakan urusan
p11        | Pemerintahan di bidang perindustrian, bidang
p11        | tenaga kerja dan bidang transmigrasi.
p11        [AYAT]
p11        | (4) Satuan tugas KTR sebagaimana dimaksud pada ayat
p11        [AYAT]
p11        | (1)  ditetapkan dengan Keputusan Bupati.
==================== PAGE 12 ====================
p12        | 12
p12        [HEADING:BAB]
p12        | BAB VI
p12        | PEMBINAAN DAN PENGAWASAN
p12        [HEADING:BAGIAN]
p12        | Bagian Kesatu
p12        | Pembinaan
p12        | Pasal 17
p12        [AYAT]
p12        | (1) Bupati melakukan pembinaan terhadap penataan dan
p12        | pengelolaan KTR.
p12        [AYAT]
p12        | (2) Pembinaan sebagaimana dimaksud pada ayat (1)
p12        | dilakukan Perangkat Daerah yang menyelenggarakan
p12        | urusan pemerintahan di bidang Kesehatan.
p12        [AYAT]
p12        | (3) Pembinaan sebagaimana dimaksud pada ayat (1)
p12        | meliputi:
p12        [SUB-ITEM]
p12        | a. penyebarluasan informasi dan sosialisasi tentang
p12        | KTR;
p12        [SUB-ITEM]
p12        | b. koordinasi di bidang penataan dan pengelolaan
p12        | KTR dengan seluruh lembaga pemerintah dan non-
p12        | pemerintah;
p12        [SUB-ITEM]
p12        | c. memberikan motivasi tidak Merokok dalam KTR;
p12        [SUB-ITEM]
p12        | d. perumusan kebijakan; dan
p12        [SUB-ITEM]
p12        | e. bekerja sama di bidang penataan dan pengelolaan
p12        | dengan lembaga pemerintah dan non-pemerintah,
p12        | baik nasional maupun internasional.
p12        [HEADING:BAGIAN]
p12        | Bagian Kedua
p12        | Pengawasan
p12        | Pasal 18
p12        [AYAT]
p12        | (1) Bupati melakukan pengawasan terhadap pelaksanaan
p12        | KTR.
p12        [AYAT]
p12        | (2) Pengawasan sebagaimana dimaksud pada ayat (1)
p12        | dilaksanakan
p12        | Perangkat
p12        | Daerah
p12        | sebagaimana
p12        | dimaksud dalam Pasal 16 ayat (3).
p12        [HEADING:BAB]
p12        | BAB VII
p12        | PENDANAAN
p12        | Pasal 19
p12        [AYAT]
p12        | (1) Pendanaan pelaksanaan KTR di Daerah bersumber
p12        | dari anggaran pendapatan dan belanja Daerah.
p12        [AYAT]
p12        | (2) Selain bersumber dari anggaran pendapatan dan
p12        | belanja Daerah sebagaimana dimaksud pada ayat (1)
p12        | pendanaan dapat bersumber dari sumber lain yang
p12        | sah dan tidak mengikat sesuai dengan ketentuan
p12        | peraturan perundang-undangan.
==================== PAGE 13 ====================
p13        | 13
p13        | Salinan sesuai dengan aslinya
p13        | KEPALA BAGIAN HUKUM,
p13        | TEGUH PRAMONO,SH,MH
p13        | Pembina Tingkat I
p13        | NIP. 19710429 199803 1 003
p13        [HEADING:BAB]
p13        | BAB VIII
p13        | KETENTUAN PENUTUP
p13        | Pasal 20
p13        | Pada saat Peraturan Daerah ini mulai berlaku, ketentuan
p13        | mengenai Tertib Kawasan Tanpa Rokok dalam BAB V
p13        | Pasal 10 sampai dengan Pasal 14 dalam Peraturan Daerah
p13        | Nomor
p13        | 3
p13        | Tahun
p13        | 2014
p13        | tentang
p13        | Ketertiban
p13        | Umum
p13        | (Lembaran Daerah Kabupaten Sukoharjo Tahun 2014
p13        | Nomor
p13        | 3,
p13        | Tambahan
p13        | Lembaran
p13        | Daerah
p13        | Kabupaten
p13        | Sukoharjo Nomor 210), dicabut dan dinyatakan tidak
p13        | berlaku.
p13        | Pasal 21
p13        | Peraturan Pelaksanaan Peraturan Daerah ini harus
p13        | ditetapkan paling lama 1 (satu) tahun terhitung sejak
p13        | Peraturan Daerah ini diundangkan.
p13        | Pasal 22
p13        | Peraturan
p13        | Daerah
p13        | ini
p13        | mulai
p13        | berlaku
p13        | pada
p13        | tanggal
p13        | diundangkan.
p13        | Agar
p13        | setiap
p13        | orang
p13        | mengetahuinya,
p13        | memerintahkan
p13        | pengundangan
p13        | Peraturan
p13        | Daerah
p13        | ini
p13        | dengan
p13        | penempatannya dalam Lembaran Daerah Kabupaten
p13        | Sukoharjo.
p13        | Ditetapkan di Sukoharjo
p13        | pada tanggal 23 Juli 2025
p13        | BUPATI SUKOHARJO,
p13        | ttd.
p13        | ETIK SURYANI
p13        | Diundangkan di Sukoharjo
p13        | pada tanggal 23 Juli 2025
p13        I1 | Pj. SEKRETARIS DAERAH
p13        | KABUPATEN SUKOHARJO,
p13        I1 | ttd.
p13        I1 | SUYAMTO
p13        I2 | LEMBARAN DAERAH KABUPATEN SUKOHARJO TAHUN 2025 NOMOR 1
p13        I1 | NOMOR REGISTER PERATURAN DAERAH KABUPATEN SUKOHARJO,
p13        I1 | PROVINSI JAWA TENGAH : (1-60/2025)
==================== PAGE 14 ====================
p14        | 14
p14        | PENJELASAN
p14        | ATAS
p14        | PERATURAN DAERAH KABUPATEN SUKOHARJO
p14        | NOMOR 1 TAHUN 2025
p14        | TENTANG
p14        | KAWASAN TANPA ROKOK
p14        I2 | I. UMUM
p14        | Rokok mengandung zat adiktif yang sangat berbahaya bagi kesehatan
p14        I2 | manusia. Hal tersebut dinyatakan secara tegas di dalam Pasal 149 Undang-
p14        I2 | Undang Nomor 17 Tahun 2023 tentang Kesehatan, bahwa produk
p14        I2 | tembakau merupakan zat adiktif. Zat adiktif merupakan zat yang jika
p14        I2 | dikonsumsi manusia akan menimbulkan adiksi atau ketagihan, dan dapat
p14        I2 | memicu timbulnya berbagai penyakit seperti penyakit jantung dan
p14        I2 | pembuluh darah, stroke, penyakit paru obstruktif kronik, kanker paru,
p14        I2 | kanker mulut, impotensi,  serta kelainan kehamilan dan janin. Hal tersebut
p14        I2 | karena di dalam Rokok yang dibakar terdapat lebih dari 4.000 (empat ribu)
p14        I2 | zat kimia antara lain nikotin yang bersifat adiktif dan tar yang bersifat
p14        I2 | karsinogenik. Asap Rokok tidak hanya membahayakan perokok, tetapi juga
p14        I2 | orang lain yang berada di sekitar perokok (perokok pasif).
p14        | Asap Rokok pasif merupakan zat sangat kompleks berisi campuran gas
p14        I2 | dan partikel halus yang dikeluarkan dari pembakaran Rokok. Asap Rokok
p14        I2 | orang lain sangat berbahaya bagi orang yang tidak merokok yang
p14        I2 | menghirup asap Rokok yang dihisap orang lain. Perokok pasif menanggung
p14        I2 | risiko sama tingginya dengan orang yang merokok. Sehingga tidak ada
p14        I2 | batas aman untuk paparan asap Rokok orang lain. Bahaya asap orang lain
p14        I2 | juga dihadapi oleh bayi dalam kandungan ibu yang merokok dan orang-
p14        I2 | orang yang berada dalam ruangan yang terdapat asap rokok yang telah
p14        I2 | ditinggalkan perokok. Ibu hamil yang merokok selama kehamilan akan
p14        I2 | mempengaruhi pertumbuhan bayi yang menyebabkan berat badan lahir
p14        I2 | rendah (BBLR) kelahiran prematur, dan kematian.
p14        | Berdasarkan hal tersebut maka Undang-Undang Nomor 17 Tahun 2023
p14        I2 | tentang Kesehatan mengamanatkan Pemerintah Daerah untuk mengatur
p14        I2 | penetapan Kawasan Tanpa Rokok. Pengaturan ini bertujuan untuk
p14        I2 | mencegah dan mengatasi dampak buruk asap Rokok. Berdasarkan
p14        I2 | ketentuan Pasal 151 ayat (2) Undang-Undang Nomor 17 Tahun 2023
p14        I2 | tentang
p14        | Kesehatan,
p14        | menentukan
p14        | bahwa
p14        | Pemerintah
p14        | Daerah
p14        | wajib
p14        I2 [PREAMBLE:MENETAPKAN]
p14        I2 | menetapkan
p14        | dan
p14        | mengimplementasikan
p14        | Kawasan
p14        | Tanpa
p14        | Rokok
p14        | di
p14        I2 | wilayahnya.
p14        | Kawasan
p14        | Tanpa
p14        | Rokok
p14        | mencakup
p14        | Fasilitas
p14        | Pelayanan
p14        I2 | Kesehatan, Tempat Proses Belajar-Mengajar, Tempat Anak Bermain, Tempat
p14        I2 | Ibadah, Angkutan Umum, Tempat Kerja Tertentu, dan Tempat Umum yang
p14        I2 | ditetapkan. Peraturan Daerah ini melarang kegiatan merokok, iklan Rokok
p14        I2 | dan penjualan Rokok di Kawasan Tanpa Rokok yang telah diuraikan
p14        I2 | sebelumnya kecuali di Tempat Umum, masih diperbolehkan transaksi jual
p14        I2 | beli Rokok.
p14        | Kawasan Tanpa Rokok merupakan tanggung jawab seluruh komponen
p14        I2 | bangsa, baik individu, masyarakat, lembaga-lembaga pemerintah dan non-
p14        I2 | pemerintah, untuk melindungi hak-hak generasi sekarang maupun yang
p14        I2 | akan datang atas Kesehatan diri dan lingkungan hidup yang sehat.
==================== PAGE 15 ====================
p15        | 15
p15        I2 | Komitmen bersama lintas sektor dan berbagai elemen akan sangat
p15        I2 | berpengaruh terhadap keberhasilan kawasan tanpa Rokok.
p15        I2 | II. PASAL DEMI PASAL
p15        | Pasal 1
p15        | Cukup jelas.
p15        | Pasal 2
p15        | Huruf a
p15        | Yang
p15        | dimaksud
p15        | dengan
p15        | asas
p15        | “kepentingan
p15        | kualitas
p15        | Kesehatan manusia” yaitu bahwa pelaksanaan KTR harus
p15        | dilaksanakan
p15        | berdasarkan
p15        | atas
p15        | kepentingan
p15        | kualitas
p15        | Kesehatan manusia.
p15        | Huruf b
p15        | Yang dimaksud asas “keseimbangan” yaitu pelaksanaan KTR
p15        | harus dilaksanakan antara kepentingan individu dan
p15        | masyarakat.
p15        | Huruf c
p15        | Yang dimaksud asas “kemanfaatan” yaitu pelaksanaan KTR
p15        | harus memberikan manfaat yang sebesar-besarnya bagi
p15        | kemanusiaan
p15        | dan
p15        | perikehidupan
p15        | yang
p15        | sehat
p15        | bagi
p15        | masyarakat.
p15        | Huruf d
p15        | Yang
p15        | dimaksud
p15        | dengan
p15        | asas
p15        | “keterpaduan”
p15        | yaitu
p15        | pelaksanaan
p15        | KTR
p15        | harus
p15        | ada
p15        | keterpaduan
p15        | antara
p15        | kepentingan pemerintah, individu, dan masyarakat.
p15        | Huruf e
p15        | Yang dimaksud dengan asas “keserasian” yaitu pelaksanaan
p15        | KTR harus ada keserasian antara pemerintah, individu, dan
p15        | masyarakat.
p15        | Huruf f
p15        | Yang dimaksud dengan “partisipasi” yaitu pelaksaan KTR
p15        | harus melibatkan partisipasi masyarakat.
p15        I1 | Huruf g
p15        | Yang dimaksud dengan “keadilan” yaitu penyelenggaraan
p15        | KTR harus dapat memberikan pelayanan yang adil dan
p15        | merata kepada semua lapisan masyarakat.
p15        | Huruf h
p15        | Yang dimaksud dengan “transparansi dan akuntabilitas”
p15        | yaitu
p15        | pelaksanaan
p15        | KTR
p15        | harus
p15        | dilaksanakan
p15        | secara
p15        | transparan dan akuntabel. Artinya bahwa masyarakat dapat
p15        | dengan
p15        | mudah
p15        | untuk
p15        | mengakses
p15        | dan
p15        | mendapatkan
p15        | informasi tentang KTR, dan dapat dipertanggungjawabkan
p15        | sesuai dengan peraturan perundang undangan.
p15        | Pasal 3
p15        | Cukup jelas.
p15        | Pasal 4
p15        | Cukup jelas.
==================== PAGE 16 ====================
p16        | 16
p16        | Pasal 5
p16        | Cukup jelas.
p16        | Pasal 6
p16        | Cukup jelas.
p16        | Pasal 7
p16        | Ayat (1)
p16        | Cukup jelas.
p16        | Ayat (2)
p16        | Penyediaan tempat khusus untuk Merokok disesuaikan
p16        | dengan kewenangan masing-masing.
p16        | Ayat (3)
p16        | Cukup jelas.
p16        | Ayat (4)
p16        | Cukup jelas.
p16        | Ayat (5)
p16        | Cukup jelas.
p16        | Ayat (6)
p16        | Cukup jelas.
p16        | Pasal 8
p16        | Cukup jelas.
p16        | Pasal 9
p16        | Cukup jelas.
p16        | Pasal 10
p16        | Cukup jelas.
p16        | Pasal 11
p16        | Cukup jelas.
p16        | Pasal 12
p16        | Cukup jelas.
p16        | Pasal 13
p16        | Cukup jelas.
p16        | Pasal 14
p16        | Cukup jelas.
p16        | Pasal 15
p16        | Ayat (1)
p16        | Huruf a
p16        | Cukup jelas.
==================== PAGE 17 ====================
p17        | 17
p17        | Huruf b
p17        | Yang dimaksud dengan “sarana prasarana yang
p17        | diperlukan untuk mewujudkan KTR” berupa:
p17        [SUB-ITEM]
p17        | a. papan dan rambu larangan Merokok dan informasi
p17        | di area KTR;
p17        [SUB-ITEM]
p17        | b. papan dan rambu area Merokok;
p17        [SUB-ITEM]
p17        | c. tempat sampah khusus puntung rokok di area
p17        | Merokok;
p17        [SUB-ITEM]
p17        | d. media sosialisasi dan edukasi berupa poster,
p17        | banner, leaflet tentang kesehatan dan bahaya
p17        | Merokok;
p17        [SUB-ITEM]
p17        | e. seragam dan tanda pengenal satuan tugas KTR;
p17        | dan
p17        [SUB-ITEM]
p17        | f. blangko
p17        | teguran/berita
p17        | acara
p17        | jika
p17        | terjadi
p17        | pelanggaran.
p17        | Huruf c
p17        | Cukup jelas.
p17        | Huruf d
p17        | Yang
p17        | dimaksud
p17        | dengan
p17        | “mengingatkan”
p17        | yaitu
p17        | memberi pengetahuan kepada setiap orang yang
p17        | melakukan pelanggaran secara persuasif atau secara
p17        | baik, dengan nada ramah dan informatif.
p17        | Huruf e
p17        | Cukup jelas
p17        | Ayat (2)
p17        | Cukup jelas.
p17        | Ayat (3)
p17        | Cukup jelas.
p17        | Ayat (4)
p17        | Cukup jelas.
p17        | Ayat (5)
p17        | Cukup jelas.
p17        | Ayat (6)
p17        | Cukup jelas.
p17        | Pasal 16
p17        | Ayat (1)
p17        | Yang dimaksud dengan “unsur terkait lainnya” terdiri  atas:
p17        [SUB-ITEM]
p17        | a. Kementerian Agama Kabupaten Sukoharjo;
p17        [SUB-ITEM]
p17        | b. Kejaksaan Negeri Sukoharjo;
p17        [SUB-ITEM]
p17        | c. Kepolisian Resor Sukoharjo; dan
p17        [SUB-ITEM]
p17        | d. Komando Distrik Militer 0726 Sukoharjo.
p17        | Ayat (2)
p17        | Cukup jelas.
==================== PAGE 18 ====================
p18        | 18
p18        | Ayat (3)
p18        | Cukup jelas.
p18        | Ayat (4)
p18        | Cukup jelas.
p18        | Pasal 17
p18        | Cukup jelas.
p18        | Pasal 18
p18        | Cukup jelas.
p18        | Pasal 19
p18        | Cukup jelas.
p18        | Pasal 20
p18        | Cukup jelas.
p18        | Pasal 21
p18        | Cukup jelas.
p18        | Pasal 22
p18        | Cukup jelas.
p18        I1 | TAMBAHAN LEMBARAN DAERAH KABUPATEN SUKOHARJO NOMOR 330
```

---


## keppres

- **File**: `keppres/keppres-no-5-tahun-2015_Dewan Kawasan Kawasan Ekonomi Khusus Provinsi Kalimantan Timur.pdf`
- **Document Type**: Keputusan Presiden (Decision)
- **Issued by**: Presiden
- **Pages**: 5 | **Lines**: 193
- **Font sizes**: [12.0]
- **Most common font**: 12.0 (100% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [72.0, 119.0, 148.0, 164.0, 192.0, 230.0, 246.0, 387.0]
- **Expected hierarchy**: Consideranda > MEMUTUSKAN > Items

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01        I3 | KEPUTUSAN PRESIDEN REPUBLIK INDONESIA
p01        I6 | NOMOR  5  TAHUN  2015
p01        | TENTANG
p01        I6 | DEWAN KAWASAN
p01        I2 | KAWASAN EKONOMI KHUSUS PROVINSI KALIMANTAN TIMUR
p01        I4 [PREAMBLE:DENGAN RAHMAT]
p01        I4 | DENGAN RAHMAT TUHAN YANG MAHA ESA
p01        I5 | PRESIDEN REPUBLIK INDONESIA,
p01        I1 [PREAMBLE:MENIMBANG]
p01        I1 | Menimbang :
p01        I3 [SUB-ITEM]
p01        I3 | a.
p01        I5 | bahwa
p01        I6 | dalam
p01        | rangka
p01        | mempercepat
p01        | pembangunan
p01        I5 | perekonomian di wilayah Provinsi Kalimantan Timur dan
p01        I5 | untuk
p01        I6 | menunjang
p01        | percepatan
p01        | dan
p01        | perluasan
p01        I5 | pembangunan
p01        | ekonomi
p01        I8 | nasional
p01        | telah
p01        | ditetapkan
p01        I5 | Peraturan Pemerintah Nomor 85 Tahun 2014 tentang
p01        I5 | Kawasan
p01        I7 | Ekonomi
p01        | Khusus
p01        I8 | Maloy
p01        | Batuta
p01        | Trans
p01        I5 | Kalimantan;
p01        I3 [SUB-ITEM]
p01        I3 | b.  bahwa dalam rangka penyelenggaraan pengembangan
p01        I5 | Kawasan Ekonomi Khusus sebagaimana dimaksud pada
p01        I5 | huruf a dan berdasarkan Pasal 19 ayat (2) Undang-
p01        I5 | Undang Nomor 39 Tahun 2009 tentang Kawasan Ekonomi
p01        I5 | Khusus,
p01        I7 | perlu
p01        | dibentuk
p01        I8 | Dewan
p01        | Kawasan
p01        | Kawasan
p01        I5 | Ekonomi Khusus Provinsi Kalimantan Timur dengan
p01        I5 | Keputusan Presiden;
p01        I3 [SUB-ITEM]
p01        I3 | c.  bahwa berdasarkan pertimbangan sebagaimana dimaksud
p01        I5 | pada huruf a dan huruf b, perlu menetapkan Keputusan
p01        I5 | Presiden tentang Dewan Kawasan Kawasan Ekonomi
p01        I5 | Khusus Provinsi Kalimantan Timur;
p01        I1 [PREAMBLE:MENGINGAT]
p01        I1 | Mengingat
p01        I3 | :
p01        I3 [ITEM]
p01        I3 | 1.
p01        I5 | Pasal 4 ayat (1) Undang-Undang Dasar Republik Indonesia
p01        I5 | Tahun 1945;
p01        I3 [ITEM]
p01        I3 | 2.  Undang-Undang Nomor 39 Tahun 2009 tentang Kawasan
p01        I5 | Ekonomi Khusus (Lembaran Negara Republik Indonesia
p01        I5 | Tahun 2009 Nomor 147, Tambahan Lembaran Negara
p01        I5 | Republik Indonesia Nomor 5066);
p01        [ITEM]
p01        | 3. Peraturan ...
==================== PAGE 2 ====================
p02        | - 2 -
p02        I3 [ITEM]
p02        I3 | 3.  Peraturan Pemerintah Nomor 2 Tahun 2011 tentang
p02        I5 | Penyelengaraan Kawasan Ekonomi Khusus (Lembaran
p02        I5 | Negara Republik Indonesia Tahun 2011 Nomor 3,
p02        I5 | Tambahan Lembaran Negara Republik Indonesia Nomor
p02        I5 | 5186) sebagaimana telah diubah dengan Peraturan
p02        I5 | Pemerintah Nomor 100 Tahun 2012 tentang Perubahan
p02        I5 | atas Peraturan Pemerintah Nomor 2 Tahun 2011 tentang
p02        I5 | Penyelengaraan Kawasan Ekonomi Khusus (Lembaran
p02        I5 | Negara Republik Indonesia Tahun 2012 Nomor 263,
p02        I5 | Tambahan Lembaran Negara Republik Indonesia Nomor
p02        I5 | 5371);
p02        I3 [ITEM]
p02        I3 | 4.  Peraturan Pemerintah Nomor 85 Tahun 2014 tentang
p02        I5 | Kawasan
p02        I7 | Ekonomi
p02        | Khusus
p02        I8 | Maloy
p02        | Batuta
p02        | Trans
p02        I5 | Kalimantan (Lembaran Negara Republik Indonesia Tahun
p02        I5 | 2014 Nomor 306, Tambahan Lembaran Negara Republik
p02        I5 | Indonesia Nomor 5611);
p02        I3 [ITEM]
p02        I3 | 5.  Peraturan Presiden Nomor 33 Tahun 2010 tentang Dewan
p02        I5 | Nasional dan Dewan Kawasan Kawasan Ekonomi Khusus
p02        I5 | sebagaimana telah diubah beberapa kali terakhir dengan
p02        I5 | Peraturan Presiden Nomor 150 Tahun 2014 tentang
p02        I5 | Perubahan Kedua atas Peraturan Presiden Nomor 33
p02        I5 | Tahun 2010 tentang Dewan Nasional dan Dewan Kawasan
p02        I5 | Kawasan Ekonomi Khusus (Lembaran Negara Republik
p02        I5 | Indonesia Tahun 2014 Nomor 277);
p02        I7 [KEPUTUSAN:MEMUTUSKAN]
p02        I7 | MEMUTUSKAN :
p02        I1 [PREAMBLE:MENETAPKAN]
p02        I1 | Menetapkan :
p02        I3 | KEPUTUSAN
p02        I7 | PRESIDEN
p02        | TENTANG
p02        | DEWAN
p02        | KAWASAN
p02        I3 | KAWASAN
p02        I6 | EKONOMI
p02        | KHUSUS
p02        I8 | PROVINSI
p02        | KALIMANTAN
p02        I3 | TIMUR.
p02        | Pasal 1 ...
==================== PAGE 3 ====================
p03        | - 3 -
p03        | Pasal 1
p03        I3 [PREAMBLE:MENETAPKAN]
p03        I3 | Menetapkan Dewan Kawasan Kawasan Ekonomi Khusus
p03        I3 | Provinsi Kalimantan Timur, yang selanjutnya disebut Dewan
p03        I3 | Kawasan, dengan susunan keanggotaan sebagai berikut:
p03        I3 | Ketua merangkap
p03        | : Gubernur Kalimantan Timur;
p03        I3 | Anggota
p03        I3 | Wakil Ketua
p03        | :  Bupati Kutai Timur;
p03        I3 | merangkap Anggota
p03        I3 | Anggota
p03        | : 1. Kepala Kantor Wilayah Badan
p03        | Pertanahan
p03        | Nasional
p03        | Provinsi
p03        | Kalimantan Timur;
p03        [ITEM]
p03        | 2. Kepala Kantor Wilayah Direktorat
p03        | Jenderal
p03        | Bea
p03        | dan
p03        | Cukai
p03        | Kalimantan Bagian Timur;
p03        [ITEM]
p03        | 3. Kepala Kantor Wilayah Direktorat
p03        | Jenderal
p03        | Pajak
p03        | Kalimantan
p03        | Timur;
p03        [ITEM]
p03        | 4. Kepala
p03        I8 | Badan
p03        | Perencanaan
p03        | Pembangunan Daerah Provinsi
p03        | Kalimantan Timur;
p03        [ITEM]
p03        | 5. Kepala
p03        I8 | Dinas
p03        | Perhubungan
p03        | Provinsi Kalimantan Timur;
p03        [ITEM]
p03        | 6. Kepala Dinas Pekerjaan Umum
p03        | Provinsi Kalimantan Timur;
p03        [ITEM]
p03        | 7. Kepala ...
==================== PAGE 4 ====================
p04        | - 4 -
p04        [ITEM]
p04        | 7. Kepala Dinas Pekerjaan Umum
p04        | Kabupaten Kutai Timur;
p04        [ITEM]
p04        | 8. Kepala
p04        I8 | Dinas
p04        | Perhubungan,
p04        | Komunikasi,
p04        | dan
p04        | Informatika
p04        | Kabupaten Kutai Timur;
p04        [ITEM]
p04        | 9. Kepala Dinas Perindustrian dan
p04        | Perdagangan Kabupaten Kutai
p04        | Timur.
p04        | Pasal 2
p04        I3 | Dewan Kawasan bertanggung jawab dan melaporkan hasil
p04        I3 | pelaksanaan tugasnya kepada Dewan Nasional Kawasan
p04        I3 | Ekonomi Khusus paling kurang 1 (satu) kali dalam 6 (enam)
p04        I3 | bulan atau sewaktu-waktu bila diperlukan.
p04        | Pasal 3
p04        I3 | Segala biaya yang diperlukan untuk pelaksanaan tugas
p04        I3 | Dewan Kawasan dibebankan pada Anggaran Pendapatan dan
p04        I3 | Belanja Daerah Provinsi Kalimantan Timur dan sumber lain
p04        I3 | yang
p04        I5 | tidak
p04        I7 | bertentangan
p04        | dengan
p04        I8 | ketentuan
p04        | peraturan
p04        I3 | perundang-undangan.
p04        | Pasal 4 ...
==================== PAGE 5 ====================
p05        | - 5 -
p05        | Pasal 4
p05        I3 | Keputusan
p05        I6 | Presiden
p05        | ini
p05        | mulai
p05        I8 | berlaku
p05        | pada
p05        | tanggal
p05        I3 | ditetapkan.
p05        | Ditetapkan di Jakarta
p05        | pada tanggal 11 Februari 2015
p05        | PRESIDEN REPUBLIK INDONESIA,
p05        | ttd.
p05        I8 | JOKO WIDODO
p05        I1 | Salinan sesuai dengan aslinya
p05        I1 | SEKRETARIAT KABINET RI
p05        I1 | Deputi Bidang Perekonomian,
p05        I3 | ttd.
p05        I2 | Ratih Nurdiati
```

---


## inpres

- **File**: `inpres/inpres-no-3-tahun-2023_Percepatan Peningkatan Konektivitas Jalan Daerah.pdf`
- **Document Type**: Instruksi Presiden (Instruction)
- **Issued by**: Presiden
- **Pages**: 6 | **Lines**: 232
- **Font sizes**: [10.0, 10.5, 11.0, 11.5, 12.0, 12.5, 13.0, 13.5, 14.0, 14.5, 15.0, 15.5, 16.0, 16.5, 17.0, 17.5, 18.0, 22.5]
- **Most common font**: 18.0 (20% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [41.0, 65.0, 107.0, 150.0, 195.0, 272.0, 348.0, 368.0, 386.0, 414.0, 453.0]
- **Expected hierarchy**: Consideranda > INSTRUKSI > Numbered

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01   F22  I10 | SALINAN
p01   F12  I6 | PRESIDEN
p01   F12  | REPUBLIK INDONESIA
p01   F12  I4 | INSTRUKSI PRESIDEN REPUBLIK INDONESIA
p01   F14  | NOMOR 3 TAHUN 2023
p01   F12  I6 | TENTANG
p01   F13  I3 | PERCEPATAN PENINGKATAN KONEKTIVITAS JALAN DAERAH
p01   F12  I5 | PRESIDEN REPUBLIK INDONESIA,
p01   F17  I2 | Dalam rangka percepatan peningkatan konektivitas jalan daerah untuk
p01   F14  I2 | memberikan manfaat maksimal dalam mendorong perekonomian nasional
p01   F14  I2 | maupun daerah, menurunkan biaya logistik nasional, menghubungkan dan
p01   F14  I2 | mengintegrasikan dengan sentra-sentra ekonomi, dan membantu pemerataan
p01   F16  I2 | kondisi jalan yang mantap, sebagai upaya mendukung pencapaian target
p01   F14  I2 | Rencana Pembangunan Jangka Menengah Nasional Tahun 2O2O-2O24, dengan
p01   F14  I2 | ini menginstruksikan:
p01   F12  I2 | Kepada
p01        I4 [ITEM]
p01        I4 | 1. Menteri Perencanaan Pembangunan Nasional/
p01   F13  I5 | Kepala Badan Perencanaan Pembangunan Nasional;
p01        I4 [ITEM]
p01        I4 | 2. Menteri Pekerjaan Umum dan Perumahan Ralryat;
p01        I4 [ITEM]
p01        I4 | 3. Menteri Keuangan;
p01        I4 [ITEM]
p01        I4 | 4. Menteri Dalam Negeri;
p01        I4 [ITEM]
p01        I4 | 5. Para Gubernur; dan
p01        I4 [ITEM]
p01        I4 | 6. Para Bupati/Wali Kota.
p01   F14  I2 | Untuk
p01   F12  I2 | KESATU
p01        I4 | Mengambil langkah-langkah yang terkoordinasi dan
p01   F15  I4 | terintegrasi sesuai tugas, fungsi, dan kewenangan masing-
p01   F14  I4 | masing untuk:
p01        I4 [ITEM]
p01        I4 | 1. melaksanakan kegiatan pembangunan jalan daerah yang
p01   F14  I5 | terhubung dan terintegrasi, utamanya untuk mendukung
p01   F16  I5 | produktivitas kawasan industri, kawasan pariwisata,
p01   F14  I5 | kawasan perkebunan, kawasan pertanian, dan kawasan
p01   F14  I5 | produktif lainnya;
p01   F14  I1 | SK No 145524A
p01   F15  I10 [ITEM]
p01   F15  I10 | 2.melaksanakan...
==================== PAGE 2 ====================
p02   F12  I4 | 2
p02   F12  I6 | PRESIDEN
p02   F12  | REPUBLIK INDONESIA
p02   F16  | 2-
p02   F14  I5 | melaksanakan kegiatan pembangunan jalan daerah dan
p02   F14  I5 | membantu meningkatkan kemantapan jalan, utamanya:
p02        I5 [SUB-ITEM]
p02        I5 | a. di sekitar kawasan industri strategis, antara lain
p02   F15  | Morowali, Konawe, Weda Bay, dan Tanjung Selor
p02        | untuk mengantisipasi pertumbuhan kawasan
p02   F14  | kumuh; dan
p02        I5 [SUB-ITEM]
p02        I5 | b. kondisi jalan daerah yang belum mantap;
p02   F14  I5 | melaksanakan pembangunan jalan di sekitar kawasan Ibu
p02   F14  I5 | Kota Nusantara dengan melakukan pelebaran jalan untuk
p02   F13  I5 | mengantisipasi kemacetan;
p02        I5 | merencanakan dan
p02   F18  I7 | menyediakan anggaran,
p02   F17  I5 | melaksanakan, memantau, mengevaluasi serta
p02   F17  I5 | mengendalikan kegiatan percepatan peningkatan
p02   F14  I5 | konektivitas jalan daerah; dan
p02   F15  I5 | mengatasi kendala dan hambatan dalam pelaksanaan
p02   F15  I5 | kegiatan percepatan peningkatan konektivitas jalan
p02   F12  I5 | daerah.
p02   F13  I4 | Khusus kepada:
p02        I4 [ITEM]
p02        I4 | 1. Menteri Perencanaan Perribangunan Nasional/
p02   F16  I5 | Kepala Badan Perencanaan Pembangunan Nasional
p02   F14  I5 | untuk:
p02        I5 [SUB-ITEM]
p02        I5 | a. mengoordinasikan kegiatan percepatan peningkatan
p02   F14  | konektivitas jalan daerah;
p02        I5 [SUB-ITEM]
p02        I5 | b. merumuskan kriteria pemilihan ruas dan
p02   F15  | pemanfaatannya serta men5rusun indikasi lokasi,
p02   F18  | ruas, dan volume dalam kegiatan percepatan
p02   F18  | peningkatan konektivitas jalan daerah bersama
p02   F13  | Menteri Pekerjaan Umum dan Perumahan Rakyat;
p02        I5 [SUB-ITEM]
p02        I5 | c. melakukan verifikasi dan penilaian sebagai dasar
p02        | penentuan ruas dan jenis penanganan serta
p02   F15  | memastikan tidak ada tumpang tindih penanganan
p02   F15  | kegiatan jalan daerah yang dikerjakan daerah dan
p02   F16  | pusat bersama Menteri Pekerjaan Umum dan
p02   F13  | Perumahan Ralryat;
p02   F12  I4 | 3
p02   F12  I4 | 4
p02   F12  I4 | 5
p02   F12  I2 | KEDUA
p02   F14  I1 | SK No 145509 A
p02   F13  I10 [SUB-ITEM]
p02   F13  I10 | d. menetapkan
==================== PAGE 3 ====================
p03   F12  I5 | d
p03   F10  I5 | e
p03   F12  I6 | PRESIDEN
p03   F12  | REPUBLIK INDONESIA
p03   F12  | 3
p03   F14  [PREAMBLE:MENETAPKAN]
p03   F14  | menetapkan daftar kegiatan percepatan peningkatan
p03   F14  | konektivitas jalan daerah bersama Menteri Pekerjaan
p03   F13  | Umum dan Perumahan Ralryat;
p03   F14  | melakukan pemantauan, evaluasi, dan pengendalian
p03   F16  | pelaksanaan kegiatan percepatan peningkatan
p03   F14  | konektivitas jalan daerah bersama Menteri Pekerjaan
p03   F13  | Umum dan Perumahan Rakyat;
p03   F16  [PREAMBLE:MENETAPKAN]
p03   F16  | menetapkan pedoman pelaksanaan kegiatan
p03   F13  | percepatan peningkatan konektivitas j alan daerah;
p03   F17  | mengoordinasikan penyelesaian kendala dan
p03   F14  | hambatan dalam pelaksanaan kegiatan percepatan
p03   F14  | peningkatan konektivitas jalan daerah; dan
p03   F14  | melaporkan hasil pelaksanaan Instrrrksi Presiden ini
p03   F12  | kepada Presiden.
p03   F14  I5 | Menteri Pekerjaan Umum dan Perumahan Ralryat untuk:
p03        I5 [SUB-ITEM]
p03        I5 | a. merumuskan kriteria pemilihan ruas dan
p03   F15  | pemanfaatannya serta men)rusun indikasi lokasi,
p03   F18  | nras, dan volume dalam kegiatan percepatan
p03   F18  | peningkatan konektivitas jalan daerah bersama
p03   F16  | Menteri Perencanaan Pembangunan Nasional/
p03   F13  | Kepala Badan Perencanaan Pembangunan Nasional;
p03        I5 [SUB-ITEM]
p03        I5 | b. menentukan kriteria teknis sebagai dasar verifikasi
p03        | dan penilaian dalam kegiatan percepatan
p03   F14  | peningkatan konektivitas jalan daerah;
p03        I5 [SUB-ITEM]
p03        I5 | c. melakukan verifikasi dan penilaian sebagai dasar
p03        | penentuan ruas dan jenis penanganan serta
p03   F14  | memastikan tidak ada tumpang tindih penanganan
p03   F15  | kegiatan jalan daerah yang dikerjakan daerah dan
p03   F14  | pusat bersama Menteri Perencanaan Pembangunan
p03   F13  | Nasional/Kepala Badan Perencanaan Pembangunan
p03   F12  | Nasional;
p03        I5 [SUB-ITEM]
p03        I5 | d. menJrusun besaran pagu pada setiap ruas jalan yang
p03   F16  | direncanakan berdasarkan kriteria teknis, jenis
p03   F13  | penanganan, dan volume dalam kegiatan percepatan
p03   F14  | peningkatan konektivitas jalan daerah;
p03   F11  I5 [SUB-ITEM]
p03   F11  I5 | f.
p03   F10  I5 | ob
p03   F13  I5 | h
p03   F11  I4 | 2
p03   F14  I1 | SK No 145510 A
p03   F14  I10 [SUB-ITEM]
p03   F14  I10 | e. menetapkan. . .
==================== PAGE 4 ====================
p04   F12  I4 | 3
p04   F12  I6 | PRESIDEN
p04   F12  | REPUBLIK INDONESIA
p04   F14  I6 | -4 -
p04        I5 [SUB-ITEM]
p04        I5 | e. menetapkan daftar kegiatan percepatan peningkatan
p04        | konektivitas jalan daerah bersama Menteri
p04   F13  | Perencanaan Pembangunan Nasional/Kepala Badan
p04   F12  | Perencanaan Pembangunan Nasional;
p04        I5 [SUB-ITEM]
p04        I5 | f.
p04   F14  | memastikan rincian lokasi, mas, volume, dan pagu
p04   F16  | setiap ruas jalan dalam Daftar Isian Pelaksanaan
p04   F17  | Anggaran Kementerian Pekerjaan Umum dan
p04   F13  | Perumahan Rakyat;
p04        I5 [SUB-ITEM]
p04        I5 | g. melaksanakan kegiatan percepatan peningkatan
p04   F16  | konektivitas jalan daerah yang dapat melibatkan
p04   F14  | perangkat daerah terkait;
p04        I5 [SUB-ITEM]
p04        I5 | h. melakukan pemantauan, evaluasi, dan pengendalian
p04   F16  | pelaksanaan kegiatan percepatan peningkatan
p04        | konektivitas jalan daerah bersama Menteri
p04   F13  | Perencanaan Pembangunan Nasional I Kepala Badan
p04   F13  | Perencanaan Pembangunan Nasional; dan
p04        I5 [SUB-ITEM]
p04        I5 | i.
p04   F14  | melakukan serah terima hasil kegiatan percepatan
p04        | peningkatan konektivitas jalan daerah kepada
p04   F16  | pemerintah daerah dalam bentuk hibah sesuai
p04   F14  | dengan ketentuan peraturan perundang-undangan.
p04   F14  I5 | Menteri Keuangan untuk:
p04        I5 [SUB-ITEM]
p04        I5 | a. menyiapkan anggaran untuk pelaksanaan kegiatan
p04   F16  | percepatan peningkatan konektivitas jalan daerah
p04   F14  | pada tahun 2023 dan tahun 2024;
p04        I5 [SUB-ITEM]
p04        I5 | b. menyiapkan tambahan anggaran yang bersumber
p04   F17  | dari Anggaran Pendapatan dan Belanja Negara
p04   F15  | Tahun Anggaran 2023 untuk pelaksanaan kegiatan
p04   F16  | percepatan peningkatan konektivitas jalan daerah
p04   F16  | dengan menggunakan mekanisme kontrak tahun
p04   F14  | tunggal dan/atau kontrak tahun jamak; dan
p04        I5 [SUB-ITEM]
p04        I5 | c. memfasilitasi untuk melakukan percepatan proses
p04        | hibah hasil kegiatan percepatan peningkatan
p04   F16  | konektivitas jalan daerah dari Menteri Pekerjaan
p04   F14  | Umum dan Perumahan Ralryat kepada pemerintah
p04   F13  | daerah, bersama Menteri Dalam Negeri.
p04   F14  I1 | SK No 145526 A
p04   F14  I11 [ITEM]
p04   F14  I11 | 4. Menteri
==================== PAGE 5 ====================
p05   F12  I4 | 4
p05   F12  I6 | PRESiDEN
p05   F12  | REPUBLTK INDONESIA
p05   F14  I6 | -5
p05   F14  I5 | Menteri Dalam Negeri untuk:
p05        I5 [SUB-ITEM]
p05        I5 | a. memberikan sosialisasi kepada Gubernur dan
p05   F14  | Bupati/Wali Kota mengenai pelaksanaan kebijakan
p05   F14  | percepatan peningkatan konektivitas jalan daerah;
p05        I5 [SUB-ITEM]
p05        I5 | b. menyiapkan dukungan kebijakan yang dibutuhkan
p05   F16  | pemerintah daerah dalam kegiatan percepatan
p05   F14  | peningkatan konektivitas jalan daerah;
p05        I5 [SUB-ITEM]
p05        I5 | c. melaksanakan pembinaan dan pengawasan dalam
p05   F16  | pelaksanaan kegiatan percepatan peningkatan
p05   F14  | konektivitas jalan daerah yang menjadi kewenangan
p05   F13  | pemerintah daerah; dan
p05        I5 [SUB-ITEM]
p05        I5 | d. memfasilitasi untuk melakukan percepatan proses
p05        | hibah hasil kegiatan percepatan peningkatan
p05   F16  | konektivitas jalan daerah dari Menteri Pekerjaan
p05   F14  | Umum dan Perumahan Ralryat kepada pemerintah
p05   F13  | daerah, bersama Menteri Keuangan.
p05   F14  I5 | Gubernur dan Bupati/Wali Kota untuk:
p05        I5 [SUB-ITEM]
p05        I5 | a. menyediakan dukungan program dan anggaran
p05   F16  | dalam rangka menyiapkan dokumen kesiapan
p05   F16  | pelaksanaan kegiatan percepatan peningkatan
p05   F14  | konektivitas jalan daerah;
p05        I5 [SUB-ITEM]
p05        I5 | b. men5rusun dokumen perencanaan dan kelengkapan
p05   F15  | perizinan sesuai dengan kewenangannya untuk
p05   F14  | kegiatan percepatan peningkatan konektivitas jalan
p05   F12  | daerah;
p05        I5 [SUB-ITEM]
p05        I5 | c. menyediakan dukungan lahan siap bangun dalam
p05        | rangka pelaksanaan kegiatan percepatan
p05   F14  | peningkatan konektivitas jalan daerah; dan
p05        I5 [SUB-ITEM]
p05        I5 | d. mengoperasikan dan melakukan pemeliharaan jalan
p05   F14  | daerah yang telah diserahterimakan dalam bentuk
p05        | hibah hasil kegiatan percepatan peningkatan
p05   F16  | konektivitas jalan daerah dari Menteri Pekerjaan
p05   F13  | Umum dan Perumahan Rakyat.
p05   F12  I4 | 5
p05   F14  I1 | SK No 145520 A
p05   F12  I11 | KETIGA
==================== PAGE 6 ====================
p06   F12  I6 | PRESIDEN
p06   F12  | REPUBLTK INDONESIA
p06   F16  | 6-
p06        I4 | : Mendukung secara penuh tanggung jawab dan bersinergi
p06   F14  I4 | dalam melaksanakan Instruksi Presiden ini.
p06   F14  I4 | Pendanaan pelaksanaan Instruksi Presiden ini bersumber dari
p06   F16  I4 | Anggaran Pendapatan dan Belanja Negara dan Anggaran
p06   F13  I4 | Pendapatan dan Belanja Daerah.
p06   F14  I2 | Instruksi Presiden ini mulai berlaku pada tanggal dikeluarkan.
p06   F14  I6 | Dikeluarkan di Jakarta
p06   F14  I6 | pada tanggal 16 Maret 2023
p06   F12  I6 | PRESIDEN REPUBLIK INDONESIA,
p06   F12  I7 | JOKO WIDODO
p06   F13  I3 | Salinan sesuai dengan aslinya
p06   F12  I2 | KEMENTERIAN SEKRETARIAT NEGARA
p06   F12  I3 | REPUBLIK INDONESIA
p06   F13  I4 | Perundang-undangan dan
p06   F15  I5 | Hukunl
p06   F13  I5 | Djaman
p06   F12  I1 | KETIGA
p06   F12  I1 | KEEMPAT
p06   F14  I8 | ttd.
p06   F14  I1 | SK No 145525 A
```

---


## tap_mpr

**ERROR**: list index out of range

---


## uud-1945

- **File**: `uud-1945/uud_1945.pdf`
- **Document Type**: UUD 1945 (Constitution)
- **Issued by**: BPUPKI/PPKI
- **Pages**: 27 | **Lines**: 1143
- **Font sizes**: [12.0, 14.0]
- **Most common font**: 12.0 (99% of lines)
- **Bold font sizes**: [12.0, 14.0]
- **Indent clusters**: [90.0, 126.0, 162.0, 332.0, 374.0, 396.0, 420.0]
- **Expected hierarchy**: BAB > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01 B F14  | UNDANG-UNDANG DASAR
p01 B F14  | NEGARA REPUBLIK INDONESIA
p01 B F14  | TAHUN 1945
p01 B      I3 | (yang dipadukan dengan Perubahan I, II, III & IV)
p01 B      | PEMBUKAAN
p01 B      | (Preambule)
p01        I2 | Bahwa sesungguhnya kemerdekaan itu ialah hak segala bangsa dan oleh sebab itu,
p01        I1 | maka penjajahan di atas dunia harus dihapuskan, karena tidak sesuai dengan
p01        I1 | perikemanusiaan dan perikeadilan.
p01        I2 | Dan perjuangan pergerakan kemerdekaan Indonesia telah sampailah kepada saat
p01        I1 | yang berbahagia dengan selamat sentausa mengantarkan rakyat Indonesia ke depan pintu
p01        I1 | gerbang kemerdekaan Negara Indonesia yang merdeka, bersatu, berdaulat, adil dan
p01        I1 | makmur.
p01        I2 | Atas berkat rakhmat Allah Yang Maha Kuasa dan dengan didorongkan oleh
p01        I1 | keinginan luhur, supaya berkehidupan kebangsaan yang bebas, maka rakyat Indonesia
p01        I1 | menyatakan dengan ini kemerdekaannya.
p01        I2 | Kemudian dari pada itu untuk membentuk suatu Pemerintah Negara Indonesia
p01        I1 | yang melindungi segenap bangsa Indonesia dan seluruh tumpah darah Indonesia dan
p01        I1 | untuk memajukan kesejahteraan umum, mencerdaskan kehidupan bangsa dan ikut
p01        I1 | melaksanakan ketertiban dunia yang berdasarkan kemerdekaan, perdamaian abadi dan
p01        I1 | keadilan sosial, maka disusunlah Kemerdekaan Kebagsaan Indonesia itu dalam suatu
p01        I1 | Undang-Undang Dasar Negara Indonesia, yang terbentuk dalam suatu susunan Negara
p01        I1 | Republik Indonesia, yang berkedaulatan rakyat dengan berdasar kepada : Ketuhanan
p01        I1 | Yang Maha Esa, Kemanusiaan yang adil dan beradab, Persatuan Indonesia dan
p01        I1 | Kerakyatan
p01        I3 | yang
p01        | dipimpin
p01        | oleh
p01        I4 | hikmat
p01        I5 | kebijaksanaan
p01        | dalam
p01        I1 | permusyawaratan/perwakilan, serta dengan mewujudkan suatu Keadilan sosial bagi
p01        I1 | seluruh rakyat Indonesia.
==================== PAGE 2 ====================
p02 B      [HEADING:BAB]
p02 B      | BAB I
p02 B      | BENTUK DAN KEDAULATAN
p02 B      | Pasal 1
p02        I1 [AYAT]
p02        I1 | (1)
p02        I2 | Negara Indonesia ialah Negara kesatuan yang berbentuk Republik.
p02        I1 | * (2)
p02        I2 | Kedaulatan berada di tangan rakyat dan dilaksanakan menurut Undang-Undang
p02        I2 | Dasar.
p02        I1 | * (3)
p02        I2 | Negara Indonesia adalah negara hukum.
p02        I2 | * Perubahan III  9 November 2001,  sebelumnya berbunyi :
p02        I3 [AYAT]
p02        I3 | (1)
p02        | Negara Indonesia ialah Negara kesatuan yang berbentuk
p02        | Republik.
p02        I3 [AYAT]
p02        I3 | (2)
p02        | Kedaulatan adalah di tangan rakyat dan dilakukan sepenuhnya
p02        | oleh Majelis Permusyawaratan Rakyat.
p02 B      [HEADING:BAB]
p02 B      | BAB II
p02 B      I3 | MAJELIS PERMUSYAWARATAN RAKYAT
p02 B      | Pasal 2
p02        I1 | * (1)
p02        I2 | Majelis Permusyawaratan Rakyat terdiri atas anggota Dewan Perwakilan Rakyat
p02        I2 | dan anggota Dewan Perwakilan Daerah yang dipilih melalui pemilihan umum dan
p02        I2 | diatur lebih lanjut dengan undang-undang.
p02        I1 [AYAT]
p02        I1 | (2)
p02        I2 | Majelis Permusyawaratan Rakyat bersidang sedikitnya sekali dalam lima tahun di
p02        I2 | ibu kota Negara.
p02        I1 [AYAT]
p02        I1 | (3)
p02        I2 | Segala putusan Majelis Permusyawaratan Rakyat ditetapkan dengan suara yang
p02        I2 | terbanyak
p02        I2 | * Perubahan IV  10 Agustus 2002,  sebelumnya berbunyi :
p02        I3 [AYAT]
p02        I3 | (1)
p02        | Majelis permusyawaratan Rakyat terdiri atas anggota-anggota
p02        | Dewan Perwakilan Rakyat, ditambah dengan utusan-utusan dari
p02        | daerah-daerah dan golongan-golongan, menurut aturan yang
p02        | ditetapkan dengan Undang-Undang.
p02 B      | *Pasal 3
p02        I1 [AYAT]
p02        I1 | (1)
p02        I2 | Majelis Permusyawaratan Rakyat berwenang mengubah dan menetapkan
p02        I2 | Undang-Undang Dasar.
p02        I1 [AYAT]
p02        I1 | (2)
p02        I2 | Majelis Permusyawaratan Rakyat melantik Presiden dan/atau Wakil Presiden.
p02        I1 [AYAT]
p02        I1 | (3)
p02        I2 | Majelis Permusyawaratan Rakyat hanya dapat memberhentikan Presiden dan/atau
p02        I2 | Wakil Presiden dalam masa jabatannya menurut Undang-Undang Dasar.
p02        I2 | * Perubahan III  9 November 2001,  sebelumnya berbunyi :
==================== PAGE 3 ====================
p03        I3 [AYAT]
p03        I3 | (1)
p03        | Majelis Permusyawaratan Rakyat menetapkan Undang-Undang
p03        | Dasar dan Garis-garis besar dari pada haluan negara.
p03 B      [HEADING:BAB]
p03 B      | BAB III
p03 B      | KEKUASAAN PEMERINTAHAN NEGARA
p03 B      | Pasal 4
p03        I1 [AYAT]
p03        I1 | (1)
p03        I2 | Presiden Republik Indonesia memegang kekuasaan Pemerintahan menurut
p03        I2 | Undang-Undang Dasar.
p03        I1 [AYAT]
p03        I1 | (2)
p03        I2 | Dalam melakukan kewajibannya Presiden dibantu oleh satu orang Wakil
p03        I2 | Presiden.
p03 B      | Pasal 5
p03        I1 | * (1)
p03        I2 | Presiden berhak mengajukan rancangan Undang-undang kepada Dewan
p03        I2 | Perwakilan Rakyat.
p03        I1 [AYAT]
p03        I1 | (2)
p03        I2 | Presiden menetapkan Peraturan Pemerintah untuk menjalankan Undang-undang
p03        I2 | sebagaimana mestinya.
p03        I2 | * Perubahan I  19 Oktober 1999,  sebelumnya berbunyi :
p03        I3 [AYAT]
p03        I3 | (1)
p03        | Presiden memegang kekuasaan membentuk Undang-undang
p03        | dengan persetujuan Dewan Perwakilan Rakyat.
p03 B      | Pasal 6
p03        I1 | * (1)
p03        I2 | Calon Presiden dan calon Wakil Presiden harus warga negara Indonesia sejak
p03        I2 | kelahirannya dan tidak pernah menerima kewarganegaraan lain karena
p03        I2 | kehendaknya sendiri, tidak pernah mengkhianati negara, serta mampu secara
p03        I2 | rohani dan jasmani untuk melaksanakan tugas dan kewajiban sebagi Presiden dan
p03        I2 | Wakil Presiden.
p03        I1 | * (2)
p03        I2 | Syarat-syarat untuk menjadi Presiden dan Wakil Presiden diatur lebih lanjut
p03        I2 | dengan undang-undang.
p03        I2 | * Perubahan III  9 November 2001,  sebelumnya berbunyi :
p03        I3 [AYAT]
p03        I3 | (1)
p03        | Presiden ialah orang Indonesia asli.
p03        I3 [AYAT]
p03        I3 | (2)
p03        | Presiden
p03        | dan
p03        | Wakil
p03        I4 | Presiden
p03        I6 | dipilih
p03        | oleh
p03        | Majelis
p03        | Permusyawaratan Rakyat dengan suara yang terbanyak.
p03 B      | *Pasal 6A
p03        I1 [AYAT]
p03        I1 | (1)
p03        I2 | Presiden dan Wakil Presiden dipilih dalam satu pasangan secara langsung oleh
p03        I2 | rakyat.
p03        I1 [AYAT]
p03        I1 | (2)
p03        I2 | Pasangan calon Presiden dan Wakil Presiden diusulkan oleh partai politik atau
p03        I2 | gabungan partai politik peserta pemilihan umum sebelum pelaksanaan pemilihan
p03        I2 | umum.
==================== PAGE 4 ====================
p04        I1 [AYAT]
p04        I1 | (3)
p04        I2 | Pasangan calon Presiden dan Wakil Presiden yang mendapatkan suara lebih dari
p04        I2 | lima puluh persen dari jumlah suara dalam pemilihan umum dengan sedikitnya
p04        I2 | dua puluh persen suara di setiap provinsi yang tersebar di lebih dari setengah
p04        I2 | jumlah provinsi di Indonesia, dilantik menjadi Presiden dan Wakil Presiden.
p04        I2 | * Perubahan III  9 November 2001
p04        I1 | * (4)
p04        I2 | Dalam hal tidak ada pasangan calon Presiden dan Wakil Presiden terpilih, dua
p04        I2 | pasangan calon yang memperoleh suara terbanyak pertama dan kedua dalam
p04        I2 | pemilihan umum dipilih oleh rakyat secara langsung dan pasangan yang
p04        I2 | memperoleh suara rakyat terbanyak dilantik sebagai Presiden dan Wakil Presiden.
p04        I2 | * Perubahan IV  10 Agustus 2002
p04        I1 | * (5)
p04        I2 | Tata cara pelaksanaan pemilihan Presiden dan Wakil Presiden lebih lanjut diatur
p04        I2 | dalam undang-undang.
p04        I2 | * Perubahan III  9 November 2001
p04 B      | *Pasal 7
p04        I2 | Presiden dan Wakil Presiden memegang jabatan selama lima tahun dan
p04        I1 | sesudahnya dapat dipilih kembali dalam jabatan yang sama, hanya untuk satu kali masa
p04        I1 | jabatan.
p04        I1 | * Perubahan I 19 Oktober 1999,  sebelumnya berbunyi :
p04        I2 | Presiden dan Wakil Presiden memegang jabatannya selama masa lima tahun,
p04        I2 | dan sesudahnya dapat dipilih kembali.
p04 B      | *Pasal 7A
p04        I2 | Presiden dan/atau Wakil Presiden dapat diberhentikan dalam masa jabatannya
p04        I1 | oleh Majelis Permusyawaratan Rakyat atas usul Dewan Perwakilan Rakyat, baik apabila
p04        I1 | terbukti telah melakukan pelanggaran hukum berupa pengkhianatan terhadap negara,
p04        I1 | korupsi, penyuapan, tindak pidana berat lainnya, atau perbuatan tercela maupun apabila
p04        I1 | terbukti tidak lagi memenuhi syarat sebagai Presiden dan/atau Wakil Presiden.
p04        I1 | * Perubahan III  9 November 2001
p04 B      | Pasal 7B
p04        I1 | * (1)
p04        I2 | Usul pemberhentian Presiden dan/atau Wakil Presiden dapat diajukan oleh Dewan
p04        I2 | Perwakilan Rakyat kepada Majelis Permusyawaratan Rakyat hanya dengan
p04        I2 | terlebih dahulu mengajukan permintaan kepada Mahkamah Konstitusi untuk
p04        I2 | memeriksa, mengadili, dan memutus pendapat Dewan Perwakilan Rakyat bahwa
p04        I2 | Presiden dan/atau Wakil Presiden telah melakukan pelanggaran hukum berupa
==================== PAGE 5 ====================
p05        I2 | pengkhianatan terhadap negara, korupsi, penyuapan, tindak pidana berat lainnya,
p05        I2 | atau perbuatan tercela; dan/atau pendapat bahwa Presiden dan/atau Wakil
p05        I2 | Presiden tidak lagi memenuhi syarat sebagai Presiden dan/atau Wakil Presiden.
p05        I1 | * (2)
p05        I2 | Pendapat Dewan Perwakilan Rakyat bahwa Presiden dan/atau Wakil Presiden
p05        I2 | telah melakukan pelanggaran hukum tersebut ataupun telah tidak lagi memenuhi
p05        I2 | syarat sebagai Presiden dan/atau Wakil Presiden adalah dalam rangka
p05        I2 | pelaksanaan fungsi pengawasan Dewan Perwakilan Rakyat.
p05        I1 | * (3)
p05        I2 | Pengajuan permintaan Dewan Perwakilan Rakyat kepada Mahkamah Konstitusi
p05        I2 | hanya dapat dilakukan dengan dukungan sekurang-kurangnya 2/3 dari jumlah
p05        I2 | anggota Dewan Perwakilan Rakyat yang hadir dalam sidang paripurna yang
p05        I2 | dihadiri oleh sekurang-kurangnya 2/3 dari jumlah anggota Dewan Perwakilan
p05        I2 | Rakyat.
p05        I1 | * (4)
p05        I2 | Mahkamah Konstitusi wajib memeriksa, mengadili, dan memutus dengan seadil-
p05        I2 | adilnya terhadap pendapat Dewan Perwakilan Rakyat tersebut paling lama
p05        I2 | sembilan puluh hari setelah permintaan Dewan Perwakilan Rakyat itu diterima
p05        I2 | oleh Mahkamah Konstitusi.
p05        I1 | * (5)
p05        I2 | Apabila Mahkamah Konstitusi memutuskan bahwa Presiden dan/atau Wakil
p05        I2 | Presiden terbukti melakukan pelanggaran hukum berupa pengkhianatan terhadap
p05        I2 | negara, korupsi, penyuapan, tindak pidana berat lainnya, atau perbuatan tercela;
p05        I2 | dan/atau terbukti bahwa Presiden dan/atau Wakil Presiden tidak lagi memenuhi
p05        I2 | syarat sebagai Presiden dan/atau Wakil Presiden, Dewan Perwakilan Rakyat
p05        I2 | menyelenggarakan sidang paripurna untuk meneruskan usul pemberhentian
p05        I2 | Presiden dan/atau Wakil Presiden kepada Majelis Permusyawaratan Rakyat.
p05        I1 | * (6)
p05        I2 | Majelis Permusyawaratan Rakyat wajib menyelenggarakan sidang untuk
p05        I2 [KEPUTUSAN:MEMUTUSKAN]
p05        I2 | memutuskan usul Dewan Perwakilan Rakyat tersebut paling lambat tiga puluh
p05        I2 | hari sejak Majelis Permusyawaratan Rakyat menerima usul tersebut.
p05        I2 | * Perubahan III  9 November 2001
p05        I1 | * (7)
p05        I2 | Keputusan Majelis Permusyawaratan Rakyat atas usul pemberhentian Presiden
p05        I2 | dan/atau Wakil Presiden harus diambil dalam rapat paripurna Majelis
p05        I2 | Permusyawaratan Rakyat yang dihadiri oleh sekurang-kurangnya 3/4 dari jumlah
p05        I2 | anggota dan disetujui oleh sekurang-kurangnya 2/3 dari jumlah anggota yang
p05        I2 | hadir, setelah Presiden dan/atau Wakil Presiden diberi kesempatan menyampaikan
p05        I2 | penjelasan dalam rapat paripurna Majelis Permusyawaratan Rakyat.
p05        I2 | * Perubahan III November 2001
p05 B      | *Pasal 7C
p05        I2 | Presiden tidak dapat membekukan dan/atau membubarkan Dewan Perwakilan
p05        I1 | Rakyat.
p05        I1 | * Perubahan III  November 2001
==================== PAGE 6 ====================
p06 B      | Pasal 8
p06        I1 | * (1)
p06        I2 | Jika Presiden mangkat, berhenti, diberhentikan, atau tidak dapat melakukan
p06        I2 | kewajibannya dalam masa jabatannya, ia digantikan oleh Wakil Presiden sampai
p06        I2 | habis masa jabatannya.
p06        I1 | * (2)
p06        I2 | Dalam hal terjadi kekosongan Wakil Presiden, selambat-lambatnya dalam waktu
p06        I2 | enam puluh hari, Majelis Permusyawaratan Rakyat menyelenggarakan sidang
p06        I2 | untuk memilih Wakil Presiden dari dua calon yang diusulkan oleh Presiden.
p06        I2 | * Perubahan III  November 2001
p06        I1 | * (3)
p06        I2 | Jika Presiden dan Wakil Presiden mangkat, berhenti, diberhentikan, atau tidak
p06        I2 | dapat melakukan kewajibannya dalam masa jabatannya secara bersamaan,
p06        I2 | pelaksana tugas kepresidenan adalah Menteri Luar Negeri, Menteri Dalam Negeri
p06        I2 | dan Menteri Pertahanan secara bersama-sama. Selambat-lambatnya tiga puluh
p06        I2 | hari setelah itu, Majelis Permusyawaratan Rakyat menyelenggarakan siding untuk
p06        I2 | memilih Presiden dan Wakil Presiden dari dua pasangan calon Presiden dan
p06        I2 | Wakil Presiden yang diusulkan oleh partai politik atau gabungan partai politik
p06        I2 | yang yang pasangan calon Presiden dan Wakil Presidennya meraih suara
p06        I2 | terbanyak pertama dan kedua dalam pemilihan umum sebelumnya, samapi
p06        I2 | berakhir masa jabatannya.
p06        I2 | * Perubahan IV  10 Agustus 2002, sebelumnya berbunyi :
p06        I3 | Jika Presiden mangkat, berhenti, atau tidak dapat melakukan
p06        I3 | kewajibannya dalam masa jabatannya, ia diganti oleh Wakil Presiden
p06        I3 | sampai habis batas waktunya.
p06 B      | *Pasal 9
p06        I1 [AYAT]
p06        I1 | (1)
p06        I2 | Sebelum memangku jabatannya, Presiden dan Wakil Presiden bersumpah
p06        I2 | menurut agama, atau berjanji dengan sungguh-sungguh di hadapan Majelis
p06        I2 | Permusyawaratan Rakyat atau Dewan Perwakilan Rakyat sebagai berikut :
p06        I2 | Sumpah Presiden (Wakil Presiden)
p06        I2 | “Demi Allah, saya bersumpah akan memenuhi kewajiban Presiden Republik
p06        I2 | Indonesia (Wakil Presiden Republik Indonesia) dengan sebaik-baiknya dan
p06        I2 | seadil-adilnya, memegang teguh Undang-Undang Dasar dan menjalankan segala
p06        I2 | Undang-Undang dan peraturannya dengan selurus-lurusnya serta berbakti kepada
p06        I2 | Nusa dan Bangsa”.
p06        I2 | Janji Presiden (Wakil Presiden) :
p06        I2 | “Saya berjanji dengan sungguh-sungguh akan memenuhi kewajiban Presiden
p06        I2 | Republik Indonesia (Wakil Presiden Republik Indonesia) dengan sebaik-baiknya
p06        I2 | dan seadil-adilnya, memegang teguh Undang-Undang Dasar dan menjalankan
==================== PAGE 7 ====================
p07        I2 | segala Undang-Undang dan peraturannya dengan selurus-lurusnya serta berbakti
p07        I2 | kepada Nusa dan bangsa”.
p07        I1 [AYAT]
p07        I1 | (2)
p07        I2 | Jika Majelis Permusyawaratan Rakyat atau Dewan Perwakilan Rakyat tidak dapat
p07        I2 | mengadakan sidang Presiden dan Wakil Presiden bersumpah menurut agama, atau
p07        I2 | berjanji dengan sungguh-sungguh dihadapan pimpinan Majelis Permusyawaratan
p07        I2 | Rakyat dengan disaksikan oleh pimpinan Mahkamah Agung.
p07        I2 | * Perubahan I 19 Oktober 1999, sebelumnya berbunyi :
p07        I3 | Sebelum memangku jabatannya, Presiden dan Wakil Presiden bersumpah
p07        I3 | menurut agama, atau berjanji dengan sungguh-sungguh dihadapan
p07        I3 | Majelis Permusyawaratan Rakyat atau Dewan Perwakilan Rakyat sebagai
p07        I3 | berikut :
p07        I3 | Sumpah Presiden (Wakil Presiden) :
p07        I3 | “Demi Allah, saya bersumpah akan memenuhi kewajiban Presiden
p07        I3 | Republik Indonesia (Wakil Presiden Republik Indonesia) dengan sebaik-
p07        I3 | baiknya dan seadil-adilnya, memegang teguh Undang-Undang Dasar dan
p07        I3 | menjalankan segala Undang-undang dan Peraturannya dengan selurus-
p07        I3 | lurusnya serta berbakti kepada Nusa dan Bangsa”.
p07        I3 | Janji Presiden (Wakil Presiden) :
p07        I3 | Saya berjanji dengan sungguh-sungguh akan memenuhi kewajiban
p07        I3 | Presiden Republik Indonesia (Wakil Presiden Republik Indonesia) dengan
p07        I3 | sebaik-baiknya dan seadil-adilnya, memegang teguh Undang-Undang
p07        I3 | Dasar dan menjalankan segala Undang-undang dan Peraturannya
p07        I3 | dengan selurus-lurusnya serta berbakti kepada Nusa dan Bangsa.
p07 B      | Pasal 10
p07        I2 | Presiden memegang kekuasaan yang tertinggi atas Angkatan Darat, Angkatan
p07        I1 | Laut, dan Angkatan Udara.
p07 B      | Pasal 11
p07        I1 | * (1)
p07        I2 | Presiden dengan persetujuan Dewan Perwakilan Rakyat menyatakan perang,
p07        I2 | membuat perdamaian dan perjanjian dengan negara lain.
p07        I2 | * Perubahan IV 10 Agustus 2002
p07        I1 | * (2)
p07        I2 | Presiden dalam membuat perjanjian internasional lainnya yang menimbulkan
p07        I2 | akibat yang luas dan mendasar bagi kehidupan rakyat yang terkait dengan beban
p07        I2 | keuangan negara, dan/atau mengharuskan perubahan atau pembentukan undang-
p07        I2 | undang harus dengan persetujuan Dewan Perwakilan Rakyat.
p07        I1 | * (3)
p07        I2 | Ketentuan lebih lanjut tentang perjanjian internasional diatur dengan undang-
p07        I2 | undang.
==================== PAGE 8 ====================
p08        I2 | * Perubahan III  November 2001, sebelumnya berbunyi :
p08        I3 | Presiden dengan persetujuan Dewan Perwakilan Rakyat menyatakan
p08        I3 | keadaan bahaya ditetapkan dengan undang-undang.
p08 B      | Pasal 12
p08        I1 | Presiden menyatakan keadaan bahaya. Syarat-syarat dan akibatnya keadaan bahaya
p08        I1 | ditetapkan dengan Undang-undang.
p08 B      | *Pasal 13
p08        I1 [AYAT]
p08        I1 | (1)
p08        I2 | Presiden mengangkat Duta dan Konsul
p08        I1 [AYAT]
p08        I1 | (2)
p08        I2 | Dalam hal mengangkat duta, Presiden memperhatikan pertimbangan Dewan
p08        I2 | Perwakilan Rakyat.
p08        I1 [AYAT]
p08        I1 | (3)
p08        I2 | Presiden menerima penempatan duta negara lain dengan memperhatikan
p08        I2 | pertimbangan Dewan Perwakilan Rakyat.
p08        I2 | * Perubahan I 19 Oktober 1999, sebelumnya berbunyi :
p08        I3 | Pasal 13
p08        [AYAT]
p08        | (1)
p08        | Presiden mengangkat Duta dan Konsul.
p08        [AYAT]
p08        | (2)
p08        | Presiden menerima Duta negara lain.
p08 B      | *Pasal 14
p08        I1 [AYAT]
p08        I1 | (1)
p08        I2 | Presiden memberi grasi dan rehabilitasi dengan memperhatikan pertimbangan
p08        I2 | Mahkamah Agung.
p08        I1 [AYAT]
p08        I1 | (2)
p08        I2 | Presiden memberi amnesti dan abolisi dengan memperhatikan pertimbangan
p08        I2 | Dewan Perwakilan Rakyat.
p08        I2 | * Perubahan I 19 Oktober 1999, sebelumnya berbunyi :
p08        I3 | Pasal 14
p08        | Presiden memberi grasi, amnesti, abolisi dan rehabilitasi.
p08 B      | * Pasal 15
p08        I2 | Presiden memberi gelar, tanda jasa, dan lain-lain tanda kehormatan yang diatur
p08        I1 | dengan Undang-undang.
p08        I1 | * Perubahan I 19 Oktober 1999, sebelumnya berbunyi :
p08        I2 | Pasal 15
p08        I3 | Presiden memberi gelar, tanda jasa, dan lain-lain tanda kehormatan.
==================== PAGE 9 ====================
p09 B      | * Pasal 16
p09        I2 | Presiden membentuk suatu dewan pertimbangan yang bertugas memberikan
p09        I1 | nasihat dan pertimbangan kepada Presiden, yang selanjutnya diatur dalam undang-
p09        I1 | undang.
p09        I1 | * Perubahan IV 10 Agustus 2002
p09 B      | * BAB IV
p09 B      | DEWAN PERTIMBANGAN AGUNG
p09 B      | Dihapus.
p09        I1 | * Perubahan IV 10 Agustus 2002, sebelumnya berbunyi :
p09        I2 [HEADING:BAB]
p09        I2 | BAB IV
p09        I2 | DEWAN PERTIMBANGAN AGUNG
p09        I2 | Pasal 16
p09        I3 [AYAT]
p09        I3 | (1)
p09        | Susunan Dewan Pertimbangan Agung ditetapkan dengan Undang-
p09        | undang.
p09        I3 [AYAT]
p09        I3 | (2)
p09        | Dewan ini berkewajiban memberi jawab atas pertanyaan Presiden
p09        | dan berhak memajukan usul kepada pemerintah.
p09 B      [HEADING:BAB]
p09 B      | BAB V
p09 B      | KEMENTERIAN NEGARA
p09 B      | Pasal 17
p09        I1 [AYAT]
p09        I1 | (1)
p09        I2 | Presiden dibantu oleh Menteri-menteri negara.
p09        I1 | * (2)
p09        I2 | Menteri-menteri itu diangkat dan diberhentikan oleh Presiden.
p09        I1 | * (3)
p09        I2 | Setiap menteri membidangi urusan tertentu dalam pemerintahan.
p09        I2 | * Perubahan I 19 Oktober 1999, sebelumnya berbunyi :
p09        I3 [AYAT]
p09        I3 | (2)
p09        | Menteri-menteri itu diangkat dan diberhentikan oleh Presiden.
p09        I3 [AYAT]
p09        I3 | (3)
p09        | Menteri-menteri itu memimpin Departemen Pemerintahan.
p09        I1 | * (4)
p09        I2 | Pembentukan, pengubahan, dan pembubaran kementerian negara diatur dalam
p09        I2 | undang-undang.
p09        I2 | * Perubahan III 9 November 2001
p09 B      [HEADING:BAB]
p09 B      | BAB VI
p09 B      | PEMERINTAHAN DAERAH
p09 B      | Pasal 18
p09        I1 | * (1)
p09        I2 | Negara Kesatuan Republik Indonesia dibagi atas daerah-daerah propinsi dan
p09        I2 | daerah propinsi itu dibagi atas kabupaten dan kota, yang tiap-tiap propinsi,
==================== PAGE 10 ====================
p10        I2 | kabupaten, dan kota itu mempunyai pemerintahan daerah, yang diatur dengan
p10        I2 | undang-undang.
p10        I2 | * Perubahan II 18 Agustus 2000, sebelumnya berbunyi :
p10        I3 | Pembagian daerah Indonesia atas daerah besar dan kecil dengan bentuk
p10        I3 | susunan pemerintahannya ditetapkan dengan Undang-undang, dengan
p10        I3 | memandang dan mengingati dasar permusyawaratan dalam sistim
p10        I3 | Pemerintahan Negara, dan hak-hak asal-usul dalam daerah yang bersifat
p10        I3 | istimewa.
p10        I1 | * (2)
p10        I2 | Pemerintahan daerah propinsi, daerah kabupaten, dan kota mengatur dan
p10        I2 | mengurus sendiri urusan pemerintahan menurut asas otonomi dan tugas
p10        I2 | pembantuan.
p10        I1 | * (3)
p10        I2 | Pemerintahan daerah propinsi, daerah kabupaten, dan kota memiliki Dewan
p10        I2 | Perwakilan Rakyat Daerah yang anggota-anggotanya dipilih melalui pemilihan
p10        I2 | umum.
p10        I1 | * (4)
p10        I2 | Gubernur, Bupati, dan Walikota masing-masing sebagai kepala pemerintahan
p10        I2 | daerah propinsi, kabupaten, dan kota dipilih secara demokratis.
p10        I1 | * (5)
p10        I2 | Pemerintahan daerah menjalankan otonomi seluas-luasnya, kecuali urusan
p10        I2 | pemerintahan yang oleh undang-undang ditentukan sebagai urusan Pemerintah.
p10        I1 | * (6)
p10        I2 | Pemerintahan daerah berhak menetapkan peraturan daerah dan peraturan-
p10        I2 | peraturan lain untuk melaksanakan otonomi dan tugas pembantuan.
p10        I1 | * (7)
p10        I2 | Susunan dan tata cara penyelenggaraan pemerintahan daerah diatur dalam
p10        I2 | undang-undang.
p10        I2 | * Perubahan II 18 Agustus 2000.
p10 B      | Pasal 18A
p10        I1 | * (1)
p10        I2 | Hubungan wewenang antara pemerintahan pusat dan pemerintahan daerah
p10        I2 | propinsi, kabupaten, kota, atau antara propinsi dan kabupaten dan kota, diatur
p10        I2 | dengan undang-undang dengan memperhatikan kekhususan dan keragaman
p10        I2 | daerah.
p10        I1 | * (2)
p10        I2 | Hubungan keuangan, pelayanan umum, pemanfaatan sumber daya alam dan
p10        I2 | sumber daya lainnya antara pemerintahan pusat dan pemerintahan daerah diatur
p10        I2 | dan dilaksanakan secara adil dan selaras berdasarkan undang-undang.
p10        I2 | * Perubahan II 18 Agustus 2000.
p10 B      | Pasal 18B
p10        I1 | * (1)
p10        I2 | Negara mengakui dan menghormati satuan-satuan pemerintahan daerah yang
p10        I2 | bersifat khusus atau bersifat istimewa yang diatur dengan undang-undang.
p10        I1 | * (2)
p10        I2 | Negara mengakui dan menghormati kesatuan-kesatuan masyarakat hukum adat
p10        I2 | beserta hak-hak tradisionalnya sepanjang masih hidup dan sesuai dengan
==================== PAGE 11 ====================
p11        I2 | perkembangan masyarakat dan prinsip Negara Kesatuan Republik Indonesia,
p11        I2 | yang diatur dalam undang-undang.
p11        I2 | * Perubahan II 18 Agustus 2000.
p11 B      [HEADING:BAB]
p11 B      | BAB VII
p11 B      | DEWAN PERWAKILAN RAKYAT
p11 B      | Pasal 19
p11        I1 | * (1)
p11        I2 | Anggota Dewan Perwakilan Rakyat dipilih melalui pemilihan umum.
p11        I1 | * (2)
p11        I2 | Susunan Dewan Perwakilan Rakyat diatur dengan undang-undang.
p11        I1 | * (3)
p11        I2 | Dewan Perwakilan Rakyat bersidang sedikitnya sekali dalam setahun.
p11        I2 | * Perubahan II 18 Agustus 2000, sebelumnya berbunyi :
p11        I3 [AYAT]
p11        I3 | (1)
p11        | Susunan Dewan Perwakilan Rakyat ditetapkan dengan Undang-
p11        | undang.
p11        I3 [AYAT]
p11        I3 | (2)
p11        | Dewan Perwakilan Rakyat bersidang sedikitnya sekali dalam
p11        | setahun.
p11 B      | Pasal 20
p11        I1 | * (1)
p11        I2 | Dewan Perwakilan Rakyat memegang kekuasaan membentuk Undang-undang.
p11        I1 | * (2)
p11        I2 | Setiap rancangan Undang-undang dibahas oleh Dewan Perwakilan Rakyat dan
p11        I2 | Presiden untuk mendapat persetujuan bersama.
p11        I1 | * (3)
p11        I2 | Jika rancangan Undang-undang itu tidak mendapat persetujuan bersama,
p11        I2 | rancangan Undang-undang itu tidak boleh diajukan lagi dalam persidangan
p11        I2 | Dewan Perwakilan Rakyat masa itu.
p11        I1 | * (4)
p11        I2 | Persidangan mengesahkan rancangan Undang-undang yang telah disetujui
p11        I2 | bersama untuk menjadi Undang-undang.
p11        I1 | * (5)
p11        I2 | Dalam rancangan undang-undang yang telah disetujui bersama tersebut tidak
p11        I2 | disahkan oleh Presiden dalam waktu tiga puluh hari semenjak rancangan undang-
p11        I2 | undang tersebut disetujui, rancangan undang-undang tersebut sah menjadi
p11        I2 | Undang-undang dan wajib diundangkan.
p11        I2 | * Perubahan I 19 Oktober 1999, sebelumnya berbunyi :
p11        I3 [AYAT]
p11        I3 | (1)
p11        | Tiap-tiap undang-undang menghendaki persetujuan Dewan
p11        | Perwakilan Rakyat.
p11        I3 [AYAT]
p11        I3 | (2)
p11        | Jika
p11        | sesuatu
p11        | rancangan
p11        I4 | Undang-undang
p11        I7 | tidak
p11        | mendapat
p11        | persetujuan Dewan Perwakilan Rakyat, maka rancangan tadi tidak
p11        | boleh dimajukan lagi dalam persidangan Dewan Perwakilan
p11        | Rakyat masa itu.
==================== PAGE 12 ====================
p12 B      | Pasal 20A
p12        I1 | * (1)
p12        I2 | Dewan Perwakilan Rakyat memiliki fungsi legislasi, fungsi anggaran, dan fungsi
p12        I2 | pengawasan.
p12        I1 | * (2)
p12        I2 | Dalam melaksanakan fungsinya, selain hak yang diatur dalam pasal-pasal lain
p12        I2 | Undang-Undang Dasar ini, Dewan Perwakilan Rakyat mempunyai hak
p12        I2 | interpelasi, hak angket, dan hak menyatakan pendapat.
p12        I1 | * (3)
p12        I2 | Selain hak yang diatur dalam pasal-pasal lain Undang-Undang Dasar ini, setiap
p12        I2 | anggota Dewan Perwakilan Rakyat mempunyai hak mengajukan pertanyaan,
p12        I2 | menyampaikan usul dan pendapat, serta hak imunitas.
p12        I1 | * (4)
p12        I2 | Ketentuan lebih lanjut tentang hak Dewan Perwakilan Rakyat dan hak anggota
p12        I2 | Dewan Perwakilan Rakyat diatur dalam undang-undang.
p12        I2 | * Perubahan II 18 Agustus 2000.
p12 B      | * Pasal 21
p12        I1 | Anggota Dewan Perwakilan Rakyat berhak mengajukan usul rancangan Undang-undang.
p12        I1 | * Perubahan I 19 Oktober 1999, sebelumnya berbunyi :
p12        I2 [AYAT]
p12        I2 | (1)
p12        I3 | Anggota-anggota
p12        | Dewan
p12        | Perwakilan
p12        I5 | Rakyat
p12        I6 | berhak
p12        | memajukan
p12        I3 | rancangan Undang-undang.
p12        I2 [AYAT]
p12        I2 | (2)
p12        I3 | Jika rancangan itu, meskipun disetujui oleh Dewan Perwakilan Rakyat,
p12        I3 | tidak disahkan oleh Presiden, maka rancangan tadi tidak boleh dimajukan
p12        I3 | lagi dalam persidangan Dewan Perwakilan Rakyat masa itu.
p12 B      | Pasal 22
p12        I1 [AYAT]
p12        I1 | (1)
p12        I2 | Dalam hal ikhwal kegentingan yang memaksa, Presiden berhak menetapkan
p12        I2 | peraturan pemerintah sebagai pengganti undang-undang.
p12        I1 [AYAT]
p12        I1 | (2)
p12        I2 | Peraturan Pemerintah itu harus mendapat persetujuan Dewan Perwakilan Rakyat
p12        I2 | dalam persidangan yang berikut.
p12        I1 [AYAT]
p12        I1 | (3)
p12        I2 | Jika tidak mendapat persetujuan, maka Peraturan Pemerintah itu harus dicabut.
p12 B      | * Pasal 22A
p12        I1 | Ketentuan lebih lanjut tentang tata cara pembentukan undang-undang diatur dengan
p12        I1 | undang-undang.
p12        I1 | * Perubahan II 18 Agustus 2000.
p12 B      | * Pasal 22B
p12        I1 | Anggota Dewan Perwakilan Rakyat dapat diberhentikan dari jabatannya, yang syarat-
p12        I1 | syarat dan tata caranya diatur dalam undang-undang.
==================== PAGE 13 ====================
p13        I1 | * Perubahan II 18 Agustus 2000.
p13 B      | * BAB VIIA
p13 B      | DEWAN PERWAKILAN DAERAH
p13 B      | * Pasal 22C
p13        I1 [AYAT]
p13        I1 | (1)
p13        I2 | Anggota Dewan Perwakilan Daerah dipilih dari setiap provinsi melalui pemilihan
p13        I2 | umum.
p13        I1 [AYAT]
p13        I1 | (2)
p13        I2 | Anggota Dewan Perwakilan Daerah dari setiap provinsi jumlahnya sama dan
p13        I2 | jumlah seluruh anggota Dewan Perwakilan Daerah itu tidak lebih dari sepertiga
p13        I2 | jumlah anggota Dewan Perwakilan Rakyat.
p13        I1 [AYAT]
p13        I1 | (3)
p13        I2 | Dewan Perwakilan Daerah bersidang sedikitnya sekali dalam setahun.
p13        I1 [AYAT]
p13        I1 | (4)
p13        I2 | Susunan dan kedudukan Dewan Perwakilan Daerah diatur dengan undang-
p13        I2 | undang.
p13        I2 | * Perubahan III 9 November 2001.
p13 B      | * Pasal 22D
p13        I1 [AYAT]
p13        I1 | (1)
p13        I2 | Dewan Perwakilan Daerah dapat mengajukan kepada Dewan Perwakilan Rakyat
p13        I2 | rancangan undang-undang yang berkaitan dengan otonomi daerah, hubungan
p13        I2 | pusat dan daerah, pembentukan dan pemekaran serta penggabungan daerah,
p13        I2 | pengelolaan sumber daya alam dan sumber daya ekonomi lainnya, serta yang
p13        I2 | berkaitan dengan perimbangan keuangan pusat dan daerah.
p13        I1 [AYAT]
p13        I1 | (2)
p13        I2 | Dewan Perwakilan Daerah ikut membahas rancangan undang-undang yang
p13        I2 | berkaitan dengan otonomi daerah; hubungan pusat dan daerah; pembentukan,
p13        I2 | pemekaran, dan penggabungan daerah; pengelolaan sumber daya alam dan
p13        I2 | sumber daya ekonomi lainnya, serta perimbangan keuangan pusat dan daerah;
p13        I2 | serta memberikan pertimbangan kepada Dewan Perwakilan Rakyat atas
p13        I2 | rancangan undang-undang anggaran pendapatan dan belanja negara dan
p13        I2 | rancangan undang-undang yang berkaitan dengan pajak, pendidikan, dan agama.
p13        I1 [AYAT]
p13        I1 | (3)
p13        I2 | Dewan Perwakilan Daerah dapat melakukan pengawasan atas pelaksanaan
p13        I2 | undang-undang mengenai : otonomi daerah, pembentukan, pemekaran dan
p13        I2 | penggabungan daerah, hubungan pusat dan daerah, pengelolaan sumber daya
p13        I2 | alam dan sumber daya ekonomi lainnya, pelaksanaan anggaran pendapatan dan
p13        I2 | belanja negara, pajak, pendidikan, dan agama serta menyampaikan hasil
p13        I2 | pengawasannya itu kepada Dewan Perwakilan Rakyat sebagai bahan
p13        I2 | pertimbangan untuk ditindaklanjuti.
p13        I1 [AYAT]
p13        I1 | (4)
p13        I2 | Anggota Dewan Perwakilan Daerah dapat diberhentikan dari jabatannya, yang
p13        I2 | syarat-syarat dan tata caranya diatur dalam undang-undang.
p13        I2 | * Perubahan III 9 November 2001.
==================== PAGE 14 ====================
p14 B      [HEADING:BAB]
p14 B      | BAB VIIB
p14 B      | PEMILIHAN UMUM
p14 B      | * Pasal 22E
p14        I1 [AYAT]
p14        I1 | (1)
p14        I2 | Pemilihan umum dilaksanakan secara langsung, umum, bebas, rahasia, jujur, dan
p14        I2 | adil setiap lima tahun sekali.
p14        I1 [AYAT]
p14        I1 | (2)
p14        I2 | Pemilihan umum diselenggarakan untuk memilih anggota Dewan Perwakilan
p14        I2 | Rakyat, Dewan Perwakilan Daerah, Presiden dan Wakil Presiden dan Dewan
p14        I2 | Perwakilan Rakyat Daerah.
p14        I1 [AYAT]
p14        I1 | (3)
p14        I2 | Peserta pemilihan umum untuk memilih anggota Dewan Perwakilan Rakyat dan
p14        I2 | anggota Dewan Perwakilan Rakyat Daerah adalah partai politik.
p14        I1 [AYAT]
p14        I1 | (4)
p14        I2 | Peserta pemilihan umum untuk memilih anggota Dewan Perwakilan Daerah
p14        I2 | adalah perseorangan.
p14        I1 [AYAT]
p14        I1 | (5)
p14        I2 | Pemilihan umum diselenggarakan oleh suatu komisi pemilihan umum yang
p14        I2 | bersifat nasional, tetap, dan mandiri.
p14        I1 [AYAT]
p14        I1 | (6)
p14        I2 | Ketentuan lebih lanjut tentang pemilihan umum diatur dengan undang-undang.
p14        I2 | * Perubahan III 9 November 2001.
p14 B      [HEADING:BAB]
p14 B      | BAB VIII
p14 B      | HAL KEUANGAN
p14 B      | Pasal 23
p14        I1 | * (1)
p14        I2 | Anggaran Pendapatan dan Belanja Negara sebagai wujud dari pengelolaan
p14        I2 | keuangan negara ditetapkan setiap tahun dengan undang-undang dan dilaksanakan
p14        I2 | secara terbuka dan bertanggung jawab untuk sebesar-besarnya kemakmuran
p14        I2 | rakyat.
p14        I1 | * (2)
p14        I2 | Rancangan undang-undang anggaran pendapatan dan belanja negara diajukan
p14        I2 | oleh Presiden untuk dibahas bersama Dewan Perwakilan Rakyat dengan
p14        I2 | memperhatikan pertimbangan Dewan Perwakilan Daerah.
p14        I1 | * (3)
p14        I2 | Apabila Dewan Perwakilan Rakyat tidak menyetujui rancangan anggaran
p14        I2 | pendapatan dan belanja negara yang diusulkan oleh Presiden, Pemerintah
p14        I2 | menjalankan Anggaran Pendapatan dan Belanja Negara tahun yang lalu.
p14        I2 | * Perubahan III 9 November 2001, sebelumnya berbunyi :
p14        I3 [AYAT]
p14        I3 | (1)
p14        | Anggaran Pendapatan dan Belanja ditetapkan tiap-tiap tahun
p14        | dengan Undang-undang. Apabila Dewan Perwakilan Rakyat tidak
p14        | menyetujui anggaran yang diusulkan Pemerintah, maka Pemerintah
p14        | menjalankan anggaran tahun yang lalu.
p14        I3 [AYAT]
p14        I3 | (2)
p14        | Segala pajak untuk keperluan negara berdasarkan Undang-undang.
p14        I3 [AYAT]
p14        I3 | (3)
p14        | Macam dan harga mata uang ditetapkan dengan Undang-undang.
p14        I3 [AYAT]
p14        I3 | (4)
p14        | Hal keuangan negara selanjutnya diatur dengan Undang-undang.
p14        I3 [AYAT]
p14        I3 | (5)
p14        | Untuk memeriksa tanggung-jawab tentang keuangan negara
p14        | diadakan suatu Badan Pemeriksa Keuangan yang peraturannya
==================== PAGE 15 ====================
p15        | ditetapkan
p15        | dengan
p15        | Undang-undang.
p15        I5 | Hasil
p15        I7 | pemeriksaan
p15        | itu
p15        | diberitahukan kepada Dewan Perwakilan Rakyat.
p15 B      | * Pasal 23A
p15        I2 | Pajak dan pungutan lain yang bersifat memaksa untuk keperluan negara diatur
p15        I1 | dengan undang-undang.
p15        I1 | * Perubahan III 9 November 2001.
p15 B      | * Pasal 23B
p15        I2 | Macam dan harga mata uang ditetapkan dengan undang-undang.
p15        I1 | * Perubahan IV 10 Agustus 2002
p15 B      | * Pasal 23C
p15        I2 | Hal-hal lain mengenai keuangan negara diatur dengan undang-undang.
p15        I1 | * Perubahan III 9 November 2001.
p15 B      | * Pasal 23D
p15        I2 | Negara memiliki suatu bank sentral yang susunan, kedudukan, kewenangan,
p15        I1 | tanggung jawab, dan independensinya diatur dengan undang-undang.
p15        I1 | * Perubahan IV 10 Agustus 2002.
p15 B      [HEADING:BAB]
p15 B      | BAB VIIIA
p15 B      | BADAN PEMERIKSA KEUANGAN
p15 B      | * Pasal 23E
p15        I1 [AYAT]
p15        I1 | (1)
p15        I2 | Untuk memeriksa pengelolaan dan tanggung jawab tentang keuangan negara
p15        I2 | diadakan suatu Badan Pemeriksa Keuangan yang bebas dan mandiri.
p15        I1 [AYAT]
p15        I1 | (2)
p15        I2 | Hasil pemeriksa keuangan negara diserahkan kepada Dewan Perwakilan Rakyat,
p15        I2 | Dewan Perwakilan Daerah, dan Dewan Perwakilan Rakyat Daerah, sesuai dengan
p15        I2 | kewenangannya.
p15        I1 [AYAT]
p15        I1 | (3)
p15        I2 | Hasil pemeriksaan tersebut ditindaklanjuti oleh lembaga perwakilan dan/atau
p15        I2 | badan sesuai dengan undang-undang.
p15        I2 | * Perubahan III 9 November 2001.
==================== PAGE 16 ====================
p16 B      | * Pasal 23F
p16        I1 [AYAT]
p16        I1 | (1)
p16        I2 | Anggota Badan Pemeriksa Keuangan dipilih oleh Dewan Perwakilan Rakyat
p16        I2 | dengan memperhatikan pertimbangan Dewan Perwakilan Daerah dan diresmikan
p16        I2 | oleh Presiden.
p16        I1 [AYAT]
p16        I1 | (2)
p16        I2 | Pimpinan Badan Pemeriksa Keuangan dipilih dari dan oleh anggota.
p16        I2 | * Perubahan III 9 November 2001.
p16 B      | * Pasal 23G
p16        I1 [AYAT]
p16        I1 | (1)
p16        I2 | Badan Pemeriksa Keuangan berkedudukan di ibu kota negara, dan memiliki
p16        I2 | perwakilan di setiap provinsi.
p16        I1 [AYAT]
p16        I1 | (2)
p16        I2 | Ketentuan lebih lanjut mengenai Badan Pemeriksa Keuangan diatur dengan
p16        I2 | undang-undang.
p16        I2 | * Perubahan III 9 November 2001.
p16 B      [HEADING:BAB]
p16 B      | BAB IX
p16 B      | KEKUASAAN KEHAKIMAN
p16 B      | Pasal 24
p16        I1 | * (1)
p16        I2 | Kekuasaan
p16        | kehakiman
p16        | merupakan
p16        I4 | kekuasaan
p16        I5 | yang
p16        I7 | merdeka
p16        | untuk
p16        I2 | menyelenggarakan peradilan guna menegakkan hukum dan keadilan.
p16        I1 | * (2)
p16        I2 | Kekuasaan kehakiman dilakukan oleh sebuah Mahkamah Agung dan badan
p16        I2 | peradilan yang berada di bawahnya dalam lingkungan peradilan umum,
p16        I2 | lingkungan peradilan agama, lingkungan peradilan militer, lingkungan peradilan
p16        I2 | tata usaha negara, dan oleh sebuah Mahkamah Konstitusi.
p16        I2 | * Perubahan III 19 November 2001, sebelumnya berbunyi :
p16        I3 [AYAT]
p16        I3 | (1)
p16        | Kekuasaan kehakiman dilakukan oleh sebuah Mahkamah Agung dan
p16        | lain-lain badan kehakiman menurut Undang-undang.
p16        I3 [AYAT]
p16        I3 | (2)
p16        | Susunan dan kekuasaan Badan-badan Kehakiman itu diatur dengan
p16        | Undang-undang.
p16        I1 | * (3)
p16        I2 | Badan-badan lain yang fungsinya berkaitan dengan kekuasaan kehakiman diatur
p16        I2 | dalam undang-undang.
p16        I2 | * Perubahan IV 10 Agustus 2002.
p16 B      | * Pasal 24A
p16        I1 [AYAT]
p16        I1 | (1)
p16        I2 | Mahkamah Agung berwenang mengadili pada tingkat kasasi, meguji peraturan
p16        I2 | perundang-undangan di bawah undang-undang terhadap undang-undang, dan
p16        I2 | mempunyai wewenang lainnya yang diberikan oleh undang-undang.
==================== PAGE 17 ====================
p17        I1 [AYAT]
p17        I1 | (2)
p17        I2 | Hakim Agung harus memiliki integritas dan kepribadian yang tidak tercela, adil,
p17        I2 | profesional, dan berpengalaman di bidang hukum.
p17        I1 [AYAT]
p17        I1 | (3)
p17        I2 | Calon hakim agung diusulkan Komisi Yudisial kepada Dewan Perwakilan Rakyat
p17        I2 | untuk mendapatkan persetujuan dan selanjutnya ditetapkan sebagai hakim agung
p17        I2 | oleh Presiden.
p17        I1 [AYAT]
p17        I1 | (4)
p17        I2 | Ketua dan wakil ketua Mahkamah Agung dipilih dari dan oleh hakim agung.
p17        I1 [AYAT]
p17        I1 | (5)
p17        I2 | Susunan, kedudukan, keanggotaan, dan hukum acara Mahkamah Agung serta
p17        I2 | badan peradilan di bawahnya diatur dengan undang-undang.
p17        I2 | * Perubahan III 19 November 2001.
p17 B      | * Pasal 24B
p17        I1 [AYAT]
p17        I1 | (1)
p17        I2 | Komisi Yudisial bersifat mandiri yang berwenang mengusulkan pengangkatan
p17        I2 | hakim agung dan mempunyai wewenang lain dalam rangka menjaga dan
p17        I2 | menegakkan kehormatan, keluhuran martabat, serta perilaku hakim.
p17        I1 [AYAT]
p17        I1 | (2)
p17        I2 | Anggota Komisi Yudisial harus mempunyai pengetahuan dan pengalaman di
p17        I2 | bidang hukum serta memiliki integritas dan kepribadian yang tidak tercela.
p17        I1 [AYAT]
p17        I1 | (3)
p17        I2 | Anggota Komisi Yudisial diangkat dan diberhentikan oleh Presiden dengan
p17        I2 | persetujuan Dewan Perwakilan Rakyat.
p17        I1 [AYAT]
p17        I1 | (4)
p17        I2 | Susunan, kedudukan, dan keanggotaan Komisi Yudisial diatur dengan undang-
p17        I2 | undang.
p17        I2 | * Perubahan III 9 November 2001.
p17 B      | * Pasal 24C
p17        I1 [AYAT]
p17        I1 | (1)
p17        I2 | Mahkamah Konstitusi berwenang mengadili pada tingkat pertama dan terakhir
p17        I2 | yang putusannya bersifat final untuk menguji undang-undang terhadap Undang-
p17        I2 | Undang Dasar, memutus sengketa kewenangan lembaga negara yang
p17        I2 | kewenangannya diberikan oleh Undang-Undang Dasar, memutus pembubaran
p17        I2 | partai politik, dan memutus perselisihan tentang hasil pemilihan umum.
p17        I1 [AYAT]
p17        I1 | (2)
p17        I2 | Mahkamah Konstitusi wajib memberikan putusan atas pendapat Dewan
p17        I2 | Perwakilan Rakyat mengenai dugaan pelanggaran oleh Presiden dan/atau Wakil
p17        I2 | Presiden menurut Undang-Undang Dasar.
p17        I1 [AYAT]
p17        I1 | (3)
p17        I2 | Mahkamah Konstitusi mempunyai sembilan orang anggota hakim konstitusi yang
p17        I2 | ditetapkan oleh Presiden, yang diajukan masing-masing tiga orang oleh
p17        I2 | Mahkamah Agung, tiga orang oleh Dewan Perwakilan Rakyat, dan tiga orang
p17        I2 | oleh Presiden.
p17        I1 [AYAT]
p17        I1 | (4)
p17        I2 | Ketua dan Wakil Ketua Mahkamah Konstitusi dipilih dari dan oleh hakim
p17        I2 | konstitusi.
p17        I1 [AYAT]
p17        I1 | (5)
p17        I2 | Hakim Konstitusi harus memiliki integritas dan kepribadian yang tidak tercela,
p17        I2 | adil, negarawan yang menguasai konstitusi dan ketatanegaraan, serta tidak
p17        I2 | merangkap sebagai pejabat negara.
p17        I1 [AYAT]
p17        I1 | (6)
p17        I2 | Pengangkatan dan pemberhentian hakim konstitusi, hukum acara serta ketentuan
p17        I2 | lainnya tentang Mahkamah Konstitusi diatur dengan undang-undang.
==================== PAGE 18 ====================
p18        I2 | * Perubahan III 9 November 2001.
p18 B      | Pasal 25
p18        I1 | Syarat-syarat untuk menjadi dan untuk diberhentikan sebagai hakim ditetapkan dengan
p18        I1 | Undang-undang.
p18 B      | * BAB IX A
p18 B      | WILAYAH NEGARA
p18 B      | * Pasal 25A
p18        I1 | Negara Kesatuan Republik Indonesia adalah sebuah Negara kepulauan yang berciri
p18        I1 | Nusantara dengan wilayah yang batas-batas dan hak-haknya ditetapkan dengan undang-
p18        I1 | undang.
p18        I1 | * Perubahan II, 18 Agustus 2000.
p18 B      [HEADING:BAB]
p18 B      | BAB X
p18 B      | WARGA NEGARA DAN PENDUDUK
p18 B      | * Pasal 26
p18        I1 | * (1)
p18        I2 | Penduduk ialah warga negara Indonesia dan orang asing yang bertempat tinggal
p18        I2 | di Indonesia.
p18        I1 | * (2)
p18        I2 | Setiap warga negara dan penduduk diatur dengan undang-undang.
p18        I2 | Perubahan II 18 Agustus 2000, sebelumnya berbunyi :
p18        I3 | WARGA NEGARA
p18        I3 | Pasal 26
p18        [AYAT]
p18        | (1)
p18        | Yang menjadi Warga Negara ialah orang-orang bangsa
p18        | Indonesia asli dan orang-orang bangsa lain yang disyahkan
p18        | dengan undang-undang sebagai Warga Negara.
p18        [AYAT]
p18        | (2)
p18        | Syarat-syarat yang mengenai kewargaan negara ditetapkan
p18        | dengan undang-undang.
p18 B      | Pasal 27
p18        I1 [AYAT]
p18        I1 | (1)
p18        I2 | Segala warga negara bersamaan kedudukannya di dalam hukum dan
p18        I2 | pemerintahan dan wajib menjunjung hukum dan pemerintahan itu dengan tidak
p18        I2 | ada kecualinya.
p18        I1 [AYAT]
p18        I1 | (2)
p18        I2 | Tiap-tiap warga negara berhak atas pekerjaan dan penghidupan yang layak bagi
p18        I2 | kemanusiaan.
p18        I1 | * (3)
p18        I2 | Setiap warga negara berhak dan wajib ikut serta dalam upaya pembelaan negara.
p18        I2 | * Perubahan II 18 Agustus 2000.
==================== PAGE 19 ====================
p19 B      | Pasal 28
p19        I1 | Kemerdekaan berserikat dan berkumpul, mengeluarkan pikiran dengan lisan dan tulisan
p19        I1 | dan sebagainya ditetapkan dengan Undang-undang.
p19 B      | * BAB XA
p19 B      | HAK ASASI MANUSIA
p19 B      | * Pasal 28A
p19        I1 | Setiap orang berhak untuk hidup serta berhak mempertahankan hidup dan kehidupannya.
p19        I1 | * Perubahan II 18 Agustus 2000.
p19 B      | * Pasal 28B
p19        I1 [AYAT]
p19        I1 | (1)
p19        I2 | Setiap orang berhak membentuk keluarga dan melanjutkan keturunan melalui
p19        I2 | perkawinan yang sah.
p19        I1 [AYAT]
p19        I1 | (2)
p19        I2 | Setiap anak berhak atas kelangsungan hidup, tumbuh, dan berkembang serta
p19        I2 | berhak atas perlindungan dari kekerasan dan diskriminasi.
p19        I2 | * Perubahan II 18 Agustus 2000.
p19 B      | * Pasal 28C
p19        I1 [AYAT]
p19        I1 | (1)
p19        I2 | Setiap orang berhak mengembangkan diri melalui pemenuhan kebutuhan
p19        I2 | dasarnya, berhak mendapat pendidikan dan memperoleh manfaat dari ilmu
p19        I2 | pengetahuan dan teknologi, seni dan budaya, demi meningkatkan kualitas
p19        I2 | hidupnya dan demi kesejahteraan umat manusia.
p19        I1 [AYAT]
p19        I1 | (2)
p19        I2 | Setiap orang berhak untuk memajukan dirinya dalam memperjuangkan haknya
p19        I2 | secara kolektif untuk membangun masyarakat, bangsa, dan negaranya.
p19        I2 | * Perubahan II 18 Agustus 2000.
p19 B      | * Pasal 28D
p19        I1 [AYAT]
p19        I1 | (1)
p19        I2 | Setiap orang berhak atas pengakuan, jaminan, perlindungan, dan kepastian hukum
p19        I2 | yang adil serta perlakuan yang sama di hadapan hukum.
p19        I1 [AYAT]
p19        I1 | (2)
p19        I2 | Setiap orang berhak untuk bekerja serta mendapat imbalan dan perlakuan yang
p19        I2 | adil dan layak dalam hubungan kerja.
p19        I1 [AYAT]
p19        I1 | (3)
p19        I2 | Setiap warga negara berhak memperoleh kesempatan yang sama dalam
p19        I2 | pemerintahan.
p19        I1 [AYAT]
p19        I1 | (4)
p19        I2 | Setiap orang berhak atas status kewarganegaraan.
p19        I2 | * Perubahan II 18 Agustus 2000.
==================== PAGE 20 ====================
p20 B      | * Pasal 28E
p20        I1 [AYAT]
p20        I1 | (1)
p20        I2 | Setiap orang bebas memeluk agama dan beribadat menurut agamanya, memilih
p20        I2 | pendidikan dan pengajaran, memilih pekerjaan, memilih kewarganegaraan,
p20        I2 | memilih tempat tinggal di wilayah negara dan meninggalkannya, serta berhak
p20        I2 | kembali.
p20        I1 [AYAT]
p20        I1 | (2)
p20        I2 | Setiap orang berhak atas kebebasan meyakini kepercayaan, menyatakan pikiran
p20        I2 | dan sikap, sesuai dengan hati nuraninya.
p20        I1 [AYAT]
p20        I1 | (3)
p20        I2 | Setiap orang berhak atas kebebasan berserikat, berkumpul, dan mengeluarkan
p20        I2 | pendapat.
p20        I2 | * Perubahan II 18 Agustus 2000.
p20 B      | * Pasal 28F
p20        I1 | Setiap orang berhak untuk berkomunikasi dan memperoleh informasi untuk
p20        I1 | mengembangkan pribadi dan lingkungan sosialnya, serta berhak untuk mencari,
p20        I1 | memperoleh, memiliki, menyimpan, mengolah, dan menyampaikan informasi dengan
p20        I1 | menggunakan segala jenis saluran yang tersedia.
p20        I1 | * Perubahan II 18 Agustus 2000.
p20 B      | * Pasal 28G
p20        I1 [AYAT]
p20        I1 | (1)
p20        I2 | Setiap orang berhak atas perlindungan diri pribadi, keluarga, kehormatan,
p20        I2 | martabat, dan harta benda yang dibawah kekuasaannya, serta berhak atas rasa
p20        I2 | aman dan perlindungan dari ancaman ketakutan untuk berbuat atau tidak berbuat
p20        I2 | sesuatu yang merupakan hak asasi.
p20        I1 [AYAT]
p20        I1 | (2)
p20        I2 | Setiap orang berhak untuk bebas dari penyiksaan atau perlakuan yang
p20        I2 | merendahkan derajat martabat manusia dan berhak memperoleh suaka politik dari
p20        I2 | negara lain.
p20        I2 | * Perubahan II 18 Agustus 2000.
p20 B      | * Pasal 28H
p20        I1 [AYAT]
p20        I1 | (1)
p20        I2 | Setiap orang berhak hidup sejahtera lahir dan batin, bertempat tinggal, dan
p20        I2 | mendapatkan lingkungan hidup yang baik dan sehat serta berhak memperoleh
p20        I2 | pelayanan kesehatan.
p20        I1 [AYAT]
p20        I1 | (2)
p20        I2 | Setiap orang berhak mendapat kemudahan dan perlakuan khusus untuk
p20        I2 | memperoleh kesempatan dan manfaat yang sama guna mencapai persamaan dan
p20        I2 | keadilan.
p20        I1 [AYAT]
p20        I1 | (3)
p20        I2 | Setiap orang berhak atas jaminan sosial yang memungkinkan pengembangan
p20        I2 | dirinya secara utuh sebagai manusia yang bermartabat.
p20        I1 [AYAT]
p20        I1 | (4)
p20        I2 | Setiap orang berhak mempunyai hak milik pribadi dan hak milik tersebut tidak
p20        I2 | boleh diambil alih secara sewenang-wenang oleh siapa pun.
==================== PAGE 21 ====================
p21        I2 | * Perubahan II 18 Agustus 2000.
p21 B      | * Pasal 28I
p21        I1 [AYAT]
p21        I1 | (1)
p21        I2 | Hak untuk hidup, hak untuk tidak disiksa, hak kemerdekaan pikiran dan hati
p21        I2 | nurani, hak beragama, hak untuk tidak diperbudak, hak untuk diakui sebagai
p21        I2 | pribadi di hadapan hukum, dan hak untuk tidak dituntut atas dasar hukum yang
p21        I2 | berlaku surut adalah hak asasi manusia yang tidak dapat dikurangi dalam keadaan
p21        I2 | apa pun.
p21        I1 [AYAT]
p21        I1 | (2)
p21        I2 | Setiap orang berhak bebas dari perlakuan yang bersifat diskriminatif atas dasar
p21        I2 | apa pun dan berhak mendapatkan perlindungan terhadap perlakuan yang bersifat
p21        I2 | diskriminatif itu.
p21        I1 [AYAT]
p21        I1 | (3)
p21        I2 | Identitas budaya dan hak masyarakat tradisional dihormati selaras dengan
p21        I2 | perkembangan zaman dan peradaban.
p21        I1 [AYAT]
p21        I1 | (4)
p21        I2 | Perlindungan, pemajuan, penegakan, dan pemenuhan hak asasi manusia adalah
p21        I2 | tanggung jawab negara, terutama pemerintah.
p21        I1 [AYAT]
p21        I1 | (5)
p21        I2 | Untuk menegakkan dan melindungi hak asasi manusia dengan prinsip negara
p21        I2 | hukum yang demokratis, maka pelaksanaan hak asasi manusia dijamin, diatur,
p21        I2 | dan dituangkan dalam peraturan perundang-undangan.
p21        I2 | * Perubahan II 18 Agustus 2000.
p21 B      | * Pasal 28J
p21        I1 [AYAT]
p21        I1 | (1)
p21        I2 | Setiap orang wajib menghormati hak asasi manusia orang lain dalam tertib
p21        I2 | kehidupan bermasyarakat, berbangsa, dan bernegara.
p21        I1 [AYAT]
p21        I1 | (2)
p21        I2 | Dalam menjalankan hak dan kebebasannya, setiap orang wajib tunduk kepada
p21        I2 | pembatasan yang ditetapkan dengan undang-undang dengan maksud semata-mata
p21        I2 | untuk menjamin pengakuan serta penghormatan atas hak dan kebebasan orang
p21        I2 | lain dan untuk memenuhi tuntutan yang adil sesuai dengan pertimbangan moral,
p21        I2 | nilai-nilai agama, keamanan, dan ketertiban umum dalam suatu masyarakat
p21        I2 | demokratis.
p21        I2 | * Perubahan II 18 Agustus 2000.
p21 B      [HEADING:BAB]
p21 B      | BAB XI
p21 B      | AGAMA
p21 B      | Pasal 29
p21        I1 [AYAT]
p21        I1 | (1)
p21        I2 | Negara berdasar atas Ketuhahan Yang Maha Esa.
p21        I1 [AYAT]
p21        I1 | (2)
p21        I2 | Negara menjamin kemerdekaan tiap-tiap penduduk untuk memeluk agamanya
p21        I2 | masing-masing dan untuk beribadat menurut agamnya dan kepercayaannya itu.
==================== PAGE 22 ====================
p22 B      [HEADING:BAB]
p22 B      | BAB XII
p22 B      I3 | PERTAHANAN DAN KEAMANAN NEGARA
p22 B      | Pasal 30
p22        I1 | * (1)
p22        I2 | Tiap-tiap warga negara berhak dan wajib ikut serta dalam usaha pertahanan dan
p22        I2 | keamanan negara.
p22        I1 | * (2)
p22        I2 | Untuk pertahanan dan keamanan negara dilaksanakan melalui sistem pertahanan
p22        I2 | dan keamanan rakyat semesta oleh Tentara Nasional Indonesia dan Kepolisian
p22        I2 | Negara Republik Indonesia, sebagai kekuatan utama, dan rakyat, sebagai
p22        I2 | kekuatan pendukung.
p22        I1 | * (3)
p22        I2 | Tentara Nasional Indonesia terdiri atas Angkatan Darat, Angkatan Laut, dan
p22        I2 | Angkatan Udara sebagai alat negara bertugas mempertahankan, melindungi, dan
p22        I2 | memelihara keutuhan dan kedaulatan negara.
p22        I1 | * (4)
p22        I2 | Kepolisian Negara Republik Indonesia sebagai alat negara yang menjaga
p22        I2 | keamanan dan ketertiban masyarakat bertugas melindungi, mengayomi, melayani
p22        I2 | masyarakat, serta menegakkan hukum.
p22        I1 | * (5)
p22        I2 | Susunan dan kedudukan Tentara Nasional Indonesia, Kepolisian Negara Republik
p22        I2 | Indonesia, hubungan kewenangan Tentara Nasional Indonesia dan Kepolisian
p22        I2 | Negara Republik Indonesia di dalam menjalankan tugasnya, syarat-syarat
p22        I2 | keikutsertaan warga negara dalam usaha pertahanan dan keamanan negara, serta
p22        I2 | hal-hal yang terkait dengan pertahanan dan keamanan diatur dengan undang-
p22        I2 | undang.
p22        I2 | * Perubahan II 18 Agustus 2000, sebelumnya berbunyi :
p22        I3 | PERTAHANAN NEGARA
p22        I3 | Pasal 30
p22        [AYAT]
p22        | (1)
p22        | Tiap-tiap warga negara berhak dan wajib ikut serta dalam
p22        | usaha pembelaan negara.
p22        [AYAT]
p22        | (2)
p22        | Syarat-syarat tentang pembelaan diatur dengan Undang-
p22        | undang.
p22 B      [HEADING:BAB]
p22 B      | BAB XIII
p22 B      | PENDIDIKAN DAN KEBUDAYAAN
p22 B      | * Pasal 31
p22        I1 | * (1)
p22        I2 | Setiap warga negara berhak mendapat pendidikan.
p22        I1 | * (2)
p22        I2 | Setiap warga negara wajib mengikuti pendidikan dasar dan pemerintah wajib
p22        I2 | membiayainya.
p22        I1 | * (3)
p22        I2 | Pemerintah mengusahakan dan menyelenggarakan satu sistem pendidikan
p22        I2 | nasional, yang meningkatkan keimanan dan ketakwaan serta akhlak mulia dalam
p22        I2 | rangka mencerdaskan kehidupan bangsa, yang diatur dengan undang-undang.
p22        I1 | * (4)
p22        I2 | Negara memprioritaskan anggaran pendidikan sekurang-kurangnya dua puluh
p22        I2 | persen dari anggaran pendapatan dan belanja negara serta dari anggaran
==================== PAGE 23 ====================
p23        I2 | pendapatan dan belanja daerah untuk memenuhi kebutuhan penyelenggaraan
p23        I2 | pendidikan nasional.
p23        I1 | * (5)
p23        I2 | Pemerintah memajukan ilmu pengetahuan dan teknologi dengan menjunjung
p23        I2 | tinggi nilai-nilai agama dan persatuan bangsa untuk kemajuan peradaban serta
p23        I2 | kesejahteraan umat manusia.
p23        I2 | Perubahan IV 10 Agustus 2002, sebelumnya berbunyi :
p23        I3 | PENDIDIKAN
p23        I3 | Pasal 31
p23        [AYAT]
p23        | (1)
p23        | Tiap-tiap Warga Negara berhak mendapat pengajaran.
p23        [AYAT]
p23        | (2)
p23        | Pemerintah mengusahakan dan menyelenggarakan satu
p23        | sistem pengajaran nasional, yang diatur dengan Undang-
p23        | undang.
p23 B      | * Pasal 32
p23        I1 | * (1)
p23        I2 | Negara memajukan kebudayaan nasional Indonesia di tengah peradaban dunia
p23        I2 | dengan menjamin kebebasan masyarakat dalam memelihara dan mengembangkan
p23        I2 | nilai-nilai budayanya.
p23        I1 | * (2)
p23        I2 | Negara menghormati dan memelihara bahasa daerah sebagai kekayaan budaya
p23        I2 | nasional.
p23        I2 | * Perubahan IV 10 Agustus 2002, sebelumnya berbunyi :
p23        I3 | Pemerintah memajukan kebudayaan nasional Indonesia.
p23 B      | * BAB XIV
p23 B      | * PEREKONOMIAN NASIONAL DAN
p23 B      | KESEJAHTERAAN SOSIAL
p23 B      | Pasal 33
p23        I1 [AYAT]
p23        I1 | (1)
p23        I2 | Perekonomian disusun sebagai usaha bersama berdasar atas asas kekeluargaan.
p23        I1 [AYAT]
p23        I1 | (2)
p23        I2 | Cabang-cabang produksi yang penting bagi negara dan yang menguasai hajat
p23        I2 | hidup orang banyak dikuasai oleh negara.
p23        I1 [AYAT]
p23        I1 | (3)
p23        I2 | Bumi dan air dan kekayaan alam yang terkandung didalamnya dikuasai oleh
p23        I2 | Negara dan dipergunakan untuk sebesar-besarnya kemakmuran rakyat.
p23        I2 | * Perubahan IV 10 Agustus 2002, sebelumnya berbunyi :
p23        I3 [HEADING:BAB]
p23        I3 | BAB XIV
p23        I3 | KESEJAHTERAAN SOSIAL
p23        I1 | * (4)
p23        I2 | Perekonomian nasional diselenggarakan berdasar atas demokrasi ekonomi dengan
p23        I2 | prinsip
p23        I3 | kebersamaan,
p23        | efisiensi
p23        I4 | berkeadilan,
p23        I5 | berkelanjutan,
p23        | berwawasan
p23        I2 | lingkungan, kemandirian, serta dengan menjaga keseimbangan kemajuan dan
p23        I2 | kesatuan ekonomi nasional.
==================== PAGE 24 ====================
p24        I1 | * (5)
p24        I2 | Ketentuan lebih lanjut mengenai pelaksanaan pasal ini diatur dalam undang-
p24        I2 | undang.
p24        I2 | * Perubahan IV 10 Agustus 2002.
p24 B      | * Pasal 34
p24        I1 [AYAT]
p24        I1 | (1)
p24        I2 | Fakir miskin dan anak-anak yang terlantar dipelihara oleh negara.
p24        I1 [AYAT]
p24        I1 | (2)
p24        I2 | Negara mengembangkan sistem jaminan sosial bagi seluruh rakyat dan
p24        I2 | memberdayakan masyarakat yang lemah dan tidak mampu sesuai dengan
p24        I2 | martabat kemanusiaan.
p24        I1 [AYAT]
p24        I1 | (3)
p24        I2 | Negara bertanggung jawab atas penyediaan fasilitas pelayanan kesehatan dan
p24        I2 | fasilitas pelayanan umum yang layak.
p24        I1 [AYAT]
p24        I1 | (4)
p24        I2 | Ketentuan lebih lanjut mengenai pelaksanaan pasal ini diatur dalam undang-
p24        I2 | undang.
p24        I2 | * Perubahan IV 10 Agustus 2002, sebelumnya berbunyi :
p24        I3 | Fakir miskin dan anak-anak yang terlantar dipelihara oleh Negara.
p24 B      | * BAB XV
p24 B      I3 | * BENDERA, BAHASA, DAN LAMBANG NEGARA,
p24 B      | SERTA LAGU KEBANGSAAN
p24        I1 | Perubahan II 18 Agustus 2000, sebelumnya berbunyi :
p24        I2 | BENDERA DAN BAHASA
p24 B      | Pasal 35
p24        I1 | Bendera Negara Indonesia ialah Sang Merah Putih.
p24 B      | Pasal 36
p24        I1 | Bahasa Negara ialah Bahasa Indonesia.
p24 B      | * Pasal 36A
p24        I1 | Lambang Negara ialah Garuda Pancasila dengan semboyan Bhineka Tunggal Ika.
p24        I1 | * Perubahan II 18 Agustus 2000.
p24 B      | * Pasal 36B
p24        I1 | Lagu Kebangsaan ialah Indonesia Raya.
p24        I1 | * Perubahan II 18 Agustus 2000.
==================== PAGE 25 ====================
p25 B      | * Pasal 36C
p25        I1 | Ketentuan lebih lanjut mengenai Bendera, Bahasa, dan Lambang Negara, serta Lagu
p25        I1 | Kebangsaan diatur dalam undang-undang.
p25        I1 | * Perubahan II 18 Agustus 2000.
p25 B      [HEADING:BAB]
p25 B      | BAB XVI
p25 B      | PERUBAHAN UNDANG-UNDANG DASAR
p25 B      | * Pasal 37
p25        I1 [AYAT]
p25        I1 | (1)
p25        I2 | Usul perubahan pasal-pasal Undang-Undang Dasar dapat diagendakan dalam
p25        I2 | sidang Majelis Permusyawaratan Rakyat apabila diajukan oleh sekurang-
p25        I2 | kurangnya 1/3 dari jumlah anggota Majelis Permusyawaratan Rakyat.
p25        I1 [AYAT]
p25        I1 | (2)
p25        I2 | Setiap usul perubahan pasal-pasal Undang-Undang Dasar diajukan secara tertulis
p25        I2 | dan ditunjukkan dengan jelas bagian yang diusulkan untuk diubah beserta
p25        I2 | alasannya.
p25        I1 [AYAT]
p25        I1 | (3)
p25        I2 | Untuk
p25        I3 | mengubah
p25        | pasal-pasal
p25        | Undang-Undang
p25        I5 | Dasar,
p25        I7 | sidang
p25        | Majelis
p25        I2 | Permusyawaratan Rakyat dihadiri oleh sekurang-kurangnya 2/3 dari jumlah
p25        I2 | anggota Majelis Permusyawaratan Rakyat.
p25        I1 [AYAT]
p25        I1 | (4)
p25        I2 | Putusan untuk mengubah pasal-pasal Undang-Undang Dasar dilakukan dengan
p25        I2 | persetujuan sekurang-kurangnya lima puluh persen ditambah satu anggota dari
p25        I2 | seluruh anggota Majelis Permusyawaratan Rakyat.
p25        I1 [AYAT]
p25        I1 | (5)
p25        I2 | Khusus mengenai bentuk Negara Kesatuan Republik Indonesia tidak dapat
p25        I2 | dilakukan perubahan.
p25        I2 | * Perubahan IV 10 Agustus 2002, sebelumnya berbunyi :
p25        I3 [AYAT]
p25        I3 | (1)
p25        | Untuk mengubah Undang-Undang Dasar sekurang-kurangnya 2/3
p25        | dari pada jumlah anggota Majelis Permusyawaratan Rakyat harus
p25        | hadir.
p25        I3 [AYAT]
p25        I3 | (2)
p25        | Putusan diambil dengan persetujuan sekurang-kurangnya 2/3 dari
p25        | pada jumlah anggota yang hadir.
p25 B      | ATURAN PERALIHAN
p25 B      | * Pasal I
p25        I2 | Segala peraturan perundang-undangan yang ada masih tetap berlaku selama
p25        I1 | belum diadakan yang baru menurut Undang-Undang Dasar ini.
p25 B      | * Pasal II
p25        I2 | Semua lembaga negara yang ada masih tetap berfungsi sepanjang untuk
p25        I1 | melaksanakan ketentuan Undang-Undang Dasar dan belum diadakan yang baru menurut
p25        I1 | Undang-Undang Dasar ini.
==================== PAGE 26 ====================
p26 B      | * Pasal III
p26        I2 | Mahkamah Konstitusi dibentuk selambat-lambatnya pada 17 Agustus 2003 dan
p26        I1 | sebelum dibentuk segala kewenangannya dilakukan oleh Mahkamah Agung.
p26        I2 | * Perubahan IV 10 Agustus 2002, sebelumnya berbunyi :
p26        I3 | Pasal I
p26        | Panitia Persiapan Kemerdekaan Indonesia mengatur dan
p26        | menyelenggarakan kepindahan pemerintahan kepada pemerintah
p26        | Indonesia.
p26        I3 | Pasal II
p26        | Segala Badan Negara dan Peraturan yang ada masih langsung
p26        | berlaku, selama belum diadakan yang baru menurut Undang-
p26        | Undang Dasar ini.
p26        I3 | Pasal III
p26        | Untuk pertama kali Presiden dan Wakil Presiden dipilih oleh
p26        | Panitia Persiapan Kemerdekaan Indonesia.
p26        I3 | Pasal IV
p26        | Sebelum Majelis Permusyawaratan Rakyat, Dewan Perwakilan
p26        | Rakyat dan Dewan Pertimbangan Agung dibentuk menurut
p26        | Undang-Undang Dasar ini, segala kekuasaannya dijalankan oleh
p26        | Presiden dengan bantuan sebuah komite nasional.
p26 B      | ATURAN TAMBAHAN
p26 B      | * Pasal I
p26        I2 | Majelis Permusyawaratan Rakyat ditugasi untuk melakukan peninjauan terhadap
p26        I1 | materi dan status hukum Ketetapan Majelis Permusyawaratan Rakyat Sementara dan
p26        I1 | Ketetapan Majelis Permusyawaratan Rakyat untuk diambil putusan pada sidang Majelis
p26        I1 | Permusyawaratan Rakyat tahun 2003.
p26 B      | * Pasal II
p26        I2 | Dengan ditetapkannya perubahan Undang-Undang Dasar ini, Undang-Undang
p26        I1 | Dasar Negara Republik Indonesia Tahun 1945 terdiri atas Pembukaan dan pasal-pasal.
p26        I2 | Perubahan tersebut diputuskan dalam Rapat Paripurna Majelis Permusyawaratan
p26        I1 | Rakyat Republik Indonesia ke-6 (lanjutan) tanggal 10 Agustus 2002 Sidang Tahunan
p26        I1 | Majelis Permusyawaratan Rakyat Republik Indonesia, dan mulai berlaku pada tanggal
p26        I1 | ditetapkan.
p26        I2 | * Perubahan IV 10 Agustus 2002, sebelumnya berbunyi :
==================== PAGE 27 ====================
p27        I3 [AYAT]
p27        I3 | (1)
p27        | Dalam enam bulan sesudah akhirnya peperangan Asia Timur Raya,
p27        | Presiden Indonesia mengatur dan menyelenggarakan segala hal
p27        | yang ditetapkan dalam Undang-Undang Dasar ini.
p27        I3 [AYAT]
p27        I3 | (2)
p27        | Dalam enam bulan sesudah Majelis Permusyawaratan Rakyat
p27        | dibentuk, Majelis itu bersidang untuk menetapkan Undang-Undang
p27        | Dasar.
```

---


## Putusan-MK

- **File**: `Putusan-MK/putusan_mkri_5301.pdf`
- **Document Type**: Putusan MK (Court Ruling)
- **Issued by**: Mahkamah Konstitusi
- **Pages**: 4 | **Lines**: 127
- **Font sizes**: [12.0]
- **Most common font**: 12.0 (100% of lines)
- **Bold font sizes**: [12.0]
- **Indent clusters**: [108.0, 126.0, 180.0, 251.0, 384.0, 415.0, 437.0, 480.0, 514.0]
- **Expected hierarchy**: Menimbang > MENGADILI > MEMUTUSKAN > Amar

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01 B      I8 | SALINAN
p01 B      | KETETAPAN
p01 B      I4 | Nomor 96/PUU-XVI/2018
p01 B      I2 | DEMI KEADILAN BERDASARKAN KETUHANAN YANG MAHA ESA
p01 B      I3 | MAHKAMAH KONSTITUSI REPUBLIK INDONESIA,
p01        I2 | Yang mengadili perkara konstitusi pada tingkat pertama dan terakhir
p01        I1 | menjatuhkan Ketetapan dalam perkara pengujian Kitab Undang-Undang Hukum
p01        I1 | Perdata terhadap Undang-Undang Dasar Negara Republik Indonesia Tahun 1945,
p01 B      I1 | sebagai berikut:
p01        I1 [PREAMBLE:MENIMBANG]
p01        I1 | Menimbang :  1.  Bahwa Mahkamah Konstitusi telah menerima permohonan
p01        I3 | bertanggal 7 November 2018 dari Jandi Mukianto, yang
p01        I3 | berdasarkan Surat Kuasa Khusus bertanggal 1 November 2018
p01        I3 | memberi kuasa kepada: Haris Satiadi, S.H., Suheru Prayitno,
p01        I3 | S.H., Rendy Alexander, S.H., Nikite Alvinta Bujangga, S.H., Praja
p01        I3 | Wibawa, S.H., dan Ocar Puspa Dewi, S.H., kesemuanya adalah
p01 B      I3 | Advokat dan Advokat Magang [sic!] pada Lembaga Bantuan
p01 B      I3 | Hukum Lentera Keadilan Rakyat (LBH LKRA) yang beralamat
p01        I3 | di Jalan Sukarjo Wiryopranoto Nomor 8D Gambir, Jakarta Pusat,
p01        I3 | yang diterima di Kepaniteraan Mahkamah Konstitusi pada
p01        I3 | tanggal 13 November 2018 dan dicatat dalam Buku Registrasi
p01 B      I3 | Perkara Konstitusi dengan Nomor 96/PUU-XVI/2018 pada
p01        I3 | tanggal 21 November 2018 perihal Permohonan Pengujian kata
p01        I3 | “Tionghoa” dalam Kitab Undang-Undang Hukum Perdata
p01        I3 | terhadap Undang-Undang Dasar Negara Republik Indonesia
p01        I3 | Tahun 1945;
p01        I3 [ITEM]
p01        I3 | 2.  Bahwa terhadap permohonan Nomor 96/PUU-XVI/2018 tersebut,
p01        I3 | Mahkamah Konstitusi telah menerbitkan:
p01        I3 [SUB-ITEM]
p01        I3 | a.  Ketetapan
p01        | Ketua
p01        | Mahkamah
p01        I6 | Konstitusi
p01        I9 | Nomor
p01        | 234/TAP.MK/2018 tentang Pembentukan Panel Hakim Untuk
p01        | Memeriksa Perkara Nomor 96/PUU-XVI/2018, bertanggal 21
p01        | November 2018;
==================== PAGE 2 ====================
p02        | 2
p02        I3 [SUB-ITEM]
p02        I3 | b.  Ketetapan Ketua Panel Hakim Mahkamah Konstitusi Nomor
p02        | 235/TAP.MK/2018 tentang Penetapan Hari Sidang Pertama
p02        | untuk
p02        I4 | memeriksa
p02        | perkara
p02        I5 | Nomor
p02        I7 | 96/PUU-XVI/2018,
p02        | bertanggal 21 November 2018;
p02        I3 [ITEM]
p02        I3 | 3.  Bahwa
p02        I4 | Mahkamah
p02        | telah
p02        I5 | menyelenggarakan
p02        I8 | Pemeriksaan
p02        I3 | Pendahuluan terhadap permohonan tersebut melalui Sidang
p02        I3 | Panel pada tanggal 6 Desember 2018;
p02        I3 [ITEM]
p02        I3 | 4. Bahwa Mahkamah telah menerima surat dari Pemohon
p02        I3 | bertanggal 17 Desember 2018 perihal penarikan permohonan
p02        I3 | yang diterima di Kepaniteraan Mahkamah Konstitusi pada
p02        I3 | tanggal 17 Desember 2018;
p02        I3 [ITEM]
p02        I3 | 5. Bahwa Mahkamah telah menyelenggarakan Sidang Panel pada
p02        I3 | tanggal 19 Desember 2018 dengan agenda menerima Perbaikan
p02        I3 | Permohonan dan sekaligus meminta konfirmasi perihal surat
p02        I3 | sebagaimana termaktub pada angka 4 di atas, namun Pemohon
p02        I3 | tidak hadir sekalipun telah dipanggil secara sah dan patut;
p02        I3 [ITEM]
p02        I3 | 6. Bahwa terhadap penarikan kembali permohonan Pemohon
p02        I3 | tersebut, Pasal 35 ayat (1) UU MK menyatakan, “Pemohon dapat
p02        I3 | menarik kembali Permohonan sebelum atau selama pemeriksaan
p02        I3 | Mahkamah Konstitusi dilakukan”;
p02        I3 [ITEM]
p02        I3 | 7. Bahwa Rapat Permusyawaratan Hakim pada tanggal 10 Januari
p02        I3 | 2019 telah menetapkan pencabutan atau penarikan kembali
p02        I3 | permohonan Nomor 96/PUU-XVI/2018 beralasan menurut hukum
p02        I3 | dan sesuai dengan Pasal 35 ayat (2) UU MK, penarikan kembali
p02        I3 | suatu Permohonan mengakibatkan Permohonan tersebut tidak
p02        I3 | dapat diajukan kembali.
p02        I1 [PREAMBLE:MENGINGAT]
p02        I1 | Mengingat :    1. Undang-Undang Dasar Negara Republik Indonesia Tahun 1945;
p02        I3 [ITEM]
p02        I3 | 2. Undang-Undang Nomor 24 Tahun 2003 tentang Mahkamah
p02        I3 | Konstitusi sebagaimana telah diubah dengan Undang-Undang
p02        I3 | Nomor 8 Tahun 2011 tentang Perubahan Atas Undang-Undang
p02        I3 | Nomor 24 Tahun 2003 tentang Mahkamah Konstitusi (Lembaran
p02        I3 | Negara Republik Indonesia Tahun 2011 Nomor 70, Tambahan
p02        I3 | Lembaran Negara Republik Indonesia Nomor 5226);
==================== PAGE 3 ====================
p03        | 3
p03        I3 [ITEM]
p03        I3 | 3. Undang-Undang Nomor 48 Tahun 2009 tentang Kekuasaan
p03        I3 | Kehakiman (Lembaran Negara Republik Indonesia Tahun 2009
p03        I3 | Nomor 157, Tambahan Lembaran Negara Republik Indonesia
p03        I3 | Nomor 5076);
p03 B      [PREAMBLE:MENETAPKAN]
p03 B      | MENETAPKAN:
p03        I1 [ITEM]
p03        I1 | 1. Mengabulkan penarikan kembali permohonan Pemohon;
p03        I1 [ITEM]
p03        I1 | 2. Menyatakan Permohonan Nomor 96/PUU-XVI/2018 ditarik kembali dan
p03        I1 | Pemohon tidak dapat mengajukan kembali permohonan a quo;
p03        I1 [ITEM]
p03        I1 | 3. Memerintahkan kepada Panitera Mahkamah Konstitusi untuk menerbitkan Akta
p03        I1 | Pembatalan Registrasi Permohonan dan mengembalikan berkas permohonan
p03        I1 | kepada Pemohon;
p03        | Demikian diputus dalam Rapat Permusyawaratan Hakim oleh sembilan
p03        I1 | Hakim Konstitusi, yaitu Anwar Usman, selaku Ketua merangkap Anggota,
p03        I1 | Aswanto, Arief Hidayat, Suhartoyo, Enny Nurbaningsih, I Dewa Gede Palguna,
p03        I1 | Manahan M.P. Sitompul, Saldi Isra, dan Wahiduddin Adams, masing-masing
p03 B      I1 | sebagai Anggota, pada hari Kamis, tanggal sepuluh, bulan Januari, tahun dua
p03 B      I1 | ribu sembilan belas, yang diucapkan dalam Sidang Pleno Mahkamah Konstitusi
p03 B      I1 | terbuka untuk umum pada hari Kamis, tanggal dua puluh empat, bulan Januari,
p03 B      I1 | tahun dua ribu sembilan belas, selesai diucapkan pukul 11.20 WIB, oleh
p03        I1 | sembilan Hakim Konstitusi, yaitu Anwar Usman selaku Ketua merangkap Anggota,
p03        I1 | Aswanto, Arief Hidayat, Suhartoyo, Enny Nurbaningsih, I Dewa Gede Palguna,
p03        I1 | Manahan M.P. Sitompul, Saldi Isra, dan Wahiduddin Adams, masing-masing
p03        I1 | sebagai Anggota, dengan dibantu oleh Hani Adhani sebagai Panitera Pengganti,
p03        I1 | serta dihadiri oleh Presiden atau yang mewakili dan Dewan Perwakilan Rakyat
p03        I1 | atau yang mewakili serta tanpa dihadiri oleh Pemohon/kuasanya.
p03 B      | KETUA,
p03 B      | ttd.
p03 B      | Anwar Usman
==================== PAGE 4 ====================
p04        | 4
p04 B      I4 | ANGGOTA-ANGGOTA,
p04 B      | ttd.
p04 B      I3 | Aswanto
p04 B      I7 | ttd.
p04 B      I6 | Arief Hidayat
p04 B      | ttd.
p04 B      I3 | Suhartoyo
p04 B      I7 | ttd.
p04 B      I5 | Enny Nurbaningsih
p04 B      | ttd.
p04 B      I3 | I Dewa Gede Palguna
p04 B      I7 | ttd.
p04 B      I5 | Manahan M.P. Sitompul
p04 B      | ttd.
p04 B      I3 | Saldi Isra
p04 B      I7 | ttd.
p04 B      I5 | Wahiduddin Adams
p04 B      I4 | PANITERA PENGGANTI,
p04 B      | ttd.
p04 B      | Hani Adhani
```

---


## JDIH_Kemnaker

- **File**: `JDIH_Kemnaker/Permenaker No. 90 Tahun 2013.pdf`
- **Document Type**: Peraturan Menteri
- **Issued by**: Menteri Ketenagakerjaan
- **Pages**: 5 | **Lines**: 151
- **Font sizes**: [12.0]
- **Most common font**: 12.0 (100% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [85.0, 103.0, 139.0, 204.0, 242.0, 269.0, 355.0, 408.0, 453.0, 497.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01        | PERATURAN PEMERINTAH REPUBLIK INDONESIA
p01        I5 | NOMOR 90 TAHUN 2013
p01        I6 | TENTANG
p01        I2 | PENCABUTAN PERATURAN PEMERINTAH NOMOR 28 TAHUN 2003
p01        I3 | TENTANG SUBSIDI DAN IURAN PEMERINTAH DALAM
p01        I3 | PENYELENGGARAAN ASURANSI KESEHATAN BAGI
p01        | PEGAWAI NEGERI SIPIL DAN PENERIMA PENSIUN
p01        I4 [PREAMBLE:DENGAN RAHMAT]
p01        I4 | DENGAN RAHMAT TUHAN YANG MAHA ESA
p01        I4 | PRESIDEN REPUBLIK INDONESIA,
p01        I1 [PREAMBLE:MENIMBANG]
p01        I1 | Menimbang :   a.   bahwa   ketentuan   mengenai   kewajiban   pemerintah
p01        I4 | sebagai pemberi kerja bagi pegawai negeri sipil dan
p01        I4 | penerima pensiun untuk membayar dan menyetor iuran
p01        I4 | program jaminan kesehatan yang menjadi tanggung
p01        I4 | jawabnya, telah diatur dalam Undang-Undang Nomor 24
p01        I4 | Tahun  2011  tentang  Badan  Peneyelenggara  Jaminan
p01        I4 | Sosial dan peraturan pelaksananya;
p01        I4 [SUB-ITEM]
p01        I4 | b.
p01        I4 | bahwa   dalam   rangka   mensinkronisasi   pengaturan
p01        I4 | mengenai  penyelenggaraan   jaminan   kesehatan   bagi
p01        I4 | pegawai negeri  sipil  dan  penerima  pensiun  sesuai
p01        I4 | dengan Undang-Undang Nomor 24 Tahun 2011 tentang
p01        I4 | Badan Peneyelenggara Jaminan Sosial, perlu mencabut
p01        I4 | Peraturan Pemerintah Nomor 28 Tahun 2003 tentang
p01        I4 | Subsidi dan Iuran Pemerintah dalam Penyelenggaraan
p01        I4 | Asuransi  Kesehatan  Bagi  Pegawai  Negeri  Sipil  dan
p01        I4 | Penerima Pensiun;
p01        I4 [SUB-ITEM]
p01        I4 | c.
p01        I4 | bahwa     berdasarkan     pertimbangan     sebagaimana
p01        I4 | dimaksud  dalam   huruf   a   dan   huruf   b,   perlu
p01        I4 [PREAMBLE:MENETAPKAN]
p01        I4 | menetapkan Peraturan Pemerintah tentang Pencabutan
p01        I4 | Peraturan Pemerintah Nomor 28 Tahun 2003 tentang
p01        I4 | Subsidi dan Iuran Pemerintah dalam Penyelenggaraan
p01        I4 | Asuransi Kesehatan  Bagi  Pegawai  Negeri  Sipil  dan
p01        I4 | Penerima Pensiun;
p01        I1 [PREAMBLE:MENGINGAT]
p01        I1 | Mengingat     :   1.  Pasal 5 ayat (2) Undang-Undang Dasar Negara Republik
p01        I4 | Indonesia Tahun 1945;
p01        I8 [ITEM]
p01        I8 | 2. Undang-Undang . . .
==================== PAGE 2 ====================
p02        I6 | - 2 -
p02        I4 [ITEM]
p02        I4 | 2. Undang-Undang Nomor 8  Tahun 1974  tentang Pokok-
p02        I4 | Pokok  Kepegawaian    (Lembaran    Negara    Republik
p02        I4 | Indonesia Tahun 1974 Nomor 55, Tambahan Lembaran
p02        I4 | Negara Republik Indonesia Nomor 3041) sebagaimana
p02        I4 | telah
p02        I5 | diubah    dengan    Undang-Undang    Nomor    43
p02        I4 | Tahun   1999   (Lembaran   Negara   Republik   Indonesia
p02        I4 | Tahun 1999   Nomor 169, Tambahan Lembaran Negara
p02        I4 | Republik Indonesia Nomor 3890);
p02        I4 [ITEM]
p02        I4 | 3. Undang-Undang Nomor 40 Tahun 2004 tentang Sistem
p02        I4 | Jaminan Sosial Nasional (Lembaran Negara Republik
p02        I4 | Indonesia Tahun 2004 Nomor 150, Tambahan Lembaran
p02        I4 | Negara Republik Indonesia Nomor 4456);
p02        I4 [ITEM]
p02        I4 | 4. Undang-Undang    Nomor    36    Tahun    2009    tentang
p02        I4 | Kesehatan
p02        I6 | (Lembaran    Negara    Republik    Indonesia
p02        I4 | Tahun 2009  Nomor 144,  Tambahan Lembaran Negara
p02        I4 | Republik Indonesia Nomor 5063);
p02        I4 [ITEM]
p02        I4 | 5.  Undang-Undang Nomor 24 Tahun 2011 tentang Badan
p02        I4 | Penyelenggara
p02        | Jaminan    Sosial    (Lembaran    Negara
p02        I4 | Republik Indonesia Tahun 2011 Nomor 116, Tambahan
p02        I4 | Lembaran Negara Republik Indonesia Nomor 5256);
p02        I6 [KEPUTUSAN:MEMUTUSKAN]
p02        I6 | MEMUTUSKAN:
p02        I1 [PREAMBLE:MENETAPKAN]
p02        I1 | Menetapkan :  PERATURAN    PEMERINTAH    TENTANG         PENCABUTAN
p02        | PERATURAN    PEMERINTAH    NOMOR    28    TAHUN    2003
p02        | TENTANG
p02        I5 | SUBSIDI
p02        | DAN
p02        I7 | IURAN
p02        I8 | PEMERINTAH
p02        I10 | DALAM
p02        | PENYELENGGARAAN ASURANSI KESEHATAN BAGI PEGAWAI
p02        | NEGERI SIPIL DAN PENERIMA PENSIUN.
p02        | Pasal 1
p02        | Peraturan Pemerintah Nomor 28 Tahun 2003 tentang Subsidi
p02        | dan Iuran   Pemerintah   dalam   Penyelenggaraan   Asuransi
p02        | Kesehatan Bagi Pegawai Negeri Sipil dan Penerima Pensiun
p02        | (Lembaran Negara Republik Indonesia Tahun 2003 Nomor 62,
p02        | Tambahan      Lembaran      Negara      Republik      Indonesia
p02        | Nomor 4294), dicabut dan dinyatakan tidak berlaku.
p02        | Pasal 2
p02        | Peraturan   Pemerintah   ini   mulai   berlaku   pada   tanggal
p02        | 1 Januari 2014.
p02        I10 | Agar . . .
==================== PAGE 3 ====================
p03        I6 | - 3 -
p03        | Agar     setiap
p03        I6 | orang     mengetahuinya,
p03        I9 | memerintahkan
p03        | pengundangan
p03        I6 | Peraturan      Pemerintah
p03        I9 | ini      dengan
p03        | penempatannya dalam Lembaran Negara Republik Indonesia.
p03        I6 | Ditetapkan di Jakarta
p03        I6 | pada tanggal 24 Desember 2013
p03        I6 | PRESIDEN REPUBLIK INDONESIA,
p03        I7 | ttd.
p03        I6 | DR. H. SUSILO BAMBANG YUDHOYONO
p03        I1 | Diundangkan di Jakarta
p03        I1 | pada tanggal 24 Desember 2013
p03        I1 | MENTERI HUKUM DAN HAK ASASI MANUSIA
p03        I3 | REPUBLIK INDONESIA,
p03        I4 | ttd.
p03        I3 | AMIR SYAMSUDIN
p03        I1 | LEMBARAN NEGARA REPUBLIK INDONESIA TAHUN 2013 NOMOR 242
==================== PAGE 4 ====================
p04        I6 | PENJELASAN
p04        | ATAS
p04        | PERATURAN PEMERINTAH REPUBLIK INDONESIA
p04        I5 | NOMOR 90 TAHUN 2013
p04        I6 | TENTANG
p04        I2 | PENCABUTAN PERATURAN PEMERINTAH NOMOR 28 TAHUN 2003
p04        I3 | TENTANG SUBSIDI DAN IURAN PEMERINTAH DALAM
p04        I3 | PENYELENGGARAAN ASURANSI KESEHATAN BAGI
p04        I3 | PEGAWAI NEGERI SIPIL DAN PENERIMA PENSIUN
p04        I1 | I.  UMUM
p04        I3 | Penyelenggaraan jaminan kesehatan secara nasional diamanatkan
p04        I1 | dalam Pasal 19 Undang-Undang Nomor 40 Tahun 2004 tentang Sistem
p04        I1 | Jaminan Sosial Nasional. Jaminan kesehatan merupakan salah satu jenis
p04        I1 | program jaminan sosial. Dalam rangka membentuk Sistem Jaminan Sosial
p04        I1 | Nasional perlu pengaturan yang terpadu dalam penyelenggaraan jaminan
p04        I1 | kesehatan.
p04        I3 | Pasal 19 Undang-Undang Nomor 24 Tahun 2011 tentang Badan
p04        I1 | Peneyelenggara Jaminan Sosial mengatur bahwa pemberi kerja wajib
p04        I1 | membayar dan menyetor iuran jaminan kesehatan yang menjadi tanggung
p04        I1 | jawabnya kepada Badan Peneyelenggara Jaminan Sosial. Pemerintah yang
p04        I1 | merupakan pemberi kerja bagi pegawai negeri sipil, membayar iuran yang
p04        I1 | menjadi tanggung jawabnya tersebut. Adapun besaran dan tata cara
p04        I1 | pembayaran iuran program jaminan kesehatan sebagaimana dimaksud
p04        I1 | akan diatur lebih lanjut dalam Peraturan Presiden.
p04        I3 | Selain membayar iuran program jaminan kesehatan pegawai negeri
p04        I1 | sipil yang  menjadi  tanggungannya  sebagaimana dimaksud, pemerintah
p04        I1 | turut pula membayar iuran program jaminan kesehatan bagi penerima
p04        I1 | pensiun  yang  meliputi  pegawai  negeri  sipil  yang  berhenti dengan  hak
p04        I1 | pensiun, anggota TNI/Polri yang berhenti dengan hak pensiun, pejabat
p04        I1 | negara yang berhenti dengan hak pensiun, dan janda, duda, anak yatim
p04        I1 | piatu dari penerima pensiun pegawai negeri sipil, anggota TNI/Polri, dan
p04        I1 | pejabat negara yang mendapat hak pensiun.
p04        I10 | Dengan . . .
==================== PAGE 5 ====================
p05        I6 | - 2 -
p05        I3 | Dengan diaturnya ketentuan mengenai subsidi dan iuran program
p05        I1 | jaminan kesehatan bagi pegawai negeri sipil dan penerima pensiun yang
p05        I1 | dibayar  pemerintah  dalam  Peraturan  Presiden, sebagaimana ketentuan
p05        I1 | Undang-Undang Nomor 24 Tahun 2011 tentang Badan Peneyelenggara
p05        I1 | Jaminan  Sosial  maka  Peraturan  Pemerintah  Nomor  28  Tahun  2003
p05        I1 | tentang Subsidi dan Iuran Pemerintah dalam Penyelenggaraan Asuransi
p05        I1 | Kesehatan Bagi Pegawai Negeri Sipil dan Penerima Pensiun perlu dicabut
p05        I1 | dan dinyatakan tidak berlaku.
p05        I1 | II.  PASAL DEMI PASAL
p05        I1 | Pasal 1
p05        I3 | Cukup jelas.
p05        I1 | Pasal 2
p05        I3 | Cukup jelas.
p05        I1 | TAMBAHAN LEMBARAN NEGARA REPUBLIK INDONESIA NOMOR 5485
```

---


## JDIH_Kemenkeu

- **File**: `JDIH_Kemenkeu/PMK_No__9_Tahun_2025_2024pmkeuangan009.pdf`
- **Document Type**: Peraturan Menteri
- **Issued by**: Menteri Keuangan
- **Pages**: 6 | **Lines**: 290
- **Font sizes**: [8.0, 12.0, 14.0]
- **Most common font**: 12.0 (97% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [71.0, 89.0, 127.0, 156.0, 198.0, 252.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01        I3 | PERATURAN MENTERI KEUANGAN REPUBLIK INDONESIA
p01        I6 | NOMOR 9 TAHUN 2025
p01        I6 | TENTANG
p01        I5 | PENGENAAN BEA MASUK ANTIDUMPING
p01        I4 | TERHADAP IMPOR PRODUK HOT ROLLED PLATE
p01        I2 | DARI REPUBLIK RAKYAT TIONGKOK, SINGAPURA, DAN UKRAINA
p01        I4 [PREAMBLE:DENGAN RAHMAT]
p01        I4 | DENGAN RAHMAT TUHAN YANG MAHA ESA
p01        I4 | MENTERI KEUANGAN REPUBLIK INDONESIA,
p01        I1 [PREAMBLE:MENIMBANG]
p01        I1 | Menimbang
p01        I4 | : a.
p01        I5 | bahwa Indonesia sebagai Negara anggota Organisasi
p01        I5 | Perdagangan
p01        | Dunia
p01        | (World
p01        | Trade
p01        | Organization)
p01        I5 | berkewajiban untuk berperan aktif dalam mewujudkan
p01        I5 | tatanan perdagangan dunia yang adil;
p01        I4 [SUB-ITEM]
p01        I4 | b.
p01        I5 | bahwa berdasarkan ketentuan Pasal 2 ayat (1) Peraturan
p01        I5 | Pemerintah Nomor 34 Tahun 2011 tentang Tindakan
p01        I5 | Antidumping,
p01        | Tindakan
p01        | Imbalan,
p01        | dan
p01        | Tindakan
p01        I5 | Pengamanan Perdagangan, terhadap barang impor selain
p01        I5 | dikenakan bea masuk juga dapat dikenakan bea masuk
p01        I5 | antidumping jika harga ekspor dari barang yang diimpor
p01        I5 | lebih rendah dari nilai normalnya dan menyebabkan
p01        I5 | kerugian;
p01        I4 [SUB-ITEM]
p01        I4 | c.
p01        I5 | bahwa hasil penyelidikan Komite Anti Dumping Indonesia
p01        I5 | telah membuktikan praktik dumping atas impor produk
p01        I5 | Hot
p01        I6 | Rolled
p01        | Plate
p01        | dari
p01        | Republik
p01        | Rakyat
p01        | Tiongkok,
p01        I5 | Singapura, dan Ukraina masih berlanjut, sehingga
p01        I5 | pengenaan bea masuk antidumping perlu dilakukan;
p01        I4 [SUB-ITEM]
p01        I4 | d.
p01        I5 | bahwa pengenaan bea masuk antidumping terhadap
p01        I5 | impor produk Hot Rolled Plate dari Republik Rakyat
p01        I5 | Tiongkok, Singapura, dan Ukraina yang telah diatur
p01        I5 | dengan
p01        I6 | Peraturan
p01        | Menteri
p01        | Keuangan
p01        | Nomor
p01        I5 | 111/PMK.010/2019 tentang Pengenaan Bea Masuk
p01        I5 | Antidumping terhadap Impor Produk Hot Rolled Plate (HRP)
p01        I5 | dari Republik Rakyat Tiongkok, Singapura, dan Ukraina,
p01        I5 | telah berakhir masa berlakunya;
==================== PAGE 2 ====================
p02        | - 2 -
p02        I4 [SUB-ITEM]
p02        I4 | e.
p02        I5 | bahwa berdasarkan pertimbangan sebagaimana dimaksud
p02        I5 | dalam huruf a sampai dengan huruf d, serta untuk
p02        I5 | melaksanakan ketentuan Pasal 23D ayat (2) Undang-
p02        I5 | Undang Nomor 17 Tahun 2006 tentang Perubahan atas
p02        I5 | Undang-Undang
p02        | Nomor
p02        | 10
p02        | Tahun
p02        | 1995
p02        | tentang
p02        I5 | Kepabeanan,
p02        | perlu
p02        [PREAMBLE:MENETAPKAN]
p02        | menetapkan
p02        | Peraturan
p02        | Menteri
p02        I5 | Keuangan tentang Pengenaan Bea Masuk Antidumping
p02        I5 | Terhadap Impor Produk Hot Rolled Plate dari Republik
p02        I5 | Rakyat Tiongkok, Singapura, dan Ukraina;
p02        I1 [PREAMBLE:MENGINGAT]
p02        I1 | Mengingat
p02        I4 | : 1.
p02        I5 | Pasal 17 ayat (3) Undang-Undang Dasar Negara Republik
p02        I5 | Indonesia Tahun 1945;
p02        I4 [ITEM]
p02        I4 | 2.
p02        I5 | Undang-Undang
p02        | Nomor
p02        | 10
p02        | Tahun
p02        | 1995
p02        | tentang
p02        I5 | Kepabeanan (Lembaran Negara Republik Indonesia Tahun
p02        I5 | 1995 Nomor 75, Tambahan Lembaran Negara Republik
p02        I5 | Indonesia Nomor 3612) sebagaimana telah diubah dengan
p02        I5 | Undang-Undang
p02        | Nomor
p02        | 17
p02        | Tahun
p02        | 2006
p02        | tentang
p02        I5 | Perubahan atas Undang-Undang Nomor 10 Tahun 1995
p02        I5 | tentang
p02        I6 | Kepabeanan
p02        | (Lembaran
p02        | Negara
p02        | Republik
p02        I5 | Indonesia Tahun 2006 Nomor 93, Tambahan Lembaran
p02   F14  I5 | Negara Republik Indonesia Nomor 4661);
p02        I4 [ITEM]
p02        I4 | 3.
p02        I5 | Undang-Undang
p02        | Nomor
p02        | 39
p02        | Tahun
p02        | 2008
p02        | tentang
p02        I5 | Kementerian
p02        | Negara
p02        | (Lembaran
p02        | Negara
p02        | Republik
p02        I5 | Indonesia Tahun 2008 Nomor 166, Tambahan Lembaran
p02        I5 | Negara Republik Indonesia Nomor 4916) sebagaimana
p02        I5 | telah diubah dengan Undang-Undang Nomor 61 Tahun
p02        I5 | 2024 tentang Perubahan atas Undang-Undang Nomor 39
p02        I5 | Tahun 2008 tentang Kementerian Negara (Lembaran
p02        I5 | Negara Republik Indonesia Tahun 2024 Nomor 225,
p02        I5 | Tambahan Lembaran Negara Republik Indonesia Nomor
p02        I5 | 6994);
p02        I4 [ITEM]
p02        I4 | 4.
p02        I5 | Peraturan Pemerintah Nomor 34 Tahun 2011 tentang
p02        I5 | Tindakan Antidumping, Tindakan Imbalan, dan Tindakan
p02        I5 | Pengamanan Perdagangan (Lembaran Negara Republik
p02        I5 | Indonesia Tahun 2011 Nomor 66, Tambahan Lembaran
p02   F14  I5 | Negara Republik Indonesia Nomor 5225);
p02        I4 [ITEM]
p02        I4 | 5.
p02        I5 | Peraturan Presiden Nomor 158 Tahun 2024 tentang
p02        I5 | Kementerian Keuangan (Lembaran Negara Republik
p02   F14  I5 | Indonesia Tahun 2024 Nomor 354);
p02        I4 [ITEM]
p02        I4 | 6.
p02        I5 | Peraturan Menteri Keuangan Nomor 124 Tahun 2024
p02        I5 | tentang Organisasi dan Tata Kerja Kementerian Keuangan
p02        I5 | (Berita Negara Republik Indonesia Tahun 2024 Nomor
p02   F14  I5 | 1063);
p02        I6 [KEPUTUSAN:MEMUTUSKAN]
p02        I6 | MEMUTUSKAN:
p02        I1 [PREAMBLE:MENETAPKAN]
p02        I1 | Menetapkan : PERATURAN MENTERI KEUANGAN TENTANG PENGENAAN
p02        I4 | BEA MASUK ANTIDUMPING TERHADAP IMPOR PRODUK HOT
p02        I4 | ROLLED
p02        I6 | PLATE
p02        | DARI
p02        | REPUBLIK
p02        | RAKYAT
p02        | TIONGKOK,
p02        I4 | SINGAPURA, DAN UKRAINA.
==================== PAGE 3 ====================
p03        | - 3 -
p03        | Pasal 1
p03        I4 | Dalam Peraturan Menteri ini yang dimaksud dengan Bea
p03        I4 | Masuk Antidumping adalah pungutan negara yang dikenakan
p03        I4 | terhadap barang dumping yang menyebabkan kerugian.
p03        | Pasal 2
p03        I4 | Terhadap impor produk Hot Rolled Plate dengan spesifikasi:
p03        I4 [ITEM]
p03        I4 | 1.
p03        I5 | produk canai lantaian dari besi atau baja bukan paduan,
p03        I5 | dengan lebar 600 mm (enam ratus milimeter) atau lebih,
p03        I5 | dicanai panas, tidak dipalut, tidak disepuh atau tidak
p03        I5 | dilapisi, tidak dalam gulungan, tidak dikerjakan lebih
p03        I5 | lanjut selain dicanai panas, dengan ketebalan melebihi 10
p03        I5 | mm (sepuluh milimeter) yang termasuk dalam pos tarif
p03        I5 [ITEM]
p03        I5 | 7208.51.00; dan
p03        I4 [ITEM]
p03        I4 | 2.
p03        I5 | produk canai lantaian dari besi atau baja bukan paduan,
p03        I5 | dengan lebar 600 mm (enam ratus milimeter) atau lebih,
p03        I5 | dicanai panas, tidak dipalut, tidak disepuh atau tidak
p03        I5 | dilapisi, tidak dalam gulungan, tidak dikerjakan lebih
p03        I5 | lanjut selain dicanai panas, dengan ketebalan 4,75 mm
p03        I5 | (empat koma tujuh puluh lima milimater) atau lebih tetapi
p03        I5 | tidak melebihi 10 mm (sepuluh milimeter) yang termasuk
p03        I5 | dalam pos tarif 7208.52.00,
p03        I4 | yang berasal dari Republik Rakyat Tiongkok, Singapura, dan
p03        I4 | Ukraina, dikenakan Bea Masuk Antidumping.
p03        | Pasal 3
p03        I4 | Bea Masuk Antidumping sebagaimana dimaksud dalam   Pasal
p03        I4 | 2 dikenakan selama 5 (lima) tahun dengan besaran tarif Bea
p03        I4 | Masuk Antidumping sebagaimana tercantum dalam Lampiran
p03        I4 | yang merupakan bagian tidak terpisahkan dari Peraturan
p03        I4 | Menteri ini.
p03        | Pasal 4
p03        I4 | Negara asal yang dikenakan Bea Masuk Antidumping
p03        I4 | sebagaimana dimaksud dalam Pasal 2 tercantum dalam
p03        I4 | Lampiran yang merupakan bagian tidak terpisahkan dari
p03        I4 | Peraturan Menteri ini.
p03        | Pasal 5
p03        I4 [AYAT]
p03        I4 | (1)
p03        I5 | Pengenaan
p03        | Bea
p03        | Masuk
p03        | Antidumping
p03        | sebagaimana
p03        I5 | dimaksud dalam Pasal 2 merupakan tambahan dari:
p03        I5 [SUB-ITEM]
p03        I5 | a.
p03        | bea masuk umum (most favoured nation); atau
p03        I5 [SUB-ITEM]
p03        I5 | b.
p03        | bea masuk preferensi berdasarkan perjanjian atau
p03        | kesepakatan internasional,
p03        I5 | yang telah dikenakan.
p03        I4 [AYAT]
p03        I4 | (2)
p03        I5 | Dalam hal ketentuan dalam perjanjian atau kesepakatan
p03        I5 | internasional tidak terpenuhi, pengenaan Bea Masuk
p03        I5 | Antidumping atas importasi dari negara yang termasuk
p03        I5 | dalam
p03        I6 | perjanjian
p03        | atau
p03        | kesepakatan
p03        | internasional
p03        I5 | sebagaimana dimaksud pada ayat (1) huruf b merupakan
p03        I5 | tambahan dari bea masuk umum (most favoured nation).
==================== PAGE 4 ====================
p04        | - 4 -
p04        | Pasal 6
p04        I4 [AYAT]
p04        I4 | (1)
p04        I5 | Besaran Bea Masuk Antidumping sebagaimana dimaksud
p04        I5 | dalam Pasal 3 berlaku terhadap barang impor Hot Rolled
p04   F14  I5 | Plate yang:
p04        I5 [SUB-ITEM]
p04        I5 | a.
p04        | dokumen pemberitahuan pabean impornya telah
p04        | mendapat nomor pendaftaran dari kantor pabean
p04        | tempat penyelesaian kewajiban pabean, dalam hal
p04        | penyelesaian kewajiban pabean dilakukan dengan
p04   F14  | pengajuan pemberitahuan pabean; atau
p04        I5 [SUB-ITEM]
p04        I5 | b.
p04        | tarif dan nilai pabeannya ditetapkan oleh kantor
p04        | pabean tempat penyelesaian kewajiban pabean,
p04        | dalam hal penyelesaian kewajiban pabean dilakukan
p04   F14  | tanpa pengajuan pemberitahuan pabean.
p04        I4 [AYAT]
p04        I4 | (2)
p04        I5 | Pemasukan dan/atau pengeluaran barang ke dan dari
p04        I5 | kawasan perdagangan bebas dan pelabuhan bebas,
p04        I5 | tempat penimbunan berikat, atau kawasan ekonomi
p04        I5 | khusus, dilaksanakan sesuai dengan ketentuan peraturan
p04        I5 | perundang-undangan mengenai pemasukan dan/atau
p04        I5 | pengeluaran barang ke dan dari kawasan perdagangan
p04        I5 | bebas dan pelabuhan bebas, tempat penimbunan berikat,
p04        I5 | atau kawasan ekonomi khusus.
p04        | Pasal 7
p04        I4 | Peraturan Menteri ini berlaku selama 5 (lima) tahun terhitung
p04        I4 | sejak tanggal berlakunya Peraturan Menteri ini.
p04        | Pasal 8
p04        I4 | Peraturan Menteri ini mulai berlaku setelah 10 (sepuluh) hari
p04        I4 | kerja terhitung sejak tanggal diundangkan.
==================== PAGE 5 ====================
p05        | - 5 -
p05   F8   | Ditandatangani secara elektronik
p05        I4 | Agar
p05        I5 | setiap
p05        I6 | orang
p05        | mengetahuinya,
p05        | memerintahkan
p05        I4 | pengundangan Peraturan Menteri ini dengan penempatannya
p05        I4 | dalam Berita Negara Republik Indonesia.
p05        I6 | Ditetapkan di Jakarta
p05        I6 | pada tanggal 24 Januari 2025
p05        I6 | MENTERI KEUANGAN REPUBLIK INDONESIA,
p05        | SRI MULYANI INDRAWATI
p05        I1 | Diundangkan di Jakarta
p05        I1 | pada tanggal                  Д
p05        I1 | DIREKTUR JENDERAL
p05        I1 | PERATURAN PERUNDANG-UNDANGAN
p05        I1 | KEMENTERIAN HUKUM REPUBLIK INDONESIA,
p05        I2 | Ѽ
p05        I1 | DHAHANA PUTRA
p05        I1 | BERITA NEGARA REPUBLIK INDONESIA TAHUN 2025 NOMOR       Ж
==================== PAGE 6 ====================
p06        | - 6 -
p06        I4 | LAMPIRAN
p06        I4 | PERATURAN MENTERI KEUANGAN REPUBLIK INDONESIA
p06        I4 | NOMOR 9 TAHUN 2025
p06        I4 | TENTANG
p06        I4 | PENGENAAN BEA MASUK ANTIDUMPING TERHADAP IMPOR
p06        I4 | PRODUK HOT ROLLED PLATE DARI REPUBLIK RAKYAT
p06        I4 | TIONGKOK, SINGAPURA, DAN UKRAINA
p06        I3 | NEGARA ASAL DAN BESARAN BEA MASUK ANTIDUMPING
p06        I1 | No.
p06        I5 | Negara
p06        | Besaran Bea Masuk
p06        | Antidumping dalam Persentase
p06        | (%)
p06        I1 [ITEM]
p06        I1 | 1.
p06        I3 | Republik Rakyat Tiongkok
p06        | 10,47
p06        I1 [ITEM]
p06        I1 | 2.
p06        I3 | Singapura
p06        | 12,50
p06        I1 [ITEM]
p06        I1 | 3.
p06        I3 | Ukraina
p06        | 12,33
p06        I6 | MENTERI KEUANGAN REPUBLIK INDONESIA,
p06        | ttd.
p06        | SRI MULYANI INDRAWATI
```

---


## JDIH_Kemendag

- **File**: `JDIH_Kemendag/Kepmendag_No__123_Tahun_2025_download_3142_2.pdf`
- **Document Type**: Keputusan Menteri (Decision)
- **Issued by**: Menteri Perdagangan
- **Pages**: 3 | **Lines**: 195
- **Font sizes**: [11.4, 11.5, 11.8, 11.9, 12.0, 12.8, 13.1, 13.2, 13.5, 13.8, 14.2, 14.4, 15.4, 16.5, 17.1, 18.1, 19.1, 19.3, 19.5, 19.8, 20.0, 20.2, 21.5, 22.6, 23.0, 24.5, 27.5]
- **Most common font**: 12.0 (79% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [63.0, 148.0, 187.0, 249.0, 272.0, 328.0, 422.0, 441.0, 470.0]
- **Expected hierarchy**: Consideranda > MENETAPKAN > Items + Tables

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01        | KEPUTUSAN MENTERI PERDAGANGAN REPUBLIK INDONESIA
p01        | NOMOR 123 TAHUN 2025
p01        I4 | TENTANG
p01        I1 | HARGA REFERENSI CRUDE PALM OIL YANG DIKENAKAN BEA KELUAR DAN
p01        | TARIF LAYANAN BADAN LAYANAN UMUM BADAN PENGELOLA DANA
p01        I3 | PERKEBUNAN KELAPA SAWIT
p01        I2 | MENTERI PERDAGANGAN REPUBLIK INDONESIA,
p01        I1 [PREAMBLE:MENIMBANG]
p01        I1 | Menimbang
p01        I2 | : a.  bahwa untuk melaksanakan ketentuan Pasal 3 ayat (1)
p01        I3 | Peraturan Menteri Perdagangan Nomor 46 Tahun 2022
p01        I3 | tentang Tata Cara Penetapan Harga Patokan Ekspor atas
p01        I3 | Produk Pertanian dan Kehutanan yang Dikenakan Bea
p01        I3 | Keluar, Harga Referensi atas Produk Pertanian dan
p01        I3 | Kehutanan dan Daftar Merek Refined, Bleached and
p01        I3 | Deodorized Palm Olein yang Dikenakan Bea Keluar dan
p01        I3 | Tarif Layanan Badan Layanan Umum Badan Pengelola
p01        I3 | Dana Perkebunan Kelapa Sawit, perlu menetapkan harga
p01        I3 | referensi crude palm oil yang dikenakan bea keluar dan
p01        I3 | tarif layanan Badan Layanan Umum Badan Pengelola
p01        I3 | Dana Perkebunan Kelapa Sawit;
p01        I2 [SUB-ITEM]
p01        I2 | b.
p01        I3 | bahwa penetapan harga referensi atas produk crude palm
p01        I3 | oil yang dikenakan bea keluar dan tarif layanan Badan
p01        I3 | Layanan Umum Badan Pengelola Dana Perkebunan Kelapa
p01        I3 | Sawit dilakukan setelah memperhatikan usulan tertulis
p01        I3 | dan hasil rapat koordinasi dengan kementerian, lembaga
p01        I3 | pemerintah non kementerian, dan/atau badan teknis
p01        I3 | terkait;
p01        I2 [SUB-ITEM]
p01        I2 | c.  bahwa berdasarkan pertimbangan sebagaimana dimaksud
p01        I3 | dalam huruf a dan huruf b, perlu menetapkan Keputusan
p01        I3 | Menteri Perdagangan tentang Harga Referensi Crude Palm
p01        I3 | Oil yang Dikenakan Bea Keluar dan Tarif Layanan Badan
p01        I3 | Layanan Umum Badan Pengelola Dana Perkebunan
p01        I3 | Kelapa Sawit;
p01        I1 [PREAMBLE:MENGINGAT]
p01        I1 | Mengingat
p01        I2 | : 1.
p01        I3 | Pasal 17 ayat (3) Undang-Undang Dasar Negara Republik
p01        I3 | Indonesia Tahun 1945;
p01        I2 [ITEM]
p01        I2 | 2.
p01        I3 | Undang-Undang
p01        I6 | Nomor
p01        | 7
p01        | Tahun
p01        I7 | 1994
p01        I9 | tentang
p01        I3 | Pengesahan Agreement Establishing The World Trade
p01        I3 | Organization
p01        I5 | (Persetujuan
p01        | Pembentukan
p01        I9 | Organisasi
p01        I3 | Perdagangan
p01        I5 | Dunia)
p01        I6 | (Lembaran
p01        I7 | Negara
p01        I9 | Republik
p01        I3 | Indonesia Tahun 1994 Nomor 57, Tambahan Lembaran
p01        I3 | Negara Republik Indonesia Nomor 3564);
p01        I2 [ITEM]
p01        I2 | 3.
p01        I3 | Undang-Undang
p01        | Nomor
p01        | 10
p01        | Tahun
p01        I7 | 1995
p01        I9 | tentang
p01        I3 | Kepabeanan (Lembaran Negara Republik Indonesia Tahun
==================== PAGE 2 ====================
p02        | -2-
p02        I3 | 1995 Nomor 75, Tambahan Lembaran Negara Republik
p02        I3 | Indonesia Nomor 3612) sebagaimana telah diubah dengan
p02        I3 | Undang-Undang
p02        | Nomor
p02        | 17
p02        | Tahun
p02        I7 | 2006
p02        I9 | tentang
p02        I3 | Perubahan atas Undang-Undang Nomor 10 Tahun 1995
p02        I3 | tentang
p02        I4 | Kepabeanan
p02        I6 | (Lembaran
p02        I7 | Negara
p02        I9 | Republik
p02        I3 | Indonesia Tahun 2006 Nomor 93, Tambahan Lembaran
p02        I3 | Negara Republik Indonesia Nomor 4661);
p02        I2 [ITEM]
p02        I2 | 4.
p02        I3 | Undang-Undang
p02        | Nomor
p02        | 39
p02        | Tahun
p02        I7 | 2008
p02        I9 | tentang
p02        I3 | Kementerian
p02        I5 | Negara
p02        I6 | (Lembaran
p02        I7 | Negara
p02        I9 | Republik
p02        I3 | Indonesia Tahun 2008 Nomor 166, Tambahan Lembaran
p02        I3 | Negara Republik Indonesia Nomor 4916) sebagaimana
p02        I3 | telah diubah dengan Undang-Undang Nomor 61 Tahun
p02        I3 | 2024 tentang Perubahan atas Undang-Undang Nomor 39
p02        I3 | Tahun 2008 tentang Kementerian Negara (Lembaran
p02        I3 | Negara Republik Indonesia Tahun 2024 Nomor 225,
p02        I3 | Tambahan Lembaran Negara Republik Indonesia Nomor
p02        I3 | 6994);
p02        I2 [ITEM]
p02        I2 | 5.
p02        I3 | Undang-Undang
p02        I6 | Nomor
p02        | 7
p02        | Tahun
p02        I7 | 2014
p02        I9 | tentang
p02        I3 | Perdagangan (Lembaran Negara Republik Indonesia Tahun
p02        I3 | 2014 Nomor 45, Tambahan Lembaran Negara Republik
p02        I3 | Indonesia Nomor 5512);
p02        I2 [ITEM]
p02        I2 | 6.
p02        I3 | Undang-Undang Nomor 6 Tahun 2023 tentang Penetapan
p02        I3 | Peraturan Pemerintah Pengganti Undang-Undang Nomor 2
p02        I3 | Tahun 2022 tentang Cipta Kerja Menjadi Undang-Undang
p02        I3 | (Lembaran Negara Republik Indonesia Tahun 2023 Nomor
p02        I3 | 41, Tambahan Lembaran Negara Republik Indonesia
p02        I3 | Nomor 6856);
p02        I2 [ITEM]
p02        I2 | 7.
p02        I3 | Peraturan Pemerintah Nomor 55 Tahun 2008 tentang
p02        I3 | Pengenaan Bea Keluar terhadap Barang Ekspor (Lembaran
p02        I3 | Negara Republik Indonesia Tahun 2008 Nomor 116,
p02        I3 | Tambahan Lembaran Negara Republik Indonesia Nomor
p02        I3 | 4886);
p02        I2 [ITEM]
p02        I2 | 8.
p02        I3 | Peraturan Pemerintah Nomor 24 Tahun 2015 tentang
p02        I3 | Penghimpunan Dana Perkebunan (Lembaran Negara
p02        I3 | Republik Indonesia Tahun 2015 Nomor 104, Tambahan
p02        I3 | Lembaran Negara Republik Indonesia Nomor 5697);
p02        I2 [ITEM]
p02        I2 | 9.
p02        I3 | Peraturan Presiden Nomor 168 Tahun 2024 tentang
p02        I3 | Kementerian Perdagangan (Lembaran Negara Republik
p02        I3 | Indonesia Tahun 2024 Nomor 364);
p02        I2 [ITEM]
p02        I2 | 10. Peraturan Menteri Perdagangan Nomor 46 Tahun 2022
p02        I3 | tentang Tata Cara Penetapan Harga Patokan Ekspor atas
p02        I3 | Produk Pertanian dan Kehutanan yang Dikenakan Bea
p02        I3 | Keluar, Harga Referensi atas Produk Pertanian dan
p02        I3 | Kehutanan dan Daftar Merek Refined, Bleached and
p02        I3 | Deodorized Palm Olein yang Dikenakan Bea Keluar dan
p02        I3 | Tarif Layanan Badan Layanan Umum Badan Pengelola
p02        I3 | Dana Perkebunan Kelapa Sawit (Berita Negara Republik
p02        I3 | Indonesia Tahun 2022 Nomor 728);
p02        I2 [ITEM]
p02        I2 | 11. Peraturan Menteri Perdagangan Nomor 23 Tahun 2023
p02        I3 | tentang Kebijakan dan Pengaturan Ekspor (Berita Negara
p02        I3 | Republik Indonesia Tahun 2023 Nomor 527) sebagaimana
p02        I3 | telah diubah beberapa kali terakhir dengan Peraturan
p02        I3 | Menteri Perdagangan Nomor 21 Tahun 2024 tentang
p02        I3 | Perubahan Kedua atas Peraturan Menteri Perdagangan
p02        I3 | Nomor 23 Tahun 2023 tentang Kebijakan dan Pengaturan
p02        I3 | Ekspor (Berita Negara Republik Indonesia Tahun 2024
==================== PAGE 3 ====================
p03   F11  I5 | -3-
p03   F13  I3 | Nomor 512);
p03   F20  I2 [ITEM]
p03   F20  I2 | 12. Peraturan Menteri Keuangan Nomor 38 Tahun 2024
p03   F16  I3 | tentang Penetapan Barang Ekspor yang Dikenakan Bea
p03   F19  I3 | Keluar dan Tarif Bea Keluar (Berita Negara Republik
p03   F14  I3 | Indonesia Tahun 2024 Nomor 294);
p03   F17  I2 [ITEM]
p03   F17  I2 | 13. Peraturan Menteri Perdagangan Nomor 26 Tahun 2024
p03   F14  I3 | tentang Ketentuan Ekspor Produk Turunan Kelapa Sawit
p03   F18  I3 | (Berita Negara Republik Indonesia Tahun 2024 Nomor
p03   F12  I3 | 674);
p03        I2 [ITEM]
p03        I2 | 14. Peraturan
p03        I5 | Menteri
p03        | Perdagangan
p03   F12  I9 | Nomor
p03   F22  I3 | 6 Tahun 2025 tentang Organisasi dan Tata Kerja
p03   F24  I3 | Kementerian Perdagangan (Berita Negara Republik
p03   F14  I3 | Indonesia Tahun 2025 Nomor 53);
p03   F12  I4 [KEPUTUSAN:MEMUTUSKAN]
p03   F12  I4 | MEMUTUSKAN:
p03   F12  I1 [PREAMBLE:MENETAPKAN]
p03   F12  I1 | Menetapkan
p03   F20  I2 | : KEPUTUSAN MENTERI PERDAGANGAN TENTANG HARGA
p03   F23  I2 | REFERENSI CRUDE PALM OIL YANG DIKENAKAN BEA
p03   F20  I2 | KELUAR DAN TARIF LAYANAN BADAN LAYANAN UMUM
p03   F14  I2 | BADAN PENGELOLA DANA PERKEBUNAN KELAPA SAWIT.
p03   F12  I1 | KESATU
p03   F20  I2 | : Menetapkan Harga Referensi Crude Palm Oil yang dikenakan
p03   F15  I2 | Bea Keluar dan Tarif Layanan Badan Layanan Umum Badan
p03   F28  I2 | Pengelola Dana Perkebunan Kelapa Sawit sebesar
p03   F14  I2 | US$ 955,44/MT.
p03   F12  I1 | KEDUA
p03   F23  I2 | : Harga Referensi sebagaimana dimaksud dalam Diktum
p03   F20  I2 | KESATU berlaku terhitung sejak tanggal 1 Februari 2025
p03   F14  I2 | sampai dengan tanggal 28 Februari 2025.
p03   F12  I1 | KETIGA
p03   F19  I2 | : Keputusan Menteri ini mulai berlaku pada tanggal 1 Februari
p03   F12  I2 [ITEM]
p03   F12  I2 | 2025.
p03   F13  I3 | Ditetapkan di Jakarta
p03   F14  I3 | pada tanggal 31 Januari 2025
p03   F14  I3 [SUB-ITEM]
p03   F14  I3 | a.n. MENTERI PERDAGANGAN REPUBLIK INDONESIA
p03   F14  I4 | Pit. Direktur Jenderal Perdagangan Luar Negeri,
p03   F12  I6 | ttd.
p03   F13  I6 | ISY KARIM
p03   F13  I1 | Salinan sesuai dengan aslinya
p03        | Sekretariat Jenderal
p03   F13  | Kementerian Perdagangan
```

---


## JDIH_Komdigi

- **File**: `JDIH_Komdigi/Permenkominfo No. 5 Tahun 2024.pdf`
- **Document Type**: Peraturan Menteri
- **Issued by**: Menteri Kominfo
- **Pages**: 23 | **Lines**: 1889
- **Font sizes**: [12.0]
- **Most common font**: 12.0 (100% of lines)
- **Bold font sizes**: [12.0]
- **Indent clusters**: [71.0, 101.0, 120.0, 156.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01        I3 | PERATURAN MENTERI KOMUNIKASI DAN INFORMATIKA
p01        | REPUBLIK INDONESIA
p01        | NOMOR 5 TAHUN 2024
p01        | TENTANG
p01        I2 | PENETAPAN BALAI UJI ALAT TELEKOMUNIKASI DAN/ATAU
p01        | PERANGKAT TELEKOMUNIKASI
p01        I4 [PREAMBLE:DENGAN RAHMAT]
p01        I4 | DENGAN RAHMAT TUHAN YANG MAHA ESA
p01        I2 | MENTERI KOMUNIKASI DAN INFORMATIKA REPUBLIK INDONESIA,
p01        I1 [PREAMBLE:MENIMBANG]
p01        I1 | Menimbang  : a.
p01        | bahwa untuk menjamin pemenuhan standar teknis pada
p01        | setiap
p01        | alat
p01        | telekomunikasi
p01        | dan/atau
p01        | perangkat
p01        | telekomunikasi yang dibuat, dirakit, dimasukan, untuk
p01        | diperdagangkan, dan/atau digunakan di wilayah Negara
p01        | Kesatuan Republik Indonesia, perlu dilakukan pengujian
p01        | alat telekomunikasi dan/atau perangkat telekomunikasi
p01        | oleh laboratorium uji;
p01        I4 [SUB-ITEM]
p01        I4 | b.
p01        | bahwa berdasarkan ketentuan Pasal 38 ayat (2) Peraturan
p01        | Pemerintah
p01        | Nomor
p01        | 46
p01        | Tahun
p01        | 2021
p01        | tentang
p01        | Pos,
p01        | Telekomunikasi,
p01        | dan
p01        | Penyiaran,
p01        | laboratorium
p01        | uji
p01        | sebagaimana dimaksud dalam huruf a ditetapkan oleh
p01        | Menteri Komunikasi dan Informatika sebagai balai uji alat
p01        | telekomunikasi dan/atau perangkat telekomunikasi;
p01        I4 [SUB-ITEM]
p01        I4 | c.
p01        | bahwa ketentuan mengenai penetapan balai uji alat
p01        | telekomunikasi
p01        | dan/atau
p01        | perangkat
p01        | telekomunikasi
p01        | sebagaimana diatur dalam Peraturan Menteri Komunikasi
p01        | dan Informatika Nomor 15 Tahun 2012 tentang Petunjuk
p01        | Pelaksanaan Penetapan Balai Uji Dalam Negeri dan
p01        | Peraturan Menteri Komunikasi dan Informatika Nomor 16
p01        | Tahun 2012 tentang Petunjuk Pelaksanaan Pengakuan
p01        | Balai Uji Negara Asing sudah tidak sesuai dengan
p01        | perkembangan
p01        | kebutuhan
p01        | hukum,
p01        | sehingga
p01        | perlu
p01        | diganti;
p01        I4 [SUB-ITEM]
p01        I4 | d.
p01        | bahwa
p01        | berdasarkan
p01        | pertimbangan
p01        | sebagaimana
p01        | dimaksud dalam huruf a, huruf b, dan huruf c, perlu
p01        [PREAMBLE:MENETAPKAN]
p01        | menetapkan
p01        | Peraturan
p01        | Menteri
p01        | Komunikasi
p01        | dan
p01        | Informatika
p01        | tentang
p01        | Penetapan
p01        | Balai
p01        | Uji
p01        | Alat
p01        | Telekomunikasi dan/atau Perangkat Telekomunikasi;
==================== PAGE 2 ====================
p02        | - 2 -
p02        I1 [PREAMBLE:MENGINGAT]
p02        I1 | Mengingat
p02        I4 | : 1.  Pasal 17 ayat (3) Undang-Undang Dasar Negara Republik
p02        | Indonesia Tahun 1945;
p02        I4 [ITEM]
p02        I4 | 2.
p02        | Undang-Undang
p02        | Nomor
p02        | 36
p02        | Tahun
p02        | 1999
p02        | tentang
p02        | Telekomunikasi (Lembaran Negara Republik Indonesia
p02        | Tahun 1999 Nomor 154, Tambahan Lembaran Negara
p02        | Republik Indonesia Nomor 3881) sebagaimana telah
p02        | diubah dengan Undang-Undang Nomor 6 Tahun 2023
p02        | tentang Penetapan Peraturan Pemerintah Pengganti
p02        | Undang-Undang Nomor 2 Tahun 2022 tentang Cipta Kerja
p02        | Menjadi Undang-Undang (Lembaran Negara Republik
p02        | Indonesia Tahun 2023 Nomor 41, Tambahan Lembaran
p02        | Negara Republik Indonesia Nomor 6856);
p02        I4 [ITEM]
p02        I4 | 3.
p02        | Undang-Undang
p02        | Nomor
p02        | 39
p02        | Tahun
p02        | 2008
p02        | tentang
p02        | Kementerian
p02        | Negara
p02        | (Lembaran
p02        | Negara
p02        | Republik
p02        | Indonesia Tahun 2008 Nomor 216, Tambahan Lembaran
p02        | Negara Republik Indonesia Nomor 5584);
p02        I4 [ITEM]
p02        I4 | 4.
p02        | Undang-Undang
p02        | Nomor
p02        | 20
p02        | Tahun
p02        | 20l4
p02        | tentang
p02        | Standardisasi dan Penilaian Kesesuaian (Lembaran
p02        | Negara Republik Indonesia Tahun 20l4 Nomor 216,
p02        | Tambahan Lembaran Negara Republik Indonesia Nomor
p02        | 5584);
p02        I4 [ITEM]
p02        I4 | 5.
p02        | Peraturan Pemerintah Nomor 46 Tahun 2021 tentang Pos,
p02        | Telekomunikasi,
p02        | dan
p02        | Penyiaran
p02        | (Lembaran
p02        | Negara
p02        | Republik Indonesia Tahun 2021 Nomor 56, Tambahan
p02        | Lembaran Negara Republik Indonesia Nomor 6658);
p02        I4 [ITEM]
p02        I4 | 6.
p02        | Peraturan Presiden Nomor 22 Tahun 2023 tentang
p02        | Kementerian Komunikasi dan Informatika (Lembaran
p02        | Negara Republik Indonesia Tahun 2023 Nomor 51);
p02        I4 [ITEM]
p02        I4 | 7.
p02        | Peraturan Menteri Komunikasi dan Informatika Nomor 12
p02        | Tahun
p02        | 2021
p02        | tentang
p02        | Organisasi
p02        | dan
p02        | Tata
p02        | Kerja
p02        | Kementerian Komunikasi dan Informatika (Berita Negara
p02        | Republik Indonesia Tahun 2021 Nomor 1120);
p02        [KEPUTUSAN:MEMUTUSKAN]
p02        | MEMUTUSKAN:
p02        I1 [PREAMBLE:MENETAPKAN]
p02        I1 | Menetapkan : PERATURAN MENTERI KOMUNIKASI DAN INFORMATIKA
p02        I4 | TENTANG PENETAPAN BALAI UJI ALAT TELEKOMUNIKASI
p02        I4 | DAN/ATAU PERANGKAT TELEKOMUNIKASI.
p02        [HEADING:BAB]
p02        | BAB I
p02        | KETENTUAN UMUM
p02        | Pasal 1
p02        I4 | Dalam Peraturan Menteri ini yang dimaksud dengan:
p02        I4 [ITEM]
p02        I4 | 1.
p02        | Alat Telekomunikasi adalah setiap alat perlengkapan yang
p02        | digunakan dalam bertelekomunikasi.
p02        I4 [ITEM]
p02        I4 | 2.
p02        | Perangkat
p02        | Telekomunikasi
p02        | adalah
p02        | sekelompok
p02        | Alat
p02        | Telekomunikasi yang memungkinkan bertelekomunikasi.
p02        I4 [ITEM]
p02        I4 | 3.
p02        | Balai Uji Alat Telekomunikasi dan/atau Perangkat
p02        | Telekomunikasi yang selanjutnya disebut Balai Uji adalah
p02        | laboratorium uji yang ditetapkan oleh Menteri untuk
p02        | melaksanakan fungsi pengujian Alat Telekomunikasi
p02        | dan/atau
p02        | Perangkat
p02        | Telekomunikasi
p02        | dalam
p02        | rangka
p02        | sertifikasi Alat Telekomunikasi dan/atau Perangkat
p02        | Telekomunikasi.
==================== PAGE 3 ====================
p03        | - 3 -
p03        I4 [ITEM]
p03        I4 | 4.
p03        | Balai
p03        | Uji
p03        | Dalam
p03        | Negeri
p03        | adalah
p03        | Balai
p03        | Uji
p03        | yang
p03        | berkedudukan di wilayah Negara Kesatuan Republik
p03        | Indonesia.
p03        I4 [ITEM]
p03        I4 | 5.
p03        | Balai Uji Luar Negeri adalah Balai Uji yang berkedudukan
p03        | di luar wilayah Negara Kesatuan Republik Indonesia.
p03        I4 [ITEM]
p03        I4 | 6.
p03        | Perjanjian
p03        | Saling
p03        | Pengakuan
p03        | (Mutual
p03        | Recognition
p03        | Agreement)
p03        | yang selanjutnya disebut
p03        | MRA adalah
p03        | pengaturan atau perjanjian yang memuat kesepakatan
p03        | antara Negara Republik Indonesia dengan negara lain
p03        | untuk saling mengakui laboratorium uji dan saling
p03        | keberterimaan Laporan Hasil Uji antar negara MRA
p03        | berdasarkan standar teknis yang berlaku di negara
p03        | tujuan.
p03        I4 [ITEM]
p03        I4 | 7.
p03        | Sertifikasi Alat Telekomunikasi dan/atau Perangkat
p03        | Telekomunikasi yang selanjutnya disebut Sertifikasi
p03        | adalah rangkaian kegiatan penerbitan sertifikat Alat
p03        | Telekomunikasi dan/atau Perangkat Telekomunikasi.
p03        I4 [ITEM]
p03        I4 | 8.
p03        | Standar
p03        | Teknis
p03        | adalah
p03        | persyaratan
p03        | teknis
p03        | Alat
p03        | Telekomunikasi dan/atau Perangkat Telekomunikasi
p03        | yang mencakup aspek elektris, elektronis, keselamatan,
p03        | kesehatan, keamanan, dan/atau lingkungan.
p03        I4 [ITEM]
p03        I4 | 9.
p03        | Laporan Hasil Uji adalah
p03        | laporan hasil uji
p03        | Alat
p03        | Telekomunikasi dan/atau Perangkat Telekomunikasi
p03        | yang diterbitkan oleh Balai Uji.
p03        I4 [ITEM]
p03        I4 | 10. Mitra MRA adalah negara lain yang melakukan MRA
p03        | dengan Negara Republik Indonesia.
p03        I4 [ITEM]
p03        I4 | 11. Badan Penetap Mitra MRA adalah  badan yang berwenang
p03        | untuk menetapkan laboratorium uji di dalam wilayah
p03        | hukumnya.
p03        I4 [ITEM]
p03        I4 | 12. Lembaga Akreditasi adalah lembaga yang melakukan
p03        | akreditasi terhadap laboratorium uji di dalam wilayah
p03        | hukumnya.
p03        I4 [ITEM]
p03        I4 | 13. Komite Akreditasi Nasional yang selanjutnya disingkat
p03        | KAN adalah lembaga nonstruktural yang bertugas dan
p03        | bertanggung jawab di bidang akreditasi lembaga penilaian
p03        | kesesuaian.
p03        I4 [ITEM]
p03        I4 | 14. Menteri adalah menteri yang menyelenggarakan urusan
p03        | pemerintahan di bidang komunikasi dan informatika.
p03        I4 [ITEM]
p03        I4 | 15. Direktur Jenderal adalah  Direktur Jenderal Sumber Daya
p03        | dan Perangkat Pos dan Informatika.
p03        I4 [ITEM]
p03        I4 | 16. Kementerian adalah kementerian yang menyelenggarakan
p03        | urusan
p03        | pemerintahan
p03        | di
p03        | bidang
p03        | komunikasi
p03        | dan
p03        | informatika.
p03        I4 [ITEM]
p03        I4 | 17. Direktorat Jenderal adalah Direktorat Jenderal  Sumber
p03        | Daya dan Perangkat Pos dan Informatika.
p03        I4 [ITEM]
p03        I4 | 18. Hari adalah hari kerja sesuai dengan yang ditetapkan oleh
p03        | Pemerintah Pusat.
p03        | Pasal 2
p03        I4 [AYAT]
p03        I4 | (1)
p03        | Setiap
p03        | Alat
p03        | Telekomunikasi
p03        | dan/atau
p03        | Perangkat
p03        | Telekomunikasi yang dibuat, dirakit, atau dimasukkan,
p03        | untuk diperdagangkan dan/atau digunakan di wilayah
p03        | Negara Kesatuan Republik Indonesia wajib dilakukan
p03        | pengujian untuk memastikan terpenuhinya Standar
p03        | Teknis sesuai dengan ketentuan peraturan perundang-
==================== PAGE 4 ====================
p04        | - 4 -
p04        | undangan.
p04        I4 [AYAT]
p04        I4 | (2)
p04        | Pengujian
p04        | sebagaimana
p04        | dimaksud
p04        | pada
p04        | ayat
p04        [AYAT]
p04        | (1)
p04        | dilaksanakan oleh laboratorium uji yang ditetapkan
p04        | sebagai Balai Uji.
p04        | Pasal 3
p04        I4 | Balai Uji sebagaimana dimaksud dalam Pasal 2 ayat (2) terdiri
p04        I4 | atas:
p04        I4 [SUB-ITEM]
p04        I4 | a.
p04        | Balai Uji Dalam Negeri; dan
p04        I4 [SUB-ITEM]
p04        I4 | b.
p04        | Balai Uji Luar Negeri.
p04        [HEADING:BAB]
p04        | BAB II
p04        | BALAI UJI DALAM NEGERI
p04        | Pasal 4
p04        I4 [AYAT]
p04        I4 | (1)
p04        | Laboratorium
p04        | uji
p04        | yang
p04        | melakukan
p04        | pengujian
p04        | Alat
p04        | Telekomunikasi dan/atau Perangkat Telekomunikasi
p04        | untuk menjadi Balai Uji Dalam Negeri sebagaimana
p04        | dimaksud dalam Pasal 3 huruf a harus mendapatkan:
p04        [SUB-ITEM]
p04        | a.
p04        | akreditasi; dan
p04        [SUB-ITEM]
p04        | b.
p04        | penetapan sebagai Balai Uji Dalam Negeri.
p04        I4 [AYAT]
p04        I4 | (2)
p04        | Akreditasi sebagaimana dimaksud pada ayat (1) huruf a
p04        | dilaksanakan
p04        | oleh
p04        | KAN
p04        | sesuai
p04        | dengan
p04        | ketentuan
p04        | peraturan perundang-undangan.
p04        I4 [AYAT]
p04        I4 | (3)
p04        | Akreditasi
p04        | sebagaimana
p04        | dimaksud
p04        | pada
p04        | ayat
p04        [AYAT]
p04        | (2)
p04        | dibuktikan dengan sertifikat akreditasi SNI ISO/IEC
p04        [ITEM]
p04        | 17025.
p04        I4 [AYAT]
p04        I4 | (4)
p04        | Penetapan sebagai Balai Uji Dalam Negeri sebagaimana
p04        | dimaksud pada ayat (1) huruf b, dilaksanakan oleh
p04        | Kementerian berkoordinasi dengan KAN.
p04        I4 [AYAT]
p04        I4 | (5)
p04        | Koordinasi
p04        | sebagaimana
p04        | dimaksud
p04        | pada
p04        | ayat
p04        [AYAT]
p04        | (4)
p04        | dilakukan untuk memeriksa:
p04        [SUB-ITEM]
p04        | a.
p04        | ruang lingkup laboratorium uji sesuai dengan
p04        | Standar Teknis yang berlaku di Indonesia; dan
p04        [SUB-ITEM]
p04        | b.
p04        | kesiapan
p04        | laboratorium
p04        | uji
p04        | dalam
p04        | melakukan
p04        | pengujian Alat Telekomunikasi dan/atau Perangkat
p04        | Telekomunikasi, paling sedikit mencakup:
p04        [ITEM]
p04        | 1.
p04        | kompetensi penguji terhadap Standar Teknis;
p04        [ITEM]
p04        | 2.
p04        | pelaksanan pengujian berdasarkan metode uji
p04        | sesuai Standar Teknis; dan
p04        [ITEM]
p04        | 3.
p04        | sarana dan prasarana pengujian yang dimiliki
p04        | dan kesesuaiannya dengan ruang lingkup
p04        | pengujian berdasarkan kebutuhan parameter
p04        | uji sesuai Standar Teknis.
p04        I4 [AYAT]
p04        I4 | (6)
p04        | Laboratorium uji yang telah mendapatkan sertifikat
p04        | akreditasi SNI ISO/IEC 17025 dari KAN sebagaimana
p04        | dimaksud pada ayat (3) dapat mengajukan permohonan
p04        | penetapan sebagai Balai Uji Dalam Negeri kepada Menteri
p04        | dengan melampirkan dokumen persyaratan sebagai
p04        | berikut:
p04        [SUB-ITEM]
p04        | a.
p04        | surat permohonan penetapan sebagai Balai Uji
p04        | Dalam Negeri;
p04        [SUB-ITEM]
p04        | b.
p04        | akta pendirian perusahaan dan akta perubahan
p04        | terakhir, jika ada perubahan yang mencantumkan
p04        | bidang usaha jasa pengujian laboratorium atau
==================== PAGE 5 ====================
p05        | - 5 -
p05        | peraturan/penetapan
p05        | mengenai
p05        | pembentukan
p05        | laboratorium uji dari  kementerian/lembaga sesuai
p05        | dengan ketentuan peraturan perundang-undangan;
p05        [SUB-ITEM]
p05        | c.
p05        | salinan sertifikat dan ruang lingkup akreditasi SNI
p05        | ISO/IEC 17025 termutakhir yang diterbitkan oleh
p05        | KAN sesuai dengan Standar Teknis yang berlaku di
p05        | Indonesia;
p05        [SUB-ITEM]
p05        | d.
p05        | struktur organisasi dan daftar riwayat hidup personil
p05        | laboratorium uji yang sesuai dengan ketentuan SNI
p05        | ISO/IEC 17025 termutakhir;
p05        [SUB-ITEM]
p05        | e.
p05        | bukti kompetensi dari penguji untuk melakukan
p05        | pengujian Alat Telekomunikasi dan/atau Perangkat
p05        | Telekomunikasi berupa:
p05        [ITEM]
p05        | 1.
p05        | salinan ijazah pendidikan dengan bidang yang
p05        | berkesesuaian;
p05        [ITEM]
p05        | 2.
p05        | tanda bukti telah mengikuti pelatihan teknis;
p05        | dan/atau
p05        [ITEM]
p05        | 3.
p05        | bukti pengalaman telah melakukan pengujian
p05        | Alat
p05        | Telekomunikasi
p05        | dan/atau
p05        | Perangkat
p05        | Telekomunikasi;
p05        [SUB-ITEM]
p05        | f.
p05        | daftar peralatan pengujian yang memuat informasi
p05        | mengenai fungsi alat, model, manufaktur/pabrikan,
p05        | jumlah, dan masa laku kalibrasi terakhir, serta
p05        | metode pengujian Alat Telekomunikasi dan/atau
p05        | Perangkat Telekomunikasi berdasarkan Standar
p05        | Teknis;
p05        [SUB-ITEM]
p05        | g.
p05        | surat pernyataan secara mandiri (self declaration)
p05        | yang menyatakan tidak memiliki potensi terjadinya
p05        | konflik kepentingan dalam pelaksanaan operasional
p05        | laboratorium dengan Direktorat Jenderal;
p05        [SUB-ITEM]
p05        | h.
p05        | contoh salinan Laporan Hasil Uji terbaru yang
p05        | diterbitkan oleh laboratorium uji pemohon dengan
p05        | menggunakan acuan uji Standar Teknis untuk ruang
p05        | lingkup pengujian yang dimohonkan;
p05        [SUB-ITEM]
p05        | i.
p05        | dokumen mutu (quality document);
p05        [SUB-ITEM]
p05        | j.
p05        | instruksi kerja yang digunakan untuk menguji Alat
p05        | Telekomunikasi dan/atau Perangkat Telekomunikasi
p05        | terhadap Standar Teknis;
p05        [SUB-ITEM]
p05        | k.
p05        | Laporan Hasil Uji profisiensi (proficiency testing
p05        | document) untuk ruang lingkup pengujian yang
p05        | dimohonkan; dan
p05        [SUB-ITEM]
p05        | l.
p05        | laporan audit internal dan eksternal yang dilakukan
p05        | secara berkala (periodic audit report).
p05        I4 [AYAT]
p05        I4 | (7)
p05        | Dalam hal laboratorium uji tidak memiliki Laporan Hasil
p05        | Uji profisiensi sebagaimana dimaksud pada ayat (6) huruf
p05        | k karena program uji profisiensi untuk ruang lingkup
p05        | pengujian yang dimohonkan tidak tersedia, pemohon
p05        | dapat
p05        | menyampaikan
p05        | dokumen
p05        | uji
p05        | banding
p05        | antarlaboratorium (inter-laboratory comparison test) pada
p05        | ruang lingkup pengujian yang dimohonkan.
p05        | Pasal 5
p05        I4 [AYAT]
p05        I4 | (1)
p05        | Direktur Jenderal melaksanakan verifikasi terhadap
p05        | permohonan
p05        | penetapan
p05        | Balai
p05        | Uji
p05        | Dalam
p05        | Negeri
==================== PAGE 6 ====================
p06        | - 6 -
p06        | sebagaimana dimaksud dalam Pasal 4 setelah dokumen
p06        | persyaratan permohonan dinyatakan lengkap.
p06        I4 [AYAT]
p06        I4 | (2)
p06        | Verifikasi sebagaimana dimaksud pada ayat (1) dilakukan
p06        | terhadap:
p06        [SUB-ITEM]
p06        | a.
p06        | kesiapan
p06        | laboratorium
p06        | uji
p06        | berdasarkan
p06        | hasil
p06        | koordinasi sebagaimana dimaksud dalam Pasal 4
p06        | ayat (5); dan
p06        [SUB-ITEM]
p06        | b.
p06        | keabsahan
p06        | dokumen
p06        | persyaratan
p06        | permohonan
p06        | sebagaimana dimaksud dalam Pasal 4 ayat (6).
p06        | Pasal 6
p06        I4 [AYAT]
p06        I4 | (1)
p06        | Berdasarkan hasil verifikasi sebagaimana dimaksud
p06        | dalam Pasal 5, Menteri menyetujui atau menolak
p06        | permohonan penetapan Balai Uji Dalam Negeri.
p06        I4 [AYAT]
p06        I4 | (2)
p06        | Dalam hal permohonan disetujui, Menteri menerbitkan
p06        | penetapan Balai Uji Dalam Negeri.
p06        I4 [AYAT]
p06        I4 | (3)
p06        | Dalam hal permohonan ditolak, Direktur Jenderal
p06        | menyampaikan surat penolakan kepada pemohon.
p06        I4 [AYAT]
p06        I4 | (4)
p06        | Persetujuan atau penolakan permohonan penetapan Balai
p06        | Uji Dalam Negeri sebagaimana dimaksud pada ayat (1)
p06        | ditetapkan paling lama 40 (empat puluh) Hari sejak
p06        | dokumen
p06        | persyaratan
p06        | permohonan
p06        | sebagaimana
p06        | dimaksud dalam Pasal 4 ayat (6) diterima secara lengkap.
p06        | Pasal 7
p06        I4 | Penetapan Balai Uji Dalam Negeri sebagaimana dimaksud
p06        I4 | dalam Pasal 6 ayat (2) diberikan untuk masa laku 5 (lima)
p06        I4 | tahun dan dapat diperpanjang.
p06        | Pasal 8
p06        I4 [AYAT]
p06        I4 | (1)
p06        | Balai Uji Dalam Negeri dapat mengajukan permohonan
p06        | perpanjangan penetapan Balai Uji Dalam Negeri kepada
p06        | Menteri.
p06        I4 [AYAT]
p06        I4 | (2)
p06        | Permohonan perpanjangan penetapan Balai Uji Dalam
p06        | Negeri sebagaimana dimaksud pada ayat (1) disampaikan
p06        | dengan melampirkan dokumen persyaratan sebagai
p06        | berikut:
p06        [SUB-ITEM]
p06        | a.
p06        | surat permohonan perpanjangan penetapan sebagai
p06        | Balai Uji Dalam Negeri;
p06        [SUB-ITEM]
p06        | b.
p06        | akta pendirian perusahaan dan akta perubahan
p06        | terakhir
p06        | apabila
p06        | ada
p06        | perubahan,
p06        | yang
p06        | mencantumkan
p06        | bidang
p06        | usaha
p06        | jasa
p06        | pengujian
p06        | laboratorium atau peraturan/penetapan mengenai
p06        | pembentukan
p06        | laboratorium
p06        | uji
p06        | dari
p06        | kementerian/lembaga
p06        | sesuai
p06        | dengan
p06        | ketentuan
p06        | peraturan perundang-undangan;
p06        [SUB-ITEM]
p06        | c.
p06        | salinan sertifikat dan ruang lingkup akreditasi SNI
p06        | ISO/IEC 17025 termutakhir yang diterbitkan oleh
p06        | KAN sesuai dengan Standar Teknis yang berlaku di
p06        | Indonesia;
p06        [SUB-ITEM]
p06        | d.
p06        | struktur organisasi dan daftar riwayat hidup personil
p06        | laboratorium uji yang sesuai dengan ketentuan SNI
p06        | ISO/IEC 17025 termutakhir;
==================== PAGE 7 ====================
p07        | - 7 -
p07        [SUB-ITEM]
p07        | e.
p07        | bukti kompetensi dari penguji untuk melakukan
p07        | pengujian Alat Telekomunikasi dan/atau Perangkat
p07        | Telekomunikasi berupa:
p07        [ITEM]
p07        | 1.
p07        | salinan ijazah pendidikan dengan bidang yang
p07        | berkesesuaian;
p07        [ITEM]
p07        | 2.
p07        | tanda bukti telah mengikuti pelatihan teknis;
p07        | dan/atau
p07        [ITEM]
p07        | 3.
p07        | bukti pengalaman telah melakukan pengujian
p07        | Alat
p07        | Telekomunikasi
p07        | dan/atau
p07        | Perangkat
p07        | Telekomunikasi;
p07        [SUB-ITEM]
p07        | f.
p07        | daftar peralatan pengujian yang memuat informasi
p07        | mengenai fungsi alat, model, manufaktur/pabrikan,
p07        | jumlah, dan masa laku kalibrasi terakhir, serta
p07        | metode pengujian Alat Telekomunikasi dan/atau
p07        | Perangkat Telekomunikasi berdasarkan Standar
p07        | Teknis;
p07        [SUB-ITEM]
p07        | g.
p07        | surat pernyataan secara mandiri (self declaration)
p07        | yang menyatakan tidak memiliki potensi terjadinya
p07        | konflik kepentingan dalam pelaksanaan operasional
p07        | laboratorium dengan Direktorat Jenderal;
p07        [SUB-ITEM]
p07        | h.
p07        | contoh salinan Laporan Hasil Uji terbaru yang
p07        | diterbitkan oleh laboratorium uji pemohon dengan
p07        | menggunakan acuan uji Standar Teknis untuk ruang
p07        | lingkup pengujian yang dimohonkan;
p07        [SUB-ITEM]
p07        | i.
p07        | dokumen mutu (quality document);
p07        [SUB-ITEM]
p07        | j.
p07        | instruksi kerja yang digunakan untuk menguji Alat
p07        | Telekomunikasi dan/atau Perangkat Telekomunikasi
p07        | terhadap Standar Teknis;
p07        [SUB-ITEM]
p07        | k.
p07        | Laporan Hasil Uji profisiensi (proficiency testing
p07        | document) untuk ruang lingkup pengujian yang
p07        | dimohonkan; dan
p07        [SUB-ITEM]
p07        | l.
p07        | laporan audit internal dan eksternal yang dilakukan
p07        | secara berkala (periodic audit report).
p07        I4 [AYAT]
p07        I4 | (3)
p07        | Dalam hal Balai Uji Dalam Negeri tidak memiliki Laporan
p07        | Hasil Uji profisiensi sebagaimana dimaksud pada ayat (2)
p07        | huruf k karena program uji profisiensi untuk ruang
p07        | lingkup pengujian yang dimohonkan tidak tersedia,
p07        | pemohon dapat menyampaikan dokumen uji banding
p07        | antarlaboratorium (inter-laboratory comparison test) pada
p07        | ruang lingkup pengujian yang dimohonkan.
p07        I4 [AYAT]
p07        I4 | (4)
p07        | Permohonan perpanjangan penetapan Balai Uji Dalam
p07        | Negeri sebagaimana dimaksud pada ayat (1) diajukan
p07        | paling lambat 40 (empat puluh) Hari sebelum masa laku
p07        | penetapan Balai Uji Dalam Negeri sebagaimana dimaksud
p07        | dalam Pasal 7 berakhir.
p07        | Pasal 9
p07        I4 [AYAT]
p07        I4 | (1)
p07        | Direktur Jenderal melaksanakan verifikasi terhadap
p07        | permohonan perpanjangan penetapan Balai Uji Dalam
p07        | Negeri sebagaimana dimaksud dalam Pasal 8 ayat (1)
p07        | setelah dokumen persyaratan permohonan perpanjangan
p07        | dinyatakan lengkap.
p07        I4 [AYAT]
p07        I4 | (2)
p07        | Verifikasi
p07        | sebagaimana
p07        | dimaksud
p07        | pada
p07        | ayat
p07        [AYAT]
p07        | (1)
p07        | dilaksanakan dengan melibatkan instansi terkait.
p07        I4 [AYAT]
p07        I4 | (3)
p07        | Ketentuan mengenai verifikasi penetapan Balai Uji Dalam
==================== PAGE 8 ====================
p08        | - 8 -
p08        | Negeri sebagaimana dimaksud dalam Pasal 5 berlaku
p08        | mutatis
p08        | mutandis
p08        | untuk
p08        | verifikasi
p08        | permohonan
p08        | perpanjangan
p08        | penetapan
p08        | Balai
p08        | Uji
p08        | Dalam
p08        | Negeri
p08        | sebagaimana dimaksud pada ayat (1).
p08        | Pasal 10
p08        I4 [AYAT]
p08        I4 | (1)
p08        | Berdasarkan hasil verifikasi sebagaimana dimaksud
p08        | dalam Pasal 9, Menteri menyetujui atau menolak
p08        | permohonan perpanjangan penetapan Balai Uji Dalam
p08        | Negeri.
p08        I4 [AYAT]
p08        I4 | (2)
p08        | Dalam hal permohonan disetujui, Menteri menerbitkan
p08        | perpanjangan penetapan Balai Uji Dalam Negeri.
p08        I4 [AYAT]
p08        I4 | (3)
p08        | Dalam hal permohonan ditolak, Direktur Jenderal
p08        | menyampaikan surat penolakan kepada Balai Uji Dalam
p08        | Negeri untuk permohonan perpanjangan penetapan Balai
p08        | Uji Dalam Negeri yang:
p08        [SUB-ITEM]
p08        | a.
p08        | diajukan tidak sesuai batas waktu sebagaimana
p08        | dimaksud dalam Pasal 8 ayat (4); atau
p08        [SUB-ITEM]
p08        | b.
p08        | berdasarkan hasil verifikasi dinyatakan ditolak.
p08        I4 [AYAT]
p08        I4 | (4)
p08        | Persetujuan atau penolakan permohonan perpanjangan
p08        | penetapan Balai Uji Dalam Negeri sebagaimana dimaksud
p08        | pada ayat (1) ditetapkan paling lama 40 (empat puluh)
p08        | Hari
p08        | sejak
p08        | dokumen
p08        | persyaratan
p08        | permohonan
p08        | sebagaimana dimaksud dalam Pasal 8 ayat (2) diterima
p08        | secara lengkap.
p08        | Pasal 11
p08        I4 [AYAT]
p08        I4 | (1)
p08        | Balai Uji Dalam Negeri dapat mengajukan permohonan
p08        | penambahan ruang lingkup pengujian kepada Menteri.
p08        I4 [AYAT]
p08        I4 | (2)
p08        | Permohonan penambahan ruang lingkup pengujian
p08        | sebagaimana dimaksud pada ayat (1) disampaikan dengan
p08        | melampirkan dokumen persyaratan sebagai berikut:
p08        [SUB-ITEM]
p08        | a.
p08        | surat permohonan penambahan ruang lingkup
p08        | pengujian Balai Uji Dalam Negeri;
p08        [SUB-ITEM]
p08        | b.
p08        | salinan penetapan Balai Uji Dalam Negeri yang masih
p08        | berlaku;
p08        [SUB-ITEM]
p08        | c.
p08        | salinan sertifikat akreditasi SNI ISO/IEC 17025
p08        | termutakhir
p08        | dengan
p08        | lampiran
p08        | ruang
p08        | lingkup
p08        | pengujian yang akan ditambahkan sesuai dengan
p08        | Standar Teknis Alat Telekomunikasi dan/atau
p08        | Perangkat Telekomunikasi;
p08        [SUB-ITEM]
p08        | d.
p08        | struktur organisasi dan daftar riwayat hidup penguji
p08        | Balai Uji Dalam Negeri yang sesuai dengan ketentuan
p08        | SNI ISO/IEC 17025 termutakhir;
p08        [SUB-ITEM]
p08        | e.
p08        | bukti kompetensi dari penguji untuk melakukan
p08        | pengujian Alat Telekomunikasi dan/atau Perangkat
p08        | Telekomunikasi berupa:
p08        [ITEM]
p08        | 1.
p08        | salinan ijazah pendidikan dengan bidang yang
p08        | berkesesuaian;
p08        [ITEM]
p08        | 2.
p08        | tanda bukti telah mengikuti pelatihan teknis;
p08        | dan/atau
p08        [ITEM]
p08        | 3.
p08        | bukti pengalaman telah melakukan pengujian
p08        | Alat
p08        | Telekomunikasi
p08        | dan/atau
p08        | Perangkat
p08        | Telekomunikasi;
==================== PAGE 9 ====================
p09        | - 9 -
p09        [SUB-ITEM]
p09        | f.
p09        | surat pernyataan mengenai fasilitas dan metode
p09        | pengujian Alat Telekomunikasi dan/atau Perangkat
p09        | Telekomunikasi sesuai dengan Standar Teknis Alat
p09        | Telekomunikasi
p09        | dan/atau
p09        | Perangkat
p09        | Telekomunikasi;
p09        [SUB-ITEM]
p09        | g.
p09        | daftar peralatan pengujian yang memuat informasi
p09        | mengenai fungsi alat, model, manufaktur/pabrikan,
p09        | jumlah, dan masa laku kalibrasi terakhir;
p09        [SUB-ITEM]
p09        | h.
p09        | contoh salinan Laporan Hasil Uji terbaru yang
p09        | diterbitkan oleh Balai Uji Dalam Negeri pemohon
p09        | sesuai dengan Standar Teknis;
p09        [SUB-ITEM]
p09        | i.
p09        | dokumen mutu (quality document) yang terbaru;
p09        [SUB-ITEM]
p09        | j.
p09        | Laporan Hasil Uji profisiensi (proficiency testing
p09        | document) untuk ruang lingkup pengujian yang
p09        | dimohonkan; dan
p09        [SUB-ITEM]
p09        | k.
p09        | instruksi kerja yang digunakan untuk menguji Alat
p09        | Telekomunikasi dan/atau Perangkat Telekomunikasi
p09        | terhadap Standar Teknis.
p09        | untuk ruang lingkup pengujian yang akan ditambahkan.
p09        I4 [AYAT]
p09        I4 | (3)
p09        | Dalam hal Balai Uji Dalam Negeri tidak memiliki Laporan
p09        | Hasil Uji profisiensi sebagaimana dimaksud pada ayat (2)
p09        | huruf j karena program uji profisiensi untuk ruang
p09        | lingkup pengujian yang akan ditambahkan tidak tersedia,
p09        | pemohon dapat menyampaikan dokumen uji banding
p09        | antarlaboratorium (inter-laboratory comparison test) pada
p09        | ruang lingkup pengujian yang akan ditambahkan.
p09        I4 [AYAT]
p09        I4 | (4)
p09        | Persyaratan sebagaimana dimaksud pada ayat (2) huruf c
p09        | dapat dikecualikan untuk Balai Uji Dalam Negeri yang
p09        | belum mendapatkan sertifikat akreditasi SNI ISO/IEC
p09        | 17025 termutakhir dari KAN untuk ruang lingkup
p09        | pengujian yang akan ditambahkan.
p09        I4 [AYAT]
p09        I4 | (5)
p09        | Permohonan penambahan ruang lingkup pengujian
p09        | sebagaimana dimaksud pada ayat (1) diajukan dengan
p09        | batas waktu:
p09        [SUB-ITEM]
p09        | a.
p09        | paling lambat 1 (satu) tahun sebelum masa laku
p09        | penetapan Balai Uji Dalam Negeri berakhir; atau
p09        [SUB-ITEM]
p09        | b.
p09        | bersamaan
p09        | dengan
p09        | permohonan
p09        | perpanjangan
p09        | penetapan Balai Uji Dalam Negeri sebagaimana
p09        | dimaksud dalam Pasal 8 ayat (1).
p09        | Pasal 12
p09        I4 [AYAT]
p09        I4 | (1)
p09        | Direktur Jenderal melaksanakan verifikasi terhadap
p09        | permohonan penambahan ruang lingkup pengujian
p09        | sebagaimana dimaksud dalam Pasal 11 ayat (1) setelah
p09        | dokumen persyaratan permohonan penambahan ruang
p09        | lingkup pengujian dinyatakan lengkap.
p09        I4 [AYAT]
p09        I4 | (2)
p09        | Verifikasi
p09        | sebagaimana
p09        | dimaksud
p09        | pada
p09        | ayat
p09        [AYAT]
p09        | (1)
p09        | dilaksanakan dengan melibatkan instansi terkait.
p09        I4 [AYAT]
p09        I4 | (3)
p09        | Ketentuan mengenai verifikasi penetapan Balai Uji Dalam
p09        | Negeri sebagaimana dimaksud dalam Pasal 5 berlaku
p09        | mutatis
p09        | mutandis
p09        | untuk
p09        | verifikasi
p09        | permohonan
p09        | penambahan ruang lingkup pengujian sebagaimana
p09        | dimaksud pada ayat (1).
==================== PAGE 10 ====================
p10        | - 10 -
p10        | Pasal 13
p10        I4 [AYAT]
p10        I4 | (1)
p10        | Berdasarkan hasil verifikasi sebagaimana dimaksud
p10        | dalam Pasal 12, Menteri menyetujui atau menolak
p10        | permohonan penambahan ruang lingkup pengujian Balai
p10        | Uji Dalam Negeri.
p10        I4 [AYAT]
p10        I4 | (2)
p10        | Dalam hal permohonan disetujui, Menteri menerbitkan
p10        | penetapan penambahan ruang lingkup pengujian Balai
p10        | Uji Dalam Negeri.
p10        I4 [AYAT]
p10        I4 | (3)
p10        | Dalam hal permohonan ditolak, Direktur Jenderal
p10        | menyampaikan surat penolakan kepada Balai Uji Dalam
p10        | Negeri untuk permohonan penambahan ruang lingkup
p10        | pengujian yang:
p10        [SUB-ITEM]
p10        | a.
p10        | diajukan tidak sesuai batas waktu sebagaimana
p10        | dimaksud dalam Pasal 11 ayat (5); atau
p10        [SUB-ITEM]
p10        | b.
p10        | berdasarkan hasil verifikasi dinyatakan ditolak.
p10        I4 [AYAT]
p10        I4 | (4)
p10        | Persetujuan atau penolakan permohonan penambahan
p10        | ruang lingkup pengujian Balai Uji Dalam Negeri
p10        | sebagaimana dimaksud pada ayat (1) ditetapkan paling
p10        | lama 40 (empat puluh) Hari sejak dokumen persyaratan
p10        | permohonan penambahan ruang lingkup pengujian
p10        | sebagaimana dimaksud dalam Pasal 11 ayat (2) diterima
p10        | secara lengkap.
p10        I4 [AYAT]
p10        I4 | (5)
p10        | Penetapan penambahan ruang lingkup pengujian Balai
p10        | Uji Dalam Negeri sebagaimana dimaksud pada ayat (1)
p10        | tidak mengubah masa laku penetapan Balai Uji Dalam
p10        | Negeri.
p10        | Pasal 14
p10        I4 [AYAT]
p10        I4 | (1)
p10        | Balai Uji Dalam Negeri sebagaimana dimaksud dalam
p10        | Pasal 11 ayat (4) wajib menyampaikan salinan sertifikat
p10        | akreditasi SNI ISO/IEC 17025 termutakhir yang memuat
p10        | informasi ruang lingkup pengujian sesuai dengan yang
p10        | ditetapkan paling lama 2 (dua) tahun sejak penetapan
p10        | penambahan ruang lingkup pengujian sebagaimana
p10        | dimaksud dalam Pasal 13 ayat (2).
p10        I4 [AYAT]
p10        I4 | (2)
p10        | Salinan
p10        | sertifikat
p10        | akreditasi
p10        | SNI
p10        | ISO/IEC
p10        | 17025
p10        | termutakhir sebagaimana dimaksud pada ayat (1)
p10        | disampaikan kepada Direktur Jenderal.
p10        I4 [AYAT]
p10        I4 | (3)
p10        | Dalam hal sampai dengan batas waktu sebagaimana
p10        | dimaksud pada ayat (1), Balai Uji Dalam Negeri belum
p10        | menyampaikan salinan sertifikat akreditasi SNI ISO/IEC
p10        | 17025 termutakhir yang memuat informasi ruang lingkup
p10        | pengujian
p10        | sesuai
p10        | dengan
p10        | yang
p10        | telah
p10        | ditetapkan,
p10        | penetapan penambahan ruang lingkup pengujian Balai
p10        | Uji Dalam Negeri sebagaimana dimaksud dalam Pasal 13
p10        | ayat (2) dinyatakan batal dan tidak berlaku.
p10        I4 [AYAT]
p10        I4 | (4)
p10        | Direktur Jenderal menerbitkan surat pemberitahuan
p10        | pembatalan
p10        | dan
p10        | tidak
p10        | berlakunya
p10        | penetapan
p10        | penambahan ruang lingkup pengujian sebagaimana
p10        | dimaksud pada ayat (3) kepada Balai Uji Dalam Negeri.
p10        I4 [AYAT]
p10        I4 | (5)
p10        | Laporan Hasil Uji untuk ruang lingkup pengujian yang
p10        | diterbitkan
p10        | setelah
p10        | tanggal
p10        | surat
p10        | pemberitahuan
p10        | pembatalan penetapan penambahan ruang lingkup
p10        | pengujian Balai Uji Dalam Negeri sebagaimana dimaksud
p10        | pada ayat (4), menjadi tidak berlaku dan tidak dapat
==================== PAGE 11 ====================
p11        | - 11 -
p11        | digunakan untuk permohonan Sertifikasi.
p11        | Pasal 15
p11        I4 | Balai Uji Dalam Negeri sebagaimana dimaksud dalam Pasal 14
p11        I4 | ayat (3) hanya dapat mengajukan kembali permohonan
p11        I4 | penambahan ruang lingkup pengujian yang sama jika telah
p11        I4 | mendapatkan
p11        | sertifikat
p11        | akreditasi
p11        | SNI
p11        | ISO/IEC
p11        | 17025
p11        I4 | termutakhir.
p11        | Pasal 16
p11        I4 | Balai Uji Dalam Negeri wajib:
p11        I4 [SUB-ITEM]
p11        I4 | a.
p11        | melaksanakan pengujian Alat Telekomunikasi dan/atau
p11        | Perangkat Telekomunikasi sesuai dengan Standar Teknis
p11        | Alat Telekomunikasi dan/atau Perangkat Telekomunikasi
p11        | dan ruang lingkup pengujian yang ditetapkan;
p11        I4 [SUB-ITEM]
p11        I4 | b.
p11        | menggunakan tanda tangan digital pada Laporan Hasil
p11        | Uji;
p11        I4 [SUB-ITEM]
p11        I4 | c.
p11        | menyampaikan
p11        | rekapitulasi
p11        | data
p11        | pengujian
p11        | Alat
p11        | Telekomunikasi dan/atau Perangkat Telekomunikasi
p11        | kepada Menteri c.q. Direktur Jenderal setiap 3 (tiga)
p11        | bulan;
p11        I4 [SUB-ITEM]
p11        I4 | d.
p11        | memberikan klarifikasi keabsahan Laporan Hasil Uji
p11        | dalam hal diperlukan oleh Direktur Jenderal; dan
p11        I4 [SUB-ITEM]
p11        I4 | e.
p11        | melaporkan kepada Direktur Jenderal dalam hal terjadi
p11        | perubahan:
p11        [ITEM]
p11        | 1.
p11        | perizinan berusaha;
p11        [ITEM]
p11        | 2.
p11        | struktur organisasi;
p11        [ITEM]
p11        | 3.
p11        | akreditasi;
p11        [ITEM]
p11        | 4.
p11        | alamat Balai Uji Dalam Negeri; atau
p11        [ITEM]
p11        | 5.
p11        | yang
p11        | dapat
p11        | memengaruhi
p11        | kesinambungan
p11        | pengujian.
p11        [HEADING:BAB]
p11        | BAB III
p11        | BALAI UJI LUAR NEGERI
p11        | Pasal 17
p11        I4 [AYAT]
p11        I4 | (1)
p11        | Menteri dapat melakukan saling pengakuan Laporan
p11        | Hasil Uji dengan negara lain.
p11        I4 [AYAT]
p11        I4 | (2)
p11        | Laporan Hasil Uji sebagaimana dimaksud pada ayat (1)
p11        | merupakan Laporan Hasil Uji yang diterbitkan oleh Balai
p11        | Uji Luar Negeri.
p11        I4 [AYAT]
p11        I4 | (3)
p11        | Balai Uji Luar Negeri sebagaimana dimaksud pada ayat (2)
p11        | ditetapkan melalui mekanisme MRA.
p11        | Pasal 18
p11        I4 [AYAT]
p11        I4 | (1)
p11        | MRA sebagaimana dimaksud dalam Pasal  17 ayat (3)
p11        | dibuat berdasarkan:
p11        [SUB-ITEM]
p11        | a.
p11        | asas manfaat; dan
p11        [SUB-ITEM]
p11        | b.
p11        | prinsip
p11        | timbal
p11        | balik
p11        | (reciprocal)
p11        | yang
p11        | saling
p11        | menguntungkan.
p11        I4 [AYAT]
p11        I4 | (2)
p11        | MRA sebagaimana dimaksud pada ayar (1) paling sedikit
p11        | memuat:
p11        [SUB-ITEM]
p11        | a.
p11        | ruang lingkup MRA;
p11        [SUB-ITEM]
p11        | b.
p11        | badan penetap;
p11        [SUB-ITEM]
p11        | c.
p11        | prosedur dan persyaratan penetapan laboratorium
==================== PAGE 12 ====================
p12        | - 12 -
p12        | uji;
p12        [SUB-ITEM]
p12        | d.
p12        | daftar standar atau regulasi teknis yang menjadi
p12        | acuan di masing-masing negara sesuai ruang lingkup
p12        | MRA; dan
p12        [SUB-ITEM]
p12        | e.
p12        | ketentuan
p12        | mengenai
p12        | pemberlakuan
p12        | dan
p12        | pengakhiran MRA.
p12        I4 [AYAT]
p12        I4 | (3)
p12        | MRA sebagaimana dimaksud pada ayat (1) dilaksanakan
p12        | sesuai dengan ketentuan peraturan perundang-undangan
p12        | terkait perjanjian internasional.
p12        I4 [AYAT]
p12        I4 | (4)
p12        | Direktur Jenderal mengumumkan pelaksanaan MRA
p12        | melalui situs web Direktorat Jenderal.
p12        | Pasal 19
p12        I4 [AYAT]
p12        I4 | (1)
p12        | Balai Uji Dalam Negeri dapat mengajukan permohonan
p12        | penetapan laboratorium uji kepada Mitra MRA.
p12        I4 [AYAT]
p12        I4 | (2)
p12        | Permohonan sebagaimana dimaksud pada ayat (1)
p12        | disampaikan melalui Direktur Jenderal.
p12        I4 [AYAT]
p12        I4 | (3)
p12        | Permohonan penetapan sebagaimana dimaksud pada
p12        | ayat (1) mencantumkan ruang lingkup pengujian yang
p12        | dimohonkan untuk ditetapkan oleh Mitra MRA dengan
p12        | melampirkan dokumen yang dipersyaratkan dalam MRA.
p12        I4 [AYAT]
p12        I4 | (4)
p12        | Direktur
p12        | Jenderal
p12        | melakukan
p12        | verifikasi
p12        | terhadap
p12        | permohonan penetapan laboratorium uji sebagaimana
p12        | dimaksud pada ayat (1) sesuai dengan prosedur dan
p12        | persyaratan penetapan laboratorium uji yang ditetapkan
p12        | dalam MRA sebagaimana dimaksud dalam Pasal 18 ayat
p12        [AYAT]
p12        | (2) huruf c.
p12        I4 [AYAT]
p12        I4 | (5)
p12        | Berdasarkan hasil verifikasi sebagaimana dimaksud pada
p12        | ayat (4), Direktur Jenderal menyampaikan permohonan
p12        | penetapan Balai Uji Dalam Negeri kepada Mitra MRA.
p12        | Pasal 20
p12        I4 [AYAT]
p12        I4 | (1)
p12        | Mitra MRA dapat mengajukan permohonan penetapan
p12        | laboratorium uji Mitra MRA kepada Menteri.
p12        I4 [AYAT]
p12        I4 | (2)
p12        | Permohonan penetapan sebagaimana dimaksud pada ayat
p12        [AYAT]
p12        | (1) mencantumkan ruang lingkup pengujian yang
p12        | dimohonkan
p12        | untuk
p12        | ditetapkan
p12        | dan
p12        | melampirkan
p12        | dokumen yang dipersyaratkan dalam MRA.
p12        I4 [AYAT]
p12        I4 | (3)
p12        | Dokumen persyaratan permohonan penetapan Balai Uji
p12        | Luar Negeri sebagaimana dimaksud pada ayat (2) harus
p12        | menggunakan:
p12        [SUB-ITEM]
p12        | a.
p12        | bahasa Indonesia;
p12        [SUB-ITEM]
p12        | b.
p12        | bahasa Inggris; atau
p12        [SUB-ITEM]
p12        | c.
p12        | bahasa asing lainnya, yang disertai terjemahan resmi
p12        | menggunakan bahasa Indonesia dan/atau bahasa
p12        | Inggris.
p12        | Pasal 21
p12        I4 [AYAT]
p12        I4 | (1)
p12        | Direktur
p12        | Jenderal
p12        | melakukan
p12        | verifikasi
p12        | terhadap
p12        | permohonan penetapan laboratorium uji Mitra MRA
p12        | sebagaimana dimaksud dalam Pasal 20 ayat (1).
p12        I4 [AYAT]
p12        I4 | (2)
p12        | Verifikasi
p12        | sebagaimana
p12        | dimaksud
p12        | pada
p12        | ayat
p12        [AYAT]
p12        | (1)
p12        | dilaksanakan sesuai dengan prosedur dan persyaratan
p12        | penetapan laboratorium uji yang diatur dalam MRA
p12        | sebagaimana dimaksud dalam Pasal 18 ayat (2) huruf c.
==================== PAGE 13 ====================
p13        | - 13 -
p13        | Pasal 22
p13        I4 [AYAT]
p13        I4 | (1)
p13        | Berdasarkan hasil verifikasi sebagaimana dimaksud
p13        | dalam Pasal 21, Menteri menyetujui atau menolak
p13        | permohonan penetapan Balai Uji Luar Negeri.
p13        I4 [AYAT]
p13        I4 | (2)
p13        | Dalam hal permohonan disetujui, Menteri menerbitkan
p13        | penetapan Balai Uji Luar Negeri.
p13        I4 [AYAT]
p13        I4 | (3)
p13        | Dalam hal permohonan ditolak, Direktur Jenderal
p13        | menyampaikan surat penolakan kepada Mitra MRA.
p13        I4 [AYAT]
p13        I4 | (4)
p13        | Persetujuan atau penolakan permohonan penetapan Balai
p13        | Uji Luar Negeri sebagaimana dimaksud pada ayat (1)
p13        | ditetapkan paling lama 40 (empat puluh) Hari sejak
p13        | permohonan diterima secara lengkap.
p13        | Pasal 23
p13        I4 [AYAT]
p13        I4 | (1)
p13        | Balai Uji Luar Negeri dapat mengajukan permohonan
p13        | perpanjangan penetapan Balai Uji Luar Negeri atau
p13        | penambahan ruang lingkup pengujian kepada Menteri
p13        | melalui mitra MRA.
p13        I4 [AYAT]
p13        I4 | (2)
p13        | Permohonan sebagaimana dimaksud pada ayat (1)
p13        | diajukan paling lambat 40 (empat puluh) Hari sebelum
p13        | masa berlaku penetapan Balai Uji Luar Negeri berakhir.
p13        I4 [AYAT]
p13        I4 | (3)
p13        | Permohonan sebagaimana dimaksud pada ayat (1)
p13        | disampaikan dengan melampirkan dokumen persyaratan
p13        | sesuai yang ditetapkan dalam MRA.
p13        I4 [AYAT]
p13        I4 | (4)
p13        | Direktur
p13        | Jenderal
p13        | melakukan
p13        | verifikasi
p13        | terhadap
p13        | permohonan sebagaimana dimaksud pada ayat (1) sesuai
p13        | dengan prosedur dan persyaratan yang ditetapkan dalam
p13        | MRA.
p13        I4 [AYAT]
p13        I4 | (5)
p13        | Berdasarkan hasil verifikasi sebagaimana dimaksud pada
p13        | ayat (4), Menteri menyetujui atau menolak permohonan
p13        | perpanjangan penetapan Balai Uji Luar Negeri atau
p13        | penambahan ruang lingkup pengujian.
p13        | Pasal 24
p13        I4 [AYAT]
p13        I4 | (1)
p13        | Menteri dapat melakukan pengakhiran MRA sebagaimana
p13        | dimaksud dalam Pasal 18 ayat (2).
p13        I4 [AYAT]
p13        I4 | (2)
p13        | Dalam hal terjadi pengakhiran MRA, penetapan Balai Uji
p13        | Luar
p13        | Negeri
p13        | masih
p13        | tetap
p13        | berlaku
p13        | sampai
p13        | dengan
p13        | berakhirnya masa berlaku penetapan Balai Uji Luar Negeri
p13        | sebagaimana ditetapkan dalam MRA.
p13        I4 [AYAT]
p13        I4 | (3)
p13        | Pengakhiran MRA sebagaimana dimaksud pada ayat (1),
p13        | dilakukan berdasarkan hasil evaluasi yang dilakukan
p13        | Direktur Jenderal terhadap pelaksanaan MRA.
p13        | Pasal 25
p13        I4 [AYAT]
p13        I4 | (1)
p13        | Laboratorium uji luar negeri dapat ditetapkan sebagai
p13        | Balai Uji Luar Negeri melalui mekanisme non-MRA, jika:
p13        [SUB-ITEM]
p13        | a.
p13        | laboratorium uji berasal dari negara yang belum
p13        | mempunyai MRA dengan Negara Republik Indonesia;
p13        | dan
p13        [SUB-ITEM]
p13        | b.
p13        | telah diakui sebelum Peraturan Menteri ini mulai
p13        | berlaku.
p13        I4 [AYAT]
p13        I4 | (2)
p13        | Untuk dapat ditetapkan sebagai Balai Uji Luar Negeri
p13        | melalui mekanisme non-MRA sebagaimana dimaksud
p13        | pada ayat (1), laboratorium uji luar negeri harus
==================== PAGE 14 ====================
p14        | - 14 -
p14        | mengajukan permohonan kepada Menteri paling lambat 1
p14        | November 2024.
p14        I4 [AYAT]
p14        I4 | (3)
p14        | Permohonan sebagaimana dimaksud pada ayat (2)
p14        | diajukan oleh pimpinan laboratorium uji luar negeri atau
p14        | pejabat
p14        | yang
p14        | ditunjuk
p14        | sebagai
p14        | penanggung
p14        | jawab
p14        | laboratorium uji luar negeri.
p14        I4 [AYAT]
p14        I4 | (4)
p14        | Permohonan sebagaimana dimaksud pada ayat (1)
p14        | disampaikan secara elektronik dengan melampirkan
p14        | dokumen persyaratan:
p14        [SUB-ITEM]
p14        | a.
p14        | surat permohonan penetapan sebagai Balai Uji Luar
p14        | Negeri;
p14        [SUB-ITEM]
p14        | b.
p14        | bukti
p14        | berbadan
p14        | hukum
p14        | di
p14        | negara
p14        | tempat
p14        | laboratorium uji luar negeri berkedudukan atau
p14        | dokumen lain yang setara;
p14        [SUB-ITEM]
p14        | c.
p14        | daftar peralatan pengujian yang digunakan serta
p14        | metode pengujian Alat Telekomunikasi dan/atau
p14        | Perangkat Telekomunikasi yang sesuai dengan
p14        | metode pengujian berdasarkan Standar Teknis yang
p14        | berlaku di Indonesia;
p14        [SUB-ITEM]
p14        | d.
p14        | surat pernyataan secara mandiri (self declaration)
p14        | yang menyatakan tidak memiliki potensi terjadinya
p14        | konflik kepentingan dalam pelaksanaan operasional
p14        | laboratorium uji dengan Direktorat Jenderal;
p14        [SUB-ITEM]
p14        | e.
p14        | contoh salinan Laporan Hasil Uji yang terbaru yang
p14        | telah diterbitkan oleh laboratorium uji pemohon
p14        | dengan menggunakan acuan uji Standar Teknis
p14        | terkait untuk setiap ruang lingkup pengujian yang
p14        | dimohonkan;
p14        [SUB-ITEM]
p14        | f.
p14        | dokumen mutu (quality document);
p14        [SUB-ITEM]
p14        | g.
p14        | instruksi kerja yang digunakan untuk menguji Alat
p14        | Telekomunikasi dan/atau Perangkat Telekomunikasi
p14        | terhadap Standar Teknis;
p14        [SUB-ITEM]
p14        | h.
p14        | Laporan Hasil Uji profisiensi (proficiency testing
p14        | document) untuk ruang lingkup pengujian yang
p14        | dimohonkan;
p14        [SUB-ITEM]
p14        | i.
p14        | laporan audit internal dan eksternal yang dilakukan
p14        | secara berkala (periodic audit report);
p14        [SUB-ITEM]
p14        | j.
p14        | surat pernyataan kesanggupan penggunaan tanda
p14        | tangan digital yang diterbitkan oleh penyelenggara
p14        | sistem elektronik yang terdaftar di negaranya dan
p14        | disertai
p14        | panduan
p14        | pengecekan
p14        | keaslian
p14        | atau
p14        | keabsahan tanda tangan digital; dan
p14        [SUB-ITEM]
p14        | k.
p14        | salinan sertifikat dan ruang lingkup akreditasi
p14        | ISO/IEC 17025 termutakhir yang diterbitkan oleh
p14        | Lembaga Akreditasi penandatangan Asia Pacific
p14        | Accreditation
p14        | Cooperation-Mutual
p14        | Recognition
p14        | Arrangement
p14        | (APAC-MRA)
p14        | atau
p14        | International
p14        | Laboratory
p14        | Accreditation
p14        | Cooperation-Mutual
p14        | Recognition Arrangement (ILAC-MRA) di negara sesuai
p14        | negara asal laboratorium uji;
p14        [SUB-ITEM]
p14        | l.
p14        | paling sedikit 2 (dua) bukti dokumen berupa:
p14        [ITEM]
p14        | 1.
p14        | akreditasi dari Lembaga Akreditasi negara lain;
p14        [ITEM]
p14        | 2.
p14        | pengakuan dari lembaga internasional yang
p14        | melaksanakan fungsi penilaian kesesuaian Alat
==================== PAGE 15 ====================
p15        | - 15 -
p15        | Telekomunikasi
p15        | dan/atau
p15        | Perangkat
p15        | Telekomunikasi; atau
p15        [ITEM]
p15        | 3.
p15        | pengakuan administrasi telekomunikasi negara
p15        | lain.
p15        I4 [AYAT]
p15        I4 | (5)
p15        | Dalam hal laboratorium uji tidak memiliki Laporan Hasil
p15        | Uji profisiensi sebagaimana dimaksud pada ayat (4) huruf
p15        | h karena program uji profisiensi untuk ruang lingkup
p15        | pengujian yang dimohonkan tidak tersedia, pemohon
p15        | dapat
p15        | menyampaikan
p15        | dokumen
p15        | uji
p15        | banding
p15        | antarlaboratorium (inter-laboratory comparison test) pada
p15        | ruang lingkup pengujian yang dimohonkan.
p15        I4 [AYAT]
p15        I4 | (6)
p15        | Dokumen persyaratan permohonan penetapan Balai Uji
p15        | Luar Negeri sebagaimana dimaksud pada ayat (4) harus
p15        | menggunakan:
p15        [SUB-ITEM]
p15        | a.
p15        | bahasa Indonesia;
p15        [SUB-ITEM]
p15        | b.
p15        | bahasa Inggris; atau
p15        [SUB-ITEM]
p15        | c.
p15        | bahasa asing lainnya, yang disertai terjemahan resmi
p15        | menggunakan bahasa Indonesia dan/atau bahasa
p15        | Inggris.
p15        | Pasal 26
p15        I4 [AYAT]
p15        I4 | (1)
p15        | Direktur Jenderal melaksanakan verifikasi terhadap
p15        | permohonan penetapan sebagai Balai Uji Luar Negeri
p15        | sebagaimana dimaksud dalam Pasal 25 ayat (2) setelah
p15        | dokumen persyaratan permohonan dinyatakan lengkap.
p15        I4 [AYAT]
p15        I4 | (2)
p15        | Verifikasi sebagaimana dimaksud pada ayat (1) dilakukan
p15        | terhadap:
p15        [SUB-ITEM]
p15        | a.
p15        | keabsahan
p15        | dokumen
p15        | persyaratan
p15        | permohonan
p15        | sebagaimana dimaksud dalam Pasal 25 ayat (4);
p15        [SUB-ITEM]
p15        | b.
p15        | kesiapan laboratorium uji, yang meliputi:
p15        [ITEM]
p15        | 1.
p15        | kompetensi penguji terhadap Standar Teknis
p15        | yang berlaku di Indonesia;
p15        [ITEM]
p15        | 2.
p15        | pelaksanan pengujian berdasarkan metode uji
p15        | sesuai Standar Teknis;
p15        [ITEM]
p15        | 3.
p15        | sarana dan prasarana pengujian yang dimiliki
p15        | dan kesesuaiannya dengan ruang lingkup
p15        | pengujian berdasarkan kebutuhan parameter
p15        | uji sesuai Standar Teknis.
p15        | Pasal 27
p15        I4 [AYAT]
p15        I4 | (1)
p15        | Berdasarkan hasil verifikasi sebagaimana dimaksud
p15        | dalam Pasal 25, Menteri menyetujui atau menolak
p15        | permohonan penetapan Balai Uji Luar Negeri melalui
p15        | mekanisme non-MRA.
p15        I4 [AYAT]
p15        I4 | (2)
p15        | Dalam hal permohonan disetujui, Menteri menerbitkan
p15        | penetapan Balai Uji Luar Negeri.
p15        I4 [AYAT]
p15        I4 | (3)
p15        | Dalam hal permohonan ditolak, Direktur Jenderal
p15        | menyampaikan surat penolakan kepada pemohon paling
p15        | lama 40 (empat puluh) Hari sejak permohonan diterima
p15        | secara lengkap.
p15        I4 [AYAT]
p15        I4 | (4)
p15        | Penetapan Balai Uji Luar Negeri melalui mekanisme non-
p15        | MRA sebagaimana dimaksud pada ayat (2) berlaku sampai
p15        | dengan 31 Desember 2026.
p15        I4 [AYAT]
p15        I4 | (5)
p15        | Penetapan Balai Uji Luar Negeri atau surat penolakan
p15        | permohonan penetapan sebagaimana dimaksud pada ayat
==================== PAGE 16 ====================
p16        | - 16 -
p16        [AYAT]
p16        | (2) dan ayat (3) disampaikan kepada pemohon secara
p16        | elektronik.
p16        | Pasal 28
p16        I4 [AYAT]
p16        I4 | (1)
p16        | Balai Uji Luar Negeri sebagaimana dimaksud dalam Pasal
p16        | 22 ayat (2) dan Pasal 27 ayat (2) wajib:
p16        [SUB-ITEM]
p16        | a.
p16        | melaksanakan
p16        | pengujian
p16        | Alat
p16        | Telekomunikasi
p16        | dan/atau Perangkat Telekomunikasi sesuai Standar
p16        | Teknis yang berlaku di Indonesia dan ruang lingkup
p16        | pengujian yang telah ditetapkan;
p16        [SUB-ITEM]
p16        | b.
p16        | melampirkan rangkuman referensi halaman Laporan
p16        | Hasil Uji yang terkait dengan persyaratan teknis
p16        | Indonesia yang menjadi acuan pengujian;
p16        [SUB-ITEM]
p16        | c.
p16        | menggunakan tanda tangan digital pada Laporan
p16        | Hasil Uji;
p16        [SUB-ITEM]
p16        | d.
p16        | memberikan klarifikasi keabsahan Laporan Hasil Uji
p16        | dalam hal diperlukan oleh Menteri; dan
p16        [SUB-ITEM]
p16        | e.
p16        | melaporkan kepada Direktur Jenderal dalam hal
p16        | terjadi perubahan:
p16        [ITEM]
p16        | 1.
p16        | status badan hukum;
p16        [ITEM]
p16        | 2.
p16        | bidang usaha;
p16        [ITEM]
p16        | 3.
p16        | struktur organisasi;
p16        [ITEM]
p16        | 4.
p16        | akreditasi;
p16        [ITEM]
p16        | 5.
p16        | alamat Balai Uji Luar Negeri;
p16        [ITEM]
p16        | 6.
p16        | penanggung jawab Balai Uji Luar Negeri;
p16        | dan/atau
p16        [ITEM]
p16        | 7.
p16        | yang
p16        | dapat
p16        | memengaruhi
p16        | kesinambungan
p16        | pengujian.
p16        I4 [AYAT]
p16        I4 | (2)
p16        | Laporan sebagaimana dimaksud pada ayat (1) huruf e dari
p16        | Balai Uji Luar Negeri yang ditetapkan melalui mekanisme
p16        | MRA disampaikan melalui Mitra MRA.
p16        [HEADING:BAB]
p16        | BAB IV
p16        | PENGAWASAN DAN PENGENDALIAN
p16        | Pasal 29
p16        I4 [AYAT]
p16        I4 | (1)
p16        | Direktur
p16        | Jenderal
p16        | melakukan
p16        | pengawasan
p16        | dan
p16        | pengendalian terhadap Balai Uji Dalam Negeri dan Balai
p16        | Uji Luar Negeri.
p16        I4 [AYAT]
p16        I4 | (2)
p16        | Pengawasan sebagaimana dimaksud pada ayat (1)
p16        | dilaksanakan secara:
p16        [SUB-ITEM]
p16        | a.
p16        | rutin; dan
p16        [SUB-ITEM]
p16        | b.
p16        | insidental.
p16        | Pasal 30
p16        I4 [AYAT]
p16        I4 | (1)
p16        | Pengawasan terhadap Balai Uji Dalam Negeri secara rutin
p16        | sebagaimana dimaksud dalam Pasal 29 ayat (2) huruf a
p16        | dilaksanakan paling sedikit 1 (satu) kali pada periode
p16        | masa laku penetapan Balai Uji Dalam Negeri.
p16        I4 [AYAT]
p16        I4 | (2)
p16        | Pengawasan sebagaimana dimaksud pada ayat (1)
p16        | dilaksanakan melalui verifikasi terhadap:
p16        [SUB-ITEM]
p16        | a.
p16        | status akreditasi Balai Uji Dalam Negeri termutakhir
p16        | yang diterbitkan oleh KAN;
p16        [SUB-ITEM]
p16        | b.
p16        | pemenuhan kewajiban oleh Balai Uji Dalam Negeri
p16        | sebagaimana dimaksud dalam Pasal 16; dan
==================== PAGE 17 ====================
p17        | - 17 -
p17        [SUB-ITEM]
p17        | c.
p17        | fungsi Balai Uji Dalam Negeri dalam melaksanakan
p17        | pengujian Alat Telekomunikasi dan/atau Perangkat
p17        | Telekomunikasi.
p17        I4 [AYAT]
p17        I4 | (3) Pengawasan terhadap Balai Uji Dalam Negeri secara
p17        | insidental sebagaimana dimaksud dalam Pasal 29 ayat (2)
p17        | huruf b dilaksanakan dalam hal terdapat:
p17        [SUB-ITEM]
p17        | a.
p17        | perubahan perizinan berusaha;
p17        [SUB-ITEM]
p17        | b.
p17        | perubahan struktur organisasi;
p17        [SUB-ITEM]
p17        | c.
p17        | perubahan akreditasi;
p17        [SUB-ITEM]
p17        | d.
p17        | perubahan alamat Balai Uji Dalam Negeri;
p17        [SUB-ITEM]
p17        | e.
p17        | penurunan kualitas pengujian dan/atau fasilitas
p17        | pengujian; dan/atau
p17        [SUB-ITEM]
p17        | f.
p17        | perubahan
p17        | lainnya
p17        | yang
p17        | dapat
p17        | memengaruhi
p17        | kesinambungan pengujian.
p17        | Pasal 31
p17        I4 [AYAT]
p17        I4 | (1)
p17        | Berdasarkan hasil pengawasan sebagaimana dimaksud
p17        | dalam
p17        | Pasal
p17        | 30,
p17        | Direktur
p17        | Jenderal
p17        | melakukan
p17        | pengendalian terhadap Balai Uji Dalam Negeri dalam hal
p17        | ditemukenali:
p17        [SUB-ITEM]
p17        | a.
p17        | Balai Uji Dalam Negeri tidak dapat memenuhi
p17        | kewajiban sebagaimana dimaksud dalam Pasal 16;
p17        [SUB-ITEM]
p17        | b.
p17        | akreditasi Balai Uji Dalam Negeri telah dicabut atau
p17        | dibekukan oleh KAN; atau
p17        [SUB-ITEM]
p17        | c.
p17        | masa laku akreditasi Balai Uji Dalam Negeri yang
p17        | diterbitkan oleh KAN telah berakhir.
p17        I4 [AYAT]
p17        I4 | (2)
p17        | Pengendalian sebagaimana dimaksud pada ayat (1)
p17        | dilaksanakan melalui pembekuan atau pencabutan
p17        | penetapan Balai Uji Dalam Negeri atau sebagian ruang
p17        | lingkup pengujian yang ditetapkan.
p17        I4 [AYAT]
p17        I4 | (3)
p17        | Pembekuan atau pencabutan Balai Uji Dalam Negeri atau
p17        | sebagian ruang lingkup pengujian yang ditetapkan
p17        | sebagaimana dimaksud dalam pada ayat (2) dilakukan
p17        | oleh Menteri atau Direktur Jenderal sesuai dengan
p17        | kewenangannya.
p17        I4 [AYAT]
p17        I4 | (4)
p17        | Balai Uji Dalam Negeri yang penetapannya dibekukan
p17        | sebagaimana dimaksud pada ayat (1) dapat mengajukan
p17        | permohonan pengaktifan kembali penetapannya dengan
p17        | menunjukkan bukti bahwa hal yang menyebabkan
p17        | pembekuannya telah terpenuhi.
p17        I4 [AYAT]
p17        I4 | (5)
p17        | Permohonan pengaktifan sebagaimana dimaksud pada
p17        | ayat (3) disampaikan kepada Menteri.
p17        I4 [AYAT]
p17        I4 | (6)
p17        | Direktur
p17        | Jenderal
p17        | melakukan
p17        | evaluasi
p17        | terhadap
p17        | permohonan pengaktifan sebagaimana dimaksud pada
p17        | ayat (4).
p17        I4 [AYAT]
p17        I4 | (7)
p17        | Berdasarkan hasil evaluasi sebagaimana dimaksud pada
p17        | ayat (5), Menteri dapat menyetujui atau menolak
p17        | permohonan pengaktifan penetapan Balai Uji Dalam
p17        | Negeri.
p17        | Pasal 32
p17        I4 [AYAT]
p17        I4 | (1)
p17        | Pengawasan terhadap Balai Uji Luar Negeri secara rutin
p17        | sebagaimana dimaksud dalam Pasal 29 ayat (2) huruf a
p17        | dilaksanakan paling sedikit 1 (satu) kali pada periode
==================== PAGE 18 ====================
p18        | - 18 -
p18        | masa laku penetapan Balai Uji Luar Negeri, baik yang
p18        | ditetapkan melalui mekanisme MRA maupun non-MRA.
p18        I4 [AYAT]
p18        I4 | (2)
p18        | Pengawasan rutin terhadap Balai Uji Luar Negeri yang
p18        | ditetapkan
p18        | melalui
p18        | mekanisme
p18        | MRA
p18        | sebagaimana
p18        | dimaksud dalam Pasal 22 ayat (2) dilaksanakan melalui
p18        | evaluasi terhadap:
p18        [SUB-ITEM]
p18        | a.
p18        | status MRA;
p18        [SUB-ITEM]
p18        | b.
p18        | masa laku penetapan Balai Uji Luar Negeri dari
p18        | Mitra MRA;
p18        [SUB-ITEM]
p18        | c.
p18        | status akreditasi Balai Uji Luar Negeri yang
p18        | diterbitkan oleh Lembaga Akreditasi Mitra MRA;
p18        [SUB-ITEM]
p18        | d.
p18        | pemenuhan kewajiban oleh Balai Uji Luar Negeri
p18        | sebagaimana dimaksud Pasal 28 ayat (1); dan
p18        [SUB-ITEM]
p18        | e.
p18        | fungsi dan kemampuan atau kompetensi teknis
p18        | dalam melakukan pengujian Alat Telekomunikasi
p18        | dan/atau Perangkat Telekomunikasi sesuai dengan
p18        | Standar Teknis yang berlaku di Indonesia.
p18        I4 [AYAT]
p18        I4 | (3)
p18        | Pengawasan rutin terhadap Balai Uji Luar Negeri yang
p18        | ditetapkan melalui mekanisme non-MRA sebagaimana
p18        | dimaksud dalam Pasal 27 ayat (2) dilaksanakan melalui
p18        | evaluasi terhadap:
p18        [SUB-ITEM]
p18        | a.
p18        | status akreditasi Balai Uji Luar Negeri yang
p18        | diterbitkan oleh Lembaga Akreditasi negara dimana
p18        | Balai Uji Luar Negeri berkedudukan;
p18        [SUB-ITEM]
p18        | b.
p18        | pemenuhan kewajiban oleh Balai Uji Luar Negeri
p18        | sebagaimana dimaksud Pasal 28 ayat (1); dan
p18        [SUB-ITEM]
p18        | c.
p18        | fungsi dan kemampuan atau kompetensi teknis
p18        | dalam melakukan pengujian Alat Telekomunikasi
p18        | dan/atau Perangkat Telekomunikasi sesuai dengan
p18        | Standar Teknis yang berlaku di Indonesia.
p18        I4 [AYAT]
p18        I4 | (4)
p18        | Pengawasan terhadap Balai Uji Luar Negeri secara
p18        | insidental sebagaimana dimaksud dalam Pasal 29 ayat (2)
p18        | huruf b dilaksanakan dalam hal terdapat:
p18        [SUB-ITEM]
p18        | a.
p18        | perubahan status badan hukum;
p18        [SUB-ITEM]
p18        | b.
p18        | perubahan bidang usaha;
p18        [SUB-ITEM]
p18        | c.
p18        | perubahan struktur organisasi;
p18        [SUB-ITEM]
p18        | d.
p18        | perubahan akreditasi;
p18        [SUB-ITEM]
p18        | e.
p18        | perubahan alamat Balai Uji Luar Negeri;
p18        [SUB-ITEM]
p18        | f.
p18        | penurunan
p18        | kualitas
p18        | dan
p18        | fasilitas
p18        | pengujian;
p18        | dan/atau
p18        [SUB-ITEM]
p18        | g.
p18        | perubahan
p18        | lainnya
p18        | yang
p18        | memengaruhi
p18        | kesinambungan pengujian sesuai dengan Standar
p18        | Teknis.
p18        | Pasal 33
p18        I4 [AYAT]
p18        I4 | (1)
p18        | Berdasarkan hasil pengawasan sebagaimana dimaksud
p18        | dalam
p18        | Pasal
p18        | 32,
p18        | Direktur
p18        | Jenderal
p18        | melakukan
p18        | pengendalian terhadap Balai Uji Luar Negeri yang
p18        | ditetapkan
p18        | melalui
p18        | mekanisme
p18        | MRA
p18        | dalam
p18        | hal
p18        | ditemukenali:
p18        [SUB-ITEM]
p18        | a.
p18        | MRA dengan Mitra MRA telah berakhir;
p18        [SUB-ITEM]
p18        | b.  masa laku penetapan dari Badan Penetap Mitra MRA
p18        | berakhir dan tidak diperpanjang;
p18        [SUB-ITEM]
p18        | c.  akreditasi Balai Uji Luar Negeri telah dicabut atau
p18        | dibekukan oleh Lembaga Akreditasi Mitra MRA;
==================== PAGE 19 ====================
p19        | - 19 -
p19        [SUB-ITEM]
p19        | d.  Balai Uji Luar Negeri tidak dapat memenuhi
p19        | kewajiban sebagaimana dimaksud dalam Pasal 28
p19        | ayat (1); atau
p19        [SUB-ITEM]
p19        | e.
p19        | Balai Uji Luar Negeri tidak lagi memiliki kemampuan
p19        | atau kompetensi teknis dalam melakukan pengujian
p19        | Alat
p19        | Telekomunikasi
p19        | dan/atau
p19        | Perangkat
p19        | Telekomunikasi sesuai dengan Standar Teknis yang
p19        | berlaku di Indonesia.
p19        I4 [AYAT]
p19        I4 | (2)
p19        | Berdasarkan hasil pengawasan sebagaimana dimaksud
p19        | dalam
p19        | Pasal
p19        | 32,
p19        | Direktur
p19        | Jenderal
p19        | melakukan
p19        | pengendalian terhadap Balai Uji Luar Negeri yang
p19        | ditetapkan melalui mekanisme non-MRA dalam hal
p19        | ditemukenali:
p19        [SUB-ITEM]
p19        | a.
p19        | akreditasi Balai Uji Luar Negeri telah dicabut atau
p19        | dibekukan oleh Lembaga Akreditasi negara dimana
p19        | Balai Uji Luar Negeri berkedudukan;
p19        [SUB-ITEM]
p19        | b.
p19        | Balai Uji Luar Negeri tidak dapat memenuhi
p19        | kewajiban sebagaimana dimaksud dalam Pasal 28
p19        | ayat (1);
p19        [SUB-ITEM]
p19        | c.
p19        | Balai Uji Luar Negeri tidak lagi memiliki kemampuan
p19        | atau kompetensi teknis dalam melakukan pengujian
p19        | Alat
p19        | Telekomunikasi
p19        | dan/atau
p19        | Perangkat
p19        | Telekomunikasi sesuai dengan Standar Teknis yang
p19        | berlaku di Indonesia; atau
p19        [SUB-ITEM]
p19        | d.
p19        | telah terdapat MRA sebagaimana dimaksud dalam
p19        | Pasal 18 yang berlaku di negara di mana Balai Uji
p19        | Luar Negeri berkedudukan.
p19        I4 [AYAT]
p19        I4 | (3)
p19        | Pengendalian sebagaimana dimaksud pada ayat (1) dan
p19        | ayat
p19        [AYAT]
p19        | (2)
p19        | dilaksanakan
p19        | melalui
p19        | pembekuan
p19        | atau
p19        | pencabutan penetapan Balai Uji Luar Negeri atau
p19        | sebagian ruang lingkup pengujian yang ditetapkan.
p19        I4 [AYAT]
p19        I4 | (4)
p19        | Pembekuan atau pencabutan Balai Uji Luar Negeri atau
p19        | sebagian ruang lingkup pengujian yang ditetapkan
p19        | sebagaimana dimaksud dalam pada ayat (3) dilakukan
p19        | oleh Menteri atau Direktur Jenderal sesuai dengan
p19        | kewenangannya.
p19        I4 [AYAT]
p19        I4 | (5)
p19        | Balai Uji Luar Negeri yang penetapannya dibekukan
p19        | sebagaimana dimaksud pada ayat (4) dapat mengajukan
p19        | permohonan pengaktifan kembali penetapannya dengan
p19        | menunjukkan bukti bahwa hal yang menyebabkan
p19        | pembekuannya telah terpenuhi.
p19        I4 [AYAT]
p19        I4 | (6)
p19        | Permohonan pengaktifan sebagaimana dimaksud pada
p19        | ayat (4) disampaikan kepada Menteri, dengan ketentuan:
p19        [SUB-ITEM]
p19        | a.
p19        | diajukan melalui Badan Penetap Mitra MRA untuk
p19        | Balai Uji Luar Negeri yang ditetapkan melalui
p19        | mekanisme MRA; atau
p19        [SUB-ITEM]
p19        | b.
p19        | diajukan secara langsung oleh Balai uji Luar Negeri
p19        | yang bersangkutan untuk Balai Uji Luar Negeri yang
p19        | ditetapkan melalui mekanisme non-MRA.
p19        I4 [AYAT]
p19        I4 | (7)
p19        | Direktur
p19        | Jenderal
p19        | melakukan
p19        | evaluasi
p19        | terhadap
p19        | permohonan pengaktifan sebagaimana dimaksud pada
p19        | ayat (6).
p19        I4 [AYAT]
p19        I4 | (8)
p19        | Berdasarkan hasil evaluasi sebagaimana dimaksud pada
p19        | ayat (7), Menteri dapat menyetujui atau menolak
p19        | permohonan pengaktifan Balai Uji Luar Negeri.
==================== PAGE 20 ====================
p20        | - 20 -
p20        | Pasal 34
p20 B      I4 [AYAT]
p20 B      I4 | (1)
p20        | Pencabutan penetapan Balai Uji Dalam Negeri, penetapan
p20        | Balai Uji Luar Negeri, atau sebagian ruang lingkup
p20        | pengujian
p20        | yang
p20        | ditetapkan
p20        | juga
p20        | dapat
p20        | dilakukan
p20        | berdasarkan permohonan dari Balai Uji Dalam Negeri
p20        | atau Balai Uji Luar Negeri.
p20        I4 [AYAT]
p20        I4 | (2)
p20        | Permohonan sebagaimana dimaksud pada ayat (1)
p20        | disampaikan kepada Menteri.
p20        I4 [AYAT]
p20        I4 | (3)
p20        | Berdasarkan permohonan sebagaimana dimaksud pada
p20        | ayat (2), Menteri menetapkan pencabutan penetapan:
p20        [SUB-ITEM]
p20        | a. Balai Uji Dalam Negeri;
p20        [SUB-ITEM]
p20        | b. Balai Uji Luar Negeri; atau
p20        [SUB-ITEM]
p20        | c. sebagian ruang lingkup pengujian.
p20        | Pasal 35
p20        I4 [AYAT]
p20        I4 | (1)
p20        | Balai Uji Dalam Negeri atau Balai Uji Luar Negeri yang
p20        | penetapannya dicabut berdasarkan:
p20        [SUB-ITEM]
p20        | a.
p20        | hasil pengawasan dan pengendalian sebagaimana
p20        | dimaksud dalam Pasal 31 ayat (2) dan Pasal 33 ayat
p20        [AYAT]
p20        | (3); atau
p20        [SUB-ITEM]
p20        | b.
p20        | permohonan sebagaimana dimaksud dalam Pasal 34
p20        | ayat (1),
p20        | hanya dapat mengajukan kembali permohonan penetapan
p20        | sebagai Balai Uji Dalam Negeri, sebagai Balai Uji Luar
p20        | Negeri, atau sebagian ruang lingkup pengujian yang telah
p20        | dicabut penetapannya setelah 1 (satu) tahun sejak tanggal
p20        | pencabutan sebagaimana dimaksud dalam Pasal 34 ayat
p20        [AYAT]
p20        | (3).
p20        I4 [AYAT]
p20        I4 | (2)
p20        | Ketentuan sebagaimana dimaksud dalam Pasal 27 ayat (4)
p20        | tetap berlaku dalam hal pengajuan kembali permohonan
p20        | penetapan sebagai Balai Uji Luar Negeri sebagaimana
p20        | dimaksud pada ayat (1) diajukan oleh Balai Uji Luar
p20        | Negeri yang ditetapkan melalui mekanisme non-MRA.
p20        | Pasal 36
p20        I4 | Tata
p20        | cara
p20        | pelaksanaan
p20        | pengawasan
p20        | dan
p20        | pengendalian
p20        I4 | sebagaimana dimaksud dalam Pasal 29 sampai dengan Pasal
p20        I4 | 35 ditetapkan oleh Direktur Jenderal.
p20        [HEADING:BAB]
p20        | BAB V
p20        | KETENTUAN LAIN-LAIN
p20        | Pasal 37
p20        I4 [AYAT]
p20        I4 | (1)
p20        | Balai Uji Dalam Negeri yang telah mendapatkan:
p20        [SUB-ITEM]
p20        | a.
p20        | penetapan Balai Uji Dalam Negeri sebagaimana
p20        | dimaksud dalam Pasal 6 ayat (2);
p20        [SUB-ITEM]
p20        | b.
p20        | perpanjangan penetapan Balai Uji Dalam Negeri
p20        | sebagaimana dimaksud dalam Pasal 10 ayat (2); dan
p20        [SUB-ITEM]
p20        | c.
p20        | penetapan penambahan ruang lingkup pengujian
p20        | sebagaimana dimaksud dalam Pasal 13 ayat (2),
p20        | dicantumkan dan/atau dilakukan pembaruan informasi
p20        | pada situs web Direktorat Jenderal.
p20        I4 [AYAT]
p20        I4 | (2)
p20        | Balai Uji Dalam Negeri sebagaimana dimaksud pada ayat
p20        [AYAT]
p20        | (1) dapat mengumumkan status penetapan Balai Uji
==================== PAGE 21 ====================
p21        | - 21 -
p21        | Dalam Negeri dan ruang lingkup pengujian pada situs web
p21        | milik Balai Uji Dalam Negeri.
p21        I4 [AYAT]
p21        I4 | (3)
p21        | Balai Uji Dalam Negeri yang telah berakhir masa laku
p21        | penetapannya sebagaimana dimaksud dalam Pasal 7
p21        | dihapus dari daftar Balai Uji Dalam Negeri pada situs web
p21        | Direktorat Jenderal.
p21        I4 [AYAT]
p21        I4 | (4)
p21        | Ruang lingkup pengujian Balai Uji Dalam Negeri yang
p21        | dinyatakan
p21        | batal
p21        | dan
p21        | tidak
p21        | berlaku
p21        | sebagaimana
p21        | dimaksud dalam Pasal 14 ayat (3), dihapus dari situs web
p21        | Direktorat Jenderal oleh Direktur Jenderal.
p21        I4 [AYAT]
p21        I4 | (5)
p21        | Informasi status penetapan dan/atau ruang lingkup
p21        | pengujian Balai Uji Dalam Negeri sebagaimana dimaksud
p21        | pada ayat (3) dan ayat (4) diperbarui dalam situs web milik
p21        | Balai Uji Dalam Negeri.
p21        | Pasal 38
p21        I4 [AYAT]
p21        I4 | (1)
p21        | Daftar Balai Uji Dalam Negeri dapat diperbarui dalam hal
p21        | terdapat perubahan status penetapan sebagai berikut:
p21        [SUB-ITEM]
p21        | a.
p21        | pembekuan dan/atau pencabutan penetapan Balai
p21        | Uji Dalam Negeri; atau
p21        [SUB-ITEM]
p21        | b.
p21        | pembekuan dan/atau pencabutan sebagian ruang
p21        | lingkup pengujian Balai Uji Dalam Negeri.
p21        I4 [AYAT]
p21        I4 | (2)
p21        | Direktur Jenderal mengumumkan perubahan status
p21        | penetapan atau ruang lingkup pengujian Balai Uji Dalam
p21        | Negeri sebagaimana dimaksud pada ayat (1) dalam situs
p21        | web Direktorat Jenderal.
p21        I4 [AYAT]
p21        I4 | (3)
p21        | Balai Uji Dalam Negeri mengumumkan perubahan status
p21        | penetapan atau ruang lingkup pengujian Balai Uji Dalam
p21        | Negeri sebagaimana dimaksud pada ayat (1) dalam situs
p21        | web milik Balai Uji Dalam Negeri.
p21        | Pasal 39
p21        I4 [AYAT]
p21        I4 | (1)
p21        | Balai Uji Luar Negeri yang telah mendapatkan:
p21        [SUB-ITEM]
p21        | a.
p21        | penetapan Balai Uji Luar Negeri sebagaimana
p21        | dimaksud dalam Pasal 22 ayat (2) dan Pasal 27 ayat
p21        [AYAT]
p21        | (2);
p21        [SUB-ITEM]
p21        | b.
p21        | perpanjangan penetapan Balai Uji Luar Negeri
p21        | sebagaimana dimaksud dalam Pasal 23 ayat (5); dan
p21        [SUB-ITEM]
p21        | c.
p21        | penetapan penambahan ruang lingkup pengujian
p21        | sebagaimana dimaksud dalam Pasal 23 ayat (5),
p21        | dicantumkan dan/atau dilakukan pembaruan informasi
p21        | pada situs web Direktorat Jenderal.
p21        I4 [AYAT]
p21        I4 | (2)
p21        | Balai Uji Luar Negeri sebagaimana dimaksud pada ayat (1)
p21        | dapat mengumumkan status penetapan sebagai Balai Uji
p21        | Luar Negeri dan ruang lingkup pengujian pada situs web
p21        | milik Balai Uji Luar Negeri.
p21        I4 [AYAT]
p21        I4 | (3)
p21        | Direktorat Jenderal menghapus Balai Uji Luar Negeri dari
p21        | situs web Direktorat Jenderal setelah berakhirnya masa
p21        | laku penetapan.
p21        | Pasal 40
p21        I4 [AYAT]
p21        I4 | (1)
p21        | Daftar Balai Uji Luar Negeri dapat diperbarui dalam hal
p21        | terdapat pembekuan dan/atau pencabutan penetapan
p21        | Balai Uji Luar Negeri atau sebagian ruang lingkup
p21        | pengujian yang ditetapkan sebagaimana dimaksud dalam
==================== PAGE 22 ====================
p22        | - 22 -
p22        | Pasal 33 ayat (3).
p22        I4 [AYAT]
p22        I4 | (2)
p22        | Direktur Jenderal mengumumkan pembekuan dan/atau
p22        | pencabutan penetapan Balai Uji Luar Negeri atau
p22        | sebagian ruang lingkup pengujian yang ditetapkan
p22        | sebagaimana dimaksud pada ayat (1) dalam situs web
p22        | Direktorat Jenderal.
p22        I4 [AYAT]
p22        I4 | (3)
p22        | Balai Uji Luar Negeri juga mengumumkan pembekuan
p22        | dan/atau pencabutan penetapan Balai Uji Luar Negeri
p22        | atau sebagian ruang lingkup pengujian yang ditetapkan
p22        | sebagaimana dimaksud pada ayat (1) dalam situs web
p22        | milik Balai Uji Luar Negeri.
p22        | Pasal 41
p22        I4 | Laporan Hasil Uji yang telah diterbitkan oleh Balai Uji Dalam
p22        I4 | Negeri atau Balai Uji Luar Negeri yang telah:
p22        I4 [SUB-ITEM]
p22        I4 | a.
p22        | dicabut status penetapan;
p22        I4 [SUB-ITEM]
p22        I4 | b.
p22        | dibatalkan status penetapan penambahan ruang lingkup
p22        | pengujian; atau
p22        I4 [SUB-ITEM]
p22        I4 | c.
p22        | berakhir masa berlaku status penetapan,
p22        I4 | masih dapat digunakan untuk keperluan Sertifikasi paling
p22        I4 | lama 6 (enam) bulan sejak tanggal diterbitkan pencabutan
p22        I4 | status penetapan, pembatalan status penetapan, atau berakhir
p22        I4 | masa berlaku status penetapan sebagaimana dimaksud pada
p22        I4 | huruf a, huruf b, atau huruf c.
p22        [HEADING:BAB]
p22        | BAB VI
p22        | KETENTUAN PERALIHAN
p22        | Pasal 42
p22        I4 | Balai Uji Luar Negeri yang telah diakui sebelum Peraturan
p22        I4 | Menteri ini mulai berlaku dan memenuhi ketentuan sebagai
p22        I4 | berikut:
p22        I4 [SUB-ITEM]
p22        I4 | a. tidak mengajukan permohonan penetapan sebagai Balai Uji
p22        | Luar Negeri melalui mekanisme non-MRA sampai dengan
p22        | batas waktu sebagaimana dimaksud dalam Pasal 25 ayat
p22        [AYAT]
p22        | (2); atau
p22        I4 [SUB-ITEM]
p22        I4 | b. mengajukan permohonan penetapan sebagai Balai Uji Luar
p22        | Negeri melalui mekanisme non-MRA tetapi permohonannya
p22        | ditolak sebagaimana dimaksud dalam Pasal 27 ayat (3),
p22        I4 | tetap diakui sebagai Balai Uji Luar Negeri sampai dengan
p22        I4 | tanggal 31 Desember 2024.
p22        [HEADING:BAB]
p22        | BAB VII
p22        | KETENTUAN PENUTUP
p22        | Pasal 43
p22        I4 | Pada saat Peraturan Menteri ini mulai berlaku:
p22        I4 [SUB-ITEM]
p22        I4 | a. Peraturan Menteri Komunikasi dan Informatika Nomor 15
p22        | Tahun 2012 tentang Petunjuk Pelaksanaan Penetapan
p22        | Balai Uji Dalam Negeri (Berita Negara Republik Indonesia
p22        | Tahun 2012 Nomor 577); dan
p22        I4 [SUB-ITEM]
p22        I4 | b. Peraturan Menteri Komunikasi dan Informatika Nomor 16
p22        | Tahun 2012 tentang Petunjuk Pelaksanaan Pengakuan
p22        | Balai Uji Negara Asing (Berita Negara Republik Indonesia
p22        | Tahun 2012 Nomor 578),
==================== PAGE 23 ====================
p23        | - 23 -
p23 B      I4 | dicabut dan dinyatakan tidak berlaku.
p23        | Pasal 44
p23        I4 | Peraturan
p23        | Menteri
p23        | ini
p23        | mulai
p23        | berlaku
p23        | pada
p23        | tanggal
p23        I4 | diundangkan.
p23        I4 | Agar
p23        | setiap
p23        | orang
p23        | mengetahuinya,
p23        | memerintahkan
p23        I4 | pengundangan Peraturan Menteri ini dengan penempatannya
p23        I4 | dalam Berita Negara Republik Indonesia.
p23        | Ditetapkan di Jakarta
p23        | pada tanggal 12 September 2024
p23        | MENTERI KOMUNIKASI DAN INFORMATIKA
p23        | REPUBLIK INDONESIA,
p23        | Œ
p23        | BUDI ARIE SETIADI
p23        I1 | Diundangkan di Jakarta
p23        I1 | pada tanggal
p23        I4 | Д
p23        I1 | PLT. DIREKTUR JENDERAL
p23        I1 | PERATURAN PERUNDANG-UNDANGAN
p23        I1 | KEMENTERIAN HUKUM DAN HAK ASASI MANUSIA
p23        I1 | REPUBLIK INDONESIA,
p23        I2 | Ѽ
p23        I1 | ASEP N. MULYANA
p23        I1 | BERITA NEGARA REPUBLIK INDONESIA TAHUN 2024 NOMOR
p23        | Ж
```

---


## JDIH_KPU

- **File**: `JDIH_KPU/PKPU_8_2026.pdf`
- **Document Type**: Peraturan KPU
- **Issued by**: Komisi Pemilihan Umum
- **Pages**: 6 | **Lines**: 286
- **Font sizes**: [12.0]
- **Most common font**: 12.0 (100% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [85.0, 170.0, 198.0, 227.0, 248.0, 281.0]
- **Expected hierarchy**: BAB > Bagian > Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 2 ====================
p02        I6 | - 2 -
p02        I3 | Daerah (Lembaran Negara Republik Indonesia Tahun 201
p02        I3 | Nomor 117, Tambahan Lembaran Negara Republik
p02        I3 | Indonesia Nomor 5316);
p02        I2 [ITEM]
p02        I2 | 4.
p02        I3 | Peraturan Komisi Pemilihan Umum Nomor 07 Tahun
p02        I3 | 2012
p02        I4 | tentang
p02        | Tahapan,
p02        | Program,
p02        | dan
p02        | Jadual
p02        I3 | Penyelenggaraan
p02        | Pemilihan
p02        | Umum
p02        | Anggota
p02        | Dewan
p02        I3 | Perwakilan Rakyat, Dewan Perwakilan Daerah, dan
p02        I3 | Dewan
p02        I5 | Perwakilan
p02        | Rakyat
p02        | Daerah
p02        | Tahun
p02        | 2014,
p02        I3 | sebagaimana telah beberapa kali diubah, terakhir dengan
p02        I3 | Peraturan Komisi Pemilihan Umum Nomor 21 Tahun
p02        I3 | 2013;
p02        I2 [ITEM]
p02        I2 | 5.
p02        I3 | Peraturan Komisi Pemilihan Umum Nomor 01 Tahun
p02        I3 | 2013 tentang Pedoman Pelaksanaan Kampanye Pemilihan
p02        I3 | Umum Anggota Dewan Perwakilan Rakyat, Dewan
p02        I3 | Perwakilan Daerah, dan Dewan Perwakilan Rakyat
p02        I3 | Daerah, sebagaimana telah diubah dengan Peraturan
p02        I3 | Komisi Pemilihan Umum Nomor 15 Tahun 2013;
p02        I5 [KEPUTUSAN:MEMUTUSKAN]
p02        I5 | MEMUTUSKAN :
p02        I1 [PREAMBLE:MENETAPKAN]
p02        I1 | Menetapkan: PERATURAN
p02        I5 | KOMISI
p02        | PEMILIHAN
p02        | UMUM
p02        | TENTANG
p02        I2 | PERUBAHAN ATAS PERATURAN KOMISI PEMILIHAN UMUM
p02        I2 | NOMOR 17 TAHUN 2013 TENTANG PEDOMAN PELAPORAN
p02        I2 | DANA KAMPANYE PESERTA PEMILIHAN UMUM ANGGOTA
p02        I2 | DEWAN
p02        I4 | PERWAKILAN
p02        | RAKYAT,
p02        | DEWAN
p02        | PERWAKILAN
p02        I2 | DAERAH DAN DEWAN PERWAKILAN RAKYAT DAERAH.
p02        | Pasal I
p02        I2 | Beberapa ketentuan dalam Peraturan Komisi Pemilihan
p02        I2 | Umum Nomor 17 Tahun 2013 tentang Pedoman Pelaporan
p02        I2 | Dana Kampanye Peserta Pemilu Anggota Dewan Perwakilan
p02        I2 | Rakyat, Dewan Perwakilan Daerah, dan Dewan Perwakilan
p02        I2 | Rakyat Daerah, diubah sebagai berikut:
p02        I2 [ITEM]
p02        I2 | 1.
p02        I3 | Ketentuan Pasal 20 ayat (5) diubah, sehingga berbunyi
p02        I3 | sebagai berikut:
p02        | “Pasal 20
p02        I3 [AYAT]
p02        I3 | (1)
p02        I4 | Pengurus Partai Politik Peserta Pemilu sesuai dengan
p02        I4 | tingkatannya wajib menyampaikan laporan awal
p02        I4 | Dana Kampanye Partai Politik Peserta Pemilu kepada
p02        I4 | KPU, KPU Provinsi, dan KPU Kabupaten/Kota.
p02        I3 [AYAT]
p02        I3 | (2)
p02        I4 | Laporan awal Dana Kampanye Partai Politik Peserta
p02        I4 | Pemilu
p02        I6 | sebagaimana
p02        | dimaksud
p02        | pada
p02        | ayat
p02        [AYAT]
p02        | (1)
p02        | mencakup …
==================== PAGE 3 ====================
p03        I6 | - 3 -
p03        I4 | mencakup laporan awal Dana Kampanye para calon
p03        I4 | anggota
p03        I6 | DPR,
p03        | DPRD
p03        | provinsi
p03        | dan
p03        | DPRD
p03        I4 | kabupaten/kota.
p03        I3 [AYAT]
p03        I3 | (3)
p03        I4 | Laporan awal Dana Kampanye Partai Politik Peserta
p03        I4 | Pemilu sebagaimana dimaksud pada ayat (1) wajib
p03        I4 | dilampiri
p03        I6 | laporan
p03        | pencatatan
p03        | penerimaan
p03        | dan
p03        I4 | pengeluaran Dana Kampanye Calon Anggota DPR,
p03        I4 | DPRD provinsi dan DPRD kabupaten/kota.
p03        I3 [AYAT]
p03        I3 | (4)
p03        I4 | Calon Anggota DPD wajib menyampaikan laporan
p03        I4 | awal Dana Kampanye Calon Anggota DPD yang
p03        I4 | bersangkutan kepada KPU melalui KPU Provinsi.
p03        I3 [AYAT]
p03        I3 | (5)
p03        I4 | Laporan sebagaimana dimaksud pada ayat (1) dan
p03        I4 | ayat (4) disampaikan paling lambat 14 (empat belas)
p03        I4 | hari sebelum hari pertama jadual pelaksanaan
p03        I4 | Kampanye Pemilu dalam bentuk rapat umum.”
p03        I2 [ITEM]
p03        I2 | 2.
p03        I3 | Ketentuan Pasal 21 ayat (1), diubah, sehingga berbunyi
p03        I3 | sebagai berikut:
p03        | “Pasal 21
p03        I3 [AYAT]
p03        I3 | (1)
p03        I4 | Laporan
p03        I6 | awal
p03        | Dana
p03        | Kampanye
p03        | sebagaimana
p03        I4 | dimaksud dalam Pasal 20 ayat (1) dan ayat (4)
p03        I4 | mencakup:
p03        I4 [SUB-ITEM]
p03        I4 | a.
p03        I5 | informasi daftar penyumbang;
p03        I4 [SUB-ITEM]
p03        I4 | b.
p03        I5 | jumlah penerimaan dan pengeluaran Dana
p03        I5 | Kampanye berupa uang, barang dan/atau jasa
p03        I5 | setelah tanggal pembukaan rekening khusus
p03        I5 | sampai dengan paling lambat 14 (empat belas)
p03        I5 | hari sebelum hari pertama jadual pelaksanaan
p03        I5 | Kampanye Pemilu dalam bentuk rapat umum;
p03        I4 [SUB-ITEM]
p03        I4 | c.
p03        I5 | jumlah penerimaan dan pengeluaran Dana
p03        I5 | Kampanye
p03        | sebagaimana
p03        | tercatat
p03        | dalam
p03        I5 | Rekening Khusus Dana Kampanye dari bank
p03        I5 | sejak dibuka sampai dengan paling lambat 14
p03        I5 | (empat belas) hari sebelum hari pertama jadual
p03        I5 | pelaksanaan Kampanye Pemilu dalam bentuk
p03        I5 | rapat umum.
p03        I3 [AYAT]
p03        I3 | (2)
p03        I4 | Lingkup waktu laporan awal Dana Kampanye
p03        I4 | terhitung dari sejak pembukaan Rekening Khusus
p03        I4 | Dana Kampanye dan pembukuan penerimaan dan
p03        I4 | pengeluaran Dana Kampanye sampai dengan paling
p03        I4 | lambat 14 (empat belas) hari sebelum hari pertama
p03        I4 | jadual pelaksanaan Kampanye Pemilu dalam bentuk
p03        I4 | rapat umum.
p03        [AYAT]
p03        | (3) Laporan …
==================== PAGE 4 ====================
p04        I6 | - 4 -
p04        I3 [AYAT]
p04        I3 | (3)
p04        I4 | Laporan awal Dana Kampanye yang tidak mencakup
p04        I4 | semua informasi/data sebagaimana dimaksud pada
p04        I4 | ayat (1) dikembalikan kepada Peserta Pemilu.
p04        I3 [AYAT]
p04        I3 | (4)
p04        I4 | Peserta Pemilu wajib menyampaikan laporan hasil
p04        I4 | perbaikan kepada KPU, KPU Provinsi, dan KPU
p04        I4 | Kabupaten/Kota paling lambat 5 (lima) hari sejak
p04        I4 | diterima
p04        I6 | dari
p04        | KPU,
p04        | KPU
p04        | Provinsi
p04        | dan
p04        | KPU
p04        I4 | Kabupaten/Kota.
p04        I3 [AYAT]
p04        I3 | (5)
p04        I4 | Dalam hal Peserta Pemilu tidak menyampaikan
p04        I4 | laporan hasil perbaikan sebagaimana dimaksud
p04        I4 | pada ayat (4), KPU, KPU Provinsi dan KPU
p04        I4 | Kabupaten/Kota
p04        | mengumumkan
p04        | kepada
p04        I4 | masyarakat melalui papan pengumuman dan/atau
p04        I4 | website
p04        I6 | KPU,
p04        | KPU
p04        | Provinsi
p04        | dan
p04        | KPU
p04        I4 | Kabupaten/Kota  paling lambat 3 (tiga) hari setelah
p04        I4 | batas waktu Peserta Pemilu tidak menyampaikan
p04        I4 | laporan hasil perbaikan.”
p04        I2 [ITEM]
p04        I2 | 3.
p04        I3 | Ketentuan Pasal 22 ayat (4) diubah, sehingga berbunyi
p04        I3 | sebagai berikut:
p04        | “Pasal 22
p04        I3 [AYAT]
p04        I3 | (1)
p04        I4 | Pengurus Partai Politik Peserta Pemilu pada setiap
p04        I4 | tingkatan
p04        | wajib
p04        | melaporkan
p04        | sumbangan
p04        I4 | sebagaimana dimaksud dalam Pasal 19 kepada KPU,
p04        I4 | KPU Provinsi dan KPU Kabupaten/Kota.
p04        I3 [AYAT]
p04        I3 | (2)
p04        I4 | Calon Anggota DPD wajib melaporkan sumbangan
p04        I4 | sebagaimana dimaksud dalam Pasal 19 kepada KPU
p04        I4 | melalui KPU Provinsi.
p04        I3 [AYAT]
p04        I3 | (3)
p04        I4 | Laporan
p04        I6 | penerimaan
p04        | sumbangan
p04        | mencakup
p04        I4 | informasi sebagaimana dimaksud dalam Pasal 19
p04        I4 | ayat (2), ayat (3) dan ayat (4).
p04        I3 [AYAT]
p04        I3 | (4)
p04        I4 | Laporan
p04        I6 | penerimaan
p04        | sumbangan
p04        | sebagaimana
p04        I4 | dimaksud pada ayat (3) disampaikan secara periodik
p04        I4 | per Desember 2013 dan per Maret 2014.”
p04        I2 [ITEM]
p04        I2 | 4.
p04        I3 | Ketentuan Pasal 28 ayat (4) huruf d diubah, sehingga
p04        I3 | berbunyi sebagai berikut:
p04        | “Pasal 28
p04        I3 [AYAT]
p04        I3 | (1)
p04        I4 | KPU menunjuk Kantor Akuntan Publik untuk
p04        I4 | melakukan audit Dana Kampanye Peserta Pemilu.
p04        I3 [AYAT]
p04        I3 | (2)
p04        I4 | Pengadaan Kantor Akuntan Publik sebagaimana
p04        I4 | dimaksud pada ayat (1) dilakukan berdasarkan
p04        I4 | peraturan perundang-undangan.
p04        [AYAT]
p04        | (3) Biaya …
==================== PAGE 5 ====================
p05        I6 | - 5 -
p05        I3 [AYAT]
p05        I3 | (3)
p05        I4 | Biaya
p05        I6 | pengadaan
p05        | Kantor
p05        | Akuntan
p05        | Publik
p05        I4 | sebagaimana dimaksud pada ayat (1) dibebankan
p05        I4 | pada Anggaran Belanja dan Pendapatan Negara
p05        I4 [HEADING:BAGIAN]
p05        I4 | Bagian KPU dan KPU Provinsi.
p05        I3 [AYAT]
p05        I3 | (4)
p05        I4 | Kantor Akuntan Publik yang ditunjuk sebagaimana
p05        I4 | dimaksud pada ayat (2) wajib membuat pernyataan
p05        I4 | tertulis di atas kertas bermeterai cukup bahwa
p05        I4 | rekan yang bertanggung jawab atas pemeriksaan
p05        I4 | laporan Dana Kampanye:
p05        I4 [SUB-ITEM]
p05        I4 | a.
p05        I5 | tidak berafiliasi secara langsung ataupun tidak
p05        I5 | langsung dengan Partai Politik Peserta Pemilu
p05        I5 | dan Calon Anggota DPD;
p05        I4 [SUB-ITEM]
p05        I4 | b.
p05        I5 | bukan merupakan anggota atau pengurus
p05        I5 | Partai Politik Peserta Pemilu;
p05        I4 [SUB-ITEM]
p05        I4 | c.
p05        I5 | Akuntan Publik yang bertanggung jawab atas
p05        I5 | pemeriksaan laporan dana kampanye telah
p05        I5 | mengikuti pelatihan audit dana kampanye yang
p05        I5 | diselenggarakan oleh asosiasi profesi akuntan
p05        I5 | publik; dan
p05        I4 [SUB-ITEM]
p05        I4 | d.
p05        I5 | telah mendapatkan 1 (satu) surat rekomendasi
p05        I5 | dari asosiasi profesi akuntan publik yang
p05        I5 | dijadikan sebagai nilai tambah dalam proses
p05        I5 | pengadaan jasa audit Partai Politik dan Calon
p05        I5 | Anggota DPD.”
p05        I2 [ITEM]
p05        I2 | 5.
p05        I3 | Ketentuan Pasal 37 diubah, sehingga berbunyi sebagai
p05        I3 | berikut:
p05        | ”Pasal 37
p05        I3 | Dalam hal pengurus Partai Politik Peserta Pemilu pada
p05        I3 | setiap
p05        I4 | tingkatan
p05        | dan
p05        | Calon
p05        | Anggota
p05        | DPD
p05        | tidak
p05        I3 | menyampaikan laporan awal Dana Kampanye kepada
p05        I3 | KPU/KPU Provinsi/ KPU Kabupaten/Kota sampai batas
p05        I3 | waktu sebagaimana dimaksud dalam Pasal 20 ayat (5),
p05        I3 | Partai Politik dan Calon Anggota DPD yang bersangkutan
p05        I3 | dikenai sanksi sebagaimana diatur dalam Peraturan
p05        I3 | Perundang-undangan tentang Pemilu Anggota DPR, DPD
p05        I3 | dan DPRD.”
p05        I2 [ITEM]
p05        I2 | 6.
p05        I3 | Ketentuan Pasal 38 diubah, sehingga berbunyi sebagai
p05        I3 | berikut:
p05        | “Pasal 38
p05        I3 | Dalam hal pengurus Partai Politik Peserta Pemilu pada
p05        I3 | setiap
p05        I4 | tingkatan
p05        | dan
p05        | Calon
p05        | Anggota
p05        | DPD
p05        | tidak
p05        I3 | menyampaikan laporan penerimaan dan pengeluaran
p05        | Dana …
```

---


## peraturan

- **File**: `peraturan/PP0201962.pdf`
- **Document Type**: Peraturan Pemerintah (simple)
- **Issued by**: Presiden
- **Pages**: 2 | **Lines**: 87
- **Font sizes**: [10.0, 13.0]
- **Most common font**: 13.0 (95% of lines)
- **Bold font sizes**: None
- **Indent clusters**: [85.0, 168.0, 210.0, 244.0, 345.0, 374.0, 417.0, 446.0, 492.0, 518.0]
- **Expected hierarchy**: Pasal > Ayat

### Full Text Skeleton

```python
==================== PAGE 1 ====================
p01   F10  | PRESIDEN
p01   F10  | REPUBLIK INDONESIA
p01        I2 | PERATURAN PEMERINTAH REPUBLIK INDONESIA
p01        I4 | NOMOR 20 TAHUN 1962
p01        | TENTANG
p01        I3 | LAFAL SUMPAH JANJI APOTEKER
p01        I4 | Presiden Republik Indonesia,
p01        I1 [PREAMBLE:MENIMBANG]
p01        I1 | Menimbang
p01        I2 | :
p01        I2 | perlu menetapkan lafal sumpah/janji apoteker: teker;
p01        I1 [PREAMBLE:MENGINGAT]
p01        I1 | Mengingat
p01        I2 | :
p01        I2 [ITEM]
p01        I2 | 1.
p01        I3 | pasal 5 ayat 2 Undang-undang Dasar:
p01        I2 [ITEM]
p01        I2 | 2.
p01        I3 | pasal
p01        I4 | 10 ayat (3) Undang-undang No. 9 tahun 1960 tentang
p01        I3 | Pokok-pokok Kesehatan (Lembaran-Negara tahun 1960 No.131);
p01        I1 | Mendengar
p01        I2 | :
p01        I2 | Menteri Pertama , Wakil Menteri Pertama Bidang Kesejahteraan
p01        I2 | Rakyat, Menteri Kesehatan dan Menteri Kehakiman :
p01        [KEPUTUSAN:MEMUTUSKAN]
p01        | Memutuskan :
p01        I1 [PREAMBLE:MENETAPKAN]
p01        I1 | Menetapkan
p01        I2 | :
p01        I2 | Peraturan Pemerintah tentang lafal sumpah/janji apoteker.
p01        | Pasal 1.
p01        I2 [AYAT]
p01        I2 | (1) Sebelum seorang Apoteker melakukan jabatannya, maka ia harus
p01        I2 | mengucapkan sumpah menurut cara agama yang dipeluknya, atau
p01        I2 | mengucapkan janji, Ucapan sumpah dimulai dengan kata-kata "Demi
p01        I2 | Allah" bagi mereka yang beragama Islam, dan sumpah untuk agama
p01        I2 | lain,
p01        I3 | pemakaian
p01        | kata-kata
p01        I5 | "Demi
p01        I7 | Allah"
p01        I8 | disesuaikan
p01        I10 | dengan
p01        I2 | kebiasaan agama masing-masing.
p01        I2 [AYAT]
p01        I2 | (2) Sumpah/janji itu berbunyi sebagai berikut:
p01        I2 [ITEM]
p01        I2 | 1. Saya
p01        I4 | akan
p01        | membaktikan
p01        I6 | hidup
p01        I7 | saya
p01        I8 | guna
p01        I9 | kepentingan
p01        I3 | perikemanusiaan, terutama dalam bidang kesehatan:
p01        I2 [ITEM]
p01        I2 | 2. Saya akan merahasiakan segala sesuatu yang saya ketahui karena
p01        I3 | pekerjaan saya dan keilmuan saya sebagai apoteker;
p01        I8 [ITEM]
p01        I8 | 3. Sekalipun …
==================== PAGE 2 ====================
p02   F10  | PRESIDEN
p02   F10  | REPUBLIK INDONESIA
p02        | - 2 -
p02        I2 [ITEM]
p02        I2 | 3. Sekalipun diancam,saya tidak akan mempergunakan pengetahuan
p02        I3 | kefarmasian saya untuk sesuatu yang bertentangan dengan hukum
p02        I3 | perikemanusiaan;
p02        I2 [ITEM]
p02        I2 | 4. Saya akan menjalankan tugas saya dengan sebaik-baiknya sesuai
p02        I3 | dengan martabat dan tradisi luhur jabatan kefar masian:
p02        I2 [ITEM]
p02        I2 | 5. Dalam menunaikan kewajiban saya, saya akan berihtiar dengan
p02        I3 | sungguh-sungguh supaya tidak terpengaruh oleh pertimbangan
p02        I3 | Keagamaan, Kebangsaan, Kesukuan, Politik,Kepartaian atau
p02        I3 | Kedudukan Sosial:
p02        I2 [ITEM]
p02        I2 | 6. Saya ikrarkan sumpah/janji ini dengan sungguh-sungguh dan
p02        I3 | dengan penuh keinsyafan.
p02        | Pasal 2.
p02        I2 | Peraturan Pemerintah ini mulai berlaku pada hari diundangkannya.
p02        I2 | Agar
p02        I3 | supaya
p02        I4 | setiap
p02        | orang
p02        I5 | dapat
p02        I6 | mengetahunya
p02        I9 | memerintahkan
p02        I2 | pengundangan Peraturan Pemerintah ini dengan penempatan dalam
p02        I2 | Lembaran-Negara Republik Indonesia.
p02        | Ditetapkan di Jakarta.
p02        | pada tanggal 20 September 1962.
p02        | Presiden Republik Indonesia.
p02        | ttd
p02        | SUKARNO.
p02        I1 | Diundangkan di Jakarta
p02        I1 | pada tanggal 20 September 1962.
p02        I1 | Sekretaris Negara,
p02        I1 | ttd
p02        I1 | MOHD. ICHSAN.
p02        I2 | LEMBARAN NEGARA TAHUN 1962 NOMOR 69
```

---

