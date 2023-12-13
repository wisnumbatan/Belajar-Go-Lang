import math

def hitung_jarak(planet1, planet2):
    # Koordinat planet dalam satuan kilometer
    koordinat_planet = {
        'mercury': (0, 0),
        'venus': (100, 200),
        'earth': (300, 400),
        'mars': (500, 600),
        'jupiter': (700, 800),
        'saturn': (900, 1000),
        'uranus': (1100, 1200),
        'neptune': (1300, 1400),
        'pluto': (1500, 1600)
    }

    # Memeriksa apakah planet ada dalam daftar koordinat
    if planet1 not in koordinat_planet or planet2 not in koordinat_planet:
        return 'Planet tidak valid'

    # Menghitung jarak antara dua planet menggunakan rumus jarak Euclidean
    x1, y1 = koordinat_planet[planet1]
    x2, y2 = koordinat_planet[planet2]
    jarak = math.sqrt((x2 - x1)**2 + (y2 - y1)**2)

    return jarak

# Contoh penggunaan
jarak_mars_venus = hitung_jarak('mars', 'venus')
print(f"Jarak antara Mars dan Venus adalah {jarak_mars_venus} kilometer")
