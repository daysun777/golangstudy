#### go语言是静态类型语言。	
编译时类型已经确定，比如对已基本数据类型的再定义后的类型，反射时候需要确认返回的是何种类型。
#### 空接口interface{}	
go的反射机制是要通过接口来进行的，而类似于Java的Object的空接口可以和任何类型进行交互，因此对基本数据类型等的反射也直接利用了这一特点

#### 反射定律
go语言使用时，有三条定律，简单整理如下  
No.1	反射可以将“接口类型变量”转换为“反射类型对象”。	Reflection goes from interface value to reflection object.
No.2	反射可以将“反射类型对象”转换为“接口类型变量”。	Reflection goes from reflection object to interface value.
No.3	如果要修改“反射类型对象”，其值必须是“可写的”。	To modify a reflection object, the value must be settable.
