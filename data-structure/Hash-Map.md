# HashMap

在Hashmap前，可以先看看c++基础STL中的Map容器.

Map容器就是一个最基础的，提供键值一一映射的容器，map始终保证遍历的时候是按key的大小顺序的

底层的实现基于红黑树, 本质就是一颗二叉查找树 , 由于没有做特殊处理，查找的效率是log(n), 下面是一些常用的方法.

```c++
// 定义一个map对象
map<int, string> mapStudent;

// 定义+初始化
map<int, string> mapStudent2 = {{1,"123"},{2,"456"}};
 
// 第一种 用insert函數插入pair
mapStudent.insert(pair<int, string>(000, "student_zero"));
 
// 第二种 用insert函数插入value_type数据
mapStudent.insert(map<int, string>::value_type(001, "student_one"));
 
// 第三种 用"array"方式插入
mapStudent[123] = "student_first";
mapStudent[456] = "student_second";

// find 返回迭代器指向当前查找元素的位置否则返回map::end()位置
map<int, string> iter;
iter = mapStudent.find("123");
 
if(iter != mapStudent.end())
       cout<<"Find, the value is"<<iter->second<<endl;
else
   cout<<"Do not Find"<<endl;

// 其他常用方法
count()         返回指定元素出现的次数
empty()         如果map为空则返回true
end()           返回指向map末尾的迭代器
equal_range()   返回特殊条目的迭代器对
erase()         删除一个元素
find()          查找一个元素
get_allocator() 返回map的配置器
insert()        插入元素
key_comp()      返回比较元素key的函数
lower_bound()   返回键值>=给定元素的第一个位置
max_size()      返回可以容纳的最大元素个数
rbegin()        返回一个指向map尾部的逆向迭代器
rend()          返回一个指向map头部的逆向迭代器
size()          返回map中元素的个数
swap()           交换两个map
upper_bound()    返回键值>给定元素的第一个位置
value_comp()     返回比较元素value的函数
```

Hashmap在c++中的实现为unordered_map，在c++11中被正式引入.

他基于哈希表，数据插入和查找的时间复杂度很低，几乎是常数时间，而代价是消耗比较多的内存.

底层实现上，使用一个下标范围比较大的数组来存储元素，形成很多的桶.

也就是我们会有一个数组，数组中的每个位置都是一个链表或者红黑树，在一个key-value被插入unordered_map时，会进行一个hash函数，判断应被插入到数组的哪个位置，然后再被插入。(参考java，jkd1.8中，当数组某个下标位置的链表中的元素大于8个时，就会转换为红黑树)

查找时，也会同样先进行hash判断，然后到对应的链表/红黑树中去查找，所以他的效率更高，基本为一个常量级。

在来说这个hash函数(散列函数)，不难看出，他应该具有如下特性：

- 如果两个散列值是不相同的（根据同一函数），那么这两个散列值的原始输入也是不相同的
- 就算散列值相同，也不一定能说明这两个数相同

区别于Map的顺序排列，他是无序的。

在HashMap中有两个很重要的参数，容量(Capacity)和负载因子(Load factor)

简单的说，Capacity就是buckets的数目，Load factor就是buckets填满程度的最大比例.

unordered_map的基本操作和map类似，例如声明，插入，删除等.

```c++
find() 根据键值，查找某个元素，返回迭代器，如果没找到元素，则返回unordered_map.end()迭代器，指示没有该元素。
count() 查找无序map中元素为指定键值的元素的数量，因为无序map不允许重复的键值，因此如果能找到该元素，则返回1，否则返回0。
insert() 插入元素，如果是重复键值，注意，该插入将会被忽略。
erase() 通过指定键值或者迭代器，可以删除元素。
clear() 清空容器内容。
empty() 判断是否为空容器。
size() 返回容器大小。
```



