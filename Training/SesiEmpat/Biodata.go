package SesiEmpat

import (
	"fmt"
	"os"
	"strconv"
)

func ExcerciseVersion1() {
	student1 := Student{"Heru1", "Bekasi", "IT", "Mengikuti Perkembangan Teknologi"}
	student2 := Student{"Heru2", "Bekasi", "IT", "Mengikuti Perkembangan Teknologi"}
	student3 := Student{"Heru3", "Bekasi", "IT", "Mengikuti Perkembangan Teknologi"}
	student4 := Student{"Heru4", "Bekasi", "IT", "Mengikuti Perkembangan Teknologi"}

	index := os.Args[1]
	switch index {
	case "1":
		fmt.Println("Nama saya adalah ", student1.Nama)
		fmt.Println("Alamat saya di ", student1.Alamat)
		fmt.Println("Pekerjaan saya ", student1.Pekerjaan)
		fmt.Println("Saya ikut kelas golang karena ", student1.Alasan)
	case "2":
		fmt.Println("Nama saya adalah ", student2.Nama)
		fmt.Println("Alamat saya di ", student2.Alamat)
		fmt.Println("Pekerjaan saya ", student2.Pekerjaan)
		fmt.Println("Saya ikut kelas golang karena ", student2.Alasan)
	case "3":
		fmt.Println("Nama saya adalah ", student3.Nama)
		fmt.Println("Alamat saya di ", student3.Alamat)
		fmt.Println("Pekerjaan saya ", student3.Pekerjaan)
		fmt.Println("Saya ikut kelas golang karena ", student3.Alasan)
	case "4":
		fmt.Println("Nama saya adalah ", student4.Nama)
		fmt.Println("Alamat saya di ", student4.Alamat)
		fmt.Println("Pekerjaan saya ", student4.Pekerjaan)
		fmt.Println("Saya ikut kelas golang karena ", student4.Alasan)
	default:
		fmt.Println("index 1 - 4")
	}
}

func ExcerciseVersion2() {

	var arrayStudent = []Student{
		Student{
			Nama:      "Heru1",
			Alamat:    "Bekasi",
			Pekerjaan: "IT",
			Alasan:    "Coba-coba",
		},
		Student{
			Nama:      "Heru2",
			Alamat:    "Bekasi",
			Pekerjaan: "IT",
			Alasan:    "Coba-coba",
		},
		Student{
			Nama:      "Heru3",
			Alamat:    "Bekasi",
			Pekerjaan: "IT",
			Alasan:    "Coba-coba",
		},
		Student{
			Nama:      "Heru4",
			Alamat:    "Bekasi",
			Pekerjaan: "IT",
			Alasan:    "Coba-coba",
		},
	}

	var Studentx = Student{
		Nama:      "Heru5",
		Alamat:    "Bekasi",
		Pekerjaan: "IT",
		Alasan:    "Coba-coba",
	}

	arrayStudent = append(arrayStudent, Studentx)

	if len(os.Args) < 2 {
		fmt.Println("Masukkan input numeric")
		return
	}

	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Input harus numeric")
		return
	}

	index = index - 1

	if index >= 0 && index < 5 {
		fmt.Println("Nama saya adalah ", arrayStudent[index].Nama)
		fmt.Println("Alamat saya di ", arrayStudent[index].Alamat)
		fmt.Println("Pekerjaan saya ", arrayStudent[index].Pekerjaan)
		fmt.Println("Saya ikut kelas golang karena ", arrayStudent[index].Alasan)
	} else {
		fmt.Println("Input 1 - 5")
	}

}
