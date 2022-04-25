## design-patterns

#### 1. abstract factory

---

**Abstract Factory** is a creational design pattern, it solves the problem about the complexity of different classes who have the same catalog of feature. 

For example, we have there class A, B and C, and for everyone of them, it has three variant,  such as Class A has three variants A1,A2 and A3, so now we have 3x3 9 different class, and with the increasement of class and variants, it will be diffcult to distinguish different classes. In this case, abstract factory will be useful.

The core idea of Abstract Factory is to provide with different variants with different Factory class, e.g. for class A1,B1 and C1, we has a Factory class Factory1. The same, we have Factory2 and Factory3, so when we need A1, we request a Factory1, and tell it we need a class A, it will give us back a Class A1, we needn't wo request a A1 directly.

Here is an example:

```C++
class A{
 public:
    virtual ~A(){};
    virtual void fun_A() = 0;
};

class B{
 public:
    virtual ~B(){};
    virtual void fun_B() = 0;
};
```

We define two abstract classes A and B.

Then we implement them.

```c++
class A1:public A{
 public:
    void fun_A(){cout<<"A1"<<endl;};
}

class A2:public A{
 public:
    void fun_A(){cout<<"A2"<<endl;};
}

class B1:public B{
 public:
    void fun_B(){cout<<"B1"<<endl;};
}

class B2:public B{
 public:
    void fun_B(){cout<<"B2"<<endl;};
}

```

Now we have four classes, then we design factories.

```c++
class AbstractFactory {
 public:
  virtual A *CreateProductA() const = 0;
  virtual B *CreateProductB() const = 0;
};

class ConcreteFactory1 : public AbstractFactory {
 public:
  A *CreateProductA() const override {
    return new A1();
  }
  B *CreateProductB() const override {
    return new B1();
  }
};

class ConcreteFactory2 : public AbstractFactory {
 public:
  A *CreateProductA() const override {
    return new A2();
  }
  B *CreateProductB() const override {
    return new B2();
  }
};
```

Now, we have two different kinds of factories.

```
//if we want a A1 or B1
ConcreteFactory1 concreteFactory1;
A1* pt1 = concreteFactory1.CreateProductA();
B1* pt2 = concreteFactory1.CreateProductB();

//if we want a A2 or B2
ConcreteFactory2 concreteFactory2;
A2* pt3 = concreteFactory2.CreateProductA();
B2* pt4 = concreteFactory2.CreateProductB();
```





