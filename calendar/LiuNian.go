package calendar

import "github.com/6tail/lunar-go/LunarUtil"

type LiuNian struct {
	// 序数，0-9
	index int
	// 大运
	daYun *DaYun
	// 年
	year int
	// 年龄
	age   int
	lunar *Lunar
}

func NewLiuNian(daYun *DaYun, index int) *LiuNian {
	liuNian := new(LiuNian)
	liuNian.daYun = daYun
	liuNian.lunar = daYun.GetLunar()
	liuNian.index = index
	liuNian.year = daYun.GetStartYear() + liuNian.index
	liuNian.age = daYun.GetStartAge() + liuNian.index
	return liuNian
}

func (liuNian *LiuNian) GetIndex() int {
	return liuNian.index
}

func (liuNian *LiuNian) GetYear() int {
	return liuNian.year
}

func (liuNian *LiuNian) GetAge() int {
	return liuNian.age
}

// 获取干支
func (liuNian *LiuNian) GetGanZhi() string {
	offset := LunarUtil.GetJiaZiIndex(liuNian.lunar.GetYearInGanZhiExact()) + liuNian.index
	if liuNian.daYun.GetIndex() > 0 {
		offset += liuNian.daYun.GetStartAge() - 1
	}
	offset %= len(LunarUtil.JIA_ZI)
	return LunarUtil.JIA_ZI[offset]
}

func (liuNian *LiuNian) GetLiuYue() []*LiuYue {
	n := 12
	l := make([]*LiuYue, n)
	for i := 0; i < n; i++ {
		l[i] = NewLiuYue(liuNian, i)
	}
	return l
}
