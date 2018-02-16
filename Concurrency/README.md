Dilakukan pengambilan data museum di Indonesia dengan API yang telah disediakan, ambil API Provinsi dahulu untuk mendapatkan kode per Provinsi, kemudian berdasarkan Kode Provinsi ambil API Kota untuk mendapatkan Kode Kota, 
setelah dapat Kode Kota maka dapat di ambil data Museum dari API Museum per Kota/Kabupaten.

Gunakan goroutine untuk mengambil data di tiap API
Setelah data didapat dengan goroutine maka simpan kedalam file.csv berdasarkan masing-masing kota/kabupaten.

Tujuan dibuatnya sistem ini untuk mempelajari penggunaan API sebagai pusat data dengan metode pengambilan layaknya multi thread untuk mempercepat proses pengambilan data.