package main

import (
	"fmt"
)

//显示切片的长度，容量信息
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	fmt.Println("lesson12 切片的扩展")

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	//使用append函数为切片增加元素
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs) //	[Ceres Pluto Haumea Makemake Eris Orcus]

	//特殊情况：如果对切片后追加元素，会替换原数组中的元素
	dwarfsSlice1 := dwarfs[2:4]
	dwarfsSlice1 = append(dwarfsSlice1, "EEEE")
	fmt.Println(dwarfsSlice1) //[Haumea Makemake EEEE]
	fmt.Println(dwarfs)       //[Ceres Pluto Haumea Makemake EEEE Orcus]

	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs) //[Ceres Pluto Haumea Makemake EEEE Orcus Salacia Quaoar Sedna]

	//研究切片的长度和容量
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	dump("planets", planets)           //planets: length 8, capacity 8 [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
	dump("planets[1:4]", planets[1:4]) //planets[1:4]: length 3, capacity 7 [Venus Earth Mars]
	dump("planets[2:5]", planets[2:5]) //planets[2:5]: length 3, capacity 6 [Earth Mars Jupiter]
	dump("planets[5:7]", planets[5:7]) //planets[5:7]: length 2, capacity 3 [Saturn Uranus]

	//探究append函数
	dwarfsRaw := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs1 := dwarfsRaw[:]                                  //length=5， capacity=5
	dwarfs2 := append(dwarfs1, "Orcus")                      //length=6， capacity=10
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna") //length=9， capacity=10

	dump("dwarfs1", dwarfs1) //dwarfs1: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]
	dump("dwarfs2", dwarfs2) //dwarfs2: length 6, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus]
	dump("dwarfs3", dwarfs3) //dwarfs3: length 9, capacity 10 [Ceres Pluto Haumea Makemake Eris Orcus Salacia Quaoar Sedna]
	fmt.Println(dwarfsRaw)   //[Ceres Pluto Haumea Makemake Eris]

	//修改dwarfs3中的元素，查看dwarfs1和dwarfs2中的影响
	dwarfs3[1] = "A"
	dump("dwarfs1", dwarfs1) //dwarfs1: length 5, capacity 5 [Ceres Pluto Haumea Makemake Eris]
	dump("dwarfs2", dwarfs2) //dwarfs2: length 6, capacity 10 [Ceres A Haumea Makemake Eris Orcus]
	fmt.Println(dwarfsRaw)   //[Ceres Pluto Haumea Makemake Eris]

	//planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	//使用三索引方式来声明切片
	terrestrial := planets[0:4:4]
	terrestrial1 := planets[0:4]
	dump("terrestrial", terrestrial)   //length 4, capacity 4 [Mercury Venus Earth Mars]
	dump("terrestrial1", terrestrial1) //length 4, capacity 8 [Mercury Venus Earth Mars]

	worlds := append(terrestrial, "Ceres")

	dump("planets", planets)         //planets: length 8, capacity 8 [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
	dump("terrestrial", terrestrial) //terrestrial: length 4, capacity 4 [Mercury Venus Earth Mars]
	dump("worlds", worlds)           //worlds: length 5, capacity 8 [Mercury Venus Earth Mars Ceres]

	//使用make函数
	dwarfsWithMake := make([]string, 0, 10)
	dwarfsWithMake = append(dwarfsWithMake, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	dump("dwarfsWithMake", dwarfsWithMake) //dwarfsWithMake: length 5, capacity 10 [Ceres Pluto Haumea Makemake Eris]

	//调用“可变参数函数”【方式1】
	twoWorld := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorld) //[New Venus New Mars]

	//调用“可变参数函数”【方式2】
	oldWorlds := []string{"Venus", "Mars", "Jupiter"}
	newWorld := terraform("New", oldWorlds...)
	fmt.Println(newWorld) //[New Venus New Mars New Jupiter]
}

//声明一个新的切片，切片的内容是将传入的切片元素前加上前缀，前缀是该函数的第一个参数
func terraform(prefix string, worlds ...string) []string {

	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
