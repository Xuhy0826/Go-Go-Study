package counters

//声明了未公开的类型
type alertCounter int

// New 创建并返回一个未公开的alertCounter 类型的值
func New(value int) alertCounter {
	return alertCounter(value)
}
