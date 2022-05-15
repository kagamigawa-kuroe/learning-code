## design-patterns(1)

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

#### 2. factory method 

---

**Factory Method** is a creational design pattern that provides an interface for creating objects in a superclass and realise it in subclasses, return differents type of classes.

For example, we have a super factory class A, and two sub factory classes A1,A2. 

Beside we have a super product class P, and two sub product classes P1,P2.

Factory A provide a interface ``P create()``, and in class A1,A2, we realize the interface, new and return P1 and P2 respectively.

It's not so difficult, here is the code.

``` c++
class Product{
  public:
    int a;
    void f1() = 0;
}

class Product1:public Product {
  public:
    void f1(){cout<<"product1"<,endl};
}

class Product2:public Product {
  public:
    void f1(){cout<<"product2"<,endl};
}

class Factory {
 public:
    product& create();
}

class Factory1:public Factory{
  public:
    product& create(){
        product1 * p = new A1();
        return *p;
    }
}

class Factory2:public Factory{
  public:
    product& create(){
        product1 * p = new A2();
        return *p;
    }
}
```

#### 3. Builder

---

Builder Pattern is a design pattern who can help you generate an instance of the same class with different attributions.

Such as we have a class P with 3 attributions a1, a2, a3, for every  attribution it has several option. Now we want to set the value of those attribution by a new class, **Builder**.

In different variants of class builder, we implement interface of attribution value setter function, for different classes we want to create.

Here is a example:

```c++
class P {
 public:
  int a;
  int b;
}

class abstractbuilderofP {
  public:
    P* p;
    virtual void set_a() = 0;
    virtual void set_b() = 0;
    virtual P& make() = 0;
}

class builder1:public abstractbuilderofP{
  builder1(){this->P = new P()};
  void set_a(){p->a = 1};
  void set_b(){p->b = 1};
  P& make(){this->set_a();this->set_b();return *p;}
}

class builder2:public abstractbuilderofP{
  builder2(){this->P = new P()};
  void set_a(){p->a = 2};
  void set_b(){p->b = 2};
  P& make(){this->set_a();this->set_b();return *p;}
}
```

So now, if we want an instance of class P with value 1 and 1, we request it by builder1 and for 2 and 2, obviously by builder2.

if you want to do encapsulation of the create function, make add new class director and pass builder as a parameter or a attribution of it, it will make sense.







