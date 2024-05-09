# 设计原则与设计模式

## 1、设计原则

### 1.1、单一职责原则

单一职责原则：类的职责单一，对外只提供一种功能，而且引起类变化的原因只有一个

举个打电话的例子:拨号、通话、回应、挂机

可以设计一个这样的接口

```go
type PhoneCall interface{
    Dial(phoneNumber int64)
    Talk(s struct{})
    Hangup()
}
```

单一职责原则要求一个接口或一个类只能有一个原因引起变化，也就是一个接口或者类只能有一个职责，它就负责一件事情
而 PhoneCall 这个接口做了两件事情：
协议管理和数据传送。Dial 和 Hangup 这两个方法实现的是协议管理，分别负责拨号接通和挂机.
Talk 方法实现的是数据传送。
不管是协议接通的变化还是输出传送的变化，都会引起这个接口的变化。

在面向对象编程的过程中，设计一个类，建议对外提供的功能单一，接口单一
影响一个类的范围就只限定在这一个接口上，一个类的一个接口具备这个类的功能含义，职责单一不复杂

举个例子：

```go
package main

import "fmt"

type workdress struct{}

func (w *workdress)OnWorkDress(){
    fmt.Println("工作的装扮...")
}

type shoppingdress struct{}

func (s *shoppingdress)OnShoppingDress(){
    fmt.Println("逛街的装扮....")
}

```

#### 1.2、开闭原则

开闭原则: 开:对扩展开放 闭：对修改关闭

```go
// 开闭原则
// 对扩展开放
// 对修改关闭

// 平铺式的代码
// type Banker struct{}

// func (b *Banker) Save() { fmt.Println("银行职员进行了存款的业务...") }

// func (b *Banker) Trans() { fmt.Println("银行职员进行了转账的业务....") }

// func (b *Banker) Stack() { fmt.Println("银行职员进行了股票的业务....") }

// 这种平铺的设计  在增加职员的作用的时候 需要修改原有的类
// 这中修改是危险的

// 基于开闭原则的代码设计

type Banker interface {
 DoBuz()
}

type SaveBanker struct{}

func (s *SaveBanker) DoBuz() { fmt.Println("银行职员进行了存款的业务....") }

type TransBanker struct{}

func (t *TransBanker) DoBuz() { fmt.Println("银行职员进行了转账的业务....") }

type StackBanker struct{}

func (s *StackBanker) DoBuz() { fmt.Println("银行职员进行了股票的业务....") }


```

当增加车辆或者司机的时候，不会对原有的代码进行破坏

接口的意义:
实现多态 调用未来

### 1.3、依赖倒转原则

依赖于抽象 而不是依赖于具体的类

高耦合度的代码:

```go
package main

import "fmt"

// === > 奔驰汽车 <===
type Benz struct {

}

func (this *Benz) Run() {
 fmt.Println("Benz is running...")
}

// === > 宝马汽车  <===
type BMW struct {

}

func (this *BMW) Run() {
 fmt.Println("BMW is running ...")
}


//===> 司机张三  <===
type Zhang3 struct {
 //...
}

func (zhang3 *Zhang3) DriveBenZ(benz *Benz) {
 fmt.Println("zhang3 Drive Benz")
 benz.Run()
}

func (zhang3 *Zhang3) DriveBMW(bmw *BMW) {
 fmt.Println("zhang3 drive BMW")
 bmw.Run()
}

//===> 司机李四 <===
type Li4 struct {
 //...
}

func (li4 *Li4) DriveBenZ(benz *Benz) {
 fmt.Println("li4 Drive Benz")
 benz.Run()
}

func (li4 *Li4) DriveBMW(bmw *BMW) {
 fmt.Println("li4 drive BMW")
 bmw.Run()
}

func main() {
 //业务1 张3开奔驰
 benz := &Benz{}
 zhang3 := &Zhang3{}
 zhang3.DriveBenZ(benz)

 //业务2 李四开宝马
 bmw := &BMW{}
 li4 := &Li4{}
 li4.DriveBMW(bmw)
}
```

我们将这种代码改成面对抽象层依赖倒转

```go
package main

import "fmt"

// ===== >   抽象层  < ========
type Car interface {
 Run()
}

type Driver interface {
 Drive(car Car)
}

// ===== >   实现层  < ========
type BenZ struct {
 //...
}

func (benz * BenZ) Run() {
 fmt.Println("Benz is running...")
}

type Bmw struct {
 //...
}

func (bmw * Bmw) Run() {
 fmt.Println("Bmw is running...")
}

type Zhang_3 struct {
 //...
}

func (zhang3 *Zhang_3) Drive(car Car) {
 fmt.Println("Zhang3 drive car")
 car.Run()
}

type Li_4 struct {
 //...
}

func (li4 *Li_4) Drive(car Car) {
 fmt.Println("li4 drive car")
 car.Run()
}


// ===== >   业务逻辑层  < ========
func main() {
 //张3 开 宝马
 var bmw Car
 bmw = &Bmw{}

 var zhang3 Driver
 zhang3 = &Zhang_3{}

 zhang3.Drive(bmw)

 //李4 开 奔驰
 var benz Car
 benz = &BenZ{}

 var li4 Driver
 li4 = &Li_4{}

 li4.Drive(benz)
}
```

### 1.4、里氏代换原则

任何抽象类出现的地方都可以用它的实现类进行替换
基类适用的，子类一定适用（子类可以扩展父类的功能，但不能改变父类原有的功能）

```go
package main

import "fmt"

// 里氏代换原则
// 基类适用的，子类一定适用（子类可以扩展父类的功能，但不能改变父类原有的功能）
// 这个在go里面似乎不太需要这样考虑

type ClassA struct{}

func (ca *ClassA) Add(x int, y int) int {
 return x + y
}

type ClassB struct {
 ClassA
}

func (cb *ClassB) Add(a string, b string) string {
 return a + b
}

func main() {
 a := 10
 b := 20
 cb := ClassB{}
 fmt.Println(cb.ClassA.Add(a, b))
 x := "hello,"
 y := "hi"
 fmt.Println(cb.Add(x, y))
}
```

### 1.5、接口隔离原则

不应该强迫用户的程序依赖他们不需要的接口方法。
一个接口应该只提供一种对外功能，不应该把所有操作都封装到一个接口中去。

```go
package main

// 接口隔离 强调 不要把所有的方法封装到一个接口上
// 将大的接口设计成小的接口
// 不把所有的操作封装到一个接口上

// 假设A类 需要func1 func2 func5
// 假设B类 需要func1 func3 func4

type SomeISP interface {
 Func1()
 Func2()
 Func3()
 Func4()
 Func5()
}

// 类A 需要的是1 2 5 但是它还需要实现3 4
// 类B 需要的是1 3 4 但是它还需要实现2 5

func main() {

}
```

### 1.6、合成复用原则

如果使用继承 会导致父类的任何变换都可能影响到子类，使用组合，则降低了这种
依赖关系，推荐使用组合而不是继承

```go
package main

import "fmt"

// 合成复用原则

type Cat struct{}

func (c *Cat) Eat() {
 fmt.Println("小猫吃饭")
}

// 给小猫添加睡觉的方法
// 通过继承的方式

type CatB struct {
 Cat
}

func (cb *CatB) Sleep() {
 fmt.Println("小猫睡觉")
}

// 给小猫添加睡觉的方法
// 通过组合的方式

type CatC struct {
 Ca Cat
}

func (cc *CatC) Sleep() {
 fmt.Println("小猫睡觉")
}

func main() {
 c:=Cat{}
 c.Eat()

 cb:=CatB{c}
 cb.Sleep()
 cb.Eat() //继承

 // 组合
 cc:=CatC{Ca: c}
 cc.Sleep()
 cc.Ca.Eat() // 组合
}

```

### 1.7、迪米特法则

一个对象应当对其他对象尽可能少的了解，从而降低各个对象之间的耦合，提高系统的可维护性。例如在一个程序中，各个模块之间相互调用时，通常会提供一个统一的接口来实现。这样其他模块不需要了解另外一个模块的内部实现细节，这样当一个模块内部的实现发生改变时，不会影响其他模块的使用。（黑盒原理）

```go
package main

import "fmt"

// 迪米特法则
// 一个对象应该对其他对象尽可能少的了解
// 从而降低各个系统之间的耦合 提高系统的可维护性
// 迪米特法则 又叫最少知道原则
// 如果两个软件实体无须直接通信，那么就不应当发生直接的相互调用，可以通过第三方转发该调用

type Student struct {
 Id   string
 Name string
}

type Class struct {
 Id       string
 Name     string
 Students []Student
}

func (c Class) PrintStudents() {
 for _, student := range c.Students {
  fmt.Println(student)
 }
}

type School struct {
 Id      string
 Name    string
 Classes []Class
}

// 这里 school 与 student 没有直接的关系
func (s School) PrintAllStudents() {

 // 高耦合度的写法
 // for _, class := range s.Classes {
 //  for _, student := range class.Students {
 //   fmt.Println(student)
 //  }
 // }

 // 降低了依赖关系的写法
 for _, class := range s.Classes {
  class.PrintStudents()
 }
}
```

## 2、设计模式

### 2.1、简单工厂模式

如果没有工厂类

```go
package main

import "fmt"

//水果类
type Fruit struct {
 //...
 //...
 //...
}

func (f *Fruit) Show(name string) {
 if name == "apple" {
  fmt.Println("我是苹果")
 } else if name == "banana" {
  fmt.Println("我是香蕉")
 } else if name == "pear" {
  fmt.Println("我是梨")
 }
}

//创建一个Fruit对象
func NewFruit(name string) *Fruit {
 fruit := new(Fruit)

 if name == "apple" {
  //创建apple逻辑
 } else if name == "banana" {
  //创建banana逻辑
 } else if name == "pear" {
  //创建pear逻辑
 }

 return fruit
}

func main() {
 apple := NewFruit("apple")
 apple.Show("apple")

 banana := NewFruit("banana")
 banana.Show("banana")

 pear := NewFruit("pear")
 pear.Show("pear")
}
```

不难看出，Fruit 类是一个“巨大的”类，在该类的设计中存在如下几个问题：
(1) 在 Fruit 类中包含很多“if…else…”代码块，整个类的代码相当冗长，代码越长，阅读难度、维护难度和测试难度也越大；而且大量条件语句的存在还将影响系统的性能，程序在执行过程中需要做大量的条件判断。
(2) Fruit 类的职责过重，它负责初始化和显示所有的水果对象，将各种水果对象的初始化代码和显示代码集中在一个类中实现，违反了“单一职责原则”，不利于类的重用和维护；  
(3) 当需要增加新类型的水果时，必须修改 Fruit 类的构造函数 NewFruit()和其他相关方法源代码，违反了“开闭原则”

简单工厂模式

```go
package main

import "fmt"

// 简单工厂模式
// 工厂（Factory）角色：简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。
// 工厂类可以被外界直接调用，创建所需的产品对象。
// 抽象产品（AbstractProduct）角色：简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
// 具体产品（Concrete Product）角色：简单工厂模式所创建的具体实例对象

// ======= 抽象层 =========

// 水果类(抽象接口)
type Fruit interface {
 Show() //接口的某方法
}

// ======= 基础类模块 =========

type Apple struct {
 Fruit //为了易于理解显示继承(此行可以省略)
}

func (apple *Apple) Show() {
 fmt.Println("我是苹果")
}

type Banana struct {
 Fruit
}

func (banana *Banana) Show() {
 fmt.Println("我是香蕉")
}

type Pear struct {
 Fruit
}

func (pear *Pear) Show() {
 fmt.Println("我是梨")
}

// ========= 工厂模块  =========
// 一个工厂， 有一个生产水果的机器，返回一个抽象水果的指针
type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruit {
 var fruit Fruit

 if kind == "apple" {
  fruit = new(Apple)
 } else if kind == "banana" {
  fruit = new(Banana)
 } else if kind == "pear" {
  fruit = new(Pear)
 }

 return fruit
}

// ==========业务逻辑层==============
func main() {
 factory := new(Factory)

 apple := factory.CreateFruit("apple")
 apple.Show()

 banana := factory.CreateFruit("banana")
 banana.Show()

 pear := factory.CreateFruit("pear")
 pear.Show()
}
```

简单工厂模式的优缺点:

优点：

1. 实现了对象创建和使用的分离。
2. 不需要记住具体类名，记住参数即可，减少使用者记忆量。

缺点：

1. 对工厂类职责过重，一旦不能工作，系统受到影响。
2. 增加系统中类的个数，复杂度和理解度增加。
3. 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂。

适用场景：

1. 工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
2. 客户端只知道传入工厂类的参数，对于如何创建对象并不关心。

### 2.2、工厂模式

工厂模式里的角色和职责

抽象工厂（Abstract Factory）角色：工厂方法模式的核心，任何工厂类都必须实现这个接口。
工厂（Concrete Factory）角色：具体工厂类是抽象工厂的一个实现，负责实例化产品对象。
抽象产品（Abstract Product）角色：工厂方法模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
具体产品（Concrete Product）角色：工厂方法模式所创建的具体实例对象。

```go
package main

import "fmt"

// ======= 抽象层 =========

//水果类(抽象接口)
type Fruit interface {
 Show()  //接口的某方法
}

//工厂类(抽象接口)
type AbstractFactory interface {
 CreateFruit() Fruit //生产水果类(抽象)的生产器方法
}

// ======= 基础类模块 =========
type Apple struct {
 Fruit  //为了易于理解显示继承(此行可以省略)
}

func (apple *Apple) Show() {
 fmt.Println("我是苹果")
}

type Banana struct {
 Fruit
}

func (banana *Banana) Show() {
 fmt.Println("我是香蕉")
}

type Pear struct {
 Fruit
}

func (pear *Pear) Show() {
 fmt.Println("我是梨")
}

// ========= 工厂模块  =========
//具体的苹果工厂
type AppleFactory struct {
 AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
 var fruit Fruit

 //生产一个具体的苹果
 fruit = new(Apple)

 return fruit
}

//具体的香蕉工厂
type BananaFactory struct {
 AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
 var fruit Fruit

 //生产一个具体的香蕉
 fruit = new(Banana)

 return fruit
}


//具体的梨工厂
type PearFactory struct {
 AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
 var fruit Fruit

 //生产一个具体的梨
 fruit = new(Pear)

 return fruit
}

//======= 业务逻辑层 =======
func main() {
 //需求1：需要一个具体的苹果对象
 //1-先要一个具体的苹果工厂
 var appleFac AbstractFactory
 appleFac = new(AppleFactory)
 //2-生产相对应的具体水果
 var apple Fruit
 apple = appleFac.CreateFruit()

 apple.Show()


 //需求2：需要一个具体的香蕉对象
 //1-先要一个具体的香蕉工厂
 var bananaFac AbstractFactory
 bananaFac = new(BananaFactory)
 //2-生产相对应的具体水果
 var banana Fruit
 banana = bananaFac.CreateFruit()

 banana.Show()

 //需求3：需要一个具体的梨对象
 //1-先要一个具体的梨工厂
 var pearFac AbstractFactory
 pearFac = new(PearFactory)
 //2-生产相对应的具体水果
 var pear Fruit
 pear = pearFac.CreateFruit()

 pear.Show()

 //需求4：需要一个日本的苹果？
}
```

当我们需要一个日本的苹果的时候 可以对代码作如下修改:

```go
package main

import "fmt"

// ======= 抽象层 =========

//水果类(抽象接口)
type Fruit interface {
 Show()  //接口的某方法
}

//工厂类(抽象接口)
type AbstractFactory interface {
 CreateFruit() Fruit //生产水果类(抽象)的生产器方法
}

// ======= 基础类模块 =========
type Apple struct {
 Fruit  //为了易于理解显示继承(此行可以省略)
}

func (apple *Apple) Show() {
 fmt.Println("我是苹果")
}

type Banana struct {
 Fruit
}

func (banana *Banana) Show() {
 fmt.Println("我是香蕉")
}

type Pear struct {
 Fruit
}

func (pear *Pear) Show() {
 fmt.Println("我是梨")
}

//(+) 新增一个"日本苹果"
type JapanApple struct {
 Fruit
}

func (jp *JapanApple) Show() {
 fmt.Println("我是日本苹果")
}

// ========= 工厂模块  =========
//具体的苹果工厂
type AppleFactory struct {
 AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
 var fruit Fruit

 //生产一个具体的苹果
 fruit = new(Apple)

 return fruit
}

//具体的香蕉工厂
type BananaFactory struct {
 AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
 var fruit Fruit

 //生产一个具体的香蕉
 fruit = new(Banana)

 return fruit
}


//具体的梨工厂
type PearFactory struct {
 AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
 var fruit Fruit

 //生产一个具体的梨
 fruit = new(Pear)

 return fruit
}

//具体的日本工厂
type JapanAppleFactory struct {
 AbstractFactory
}

func (fac *JapanAppleFactory) CreateFruit() Fruit {
 var fruit Fruit

 //生产一个具体的日本苹果
 fruit = new(JapanApple)

 return fruit
}

// ========= 业务逻辑层  =========
func main() {
 /*
  本案例为了突出根据依赖倒转原则与面向接口编程特性。
     一些变量的定义将使用显示类型声明方式
 */

 //需求1：需要一个具体的苹果对象
 //1-先要一个具体的苹果工厂
 var appleFac AbstractFactory
 appleFac = new(AppleFactory)
 //2-生产相对应的具体水果
 var apple Fruit
 apple = appleFac.CreateFruit()

 apple.Show()


 //需求2：需要一个具体的香蕉对象
 //1-先要一个具体的香蕉工厂
 var bananaFac AbstractFactory
 bananaFac = new(BananaFactory)
 //2-生产相对应的具体水果
 var banana Fruit
 banana = bananaFac.CreateFruit()

 banana.Show()

 //需求3：需要一个具体的梨对象
 //1-先要一个具体的梨工厂
 var pearFac AbstractFactory
 pearFac = new(PearFactory)
 //2-生产相对应的具体水果
 var pear Fruit
 pear = pearFac.CreateFruit()

 pear.Show()

 //需求4：需要一个日本的苹果？
 //1-先要一个具体的日本评估工厂
 var japanAppleFac AbstractFactory
 japanAppleFac = new(JapanAppleFactory)
 //2-生产相对应的具体水果
 var japanApple Fruit
 japanApple = japanAppleFac.CreateFruit()

 japanApple.Show()
}

```

工厂方法模式的优缺点:

优点：

1. 不需要记住具体类名，甚至连具体参数都不用记忆。
2. 实现了对象创建和使用的分离。
3. 系统的可扩展性也就变得非常好，无需修改接口和原类。 4.对于新产品的创建，符合开闭原则。

缺点：

1. 增加系统中类的个数，复杂度和理解度增加。
2. 增加了系统的抽象性和理解难度。

适用场景：

1. 客户端不知道它所需要的对象的类。
2. 抽象工厂类通过其子类来指定创建哪个对象。

### 2.3、抽象工厂模式

（1）当添加一个新产品的时候，比如葡萄，虽然不用修改代码，但是需要添加大量的类，而且还需要添加相对的工厂。（系统开销，维护成本）
（2）如果使用同一地域的水果（日本苹果，日本香蕉，日本梨），那么需要分别创建具体的工厂，如果选择出现失误，将会造成混乱，虽然可以加一些约束，但是代码实现变得复杂。
所以“抽象工厂方法模式”引出了“产品族”和“产品等级结构”概念，其目的是为了更加高效的生产同一个产品组产品。

抽象工厂的角色和职责:

抽象工厂（Abstract Factory）角色：它声明了一组用于创建一族产品的方法，每一个方法对应一种产品。
具体工厂（Concrete Factory）角色：它实现了在抽象工厂中声明的创建产品的方法，生成一组具体产品，这些产品构成了一个产品族，每一个产品都位于某个产品等级结构中。
抽象产品（Abstract Product）角色：它为每种产品声明接口，在抽象产品中声明了产品所具有的业务方法。
具体产品（Concrete Product）角色：它定义具体工厂生产的具体产品对象，实现抽象产品接口中声明的业务方法。

```go
package main

import "fmt"

// ======= 抽象层 =========
type AbstractApple interface {
 ShowApple()
}

type AbstractBanana interface {
 ShowBanana()
}

type AbstractPear interface {
 ShowPear()
}

//抽象工厂
type AbstractFactory interface {
 CreateApple() AbstractApple
 CreateBanana() AbstractBanana
 CreatePear() AbstractPear
}

// ======== 实现层 =========
/*  中国产品族 */
type ChinaApple struct {}

func (ca *ChinaApple) ShowApple() {
 fmt.Println("中国苹果")
}

type ChinaBanana struct {}

func (cb *ChinaBanana) ShowBanana() {
 fmt.Println("中国香蕉")
}

type ChinaPear struct {}

func (cp *ChinaPear) ShowPear() {
 fmt.Println("中国梨")
}

type ChinaFactory struct {}

func (cf *ChinaFactory) CreateApple() AbstractApple {
 var apple AbstractApple

 apple = new(ChinaApple)

 return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
 var banana AbstractBanana

 banana = new(ChinaBanana)

 return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
 var pear AbstractPear

 pear = new(ChinaPear)

 return pear
}

/*  日本产品族 */
type JapanApple struct {}

func (ja *JapanApple) ShowApple() {
 fmt.Println("日本苹果")
}

type JapanBanana struct {}

func (jb *JapanBanana) ShowBanana() {
 fmt.Println("日本香蕉")
}

type JapanPear struct {}

func (cp *JapanPear) ShowPear() {
 fmt.Println("日本梨")
}

type JapanFactory struct {}

func (jf *JapanFactory) CreateApple() AbstractApple {
 var apple AbstractApple

 apple = new(JapanApple)

 return apple
}

func (jf *JapanFactory) CreateBanana() AbstractBanana {
 var banana AbstractBanana

 banana = new(JapanBanana)

 return banana
}

func (cf *JapanFactory) CreatePear() AbstractPear {
 var pear AbstractPear

 pear = new(JapanPear)

 return pear
}

/*  美国产品族 */
type AmericanApple struct {}

func (aa *AmericanApple) ShowApple() {
 fmt.Println("美国苹果")
}

type AmericanBanana struct {}

func (ab *AmericanBanana) ShowBanana() {
 fmt.Println("美国香蕉")
}

type AmericanPear struct {}

func (ap *AmericanPear) ShowPear() {
 fmt.Println("美国梨")
}

type AmericanFactory struct {}

func (af *AmericanFactory) CreateApple() AbstractApple {
 var apple AbstractApple

 apple = new(AmericanApple)

 return apple
}

func (af *AmericanFactory) CreateBanana() AbstractBanana {
 var banana AbstractBanana

 banana = new(AmericanBanana)

 return banana
}

func (af *AmericanFactory) CreatePear() AbstractPear {
 var pear AbstractPear

 pear = new(AmericanPear)

 return pear
}

// ======== 业务逻辑层 =======
func main() {
 //需求1: 需要美国的苹果、香蕉、梨 等对象
 //1-创建一个美国工厂
 var aFac AbstractFactory
 aFac = new(AmericanFactory)

 //2-生产美国苹果
 var aApple AbstractApple
 aApple = aFac.CreateApple()
 aApple.ShowApple()

 //3-生产美国香蕉
 var aBanana AbstractBanana
 aBanana = aFac.CreateBanana()
 aBanana.ShowBanana()

 //4-生产美国梨
 var aPear AbstractPear
 aPear = aFac.CreatePear()
 aPear.ShowPear()

 //需求2: 需要中国的苹果、香蕉
 //1-创建一个中国工厂
 cFac := new(ChinaFactory)

 //2-生产中国苹果
 cApple := cFac.CreateApple()
 cApple.ShowApple()

 //3-生产中国香蕉
 cBanana := cFac.CreateBanana()
 cBanana.ShowBanana()
}

```

### 2.4、单例模式

在单例类的内部实现只生成一个实例，同时它提供一个静态的 getInstance()工厂方法，
让客户可以访问它的唯一实例；为了防止在外部对其实例化，将其构造函数设计为私有；
在单例类内部定义了一个 Singleton 类型的静态对象，作为外部共享的唯一实例。

单例解决的问题:

保证一个类永远只能有一个对象，且该对象的功能依然能被其他模块使用。

```go
package main

import "fmt"

/*
三个要点：
  一是某个类只能有一个实例；
  二是它必须自行创建这个实例；
  三是它必须自行向整个系统提供这个实例。
*/

/*
 保证一个类永远只能有一个对象
*/


//1、保证这个类非公有化，外界不能通过这个类直接创建一个对象
//   那么这个类就应该变得非公有访问 类名称首字母要小写
type singelton struct {}

//2、但是还要有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向
//   Golang中没有常指针概念，所以只能通过将这个指针私有化不让外部模块访问
var instance *singelton = new(singelton)

//3、如果全部为私有化，那么外部模块将永远无法访问到这个类和对象，
//   所以需要对外提供一个方法来获取这个唯一实例对象
//   注意：这个方法是否可以定义为singelton的一个成员方法呢？
//       答案是不能，因为如果为成员方法就必须要先访问对象、再访问函数
//        但是类和对象目前都已经私有化，外界无法访问，所以这个方法一定是一个全局普通函数
func GetInstance() *singelton {
 return instance
}

func (s *singelton) SomeThing() {
 fmt.Println("单例对象的某方法")
}

func main() {
 s := GetInstance()
 s.SomeThing()
}
```

这种单例 属于懒汉式单例 不用考虑并发安全问题
缺点是 无论使用不使用 都会创建内存

饿汉式单例:

```go
package main

import (
 "fmt"
 "sync"
)

//定义锁
var lock sync.Mutex

type singelton struct {}

var instance *singelton

func GetInstance() *singelton {
 //为了线程安全，增加互斥
 lock.Lock()
 defer lock.Unlock()

 if instance == nil {
  return new(singelton)
 } else {
  return instance
 }
}

func (s *singelton) SomeThing() {
 fmt.Println("单例对象的某方法")
}


func main() {
 s := GetInstance()
 s.SomeThing()
}
```

### 2.5、代理模式

是指具有与代理元（被代理的对象）具有相同的接口的类，客户端必须通过代理与被代理的目标类交互，而代理一般在交互的过程中（交互前后），进行某些特别的处理
这里假设有一个“自己”的角色，正在玩一款网络游戏。称这个网络游戏就是代理模式的“Subject”，表示要做一件事的目标或者对象事件主题。
（1）“自己”有一个给游戏角色升级的需求或者任务，当然“自己”可以独自完成游戏任务的升级。
（2）或者“自己”也可以邀请以为更加擅长游戏的“游戏代练”来完成升级这件事，这个代练就是“Proxy”代理。
（3）“游戏代练”不仅能够完成升级的任务需求，还可以额外做一些附加的能力。比如打到一些好的游戏装备、加入公会等等周边收益。
所以代理的出现实则是为了能够覆盖“自己”的原本的需求，且可以额外做其他功能，这种额外创建的类是不影响已有的“自己”和“网络游戏”的的关系。是额外添加，在设计模式原则上，是符合“开闭原则”思想。那么当需要给“自己”增加额外功能的时候，又不想改变自己，那么就选择邀请一位”代理”来完成吧。

代理的类结构:

- subject（抽象主题角色）：真实主题与代理主题的共同接口。
- RealSubject（真实主题角色）：定义了代理角色所代表的真实对象。
- Proxy（代理主题角色）：含有对真实主题角色的引用，代理角色通常在将客户端调用传递给真是主题对象之前或者之后执行某些操作，而不是单纯返回真实的对象。

```go
package main

import "fmt"

type Goods struct {
 Kind string   //商品种类
 Fact bool   //商品真伪
}

// =========== 抽象层 ===========
//抽象的购物主题Subject
type Shopping interface {
 Buy(goods *Goods) //某任务
}


// =========== 实现层 ===========
//具体的购物主题， 实现了shopping， 去韩国购物
type KoreaShopping struct {}

func (ks *KoreaShopping) Buy(goods *Goods) {
 fmt.Println("去韩国进行了购物, 买了 ", goods.Kind)
}


//具体的购物主题， 实现了shopping， 去美国购物
type AmericanShopping struct {}

func (as *AmericanShopping) Buy(goods *Goods) {
 fmt.Println("去美国进行了购物, 买了 ", goods.Kind)
}

//具体的购物主题， 实现了shopping， 去非洲购物
type AfrikaShopping struct {}

func (as *AfrikaShopping) Buy(goods *Goods) {
 fmt.Println("去非洲进行了购物, 买了 ", goods.Kind)
}


//海外的代理
type OverseasProxy struct {
 shopping Shopping //代理某个主题，这里是抽象类型
}

func (op *OverseasProxy) Buy(goods *Goods) {
 // 1. 先验货
 if (op.distinguish(goods) == true) {
  //2. 进行购买
  op.shopping.Buy(goods) //调用原被代理的具体主题任务
  //3 海关安检
  op.check(goods)
 }
}

//创建一个代理,并且配置关联被代理的主题
func NewProxy(shopping Shopping) Shopping {
 return &OverseasProxy{shopping}
}

//验货流程
func (op *OverseasProxy) distinguish(goods *Goods) bool {
 fmt.Println("对[", goods.Kind,"]进行了辨别真伪.")
 if (goods.Fact == false) {
  fmt.Println("发现假货",goods.Kind,", 不应该购买。")
 }
 return goods.Fact
}

//安检流程
func (op *OverseasProxy) check(goods *Goods) {
 fmt.Println("对[",goods.Kind,"] 进行了海关检查， 成功的带回祖国")
}


func main() {
 g1 := Goods{
  Kind: "韩国面膜",
  Fact: true,
 }

 g2 := Goods{
  Kind: "CET4证书",
  Fact: false,
 }

 //如果不使用代理来完成从韩国购买任务
 var shopping Shopping
 shopping = new(KoreaShopping) //具体的购买主题

 //1-先验货
 if g1.Fact == true {
  fmt.Println("对[", g1.Kind,"]进行了辨别真伪.")
  //2-去韩国购买
  shopping.Buy(&g1)
  //3-海关安检
  fmt.Println("对[",g1.Kind,"] 进行了海关检查， 成功的带回祖国")
 }

 fmt.Println("---------------以下是 使用 代理模式-------")
 var overseasProxy Shopping
 overseasProxy = NewProxy(shopping)
 overseasProxy.Buy(&g1)
 overseasProxy.Buy(&g2)
}

```

### 2.6、装饰模式

装饰模式(Decorator Pattern)：动态地给一个对象增加一些额外的职责，就增加对象功能来说，装饰模式比生成子类实现更为灵活。装饰模式是一种对象结构型模式。

Component（抽象构件）：它是具体构件和抽象装饰类的共同父类，声明了在具体构件中实现的业务方法，
它的引入可以使客户端以一致的方式处理未被装饰的对象以及装饰之后的对象，实现客户端的透明操作。

ConcreteComponent（具体构件）：它是抽象构件类的子类，用于定义具体的构件对象，实现了在抽象构件中声明的方法，装饰器可以给它增加额外的职责（方法）

```go
package main

import "fmt"

// ---------- 抽象层 ----------
//抽象的构件
type Phone interface {
 Show() //构件的功能
}

//装饰器基础类（该类本应该为interface，但是Golang interface语法不可以有成员属性）
type Decorator struct {
 phone Phone
}

func (d *Decorator) Show() {}


// ----------- 实现层 -----------
// 具体的构件
type HuaWei struct {}

func (hw *HuaWei) Show() {
 fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct{}

func (xm *XiaoMi) Show() {
 fmt.Println("秀出了XiaoMi手机")
}

// 具体的装饰器类
type MoDecorator struct {
 Decorator   //继承基础装饰器类(主要继承Phone成员属性)
}

func (md *MoDecorator) Show() {
 md.phone.Show() //调用被装饰构件的原方法
 fmt.Println("贴膜的手机") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
 return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
 Decorator   //继承基础装饰器类(主要继承Phone成员属性)
}

func (kd *KeDecorator) Show() {
 kd.phone.Show()
 fmt.Println("手机壳的手机") //装饰额外的方法
}

func NewKeDecorator(phone Phone) Phone {
 return &KeDecorator{Decorator{phone}}
}


// ------------ 业务逻辑层 ---------
func main() {
 var huawei Phone
 huawei = new(HuaWei)
 huawei.Show()  //调用原构件方法

 fmt.Println("---------")
 //用贴膜装饰器装饰，得到新功能构件
 var moHuawei Phone
 moHuawei = NewMoDecorator(huawei) //通过HueWei ---> MoHuaWei
 moHuawei.Show() //调用装饰后新构件的方法

 fmt.Println("---------")
 var keHuawei Phone
 keHuawei = NewKeDecorator(huawei) //通过HueWei ---> KeHuaWei
 keHuawei.Show()

 fmt.Println("---------")
 var keMoHuaWei Phone
 keMoHuaWei = NewMoDecorator(keHuawei) //通过KeHuaWei ---> KeMoHuaWei
 keMoHuaWei.Show()
}
```

### 2.7、适配器模式

将一个类的接口转换成客户希望的另外一个接口。使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

Target（目标抽象类）：目标抽象类定义客户所需接口，可以是一个抽象类或接口，也可以是具体类。
Adapter（适配器类）：适配器可以调用另一个接口，作为一个转换器，对 Adaptee 和 Target 进行适配，适配器类是适配器模式的核心，在对象适配器中，它通过继承 Target 并关联一个 Adaptee 对象使二者产生联系。
Adaptee（适配者类）：适配者即被适配的角色，它定义了一个已经存在的接口，这个接口需要适配，适配者类一般是一个具体类，包含了客户希望使用的业务方法，在某些情况下可能没有适配者类的源代码。

```go
package main

import "fmt"

//适配的目标
type V5 interface {
 Use5V()
}

//业务类，依赖V5接口
type Phone struct {
 v V5
}

func NewPhone(v V5) *Phone {
 return &Phone{v}
}

func (p *Phone) Charge() {
 fmt.Println("Phone进行充电...")
 p.v.Use5V()
}


//被适配的角色，适配者
type V220 struct {}

func (v *V220) Use220V() {
 fmt.Println("使用220V的电压")
}

//电源适配器
type Adapter struct {
 v220 *V220
}

func (a *Adapter) Use5V() {
 fmt.Println("使用适配器进行充电")

 //调用适配者的方法
 a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
 return &Adapter{v220}
}



// ------- 业务逻辑层 -------
func main() {
 iphone := NewPhone(NewAdapter(new(V220)))

 iphone.Charge()
}
```

### 2.8、外观模式

根据迪米特法则，如果两个类不必彼此直接通信，那么这两个类就不应当发生直接的相互作用
Facade 模式也叫外观模式，是由 GoF 提出的 23 种设计模式中的一种。Facade 模式为一组具有类似功能的类群，
比如类库，子系统等等，提供一个一致的简单的界面。这个一致的简单的界面被称作 facade。

**外观模式中的角色与职责**
Facade(外观角色)：为调用方, 定义简单的调用接口。
SubSystem(子系统角色)：功能提供者。指提供功能的类群（模块或子系统）

```go
package main

import "fmt"

type SubSystemA struct {}

func (sa *SubSystemA) MethodA() {
 fmt.Println("子系统方法A")
}

type SubSystemB struct {}

func (sb *SubSystemB) MethodB() {
 fmt.Println("子系统方法B")
}

type SubSystemC struct {}

func (sc *SubSystemC) MethodC() {
 fmt.Println("子系统方法C")
}

type SubSystemD struct {}

func (sd *SubSystemD) MethodD() {
 fmt.Println("子系统方法D")
}

//外观模式，提供了一个外观类， 简化成一个简单的接口供使用
type Facade struct {
 a *SubSystemA
 b *SubSystemB
 c *SubSystemC
 d *SubSystemD
}

func (f *Facade) MethodOne() {
 f.a.MethodA()
 f.b.MethodB()
}


func (f *Facade) MethodTwo() {
 f.c.MethodC()
 f.d.MethodD()
}

func main() {
 //如果不用外观模式实现MethodA() 和 MethodB()
 sa := new(SubSystemA)
 sa.MethodA()
 sb := new(SubSystemB)
 sb.MethodB()

 fmt.Println("-----------")
 //使用外观模式
 f := Facade{
  a: new(SubSystemA),
  b: new(SubSystemB),
  c: new(SubSystemC),
  d: new(SubSystemD),
 }

 //调用外观包裹方法
 f.MethodOne()
}
```

### 2.9、模板方法模式

AbstractClass（抽象类）：在抽象类中定义了一系列基本操作(PrimitiveOperations)，这些基本操作可以是具体的，也可以是抽象的，每一个基本操作对应算法的一个步骤，在其子类中可以重定义或实现这些步骤。同时，在抽象类中实现了一个模板方法(Template Method)，用于定义一个算法的框架，模板方法不仅可以调用在抽象类中实现的基本方法，也可以调用在抽象类的子类中实现的基本方法，还可以调用其他对象中的方法。
ConcreteClass（具体子类）：它是抽象类的子类，用于实现在父类中声明的抽象基本操作以完成子类特定算法的步骤，也可以覆盖在父类中已经实现的具体基本操作。

### 2.10、命令模式

将一个请求封装为一个对象，从而让我们可用不同的请求对客户进行参数化；对请求排队或者记录请求日志，以及支持可撤销的操作。命令模式是一种对象行为型模式，其别名为动作(Action)模式或事务(Transaction)模式。

命令模式可以将请求发送者和接收者完全解耦，发送者与接收者之间没有直接引用关系，发送请求的对象只需要知道如何发送请求，而不必知道如何完成请求。

Command（抽象命令类）：抽象命令类一般是一个抽象类或接口，在其中声明了用于执行请求的 execute()等方法，通过这些方法可以调用请求接收者的相关操作。

ConcreteCommand（具体命令类）：具体命令类是抽象命令类的子类，实现了在抽象命令类中声明的方法，它对应具体的接收者对象，将接收者对象的动作绑定其中。在实现 execute()方法时，将调用接收者对象的相关操作(Action)。

Invoker（调用者）：调用者即请求发送者，它通过命令对象来执行请求。一个调用者并不需要在设计时确定其接收者，因此它只与抽象命令类之间存在关联关系。在程序运行时可以将一个具体命令对象注入其中，再调用具体命令对象的 execute()方法，从而实现间接调用请求接收者的相关操作。

Receiver（接收者）：接收者执行与请求相关的操作，它具体实现对请求的业务处理。

### 2.10、策略模式

策略模式的角色:

Context (环境类):环境类时使用算法的角色,它在解决某个问题(即实现某个方法)时可以采用多种策略，在环境类种维护一个对抽象策略类的引用实例，用于定义所采用的策略

Strategy (抽象策略类)：它为所支持的算法声明了抽象方法，是所有策略类的父类，它可以是抽象类或者具体类，也可以是接口。环境类通过抽象策略类中声明的方法运行时调用具体的策略类中实现算法

ConcreteStrategy(具体策略类):它实现了在抽象类中声明的算法，在运行时，具体策略类将覆盖在环境类中定义的抽象策略类对象，使用一种具体的算法实现某个业务处理

策略模式的有缺点:

优点：
(1) 策略模式提供了对“开闭原则”的完美支持，用户可以在不修改原有系统的基础上选择算法或行为，也可以灵活地增加新的算法或行为。
(2) 使用策略模式可以避免多重条件选择语句。多重条件选择语句不易维护，它把采取哪一种算法或行为的逻辑与算法或行为本身的实现逻辑混合在一起，将它们全部硬编码(Hard Coding)在一个庞大的多重条件选择语句中，比直接继承环境类的办法还要原始和落后。
(3) 策略模式提供了一种算法的复用机制。由于将算法单独提取出来封装在策略类中，因此不同的环境类可以方便地复用这些策略类。

缺点：
(1) 客户端必须知道所有的策略类，并自行决定使用哪一个策略类。这就意味着客户端必须理解这些算法的区别，以便适时选择恰当的算法。换言之，策略模式只适用于客户端知道所有的算法或行为的情况。
(2) 策略模式将造成系统产生很多具体策略类，任何细小的变化都将导致系统要增加一个新的具体策略类。

### 2.11、观察者模式

观察者模式是用于建立一种对象与对象之间的依赖关系，一个对象发生改变时将自动通知其他对象，其他对象将相应作出反应。
在观察者模式中，发生改变的对象称为观察目标，而被通知的对象称为观察者，一个观察目标可以对应多个观察者，
而且这些观察者之间可以没有任何相互联系，可以根据需要增加和删除观察者，使得系统更易于扩展。

Subject（被观察者或目标，抽象主题）：被观察的对象。当需要被观察的状态发生变化时，需要通知队列中所有观察者对象。Subject 需要维持（添加，删除，通知）一个观察者对象的队列列表。
ConcreteSubject（具体被观察者或目标，具体主题）：被观察者的具体实现。包含一些基本的属性状态及其他操作。
Observer（观察者）：接口或抽象类。当 Subject 的状态发生变化时，Observer 对象将通过一个 callback 函数得到通知。
ConcreteObserver（具体观察者）：观察者的具体实现。得到通知后将完成一些具体的业务逻辑处理。
