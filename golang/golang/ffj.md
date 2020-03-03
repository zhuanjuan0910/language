## 方法集：影响接口实现规则
+  类型 T 方法集包含全部 receiver T 方法。
+ 类型 *T 方法集包含全部 receiver T + *T 方法。
+  如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。 
+  如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。 
+  不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。
###  用实例 value 和 pointer 调用方法 (含匿名字段) 不受方法集约束，编译器总是查找全部方法，并自动转换 receiver 实参。
## 方法表达式
+ instance.method
+ type.method