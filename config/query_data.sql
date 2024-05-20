CREATE DATABASE data_makanan;

CREATE TABLE `makanan`(
	`id` int NOT NULL AUTO_INCREMENT,
    `nama` varchar(255) NOT NULL,
    `harga` float NOT NULL,
    `rating` float NOT NULL,
	`jarak` float NOT NULL ,
	`lokasi` varchar(255) NOT NULL,
	PRIMARY KEY (`id`)
);

INSERT INTO makanan (Nama, Harga, Rating, Jarak, Lokasi) VALUES
('Ayam Gepuk Pak Gembus', 24, 4.4, 1.2, 'https://maps.app.goo.gl/5Ss9Ry4gCaM1GV7z6'),
('Mie Ayam "Baso Budi"', 14, 4.8, 1, 'https://maps.app.goo.gl/BtxLJo47FzSzrCmL9'),
('Batagor MSU', 10, 5, 1.3, 'https://maps.app.goo.gl/Sx9945WJkQXs24UP7'),
('Sate Suramadu', 25, 4.3, 1.3, 'https://maps.app.goo.gl/skXXHUmvV1WVZVXx8'),
('Ketoprak Mas No', 15, 4.9, 0.8, 'https://maps.app.goo.gl/GMP4A8dHAc9pS9CL6'),
('Nasi Pecel Lele "Lumintu 768"', 15, 4.5, 1, 'https://maps.app.goo.gl/DXgV4WZjcoSELyaAA'),
('Geprek Crispy "Ayam Crisbar"', 14, 4.4, 1.1, 'https://maps.app.goo.gl/FcjSGXRb5xTKavMf6'),
('Mie Baso "Mas Japra Solo"', 15, 4.7, 1.2, 'https://maps.app.goo.gl/nHSgDhB6JcXZHTVT8'),
('Mie Ayam "Jabrig"', 10, 4.5, 1.2, 'https://maps.app.goo.gl/t9giWhD6E3MdT7o78'),
('Yakiniku Rice Bowl "Spicy Yakiniku"', 20, 4.8, 0.95, 'https://maps.app.goo.gl/EgWp5aQSPKZvhNPS7'),
('Classic Beef Kebab "Merhaba Kebab"', 18, 5, 1.4, 'https://maps.app.goo.gl/fpLQPxgiSeJDwdjx8'),
('Paket Jumbo Chicken "Chicken William"', 15, 4.3, 1.7, 'https://maps.app.goo.gl/drA7Kznw118A6uiG6'),
('Paket Nila Merah Nasi "Ikan Dan Ayam Bakar Pesona Bali"', 26, 5, 3.1, 'https://maps.app.goo.gl/BHxZu27AUfJAmxum7'),
('Seblak "Azka"', 15, 4, 1, 'https://maps.app.goo.gl/3ithL5ePS8kyPU4NA'),
('Nasi Goreng Telor "Kedai Aas"', 12, 4.6, 1.4, 'https://maps.app.goo.gl/2Lh7thn32hUgzbR69'),
('Kwetiau Goreng Seafood "Kedai Aas"', 14, 4.6, 1.4, 'https://maps.app.goo.gl/2Lh7thn32hUgzbR69'),
('Nasi Goreng "Moro Tresno"', 13, 2.3, 0.9, 'https://maps.app.goo.gl/4gkAkb4WL5AN5WAt6'),
('Mie Tek-Tek "Warkop Djoeang"', 10, 3, 1.1, 'https://maps.app.goo.gl/RfnkhfRrsA6qjum49'),
('Dori Fish "Mr. Mangkok"', 18, 4.5, 1.8, 'https://maps.app.goo.gl/hPWTHf8jFhcSNnK66'),
('Capcay Ayam "Kedai Aas"', 13, 4.6, 1.4, 'https://maps.app.goo.gl/2Lh7thn32hUgzbR69');
