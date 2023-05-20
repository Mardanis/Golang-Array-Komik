package main

import "fmt"

func main() {

	// var jenisKomik = [...]string{
	// 	"Manhwa",
	// 	"Manga",
	// 	"mahwaa",
	// }

	// tampil dalam bentuk gini
	// Judul : I Became the Tyrant of a Defense Game : string
	// Jenis Komik : Manhwa : string
	// Konsep Cerita : 	Isekai : string
	// Tags : Action, Adventure, Fantasy, Isekai, Shounen : string
	// Chapter : 1, 2, 3, 4, 5, 6 ,7, 8, 9, 10  : string

	var judul_komik = [5]string{
		"Blue Lock (Manga)",
		"How to Live as a Villain (Manhwa)",
		"Omniciscient Reader's Viewpoint (Manhwa)",
		"The World After The Fall (Manhwa)",
		"Kill The Hero (Manhwa)",
	}

	var tags = [][6]string{
		{"Action", "Sport", "Shounen"},
		{"Isekai", "Fantasy", "Action", "Adventure", "Magic"},
		{"Isekai", "Fantasy", "Action", "Adventure", "Game", "Magic"},
		{"Isekai", "Fantasy", "Action", "Adventure"},
		{"Isekai", "Fantasy", "Action", "Adventure", "Reincarnation", "Magic"},
	}

	var sinopsis = [5]string{
		// Blue Lock

		"Isagi yang menyukai bola memulai karir nya di Blue Lock yang tingkat Kesulitan nya hingga 99%",

		// How to Live as a Villain

		"Tokoh Utama yang lemah namun licik dan jahat dikirim ke dunia lain sebagai hiburan dewa",

		// Omniciscient Reader's Viewpoint

		"Seorang Pembaca Novel terakhir dari serial dewa menjadi kenyataan. Apakah pembaca bisa mengubah takdir kejam tokoh utama dan membantu mencapai ending yang baru?",

		// The World After The Fall

		"Sebuah tower muncul dari bawah tanah membuat manusia mengalami kepunahan. Mulailah Era Baru di dalam tower yang seperti neraka demi umat manusia",

		// Kill The Hero

		"Hero telah menghianatiku. Aku melakukan regresi ke masalalu demi membunuh Hero Yang telah menghianatiku",
	}

	var Chapter = [][10]string{
		// Blue Lock
		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},

		// How to Live as a Villain

		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},

		// Omniciscient Reader's Viewpoint

		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},

		// The World After The Fall

		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		// Kill The Hero

		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	}

	for i := 0; i < len(judul_komik); i++ {
		fmt.Println("Judul Komik:", judul_komik[i])

		fmt.Println(("Tags"))
		for _, tag := range tags[i] {
			fmt.Println("-", tag)
		}

		fmt.Println("Sinopsis:", sinopsis[i])
		fmt.Println()

		fmt.Println(("Chapter"))
		for _, Chapter := range Chapter[i] {
			fmt.Println("-", Chapter)
		}
	}

}
