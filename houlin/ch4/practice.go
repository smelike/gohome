package main

/*
	实战：


	开发完成一个可运行的软件，并通过【基本的功能】测试之后，对软件进行【性能评测】。


		可运行的软件：

			1、到底能跑多快？

			2、高负载下，是否还能保证正确性。或是，载荷数量与软件正确性之间的关系是怎样的？
			载荷数量是一个笼统的词，可以是 HTTP 请求的数量，也可以是 API 调用的次数。

			3、保证一定正确性的条件下，该软件的可伸缩性是怎样的？

			4、正常工作的情况下，负载与系统资源（包括 CPU、内存和各种 I/O 资源等）使用率之间的关系怎样的？

		只有为这些问题找到了答案，我们才能够真正了解到软件的性能，也只有这样才会知道需要进行怎样的软件设计，
		以及提供怎样的系统资源，才能够让它在承受一定量的载荷的同时保证正确性。这也是分析和定位软件性能瓶颈所需的重要参考资料。
		其中，这个载荷的量是我们在对性能评测所得到的一系列数值进行统计和分析后得出的。

		通过对这些数值的掌握，我们也可以了解软件运行环境下的性能。
		另外，正确性的比率应该满足软件使用者（客户端软件的开发者或者终端用户）对软件的刚性需求。

		所谓刚性需求，就是关乎软件质量和使用者体验的硬性指标，是软件必须满足的需求。

*/
