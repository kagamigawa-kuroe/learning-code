# STL

#### vector

1. vector的初始化：可以有五种方式,举例说明如下：

   （1） vector<int> a(10); //定义了10个整型元素的向量（尖括号中为元素类型名，它可以是任何合法的数据类型），但没有给出初值，其值是不确定的。
   （2）vector<int>a(10,1); //定义了10个整型元素的向量,且给出每个元素的初值为1
   （3）vector<int>a(b); //用b向量来创建a向量，整体复制性赋值
   （4）vector<int>a(b.begin(),b.begin+3); //定义了a值为b中第0个到第2个（共3个）元素
   （5）intb[7]={1,2,3,4,5,9,8};vector<int> a(b,b+7); //从数组中获得初值

2. vector对象的几个重要操作，举例说明如下：

   （1）a.assign(b.begin(), b.begin()+3); //b为向量，将b的0~2个元素构成的向量赋给a，会清除掉vector容器中以前的内容
   （2）a.assign(4,2);//是a只含4个元素，且每个元素为2
   （3）a.back(); //返回a的最后一个元素
   （4）a.front(); //返回a的第一个元素
   （5）a[i]; //返回a的第i个元素，当且仅当a[i]存在
   （6）a.clear();//清空a中的元素
   （7）a.empty();//判断a是否为空，空则返回ture,不空则返回false
   （8）a.pop_back(); //删除a向量的最后一个元素
   （9）a.erase(a.begin()+1,a.begin()+3);//删除a中第1个（从第0个算起）到第2个元素，也就是说删除的元素从a.begin()+1算起（包括它）一直到a.begin()+3（不包括它）
   （10）a.push_back(5);//在a的最后一个向量后插入一个元素，其值为5
   （11）a.insert(a.begin()+1,5);//在a的第1个元素（从第0个算起）的位置插入数值5，如a为1,2,3,4，插入元素后为1,5,2,3,4
   （12）a.insert(a.begin()+1,3,5);//在a的第1个元素（从第0个算起）的位置插入3个数，其值都为5
   （13）a.insert(a.begin()+1,b+3,b+6);//b为数组，在a的第1个元素（从第0个算起）的位置插入b的第3个元素到第5个元素（不包括b+6），如b为1,2,3,4,5,9,8，插入元素后为1,4,5,9,2,3,4,5,9,8
   （14）a.size();//返回a中元素的个数；
   （15）a.capacity();//返回a在内存中总共可以容纳的元素个数
   （16）a.rezize(10);//将a的现有元素个数调至10个，多则删，少则补，其值随机
   （17）a.rezize(10,2);//将a的现有元素个数调至10个，多则删，少则补，其值为2
   （18）a.reserve(100);//将a的容量（capacity）扩充至100，也就是说现在测试a.capacity();的时候返回值是100.这种操作只有在需要给a添加大量数据的时候才 显得有意义，因为这将避免内存多次容量扩充操作（当a的容量不足时电脑会自动扩容，当然这必然降低性能） 
   （19）a.swap(b);//b为向量，将a中的元素和b中的元素进行整体性交换
   （20）a==b; //b为向量，向量的比较操作还有!=,>=,<=,>,<

---

#### stack

C++ stack（堆栈）实现了一个**先进后出**（FILO）的数据结构。

构造函数：

- `stack<T> stkT;` : 采用模板类实现，stack对象的默认构造形式
- `stack(const stack &stk);` : 拷贝构造函数

常用方法：

- `size()`: 返回栈中的元素数
- `top()`: 返回栈顶的元素
- `pop()`: 从栈中取出并删除元素
- `push(x)`: 向栈中添加元素x
- `empty()`: 在栈为空时返回true

---

#### set

Sets are containers that store unique elements following a specific order.

集合中以一种特定的顺序保存唯一的元素。

构造函数:

- `set();` 无参数 - 构造一个空的set
- `set(InputIterator first, InputIterator last)` : 迭代器的方式构造set
- `set(const set &from);` : copyd的方式构造一个与set from 相同的set
- `set(input_iterator start, input_iterator end);` 迭代器(start)和迭代器(end) - 构造一个初始值为[start,end)区间元素的Vector(注:半开区间).
- 

- `std::set<int>a{1, 2, 3, 4, 5};`

常用API：

- `begin()` : 返回指向第一个元素的迭代器
- `clear()` : 清除所有元素
- `count()` : 返回某个值元素的个数
- `empty()` : 如果集合为空，返回true
- `end()` : 返回指向最后一个元素的迭代器
- `equal_range()` : 返回集合中与给定值相等的上下限的两个迭代器
- `erase()` : 删除集合中的元素
- `find()` : 返回一个指向被查找到元素的迭代器
- `get_allocator()` : 返回集合的分配器
- `insert()` : 在集合中插入元素
- `lower_bound()` : 返回指向大于（或等于）某值的第一个元素的迭代器
- `key_comp()` : 返回一个用于元素间值比较的函数
- `max_size()` : 返回集合能容纳的元素的最大限值
- `rbegin()` : 返回指向集合中最后一个元素的反向迭代器
- `rend()` : 返回指向集合中第一个元素的反向迭代器
- `size()` : 集合中元素的数目
- `swap()` : 交换两个集合变量
- `upper_bound()` : 返回大于某个值元素的迭代器
- `value_comp()` : 返回一个用于比较元素间的值的函数

---

#### queue

C++队列是一种容器适配器，它给予程序员一种先进先出(FIFO)的数据结构。

构造函数：

- `std::queue<std::string> words;`
- `std::queue<std::string> copy_words {words};`
- `std::deque<int> mydeck (3,100);`
- `std::queue<int,std::list<int>> third;`

常用API：

- `back()` : 返回最后一个元素
- `empty()` : 如果队列空则返回真
- `front()` : 返回第一个元素
- `pop()` : 删除第一个元素
- `push()` : 在末尾加入一个元素
- `size()` : 返回队列中元素的个数

---

#### list

构造函数：

- `list<string> a;`
- `list<string> a {20};`
- `list<string> a(20,0);`
- `list (const list& x);`

常用API：

- `assign()` : 给list赋值
- `back()` : 返回最后一个元素
- `begin()` : 返回指向第一个元素的迭代器
- `clear()` : 删除所有元素
- `empty()` : 如果list是空的则返回true
- `end()` : 返回末尾的迭代器
- `erase()` : 删除一个元素
- `front()` : 返回第一个元素
- `get_allocator()` : 返回list的配置器
- `insert()` : 插入一个元素到list中
- `max_size()` : 返回list能容纳的最大元素数量
- `merge()` : 合并两个list
- `pop_back()` : 删除最后一个元素
- `pop_front()` : 删除第一个元素
- `push_back()` : 在list的末尾添加一个元素
- `push_front()` : 在list的头部添加一个元素
- `rbegin()` : 返回指向第一个元素的逆向迭代器
- `remove()` : 从list删除元素
- `remove_if()` : 按指定条件删除元素
- `rend()` : 指向list末尾的逆向迭代器
- `resize()` : 改变list的大小
- `reverse()` : 把list的元素倒转
- `size()` : 返回list中的元素个数
- `sort()` : 给list排序
- `splice()` : 合并两个list
- `swap()` : 交换两个list
- `unique()` : 删除list中重复的元素