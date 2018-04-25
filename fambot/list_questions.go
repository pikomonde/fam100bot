package fambot

const listofquestions = `
Sebutkan benda apa saja yang ada di dalam kamar tidur|TEMPAT TIDUR~KASUR~SPRING BED|LEMARI~LEMARI BAJU~LEMARI PAKAIAN|CERMIN~KACA|MEJA~MEJA BELAJAR~MEJA RIAS|TELEVISI~TV~TIVI;
Sebutkan benda apa saja yang ada di dalam kamar mandi|BAK~BAK MANDI~BAK AIR|SABUN|GAYUNG|SAMPO~SHAMPO~SHAMPOO|HANDUK~ANDUK;
Sebutkan hewan yang memiliki kaki 2|AYAM|BURUNG|BEBEK|ANGSA|PENGUIN~PINGUIN;
Sebutkan profesi yang menggunakan peluit|WASIT|TUKANG PARKIR~PAK OGAH|POLISI|SATPAM~SECURITY~HANSIP|PRAMUKA;
Sebutkan benda-benda yang dibawa oleh siswa saat pergi ke sekolah|BUKU~BUKU TULIS|TAS~RANSEL~TAS RANSEL|PULPEN~BOLPOIN~PENA|PENSIL|PENGGARIS;
Sebutkan nama-nama sosial media yang digunakan oleh banyak orang|FACEBOOK~FB|TWITTER|INSTAGRAM~IG|PATH|GOOGLE PLUS~GOOGLE+~GOOGLE +;
Sebutkan kendaraan yang sering digunakan oleh masyarakat Indonesia|MOTOR~SEPEDA MOTOR|MOBIL|BUS~BIS|KERETA~KRL~KERETA API|DELMAN~ANDONG;
Sebutkan bagian-bagian dari komputer|MONITOR|KEYBOARD|MOUSE|CPU|PRINTER;
Sebutkan pelajaran yang tidak disukai di sekolah|MATEMATIKA|FISIKA|BAHASA INGGRIS~INGGRIS|SEJARAH|BIOLOGI;
Sebutkan nama-nama makanan Indonesia yang menggunakan bahan daging sapi|BAKSO~BASO|RENDANG|SEMUR|DENDENG|ABON;
Sebutkan benda-benda yang dicari ketika mati lampu|LILIN|KOREK~KOREK API|SENTER|HANDPHONE~HP~SMARTPHONE|GENSET~JENSET;
Sebutkan benda-benda elektronik yang menggeluarkan suara|RADIO|HANDPHONE~HP~SMARTPHONE|TELEVISI~TV~TIVI|KOMPUTER~PC~LAPTOP|MP3 PLAYER~WALKMAN;
Sebutkan jenis tagihan atau pembayaran yang biasa dibayar per bulan|LISTRIK~PLN|AIR~PDAM|KREDIT~KARTU KREDIT~CREDIT CARD|BPJS~JAMSOSTEK|ASURANSI;
Sebutkan benda-benda yang dapat menampung air|GELAS~CANGKIR~MUG|EMBER|BOTOL|BAK~BAK MANDI~BAK AIR|GENTONG~DRUM~TONG;
Sebutkan hewan apa saja yang memakan rumput|KAMBING|SAPI|KUDA|RUSA~KIJANG|KERBAU~KEBO;
Sebutkan nama-nama negara yang menggunakan bahasa Inggris dalam kesehariannya|INGGRIS|AMERIKA~AMERIKA SERIKAT~USA|SINGAPURA~SINGAPORE~SINGAPUR|AUSTRALIA|IRLANDIA;
Sebutkan hewan yang sering dipelihara oleh manusia|KUCING|ANJING|BURUNG|IKAN|AYAM;
Sebutkan negara-negara yang terkenal akan sepakbolanya|BRAZIL|ARGENTINA|JERMAN|ITALI|INGGRIS;
Sebutkan aliran musik yang terpopuler di dunia|POP|ROCK|JAZZ|METAL|HIP HOP~HIP-HOP;
Sebutkan nama-nama ikan yang sering dikonsumsi oleh masyarakat Indonesia|TONGKOL|TERI|MUJAER~MUJAIR|LELE|TENGGIRI;
Sebutkan jenis pakaian yang digunakan di badan|KAOS~T-SHIRT~T SHIRT|KEMEJA|JAKET|SWEATER|JAS;
Sebutkan benda-benda yang digunakan di kepala|TOPI|HELM|KUPLUK|BANDO|MAHKOTA;
Sebutkan nama buah yang berwarna kuning|PISANG|JERUK|NANAS|LEMON|PIR~PEAR;
Sebutkan benda-bendabenda yang berputar|KIPAS ANGIN|RODA~BAN|JAM~JAM DINDING|BLENDER|STIR~STIR MOBIL;
Sebutkan hewan-hewan yang memiliki racun|ULAR~COBRA~KOBRA|LEBAH~TAWON|KALAJENGKING|UBUR-UBUR|IKAN BUNTAL~IKAN FUGU~FUGU;
Sebutkan lomba-lomba yang ada di 17 Agustus-an|PANJAT PINANG|BALAP KARUNG~BALAPAN KARUNG|TARIK TAMBANG|MAKAN KERUPUK~MAKAN KRUPUK|GIGIT KOIN;
Sebutkan nama binatang yang suka menggigit|NYAMUK|ULAR|ANJING|TIKUS|KUCING;
Sebutkan nama club sepakbola dari Liga Inggris|MANCHESTER UNITED~MU|LIVERPOOL|ARSENAL|CHELSEA|MANCHESTER CITY;
Sebutkan nama-nama olahraga yang dimainkan secara tim|SEPAK BOLA~SEPAKBOLA|BASKET|VOLI~VOLLY|HOKI|BASEBALL~KASTI;
Sebutkan permainan apa saja yang ada di Dunia Fantasi (Dufan)|HALILINTAR|KORA-KORA~KORA KORA|TORNADO|ARUNG JERAM|KOMEDI PUTAR;
Sebutkan nama mata uang negara yang ada di dunia|DOLAR~DOLLAR|RUPIAH|POUND STERLING~POUNDSTERLING|RUPEE~RUPE|PESO;
Sebutkan nama anggota tubuh yang berada di kepada|MATA|HIDUNG|TELINGA~KUPING|MULUT~BIBIR|PIPI;
Sebutkan nama senjata tradisional dari Indonesia|KERIS|BADIK|KUJANG|MANDAU|RENCONG;
Sebutkan nama-nama dari jenis-jenis burung|MERPATI|KAKATUA~KAKA TUA|ELANG~RAJAWALI|BEO|CENDRAWASIH~CENDRAWASI;
Sebutkan profesi yang mensyarakatkan tentang tinggi badan|PRAMUGARI|MODEL|TENTARA~TNI~ABRI|POLISI|PILOT;
Sebutkan negara-negara yang ada di benua ASIA|INDONESIA|JEPANG|CHINA~RRC|SINGAPURA~SINGAPUR|MALAYSIA;
Sebutkan film animasi atau kartun dari Jepang yang populer di Indonesia|NARUTO|DORAEMON|ONE PIECE|DRAGON BALL|SHINCHAN~SINCAN~SINCHAN;
Sebutkan aliran beladiri yang ada di dunia|KARATE|PENCAK SILAT~SILAT|KUNGFU~KUNG FU~KUNG-FU|TAEKWONDO~TEKONDO~TEKWONDO|JUDO;
Sebutkan nama-nama benua|ASIA|EROPA|AMERIKA|AUSTRALIA|AFRIKA;
Sebutkan benda-benda yang sering lupa dibawa|HANDPHONE~HP~SMARTPHONE|KUNCI|DOMPET|KACAMATA~KACA MATA|ROKOK;
Sebutkan tempat-tempat yang sering digunakan untuk pacaran|BIOSKOP|CAFE~KAFE|MALL|TAMAN~TAMAN BERMAIN|PANTAI;
Sebutkan hal-hal yang terkenal dari negara Mesir|PIRAMIDA~PIRAMID|FIRAUN|PADANG PASIR~GURUN~GURUN PASIR|SUNGAI NIL~NIL|SPHINX;
Sebutkan nama binatang yang akan membuat cewek berteriak|KECOA~KECOAK|TIKUS|CACING|CICAK|BELUT;
Sebutkan nama stasiun tv (channel) yang ada di Indonesia|RCTI|SCTV|TRANS TV~TRANS~TRANSTV|TV ONE~TVONE|INDOSIAR;
Sebutkan nama alat musik yang memiliki senar|GITAR|BIOLA|KECAPI|HARPA|BASS;
Sebutkan benda-benda yang ada di dapur|KOMPOR|PANCI|KUALI~PENGGORENGAN~WAJAN|PISAU|TALENAN~TELENAN;
Sebutkan nama-nama gaya berenang|KATAK|BEBAS|KUPU-KUPU|PUNGGUNG|DADA;
Sebutkan kegiatan yang biasa diperintahkan oleh orangtua kepada anaknya|BELAJAR|TIDUR|MAKAN|MANDI|PULANG;
Sebutkan nama binatang yang biasa digunakan untuk menghina orang|ANJING|MONYET|BABI|BANGSAT|KAMPRET;
Sebutkan nama minuman yang biasa diberi gula|KOPI|TEH|SUSU|CENDOL|DAWET;
Sebutkan nama-nama kota di Indonesia yang berawalan \"M\"|MEDAN|MAKASSAR~MAKASAR|MALANG|MANADO|MATARAM;
Sebutkan tipe cewek yang disukai oleh cowok|CANTIK|SEXY~SEXI~SEKSI|PINTAR~PANDAI|PUTIH~BERSIH~BENING|KAYA~TAJIR;
Sebutkan benda-benda yang hanya digunakan oleh cewek|LIPSTIK~LIPSTICK|BEDAK|PEMBALUT|MASKARA|PENSIL ALIS~PINSIL ALIS;
Sebutkan nama hari-hari besar keagamaan yang ada di Indonesia|IDUL FITRI~LEBARAN|NATAL|IDUL ADHA~IDUL KORBAN~IDUL KURBAN|WAISAK|NYEPI;
Sebutkan cita-cita yang sering diucapkan oleh anak SD|DOKTER|TENTARA|GURU|POLISI|PRESIDEN;
Sebutkan nama-nama negara di dunia yang berawalan \"B\"|BELANDA|BRAZIL|BOLIVIA|BELGIA|BRUNEI DARUSSALAM~BRUNEI;
Sebutkan nama-nama merk motor yang ada di MotoGP|YAMAHA|HONDA|DUCATI|SUZUKI|KAWASAKI;
Sebutkan nama-nama penyakit berbahaya dan mematikan|JANTUNG|KANKER|AIDS~HIV~HIV AIDS|STROKE|MALARIA;
Sebutkan nama-nama minimarket yang ada di Indonesia|INDOMARET|ALFAMART|ALFAMIDI|7-ELEVEN~SEVEL~7 ELEVEN|CIRCLE K~CIRCLE-K;
Sebutkan nama-nama hewan yang ditunggangi manusia|KUDA|UNTA~ONTA|GAJAH|KELEDAI|BANTENG;
Sebutkan benda-benda yang dibuat dari emas|CINCIN|KALUNG|ANTING|GELANG|MAHKOTA;
Sebutkan nama-nama tokoh Pandawa Lima|ARJUNA|BIMA|NAKULA|SADEWA|YUDISTIRA~YUDHISTIRA;
Sebutkan jenis-jenis kain|SUTRA~SUTERA|KATUN|SATIN|WOL~WOOLS~WOLS|RAYON;
Sebutkan nama-nama bencana alam|GEMPA BUMI~GEMPA|BANJIR|GUNUNG MELETUS|BADAI|TORNADO~ANGIN TORNADO;
Sebutkan nama-nama Selat yang ada di Indonesia|SUNDA~SELAT SUNDA|MALAKA~SELAT MALAKA|MADURA~SELAT MADURA|LOMBOK~SELAT LOMBOK|BALI~SELAT BALI;
Sebutkan akibat yang ditimbulkan setelah putus cinta|GALAU|MENANGIS~SEDIH~NANGIS|STRESS~STRES~DEPRESI|BUNUH DIRI|MOVE ON~CARI PACAR LAGI;
Sebutkan jenis rasa yang dapat dirasakan oleh lidah manusia|MANIS|PAHIT|ASIN|ASAM~ASEM|UMAMI~GURIH;
Sebutkan alat-alat pernafasan pada mahluk hidup|PARU-PARU|INSANG|KULIT~PORI-PORI|TRAKEA|PUNDI UDARA;
Sebutkan nama tokoh-tokoh yang ada di serial animasi Naruto|NARUTO|SASUKE|SAKURA|KAKASHI|GAARA;
Sebutkan nama-nama telur yang biasa dimakan oleh masyarakat Indonesia|AYAM~TELUR AYAM|BEBEK~TELUR BEBEK|PUYUH~TELUR PUYUH|ANGSA~TELUR ANGSA|KAVIAR~TELUR IKAN~IKAN;
Sebutkan nama-nama binatang yang paling dibenci manusia|NYAMUK|KECOA~KECOAK|TIKUS|LALAT|LINTAH;
Sebutkan rumus-rumus matematika yang diajarkan di sekolah dasar|TAMBAH|KURANG|KALI~PERKALIAN|BAGI~PEMBAGIAN|PANGKAT;
Apa saja sebutan yang biasa digunakan masyarakat Indonesia untuk memanggil ayah|PAPA|BAPAK~BAPA|BOKAP|PAPI|BABE~BABEH;
Sebutkan negara yang memiliki jumlah penduduk terpadat di dunia|INDONESIA|CHINA~CINA~RRC|INDIA|AMERIKA~AMERIKA SERIKAT~USA|BRAZIL;
Sebutkan nama-nama browser terpopuler|MOZILLA~MOZILLA FIREFOX~FIREFOX|CHROME~GOOGLE CHROME|OPERA~OPERA MINI|SAFARI|INTERNET EXPLORER~IE;
Sebutkan nama-nama sepatu pria yang terpopuler di dunia|NIKE|ADIDAS|JORDAN|CONVERSE|REEBOK;
Sebutkan benda-benda apa saja yang dimasukkan ke dalam dompet|UANG~DUIT|KTP|SIM|KARTU ATM~ATM|KARTU NAMA;
Sebutkan hewan bertanduk|KIJANG|KAMBING|BANTENG|KERBAU~KEBO|BISON~BYSON;
Sebutkan nama hewan yang suaranya sering ditiru oleh manusia|KUCING|ANJING|KAMBING|AYAM|TOKEK;
Sebutkan nama-nama klub sepakbola yang bertanding di Liga Seri A|JUVENTUS|AC MILAN|INTER MILAN~INTERMILAN|AS ROMA|LAZIO;
Sebutkan tempat wisata Indonesia terpopuler|BALI~PULAU BALI|RAJA AMPAT|BUNAKEN|WAKATOBI|PULAU LOMBOK~LOMBOK;
Sebutkan rukun Islam|SYAHADAT~SAHADAT|SHOLAT~SOLAT|ZAKAT|PUASA|HAJI;
Sebutkan group band Indonesia yang terpopuler|SLANK|NOAH~PETERPAN~PETER PAN|NIDJI|DEWA~DEWA 19|KANGEN BAND~KANGEN;
Sebutkan maskapai penerbangan yang ada di Indonesia|GARUDA~GARUDA INDONESIA|LION AIR|CITILINK|AIRASIA~AIR ASIA|BATIK AIR;
Sebutkan penjahat atau musuh superhero yang terpopuler|JOKER|LOKI|MAGNETO|VENOM|ULTRON;
Sebutkan bahan-bahan untuk membuat roti|TEPUNG~TERIGU~TEPUNG TERIGU|TELUR~TELOR|GULA|SUSU|MENTEGA;
Sebutkan anggota tubuh yang diciptakan sepasang|TANGAN|KAKI|MATA|TELINGA~KUPING|PUTING;
Sebutkan negara-negara yang ada di benua Eropa|INGGRIS|PERANCIS~PRANCIS|ITALI~ITALIA|JERMAN|RUSIA;
Sebukan mamalia yang hidup di air laut|PAUS~IKAN PAUS|LUMBA-LUMBA|ANJING LAUT|SINGA LAUT|IKAN PESUT~PESUT~DUGONG;
Sebutkan nama-nama hantu yang ada di Indonesia|KUNTILANAK|POCONG|GENDERUWO~GENDRUWO|TUYUL|WEWE GOMBEL;
Sebutkan nama media cetak atau koran yang ada di Indonesia|KOMPAS~KORAN KOMPAS|REPUBLIKA|JAWA POS|LAMPU MERAH|MEDIA INDONESIA;
Sebutkan unsur-unsur dalam kandungan udara|OKSIGEN|NITROGEN|HIDROGEN|KARBON DIOKSIDA~KARBONDIOKSIDA|HELIUM;
Sebutkan barang atau benda apa saja yang biasanya dibeli secara kredit|MOBIL|MOTOR~SEPEDA MOTOR|RUMAH~APARTEMENT~APARTEMEN|HANDPHONE~SMARTPHONE~HP|GADGET;
Sebutkan tokoh-tokoh yang ada di film kartun Doraemon|DORAEMON|NOBITA|SIZUKA|GIANT~TAKESHI|SUNEO;
Sebutkan merk minuman bersoda yang dijual di Indonesia|COCA COLA~COCA|SPRITE|FANTA|PEPSI|TEBS;
Sebutkan merk bolpoin atau pulpen yang dijual di Indonesia|PILOT|SNOWMAN|FABER CASTELL|FASTER|BOXY;
Sebutkan apa yang ada pada binatang ayam|TELUR|JENGGER|BULU|CEKER|DAGING;
Sebutkan bagian-bagian dari bunga|TANGKAI|KELOPAK|SERBUK SARI|PUTIK|MAHKOTA;
Sebutkan buah apa yang dikupas terlebih dahulu sebelum dimakan|JERUK|MANGGA|DURIAN~DUREN|PISANG|MANGGIS;
Sebutkan nama pangkat dalam dunia militer|JENDERAL~JENDRAL|KOPRAL|MAYOR|LETNAN|SERSAN;
Sebutkan buah apa saja yang mengandung vitamin C|JERUK|JAMBU~JAMBU BIJI|MANGGA|TOMAT|APEL;
Sebutkan alat musik tiup yang kamu ketahui|TEROMPET|SULING~SERULING|HARMONIKA|KLARINET~CLARINET|PIANIKA;
Sebutkan jenis-jenis Tepung yang biasa digunakan|TERIGU~TEPUNG TERIGU|TAPIOKA~TEPUNG TAPIOKA|SAGU~TEPUNG SAGU|BERAS~TEPUNG BERAS|MAIZENA~TEPUNG MAIZENA;
Sebutkan nama-nama gunung yang ada di Indonesia|KRAKATAU|BROMO|MERAPI|SALAK|SEMERU;
Sebutkan hewan-hewab buas bertaring panjang|SINGA|HARIMAU|ULAR|SERIGALA~SRIGALA|MACAN;
Sebutkan tari-tarian tradisional yang ada di Indonesia|KECAK~TARI KECAK|JAIPONG~TARI JAIPONG|PIRING~TARI PIRING|TOPENG~TARI TOPENG|SAMAN~TARI SAMAN;
Sebutkan istilah yang ada dalam olahraga sepak bola|GOL~GOAL|OFFSIDE|PINALTI|KIPER|GAWANG;
Sebutkan nama-nama cabang olahraga yang menggunakan pemukul|BASEBALL|BULU TANGKIS~BADMINTON|TENIS|GOLF|KASTI;
Sebutkan suku-suku yang ada di Indonesia|JAWA|BATAK|BETAWI|BUGIS|DAYAK;
Sebutkan anggota tubuh panca indera yang ada pada manusia|MATA|HIDUNG|TELINGA~KUPING|LIDAH|KULIT;
Sebutkan tokoh-tokoh yang ada dalam video game Mortal Kombat|SCORPION|SUB ZERO~SUB-ZERO|LIU KANG|SONYA|KITANA;
Sebutkan nama-nama website yang populer yang dikunjungi setiap hari|FACEBOOK|GOOGLE|YOUTUBE|KASKUS|DETIK~DETIK.COM;
Sebutkan nama-nama olahraga yang menggunakan bola besar|SEPAK BOLA~SEPAKBOLA|BASKET|VOLI~VOLLI~VOLLY|BOWLING|DODGEBALL;
Sebutkan hewan bercangkang|SIPUT~KEONG~BEKICOT|KURA-KURA|PENYU|KUMANG~KELOMANG|KERANG;
Sebutkan angka-angka yang sering digunakan dalam tulisan alay untuk mengganti huruf|4~EMPAT|1~SATU|7~TUJUH|5~LIMA|0~NOL;
Sebutkan apa saja yang biasa kamu lihat di langit|MATAHARI|BULAN|BINTANG|AWAN|PESAWAT~PESAWAT TERBANG;
Sebutkan apa yang biasanya dilakukan jika pergi ke mall|BELANJA~SHOPPING|JALAN-JALAN~MUTER-MUTER|NONGKRONG|PACARAN|JAJAN~MAKAN;
Sebutkan merk minuman gelas yang dijual di warung|AQUA|TEH GELAS|FRUTANG|ALE-ALE|OKKY JELLY~OKKY JELLY DRINK;
Sebutkan nama-nama anggota personel grup band SLANK saat ini|BIMBIM~BIM BIM~BIM-BIM|KAKA|ABDEE~ABDEE NEGARA~ABDEE NEGARA NURDIN|RIDHO|IVANKA;
Sebutkan benda-benda yang suka pecah|GELAS|TELUR~TELOR|KACA~CERMIN|LAMPU|GENTENG~GENTING;
Sebutkan benda-benda yang bisa terbang|PESAWAT~PESAWAT TERBANG|LAYANGAN~LAYANG-LAYANG|HELIKOPTER|UFO~PIRING TERBANG|DRONE;
Sebutkan peralatan masak yang terbuat dari besi/aluminium/stainless steel|PANCI|PENGGORENGAN~KUALI|PISAU|SENDOK|GARPU;
Sebutkan negara-negara yang ada di Timur Tengah|ARAB SAUDI|PALESTINA|MESIR|QATAR|IRAK;
Sebutkan nama buah yang kulitnya berwarna hijau|MANGGA|SEMANGKA|ALPUKAT|SIRSAK|KEDONDONG;
Sebutkan nama tokoh-tokoh yang ada dalam film animasi Dragon Ball|SON GOKU~GOKU~SONGOKU|BEJITA~VEGETA~BEZITA|SON GOHAN~GOHAN~SONGOHAN|PIKOLO~PICOLO~PICOLLO|JIN KURA-KURA;
Sebutkan apa saja yang biasanya didapat setelah memenangi lomba|PIALA|UANG|MENDALI|PIAGAM|MAHKOTA;
Sebutkan sesuatu hal yang membuat pria menjadi ilfil kepada wanita cantik|BAU BADAN~BAU KETEK|BAU MULUT~BAU JIGGONG|KENTUT~BUANG ANGIN~BUANG GAS|SENDAWA~BERSENDAWA|DANDAN MENOR~MENOR~MAKEUP MENOR;
Sebutkan hewan apa yang jalannya lambat atau pelan|KURA-KURA|SIPUT~KEONG~BEKICOT|KOALA|BINTANG LAUT|DUGONG~IKAN DUYUNG;
Sebutkan benda apa saja yang dipakai ketika udara sedang dingin|SELIMUT|JAKET~SWEATER|KUPLUK|SARUNG|SYAL;
Sebutkan jenis-jenis atau contoh-contoh dari tanaman umbi-umbian|UBI|SINGKONG~KETELA~KETELA POHON|KENTANG|LOBAK|WORTEL;
Sebutkan apa saja yang dipamerkan oleh cewek ke teman-teman ceweknya|PACAR|TAS|BAJU|SEPATU|RAMBUT;
Sebutkan benda-benda yang digunakan saat sedang hujan|PAYUNG|JAS HUJAN|TOPI|KANTONG KRESEK~KRESEK~KANTONG PLASTIK|JAKET~TUDUNG;
Sebutkan apa saja bentuk karya sastra yang populer|PUISI~SAJAK|NOVEL|CERPEN|PANTUN|PROSA;
Sebutkan apa saja yang digunakan untuk tidur|KASUR~TEMPAT TIDUR~SPRINGBED|BANTAL|GULING|SELIMUT|PIYAMA;
Sebutkan nama penemu dunia yang terpopuler|THOMAS EDISON~THOMAS ALVA EDISON~ALVA EDISON|ALBERT EINSTEIN|ALEXANDER GRAHAM BELL~GRAHAM BELL|JAMES WATT|ISAAC NEWTON;
Sebutkan nama-nama personel JKT48|NABILAH|MELODY|VE|GABY|RACHEL;
Sebutkan game android yang booming|ANGRY BIRD|FLAPPY BIRD|CLASH OF CLAN~COC|POKEMON GO|TAHU BULAT;
Sebutkan negara-negara yang ikut dalam perang dunia ke-2|AMERIKA SERIKAT|JEPANG|JERMAN|INGGRIS|BELANDA;
Sebutkan nama-nama stasiun yang ada di daerah Jabodetabek|GAMBIR~STASIUN GAMBIR|MANGGARAI~STASIUN MANGGARAI|BOGOR~STASIUN BOGOR|KOTA~STASIUN KOTA|TANAHABANG~TANAH ABANG~STASIUN TANAH ABANG;
Sebutkan anggota tubuh yang biasanya berwarna hitam|RAMBUT|JENGGOT|KUMIS|TAHI LALAT|BULU HIDUNG;
Sebutkan merk obat nyamuk yang ada di Indonesia|BAYGON|HIT|VAPE|TIGA RODA|MORTEN;
Sebutkan nama-nama lembaga tinggi negara selain Presiden yang ada di Indonesia|DPR|MPR|BPK|Mahkamah Agung~MA|DPD;
Sebutkan pembalap MotoGP saat ini|VALENTINO ROSSI~ROSSI|JORGE LORENZO~LORENZO|MARC MARQUEZ~MARQUEZ|DANI PEDROSA~PEDROSA|ANDREA DOVIZIOSO~DOVIZIOSO;
Sebutkan nama-nama pelawak atau komedian Indonesia yang telah wafat|DONO|KASINO|BENYAMIN S~BENYAMIN~BENYAMIN SUEB|JOJON|BASUKI;
Sebutkan yang biasa kita lakukan ketika mencium bau kentut|TUTUP HIDUNG~MENUTUP HIDUNG|KABUR~PERGI~MENJAUH|TAHAN NAPAS~TAHAN NAFAS~NAHAN NAPAS|MUNTAH~MUAL~MUNTAH-MUNTAH|DIAM;
Sebutkan nama artis yang setiap hari ada di televisi|RAFFI AHMAD|UYA KUYA|JESSICA ISKANDAR~JESIKA ISKANDAR|AYU TING-TING|GILANG~GILANG BHASKARA;
Sebutkan negara monarki yang ada di dunia|ARAB SAUDI|INGGRIS|JEPANG|MALAYSIA|BRUNEI DARUSALLAM;
Sebutkan tokoh-tokoh yang ada dalam cerita Si Unyil selain Unyilnya sendiri|PAK OGAH|PAK RADEN|UCRIT|USRO|CUPLIS;
Sebutkan benda-benda yang mengeluarkan mengeluarkan panas|KOMPOR|RICE COOKER~MAGIC JAR~MAGIC COM|SETRIKA~STRIKAAN~SETRIKAAN|SOLDIER|KNALPOT;
Sebutkan mengapa orang memakai masker|ASAP|BAU|GANTENG~CANTIK|GIGINYA GONDRONG|BULU HIDUNG KELUAR-KELUAR;
Sebutkan merk-merk helm yang ada di Indonesia|KYT|INK|BMC|INDEX|AGV;
Sebutkan kota-kota di Indonesia yang diawali huruf \"B\"|BANDUNG|BOGOR|BEKASI|BALIKPAPAN~BALIK PAPAN|BANDA ACEH;
Sebutkan jenis-jenis narkoba|GANJA|SABU-SABU~SABU|EKSTASI|HEROIN|PUTAU~PUTAW;
Sebutkan tempat dimana orang biasanya menyimpan atau menaruh uang|DOMPER|KANTONG~KANTUNG|BANK|BRANKAS~BRANGKAS|BRA~DALEMAN~BH;
Sebutkan bahan-bahan apa saja yang biasanya ditaruh di atas pizza|DAGING|SOSIS|JAMUR|BROKOLI|NANAS;
Sebutkan jenis-jenis busana muslim|HIJAB~JILBAB~KERUDUNG|SARUNG|GAMIS|KOKO~BAJU KOKO|MUKENA;
Sebutkan nama-nama laki-laki yang terlalu sering digunakan di Indonesia|BUDI|AGUNG|ANDI|ANDRE|AGUS;
Sebutkan nama binatang yang jalannya melompat-lompat|KATAK~KODOK|KANGGURU|KELINCI|BELALANG|KUTU;
Sebutkan nama-nama jenis pakaian yang hanya cocok dipakai oleh wanita|ROK~ROK MINI|HIJAB~JILBAB|DASTER|GAUN|BRA~KUTANG~BH;
Sebutkan apa saja yang biasanya ditabur di bubur ayam|DAGING AYAM~AYAM|KACANG~KACANG KEDELAI|DAUN SELEDRI~SELEDRI~SLEDRI|KERUPUK~KRUPUK|BAWANG~BAWANG GORENG;
Sebutkan nama negara apa saja yang pernah menjajah Indonesia|BELANDA|JEPANG|PORTUGIS|INGGRIS|SPANYOL;
Sebutkan istilah-istilah atau kata lain dari meninggal dunia|MATI|WAFAT|GUGUR|MAMPUS|MODAR;
Sebutkan lambang atau benda apa saja yang terdapat dalam lambang-lambang sila yang ada di Pancasila|BINTANG|POHON BERINGIN~BRINGIN~POHON BRINGIN|KEPALA BANTENG|RANTAI|PADI DAN KAPAS~PADI KAPAS;
Sebutkan console game yang pernah ada|PLAY STASION~PS~PLAYSTATION|NINTENDO|XBOX|SEGA|GAME BOY~GAMEBOY~GIMBOT;
Sebutkan superhero yang kostumnya memiliki warna merah|SPIDERMAN~SPIDER MAN~SPIDER-MAN|IRONMAN~IRON MAN~IRON-MAN|SUPERMAN~SUPER MAN~SUPER-MAN|THOR|DEADPOOL~DEAD POOL~DEAD-POOL;
Sebutkan nama-nama bajak laut yang terkenal di dunia|BLACKBEARD|WILLIAM KIDD~KIDD|ANNE BONNY|SAMUEL BELLAMY~BELAMMY|DAVY JONES;
Sebutkan merk handphone yang dulu berjaya|NOKIA|BLACKBERRY~BB~BLACK BERRY|SONY ERICSSON~ERICSSON|MOTOROLA|SIEMENS;
Sebutkan nama-nama zodiak|LEO|CANCER|LIBRA|CAPRICON|SAGITARIUS;
Sebutkan nama binatang yang sering dijadikan tokoh kartun|ANJING|KUCING|DOMBA~KAMBING|BERUANG|SINGA;
Sebutkan sesuatu yang biasanya dikunci|PINTU|GEMBOK|MOTOR|MOBIL|PAGAR;
Sebutkan binatang melata|ULAR|ULAT~ULET|CACING|LINTAH|BELUT;
Sebutkan istilah relationship antara cewek dan cowok|PACAR~PACARAN|MANTAN|KAKA ADEAN~KAKAK ADEAN~KAKAK-ADEAN|GEBETAN|SAHABAT~TEMENAN~SAHABATAN;
Sebutkan sinetron yang pernah booming|SI DOEL~SI DOEL ANAK SEKOLAHAN|TUKANG BUBUR NAIK HAJI|PUTRI YANG TERTUKAR|TERSANJUNG|GANTENG-GANTENG SRIGALA~GGS~GANTENG-GANTENG SERIGALA;
Sebutkan baju timnas negara mana saja yang sering dipakai oleh masyarakat Indonesia|INDONESIA|BRAZIL|ITALIA|BELANDA|ARGENTINA;
Sebutkan merk pasta gigi yang ada di Indonesia|PEPSODENT|CLOSE UP|ENZIM|CIPTADENT|KODOMO;
Sebutkan mahluk hidup yang berduri|KAKTUS|LANDAK|DURIAN~DUREN|MAWAR|IKAN;
Sebutkan nama dongeng yang terkenal di dunia|PINOKIO|CINDERELA~CINDERELLA|PUTRI SALJU|ALADDIN~ALADIN~ALLADIN|ITIK BURUK RUPA;
Sebutkan nama-nama pemimpin yang ada di Indonesia yang populer saat ini|JOKOWI~PRESIDEN JOKOWI~JOKO WIDODO|AHOK~BASUKI TJAHYA PURNAMA|RIDWAN KAMIL|RISMA~IBU RISMA~TRI RISMA|BIMA ARYA;
Sebutkan nama hewan yang huruf depannya adalah \"C\"|CACING|CICAK|CAPUNG|CENDRAWASIH~CENDRAWASI|CUMI-CUMI;
Sebutkan nama alat ukur|PENGGARIS~METERAN~MISTAR|TERMOMETER|SPEEDOMETER|TIMBANGAN|STOPWATCH;
Sebutkan nama binatang yang dijadikan nama mobil|KIJANG|PANTHER|KUDA|JAGUAR|ZEBRA;
Sebutkan pujian yang diberikan orangtua kepada anaknya|PINTAR|RAJIN|CANTIK~GANTENG|LUCU|NURUT;
Sebutkan bahan-bahan untuk membuat sambel|CABAI~CABE|BAWANG~BAWANG MERAH~BAWANG PUTIH|TERASI|TOMAT|GULA;
Sebutkan lembaga-lembaga khusus yang ada di PBB|UNESCO|WHO|UNICEF|IMF|WFP;
Sebutkan makanan yang dimulai dengan kata bak|BAKPIA|BAKSO|BAKMI~BAKMIE|BAKPAU~BAKPAO~BAKPAW|BAKWAN;
Sebutkan merk buku tulis yang populer di Indonesia|SINAR DUNIA~SIDU|KIKY|BIG BOSS~BIGBOSS|HIPO|TIARA;
Sebutkan apa saja yang ada di universitas|MAHASISWA/MAHASISWI~MAHASISWA~MAHASISWI|DOSEN|REKTOR|SKRIPSI|KULIAH;
Sebutkan anggota Pandawa Lima|ARJUNA|BIMA|YUDISTIRA~YUDHISTIRA|NAKULA|SADEWA;
Sebutkan pelajaran bahasa yang biasanya diajarkan di sekolah Indonesia|INDONESIA|INGGRIS|JEPANG|MANDARIN~CHINA~CINA|ARAB;
Sebutkan nama-nama hewan yang suka dipakai dalam peribahasa Indonesia|SRIGALA~SERIGALA|DOMBA|GAJAH|HARIMAU|KELEDAI;
Sebutkan bagian-bagian dari harimau|BELANG|TARING|CAKAR|EKOR|KUMIS;
Sebutkan benda yang biasa dijemur|BAJU|IKAN~IKAN ASIN|KRUPUK~KERUPUK|KASUR|BANTAL;
Sebutkan nama-nama pemeran OVJ jaman dulu|SULE|ANDRE|AZIZ~AZIS~AZIZ GAGAP|NUNUNG|PARTO;
Sebutkan sesuatu yang bercahaya|LAMPU|MATAHARI|BULAN|BINTANG|SENTER;
Sebutkan nama buah yang kulitnya kasar tidak halus|DURIAN~DUREN|SIRSAK|NANAS|RAMBUTAN|SALAK;
Sebutkan sayuran yang bisa langsung dimakan|MENTIMUN~TIMUN~BONTENG|TOMAT|KEMANGI~DAUN KEMANGI|WORTEL|SELADA~DAUN SELADA;
Sebutkan warna plat nomor kendaraan yang ada di Indonesia|HITAM|KUNING|MERAH|BIRU|PUTIH;
Sebutkan makanan khas yang ada di saat Lebaran|KETUPAT|NASTAR|OPOR~OPOR AYAM|RENDANG|PUTRI SALJU;
Sebutkan hal yang terkenal dari kota Jakarta|MACET|BANJIR|MONAS|PANAS|AHOK;
Sebutkan benda yang berbentuk bundar|RODA~BAN|KOIN|PIRING|CD~DVD|SETIR~STIR~STIR MOBIL;
Sebutkan dimana saja kamu biasa memasukkan password|FACEBOOK|GMAIL|WIFI|INSTAGRAM~IG|TWITTER;
Sebutkan bank-bank yang ada di Indonesia|BCA|MANDIRI|BRI|BUKOPIN|BNI;
Sebutkan jenis-jenis sel darah putih|NEUTROFIL|EOSINOFIL|LIMSOFIT|BASOFIL|MONOSIT;
Sebutkan antivirus yang populer di Indonesia|AVG|AVIRA|SMADAV|NORTON|AVAST;
Sebutkan istilah-istilah yang ada dalam program Microsoft Excel|FORMULA|CELL|WORKSHEET~SHEET|RANGE|ROW;
Sebutkan makanan atau minuman yang membuat asam lambung meningkat|KOPI|COKELAT~COKLAT|KACANG|SODA~MINUMAN SODA~SOFT DRINK|ALKOHOL;
Sebutkan jenis gaya yang ada dalam ilmu fisika|PEGAS~GAYA PEGAS|GESEK~GAYA GESEK|GRAVITASI~GAYA GRAVITASI|MAGNET~GAYA MAGNET|OTOT~GAYA OTOT;
Sebutkan macam-macam ukuran kertas yang biasanya ada di program Microsoft Word|A4|A3|F4|LEGAL|LETTER;
Sebutkan jenis-jenis karangan dalam bahasa Indonesia|NARASI~KARANGAN NARASI|DESKRIPSI~KARANGAN DESKRIPSI|ARGUMENTASI~KARANGAN ARGUMENTASI|EKSPOSISI~KARANGAN EKSPOSISI|PERSUASI~KARANGAN PERSUASI;
Sebutkan macam-macam emosi yang ada dalam diri manusia|MARAH|BAHAGIA~SENANG~GEMBIRA|SEDIH~MURUNG|MALU|TAKUT;
Sebutkan benda-benda yang suka ditiup|BALON|LILIN~API LILIN~LILIN ULANG TAHUN|SULING~SERULING|TEROMPET|DEBU;
Sebutkan penyanyi solo pria yang terpopuler di Indonesia|GLENN FREDLY~GLEN FREDLY~GLEN FREDLI|JUDIKA|TOMPI|AFGAN|TULUS;
Sebutkan judul lagu nasional Indonesia|INDONESIA RAYA|BAGIMU NEGERI|MAJU TAK GENTAR|GUGUR BUNGA|GARUDA PANCASILA;
Sebutkan atlet Indonesia yang meraih mendali emas di Olimpiade|SUSI SUSANTI|TONTOWI~TONTOWI AHMAD|LILIYANA|TAUFIK HIDAYAT|ALAN BUDIKUSUMA;
Sebutkan benda-benda yang suka didorong|GEROBAK|TROLI~TROLY~TROLLY|KERETA BAYI~STROLLER~KERETA DORONG BAYI|MOBIL~MOBIL MOGOK|MEJA;
Sebutkan benda-benda yang biasa kita buka|PINTU|BAJU~PAKAIAN|LEMARI~LEMARI PAKAIAN|KULKAS|CELANA;
Sebutkan nama-nama bencana alam|GEMPA~GEMPA BUMI|GUNUNG MELETUS|BANJIR|LONGSOR|TSUNAMI;
Sebutkan nama-nama nominal uang rupiah dalam bahasa Hokien|GOCENG|CEPEK|SECENG|GOBAN|JIGO;
Sebutkan nama atau istilah yang berarti uang|DUIT|DOKU|MONEY|FULUS|HEPENG;
Sebutkan istilah-istilah atau nama lain dari kata : saya|GUA~GUE~GW|AKU|ANA~ANE|BETA|AKIKA~EIKE;
Sebutkan isitilah-istilah atau nama lain dari kata : lemas|LESU|LUNGLAI|LOYO|LEMAH|LETIH;
Sebutkan merk-merk laptop|ACER|LENOVO|ASUS|DELL|TOSHIBA;
Sebutkan profesi yang memakai seragam|ANAK SEKOLAH~SISWA~PELAJAR|POLISI|TNI~TENTARA~ABRI|PNS~PEGAWAI NEGERI SIPIL|BURUH~BURUH PABRIK~KARYAWAN;
Sebutkan nama-nama superhero yang kostumnya ada warna merahnya|SPIDERMAN~SPIDER MAN~SPIDER-MAN|IRONMAN~IRON MAN~IRON-MAN|THE FLASH|DEADPOOL|SUPERMAN~SUPER MAN~SUPER-MAN;
Sebutkan nama-nama penyanyi solo wanita Indonesia (POP)|AGNES MONICA~AGMO|ROSSA|SYAHRINI|RAISA|BUNGA CITRA LESTARI~BCL;
Sebutkan nama-nama penyanyi solo wanita Indonesia (DANGDUT)|AYU TING TING|INUL~INUL DARATISTA|ZASKIA GOTIK|ELVY SUKAESIH|RITA SUGIARTO;
Sebutkan binatang yang dagingnya sering dikonsumsi oleh masyarakat Indonesia|AYAM|SAPI|KAMBING|IKAN|ULAR;
Sebutkan nama-nama band Indonesia yang memiliki vokalis wanita|KOTAK|GEISHA|COKELAT|MOCCA|UTOPIA;
Sebutkan singkatan-singkatan gaul yang sering digunakan anak muda Indonesia|PHP|LDR|MODUS|OTW|BT~BETE;
Sebutkan nama-nama masakan Indonesia yang berkuah|BAKSO~BASO|SOP|SOTO|RENDANG|GULAI;
Sebutkan benda-benda yang memiliki resleting|CELANA~JEANS|TAS~RANSEL~TAS RANSEL|JAKET|DOMPET|TENDA;
Sebutkan pelawak yang sering tampil di acara ILK|DENNY CANDRA~DENNY CHANDRA|CAK LONTONG|JARWO~JARWO KUAT|KOMENG|RICO CEPER;
Sebutkan nama-nama penyanyi wanita yang diawalin huruf R|RAISA|ROSSA|RUTH SAHANAYA|RITA SUGIARTO|REZA;
Sebutkan judul lagu yang dinyanyikan oleh grup band PETERPAN|TOPENG|ADA APA DENGANMU|MUNGKIN NANTI|SEMUA TENTANG KITA|DI BALIK AWAN;
Sebutkan merk smartphone yang dijual di Indonesia|SAMSUNG|APPLE|LG|OPPO|BLACKBERRY;
Sebutkan kata tanya dalam bahasa Inggris|WHAT|WHO|WHERE|WHEN|HOW;
Sebutkan bahan-bahan apa saja yang sering dijadikan topping pizza|SOSIS~PEPPERONI~PEPERONI|KEJU~KEJU MOZZARELLA~MOZZARELLA|DAGING~DAGING SAPI~BEEF|JAMUR|PAPRIKA;
Sebutkan rasa-rasa dari Indomie goreng|RENDANG~RASA RENDANG|SPESIAL~RASA SPESIAL|AYAM BAWANG~RASA AYAM BAWANG|CABE IJO~RASA CABE IJO|SOTO~RASA SOTO;
Sebutkan bahan-bahan untuk membuat kue|TEPUNG~TERIGU~TEPUNG TERIGU|TELUR~TELOR|GULA~GULA PASIR|SUSU|SODA KUE~BAKING SODA~BAKING POWDER;
Sebutkan merk balsam / balsem|CAPLANG~CAP LANG|GELIGA|BALPIRIK|VICKS|TJING TJAU~CINGCAU~CING CAU;
Sebutkan olahraga yang menggunakan tangan|BASKET|VOLI~VOLLY|TINJU|BADMINTON~BULU TANGKIS|TENIS;
Sebutkan merk kecap manis|BANGO|ABC|SEDAAP~SEDAP|INDOFOOD|CAP BEMO~BEMO;
Sebutkan organ-organ yang terdapat di dalam tubuh manusia|JANTUNG|GINJAL|PARU-PARU|LAMBUNG|HATI~LIVER~LEVER;
Sebutkan nama tokoh yang ada dalam animasi Boboiboy|BOBOIBOY|GOPAL|ADU DU~ADUDU|YAYA|PROBE;
Sebutkan nama-nama hewan yang ada di film animasi Masha and The Bear|BERUANG|KELINCI|SRIGALA~SERIGALA|TUPAI|HARIMAU~MACAN;
Sebutkan benda-benda yang digantung atau ditempel di dinding|JAM DINDING~JAM|FOTO~FOTO KELUARGA|LUKISAN|POSTER|CERMIN~KACA;
Sebutkan benda-benda yang ada di meja makan|PIRING|SENDOK|GARPU|GELAS|TUDUNG SAJI;
Sebutkan merk sepatu bola|NIKE|ADIDAS|SPECS|PUMA|DIADORA;
Sebutkan film animasi Jepang yang bertemakan olahraga|CAPTAIN TSUBASA~TSUBASA~KAPTEN TSUBASA|SLAM DUNK|EYESHIELD 21|HAIKYUU|PRINCE OF TENNIS;
Sebutkan merk lotion yang dijual di Indonesia|VASELINE|CITRA|NIVEA|DOVE|MARINA;
Sebutkan jenis pelayanan yang ditawarkan di salon kecantikan|CREAMBATH|POTONG RAMBUT~CUKUR RAMBUT~GUNTING RAMBUT|FACIAL|LULUR|MANICURE PEDICURE~MANI PEDI~MENI PEDI;
Sebutkan makanan yang sering kita temui saat hari Lebaran|KETUPAT|OPOR AYAM|NASTAR~KUE NASTAR|RENDANG|PUTRI SALJU~KUE PUTRI SALJU;
Sebutkan sayuran atau buah yang berbau tajam atau tidak enak|JENGKOL|PETAI~PETE|DURIAN~DUREN|MENGKUDU|BAWANG~BAWANG PUTIH~BAWANG MERAH;
Sebutkan restoran francais atau waralaba yang ada di Indonesia|MCDONALD'S~MCDONALD|PIZZA HUT|KFC~KENTUCKY FRIED CHICKEN|HOKA HOKA BENTO|STARBUCKS;
Sebutkan alat-alat musik yang memiliki senar|GITAR|BIOLA|KECAPI|HARPA|UKULELE;
Sebutkan negara-negara tujuan TKI (Tenaga Kerja Indonesia)|MALAYSIA|HONGKONG|ARAB SAUDI|TAIWAN|KOREA SELATAN;
Sebutkan makanan Jepang yang populer di Indonesia|SUSHI|MIE RAMEN~RAMEN~MI RAMEN|TAKOYAKI|TERIYAKI|YAKINIKU;
Sebutkan nama-nama hewan yang bercangkang|SIPUT~KEONG~BEKICOT|KURA-KURA~KURA KURA|PENYU|KERANG|KEPITING;
Sebutkan nama-nama buah yang berawalan huruf K|KEDONDONG|KELENGKENG|KELAPA|KIWI|KESEMEK;
Sebutkan nama-nama buah yang berawalan huruf M|MANGGA|MELON|MANGGIS|MARKISA|MENGKUDU;
Sebutkan negara-negara tujuan wisata selain Indonesia|PRANCIS|AMERIKA SERIKAT~USA|SINGAPURA~SINGAPUR|INGGRIS|MALAYSIA;
Sebutkan nama-nama hewan yang diawali huruf A|ANJING|AYAM|ANGSA|ANOA|ANAKONDA;
Sebutkan nama-nama hewan omnivora|MONYET~KERA|TIKUS|AYAM|BURUNG|BABI;
Sebutkan tombol yang ada di stick PS|X~SILANG|KOTAK|BULAT~LINGKARAN~O|SEGITIGA|START;
Sebutkan merk pengharum ruangan|AMBIPUR|STELLA|BAYFRESH|GLADE|SWALLOW;
`

// {
// 	"questions": [
// 	  {"id": 0, "text": "Apa yang dilakukan jika kebelet pipis saat berenang?", "answers": [
// 		{"score": 32, "text": "pipis dalam kolam", "tag": ["pipis", "buang air", "kencing"]},
// 		{"score": 24, "text": "lari ke toilet", "tag": ["ke toilet"]},
// 		{"score": 22, "text": "ditahan", "tag": ["tahan"]},
// 		{"score": 18, "text": "naik", "tag": ["pinggir kolam"]}
// 	  ]},
// 	  {"id": 1, "text": "Bagaimana cara agar tidak masuk angin?", "answers": [
// 		{"score": 37, "text": "pakai jaket", "tag": ["jaket"]},
// 		{"score": 23, "text": "makan teratur", "tag": ["makan"]},
// 		{"score": 15, "text": "minum obat", "tag": ["obat"]},
// 		{"score": 9, "text": "tidak begadang"},
// 		{"score": 4, "text": "minum air hangat", "tag": ["minum"]}
// 	  ]},
// 	  {"id": 2, "text": "Kegiatan apa yang dilakukan saat cuaca dingin?", "answers": [
// 		{"score": 39, "text": "tidur"},
// 		{"score": 28, "text": "minum kopi"},
// 		{"score": 13, "text": "makan"},
// 		{"score": 8, "text": "minum teh", "tag": ["teh"]},
// 		{"score": 5, "text": "pelukan"}
// 	  ]},
// 	  {"id": 3, "text": "Apa yang pria lakukan saat menemani kekasihnya di salon?", "answers": [
// 		{"score": 43, "text": "main ponsel", "tag": ["main handphone"]},
// 		{"score": 34, "text": "membaca", "tag": ["bac"]},
// 		{"score": 11, "text": "merokok"},
// 		{"score": 7, "text": "dengar musik", "tag": ["mendengar musik"]}
// 	  ]},
// 	  {"id": 4, "text": "Minuman apa yang membutuhkan air panas?", "answers": [
// 		{"score": 52, "text": "kopi"},
// 		{"score": 24, "text": "susu"},
// 		{"score": 12, "text": "teh"},
// 		{"score": 6, "text": "bandrek"}
// 	  ]},
// 	  {"id": 5, "text": "Sebutkan sesuatu yang berhubungan dengan printer?", "answers": [
// 		{"score": 37, "text": "kertas"},
// 		{"score": 32, "text": "tinta"},
// 		{"score": 16, "text": "printer"},
// 		{"score": 15, "text": "catridge"}
// 	  ]},
// 	  {"id": 6, "text": "Sebutkan cara untuk mendapatkan barang impian?", "answers": [
// 		{"score": 41, "text": "menabung"},
// 		{"score": 26, "text": "bekerja"},
// 		{"score": 14, "text": "hutang", "tag": ["menghutang"]},
// 		{"score": 5, "text": "menjual barang lain", "tag": ["jual barang"]}
// 	  ]},
// 	  {"id": 7, "text": "?", "answers": [
// 		{"score": 39, "text": "", "tag": [""]},
// 		{"score": 28, "text": "", "tag": [""]},
// 		{"score": 13, "text": "", "tag": [""]},
// 		{"score": 8, "text": "", "tag": [""]},
// 		{"score": 5, "text": "", "tag": [""]}
// 	  ]}
// 	]
//   }
