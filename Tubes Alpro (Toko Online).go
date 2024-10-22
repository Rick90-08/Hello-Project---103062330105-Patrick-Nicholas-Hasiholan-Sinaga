// Nama : Patrick Nicholas Hasiholan Sinaga
// NIM : 103062330105
// Kelas : S1 IT-KJ-23-002

package main

import (
    "fmt"
)

// Struct untuk merepresentasikan sebuah barang
type Item struct {
    ID    int
    Name  string
    Price float64
    Stock int
}

// Struct untuk pengguna
type User struct {
    ID       int
    Username string
    Password string
    Role     string // "buyer", "seller", "admin"
    Approved bool
}

// Array statis untuk barang-barang
var items [100]Item
var itemCount int

// Array statis untuk pengguna
var users [100]User
var userCount int

// Struct untuk menyimpan barang yang dibeli
type CartItem struct {
    Item  Item
    Qty   int
    Total float64
}

// Struct untuk transaksi
type Transaction struct {
    UserID int
    Cart   [maxCartItems]CartItem // Ubah menjadi array statis
    Total  float64
}

const maxCartItems = 100 // Tentukan kapasitas maksimum untuk cart dalam setiap transaksi

const maxTransactions = 100 // Tentukan kapasitas maksimum transaksi

var transactions [maxTransactions]Transaction
var transactionCount int

func main() {
    // Mengisi array items dengan dummy data
    items[0] = Item{ID: 1, Name: "Laptop", Price: 10000000, Stock: 5}
    items[1] = Item{ID: 2, Name: "Smartphone", Price: 5000000, Stock: 10}
    items[2] = Item{ID: 3, Name: "Headphone", Price: 1500000, Stock: 15}
    items[3] = Item{ID: 4, Name: "Mouse", Price: 250000, Stock: 20}
    itemCount = 4

    // Menambahkan admin default
    users[0] = User{ID: 1, Username: "admin", Password: "admin", Role: "admin", Approved: true}
    userCount = 1

    for {
        fmt.Println("\nSelamat datang di Toko Online")
        fmt.Println("1. Registrasi Akun")
        fmt.Println("2. Login")
        fmt.Println("3. Info")
        fmt.Println("4. Keluar")
        fmt.Print("Pilih (1/2/3/4): ")
        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            registerAccount()
        case 2:
            login()
        case 3:
            info()
        case 4:
            fmt.Println("Terima kasih telah menggunakan Toko Online!")
            return
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}

func info() {
    fmt.Println("\nCara Kerja Toko Online:")
    fmt.Println("Toko Online ini menyediakan beberapa fitur utama:")
    fmt.Println("1. Registrasi Akun: Pengguna dapat mendaftar sebagai pembeli atau penjual.")
    fmt.Println("2. Login: Pengguna yang telah terdaftar dapat masuk ke dalam sistem.")
    fmt.Println("3. Info: Memberikan informasi singkat tentang cara kerja program.")
    fmt.Println("4. Keluar: Keluar dari program.")
    fmt.Println("Jika ingin login sebagai admin masuk ke login dan masukkan username dan password sebagai 'admin'. ")
    fmt.Print("Untuk kembali ke menu utama, ketik '0': ")
    var back int
    fmt.Scan(&back)
    if back == 0 {
        return
    } else {
        fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        info()
    }
}

// Fungsi untuk mendaftarkan akun baru
func registerAccount() {
    // Meminta input dari pengguna untuk username, password, dan peran
    var username, password, role string
    fmt.Print("\nMasukkan username: ")
    fmt.Scan(&username)
    fmt.Print("Masukkan password: ")
    fmt.Scan(&password)
    fmt.Print("Masukkan peran (buyer/seller): ")
    fmt.Scan(&role)

    // Memeriksa apakah peran yang dimasukkan valid (buyer atau seller)
    if role != "buyer" && role != "seller" {
        fmt.Println("Peran tidak valid. Silakan coba lagi.")
        return
    }

    // Menentukan ID baru untuk pengguna yang akan didaftarkan
    id := userCount + 1

    // Menambahkan pengguna baru ke dalam array users
    users[userCount] = User{ID: id, Username: username, Password: password, Role: role, Approved: false}
    userCount++

    // Memberikan konfirmasi registrasi kepada pengguna
    fmt.Println("Registrasi berhasil! Menunggu persetujuan admin.")
}

// Fungsi untuk melakukan proses login
func login() {
    // Meminta input username dan password dari pengguna
    var username, password string
    fmt.Print("\nMasukkan username: ")
    fmt.Scan(&username)
    fmt.Print("Masukkan password: ")
    fmt.Scan(&password)

    // Memeriksa keberadaan pengguna berdasarkan username dan password yang dimasukkan
    user, found := authenticate(username, password)
    if !found {
        // Jika pengguna tidak ditemukan, cetak pesan kesalahan
        fmt.Println("Username atau password salah.")
        return
    }

    // Memeriksa apakah akun pengguna sudah disetujui oleh admin
    if !user.Approved {
        fmt.Println("Akun Anda belum disetujui oleh admin.")
        return
    }

    // Berdasarkan peran pengguna, panggil fungsi mode yang sesuai
    switch user.Role {
    case "buyer":
        buyerMode(user.ID)
    case "seller":
        sellerMode(user.ID)
    case "admin":
        adminMode()
    }
}

// Fungsi untuk mengautentikasi pengguna berdasarkan username dan password
func authenticate(username, password string) (User, bool) {
    // Iterasi melalui array pengguna untuk mencocokkan username dan password
    for i := 0; i < userCount; i++ {
        if users[i].Username == username && users[i].Password == password {
            return users[i], true // Jika ditemukan, kembalikan pengguna dan true
        }
    }
    return User{}, false // Jika tidak ditemukan, kembalikan pengguna kosong dan false
}

// Fungsi untuk masuk ke dalam mode Admin
func adminMode() {
    // Looping tanpa batas untuk menampilkan opsi dan menerima input dari admin
    for {
        fmt.Println("\nMode Admin")
        fmt.Println("1. Setujui atau tolak registrasi akun")
        fmt.Println("2. Logout")
        fmt.Print("Pilih (1/2): ")
        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            approveOrRejectAccounts() // Memanggil fungsi untuk menyetujui atau menolak registrasi akun
        case 2:
            return // Keluar dari mode Admin
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}

// Fungsi untuk menyetujui atau menolak registrasi akun
func approveOrRejectAccounts() {
    // Iterasi melalui array pengguna untuk menampilkan pengguna yang belum disetujui
    for i := 0; i < userCount; i++ {
        if !users[i].Approved {
            fmt.Printf("\nID: %d, Username: %s, Role: %s\n", users[i].ID, users[i].Username, users[i].Role)
            fmt.Println("Setujui akun ini? (yes/no): ")
            var response string
            fmt.Scan(&response)
            if response == "yes" {
                users[i].Approved = true // Menyetujui akun jika admin memilih "yes"
                fmt.Println("Akun disetujui.")
            } else {
                fmt.Println("Akun ditolak.") // Menolak akun jika admin memilih "no"
            }
        }
    }
}

// Fungsi buyerMode mengatur interaksi ketika pengguna masuk ke mode Pembeli.
// Di sini, pengguna dapat melihat daftar barang yang tersedia, membeli barang, dan menyelesaikan pembelian.
func buyerMode(userID int) {
    const maxCartItems = 100
    var cart [maxCartItems]CartItem // Membuat array statis untuk menyimpan item yang dibeli dalam keranjang belanja
    var cartItemCount int           // Menyimpan jumlah item dalam keranjang belanja

    for {
        fmt.Println("\nMode Pembeli")
        viewItems() // Menampilkan daftar barang yang tersedia

        fmt.Println("Masukkan ID barang yang ingin Anda beli (atau 0 untuk keluar) atau tekan 'F' untuk filter:")
        var input string
        fmt.Scan(&input)

        if input == "0" {
            if cartItemCount > 0 {
                completePurchase(userID, cart, cartItemCount) // Menyelesaikan pembelian jika keranjang belanja tidak kosong
            }
            return
        } else if input == "F" || input == "f" {
            filterItems() // Memfilter daftar barang jika pengguna memilih 'F' atau 'f'
            continue
        }

        var id int
        fmt.Sscanf(input, "%d", &id)

        item, found := getItemByID(id)
        if !found {
            fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
            continue
        }

        fmt.Printf("Anda memilih %s dengan harga Rp%.2f\n", item.Name, item.Price)
        fmt.Println("Masukkan jumlah yang ingin dibeli:")
        var qty int
        fmt.Scan(&qty)

        if qty > item.Stock {
            fmt.Printf("Stok tidak mencukupi. Stok tersedia: %d\n", item.Stock)
        } else {
            updateStock(id, qty) // Mengupdate stok barang setelah pembelian
            total := float64(qty) * item.Price
            cart[cartItemCount] = CartItem{Item: item, Qty: qty, Total: total} // Menambah item yang dibeli ke dalam keranjang belanja
            cartItemCount++
            fmt.Printf("Anda telah membeli %d %s dengan total harga Rp%.2f\n", qty, item.Name, total)
        }
    }
}

// Fungsi sellerMode mengatur interaksi ketika pengguna masuk ke mode Penjual.
// Di sini, penjual dapat melakukan beberapa tindakan seperti menambah, mengubah, atau menghapus barang,
// melihat daftar barang yang diurutkan, mencetak transaksi, atau keluar dari mode Penjual.
func sellerMode(userID int) {
    for {
        fmt.Println("\nMode Penjual")
        fmt.Println("1. Tambah barang")
        fmt.Println("2. Ubah barang")
        fmt.Println("3. Hapus barang")
        fmt.Println("4. Lihat barang (terurut)")
        fmt.Println("5. Cetak transaksi")
        fmt.Println("6. Logout")
        fmt.Print("Pilih (1/2/3/4/5/6): ")
        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            addItem() // Menambah barang baru ke dalam daftar barang yang tersedia
        case 2:
            editItem() // Mengubah informasi barang yang sudah ada dalam daftar
        case 3:
            deleteItem() // Menghapus barang dari daftar barang yang tersedia
        case 4:
            viewItems() // Menampilkan daftar barang yang diurutkan
        case 5:
            printTransactions() // Mencetak transaksi yang telah dilakukan
        case 6:
            return // Keluar dari mode Penjual dan kembali ke menu utama
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.") // Menampilkan pesan jika pilihan tidak valid
        }
    }
}

// Fungsi addItem digunakan untuk menambahkan barang baru ke dalam daftar barang yang tersedia.
// Jika kapasitas maksimum barang telah tercapai, fungsi akan menampilkan pesan bahwa kapasitas penuh.
// Jika belum, pengguna diminta untuk memasukkan nama, harga, dan stok barang baru.
// Setelah memasukkan informasi barang baru, barang tersebut akan ditambahkan ke dalam daftar barang,
// dan jumlah barang yang ada akan diperbarui.
func addItem() {
    if itemCount >= 100 {
        fmt.Println("Tidak bisa menambah barang lagi, kapasitas penuh.")
        return
    }
    var id int
    if itemCount > 0 {
        id = items[itemCount-1].ID + 1 // ID baru adalah ID terakhir ditambah 1
    } else {
        id = 1 // Jika belum ada barang, ID baru dimulai dari 1
    }
    fmt.Print("\nMasukkan nama barang: ")
    var name string
    fmt.Scan(&name)
    fmt.Print("Masukkan harga barang: ")
    var price float64
    fmt.Scan(&price)
    fmt.Print("Masukkan stok barang: ")
    var stock int
    fmt.Scan(&stock)

    newItem := Item{ID: id, Name: name, Price: price, Stock: stock} // Membuat barang baru dengan informasi yang dimasukkan
    items[itemCount] = newItem // Menambahkan barang baru ke dalam daftar barang
    itemCount++ // Menambah jumlah barang yang ada
    fmt.Println("Barang berhasil ditambahkan!") // Menampilkan pesan bahwa barang berhasil ditambahkan
}

// Fungsi editItem digunakan untuk mengubah informasi barang yang sudah ada berdasarkan ID barang yang diberikan.
// Pengguna diminta untuk memasukkan ID barang yang ingin diubah.
// Jika barang dengan ID tersebut tidak ditemukan, fungsi akan menampilkan pesan bahwa barang tidak ditemukan dan berhenti.
// Jika ditemukan, pengguna diminta untuk memasukkan informasi baru untuk nama, harga, dan stok barang.
// Informasi barang yang ada akan diperbarui dengan informasi baru yang dimasukkan.
func editItem() {
    fmt.Print("\nMasukkan ID barang yang ingin diubah: ")
    var id int
    fmt.Scan(&id)

    index, found := getItemIndexByID(id) // Mendapatkan indeks barang dalam daftar berdasarkan ID
    if !found {
        fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
        return
    }

    fmt.Print("Masukkan nama barang baru: ")
    var name string
    fmt.Scan(&name)
    fmt.Print("Masukkan harga barang baru: ")
    var price float64
    fmt.Scan(&price)
    fmt.Print("Masukkan stok barang baru: ")
    var stock int
    fmt.Scan(&stock)

    items[index] = Item{ID: id, Name: name, Price: price, Stock: stock} // Memperbarui informasi barang yang ada
    fmt.Println("Barang berhasil diubah!") // Menampilkan pesan bahwa barang berhasil diubah
}

// Fungsi deleteItem digunakan untuk menghapus barang dari daftar berdasarkan ID barang yang diberikan.
// Pengguna diminta untuk memasukkan ID barang yang ingin dihapus.
// Jika barang dengan ID tersebut tidak ditemukan, fungsi akan menampilkan pesan bahwa barang tidak ditemukan dan berhenti.
// Jika ditemukan, barang tersebut akan dihapus dari daftar.
// Setelah penghapusan, jumlah barang (itemCount) akan dikurangi satu.
func deleteItem() {
    fmt.Print("\nMasukkan ID barang yang ingin dihapus: ")
    var id int
    fmt.Scan(&id)

    index, found := getItemIndexByID(id) // Mendapatkan indeks barang dalam daftar berdasarkan ID
    if !found {
        fmt.Println("Barang dengan ID tersebut tidak ditemukan.")
        return
    }

    // Memindahkan semua barang setelah barang yang dihapus ke posisi sebelumnya
    for i := index; i < itemCount-1; i++ {
        items[i] = items[i+1]
    }
    itemCount-- // Mengurangi jumlah barang setelah penghapusan

    fmt.Println("Barang berhasil dihapus!") // Menampilkan pesan bahwa barang berhasil dihapus
}

// Fungsi viewItems digunakan untuk menampilkan daftar barang yang tersedia.
// Fungsi ini akan mencetak ID, nama, harga, dan stok dari setiap barang dalam daftar.
func viewItems() {
    fmt.Println("\nDaftar barang yang tersedia:")
    for i := 0; i < itemCount; i++ {
        item := items[i]
        fmt.Printf("ID: %d, Nama: %s, Harga: Rp%.2f, Stok: %d\n", item.ID, item.Name, item.Price, item.Stock)
    }
}

// Fungsi filterItems digunakan untuk memfilter dan mengurutkan daftar barang berdasarkan pilihan pengguna.
// Pengguna diminta untuk memilih metode pengurutan yang diinginkan.
// Berdasarkan pilihan pengguna, fungsi akan memanggil fungsi pengurutan yang sesuai: selectionSort untuk nama dan stok, insertionSort untuk harga.
// Setelah pengurutan selesai, fungsi akan mencetak daftar barang yang sudah terurut.
func filterItems() {
    fmt.Println("\nPilih metode pengurutan: ")
    fmt.Println("1. Nama (Ascending)")
    fmt.Println("2. Nama (Descending)")
    fmt.Println("3. Harga (Ascending)")
    fmt.Println("4. Harga (Descending)")
    fmt.Println("5. Stok (Ascending)")
    fmt.Println("6. Stok (Descending)")
    fmt.Print("Pilih (1/2/3/4/5/6): ")
    var choice int
    fmt.Scan(&choice)

    switch choice {
    case 1:
        selectionSort(items[:itemCount], "name", true)
    case 2:
        selectionSort(items[:itemCount], "name", false)
    case 3:
        insertionSort(items[:itemCount], "price", true)
    case 4:
        insertionSort(items[:itemCount], "price", false)
    case 5:
        selectionSort(items[:itemCount], "stock", true)
    case 6:
        selectionSort(items[:itemCount], "stock", false)
    default:
        fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        return
    }

    fmt.Println("Daftar barang yang tersedia (terurut):")
    for i := 0; i < itemCount; i++ {
        item := items[i]
        fmt.Printf("ID: %d, Nama: %s, Harga: Rp%.2f, Stok: %d\n", item.ID, item.Name, item.Price, item.Stock)
    }
}

// Fungsi selectionSort digunakan untuk mengurutkan slice dari struktur Item menggunakan algoritma selection sort.
// Pengguna dapat menentukan kunci pengurutan ("name", "price", atau "stock") dan arah pengurutan (ascending atau descending).
// Fungsi melakukan pengurutan berdasarkan kunci yang diberikan, dengan menggeser elemen slice sesuai urutan yang diinginkan.
// Proses pengurutan dilakukan dengan membandingkan elemen-elemen secara berpasangan dan menukar posisi jika diperlukan.
// Pada akhirnya, daftar barang akan diurutkan sesuai dengan kunci dan arah pengurutan yang ditentukan.
func selectionSort(items []Item, by string, ascending bool) {
    n := len(items)
    for i := 0; i < n-1; i++ {
        idx := i
        for j := i + 1; j < n; j++ {
            switch by {
            case "name":
                if ascending {
                    if items[j].Name < items[idx].Name {
                        idx = j
                    }
                } else {
                    if items[j].Name > items[idx].Name {
                        idx = j
                    }
                }
            case "price":
                if ascending {
                    if items[j].Price < items[idx].Price {
                        idx = j
                    }
                } else {
                    if items[j].Price > items[idx].Price {
                        idx = j
                    }
                }
            case "stock":
                if ascending {
                    if items[j].Stock < items[idx].Stock {
                        idx = j
                    }
                } else {
                    if items[j].Stock > items[idx].Stock {
                        idx = j
                    }
                }
            }
        }
        items[i], items[idx] = items[idx], items[i]
    }
}

// Fungsi insertionSort digunakan untuk mengurutkan slice dari struktur Item menggunakan algoritma insertion sort.
// Pengguna dapat menentukan kunci pengurutan ("name", "price", atau "stock") dan arah pengurutan (ascending atau descending).
// Fungsi ini membandingkan setiap elemen dengan elemen sebelumnya dalam slice, dan jika elemen tersebut tidak dalam posisi yang benar,
// maka digeser ke posisi yang sesuai. Proses ini diulangi untuk setiap elemen hingga seluruh slice diurutkan.
// Fungsi compareItems digunakan untuk membandingkan dua elemen berdasarkan kunci yang diberikan.
// Ini membantu dalam menentukan apakah elemen harus dipindahkan ke posisi yang lebih awal atau lebih akhir dalam urutan yang diinginkan.
func insertionSort(items []Item, by string, ascending bool) {
    n := len(items)
    for i := 1; i < n; i++ {
        key := items[i]
        j := i - 1
        for j >= 0 && ((ascending && compareItems(items[j], key, by) > 0) || (!ascending && compareItems(items[j], key, by) < 0)) {
            items[j+1] = items[j]
            j--
        }
        items[j+1] = key
    }
}

// Fungsi compareItems digunakan dalam algoritma pengurutan insertion sort untuk membandingkan dua elemen berdasarkan kunci tertentu.
// Kunci pengurutan dapat berupa "name", "price", atau "stock".
// Fungsi mengembalikan nilai negatif jika a kurang dari b, nilai positif jika a lebih besar dari b, dan 0 jika a sama dengan b.
func compareItems(a, b Item, by string) int {
    switch by {
    case "name":
        if a.Name < b.Name {
            return -1
        } else if a.Name > b.Name {
            return 1
        }
    case "price":
        if a.Price < b.Price {
            return -1
        } else if a.Price > b.Price {
            return 1
        }
    case "stock":
        if a.Stock < b.Stock {
            return -1
        } else if a.Stock > b.Stock {
            return 1
        }
    }
    return 0
}

// Fungsi getItemByID digunakan untuk mencari dan mengembalikan barang berdasarkan ID yang diberikan dari array items.
// Jika barang dengan ID yang sesuai ditemukan, fungsi mengembalikan barang tersebut bersama dengan true.
// Jika tidak ditemukan, fungsi mengembalikan barang kosong dan false.
// binary search
func getItemByID(id int) (Item, bool) {
    low := 0
    high := itemCount - 1

    for low <= high {
        mid := (low + high) / 2
        if items[mid].ID == id {
            return items[mid], true
        } else if items[mid].ID < id {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return Item{}, false
}

// Fungsi getItemIndexByID digunakan untuk mencari indeks dari barang berdasarkan ID yang diberikan dari array items.
// Jika barang dengan ID yang sesuai ditemukan, fungsi mengembalikan indeksnya bersama dengan true.
// Jika tidak ditemukan, fungsi mengembalikan -1 dan false.
// Sequential search
func getItemIndexByID(id int) (int, bool) {
    for i := 0; i < itemCount; i++ {
        if items[i].ID == id {
            return i, true
        }
    }
    return -1, false
}

// Fungsi updateStock digunakan untuk mengurangi stok barang berdasarkan ID dan jumlah yang diberikan.
// Fungsi ini mencari barang dengan ID yang sesuai dalam array items, lalu mengurangi stoknya sesuai dengan jumlah yang ditentukan.
func updateStock(id, qty int) {
    for i := 0; i < itemCount; i++ {
        if items[i].ID == id {
            items[i].Stock -= qty
        }
    }
}

// Fungsi completePurchase digunakan untuk menyelesaikan pembelian dengan menambahkan transaksi ke array transactions.
// Fungsi ini menerima ID pengguna dan keranjang belanja sebagai input.
// Pertama, total pembelian dihitung dari harga total setiap item dalam keranjang.
// Kemudian, sebuah array statis baru dibuat untuk menyimpan keranjang belanja, karena persyaratan meminta penggunaan array statis.
// Transaksi baru ditambahkan ke array transactions jika masih ada kapasitas transaksi yang tersedia.
// Jika kapasitas sudah penuh, pesan kesalahan akan dicetak.
// Setelah transaksi ditambahkan, detail transaksi dicetak, termasuk item yang dibeli dan total pembayarannya.
func completePurchase(userID int, cart [100]CartItem, cartItemCount int) {
    var total float64
    for i := 0; i < cartItemCount; i++ {
        total += cart[i].Total
    }

    // Menambahkan transaksi dengan cart statis ke dalam array transactions
    if transactionCount < maxTransactions {
        transactions[transactionCount] = Transaction{UserID: userID, Cart: cart, Total: total}
        transactionCount++
    } else {
        fmt.Println("Tidak bisa menambah transaksi lagi, kapasitas penuh.")
        return
    }

    fmt.Println("Transaksi berhasil! Barang yang Anda beli:")
    for i := 0; i < cartItemCount; i++ {
        item := cart[i]
        fmt.Printf("%d x %s - Rp%.2f\n", item.Qty, item.Item.Name, item.Total)
    }
    fmt.Printf("Total yang harus dibayar: Rp%.2f\n", total)
}

// Fungsi printTransactions digunakan untuk mencetak daftar transaksi yang telah dilakukan.
// Fungsi ini tidak menerima parameter input.
// Pertama-tama, pesan header "Daftar Transaksi:" dicetak.
// Setelah itu, untuk setiap transaksi dalam array transactions, informasi pengguna yang melakukan transaksi dan total pembelian dicetak.
// Kemudian, untuk setiap item dalam keranjang belanja transaksi, detail item, termasuk jumlah yang dibeli dan total harga, dicetak.
// Setelah semua transaksi dicetak, baris kosong ditambahkan untuk memisahkan transaksi yang berbeda.
func printTransactions() {
    fmt.Println("\nDaftar Transaksi:")
    for _, transaction := range transactions {
        user := users[transaction.UserID-1]
        fmt.Printf("User: %s, Total: Rp%.2f\n", user.Username, transaction.Total)
        for _, item := range transaction.Cart {
            fmt.Printf("%d x %s - Rp%.2f\n", item.Qty, item.Item.Name, item.Total)
        }
        fmt.Println()
    }
}

        fmt.Printf("User: %s, Total: Rp%.2f\n", user.Username, transaction.Total)
        for _, item := range transaction.Cart {
            fmt.Printf("%d x %s - Rp%.2f\n", item.Qty, item.Item.Name, item.Total)
        }
        fmt.Println()
    }
}
