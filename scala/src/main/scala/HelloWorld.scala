object HelloWorld {
  var num: Int = 1
  var str: String = "start"
  var short: Short = _ // initial value

  var NUM: Int = 1
  val STR: String = "2"
  var array: Array[String] = Array("a", "b", "c")

  def main(args: Array[String]) = {
    println("Hello World")

    println(num)
    println(str)
    println(short)

    println(NUM)
    println(STR)

    if (true) println(STR) else println("else")

    for (s <- array) println(s)

    val one = 1
    one match {
      case 1 => println("one")
      case 2 => println("two")
      case _ => println("other")
    }

    println(append_ten(NUM))
    println(append_eight(NUM))

    val user = new User("Tom", append_eight(3))
    println(user.name)
    println(user.age)
    user.how_old
    user.agree
  }

  def append_ten(num: Int): Int = {
    num + 10
  }

  def append_eight(num: Int): Int = num + 8

  class User(_name: String, _age: Int) {
    val name = _name
    val age = _age

    def agree = println("I'm " + name)

    def how_old = println("I'm " + fudge_the_count)

    private def fudge_the_count: Int = age - 5
  }

}
