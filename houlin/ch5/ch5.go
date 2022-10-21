package main

/*
	第五章：同步

	#5.1.1

	互斥锁：传统并发程序对共享资源，进行访问控制的主要手段。
	它由标准库代码包 sync 中的 Mutex 结构类型表示。
	sync.Mutex 类型只有两个公开的指针方法——Lock 和 Unlock。

	sync.Mutex 类型的零值表示未被锁定的互斥量。
	声明： var mutex sync.Mutex

	其他编程语言（如：C、Java）的锁类工具时，可能会犯一个低级错误：忘记及时解开已被锁住的锁，从而导致：流程执行异常、线程执行停滞、程序死锁等系列问题。
	在 Go 中，这个低级错误的发生率极低，主要原因时存在 defer 语句。

	惯用法是：在锁定互斥锁后，紧接着用 defer 语句保证该互斥锁的及时解锁。如下代码片段：

	var mutex sync.Mutex

	func write() {
		mutex.Lock()
		defer mutex.Unlock()
		// 省略若干代码
	}
	defer 语句保证了在该函数执行结束之前，互斥锁 mutex 一定会被解锁，
	因此省去了在所有 return 语句之前，以及异常发生之时重复的附加解锁操作。
	在函数的内部执行流程很复杂的情况下，该工作量不容忽视，且极易遗漏。

	对同一互斥锁的锁定操作和解锁操作要成对出现。
	如果锁定了一个已锁定的互斥锁，那进行重复锁定操作的 goroutine 将被阻塞，直到该互斥锁回到解锁状态。

	互斥锁的特性：该锁被锁后，只要还未被释放，则无法再次被锁。

	互斥锁锁定操作的逆操作，并不会引起任何 goroutine 的阻塞，但可能引发一个无法恢复的运行时恐慌。

	对未锁定的互斥锁进行解锁操作时，就会引发一个运行时恐慌。
	避免发生该情况的最简单、有效的方式是使用 defer 语句。

	强烈建议：把对同一个互斥锁的锁定和解锁操作，放在同一个层次的代码块中。

	例如：（1）在同一个函数或方法中对某个互斥锁进行锁定和解锁。
	（2）把互斥锁作为某一个结构体类型中的字段，以便该类型的多个方法可使用到。
	代表互斥锁的变量的作用域尽量小。以免在不相关流程中被误用。

	#5.1.2

	读写锁，即针对读写操作的互斥锁。与普通的互斥锁最大的不同，就是可以分别正对读操作和写操作进行锁定和解锁操作。
	读写锁控制下的多个写操作之间都是互斥的，并且写操作与读操作之间也都是互斥的。
	但多个读操作之间却不存在互斥关系。

	Go 中的读写锁由结构体类型  sync.RWMutex 表示。
	sync.RWMutex 类型的零值就是已经可用的读写锁实例了。
	此类型的方法集合中包含两对方法，即：
	func (*RWMutex) Lock() // 写操作-写锁定
	func (*RWMutex) Unlock() // 写操作-写解锁
	和
	func (*RWMutex) RLock() // 读操作-读锁定
	func (*RWMutex) RUnlock() // 读操作-读解锁


	#5.2 条件变量

	Go 标准库中的 sync.Cond 类型代表了条件变量。

	互斥锁的声明和创建：var mutex sync.Mutex
	读写锁的声明和创建：var rwm sync.RWMutex

	条件变量的声明和创建，需用到 sync.NewCond 函数，该函数声明如下：
	func NewCond(l locker) *Cond

	条件变量总要与互斥量组合使用。sync.NewCond 函数的唯一参数是 sync.Locker 类型的，
	而具体的参数值既可是一个互斥锁(sync.Mutex)，也可为一个读写锁(sync.RWMutex)。

	sync.NewCond 函数在被调用之后，会返回一个 *sync.Cond 类型的结果值，可调用该值拥有的几个方法来操纵这个条件变量。

	*sync.Cond 类型的方法集合有 3 个方法，即：Wait、Signal 和 Broadcast。
	意思分别是：等待通知、单发通知和广播通知的操作。

	问题：条件变量是如何与读写锁 fmutex 的读锁关联的？

	条件变量 rcond 是与读写锁 fmutex 的读锁关联的。这是怎样做到的呢？
	读写锁的 RLocker 方法，它会返回当前读写锁中的读锁。该读锁同时也是 sync.Locker 接口的实现。
	因此，可以把它(df.fmutex.RLocker())作为参数值传给 sync.NewCond 函数。

	Go 提供的互斥锁、读写锁和条件变量，都基本遵循了 POSIX 标准中描述的对应同步工具的行为规范。
	它们简单且高效，为复杂的类型提供并发安全的保证。一些情况下，它们比通道更加适用。（通道的应用场景是什么呢？）
	问题：通道的应用场景是什么？什么情况下，适合应用通道？
	在只需对一个或多个临界区进行保护时，使用锁往往使程序的性能损耗更小。

	使程序性能损耗更小的同步工具——原子操作。

	#5.3 原子操作

	原子操作，即执行过程中不能被中断的操作。
	在针对某个值的原子操作执行过程当中，CPU 绝不会再去执行其他对该值的操作，无论这些其他操作是否为原子操作。

	Go 提供的原子操作都是非侵入式的。它们由标准库代码包 sync/atomic 中的众多函数代表，
	通过调用这些函数对几种简单类型的值执行原子操作。
	这些类型包括 6 种：int32、int64、uint32、uint64、uintptr 和 unsafe.Pointer。
	这些函数提供的原子操作共有 5 种：增或减、比较并交换、载入、存储和交换。
	这些函数提供了不同的功能，且【适用的场景】也有所区别。

	#5.3.1 增或减

	原子增/减操作，即可实现被操作值得增大或减小。
	被操作值的类型只能使数值类型，即：int32、int64、uint32、uint64 和 uintptr 类型。

	例如，想原子地把一个 int32 类型的变量 i32 的值增大 3，可这样做：
		newi32 := atomic.AddInt32(&i32, 3)

	例如，原子地将 int64 类型的变量 i64 的值减小 3，可以：
		var i64 int64
		atomic.AddInt64(&i64, -3)

	例如，想原子地把 uint32 类型的变量 ui32 的值增加 NN（NN 代表了一个负整数），可以
		atomic.AddUint32(&ui32, ^uint32(-NN-1))

		对于 uint64 类型，调用表达式：
		atomic.AddUint64(&ui64, ^uint64(-NN-1))

	二进制补码的特性
	一个负整数的补码可通过对它按位（除了符号位之外）求反码并加一得到。
	另外，一个负整数可以由对它的绝对值减一并求补码后得到的数值的二进制形式表示。

	例如，如果 NN 是一个 int 类型的变量且其值为 -35，那么表达式
	 uint32(int32(NN)) 和 ^uint32(-NN-1) 的结果值都会是 11111111111111111111111111011101。

	 使用 ^uint32(-NN-1) 和 ^uint64(-NN-1) 来分别表示 uint32 类型和 uint64 类型的 NN。


	#5.3.2 比较并交换

	比较并交换即 "Compare And Swap", 简称 CAS。

	CAS 操作，它总是假设被操作之未曾改变（即与旧值相等），并一旦确认这个假设的真实性就立即进行值替换。

	以针对 int32 类型值的函数为例，函数名为 CompareAndSwapInt32。声明如下：
	func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)

	CAS 操作的优势：可在不创建互斥量和不形成临界区的情况，完成并发安全的值替换操作。这可大大减少同步对程序性能的损耗。
	CAS 操作的劣势：在被操作值被频繁变更的情况下，CAS 操作并不那么容易成功。
	有时不得不利用 for 循环进行多次尝试。代码片段如下：
	var value int32
	func addValue(delta int32) {
		for {
			v := value
			if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
				break
			}
		}
	}

	可以看到，为了保证 CAS 操作成功完成，仅在 CompareAndSwapInt32 函数大的结果值为 true 时
	才退出循环。这种做法与自旋锁的自选行为相似。addValue 函数会不断尝试原子地更新 value 值，
	直到这一操作成功为止。操作失败的缘由总是 value 的值已不与 v 的值相等了。如果 value 的值
	会被并发地修改，就很有可能发生这种情况。

	（value 的值已不与 v 的值相等，缘由是 value 的值在被并发地修改。）

	CAS 操作虽不会让 goroutine 阻塞，但仍可能使流程的执行暂时停滞。不过，这种停滞大都极其短暂。

	并发安全地更新一些类型的值，应总是有限选择 CAS 操作。
	进行原子的 CAS 操作的函数，共有 6 个。
	CompareAndSwapInt32、CompareAndSwapInt64、CompareAndSwapPointer、
	CompareAndSwapUint32、CompareAndSwapUint64、CompareAndSwapUintptr

	(代码包是 atomic.*)

	#5.3.3 载入

	前面 for 循环种，使用语句 v := value 为变量赋值。
	但是要注意，在读取 value 的过程，并不能保证没有对此值的并发读写操作。

	例子：在 32 位计算架构的计算机上写入一个 64 位的整数。如果在这个写操作完成前，
	有一个读操作被并发地进行了，这个读操作就可能会读取到一个只被修改了一半的数据。
	这种结果是相当糟糕的。（读取到一个只被修改了一半的数据）

	为了原子地读取某个值，sync/atomic 代码包同样提供了一系列的函数，函数的名称都以  "Load"（意为"载入"）为前缀。
	以针对 int32 类型值的函数为例：
	func addValue(delta int32) {
		for {
			v := atomic.LoadInt32(&value)
			if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
				break
			}
		}
	}

	载入函数: sync/atomic.LoadInt32\aotmic.LoadInt64\atomic.LoadPointer
	\atomic.LoadUint32\atomic.LoadUint64\atomic.LoadUintptr

	#5.3.4 存储（写入操作）

	与读取操作相对应的是写入操作。而 sync/atomic 包也提供了对应的存储函数，函数的名称均以 "Store" 为前缀。

	在原子地存储某个值的过程种，任何 CPU 都不会进行针对同一个值得读写操作。如果把所有【针对此值的写操作】都改为
	原子操作，就绝不会出现针对此值的读操作因被并发地进行，而读到【修改了一半的值】的情况。

	原子的值存储操作总会成功，因为它并不关心被操作值的旧值是什么。显然，这与前面讲到的 CAS 操作有着明显的区别。
	因此，不能把前面 addValue 函数种对 atomic.CompareAndSwapInt32 函数的调用替换为对 atomic.StoreInt32 函数的调用。

	函数 atomic.StoreInt32 会接受两个参数。第一个参数的类型是 *int32，指向被操作值的指针值。第二个参数则是 int32 类型的，
	它的值是欲存储的新值。其他同类函数也会有类似的参数声明列表。

	#5.3.5 交换

	在 sync/atomic 代码包中还存在着一类函数，它们的功能与前文所讲的 CAS（compare and swap） 操作和原子载入（load）都有相似之处。
	这里的功能可以称为”原子交换操作“，这类函数的名称都以 "Swap" 为前缀。

	与 CAS  操作不同，原子交换操作不会关心被操作值的旧值，而是直接设置新值。但又比原子存储操作多做了一步：返回被操作值的旧值。
	此类操作比 CAS 操作的约束更少，同时又比原子载入操作的功能更强。

	atomic.SwapInt32 函数为例，第一个参数是代表了被操作值的内存地址的 *int32 类型值，而第二个参数用来表示新值。函数是有结果值的，该值
	即被新值替换掉的旧值。

	sync/atomic 代码包中所有函数的功能和用法，都是用来对特定类型的值进行原子性操作。如果想以并发安全的方式操作特定的简单类型值，应该首先
	考虑使用 sync/atomic 这些函数来实现。（以并发安全的方式操作....）
	原子地减小 uint32 类型或 uint64 类型的值，其实现方式并不那么直观。

	#5.3.6 原子值

	sync/atomic.Value 是一个结构体类型，暂且称为”原子值类型“，用于存储需要原子读写的值。
	不同于 sync/atomic 包中的其他函数，sync/atomic.Value 可接受被操作值的类型不限。

	声明一个可用的原子值实例，如：var atomicVal atomic.Value

	该类型有两个公开的指针方法——Load 和 Store。
	前者（Load）用于原子地读取【原子值实例】中存储的值，它会返回一个 interface{} 类型的结果且不接受任何参数。
	后者（Store）用于原子地在原子值实例中存储一个值，它接受一个 interface{} 类型的参数而没有任何结果。
	在未曾通过 Store 方法向原子值实例存储值之前，它的 Load 方法总会返回 nil。

	对于原子值实例的 Store 方法有两个限制。
	第一、作为参数传入该方法的值不能为 nil。
	第二、作为参数传入该方法的值必须与之前传入的值（如果有的话）的类型相同。
	一旦原子值实例存储了某一个类型的值，那之后存储的值就都必须是该类型的。
	如果违反了任一限制，对该方法的调用都会引发一个运行时恐慌。(panic)

	注意情况：
	sync/atomic.Value 类型的变量一旦声明，其值就不应该被赋值到它处。
	作为源值赋给别的变量、作为参数值传入函数、作为结果值从函数返回、作为元素值通过通道传递等，都会造成值的复制。
	（所以在这类变量之上不应该实施这些操作）
	不过，sync/atomic.Value 类型的指针类型的变量却不存在这个问题。
	根本原因是，对结构体值的复制不但会生成该值的副本，还会生成其中字段的副本。如此一来，不应施加于此的并发安全保护也就失效了。甚至，
	 向副本存储的操作都已与原值无关了。

	小总结： （源值-赋值变量、参数值-传入函数、结果值-函数返回、元素值-通道传递）

	（类型）对于 sync 包中的 Mutex、RWMutex 和 Cond 类型，go vet 命令同样会检查此类复制问题，其原因也是相似的。（什么原因呢？）
	一个比较彻底的解决方案是，避免直接使用它们，而使用它们的指针值。（指针值）

	原子值类型的经典使用场景：
	由于对原子值的读写操作必是原子的，同时又不受操作值类型的限制，因此它比原子函数的使用场景大很多。
	有些时候，它可以完美替换锁。

	原子值、原子函数



*/

/*
	数组与切片的区别

	primes := [6]int{2, 3, 5, 7, 11, 13} // array

	var s []int = primes[1:4] // slice
*/

/*
	测试竞态条件： go test -race
*/
